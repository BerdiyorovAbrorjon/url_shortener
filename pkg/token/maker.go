package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	//CreateToken is function for create token string
	CreateToken(userId int32, email string, duration time.Duration) (string, error)

	//VerifyToken is function for check token to valid or not
	VerifyToken(token string) (*Payload, error)
}
