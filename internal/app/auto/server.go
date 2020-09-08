package auto

import (
	"encoding/json"
	"net/http"

	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/base62"
	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/models"
	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/store"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ctxKey int8

const (
	ctxKeyUser ctxKey = iota
	ctxKeyRequestID
)

const (
	currentHost string = "192.168.99.106:3000/"
)

// Server structure
type server struct {
	logger *logrus.Logger
	router *mux.Router
	store  store.Store
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(store store.Store) *server {
	s := &server{
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store,
	}
	s.configureRouter()
	return s
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.Use(s.accessLogMiddleware)
	s.router.Use(s.panicMiddleware)
	s.router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, map[string]bool{"ok": true})
	})
	s.router.HandleFunc("/api/shorten_url", s.handleShortenURL()).Methods("POST")
	s.router.HandleFunc("/{id:[0-9A-Za-z]{0,22}}", s.handleRedirectBase62()).Methods("GET")
}

func (s *server) handleShortenURL() http.HandlerFunc {
	type request struct {
		URL        string `json:"url"`
		ShortenURL string `json:"shorten_url"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		l := &models.Links{
			URL:        req.URL,
			ShortenURL: req.ShortenURL,
		}
		if err := s.store.Links().Create(l); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}
		s.respond(w, r, http.StatusOK, currentHost+l.ShortenURL)
	}
}

func (s *server) handleRedirectBase62() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := base62.Decode(vars["id"])
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		l, err := s.store.Links().FindByID(id)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		http.Redirect(w, r, l.URL, http.StatusMovedPermanently)
	}
}
