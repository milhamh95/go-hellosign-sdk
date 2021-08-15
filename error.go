package hellosign

// Error is an error response from hellosign
// To see list of error from hellosign,
// Please access: https://app.hellosign.com/api/reference#ErrorNames
type ErrorResponse struct {
	ErrorResponse ErrorDetail `json:"error"`
}

// ErrorDetail is an error detail from hellosign error response
type ErrorDetail struct {
	ErrorMessage string `json:"error_msg"`
	ErrorName    string `json:"error_name"`
}

func (e *ErrorResponse) Error() string {
	return e.ErrorResponse.ErrorName + ": " + e.ErrorResponse.ErrorMessage
}
