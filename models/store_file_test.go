package models

import (
	"FileOrganizer/util"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func generateExtensions(n int) []string {
	extensions := make([]string, n)

	for i := 0; i < n; i++ {
		extensions[i] = util.RandomExt()
	}

	return extensions
}

func newFileCategory() *FileCategory {
	fileType := util.RandomFileType()
	folderName := util.RandomFolderName()
	extensions := generateExtensions(4)

	return &FileCategory{
		fileType,
		extensions,
		folderName,
	}
}

func addCategory(t *testing.T) *FileCategory {
	dataStore := NewFileDataStore("")

	fileCategory := newFileCategory()

	err := dataStore.AddCategory(fileCategory.FileType, fileCategory.FolderName, fileCategory.Extensions)
	require.NoError(t, err)

	categories, err := dataStore.GetCategories()
	lastCategory := categories[len(categories)-1]

	require.Equal(t, fileCategory, lastCategory)

	return fileCategory
}

func TestGetCategories(t *testing.T) {
	dataStore := NewFileDataStore("")

	categories, err := dataStore.GetCategories()
	require.NoError(t, err)
	require.NotEmpty(t, categories)
}

func TestAddCategory(t *testing.T) {
	addCategory(t)
}

func TestGetType(t *testing.T) {
	fileCategory := addCategory(t)
	dataStore := NewFileDataStore("")

	fileName := fmt.Sprintf("%v.%v", util.RandomString(5), fileCategory.Extensions[0])

	fetchedFileCategory, err := dataStore.GetType(fileName)
	require.NoError(t, err)
	require.Equal(t, fetchedFileCategory, fileCategory)

	fileName = fmt.Sprintf("%v.%v", util.RandomString(5), util.RandomExt())

	fetchedFileCategory, err = dataStore.GetType(fileName)
	require.Error(t, err)
	require.Nil(t, fetchedFileCategory)
}

func TestRemoveCategory(t *testing.T) {
	fileCategory := addCategory(t)
	dataStore := NewFileDataStore("")

	err := dataStore.RemoveCategory(fileCategory.FileType)
	require.NoError(t, err)

	_, err = dataStore.GetType(fmt.Sprintf("%v.%v", util.RandomString(5), fileCategory.Extensions[0]))
	require.Error(t, err)
}

func TestAddExtensionsToCategory(t *testing.T) {
	fileCategory := addCategory(t)
	dataStore := NewFileDataStore("")
	extensions := generateExtensions(2)

	err := dataStore.AddExtensionsToCategory(fileCategory.FileType, extensions)
	require.NoError(t, err)

	for _, ext := range extensions {
		fetchedFileCategory, err := dataStore.GetType(fmt.Sprintf("%v.%v", util.RandomString(5), ext))
		require.NoError(t, err)

		require.Equal(t, fileCategory.FileType, fetchedFileCategory.FileType)
		require.Equal(t, fileCategory.FolderName, fetchedFileCategory.FolderName)
	}
}

func TestRemoveExtensionsToCategory(t *testing.T) {
	fileCategory := addCategory(t)
	dataStore := NewFileDataStore("")
	extensions := generateExtensions(2)

	err := dataStore.RemoveExtensionsFromCategory(fileCategory.FileType, extensions[:1])
	require.NoError(t, err)

	fetchedFileCategory, err := dataStore.GetType(fmt.Sprintf("%v.%v", util.RandomString(5), extensions[0]))
	require.Error(t, err)
	require.Nil(t, fetchedFileCategory)
}

func TestSetCategoryFolder(t *testing.T) {
	fileCategory := addCategory(t)
	dataStore := NewFileDataStore("")
	newFolderName := util.RandomFolderName()

	err := dataStore.SetCategoryFolder(fileCategory.FileType, newFolderName)
	require.NoError(t, err)

	fetchedFileCategory, err := dataStore.GetType(fmt.Sprintf("%v.%v", util.RandomString(5), fileCategory.Extensions[0]))
	require.NoError(t, err)
	require.Equal(t, fetchedFileCategory.FileType, fileCategory.FileType)
	require.Equal(t, fetchedFileCategory.FolderName, newFolderName)
}
