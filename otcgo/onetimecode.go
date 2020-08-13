package otcgo

import (
	"errors"
	"github.com/yigitsadic/onetimecode/grpc/grpc"
	"github.com/yigitsadic/onetimecode/models"
	"github.com/yigitsadic/onetimecode/shared"
	"golang.org/x/net/context"
	"log"
	"time"
)

type Server struct {
	CodeStore *models.CodeStore
}

func (s *Server) CreateCode(ctx context.Context, gen *grpc.OneTimeCodeGen) (*grpc.OneTimeCodeResponse, error) {
	s.CodeStore.Mux.Lock()
	defer s.CodeStore.Mux.Unlock()

	next := true
	var val *models.OneTimeCode
	var ok bool

	for next {
		randVal := shared.CreateRandomValue()
		val, ok = s.CodeStore.Codes[randVal]

		if !ok {
			val = &models.OneTimeCode{
				Identifier: gen.Identifier,
				Value:      randVal,
				ExpiresAt:  time.Now().UTC().Add(time.Second * time.Duration(s.CodeStore.Expiration)),
			}

			next = false
		}
	}

	s.CodeStore.Codes[val.Value] = val

	go func(c *models.OneTimeCode, s *models.CodeStore) {
		log.Printf("Enqueued delete key job with value=%s\tidentifier=%s\texpiresAt=%s\n", c.Value, c.Identifier, c.ExpiresAt)

		time.AfterFunc(time.Until(c.ExpiresAt), func() {
			log.Printf("Code %s is expired and deleted at %s\n", c.Value, time.Now())
			delete(s.Codes, c.Value)
		})
	}(val, s.CodeStore)

	return &grpc.OneTimeCodeResponse{
		Identifier: val.Identifier,
		ExpiresAt:  val.ExpiresAt.UTC().Unix(),
		Value:      val.Value,
	}, nil
}

func (s *Server) ReadCode(ctx context.Context, req *grpc.ReadCodeReq) (*grpc.OneTimeCodeResponse, error) {
	parsed, ok := s.CodeStore.Codes[req.Value]
	if !ok {
		log.Printf("Code %s is not found\n", req.Value)

		return nil, errors.New("unable to fetch given code")
	}

	log.Printf("Served Code %s\t with identifier %s at %s\n", parsed.Value, parsed.Identifier, time.Now().UTC())

	return &grpc.OneTimeCodeResponse{
		Identifier: parsed.Identifier,
		ExpiresAt:  parsed.ExpiresAt.UTC().Unix(),
		Value:      parsed.Value,
	}, nil
}
