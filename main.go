package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

var ignoreDirectory = map[string]bool{".DS_Store": false, "dockerfile": false, "hw1.md": false, ".dockerignore": false, ".idea": true, ".git": true, ".gitignore": false}

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

func dirTree(out io.Writer, path string, files bool) error {
	if path == "-f" {
		path = "."
	}
	return walkFun(out, path, files, "")
}

func walkFun(out io.Writer, path string, printFiles bool, prefix string) error {
	directoryList, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	if !printFiles {
		directoryList = filterDirectory(directoryList)
	}
	sort.Slice(directoryList, func(i, j int) bool {
		return directoryList[i].Name() < directoryList[j].Name()
	})
	for position, currentDirectory := range directoryList {
		isLast := len(directoryList)-1 == position
		if isIgnoreDirectory(currentDirectory.Name()) {
			continue
		}
		var afterPrefix string
		var newPrefix string
		if isLast {
			afterPrefix = prefix + "\u2514\u2500\u2500\u2500"
			newPrefix = prefix + "\t"
		} else {
			afterPrefix = prefix + "\u251C\u2500\u2500\u2500"
			newPrefix = prefix + "\u2502\t"
		}
		afterPrefix += currentDirectory.Name()
		if currentDirectory.IsDir() {
			if _, err := fmt.Fprintln(out, afterPrefix); err != nil {
				return err
			}
			if err := walkFun(out, path+string(os.PathSeparator)+currentDirectory.Name(), printFiles, newPrefix); err != nil {
				return err
			}
		} else {
			if printFiles {
				if currentDirectory.Size() == 0 {
					afterPrefix = afterPrefix + " (empty)"
				} else {
					afterPrefix = afterPrefix + " (" + strconv.FormatInt(currentDirectory.Size(), 10) + "b)"
				}
				_, err := fmt.Fprintln(out, afterPrefix)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func isIgnoreDirectory(directoryName string) bool {
	_, isExist := ignoreDirectory[directoryName]
	return isExist
}

func filterDirectory(directories []os.FileInfo) []os.FileInfo {
	filteredDirectories := make([]os.FileInfo, 0)
	for _, currentDirectory := range directories {
		if currentDirectory.IsDir() {
			filteredDirectories = append(filteredDirectories, currentDirectory)
		}
	}
	return filteredDirectories
}
