package calculationserver

type Calculation struct {
	ID        string `gorm:"primarykey" json:"id"`
	Expresion string `json:"expression"`
	Result    string `json:"result"`
}

type CalculationRequest struct {
	Expresion string `json:"expression"`
}
