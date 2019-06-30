package message

func Create(status bool) map[string]interface{} {
	msg := make(map[string]interface{})
	msg["status"] = status
	return msg
}

func CreateWithData(status bool, data interface{}) map[string]interface{} {
	msg := make(map[string]interface{})
	msg["status"] = status
	msg["data"] = data
	return msg
}
