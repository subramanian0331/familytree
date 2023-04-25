package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/subramanian0331/familytree/models"
)

type PostgresDB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (p *PostgresDB) createConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		p.Host, p.Port, p.User, p.Password, p.DBName)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Error connecting to database: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Println("Error pinging database: ", err.Error())
	}

	return db
}

func (p *PostgresDB) AddUser(usr models.User) error {
	// create new user in database
	db := p.createConnection()
	defer db.Close()

	GetQuery := `SELECT id FROM user_profile WHERE email = $1 `
	row := db.QueryRow(GetQuery, usr.Email)
	err := row.Scan()
	if err != sql.ErrNoRows {
		return fmt.Errorf("user exists")
	}

	Query := `INSERT INTO user_profile (firstname, lastname, nickname, email, password_hash) VALUES($1, $2, $3, $4, $5) RETURNING id`
	var id int64
	err = db.QueryRow(Query, usr.Firstname, usr.Lastname, usr.Nickname, usr.Email, usr.PassHash).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Errorf("db insertion failed", err)
	}

	tStr := usr.UserMetaData.CreatedAt.Format("2006-01-02 15:04:05")
	Query2 := `INSERT INTO user_metadata (user_id, created_at, role, gender) VALUES($1, $2, $3, $4) RETURNING id`
	err = db.QueryRow(Query2, id, tStr, usr.UserMetaData.Role, usr.UserMetaData.Gender.String()).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Errorf("db insertion failed", err)
	}
	return nil
}
func (p *PostgresDB) DeleteUser(id int) error {
	db := p.createConnection()
	defer db.Close()

	GetQuery := `SELECT id FROM user_profile WHERE id = $1 `
	row := db.QueryRow(GetQuery, id)
	err := row.Scan()
	if err == sql.ErrNoRows {
		return fmt.Errorf("user missing")
	}

	DelQuery := `DELETE FROM user_profile WHERE id = $1`
	_, err = db.Query(DelQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresDB) GetUser(email string) (*models.User, error) { return nil, nil }
func (p *PostgresDB) UpdateUser(models.User) error               { return nil }
