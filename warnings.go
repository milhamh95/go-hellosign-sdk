package hellosign

// Warnings represent warning messages response
// https://app.hellosign.com/api/reference#get_account
type Warnings struct {
	WarningMessage string `json:"warning_msg"`
	WarningName    string `json:"warning_name"`
}
