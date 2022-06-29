package rest

import (
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/runtime/middleware"

	"github.com/dsegundo2/service-management/internal/servicedb"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	db         *servicedb.ServiceDB
	serverAddr string
	router     *mux.Router
	log        *logrus.Entry
}

func New(address string, db *servicedb.ServiceDB, logger *logrus.Entry) *Server {
	return &Server{
		db:         db,
		serverAddr: address,
		router:     mux.NewRouter(),
		log:        logger,
	}
}

func (s *Server) Listen() {
	s.router.HandleFunc("/v1/services", s.HandleCreateService).Methods(http.MethodPost)
	s.router.HandleFunc("/v1/services/{service_id}", s.HandleReadService).Methods(http.MethodGet)
	s.router.HandleFunc("/v1/services", s.HandleListServices).Methods(http.MethodGet)
	s.router.HandleFunc("/v1/services", s.HandleUpdateService).Methods(http.MethodPut)
	s.router.HandleFunc("/v1/services/{service_id}:addVersion", s.HandleAddServiceVersion).Methods(http.MethodPost)
	s.router.HandleFunc("/v1/services/{service_id}:removeVersion", s.HandleRemoveServiceVersion).Methods(http.MethodPost)
	s.router.HandleFunc("/v1/services/{service_id}", s.HandleDeleteService).Methods(http.MethodDelete)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml", BasePath: "/v1"}
	sh := middleware.Redoc(opts, nil)

	s.router.Handle("/doc", sh)
	s.router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	srv := &http.Server{
		Handler:      s.router,
		Addr:         s.serverAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s.log.WithField("Address", s.serverAddr).Info("Listening for requests")
	log.Fatal(srv.ListenAndServe())
}
