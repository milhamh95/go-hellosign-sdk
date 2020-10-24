package hellosign

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	subURLSignatureRequest = "/signature_request"
)

// SignatureRequestAPI is a service to signature request API
type SignatureRequestAPI service

// SignatureRequestListParam is param to fetch signature requests list
type SignatureRequestListParam struct {
	ListInfoQueryParam
	AccountID string
	Query     string
}

// SignatureRequestList is a response for fetch signature requests
type SignatureRequestList struct {
	ListInfo          ListInfo           `json:"list_info"`
	SignatureRequests []SignatureRequest `json:"signature_requests"`
}

// SignatureRequest is a response for signature request
type SignatureRequest struct {
	SignatureRequest SignatureRequestDetail `json:"signature_request"`
	Warnings         []Warnings             `json:"warnings,omitempty"`
}

// SignatureRequestDetail is a detail for signature request
type SignatureRequestDetail struct {
	TestMode              bool                   `json:"test_mode"`
	SignatureRequestID    string                 `json:"signature_request_id"`
	RequesterEmailAddress string                 `json:"requester_email_address"`
	Title                 string                 `json:"title"`
	OriginalTitle         string                 `json:"original_title"`
	Subject               string                 `json:"subject"`
	Message               string                 `json:"message"`
	CreatedAt             int64                  `json:"created_at"`
	IsComplete            bool                   `json:"is_complete"`
	IsDeclined            bool                   `json:"is_declined"`
	HasError              bool                   `json:"has_error"`
	FilesURL              string                 `json:"has_url"`
	SigningURL            string                 `json:"signing_url"`
	DetailsURL            string                 `json:"details_url"`
	CCEmailAddresses      []string               `json:"cc_email_addresses"`
	SigningRedirectURL    string                 `json:"signing_redirect_url"`
	CustomFields          []CustomFieldsDetail   `json:"custom_fields"`
	ResponseData          []ResponseDataDetail   `json:"response_data"`
	Signatures            []SignatureDetail      `json:"signatures"`
	Metadata              map[string]interface{} `json:"metadata"`
	TemplateIDS           string                 `json:"template_ids"`
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

// SignatureRequestPayload is payload for signature request
type SignatureRequestPayload struct {
	TestMode              int                         `json:"test_mode"`
	FileURL               []string                    `json:"file_url"`
	Title                 string                      `json:"title"`
	Subject               string                      `json:"subject"`
	Message               string                      `json:"message"`
	SigningRedirectURL    string                      `json:"signing_redirect_url"`
	Signers               []SignerDetail              `json:"signers"`
	Attachments           []SignatureAttachmentDetail `json:"attachments"`
	CustomFields          []CustomFieldsDetail        `json:"custom_fields"`
	CCEmailAddresses      []string                    `json:"cc_email_addresses"`
	UseTextTags           int                         `json:"use_text_tags"`
	HideTextTags          int                         `json:"hide_text_tags"`
	Metadata              map[string]interface{}      `json:"metadata"`
	ClientID              string                      `json:"client_id"`
	AllowDecline          int                         `json:"allow_decline"`
	AllowReassign         int                         `json:"allow_reassign"`
	FormFieldsPerDocument [][]FormFieldDetail         `json:"form_fields_per_document"`
	SigningOptions        map[string]interface{}      `json:"signing_options"`
	FieldOptions          FieldOptionsDetail          `json:"field_options"`
}

// SignerDetail is detail for signer
type SignerDetail struct {
	Name         string              `json:"name"`
	EmailAddress string              `json:"email_address"`
	Order        int                 `json:"int"`
	Pin          int                 `json:"pin"`
	Group        []SignerGroupDetail `json:"signers"`
}

// SignerGroupDetail is detail for signer group
type SignerGroupDetail struct {
	Name         string `json:"name"`
	EmailAddress string `json:"email_address"`
}

// SignatureAttachmentDetail is detail for signature attachment
type SignatureAttachmentDetail struct {
	Name         string `json:"name"`
	Instructions string `json:"instructions"`
	SignerIndex  int    `json:"signer_index"`
	Required     bool   `json:"required"`
	UploadedAt   int64  `json:"uploaded_at"`
}

// FieldOptionsDetail is detail for field options
type FieldOptionsDetail struct {
	DateFormat string `json:"date_format"`
}

// FormFieldDetail is detail for form fields per document
type FormFieldDetail struct {
	APIID    string `json:"api_id"`
	Type     string `json:"type"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Page     int    `json:"page"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Required bool   `json:"required"`
	Signer   int    `json:"signer"`
}

const (
	// FieldText is text field type
	FieldText = "text"
	// FieldCheckBox is check box field type
	FieldCheckBox = "checkbox"
	// FieldDateSigned is date signed field type
	FieldDateSigned = "date_signed"
	// FieldDropdown is dropdown field type
	FieldDropdown = "dropdown"
	// FieldInitials is initals field type
	FieldInitials = "initials"
	// FieldRadio is radio field type
	FieldRadio = "radio"
	// FieldSignature is signature field type
	FieldSignature = "signature"
	// FieldTextMerge is text merge field type
	FieldTextMerge = "text-merge"
	// FieldCheckboxMerge is check box mrege field type
	FieldCheckboxMerge = "checkbox-merge"
)

const (
	// DateFormat1 is date format for MM / DD / YYYY ex: 10 / 16 / 2020
	DateFormat1 = "MM / DD/ YYYY"
	// DateFormat2 is date format for MM - DD - YYYY ex: 10 - 16 - 2020
	DateFormat2 = "MM - DD - YYYY"
	// DateFormat3 is date format for DD / MM / YYYY ex: 16 / 10 / 2020
	DateFormat3 = "DD / MM / YYYY"
	// DateFormat4 is date format for DD - MM - YYYY ex: 16 - 10 - 2020
	DateFormat4 = "DD - MM - YYYY"
	// DateFormat5 is date format for YYYY / MM / DD ex: 2020 / 10 / 16
	DateFormat5 = "YYYY / MM / DD"
	// DateFormat6 is date format for YYYY - MM - DD ex: 2020 - 10 - 16
	DateFormat6 = "YYYY - MM - DD"
)

// Get will return a signature request by signature request id
func (s *SignatureRequestAPI) Get(ctx context.Context, id string) (SignatureRequest, error) {
	path := s.client.BaseURL + subURLSignatureRequest + "/" + id
	resp, err := s.client.callAPI(
		ctx,
		requestParam{
			path:   path,
			method: http.MethodGet,
		},
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

// Fetch will return signture request list based on param
func (s *SignatureRequestAPI) Fetch(ctx context.Context, p SignatureRequestListParam) (SignatureRequestList, error) {
	path := s.client.BaseURL + subURLSignatureRequest
	req, err := s.client.prepareRequest(
		ctx,
		requestParam{
			path:   path,
			method: http.MethodGet,
		})
	if err != nil {
		return SignatureRequestList{}, err
	}

	q := req.URL.Query()
	q.Add("account_id", p.AccountID)
	q.Add("page", strconv.Itoa(p.Page))
	q.Add("page_size", strconv.Itoa(p.PageSize))
	q.Add("query", p.Query)

	req.URL.RawQuery = q.Encode()

	resp, err := s.client.executeRequest(req)
	if err != nil {
		return SignatureRequestList{}, err
	}

	signatureRequestList := SignatureRequestList{}
	err = json.NewDecoder(resp.Body).Decode(&signatureRequestList)
	if err != nil {
		return SignatureRequestList{}, err
	}

	return signatureRequestList, nil
}
