package Repositories

import (
	"database/sql"
	"enCodeTest/src/models"
)

type PersonRepository interface {
	GetByID(id int) (*models.Person, error)
	GetAll(limit, offset int, search string) ([]models.Person, error)
	Create(person *models.Person) error
	Update(person *models.Person) error
	Delete(id int) error
}

type PostgresRepository struct {
	db *sql.DB
}

// Create New REPO
func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db}

}

func (r *PostgresRepository) Create(person *models.Person) error {
	query := `INSERT INTO persons (email, phone, first_name, last_name) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, person.Email, person.Phone, person.FirstName, person.LastName)
	return err
}

func (r *PostgresRepository) GetByID(id int) (*models.Person, error) {
	var person models.Person
	query := `SELECT id, email, phone, first_name, last_name FROM persons WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&person.ID, &person.Email, &person.Phone, &person.FirstName, &person.LastName)
	return &person, err
}

func (r *PostgresRepository) GetAll(limit, offset int, search string) ([]models.Person, error) {
	var persons []models.Person
	query := "SELECT id, email, phone, first_name, last_name FROM persons"
	//Если есть дополнительный критерий по поиску (имя или фамилия)

	if search != "" {
		query += " WHERE first_name LIKE $1 OR last_name LIKE $2"
	}

	query += " LIMIT $3 OFFSET $4"
	var rows *sql.Rows
	var err error
	if search != "" {
		// Выполняем запрос с поисковым параметром
		rows, err = r.db.Query(query, "%"+search+"%", "%"+search+"%", limit, offset)
	} else {
		rows, err = r.db.Query(query, limit, offset)

	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var person models.Person
		err := rows.Scan(&person.ID, &person.Email, &person.Phone, &person.FirstName, &person.LastName)
		if err != nil {
			return nil, err
		}
		persons = append(persons, person)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	//log.Println(persons)
	return persons, nil
}
func (r *PostgresRepository) Update(person *models.Person) error {
	query := `UPDATE persons SET email = $1, phone = $2, first_name = $3, last_name = $4 WHERE id = $5`
	_, err := r.db.Exec(query, person.Email, person.Phone, person.FirstName, person.LastName, person.ID)
	return err
}

func (r *PostgresRepository) Delete(id int) error {
	query := `DELETE FROM persons WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
