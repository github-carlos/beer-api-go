package handlers

import (
	"beer-api/core/beer"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeHandlers(s beer.UseCase, r *mux.Router, n *negroni.Negroni) {
	r.Handle("/v1/beer", n.With(
		negroni.Wrap(getAll(s)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/v1/beer/{id}", n.With(
		negroni.Wrap(getOne(s)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/v1/beer", n.With(
		negroni.Wrap(createBeer(s)),
	)).Methods("POST")

	r.Handle("/v1/beer/{id}", n.With(
		negroni.Wrap(updateBeer(s)),
	)).Methods("PUT")

	r.Handle("/v1/beer/{id}", n.With(
		negroni.Wrap(deleteBeer(s)),
	)).Methods("DELETE")
}

func getAll(s beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		beers, err := s.GetAll()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(HandleError(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(beers)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(HandleError("Erro convertendo em JSON"))
			return
		}
	})
}

func getOne(s beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
		}

		beer, err := s.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(HandleError(err.Error()))
			return
		}

		json.NewEncoder(w).Encode(beer)
	})
}

func createBeer(s beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		r.Header.Set("Content-Language", "application/json")
		var beer beer.Beer

		err := json.NewDecoder(r.Body).Decode(&beer)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(HandleError(err.Error()))
			return
		}

		err = s.Store(&beer)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(HandleError(err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)

	})
}

func updateBeer(s beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var b beer.Beer
		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(HandleError(err.Error()))
			return
		}

		b.ID = id

		oldBeer, err := s.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(HandleError(err.Error()))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&b)

		if b.Name == "" {
			b.Name = oldBeer.Name
		}

		if b.Style == 0 {
			b.Style = oldBeer.Style
		}

		if b.Type == 0 {
			b.Type = oldBeer.Type
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(HandleError(err.Error()))
			return
		}
		err = s.Update(&b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(HandleError(err.Error()))
			return
		}
	})
}

func deleteBeer(s beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(HandleError(err.Error()))
			return
		}
		_, err = s.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(HandleError(err.Error()))
			return
		}

		err = s.Remove(id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(HandleError(err.Error()))
			return
		}
	})
}
