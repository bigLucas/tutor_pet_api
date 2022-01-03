package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func main() {
	const (
		RESOURCES_FOLDER    = "resources/"
		MAIN_TEMPLATE_NAME  = "base.yaml"
		FINAL_TEMPLATE_NAME = "template.yaml"
		RESOURCES_PREFIX    = "Resources:"
	)

	// get content from resources
	content, err := getContentFromResources(RESOURCES_FOLDER, RESOURCES_PREFIX)
	if err != nil {
		fmt.Println(err)
		return
	}

	// read the main template
	template, err := os.ReadFile(MAIN_TEMPLATE_NAME)
	if err != nil {
		fmt.Println(err)
		return
	}

	// write the final template
	final_content := append(template, []byte(content)...)
	err = os.WriteFile(FINAL_TEMPLATE_NAME, final_content, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File successful created!")
}

func getContentFromResources(folder string, prefix string) (string, error) {
	content := ""
	resources_list, err := os.ReadDir(folder)
	if err != nil {
		return content, err
	}
	for _, resource := range resources_list {
		content += getContent(resource, folder, prefix)
	}
	return content, err
}

func getContent(resource fs.DirEntry, folder string, prefix string) string {
	if !resource.IsDir() {
		resource_content, err := os.ReadFile(folder + resource.Name())
		if err == nil {
			return strings.TrimPrefix(string(resource_content), prefix)
		}
	}
	return ""
}
