package handlers

import (
	"context"
	"encoding/json"
	"github.com/yigitsadic/onetimecode/shared"
	"net/http"
	"time"
)

type CreateDto struct {
	Identifier string `json:"identifier"`
}

func (c CreateDto) Validate() bool {
	if len(c.Identifier) == 0 {
		return false
	}

	return true
}

type CreateResponse struct {
	Identifier string `json:"identifier"`
	Value      string `json:"value"`
	ExpiresAt  int64  `json:"expiresAt"`
}

type FailedCreationResponse struct {
	Message   string `json:"message"`
	ErrorCode int8   `json:"errorCode"`
}

func HandleCreate(redisService *shared.RedisService, ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var dto *CreateDto
		json.NewDecoder(r.Body).Decode(&dto)

		if !dto.Validate() {
			w.WriteHeader(http.StatusUnprocessableEntity)

			json.NewEncoder(w).Encode(&FailedCreationResponse{
				Message:   "Unable to create with given values",
				ErrorCode: shared.ERR_UNABLE_TO_CREATE,
			})

			return
		}

		randVal := shared.CreateRandomValue()

		err := redisService.RedisClient.HSet(*ctx, randVal, map[string]interface{}{
			"Identifier": dto.Identifier,
			"ExpiresAt":  time.Now().Add(120 * time.Second).Unix(),
		}).Err()
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)

			json.NewEncoder(w).Encode(&FailedCreationResponse{
				Message:   "Unable to create with given values",
				ErrorCode: shared.ERR_UNABLE_TO_CREATE,
			})

			return
		}

		json.NewEncoder(w).Encode(&CreateResponse{
			Identifier: dto.Identifier,
			Value:      randVal,
			ExpiresAt:  time.Now().Add(120 * time.Second).Unix(),
		})
	}
}
