package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dsegundo2/service-management/internal/rest/models"
	db "github.com/dsegundo2/service-management/internal/servicedb"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	ErrMissingRequestData = errors.New("missing required field")
)

func (s *Server) HandleCreateService(w http.ResponseWriter, r *http.Request) {
	service := &db.Service{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&service); err != nil {
		s.writeResponse(w, r, nil, ErrMissingRequestData)
		return
	}
	defer r.Body.Close()

	err := s.db.CreateService(service)
	s.writeResponse(w, r, service, err)
}

func (s *Server) HandleReadService(w http.ResponseWriter, r *http.Request) {
	var err error
	loadParam := r.URL.Query()["load_versions"]
	var loadVersions bool
	if len(loadParam) > 0 {
		loadVersions, err = strconv.ParseBool(loadParam[0])
	}
	if err != nil {
		s.writeResponse(w, r, nil, ErrMissingRequestData)
		return
	}
	path := r.URL.Path
	splitPath := strings.Split(path, "/")
	if len(splitPath) < 3 {
		s.writeResponse(w, r, nil, ErrMissingRequestData)
		return
	}

	service := &db.Service{
		ID: splitPath[3],
	}
	err = s.db.ReadService(service, loadVersions)

	s.writeResponse(w, r, service, err)
}

func (s *Server) HandleListServices(w http.ResponseWriter, r *http.Request) {
	var err error
	filterParam := r.URL.Query()["filter"]
	var filter string
	if len(filterParam) > 0 {
		filter = filterParam[0]
	}
	sortParam := r.URL.Query()["sort"]
	var sort string
	if len(sortParam) > 0 {
		sort = sortParam[0]
	}
	limitParam := r.URL.Query()["limit"]
	var limit int
	if len(limitParam) > 0 {
		limit, err = strconv.Atoi(limitParam[0])
	}
	if err != nil {
		s.writeResponse(w, r, nil, ErrMissingRequestData)
		return
	}
	// Set default for limit
	if limit == 0 {
		limit = 10
	}
	offsetParam := r.URL.Query()["offset"]
	var offset int
	if len(offsetParam) > 0 {
		offset, err = strconv.Atoi(offsetParam[0])
	}
	if err != nil {
		s.writeResponse(w, r, nil, ErrMissingRequestData)
		return
	}
	services, total, err := s.db.ListServices(filter, sort, limit, offset)

	resp := &models.ListServicesResponse{}
	if err == nil {
		resp = models.ConvertListServicesResponse(services, int(total))
	}

	s.writeResponse(w, r, resp, err)
}

func (s *Server) HandleUpdateService(w http.ResponseWriter, r *http.Request) {
	serviceUpdateRequest := &models.UpdateServiceRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&serviceUpdateRequest); err != nil {
		s.writeResponse(w, r, nil, ErrMissingRequestData)
		return
	}
	defer r.Body.Close()

	dbService := &db.Service{
		ID:          serviceUpdateRequest.ServiceId,
		Title:       serviceUpdateRequest.Title,
		Description: serviceUpdateRequest.Description,
	}

	err := s.db.UpdateService(dbService)
	s.writeResponse(w, r, dbService, err)
}

func (s *Server) HandleAddServiceVersion(w http.ResponseWriter, r *http.Request) {
	serviceVersion := &db.ServiceVersion{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&serviceVersion); err != nil {
		s.writeResponse(w, r, nil, ErrMissingRequestData)
		return
	}
	defer r.Body.Close()

	err := s.db.CreateServiceVersion(serviceVersion)
	s.writeResponse(w, r, serviceVersion, err)
}

func (s *Server) HandleRemoveServiceVersion(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) HandleDeleteService(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	splitPath := strings.Split(path, "/")
	if len(splitPath) < 3 {
		s.writeResponse(w, r, nil, ErrMissingRequestData)
		return
	}

	err := s.db.DeleteService(splitPath[3])

	s.writeResponse(w, r, models.Empty{}, err)
}

func (s *Server) writeResponse(w http.ResponseWriter, r *http.Request, payload interface{}, err error) {
	// Match on gorm errors
	s.log.WithError(err).WithFields(logrus.Fields{
		"Request":       fmt.Sprintf("%+v", r),
		"Response Body": payload,
	}).Info("request finished")

	// One off check for duplicate key. TODO: make cleaner with db exported errors
	if err != nil && strings.Contains(err.Error(), "duplicate key") {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("duplicate value. unique value already exists"))
		return
	}

	// Check for error and return correct status code
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		w.WriteHeader(http.StatusNotFound)
	case gorm.ErrNotImplemented:
		w.WriteHeader(http.StatusMethodNotAllowed)
	case ErrMissingRequestData:
		w.WriteHeader(http.StatusBadRequest)
	case gorm.ErrMissingWhereClause:
		w.WriteHeader(http.StatusBadRequest)
	case gorm.ErrPrimaryKeyRequired:
		w.WriteHeader(http.StatusBadRequest)
	case gorm.ErrInvalidData:
		w.WriteHeader(http.StatusBadRequest)
	case gorm.ErrInvalidField:
		w.WriteHeader(http.StatusBadRequest)
	case gorm.ErrEmptySlice:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// Successful response
	response, err := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(response))
}
