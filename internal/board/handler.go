package board

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

func CreateAnnouncementHandler(b *Board) func(w http.ResponseWriter, r *http.Request) {
	if b == nil {
		panic("nil board service")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := uuid.Parse(r.Context().Value("id").(string))
		if err != nil {
			slog.Error("user id not parse", slog.String("user_id", r.Context().Value("id").(string)))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		txtReq := make(map[string]string)

		err = json.Unmarshal(body, &txtReq)
		if err != nil {
			return
		}

		defer r.Body.Close()

		if text, ok := txtReq["text"]; ok {
			err = b.Add(userID, text)
			if err != nil {
				slog.Error("add error", slog.String("user_id", userID.String()))
				return
			}

			w.WriteHeader(http.StatusCreated)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
	}
}

func UpdateAnnouncementHandler(b *Board) func(w http.ResponseWriter, r *http.Request) {
	if b == nil {
		panic("nil board service")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := uuid.Parse(r.Context().Value("id").(string))
		if err != nil {
			slog.Error("user id not parse", slog.String("user_id", r.Context().Value("id").(string)))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		var req struct {
			ID   string `json:"id"`
			Text string `json:"text"`
		}
		if err := json.Unmarshal(body, &req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := uuid.Parse(req.ID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = b.Update(id, req.Text)
		if err != nil {
			if errors.Is(err, NotFoundAnnouncement) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			slog.Error("update error", slog.String("user_id", userID.String()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteAnnouncementHandler(b *Board) func(w http.ResponseWriter, r *http.Request) {
	if b == nil {
		panic("nil board service")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := uuid.Parse(r.Context().Value("id").(string))
		if err != nil {
			slog.Error("user id not parse", slog.String("user_id", r.Context().Value("id").(string)))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var req struct {
			ID string `json:"id"`
		}
		if err := json.Unmarshal(body, &req); err != nil || req.ID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := uuid.Parse(req.ID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = b.Delete(id)
		if err != nil {
			if errors.Is(err, NotFoundAnnouncement) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			slog.Error("delete error", slog.String("err", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func ListHandler(b *Board) func(w http.ResponseWriter, r *http.Request) {
	if b == nil {
		panic("nil board service")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		l := b.List()

		data, err := json.Marshal(l)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	}
}
