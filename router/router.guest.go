package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/guestbook-core/api"
	"github.com/renosyah/guestbook-core/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddGuest(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Guest

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Guest/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return guestModule.Add(ctx, param)
}

func HandlerAllGuest(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.ListQuery

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Guest/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return guestModule.All(ctx, param)
}

func HandlerOneGuest(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Guest/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return guestModule.One(ctx, model.Guest{ID: id})
}

func HandlerUpdateGuest(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.Guest

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Guest/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Guest/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return guestModule.Update(ctx, param)
}

func HandlerDeleteGuest(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Guest/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return guestModule.Delete(ctx, model.Guest{ID: id})
}
