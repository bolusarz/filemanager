package models

import (
	"FileOrganizer/constants"
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// TODO: Think about caching
const storeFileName = "store.json"

type FileDataStore struct {
	storagePath string
}

func New(storagePath string) *FileDataStore {
	storage := FileDataStore{storagePath: storagePath}
	storage.Init()
	return &storage
}

func (f FileDataStore) Init() {
	_, err := f.GetCategories()
	if err != nil {
		dataStore := make([]*FileCategory, len(constants.ALLOWED_TYPES))
		index := 0
		for key, value := range constants.ALLOWED_TYPES {
			dataStore[index] = &FileCategory{
				FileType:   key,
				Extensions: value,
				FolderName: key,
			}
			index++
		}
		err = saveToFile(dataStore, filepath.Join(f.storagePath, storeFileName))

		if err != nil {
			panic(err)
		}
	}
}

func (f FileDataStore) GetCategories() ([]*FileCategory, error) {
	file, err := os.Open(filepath.Join(f.storagePath, storeFileName))

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	jsonString := ""

	for scanner.Scan() {
		jsonString += scanner.Text()
	}

	var categories []*FileCategory
	err = json.Unmarshal([]byte(jsonString), &categories)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (f FileDataStore) GetType(fileName string) (FileCategory, error) {
	categories, err := f.GetCategories()

	if err != nil {
		return FileCategory{}, err
	}

	for _, category := range categories {
		if category.IsOfType(fileName) {
			return *category, nil
		}
	}
	return FileCategory{}, errors.New("category not found")
}

func (f FileDataStore) AddCategory(fileType string, folderName string, extensions []string) error {
	categories, err := f.GetCategories()

	if err != nil {
		return err
	}

	fileCategory := &FileCategory{
		FileType:   fileType,
		Extensions: extensions,
		FolderName: fileType,
	}

	if folderName != "" {
		fileCategory.FolderName = folderName
	}

	categories = append(categories, fileCategory)

	err = saveToFile(categories, filepath.Join(f.storagePath, storeFileName))

	if err != nil {
		return err
	}

	return nil
}

func (f FileDataStore) RemoveCategory(fileType string) error {
	categories, err := f.GetCategories()

	if err != nil {
		return err
	}

	var indexToRemove = -1

	for index, category := range categories {
		if category.FileType == fileType {
			indexToRemove = index
			break
		}
	}

	if indexToRemove == -1 {
		return errors.New("category not found")
	}

	categories = append(categories[:indexToRemove], categories[indexToRemove+1:]...)

	err = saveToFile(categories, filepath.Join(f.storagePath, storeFileName))

	if err != nil {
		return err
	}

	return nil
}

func (f FileDataStore) AddExtensionsToCategory(fileType string, extensions []string) error {
	categories, err := f.GetCategories()
	if err != nil {
		return err
	}

	for _, category := range categories {
		if category.FileType == fileType {
			category.Extensions = append(category.Extensions, extensions...)

			err = saveToFile(categories, filepath.Join(f.storagePath, storeFileName))

			if err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("category not found")

}

func (f FileDataStore) RemoveExtensionsFromCategory(fileType string, extensions []string) error {
	categories, err := f.GetCategories()
	if err != nil {
		return err
	}

	for index, category := range categories {
		if category.FileType == fileType {
			for extensionIdx, extension := range category.Extensions {
				for _, extensionToRemove := range extensions {
					if extension == extensionToRemove {
						category.Extensions = append(category.Extensions[:extensionIdx], categories[index].Extensions[extensionIdx+1:]...)
					}
				}
			}

			err = saveToFile(categories, filepath.Join(f.storagePath, storeFileName))

			if err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("category not found")
}

func (f FileDataStore) SetCategoryFolder(fileType, folderName string) error {
	categories, err := f.GetCategories()
	if err != nil {
		return err
	}

	for _, category := range categories {
		if category.FileType == fileType {
			category.FolderName = folderName

			err = saveToFile(categories, filepath.Join(f.storagePath, storeFileName))
			if err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("category not found")
}

func saveToFile(categories []*FileCategory, filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	data, err := json.Marshal(categories)

	if err != nil {
		return err
	}

	_, err = writer.Write(data)

	if err != nil {
		return err
	}

	return nil
}
