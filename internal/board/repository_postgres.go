package board

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"board/internal/announcement"
)

type postgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) BoardRepository {
	return &postgresRepository{pool: pool}
}

func (r *postgresRepository) AddAnnouncement(a announcement.Announcement) error {
	_, err := r.pool.Exec(context.Background(),
		`INSERT INTO announcements (id, text, user_id) VALUES ($1, $2, $3)`,
		a.ID, a.Text, a.UserId)
	return err
}

func (r *postgresRepository) UpdateAnnouncement(a announcement.Announcement) error {
	cmd, err := r.pool.Exec(context.Background(),
		`UPDATE announcements SET text = $1 WHERE id = $2`, a.Text, a.ID)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return NotFoundAnnouncement
	}
	return nil
}

func (r *postgresRepository) DeleteAnnouncement(id uuid.UUID) error {
	cmd, err := r.pool.Exec(context.Background(), `DELETE FROM announcements WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return NotFoundAnnouncement
	}
	return nil
}

func (r *postgresRepository) List() []announcement.Announcement {
	rows, err := r.pool.Query(context.Background(),
		`SELECT id, text, user_id FROM announcements ORDER BY id`)
	if err != nil {
		return []announcement.Announcement{}
	}
	defer rows.Close()

	var list []announcement.Announcement
	for rows.Next() {
		var a announcement.Announcement
		if err := rows.Scan(&a.ID, &a.Text, &a.UserId); err != nil {
			return list
		}
		list = append(list, a)
	}
	return list
}

func (r *postgresRepository) Get(id uuid.UUID) (announcement.Announcement, error) {
	var a announcement.Announcement
	err := r.pool.QueryRow(context.Background(),
		`SELECT id, text, user_id FROM announcements WHERE id = $1`, id).
		Scan(&a.ID, &a.Text, &a.UserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return announcement.Announcement{}, NotFoundAnnouncement
		}
		return announcement.Announcement{}, err
	}
	return a, nil
}
