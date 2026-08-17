package domain

type Signal struct {
	ID         string `json:"id"`
	Symbol     string `json:"symbol"`
	Side       string `json:"side"`
	Quantity   int    `json:"quantity"`
	StrategyID string `json:"strategy_id"`
}
