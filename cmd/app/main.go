package main

import (
	"encoding/json"
	"extension-pre-release/cli"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func updateValueFromPath(path string, file *map[string]interface{}, value string) {
	parts := strings.Split(path, ".")

	var paths []string

	paths = append(paths, parts...)

	current := *file

	for i := 0; i < len(paths)-1; i++ {
		current = current[paths[i]].(map[string]interface{})
	}

	current[paths[len(paths)-1]] = value
}

func getValueFromPath(path string, file map[string]interface{}) string {
	parts := strings.Split(path, ".")

	var paths []string

	paths = append(paths, parts...)

	current := file

	for i := 0; i < len(paths)-1; i++ {
		current = current[paths[i]].(map[string]interface{})
	}

	fromPathVersion, ok := current[paths[len(paths)-1]].(string)
	if !ok {
		log.Fatalf("Version is not a string")
	}

	return fromPathVersion
}

func main() {
	filePath := flag.String("file", "", "path to file")
	versionPath := flag.String("version_path", "version", "path to  in json like 'example.version'")

	flag.Parse()

	if *filePath == "" {
		log.Fatalf("No file path provided")
		return
	}

	rawFile, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer rawFile.Close()

	bytes, err := io.ReadAll(rawFile)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var file map[string]interface{}
	if err := json.Unmarshal(bytes, &file); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	var version string

	if *versionPath != "version" {
		fromPathVersion := getValueFromPath(*versionPath, file)
		version = fromPathVersion
	} else {
		defaultVersion, ok := file["version"].(string)
		if !ok {
			log.Fatalf("Version is not a string")
		}

		version = defaultVersion
	}

	optionIndex := cli.GetOption()

	if optionIndex == -1 {
		fmt.Println("No option selected, exiting...")
		return
	}

	parts := strings.Split(version, ".")

	var nums []int
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Error converting %s to int: %v\n", part, err)
			return
		}
		nums = append(nums, num)
	}

	nums[optionIndex]++

	for i := optionIndex + 1; i < len(nums); i++ {
		nums[i] = 0
	}

	parts = []string{}

	for _, num := range nums {
		parts = append(parts, strconv.Itoa(num))
	}

	newVersion := strings.Join(parts, ".")
	updateValueFromPath(*versionPath, &file, newVersion)

	updatedData, err := json.MarshalIndent(file, "", "  ")
	if err != nil {
		fmt.Printf("Error generating JSON: %v\n", err)
		return
	}

	err = os.WriteFile(*filePath, updatedData, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	green := "\033[32m"
	reset := "\033[0m"

	err = exec.Command("git", "add", ".").Run()
	if err != nil {
		fmt.Printf("Error adding files: %v\n", err)
	}

	fmt.Println(green + "Version updated successfully!" + reset)
}
