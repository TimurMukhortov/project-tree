package main

import (
	//"io"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

func dirTree1(out *os.File, path string, files bool) error {
	folderStruct := Folder{
		nestingLevel: 0,
	}
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			nestingLevel := len(strings.Split(path, "/"))
			if info.IsDir() {
				folderStruct.folders = append(folderStruct.folders, Folder{
					name:         info.Name(),
					nestingLevel: nestingLevel,
				})
			} else {
				folderStruct.files = append(folderStruct.files, File{
					name:         info.Name(),
					nestingLevel: nestingLevel,
				})
			}

			fmt.Println("\t"+path, info.Size(), info.IsDir())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(fullPathToProject)
	//os.Chdir(fullPathToProject + "/" + path + "/")
	//fmt.Println(filepath.Abs("./"))
	//if false {
	//	return fmt.Errorf("Kek")
	//}
	return nil
}

func dirTree(out *os.File, path string, files bool) error {
	//te, _ := os.Lstat(".")
	//fmt.Println(te)
	walkFun(".", files, 0)
	//filesTest, err := ioutil.ReadDir(".")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, file := range filesTest {
	//	fmt.Println(file.Name(), file.IsDir())
	//}
	return nil
}

func walkFun(path string, file bool, nestingLevel int) {
	directoryList, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, directory := range directoryList {
		fmt.Println(directory.Name(), directory.IsDir())
		if directory.Name() != ".idea" && directory.IsDir() {
			walkFun(directory.Name(), file, nestingLevel+1)
		}
	}
}

func tabCounter(count int) string {
	var result []string
	for position := 1; position <= count; position++ {
		result = append(result, "\t")
	}
	return strings.Join(result, "")
}
