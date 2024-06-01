package utils

import "github.com/google/uuid"

func GenerateUUID() string {
	key := uuid.New().String()
	return key
}