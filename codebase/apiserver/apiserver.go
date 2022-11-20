package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/joekesov/go-rss-reader-package"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var defaultStopTimeout = time.Second * 30

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) (*APIServer, error) {
	if addr == "" {
		return nil, errors.New("addr cannot be blank")
	}

	return &APIServer{
		addr: addr,
	}, nil
}

// Start starts a server with a stop channel
func (s *APIServer) Start(stop <-chan struct{}) error {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.router(),
	}

	go func() {
		logrus.WithField("addr", srv.Addr).Info("starting server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), defaultStopTimeout)
	defer cancel()

	logrus.WithField("timeout", defaultStopTimeout).Info("stopping server")
	return srv.Shutdown(ctx)
}

func (s *APIServer) router() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", s.defaultRoute)
	router.HandleFunc("/rss", s.otherRoute).Methods("POST")
	return router
}

func (s *APIServer) defaultRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World maraba"))
}

func (s *APIServer) otherRoute(w http.ResponseWriter, r *http.Request) {
	var us urls
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&us); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %s", err))
		return
	}
	defer r.Body.Close()

	rssItems, err := jreader.Parse(us.Urls)
	if err != nil {

	}

	fmt.Println(len(rssItems))
	respondWithJSON(w, http.StatusCreated, rssItems)
}
