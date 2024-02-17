package beer

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type UseCase interface {
	GetAll() ([]*Beer, error)
	Get(ID int64) (*Beer, error)
	Store(b *Beer) error
	Update(b *Beer) error
	Remove(ID int64) error
}

type Service struct {
	DB *sql.DB
}

// age como uma factory que instancia um service e retorna o endereço
func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll() ([]*Beer, error) {

	var result []*Beer

	// vamos usar sempre a conexáo que está dentro do service
	rows, err := s.DB.Query("SELECT * FROM beer")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var b Beer
		err := rows.Scan(&b.ID, &b.Name, &b.Type, &b.Style)

		if err != nil {
			return nil, err
		}

		result = append(result, &b)
	}

	return result, nil
}

func (s *Service) Get(ID int64) (*Beer, error) {
	query, err := s.DB.Prepare("SELECT * FROM beer WHERE id=?")
	if err != nil {
		return nil, err
	}

	defer query.Close()

	var beer Beer
	err = query.QueryRow(ID).Scan(&beer.ID, &beer.Name, &beer.Style, &beer.Type)

	if err != nil {
		return nil, err
	}

	return &beer, nil
}
func (s *Service) Store(b *Beer) error {

	tx, err := s.DB.Begin()

	if err != nil {
		return err
	}

	statement, err := tx.Prepare("INSERT INTO beer(id, name, type, style) values (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(b.ID, b.Name, b.Style, b.Type)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *Service) Update(b *Beer) error {
	if b.ID == 0 {
		return fmt.Errorf("Invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("update beer set name=?, type=?, style=? where id=?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(b.Name, b.Type, b.Style, b.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) Remove(ID int64) error {
	if ID == 0 {
		return fmt.Errorf("Invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM beer WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(ID)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
