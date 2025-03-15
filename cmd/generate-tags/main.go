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
		if !TagCompatible(k) {
			continue
		}
		template.GenDockerfile(k, v)
	}
}

func TagCompatible(tag string) bool {
	rg := regexp.MustCompile(`^(\d{2})(-.+)?$`)
	m := rg.FindStringSubmatch(tag)
	if m == nil {
		return false
	}
	majorVer, _ := strconv.Atoi(m[1])
	if majorVer < 21 {
		return false
	}
	return true
}
