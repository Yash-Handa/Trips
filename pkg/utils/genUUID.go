package utils

import uuid "github.com/satori/go.uuid"

// GenUUID generates a UUID as string
func GenUUID() string {
	return uuid.NewV4().String()
}
