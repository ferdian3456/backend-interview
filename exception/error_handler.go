package exception

import (
	"backend-interview/helper"
	"backend-interview/model"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := model.WebResponse{
		Data: nil,
		Status: []model.Status{
			{
				Id:   1,
				Name: "FAILED",
			},
		},
	}

	helper.WriteToResponseBody(writer, webResponse)
}
