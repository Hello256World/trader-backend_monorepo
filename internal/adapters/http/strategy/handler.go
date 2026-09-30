package strategy

import (
	"net/http"

	strategyUC "github.com/Hello256World/trader-backend_monorepo/internal/usecases/strategy"
	"github.com/Hello256World/trader-backend_monorepo/pkg/apierrors"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
}

type handler struct {
	svc strategyUC.Service
}

func NewHandler(svc strategyUC.Service) Handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) Create(c *gin.Context) {
	var req CreateRequest

	if err := c.ShouldBind(&req); err != nil {
		apiErr := apierrors.NewBadRequestError("invalid json body")
		c.AbortWithStatusJSON(apiErr.StatusCode(), apiErr)
		return
	}

	result, err := h.svc.Create(c, &strategyUC.CreateRequest{
		Name:       req.Name,
		Descrption: req.Description,
	})

	if err != nil {
		apiErr := apierrors.FromError(err)
		c.AbortWithStatusJSON(apiErr.StatusCode(), apiErr)
		return
	}

	c.JSON(http.StatusCreated, CreateResponse{ID: result.StrategyID})
}

func (h *handler) GetByID(c *gin.Context) {
	var req GetByIDRequest

	if err := c.ShouldBind(&req); err != nil {
		apiErr := apierrors.NewBadRequestError("error getting strategy by id")
		c.AbortWithStatusJSON(apiErr.StatusCode(), apiErr)
		return
	}

	result, err := h.svc.GetByID(c, &strategyUC.GetByIDRequest{ID: req.ID})

	if err != nil {
		apiErr := apierrors.FromError(err)
		c.AbortWithStatusJSON(apiErr.StatusCode(), apiErr)
		return
	}

	res := GetByIDResponse{Strategy: fromStrategyCoreToHTTP(result.Strategy)}

	c.JSON(http.StatusOK, res)
}
