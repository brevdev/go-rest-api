package entity

import (
	"time"
)

type AlbumID EntityID

func GenerateAlbumID() AlbumID {
	return AlbumID(GenerateID("album"))
}

func ValidateAlbumID(id string) (AlbumID, error) {
	if err := ValidateID("user", EntityID(id)); err != nil {
		return "", err
	}
	return AlbumID(id), nil
}

// Album represents an album record.
type Album struct {
	ID        AlbumID   `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
