package hellosign

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
)

const (
	subURLSignatureRequest = "/signature_request"
)

// SignatureRequestAPI is a service to signature request API
type SignatureRequestAPI service

// SignatureRequestList is a response for fetch signature requests
type SignatureRequestList struct {
	ListInfo          ListInfo           `json:"list_info"`
	SignatureRequests []SignatureRequest `json:"signature_requests"`
}

// ListInfo is a information for query parameter response
type ListInfo struct {
	Page       int `json:"page"`
	NumPages   int `json:"num_pages"`
	NumResults int `json:"num_results"`
	PageSize   int `json:"page_size"`
}

// SignatureRequestListParam is param to fetch signature requests list
type SignatureRequestListParam struct {
	AccountID string
	Page      int
	PageSize  int
	Query     string
}

// SignatureRequest is a response for signature request
type SignatureRequest struct {
	SignatureRequest SignatureRequestDetail `json:"signature_request"`
	Warnings         []Warnings             `json:"warnings,omitempty"`
}

// SignatureRequestDetail is a detail for signature request
type SignatureRequestDetail struct {
	TestMode              bool                 `json:"test_mode"`
	SignatureRequestID    string               `json:"signature_request_id"`
	RequesterEmailAddress string               `json:"requester_email_address"`
	Title                 string               `json:"title"`
	OriginalTitle         string               `json:"original_title"`
	Subject               string               `json:"subject"`
	Message               string               `json:"message"`
	CreatedAt             int64                `json:"created_at"`
	IsComplete            bool                 `json:"is_complete"`
	IsDeclined            bool                 `json:"is_declined"`
	HasError              bool                 `json:"has_error"`
	FilesURL              string               `json:"has_url"`
	SigningURL            string               `json:"signing_url"`
	DetailsURL            string               `json:"details_url"`
	CCEmailAddresses      []string             `json:"cc_email_addresses"`
	SigningRedirectURL    string               `json:"signing_redirect_url"`
	CustomFields          []CustomFieldsDetail `json:"custom_fields"`
	ResponseData          []ResponseDataDetail `json:"response_data"`
	Signatures            []SignatureDetail    `json:"signatures"`
	TemplateIDS           string               `json:"template_ids"`
}

// CustomFieldsDetail is details for custom fields
type CustomFieldsDetail struct {
	Name      string      `json:"name"`
	FieldType string      `json:"type"`
	Value     interface{} `json:"value"`
	Required  bool        `json:"required"`
	APIID     string      `json:"api_id"`
	Editor    string      `json:"editor"`
}

// ResponseDataDetail is detail for response data
type ResponseDataDetail struct {
	APIID       string      `json:"api_id"`
	SignatureID string      `json:"signature_id"`
	Name        string      `json:"name"`
	Value       interface{} `json:"value"`
	Required    bool        `json:"required"`
	FieldType   string      `json:"type"`
}

// SignatureDetail is detail for signature
type SignatureDetail struct {
	SignatureID        string `json:"signature_id"`
	SignerEmailAddress string `json:"signer_email_address"`
	SignerName         string `json:"signer_name"`
	SignerRole         string `json:"signer_role"`
	Order              int    `json:"order"`
	StatusCode         string `json:"status_code"`
	DeclineReason      string `json:"decline_reason"`
	SignedAt           int64  `json:"signed_at"`
	LastViewedAt       int64  `json:"last_viewed_at"`
	LastRemindedAt     int64  `json:"last_reminded_at"`
	HasPin             bool   `json:"has_pin"`
	ReassignedBy       string `json:"reassigned_by"`
	ReassignmentReason string `json:"reassignment_reason"`
	Error              string `json:"error"`
}

// Get will return a signature request by signature request id
func (s *SignatureRequestAPI) Get(id string) (SignatureRequest, error) {
	path := s.client.BaseURL + subURLSignatureRequest + "/" + id
	resp, err := s.client.doRequest(
		path,
		http.MethodGet,
		&bytes.Buffer{},
		&multipart.Writer{},
	)
	if err != nil {
		return SignatureRequest{}, err
	}
	defer resp.Body.Close()

	signatureRequest := SignatureRequest{}
	err = json.NewDecoder(resp.Body).Decode(&signatureRequest)
	if err != nil {
		return SignatureRequest{}, err
	}

	return signatureRequest, nil
}
