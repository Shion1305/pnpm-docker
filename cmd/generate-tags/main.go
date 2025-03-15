package main

import (
	"github.com/Shion1305/pnpm-docker/pkg/docker"
	"github.com/Shion1305/pnpm-docker/pkg/template"
	"regexp"
	"strconv"
)

func main() {
	digests, err := docker.ListImageDigests("library/node")
	if err != nil {
		return
	}

	for k, v := range digests {
		//	if k does not contain numbers
		if !tagCompatible(k) {
			continue
		}
		template.GenDockerfile(k, v)
	}
}

func tagCompatible(tag string) bool {
	rg := regexp.MustCompile(`^(\d+)`)
	if !rg.MatchString(tag) {
		return false
	}
	majorVer, _ := strconv.Atoi(rg.FindString(tag))
	if majorVer < 22 {
		return false
	}
	return true
}
