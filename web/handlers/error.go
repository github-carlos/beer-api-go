package handlers

import (
	"encoding/json"
	"fmt"
)

func HandleError(message string) []byte {
	error := struct {
		Message string `json:"message"`
	}{
		message,
	}

	fmt.Println(message)
	response, err := json.Marshal(error)

	if err != nil {
		return []byte(err.Error())
	}

	return response

}
