package cmd

import (
	"errors"
	"fmt"
	"os"
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

	err := store.AddCategory(fileType, folderName, extensions)

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
