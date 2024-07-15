package docker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const dockerHubAPI = "https://hub.docker.com/v2"

type (
	TagDigestMap map[string]string
	TagList      struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name    string `json:"name"`
			ImageID string `json:"image_id"`
			Images  []struct {
				Architecture string `json:"architecture"`
				Digest       string `json:"digest"`
			} `json:"images"`
		} `json:"results"`
	}
)

func ListImageDigests(repository string) (TagDigestMap, error) {
	repoParts := strings.Split(repository, "/")
	var repoOwner, repoName string
	if len(repoParts) == 2 {
		repoOwner = repoParts[0]
		repoName = repoParts[1]
	} else {
		repoOwner = "library"
		repoName = repoParts[0]
	}

	page := 1
	tagMap := make(TagDigestMap, 100)

	for {
		url := fmt.Sprintf("%s/repositories/%s/%s/tags/?page=%d&page_size=100", dockerHubAPI, repoOwner, repoName, page)
		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("failed to get tags: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to get tags: %s", resp.Status)
		}

		var tagList TagList
		if err := json.NewDecoder(resp.Body).Decode(&tagList); err != nil {
			return nil, fmt.Errorf("failed to decode response: %v", err)
		}
		for _, t := range tagList.Results {
			tagMap[t.Name] = t.Images[0].Digest
		}
		if tagList.Next == "" {
			break
		}
		page += 1
		fmt.Println(len(tagMap))
	}
	return tagMap, nil
}
