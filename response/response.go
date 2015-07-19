// This is the general response that will
// encompass all other responses
package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	StatusText string      `json:"status_text"`
	ErrorCode  int         `json:"error_code"`
	ErrorText  string      `json:"error_text"`
	Content    interface{} `json:"content"`
}

func New() *Response {
	resp := new(Response)
	resp.StatusCode = http.StatusOK
	resp.ErrorCode = ErrorNoError
	return resp
}

func (resp *Response) Output(writer http.ResponseWriter) {
	// Determining Unknowns
	resp.Success = resp.StatusCode == http.StatusOK &&
		resp.ErrorCode == ErrorNoError
	resp.StatusText = http.StatusText(resp.StatusCode)
	resp.ErrorText = GetErrorText(resp.ErrorCode)

	// Writing to the ResponseWriter
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(resp.StatusCode)
	fmt.Fprint(writer, toJSON(resp))
}

func toJSON(i interface{}) string {
	json, err := json.Marshal(i)
	if err != nil {
		return "Error converting response to JSON"
	}
	return string(json)
}
