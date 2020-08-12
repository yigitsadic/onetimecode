package handlers

import (
	"encoding/json"
	"github.com/yigitsadic/onetimecode/models"
	"github.com/yigitsadic/onetimecode/responses"
	"github.com/yigitsadic/onetimecode/shared"
	"log"
	"net/http"
	"time"
)

func HandleRead(codeStore *models.CodeStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		val := r.URL.Query().Get("code")

		if val == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(&responses.ReadErrorResponse{
				Message:   "Unable to parse parameters",
				ErrorCode: shared.ERR_CANNOT_PARSE,
			})

			return
		}

		parsed, ok := codeStore.Codes[val]
		if !ok {
			log.Printf("Code %s is not found\n", val)

			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(&responses.ReadErrorResponse{
				Message:   "Value parameter not found",
				ErrorCode: shared.ERR_VALUE_PARAM_NOT_FOUND,
			})

			return
		}

		log.Printf("Served Code %s\t with identifier %s at %s\n", parsed.Value, parsed.Identifier, time.Now().UTC())

		json.NewEncoder(w).Encode(&responses.ReadResponse{
			Message:    "success",
			Identifier: parsed.Identifier,
			ExpiresAt:  parsed.ExpiresAt.UTC().Unix(),
		})
	}
}
