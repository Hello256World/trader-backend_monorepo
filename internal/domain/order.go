package domain

type Order struct {
	Symbol     string 
	ID         string
	Side       string  
	Quantity   int    
	StrategyID string  
	Price      float64 
	Status     string 
}
