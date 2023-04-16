package store

import "github.com/subramanian0331/familytree/models"

type PostgresDB struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (p *PostgresDB) AddUser(models.User) error                  { return nil }
func (p *PostgresDB) DeleteUser(id string) (*models.User, error) { return nil, nil }
func (p *PostgresDB) GetUser(email string) (*models.User, error) { return nil, nil }
func (p *PostgresDB) UpdateUser(models.User) error               { return nil }
