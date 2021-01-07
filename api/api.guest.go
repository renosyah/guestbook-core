package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/renosyah/guestbook-core/model"
	uuid "github.com/satori/go.uuid"
)

type (
	GuestModule struct {
		db   *sql.DB
		Name string
	}
)

func NewGuestModule(db *sql.DB) *GuestModule {
	return &GuestModule{
		db:   db,
		Name: "module/Guest",
	}
}

func (m GuestModule) All(ctx context.Context, param model.ListQuery) ([]model.GuestResponse, *Error) {
	var all []model.GuestResponse

	data, err := (&model.Guest{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all Guest"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Guest found"
		}
		return []model.GuestResponse{}, NewErrorWrap(err, m.Name, "all/Guest",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m GuestModule) Add(ctx context.Context, param model.Guest) (model.GuestResponse, *Error) {

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Guest"

		return model.GuestResponse{}, NewErrorWrap(err, m.Name, "add/Guest",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m GuestModule) One(ctx context.Context, param model.Guest) (model.GuestResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Guest"

		return model.GuestResponse{}, NewErrorWrap(err, m.Name, "one/Guest",
			message, status)
	}

	return data.Response(), nil
}

func (m GuestModule) Update(ctx context.Context, param model.Guest) (model.GuestResponse, *Error) {
	var emptyUUID uuid.UUID

	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update Guest"

		return model.GuestResponse{}, NewErrorWrap(err, m.Name, "update/Guest",
			message, status)
	}
	return param.Response(), nil
}

func (m GuestModule) Delete(ctx context.Context, param model.Guest) (model.GuestResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete Guest"

		return model.GuestResponse{}, NewErrorWrap(err, m.Name, "delete/Guest",
			message, status)
	}
	return param.Response(), nil
}
