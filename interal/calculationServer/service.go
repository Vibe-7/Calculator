package calculationserver

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

type CalculationService interface {
	CreateCalculation(expression string) (Calculation, error)
	GetAllCalculation() ([]Calculation, error)
	GetCalculationID(id string) (Calculation, error)
	UpdateCalculation(id, expression string) (Calculation, error)
	DeleteCalculetion(id string) error
}

type caclService struct {
	repo CalculationReository
}

func NewCalculationService(r CalculationReository) CalculationService {
	return &caclService{repo: r}
}

func (s *caclService) calculationExpresion(expresion string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expresion)
	if err != nil {
		return "", err
	}
	result, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", result), nil
}

func (s *caclService) CreateCalculation(expression string) (Calculation, error) {
	result, err := s.calculationExpresion(expression)
	if err != nil {
		return Calculation{}, err
	}
	cacl := Calculation{
		ID:        uuid.NewString(),
		Expresion: expression,
		Result:    result,
	}

	if err := s.repo.CreateCalculation(cacl); err != nil {
		return Calculation{}, err
	}

	return cacl, nil
}

func (s *caclService) GetAllCalculation() ([]Calculation, error) {
	return s.repo.GetCalculation()
}

func (s *caclService) GetCalculationID(id string) (Calculation, error) {
	return s.repo.GetCalculationByID(id)
}

func (s *caclService) UpdateCalculation(id string, expression string) (Calculation, error) {
	cacl, err := s.repo.GetCalculationByID(id)
	if err != nil {
		return Calculation{}, err
	}
	result, err := s.calculationExpresion(expression)
	if err != nil {
		return Calculation{}, err
	}
	cacl.Expresion = expression
	cacl.Result = result
	if err := s.repo.UpdateCalculation(cacl); err != nil {
		return Calculation{}, err
	}

	return cacl, nil
}
func (s *caclService) DeleteCalculetion(id string) error {
	return s.repo.DeleteCalculation(id)
}
