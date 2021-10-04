package entity

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AlbumID EntityID

const albumPrefix string = "album"

func (a AlbumID) Validate() error {
	return validation.Validate(string(a), ValidateIDRule(albumPrefix))
}

func GenerateAlbumID() AlbumID {
	return AlbumID(GenerateID(albumPrefix))
}

// Album represents an album record.
type Album struct {
	ID        AlbumID   `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
