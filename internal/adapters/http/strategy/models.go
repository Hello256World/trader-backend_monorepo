package strategy

type CreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

type GetByIDRequest struct {
	ID string `json:"id" binding:"required"`
}

type GetByIDResponse struct {
	Strategy Strategy `json:"strategy"`
}

type Strategy struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
