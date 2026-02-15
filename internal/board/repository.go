package board

import (
	"github.com/google/uuid"

	"board/internal/announcement"
)

type BoardRepository interface {
	AddAnnouncement(announcement.Announcement) error
	UpdateAnnouncement(announcement.Announcement) error
	DeleteAnnouncement(id uuid.UUID) error
	List() []announcement.Announcement
	Get(id uuid.UUID) (announcement.Announcement, error)
}

type memoryRepository struct {
	announcements []announcement.Announcement
}

func NewRepository() BoardRepository {
	return &memoryRepository{
		announcements: make([]announcement.Announcement, 0),
	}
}

func (r *memoryRepository) AddAnnouncement(a announcement.Announcement) error {
	r.announcements = append(r.announcements, a)
	return nil
}

func (r *memoryRepository) UpdateAnnouncement(a announcement.Announcement) error {
	for i := range r.announcements {
		if r.announcements[i].ID == a.ID {
			r.announcements[i] = a
			return nil
		}
	}
	return NotFoundAnnouncement
}

func (r *memoryRepository) DeleteAnnouncement(id uuid.UUID) error {
	for i := range r.announcements {
		if r.announcements[i].ID == id {
			r.announcements = append(r.announcements[:i], r.announcements[i+1:]...)
			return nil
		}
	}
	return NotFoundAnnouncement
}

func (r *memoryRepository) List() []announcement.Announcement {
	return r.announcements
}

func (r *memoryRepository) Get(id uuid.UUID) (announcement.Announcement, error) {
	for _, a := range r.announcements {
		if a.ID == id {
			return a, nil
		}
	}
	return announcement.Announcement{}, NotFoundAnnouncement
}
