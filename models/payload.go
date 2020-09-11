package models

import "github.com/gbrlsnchs/jwt/v3"

// just a wrapper around jwt's recommended payload
type CustomPayload struct {
	jwt.Payload
}
