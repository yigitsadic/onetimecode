package handlers

import (
	"context"
	"encoding/json"
	"github.com/yigitsadic/onetimecode/shared"
	"math/rand"
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

		err := redisService.RedisClient.Set(*ctx, createRandomValue(dto.Identifier), dto.Identifier, time.Second*120).Err()
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
			Value:      createRandomValue(dto.Identifier),
			ExpiresAt:  time.Now().Add(120 * time.Second).Unix(),
		})
	}
}

func createRandomValue(identifier string) string {
	var built []byte
	var source []byte
	for x := byte('A'); x <= byte('Z'); x++ {
		source = append(source, x)
	}

	for x := byte('0'); x <= byte('9'); x++ {
		source = append(source, x)
	}

	for x := 0; x < 7; x++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		built = append(built, source[r1.Intn(len(source)-1)])
	}

	return string(built)
}
