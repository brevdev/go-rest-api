package entity

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type EntityID string

// GenerateID generates a unique ID that can be used as an identifier for an entity.
func GenerateID(prefix string) EntityID {
	return EntityID(fmt.Sprintf("%s-%s", prefix, uuid.New().String()))
}

func ValidateID(prefix string, id EntityID) error {
	if !strings.HasPrefix(string(id), prefix) {
		return fmt.Errorf("invalid %s id  - %s", prefix, id)
	}
	return nil
}
