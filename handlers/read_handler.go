package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/yigitsadic/onetimecode/models"
	"github.com/yigitsadic/onetimecode/shared"
	"net/http"
)

type ReadResponse struct {
	Message    string `json:"message"`
	Identifier string `json:"identifier"`
	ExpiresAt  string `json:"expiresAt"`
}

type ReadErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int8   `json:"errorCode"`
}

func HandleRead(redisService *shared.RedisService, ctx *context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(&ReadErrorResponse{
				Message:   "Unable to parse parameters",
				ErrorCode: shared.ERR_CANNOT_PARSE,
			})

			return
		}

		values, ok := r.Form["value"]
		if !ok {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(&ReadErrorResponse{
				Message:   "Value parameter not found",
				ErrorCode: shared.ERR_VALUE_PARAM_NOT_FOUND,
			})

			return
		}

		result, err := readFromRedis(redisService, ctx, values)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(&ReadErrorResponse{
				Message:   "Given value not found",
				ErrorCode: shared.ERR_NOT_FOUND,
			})

			return
		}

		json.NewEncoder(w).Encode(&ReadResponse{
			Message:    "success",
			Identifier: result.Identifier,
			ExpiresAt:  result.ExpiresAt,
		})
	}
}

func readFromRedis(redisService *shared.RedisService, ctx *context.Context, values []string) (*models.OneTimeCode, error) {
	for _, param := range values {
		if len(param) > 0 {
			val, err := redisService.RedisClient.HGetAll(*ctx, param).Result()

			if err != nil {
				continue
			}

			return &models.OneTimeCode{
				Identifier: val["Identifier"],
				Value:      param,
				ExpiresAt:  val["ExpiresAt"],
			}, nil
		}
	}

	return nil, errors.New("unable to find given value")
}
