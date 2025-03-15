package template

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Shion1305/pnpm-docker/pkg/docker"
)

const templateFilePath = "./template/Dockerfile"

type TemplateData struct {
	NodeTag            string
	RequireWgetInstall bool
}

func GenDockerfile(nodeTag string, digest docker.DigestInfo) error {
	tmpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("error parsing template file: %v", err)
	}
	wgetInstall := strings.Contains(nodeTag, "slim")
	data := TemplateData{NodeTag: nodeTag, RequireWgetInstall: wgetInstall}
	outputDir := filepath.Join("./images", nodeTag)
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating output directory: %v", err)
	}
	outputFilePath := filepath.Join(outputDir, "Dockerfile")
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	err = tmpl.Execute(writer, data)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}
	writer.Flush()

	if digest.AMD64 != "" {
		if err := writeOutDigest(outputDir, "amd64", digest.AMD64); err != nil {
			return fmt.Errorf("error writing out digest for %s:%s, %v", nodeTag, "amd64", err)
		}
	}
	if digest.ARM64 != "" {
		if err := writeOutDigest(outputDir, "arm64", digest.ARM64); err != nil {
			return fmt.Errorf("error writing out digest for %s:%s, %v", nodeTag, "arm64", err)
		}
	}
	return nil
}

func writeOutDigest(outputDir, arch string, digest string) error {
	outputFilePath := filepath.Join(outputDir, fmt.Sprintf("node-digest-%s", arch))
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()
	if _, err := outputFile.WriteString(digest + "\n"); err != nil {
		return fmt.Errorf("error writing content to file: %v", err)
	}
	return nil
}
