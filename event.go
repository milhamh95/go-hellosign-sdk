package hellosign

// Event represent callback event response
type Event struct {
	Event            EventDetail      `json:"event"`
	SignatureRequest SignatureRequest `json:"signature_request"`
}

// EventDetail is detail for event
type EventDetail struct {
	EventTime     int64         `json:"event_time"`
	EventType     string        `json:"event_type"`
	EventHash     string        `json:"event_hash"`
	EventMetadata EventMetadata `json:"event_metadata"`
}

// EventMetadata is detail for event metadata
type EventMetadata struct {
	RelatedSignatureID   string `json:"related_signature_id"`
	ReportedForAccountID string `json:"reported_for_account_id"`
	ReportedForAppID     string `json:"reported_for_app_id"`
}
