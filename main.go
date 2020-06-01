package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

func dirTree(out io.Writer, path string, files bool) error {
	var startPath string
	if path == "-f" {
		startPath = "."
	} else {
		startPath = "." + string(os.PathSeparator) + path
	}
	err := walkFun(out, startPath, files, 0, "")
	return err
}

func walkFun(out io.Writer, path string, printFiles bool, nestingLevel int, prefix string) error {
	directoryList, err := ioutil.ReadDir(path)
	if !printFiles {
		directoryList = filterDirectory(directoryList)
	}

	if err != nil {
		return err
	}
	for position, currentDirectory := range directoryList {
		isLast := len(directoryList) == 0 || len(directoryList)-1 == position
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
		afterPrefix = afterPrefix + currentDirectory.Name()
		if currentDirectory.IsDir() {
			_, err := fmt.Fprintln(out, afterPrefix)
			if err != nil {
				return err
			}
			err = walkFun(out, path+string(os.PathSeparator)+currentDirectory.Name(), printFiles, nestingLevel+1, newPrefix)
			if err != nil {
				return err
			}
		} else {
			if printFiles {
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

//// ├
//println("\u251C")

//// ─
//println("\u2500")

//// └
//println("\u2514")

//// │
//println("\u2502")
