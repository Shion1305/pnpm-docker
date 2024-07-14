package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/docker/reference"
	"github.com/containers/image/v5/types"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dockerRefString := "docker://docker.io/library/node"
	tags, err := getRepositoryTags(dockerRefString)
	if err != nil {
		fmt.Println("Error getting tags:", err)
		return
	}
	fmt.Println(tags)
	fmt.Println(len(tags))
	// filter major version lower than 12
	rgMajor := regexp.MustCompile(`^(\d+)`)

	targetTags := make([]string, 0)
	for _, tag := range tags {
		if rgMajor.MatchString(tag) {
			majorVer, _ := strconv.Atoi(rgMajor.FindString(tag))
			if majorVer < 12 {
				continue
			}
		}
		targetTags = append(targetTags, tag)
	}
	fmt.Println(targetTags)
	fmt.Println(len(targetTags))
}

func getRepositoryTags(repoName string) ([]string, error) {
	ctx := context.Background()
	ref, err := parseDockerRepositoryReference(repoName)
	if err != nil {
		return nil, err
	}
	tags, err := docker.GetRepositoryTags(ctx, nil, ref)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func parseDockerRepositoryReference(refString string) (types.ImageReference, error) {
	dockerRefString, ok := strings.CutPrefix(refString, docker.Transport.Name()+"://")
	if !ok {
		return nil, fmt.Errorf("docker: image reference %s does not start with %s://", refString, docker.Transport.Name())
	}

	ref, err := reference.ParseNormalizedNamed(dockerRefString)
	if err != nil {
		return nil, err
	}

	if !reference.IsNameOnly(ref) {
		return nil, errors.New(`No tag or digest allowed in reference`)
	}

	// Checks ok, now return a reference. This is a hack because the tag listing code expects a full image reference even though the tag is ignored
	return docker.NewReference(reference.TagNameOnly(ref))
}
