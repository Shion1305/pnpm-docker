package template

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const templateFilePath = "./template/Dockerfile"

type TemplateData struct {
	NodeTag string
}

func GenDockerfile(nodeTag, digest string) error {
	tmpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("error parsing template file: %v", err)
	}

	data := TemplateData{NodeTag: nodeTag}
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

	// write digest to file
	digestFilePath := filepath.Join(outputDir, "node-digest")
	digestFile, err := os.Create(digestFilePath)
	if err != nil {
		return fmt.Errorf("error creating digest file: %v", err)
	}
	defer digestFile.Close()
	if _, err := digestFile.WriteString(digest); err != nil {
		return fmt.Errorf("error writing digest to file: %v", err)
	}
	return nil
}
