package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const IMAGE_TYPE = 0
const AUDIO_TYPE = 1
const VIDEO_TYPE = 2
const DOCUMENT_TYPE = 3
const COMPRESSED_TYPE = 4
const UNDEFINED_TYPE = -1

var ALLOWED_TYPES = map[int][]string{
	IMAGE_TYPE:    {".jpg", ".jpeg", ".png", ".gif", ".svg", ".bmp", ".webp", ".psd", ".ico", ".heic", ".raw", ".ai"},
	AUDIO_TYPE:    {".mp3", ".wav", ".flac", ".aac", ".ogg", ".m4a", ".wma", ".alac", ".aiff", ".pcm"},
	VIDEO_TYPE:    {".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv", ".webm", ".m4v", ".mpeg", ".3gp", ".ogg"},
	DOCUMENT_TYPE: {".doc", ".docx", ".pdf", ".txt", ".rtf", ".odt", ".ppt", ".pptx", ".xls", ".xlsx", ".csv", ".html", ".xml", ".md", ".epub", ".pages"},
}

func main() {
	dir := os.Args[1]

	files, err := os.ReadDir(dir)
	var filesToOrder []string

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filesToOrder = append(filesToOrder, file.Name())
	}

	for _, fileName := range filesToOrder {
		fileType := determineType(fileName)

		switch fileType {
		case IMAGE_TYPE:
			_ = moveToFolder(filepath.Join(dir, fileName), filepath.Join(dir, "images"))
		case AUDIO_TYPE:
			_ = moveToFolder(filepath.Join(dir, fileName), filepath.Join(dir, "audio"))
		case VIDEO_TYPE:
			_ = moveToFolder(filepath.Join(dir, fileName), filepath.Join(dir, "videos"))
		case DOCUMENT_TYPE:
			_ = moveToFolder(filepath.Join(dir, fileName), filepath.Join(dir, "documents"))
		default:
			fmt.Printf("Unsupported file type: %s\n", fileName)
		}
	}
}

func determineType(file string) int {
	for key, value := range ALLOWED_TYPES {
		for _, ext := range value {
			if strings.HasSuffix(file, ext) {
				return key
			}
		}
	}
	return UNDEFINED_TYPE
}

func moveToFolder(src, dest string) error {
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		err := os.Mkdir(dest, 0777)
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
