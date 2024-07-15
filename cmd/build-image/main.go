package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type TemplateData struct {
	NodeTag string
}

func main() {
	templateFilePath := "./template/Dockerfile"

	nodeTag := "22-alpine3.19"

	tmpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		fmt.Println("Error parsing template file:", err)
		return
	}

	// テンプレートにデータを埋め込む
	data := TemplateData{NodeTag: nodeTag}
	outputDir := filepath.Join(".", nodeTag)
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}
	outputFilePath := filepath.Join(outputDir, "Dockerfile")
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	err = tmpl.Execute(writer, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
	writer.Flush()

	outputFile.Seek(0, 0)
	scanner := bufio.NewScanner(outputFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading new Dockerfile:", err)
		return
	}

	cmd := exec.Command("docker", "build", "-f", outputFilePath, "-t", "shion/pnpm:"+nodeTag, outputDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error building Docker image:", err)
		return
	}
	fmt.Println("Docker image built successfully")
}
