package constant

var AllowedExtensions = map[string]bool{
	".mp4":  true,
	".mov":  true,
	".mkv":  true,
	".webm": true,
	".mp3":  true,
	".wav":  true,
	".aac":  true,
	".flac": true,
	".ogg":  true,
}

const MaxFileSize int64 = 1 << 30 // 1 GB
