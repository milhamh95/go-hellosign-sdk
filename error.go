package hellosign

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

func prepareError(r *http.Response, errResp interface{}) error {
	errRespJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.New("read response")
	}

	err = json.Unmarshal(errRespJSON, &errResp)
	if err != nil {
		return errors.New("unmarshal json")
	}

	return err
}
