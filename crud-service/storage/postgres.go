package storage

import (
	_ "github.com/lib/pq"
	"github.com/vsPEach/HelmCRUD/crud-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(dsn string) (*Storage, error) {
	connect, err := gorm.Open(postgres.Open(dsn), nil)
	if err != nil {
		return nil, err
	}
	err = connect.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return &Storage{connect}, nil
}

func (s *Storage) Create(user models.User) (int64, error) {
	tx := s.db.Create(&user)
	return tx.RowsAffected, tx.Error
}

func (s *Storage) Update(user models.User) (int64, error) {
	tx := s.db.Save(&user)
	return tx.RowsAffected, tx.Error
}

func (s *Storage) Delete(id uint64) error {
	return s.db.Delete(models.User{}, id).Error
}

func (s *Storage) Get(id uint64) (models.User, error) {
	var user models.User
	tx := s.db.First(&user, id)
	return user, tx.Error
}
