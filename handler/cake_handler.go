package handler

import (
	"cake-api/core"
	"cake-api/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprint(`"`, time.Time(t).Format("2006-01-02 15-04-05"), `"`)), nil
}

type CakeResponse struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Rating      float32  `json:"rating"`
	Image       string   `json:"image"`
	CreatedAt   JSONTime `json:"created_at"`
	UpdatedAt   JSONTime `json:"updated_at"`
}

func mapCakeToCakeResponse(cake *core.Cake) *CakeResponse {
	if cake == nil {
		return nil
	}
	return &CakeResponse{
		ID:          cake.ID,
		Title:       cake.Title,
		Description: cake.Description,
		Rating:      cake.Rating,
		Image:       cake.Image,
		CreatedAt:   JSONTime(cake.CreatedAt),
		UpdatedAt:   JSONTime(cake.UpdatedAt),
	}
}

type CakeHandler struct {
	useCase core.CakeUseCase
}

func NewCakeHandler(useCase core.CakeUseCase) *CakeHandler {
	return &CakeHandler{useCase: useCase}
}

func (h *CakeHandler) GetListCakes(c *gin.Context) {
	result, err := h.useCase.GetListCakes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Message{Error: err.Error()})
		return
	}

	var responses []CakeResponse
	for _, c := range result {
		responses = append(responses, *mapCakeToCakeResponse(&c))
	}
	c.JSON(http.StatusOK, response.Message{
		Message: "cake list",
		Data:    responses,
	})
}

func (h *CakeHandler) GetCakeDetail(c *gin.Context) {
	cakeID := StrToUint(c.Param("cakeID"))
	cake, err := h.useCase.GetCakeDetail(c.Request.Context(), cakeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Message{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Message{
		Message: "cake",
		Data:    mapCakeToCakeResponse(cake),
	})
}

func (h *CakeHandler) CreateNewCake(c *gin.Context) {
	request := core.CreateCakeRequest{}
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.Message{Error: err.Error()})
		return
	}
	cake, err := h.useCase.CreateNewCake(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Message{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response.Message{
		Message: "cake created",
		Data:    mapCakeToCakeResponse(cake),
	})
}

func (h *CakeHandler) UpdateCake(c *gin.Context) {
	request := core.CakeUpdateRequest{}
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.Message{Error: err.Error()})
		return
	}
	cakeID := StrToUint(c.Param("cakeID"))
	cake, err := h.useCase.UpdateCake(c.Request.Context(), cakeID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Message{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response.Message{
		Message: "cake updated",
		Data:    mapCakeToCakeResponse(cake),
	})
}

func (h *CakeHandler) DeleteCakeByID(c *gin.Context) {
	cakeID := StrToUint(c.Param("cakeID"))
	if err := h.useCase.DeleteCakeByID(c.Request.Context(), cakeID); err != nil {
		c.JSON(http.StatusInternalServerError, response.Message{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Message{Message: "cake deleted"})
}

func StrToUint(str string) uint {
	u, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return uint(u)
}
