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

	errUnknownJSON := testdata.GetGolden(t, "err-unknown")
	errUnauthorizedJSON := testdata.GetGolden(t, "err-unauthorized-api-key")

	tests := map[string]struct {
		accountResponse http.Response
		expectedAccount hellosign.Account
		expectedError   error
	}{
		"success": {
			accountResponse: http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(accountJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: account,
			expectedError:   nil,
		},
		"unauthorized api key": {
			accountResponse: http.Response{
				StatusCode: http.StatusNotFound,
				Body:       ioutil.NopCloser(bytes.NewReader(errUnauthorizedJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: hellosign.Account{},
			expectedError:   errors.New("unauthorized: Unauthorized api key"),
		},
		"unknown error": {
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

			apiClient := hellosign.NewClient("123", mockHTTPClient)
			resp, err := apiClient.AccountAPI.Get()
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

	errUnauthorizedPaidPlanJSON := testdata.GetGolden(t, "err-unauthorized-paid-plan")

	notVerifiedAccount := hellosign.Account{}
	notVerifiedAccountJSON, err := json.Marshal(notVerifiedAccount)
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
		"unauthorized access to paid plan": {
			email: "rifivazu-0282@gmail.com",
			accountResponse: http.Response{
				StatusCode: http.StatusForbidden,
				Body:       ioutil.NopCloser(bytes.NewReader(errUnauthorizedPaidPlanJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: hellosign.Account{},
			expectedError:   errors.New("forbidden: A paid API plan is required to access this endpoint"),
		},
		"not verified": {
			email: "rifivazu-0282@gmail.com",
			accountResponse: http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(notVerifiedAccountJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: hellosign.Account{},
			expectedError:   nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			is := is.New(t)
			mockHTTPClient := testdata.NewClient(t, func(req *http.Request) *http.Response {
				return &test.accountResponse
			})

			apiClient := hellosign.NewClient("123", mockHTTPClient)
			resp, err := apiClient.AccountAPI.Verify(test.email)
			if test.expectedError != nil {
				is.Equal(test.expectedError.Error(), err.Error())
				return
			}
			is.NoErr(err)
			is.Equal(test.expectedAccount, resp)
		})
	}
}

func TestAccount_Update(t *testing.T) {
	is := is.New(t)

	accountJSON := testdata.GetGolden(t, "account-1")

	account := hellosign.Account{}
	err := json.Unmarshal(accountJSON, &account)
	is.NoErr(err)

	tests := map[string]struct {
		callbackURL     string
		accountResponse http.Response
		expectedAccount hellosign.Account
		expectedError   error
	}{
		"success": {
			callbackURL: "rifivazu-0282@gmail.com",
			accountResponse: http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(accountJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: account,
			expectedError:   nil,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			is := is.New(t)
			mockHTTPClient := testdata.NewClient(t, func(req *http.Request) *http.Response {
				return &test.accountResponse
			})

			apiClient := hellosign.NewClient("123", mockHTTPClient)
			resp, err := apiClient.AccountAPI.Update(test.callbackURL)
			if err != nil {
				is.Equal(test.expectedError.Error(), err.Error())
				return
			}

			is.NoErr(err)
			is.Equal(test.expectedAccount, resp)
		})
	}
}

func TestAccount_Create(t *testing.T) {
	is := is.New(t)
	accountJSON := testdata.GetGolden(t, "account-1")

	account := hellosign.Account{}
	err := json.Unmarshal(accountJSON, &account)
	is.NoErr(err)

	errBadRequestJSON := testdata.GetGolden(t, "account-err-bad-request")

	tests := map[string]struct {
		emailAddress    string
		accountResponse http.Response
		expectedAccount hellosign.Account
		expectedError   error
	}{
		"success": {
			emailAddress: "rifivazu-0282@gmail.com",
			accountResponse: http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(accountJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: account,
			expectedError:   nil,
		},
		"invalid email address parameter": {
			emailAddress: "rifivazu-0282@gmail.com",
			accountResponse: http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       ioutil.NopCloser(bytes.NewReader(errBadRequestJSON)),
				Header:     make(http.Header),
			},
			expectedAccount: hellosign.Account{},
			expectedError:   errors.New("bad_request: Invalid parameter: email_addres"),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			is := is.New(t)
			mockHTTPClient := testdata.NewClient(t, func(req *http.Request) *http.Response {
				return &test.accountResponse
			})

			apiClient := hellosign.NewClient("123", mockHTTPClient)
			res, err := apiClient.AccountAPI.Create(test.emailAddress)
			if test.expectedError != nil {
				is.Equal(test.expectedError.Error(), err.Error())
				return
			}

			is.NoErr(err)
			is.Equal(test.expectedAccount, res)
		})
	}
}
