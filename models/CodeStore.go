package models

import "sync"

type CodeStore struct {
	Codes      map[string]*OneTimeCode
	Mux        sync.Mutex
	Expiration int
}

func NewCodeStore(exp int) *CodeStore {
	return &CodeStore{
		Codes:      make(map[string]*OneTimeCode),
		Expiration: exp,
	}
}
