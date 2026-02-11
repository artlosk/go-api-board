package announcement

import "github.com/google/uuid"

type Announcement struct {
	ID     uuid.UUID `json:"id"`
	Text   string    `json:"text"`
	UserId uuid.UUID `json:"userId"`
}
