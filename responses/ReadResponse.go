package responses

type ReadResponse struct {
	Message    string `json:"message"`
	Identifier string `json:"identifier"`
	ExpiresAt  int64  `json:"expiresAt"`
}

type ReadErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int8   `json:"errorCode"`
}
