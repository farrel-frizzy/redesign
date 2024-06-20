package helper

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(writer http.ResponseWriter, result any) {
	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(result)
	if err != nil {
		panic(err)
	}
}
