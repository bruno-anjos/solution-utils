package http_utils

import (
	"encoding/json"
	"net/http"
)

func SendJSONReplyOK(w http.ResponseWriter, replyContent interface{}) {
	toSend, err := json.Marshal(replyContent)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(toSend)
	if err != nil {
		panic(err)
	}
}