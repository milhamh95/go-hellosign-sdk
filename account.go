package hellosign

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
)

// AccountAPI is a service to account API
type AccountAPI service

// Account represent account response
type Account struct {
	Account  AccountDetail `json:"account"`
	Warnings []Warnings    `json:"warnings,omitempty"`
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
func (a *AccountAPI) Get(ctx context.Context) (Account, error) {
	resp, err := a.client.doRequest(
		requestParam{
			ctx:    ctx,
			path:   a.client.BaseURL + SubURLAccount,
			method: http.MethodGet,
		},
	)
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
func (a *AccountAPI) Verify(ctx context.Context, emailAddress string) (Account, error) {
	path := a.client.BaseURL + SubURLAccount + "/verify"

	var payload bytes.Buffer
	writer := multipart.NewWriter(&payload)

	emailAddressField, err := writer.CreateFormField("email_address")
	if err != nil {
		return Account{}, err
	}
	emailAddressField.Write([]byte(emailAddress))
	writer.Close()

	resp, err := a.client.doRequest(requestParam{
		ctx:    ctx,
		path:   path,
		method: http.MethodPost,
		writer: writer,
	})
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
func (a *AccountAPI) Update(ctx context.Context, callbackURL string) (Account, error) {
	var payload bytes.Buffer
	writer := multipart.NewWriter(&payload)

	callbackURLField, err := writer.CreateFormField("callback_url")
	if err != nil {
		return Account{}, err
	}
	callbackURLField.Write([]byte(callbackURL))
	writer.Close()

	resp, err := a.client.doRequest(
		requestParam{
			ctx:    ctx,
			path:   a.client.BaseURL + SubURLAccount,
			method: http.MethodPost,
			writer: writer,
		},
	)
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
func (a *AccountAPI) Create(ctx context.Context, emailAddress string) (Account, error) {
	path := a.client.BaseURL + SubURLAccount + "/create"

	var payload bytes.Buffer
	writer := multipart.NewWriter(&payload)

	emailAddressField, err := writer.CreateFormField("email_address")
	if err != nil {
		return Account{}, err
	}
	emailAddressField.Write([]byte(emailAddress))
	writer.Close()

	resp, err := a.client.doRequest(requestParam{
		ctx:    ctx,
		path:   path,
		method: http.MethodPost,
		writer: writer,
	})
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
