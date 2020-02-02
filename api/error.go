package api

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/hlog"
)

type ErrResponse struct {
	Msg            string `json:"msg"`        // err string
	RequestID      string `json:"request_id"` // request id to track errors
	HTTPStatusCode int    `json:"-"`          // http response status code
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(r *http.Request, err error) render.Renderer {
	hlog.FromRequest(r).Error().Err(err).Send()
	reqID, _ := hlog.IDFromRequest(r)
	return &ErrResponse{
		HTTPStatusCode: http.StatusBadRequest,
		Msg:            err.Error(),
		RequestID:      reqID.String(),
	}
}

func ErrNotFound(r *http.Request, err error) render.Renderer {
	hlog.FromRequest(r).Error().Err(err).Send()
	reqID, _ := hlog.IDFromRequest(r)
	return &ErrResponse{
		HTTPStatusCode: http.StatusNotFound,
		Msg:            err.Error(),
		RequestID:      reqID.String(),
	}
}

func ErrUnauthorized(r *http.Request, err error) render.Renderer {
	hlog.FromRequest(r).Error().Err(err).Send()
	reqID, _ := hlog.IDFromRequest(r)
	return &ErrResponse{
		HTTPStatusCode: http.StatusUnauthorized,
		Msg:            err.Error(),
		RequestID:      reqID.String(),
	}
}

func ErrServerError(r *http.Request, err error) render.Renderer {
	hlog.FromRequest(r).Error().Err(err).Send()
	reqID, _ := hlog.IDFromRequest(r)
	return &ErrResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		Msg:            err.Error(),
		RequestID:      reqID.String(),
	}
}
