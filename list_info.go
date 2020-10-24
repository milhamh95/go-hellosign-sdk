package hellosign

// ListInfo is a information for query parameter response
type ListInfo struct {
	Page       int `json:"page"`
	NumPages   int `json:"num_pages"`
	NumResults int `json:"num_results"`
	PageSize   int `json:"page_size"`
}

// ListInfoQueryParam is query param for list info
type ListInfoQueryParam struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
