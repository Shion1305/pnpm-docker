package main

import (
	"fmt"
	"github.com/Shion1305/pnpm-docker/pkg/docker"
)

func main() {
	//targetTags := []string{"latest"}

	hash, err := docker.ListImageDigests("library/node")
	if err != nil {
		fmt.Println("Error getting image hash:", err)
		return
	}
	for k, v := range hash {
		fmt.Println(k, v)
	}
}
