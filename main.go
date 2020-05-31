package main

import (
	//"io"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var ignoreDirectory = map[string]bool{".dockerignore": false, ".idea": true, ".git": true, ".gitignore": false}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out *os.File, path string, files bool) error {
	var startPath string
	if path == "-f" {
		startPath = "."
	} else {
		startPath = "./" + path
	}
	err := walkFun(startPath, files, 0)
	return err
}

func walkFun(path string, printFiles bool, nestingLevel int) error {
	directoryList, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, currentDirectory := range directoryList {
		if isIgnoreDirectory(currentDirectory.Name()) {
			if currentDirectory.IsDir() {
				fmt.Println(tabCounter(nestingLevel) + currentDirectory.Name())
				walkFun(path+"/"+currentDirectory.Name(), printFiles, nestingLevel+1)
			} else {
				if printFiles {
					fmt.Println(tabCounter(nestingLevel) + currentDirectory.Name())
				}
			}
		}
	}
	return nil
}

func tabCounter(count int) string {
	var result []string
	for position := 1; position <= count; position++ {
		result = append(result, "\t")
	}
	return strings.Join(result, "")
}

func isIgnoreDirectory(directoryName string) bool {
	_, isExist := ignoreDirectory[directoryName]
	return !isExist
}
