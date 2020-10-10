package hellosign

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
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
	SubURLAccount = "/account/"
)

// Get will return an account and its settings
// based on user api key
func (a *AccountAPI) Get() (Account, error) {
	req, err := http.NewRequest(http.MethodGet, a.client.BaseURL+SubURLAccount, nil)
	if err != nil {
		return Account{}, err
	}
	req.SetBasicAuth(a.client.apiKey, "")

	resp, err := a.client.HTTPClient.Do(req)
	if err != nil {
		return Account{}, err
	}

	if resp.StatusCode >= http.StatusMultipleChoices {
		e := Error{}
		err = prepareError(resp, &e)
		if err != nil {
			return Account{}, err
		}

		msg := e.Error.ErrorName + ": " + e.Error.ErrorMessage
		return Account{}, errors.New(msg)
	}

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Account{}, err
	}

	accountDetail := Account{}
	err = json.Unmarshal(bodyResp, &accountDetail)
	if err != nil {
		return Account{}, err
	}

	return accountDetail, nil
}

// GetByID will return account and its setting based on account id
func (a *AccountAPI) GetByID(accountID string) (Account, error) {
	req, err := http.NewRequest(http.MethodGet, a.client.BaseURL+SubURLAccount+accountID, nil)
	if err != nil {
		return Account{}, err
	}
	req.SetBasicAuth(a.client.apiKey, "")

	resp, err := a.client.HTTPClient.Do(req)
	if err != nil {
		return Account{}, err
	}

	if resp.StatusCode >= http.StatusMultipleChoices {
		e := Error{}
		err = prepareError(resp, &e)
		if err != nil {
			return Account{}, err
		}

		msg := e.Error.ErrorName + ": " + e.Error.ErrorMessage
		return Account{}, errors.New(msg)
	}

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Account{}, err
	}

	accountDetail := Account{}
	err = json.Unmarshal(bodyResp, &accountDetail)
	if err != nil {
		return Account{}, err
	}

	return accountDetail, nil
}

// Verify will check whether an HelloSign Account exists for the given email address.
// This method is restricted to paid API users.
func (a *AccountAPI) Verify(emailAddress string) (Account, error) {
	path := a.client.BaseURL + SubURLAccount + "verify"

	var params bytes.Buffer
	writer := multipart.NewWriter(&params)

	emailAddressField, err := writer.CreateFormField("email_address")
	if err != nil {
		return Account{}, err
	}
	emailAddressField.Write([]byte(emailAddress))

	req, err := http.NewRequest(http.MethodPost, path, &params)
	if err != nil {
		return Account{}, err
	}
	req.SetBasicAuth(a.client.apiKey, "")

	resp, err := a.client.HTTPClient.Do(req)
	if err != nil {
		return Account{}, err
	}

	if resp.StatusCode >= http.StatusMultipleChoices {
		e := Error{}
		err = prepareError(resp, &e)
		if err != nil {
			return Account{}, err
		}

		msg := e.Error.ErrorName + ": " + e.Error.ErrorMessage
		return Account{}, errors.New(msg)
	}

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Account{}, err
	}

	accountDetail := Account{}
	err = json.Unmarshal(bodyResp, &accountDetail)
	if err != nil {
		return Account{}, err
	}

	return accountDetail, nil
}

// func (a *AccountAPI) Create(emailAddress string) string {
// 	path := BaseURL + SubURLAccount + "/create"

// 	var params bytes.Buffer
// 	writer := multipart.NewWriter(&params)

// 	emailAddressField, err := writer.CreateFormField("email_address")
// 	if err != nil {
// 		return "error"
// 	}
// 	emailAddressField.Write([]byte(emailAddress))

// 	req, err := http.NewRequest(http.MethodPost, path, &params)
// 	req.SetBasicAuth(a.client.apiKey, "")
// 	req.Header.Add("Content-Type", writer.FormDataContentType())
// 	if err != nil {
// 		return ""
// 	}

// 	defaultTransport := &http.Transport{
// 		Dial: (&net.Dialer{
// 			KeepAlive: time.Duration(5000) * time.Millisecond,
// 		}).Dial,
// 		MaxIdleConns:        10,
// 		MaxIdleConnsPerHost: 5,
// 	}
// 	httpClient := http.Client{
// 		Transport: defaultTransport,
// 		Timeout:   time.Duration(2000) * time.Millisecond,
// 	}

// 	resp, err := httpClient.Do(req)
// 	fmt.Println("-------- err -------")
// 	fmt.Printf("%+v\n", err)
// 	fmt.Println("----------------")
// 	if err != nil {
// 		return ""
// 	}

// 	if resp.StatusCode >= 300 {
// 		errResp, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return "read response error"
// 		}
// 		fmt.Println("========  ========")
// 		fmt.Printf("%+v\n", string(errResp))
// 		fmt.Println("=================")
// 		return "new error"
// 	}

// 	bodyResp, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "read response"
// 	}

// 	accountDetail := Account{}
// 	err = json.Unmarshal(bodyResp, &accountDetail)
// 	if err != nil {
// 		return "error unmarshal"
// 	}

// 	fmt.Println("--------  -------")
// 	fmt.Printf("%+v\n", accountDetail)
// 	fmt.Println("----------------")
// 	fmt.Println(a.client.apiKey)
// 	return "this is account"
// }
