package constants

var ALLOWED_TYPES = map[string][]string{
	"images":   {".jpg", ".jpeg", ".png", ".gif", ".svg", ".bmp", ".webp", ".psd", ".ico", ".heic", ".raw", ".ai"},
	"audio":    {".mp3", ".wav", ".flac", ".aac", ".ogg", ".m4a", ".wma", ".alac", ".aiff", ".pcm"},
	"video":    {".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv", ".webm", ".m4v", ".mpeg", ".3gp", ".ogg"},
	"document": {".doc", ".docx", ".pdf", ".txt", ".rtf", ".odt", ".ppt", ".pptx", ".xls", ".xlsx", ".csv", ".html", ".xml", ".md", ".epub", ".pages"},
}
