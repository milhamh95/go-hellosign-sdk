package hellosign

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
)

// AccountAPI is a service to get account
type AccountAPI service

// Account represent account response
type Account struct {
	Account  AccountDetail `json:"account"`
	Warnings []Warnings    `json:"warnings"`
}

// AccountDetail represent detail account response
type AccountDetail struct {
	AccountID       string        `json:"account_id"`
	EmailAddress    string        `json:"email_address"`
	IsLocked        bool          `json:"is_locked"`
	IsPaidHelloSign bool          `json:"is_paid_hs"`
	IsPaidHelloFax  bool          `json:"is_paid_hello_fax"`
	Quota           AccountQuotas `json:"quotas"`
	CallbackURL     string        `json:"callback_url"`
	RoleCode        string        `json:"role_code"`
}

// AccountQuotas represent account quota
type AccountQuotas struct {
	APISignatureRequestsLeft int `json:"api_signature_requests_left"`
	DocumentsLeft            int `json:"documents_left"`
	TemplatesLeft            int `json:"templates_left"`
}

const (
	// SubURLAccount is sub url path for account
	SubURLAccount = "/account"
)

// Get will return an account and its settings
// based on user api key
func (a *AccountAPI) Get() (Account, error) {
	resp, err := a.client.doRequest(
		a.client.BaseURL+SubURLAccount,
		http.MethodGet,
		&bytes.Buffer{},
		&multipart.Writer{})
	if err != nil {
		return Account{}, err
	}
	defer resp.Body.Close()

	account := Account{}
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		return Account{}, err
	}

	return account, nil
}

// Verify will check whether an HelloSign Account exists for the given email address.
// This method is restricted to paid API users.
func (a *AccountAPI) Verify(emailAddress string) (Account, error) {
	path := a.client.BaseURL + SubURLAccount + "/verify"

	var params bytes.Buffer
	writer := multipart.NewWriter(&params)

	emailAddressField, err := writer.CreateFormField("email_address")
	if err != nil {
		return Account{}, err
	}
	emailAddressField.Write([]byte(emailAddress))
	writer.Close()

	resp, err := a.client.doRequest(path, http.MethodPost, &params, writer)
	if err != nil {
		return Account{}, err
	}
	defer resp.Body.Close()

	account := Account{}
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		return Account{}, err
	}

	return account, nil
}

// Update will update account callback url
func (a *AccountAPI) Update(callbackURL string) (Account, error) {
	var params bytes.Buffer
	writer := multipart.NewWriter(&params)

	callbackURLField, err := writer.CreateFormField("callback_url")
	if err != nil {
		return Account{}, err
	}
	callbackURLField.Write([]byte(callbackURL))
	writer.Close()

	resp, err := a.client.doRequest(a.client.BaseURL+SubURLAccount, http.MethodPost, &params, writer)
	if err != nil {
		return Account{}, err
	}
	defer resp.Body.Close()

	account := Account{}
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		return Account{}, err
	}

	return account, nil
}

// Create will create a new hellosign account
func (a *AccountAPI) Create(emailAddress string) (Account, error) {
	path := a.client.BaseURL + SubURLAccount + "/create"

	var params bytes.Buffer
	writer := multipart.NewWriter(&params)

	emailAddressField, err := writer.CreateFormField("email_address")
	if err != nil {
		return Account{}, err
	}
	emailAddressField.Write([]byte(emailAddress))
	writer.Close()

	resp, err := a.client.doRequest(path, http.MethodPost, &params, writer)
	if err != nil {
		return Account{}, err
	}
	defer resp.Body.Close()

	account := Account{}
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		return Account{}, err
	}

	return account, nil
}
