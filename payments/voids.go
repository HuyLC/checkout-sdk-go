package payments

// VoidsRequest ...
type VoidsRequest struct {
	Reference string            `json:"reference,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// VoidsResponse ...
type VoidsResponse struct {
	StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
	Accepted       *Accepted                `json:"accepted,omitempty"`
}
