package hellosign_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/matryer/is"

	hellosign "go-hellosign-sdk"
	"go-hellosign-sdk/testdata"
)

func TestAccount_Get(t *testing.T) {
	is := is.New(t)

	accountJSON := testdata.GetGolden(t, "account-1")

	account := hellosign.Account{}
	err := json.Unmarshal(accountJSON, &account)
	is.NoErr(err)

	errUnknownJSON := testdata.GetGolden(t, "account-1-err-unknown")
	errUnauthorizedJSON := testdata.GetGolden(t, "err-unauthorized")

	tests := map[string]struct {
		accountID       string
		accountResponse http.Response
		expectedAccount hellosign.Account
		expectedError   error
	}{
		"success": {
			accountID: "1",
			accountResponse: http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(accountJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: account,
			expectedError:   nil,
		},
		"unauthorized": {
			accountID: "1",
			accountResponse: http.Response{
				StatusCode: http.StatusNotFound,
				Body:       ioutil.NopCloser(bytes.NewReader(errUnauthorizedJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: hellosign.Account{},
			expectedError:   errors.New("unauthorized: Unauthorized api key"),
		},
		"unknown error": {
			accountID: "1",
			accountResponse: http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       ioutil.NopCloser(bytes.NewReader(errUnknownJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: hellosign.Account{},
			expectedError:   errors.New("unknown: Unknown error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			is := is.New(t)
			mockHTTPClient := testdata.NewClient(t, func(req *http.Request) *http.Response {
				return &test.accountResponse
			})

			apiClient := hellosign.NewAPI("123", mockHTTPClient)
			resp, err := apiClient.AccountAPI.GetByID(test.accountID)
			if err != nil {
				is.Equal(test.expectedError.Error(), err.Error())
				return
			}

			is.NoErr(err)
			is.Equal(test.expectedAccount, resp)
		})
	}
}

func TestAccount_GetByID(t *testing.T) {
	is := is.New(t)

	accountJSON := testdata.GetGolden(t, "account-1")

	account := hellosign.Account{}
	err := json.Unmarshal(accountJSON, &account)
	is.NoErr(err)

	errNotFoundJSON := testdata.GetGolden(t, "account-1-err-not-found")
	errUnknownJSON := testdata.GetGolden(t, "err-unknown")

	tests := map[string]struct {
		accountID       string
		accountResponse http.Response
		expectedAccount hellosign.Account
		expectedError   error
	}{
		"success": {
			accountID: "1",
			accountResponse: http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(accountJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: account,
			expectedError:   nil,
		},
		"not found": {
			accountID: "1",
			accountResponse: http.Response{
				StatusCode: http.StatusNotFound,
				Body:       ioutil.NopCloser(bytes.NewReader(errNotFoundJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: hellosign.Account{},
			expectedError:   errors.New("not_found: User not found for ID 1"),
		},
		"unknown error": {
			accountID: "1",
			accountResponse: http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       ioutil.NopCloser(bytes.NewReader(errUnknownJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: hellosign.Account{},
			expectedError:   errors.New("unknown: Unknown error"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			is := is.New(t)
			mockHTTPClient := testdata.NewClient(t, func(req *http.Request) *http.Response {
				return &test.accountResponse
			})

			apiClient := hellosign.NewAPI("123", mockHTTPClient)
			resp, err := apiClient.AccountAPI.GetByID(test.accountID)
			if err != nil {
				is.Equal(test.expectedError.Error(), err.Error())
				return
			}

			is.NoErr(err)
			is.Equal(test.expectedAccount, resp)
		})
	}
}

func TestAccount_Verify(t *testing.T) {
	is := is.New(t)

	accountResp := hellosign.Account{
		Account: hellosign.AccountDetail{
			EmailAddress: "rifivazu-0282@gmail.com",
		},
	}
	accountRespJSON, err := json.Marshal(accountResp)
	is.NoErr(err)

	tests := map[string]struct {
		email           string
		accountResponse http.Response
		expectedAccount hellosign.Account
		expectedError   error
	}{
		"success": {
			email: "rifivazu-0282@gmail.com",
			accountResponse: http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(accountRespJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: accountResp,
			expectedError:   nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			is := is.New(t)
			mockHTTPClient := testdata.NewClient(t, func(req *http.Request) *http.Response {
				return &test.accountResponse
			})

			apiClient := hellosign.NewAPI("123", mockHTTPClient)
			resp, err := apiClient.AccountAPI.Verify(test.email)
			if err != nil {
				is.Equal(test.expectedError.Error(), err.Error())
				return
			}
			is.NoErr(err)
			is.Equal(test.expectedAccount, resp)
		})
	}
}

// func TestGet2_Account(t *testing.T) {
// 	dataResult := []byte(`{"account":{"account_id":"18b7e8510490b4125f97bd3fec40a5b53345873d","email_address":"photon628@gmail.com","is_locked":false,"is_paid_hs":false,"is_paid_hf":false,"quotas":{"templates_left":0,"documents_left":3,"api_signature_requests_left":0},"callback_url":null,"locale":"en-US","role_code":null}}`)
// 	tests := map[string]struct {
// 		accountID          string
// 		accountResponse    map[string]testdata.HTTPCall
// 		expectedStatusCode int
// 	}{
// 		"success": {
// 			accountID: "1",
// 			accountResponse: map[string]testdata.HTTPCall{
// 				"GET /account/1": {
// 					Header: map[string]string{
// 						"Content-Type": "application/json",
// 					},
// 					Method:       http.MethodDelete,
// 					Status:       http.StatusOK,
// 					ExpectedResp: dataResult,
// 				},
// 			},
// 			expectedStatusCode: http.StatusOK,
// 		},
// 	}

// 	for testName, test := range tests {
// 		t.Run(testName, func(t *testing.T) {
// 			server, serverStop := testdata.StartServer(t, test.accountResponse)
// 			defer serverStop()

// 			apiClient := hellosign.NewAPI("123", &http.Client{Timeout: 2 * time.Second})
// 			apiClient.BaseURL = server.URL
// 			resp, err := apiClient.AccountAPI.Get(test.accountID)
// 			fmt.Println(resp)
// 		})
// 	}
// }
