package utilstts


func Contains(text string, list []string) bool {
	for _, content := range list {
		if content == text {
			return true
		}
	}
	return false
}

