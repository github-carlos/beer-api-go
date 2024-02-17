package handlers

import (
	"beer-api/core/beer"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllBeersIntegration(t *testing.T) {

	b1 := &beer.Beer{
		ID:    1,
		Name:  "Beer 1",
		Type:  1,
		Style: 1,
	}
	b2 := &beer.Beer{
		ID:    2,
		Name:  "Beer 2",
		Type:  2,
		Style: 2,
	}
	// making connection
	sql, err := sql.Open("sqlite3", "../../data/beer_test.db")
	assert.Nil(t, err)

	service := beer.NewService(sql)

	clearDB(sql)
	// creating data to use as test
	assert.Nil(t, service.Store(b1))
	assert.Nil(t, service.Store(b2))

	handler := getAll(service)
	r := mux.NewRouter()
	r.Handle("/v1/beer", handler).Methods("GET")
	req, err := http.NewRequest("GET", "/v1/beer", nil)
	assert.Nil(t, err)
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var beers []*beer.Beer
	err = json.NewDecoder(rr.Body).Decode(&beers)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(beers))
	assert.Equal(t, b1.ID, beers[0].ID)
	assert.Equal(t, b2.ID, beers[1].ID)
}

func clearDB(sql *sql.DB) {
	tx, err := sql.Begin()
	if err != nil {
		log.Fatal(err.Error())
	}
	tx.Exec("delete from beer")
	tx.Commit()
}

type ServiceMock struct{}

func (s *ServiceMock) Get(id int64) (*beer.Beer, error) {
	return &beer.Beer{
		ID:    1,
		Name:  "Heineken",
		Type:  1,
		Style: 2,
	}, nil
}

func (s *ServiceMock) GetAll() ([]*beer.Beer, error) {
	return []*beer.Beer{
		&beer.Beer{ID: 1, Name: "Heineken", Type: 1, Style: 2},
		&beer.Beer{ID: 2, Name: "Original", Type: 2, Style: 1},
	}, nil
}

func (s *ServiceMock) Store(b *beer.Beer) error {
	return nil
}

func (s *ServiceMock) Remove(id int64) error {
	return nil
}

func (s *ServiceMock) Update(b *beer.Beer) error {
	return nil
}

func Test_GetAllWithMock(t *testing.T) {

	s := &ServiceMock{}
	handler := getAll(s)

	r := mux.NewRouter()
	r.Handle("/v1/beer", handler)

	req, err := http.NewRequest("GET", "/v1/beer", nil)
	assert.Nil(t, err)
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var beers []*beer.Beer
	err = json.NewDecoder(rr.Body).Decode(&beers)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(beers))
	assert.Equal(t, int64(1), beers[0].ID)
	assert.Equal(t, int64(2), beers[1].ID)
}
