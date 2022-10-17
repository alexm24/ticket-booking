package route

import (
	"encoding/json"
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"

	"github.com/alexm24/ticket-booking/internal/models"
	"github.com/alexm24/ticket-booking/internal/service"
)

type Route struct {
	service *service.Service
}

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func newErrorResponse(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(Error{Code: int32(code), Message: msg})
}

func NewRoute(service *service.Service) *Route {
	return &Route{service}
}

func (c *Route) AddRoute(w http.ResponseWriter, r *http.Request) {
	var item models.PostRoute
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "error post route")
		return
	}

	i, err := c.service.IRoutes.CreateRoute(item)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "error service")
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(i)
}

func (c *Route) AddPassenger(w http.ResponseWriter, r *http.Request, routeId openapi_types.UUID) {
	var item models.PostPassenger
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "error post passenger")
		return
	}

	i, err := c.service.IPassengers.CreatePassenger(routeId, item)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "error service")
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(i)
}
