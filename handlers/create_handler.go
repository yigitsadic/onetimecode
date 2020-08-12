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

type CreateDto struct {
	Identifier string `json:"identifier"`
}

func (c CreateDto) Validate() bool {
	if len(c.Identifier) == 0 {
		return false
	}

	return true
}

func HandleCreate(codeStore *models.CodeStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var dto *CreateDto
		json.NewDecoder(r.Body).Decode(&dto)

		if !dto.Validate() {
			w.WriteHeader(http.StatusUnprocessableEntity)

			json.NewEncoder(w).Encode(&responses.FailedCreationResponse{
				Message:   "Unable to create with given values",
				ErrorCode: shared.ERR_UNABLE_TO_CREATE,
			})

			return
		}

		codeStore.Mux.Lock()
		defer codeStore.Mux.Unlock()

		next := true
		var val *models.OneTimeCode
		var ok bool

		for next {
			randVal := shared.CreateRandomValue()
			val, ok = codeStore.Codes[randVal]

			if !ok {
				val = &models.OneTimeCode{
					Identifier: dto.Identifier,
					Value:      randVal,
					ExpiresAt:  time.Now().UTC().Add(time.Second * time.Duration(codeStore.Expiration)),
				}

				next = false
			}
		}

		codeStore.Codes[val.Value] = val

		go func(c *models.OneTimeCode, s *models.CodeStore) {
			log.Printf("Enqueued delete key job with value=%s\tidentifier=%s\texpiresAt=%s\n", c.Value, c.Identifier, c.ExpiresAt)

			time.AfterFunc(time.Until(c.ExpiresAt), func() {
				log.Printf("Code %s is expired and deleted at %s\n", c.Value, time.Now())
				delete(s.Codes, c.Value)
			})
		}(val, codeStore)

		json.NewEncoder(w).Encode(&responses.CreateResponse{
			Identifier: dto.Identifier,
			Value:      val.Value,
			ExpiresAt:  val.ExpiresAt.UTC().Unix(),
		})
	}
}
