package models

import (
	"os"
	"path/filepath"
	"strings"
)

type DataStore interface {
	AddCategory(fileType, folderName string, extensions []string) error
	RemoveCategory(fileType string) error
	AddExtensionsToCategory(fileType string, extensions []string) error
	RemoveExtensionsFromCategory(fileType string, extensions []string) error
	SetCategoryFolder(fileType, folderName string) error
	GetCategories() ([]*FileCategory, error)
	GetType(fileName string) (*FileCategory, error)
}

type FileCategory struct {
	FileType   string   `json:"file_type"`
	Extensions []string `json:"extensions"`
	FolderName string   `json:"folder_name"`
}

func (c FileCategory) IsOfType(fileName string) bool {
	for _, ext := range c.Extensions {
		if strings.HasSuffix(fileName, ext) {
			return true
		}
	}
	return false
}

func (c FileCategory) MoveToFolder(src string) error {
	dest := filepath.Join(filepath.Dir(src), c.FolderName)

	if _, err := os.Stat(dest); os.IsNotExist(err) {
		err := os.Mkdir(dest, 0755)
		if err != nil {
			return err
		}
	}

	fileName := filepath.Base(src)

	err := os.Rename(src, filepath.Join(dest, fileName))

	if err != nil {
		return err
	}

	return nil
}
