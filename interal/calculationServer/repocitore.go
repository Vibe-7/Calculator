package calculationserver

import (
	"fmt"

	"gorm.io/gorm"
)

type CalculationReository interface {
	CreateCalculation(cacl Calculation) error
	GetCalculation() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(cacl Calculation) error
	DeleteCalculation(id string) error
}

type CaclRepository struct {
	db *gorm.DB
}

func NewCalculationRepository(db *gorm.DB) CalculationReository {
	return &CaclRepository{db: db}

}

func (r *CaclRepository) CreateCalculation(cacl Calculation) error {
	if cacl.Expresion == "" {
		return fmt.Errorf("expression cannot be empty")
	}
	return r.db.Create(&cacl).Error
}
func (r *CaclRepository) GetCalculation() ([]Calculation, error) {
	var Calc []Calculation
	err := r.db.Find(&Calc).Error
	return Calc, err
}

func (r *CaclRepository) GetCalculationByID(id string) (Calculation, error) {
	var Calc Calculation
	err := r.db.First(&Calc, "id = ?", id).Error
	return Calc, err
}

func (r *CaclRepository) UpdateCalculation(cacl Calculation) error {
	return r.db.Save(&cacl).Error
}

func (r *CaclRepository) DeleteCalculation(id string) error {
	if id == "" {
		return fmt.Errorf("id cannot be empty")
	}
	return r.db.Delete(&Calculation{}, "id = ?", id).Error
}
