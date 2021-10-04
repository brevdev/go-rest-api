package entity

import (
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type EntityID string

// GenerateID generates a unique ID that can be used as an identifier for an entity.
func GenerateID(prefix string) EntityID {
	return EntityID(fmt.Sprintf("%s-%s", prefix, uuid.New().String()))
}

func ValidateIDRule(prefix string) validation.Rule {
	re := regexp.MustCompile(fmt.Sprintf(`%s_*`, prefix))
	rule := validation.Match(re)
	return rule.Error(fmt.Sprintf("invalid %s id", prefix))
}
