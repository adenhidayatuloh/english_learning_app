package pakarpostgres

import (
	"english_app/internal/user_module/entity"
	"english_app/pkg/errs"
	"errors"

	Pakarrepository "english_app/internal/user_module/repository/pakar_repository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type pakarPostgres struct {
	db *gorm.DB
}

// UpdatePakar implements Pakarrepository.PakarRepository.
func (s *pakarPostgres) UpdatePakar(oldPakar *entity.Pakar, newPakar *entity.Pakar) (*entity.Pakar, errs.MessageErr) {
	if err := s.db.Model(oldPakar).Updates(newPakar).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewUnprocessableEntity(fmt.Sprintf("Failed to update email %s", oldPakar.Email))
	}

	return oldPakar, nil
}

// CreatePakar implements Pakarrepository.PakarRepository.
func (s *pakarPostgres) CreatePakar(Pakar *entity.Pakar) (*entity.Pakar, errs.MessageErr) {

	if err := s.db.Create(Pakar).Error; err != nil {
		log.Println("Error:", err.Error())

		return nil, errs.NewInternalServerError("Failed to create Pakar ")
	}

	return Pakar, nil
}

func NewpakarMysql(db *gorm.DB) Pakarrepository.PakarRepository {
	return &pakarPostgres{db}
}

// GetPakarByEmail implements Pakarrepository.PakarRepository.
func (s *pakarPostgres) GetPakarByEmail(email string) (*entity.Pakar, errs.MessageErr) {
	var Pakar entity.Pakar

	if err := s.db.First(&Pakar, "email = ?", email).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle record not found error...

			return nil, errs.NewNotFound(fmt.Sprintf("Pakar with email %s is not found", email))
		}
		return nil, errs.NewBadRequest("cannot get Pakar")
	}

	return &Pakar, nil
}
