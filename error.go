package hellosign

import (
	"encoding/json"
	"net/http"
)

// Error is an error response from hellosign
// To see list of error from hellosign,
// Please access: https://app.hellosign.com/api/reference#ErrorNames
type Error struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail is an error detail from hellosign error response
type ErrorDetail struct {
	ErrorMessage string `json:"error_msg"`
	ErrorName    string `json:"error_name"`
}

func prepareError(r *http.Response) (string, error) {
	e := Error{}
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		return "", err
	}
	return e.Error.ErrorName + ": " + e.Error.ErrorMessage, nil
}
