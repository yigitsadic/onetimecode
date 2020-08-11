package models

type OneTimeCode struct {
	Identifier string
	Value      string
	ExpiresAt  string
}
