package user

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-tutorial/internal/apperror"
	"rest-api-tutorial/internal/handlers"
	"rest-api-tutorial/pkg/logging"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetAllUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))
}

func (h *handler) GetAllUser(w http.ResponseWriter, r *http.Request) error {
	//w.WriteHeader(200)
	//w.Write([]byte(fmt.Sprintf("this is list of users")))
	return apperror.ErrNotFound
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, request *http.Request) error {
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("this is user by UUID")))
	return nil
}

func (h *handler) CreateUser(w http.ResponseWriter, request *http.Request) error {
	//w.WriteHeader(201)
	//w.Write([]byte(fmt.Sprintf("this is create user")))
	return fmt.Errorf("this is API error")
}

func (h *handler) UpdateUser(w http.ResponseWriter, request *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte(fmt.Sprintf("this is updating user")))
	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, request *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte(fmt.Sprintf("this is partially updating user")))
	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, request *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte(fmt.Sprintf("this is delete user")))
	return nil
}
