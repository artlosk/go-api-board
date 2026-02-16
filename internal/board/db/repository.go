package db

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"board/internal/announcement"
)

type Repository struct {
	db *sql.DB
}

func (r Repository) AddAnnouncement(announcement announcement.Announcement) error {
	const q = "INSERT INTO announcement (id, user_id, text) VALUES ($1, $2, $3)"
	exec, err := r.db.Exec(q, announcement.ID, announcement.UserId, announcement.Text)
	if err != nil {
		return fmt.Errorf("failed to add announcement: %w", err)
	}

	return nil
}

func (r Repository) UpdateAnnouncement(announcement announcement.Announcement) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) DeleteAnnouncement(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) List() []announcement.Announcement {
	//TODO implement me
	panic("implement me")
}

func (r Repository) Get(id uuid.UUID) (announcement.Announcement, error) {
	//TODO implement me
	panic("implement me")
}
