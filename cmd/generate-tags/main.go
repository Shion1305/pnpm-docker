package main

import (
	"github.com/Shion1305/pnpm-docker/pkg/docker"
	"github.com/Shion1305/pnpm-docker/pkg/template"
	"regexp"
)

func main() {
	digests, err := docker.ListImageDigests("library/node")
	if err != nil {
		return
	}

	for k, v := range digests {
		//	if k does not contain numbers
		rg := regexp.MustCompile(`\d+`)
		if rg.MatchString(k) {
			continue
		}
		template.GenDockerfile(k, v)
	}
}
