package utils

func HandleResponse(message string, code int) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"status": code,
	}
}