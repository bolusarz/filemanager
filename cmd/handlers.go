package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
)

func DisplayCategories(verbose bool) {
	categories, err := store.GetCategories()

	if err != nil {
		fmt.Println(err)
	}

	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.Debug)

	defer w.Flush()

	if verbose {
		_, _ = fmt.Fprintln(w, "Category\t Folder Name\t Extensions\t")
		for _, category := range categories {
			_, _ = fmt.Fprintln(
				w,
				fmt.Sprintf("%s\t %s\t %s\t",
					category.FileType,
					category.FolderName,
					strings.Join(category.Extensions, ","),
				),
			)
		}
	} else {
		_, _ = fmt.Fprintln(w, "Category\t Folder Name\t")
		for _, category := range categories {
			_, _ = fmt.Fprintln(w, fmt.Sprintf("%s\t %s\t", category.FileType, category.FolderName))
		}
	}
}

func addCategory(fileType, folderName string, extensions []string) {
	if fileType == "" {
		fmt.Println(errors.New("category name not supplied"))
		return
	}

	var extensionsToAdd []string
	var extensionsToRemove []string

	for _, extension := range extensions {
		if strings.HasPrefix(extension, "-") {
			rawExtension := strings.ReplaceAll(extension, "-", "")
			if strings.HasPrefix(rawExtension, ".") {
				extensionsToRemove = append(extensionsToRemove, rawExtension)
			} else {
				extensionsToRemove = append(extensionsToRemove, "."+extension)
			}
		} else {
			if strings.HasPrefix(extension, ".") {
				extensionsToAdd = append(extensionsToAdd, extension)
			} else {
				extensionsToAdd = append(extensionsToAdd, "."+extension)
			}
		}
	}

	err := store.AddCategory(fileType, folderName, extensionsToAdd)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = store.RemoveExtensionsFromCategory(fileType, extensionsToRemove)

	if err != nil {
		fmt.Println(err)
		return
	}

	DisplayCategories(true)
}

func removeCategory(fileType string) {
	err := store.RemoveCategory(fileType)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully removed")
	DisplayCategories(false)
}

func organizeFiles(dir string) {
	files, err := os.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		category, err := store.GetType(fileName)

		if err != nil {
			fmt.Println(err)
			continue
		}

		err = category.MoveToFolder(filepath.Join(dir, fileName))

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("%s moved to %s\n", fileName, filepath.Join(dir, category.FolderName, fileName))
	}
}
