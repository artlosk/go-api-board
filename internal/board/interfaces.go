package board

import (
	"github.com/google/uuid"

	"board/internal/announcement"
)

type BoardRepository interface {
	AddAnnouncement(announcement announcement.Announcement) error
	UpdateAnnouncement(announcement announcement.Announcement) error
	DeleteAnnouncement(id uuid.UUID) error
	List() []announcement.Announcement
	Get(id uuid.UUID) (announcement.Announcement, error)
}
