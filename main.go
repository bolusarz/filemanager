package main

import (
	"FileOrganizer/cmd"
	"os"
)

type Directories []os.DirEntry

func (entries Directories) filterFiles() *[]string {
	var files []string

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		files = append(files, entry.Name())
	}

	return &files
}

func main() {
	//dir := os.Args[1]
	//
	//var store models.DataStore
	//
	//store = models.New("")
	//
	//files, err := os.ReadDir(dir)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var filesToOrder = *Directories(files).filterFiles()
	//
	//for _, fileName := range filesToOrder {
	//	fileType, err := store.GetType(fileName)
	//
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	err = fileType.MoveToFolder(filepath.Join(dir, fileName))
	//
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Printf("%s moved to %s\n", fileName, filepath.Join(dir, fileType.FolderName, fileName))
	//}

	cmd.Execute()
}
