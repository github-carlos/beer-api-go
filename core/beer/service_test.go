package beer_test

import (
	"beer-api/core/beer"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	b := &beer.Beer{
		ID:    1,
		Name:  "Patagonia",
		Type:  beer.TypeLarger,
		Style: beer.StylePale,
	}

	db, err := sql.Open("sqlite3", "../../data/beer_test.db")

	defer clearAndClose(db, t)

	if err != nil {
		t.Fatalf("Error trying to open database %s", err.Error())
	}

	service := beer.NewService(db)
	err = service.Store(b)

	if err != nil {
		t.Fatalf("Error trying to store new Beer %s", err.Error())
	}

	saved, err := service.Get(1)

	if err != nil {
		t.Fatalf("Error getting one beer %s", err.Error())
	}

	if saved.ID != 1 {
		t.Fatalf("Found item is not expected one 1 != %d", saved.ID)
	}

	allBeer, err := service.GetAll()

	if err != nil {
		t.Fatalf("Error trying to get all Beer %s", err.Error())
	}

	if len(allBeer) != 1 {
		t.Fatalf("Not found all beers expected 1 found %d", len(allBeer))
	}

}

func clearAndClose(db *sql.DB, t *testing.T) {
	tx, err := db.Begin()
	assert.Nil(t, err)
	if err != nil {
		t.Fatalf("Errro criando transação %s", err.Error())
	}
	_, err = tx.Exec("delete from beer")
	assert.Nil(t, err)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	db.Close()
}
