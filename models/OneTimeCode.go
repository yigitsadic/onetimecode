package models

import "time"

type OneTimeCode struct {
	Identifier string
	Value      string
	ExpiresAt  time.Time
}
