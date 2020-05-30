package main

import (
	//"io"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type File struct {
	name         string
	nestingLevel int
}

type Folder struct {
	name         string
	nestingLevel int
	folders      []Folder
	files        []File
}

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
	err := walkFun(".", files, 0)
	return err
}

func walkFun(path string, file bool, nestingLevel int) error {
	directoryList, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, currentDirectory := range directoryList {
		fmt.Println(tabCounter(nestingLevel)+currentDirectory.Name(), currentDirectory.IsDir())
		if currentDirectory.IsDir() {
			walkFun(path+"/"+currentDirectory.Name(), file, nestingLevel+1)
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
