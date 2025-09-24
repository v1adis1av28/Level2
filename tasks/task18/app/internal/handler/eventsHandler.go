package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/v1adis1av28/level2/tasks/task18/app/internal/models"
	"github.com/v1adis1av28/level2/tasks/task18/app/internal/validate"
)

type Handler struct {
	Data map[int][]models.Event
}

func NewHandler(mp map[int][]models.Event) *Handler {
	return &Handler{
		Data: mp,
	}
}

func (h *Handler) GetDayEvents(c *gin.Context) {

}

func (h *Handler) GetWeekEvents(c *gin.Context) {

}

func (h *Handler) GetMonthEvents(c *gin.Context) {

}

func (h *Handler) CreateEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if event.UserId < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id can`t be negative"})
		return
	}
	validateErr := validate.ValidateDate(event.Date)
	if errors.Is(validateErr, fmt.Errorf("you can`t create event in the past")) {
		c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		return
	}
	h.Data[event.UserId] = append(h.Data[event.UserId], event)
	fmt.Println(h.Data[event.UserId])
}

func (h *Handler) UpdateEvent(c *gin.Context) {

}

func (h *Handler) DeleteEvent(c *gin.Context) {

}
