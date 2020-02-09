package api

import (
	"encoding/json"
	"net/http"

	"github.com/aca/go-restapi-boilerplate/ent"
	"github.com/aca/go-restapi-boilerplate/ent/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func (s *server) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Ctx(ctx).Debug().Msg("CreateUser")

	u := &ent.User{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		render.Render(w, r, ErrInvalidRequest(r, err))
		return
	}

	u, err := s.db.User.Create().SetUserID(u.UserID).SetUserName(u.UserName).Save(ctx)
	if err != nil {
		render.Render(w, r, ErrServerError(r, err))
		return
	}
	render.JSON(w, r, u)
}

func (s *server) ReadUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Ctx(ctx).Debug().Msg("ReadUser")
	uid := chi.URLParam(r, "userID")
	u, err := s.db.User.Query().Where(user.UserIDEQ(uid)).Only(ctx)
	if err != nil {
		render.Render(w, r, ErrServerError(r, err))
		return
	}
	render.JSON(w, r, u)
	return
}

func (s *server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Ctx(ctx).Debug().Msg("DeleteUser")
	uid := chi.URLParam(r, "userID")
	_, err := s.db.User.Delete().Where(user.UserIDEQ(uid)).Exec(ctx)
	if err != nil {
		render.Render(w, r, ErrServerError(r, err))
		return
	}
	return
}

func (s *server) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Ctx(ctx).Debug().Msg("PatchUser")

	u := &ent.User{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		render.Render(w, r, ErrInvalidRequest(r, err))
	}

	uid := chi.URLParam(r, "userID")
	_, err := s.db.User.Update().Where(user.UserIDEQ(uid)).SetUserName(u.UserName).Save(ctx)
	if err != nil {
		render.Render(w, r, ErrServerError(r, err))
		return
	}
	render.JSON(w, r, u)
	return
}
