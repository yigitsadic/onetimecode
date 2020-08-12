package responses

type CreateResponse struct {
	Identifier string `json:"identifier"`
	Value      string `json:"value"`
	ExpiresAt  int64  `json:"expiresAt"`
}

type FailedCreationResponse struct {
	Message   string `json:"message"`
	ErrorCode int8   `json:"errorCode"`
}
