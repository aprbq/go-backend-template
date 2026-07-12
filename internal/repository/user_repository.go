package repository

import (
	"go-backend-template/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryDB struct {
	db *mongo.Database
}

func NewUserRepositoryDB(db *mongo.Database) UserRepository {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) Create(user *model.User) error {
	return nil
}

func (r userRepositoryDB) FetchAll() ([]model.User, error) {
	return nil, nil
}

func (r userRepositoryDB) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}

func (r userRepositoryDB) FindByID(id uint) (*model.User, error) {
	return nil, nil
}

func (r userRepositoryDB) Update(user *model.User) error {
	return nil
}

func (r userRepositoryDB) Delete(user *model.User) error {
	return nil
}
