package store

import "github.com/subramanian0331/familytree/models"

type Storage interface {
	AddMember(member *models.Member) error
	AddRelationship(member1 string, member2 string, Relationship models.Relationship, family string) error
	GetMemberDetails(memberName string, family string) (*models.Member, error)
	UpdateMemberDetails(member *models.Member, family string) error
	GetParents(memberName string, family string) ([]*models.Member, error)
	GetChildren(memberName string, family string) ([]*models.Member, error)
	GetAllFamilyMembers(family string) ([]*models.Member, error)
}

func NewStorage(host string) Storage {
	return &RedisGraphDB{host: host}
}

type UserStorage interface {
	AddUser(models.User) error
	DeleteUser(Id int) error
	GetUser(id string) (*models.User, error)
}

func NewUserStorage(host string, port string, username string, password string, dbname string) UserStorage {
	return &PostgresDB{
		Host:     host,
		Port:     port,
		User:     username,
		Password: password,
		DBName:   dbname,
	}
}
