package http_utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const (
	errorSendingJSONReplyFormat = "error replying with JSON: %s"
)

func SendJSONReplyOK(w http.ResponseWriter, replyContent interface{}) {
	SendJSONReplyWithStatus(w, http.StatusOK, replyContent)
}

func SendJSONReplyWithStatus(w http.ResponseWriter, status int, replyContent interface{}) {
	toSend, err := json.Marshal(replyContent)
	if err != nil {
		panic(fmt.Sprintf(errorSendingJSONReplyFormat, errors.WithStack(err)))
	}

	w.WriteHeader(status)
	_, err = w.Write(toSend)
	if err != nil {
		panic(fmt.Sprintf(errorSendingJSONReplyFormat, errors.WithStack(err)))
	}
}
