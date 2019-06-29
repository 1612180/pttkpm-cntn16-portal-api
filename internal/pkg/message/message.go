package message

func Create(status bool) (msg map[string]interface{}) {
	msg = make(map[string]interface{})
	msg["status"] = status
	return
}
