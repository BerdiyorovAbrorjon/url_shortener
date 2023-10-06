package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token error")
	ErrExpireToken  = errors.New("token has expired")
)

// Payload contains payload data of the token
type Payload struct {
	ID       uuid.UUID `json:"id"`
	UserID   int32     `json:"user_id"`
	Email    string    `json:"email"`
	IssuedAt time.Time `json:"issued_at"`
	ExpireAt time.Time `json:"expire_at"`
}

// NewPayload is function which to create new Payload with specific params: username and expire duration
func NewPayload(userId int32, email string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:       tokenID,
		UserID:   userId,
		Email:    email,
		IssuedAt: time.Now(),
		ExpireAt: time.Now().Add(duration),
	}, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpireAt) {
		return ErrExpireToken
	}
	return nil
}
