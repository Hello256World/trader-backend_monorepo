package strategy

import (
	"net/http"
	"strings"

	strategyUC "github.com/Hello256World/trader-backend_monorepo/internal/usecases/strategy"
	"github.com/Hello256World/trader-backend_monorepo/pkg/apierrors"
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
}

type handlers struct {
	svc strategyUC.Service
}

func NewHandlers(svc strategyUC.Service) Handlers {
	return &handlers{
		svc: svc,
	}
}

func (h *handlers) Create(c *gin.Context) {
	var req CreateRequest

	if err := c.ShouldBind(&req); err != nil {
		apiErr := apierrors.NewBadRequestError("invalid json body")
		c.AbortWithStatusJSON(apiErr.StatusCode(), apiErr)
		return
	}

	result, err := h.svc.Create(c.Request.Context(), &strategyUC.CreateRequest{
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

func (h *handlers) GetByID(c *gin.Context) {
	id := c.Param("strategy-id")

	if id = strings.TrimSpace(id); id == "" {
		apiErr := apierrors.NewBadRequestError("invalid id")
		c.AbortWithStatusJSON(apiErr.StatusCode(), apiErr)
		return
	}

	result, err := h.svc.GetByID(c, &strategyUC.GetByIDRequest{ID: id})

	if err != nil {
		apiErr := apierrors.FromError(err)
		c.AbortWithStatusJSON(apiErr.StatusCode(), apiErr)
		return
	}

	res := GetByIDResponse{Strategy: fromStrategyCoreToHTTP(result.Strategy)}

	c.JSON(http.StatusOK, res)
}
