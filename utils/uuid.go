package utils

import (
	"github.com/google/uuid"

	"url-shortener/log"
)

func GenerateId() (uuid.UUID, error) {
	u, err := uuid.NewV7()
	if err != nil {
		log.Logger.Printf("Generate uuid failed, error is: %s", err)
		return uuid.Nil, err
	}

	return u, nil
}
