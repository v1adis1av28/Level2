package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/v1adis1av28/level2/tasks/task18/app/internal/models"
	"github.com/v1adis1av28/level2/tasks/task18/app/internal/validate"
)

type Handler struct {
	Data map[int][]models.Event
}
type CreateRequest struct {
	UserId  int    `json:"user_id"`
	EventId int    `json:"event_id"`
	Date    string `json:"date"`
	Name    string `json:"name"`
}

type EditRequest struct { //структура для запросов update/delete
	UserId  int `json:"user_id"`
	EventId int `json:"event_id"`
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
	var eventReq CreateRequest
	err := c.ShouldBindJSON(&eventReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if eventReq.UserId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id can't be negative or zero"})
		return
	}

	if eventReq.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "event name cannot be empty"})
		return
	}

	date, err := time.Parse("2006-01-02", eventReq.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, use YYYY-MM-DD"})
		return
	}

	validateErr := validate.ValidateDate(date)
	if validateErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		return
	}

	event := models.Event{
		UserId:  eventReq.UserId,
		EventId: eventReq.EventId,
		Date:    date,
		Name:    eventReq.Name,
	}

	h.Data[event.UserId] = append(h.Data[event.UserId], event)
	fmt.Println(len(h.Data[event.UserId]))
	c.JSON(http.StatusOK, gin.H{"result": event})
}

func (h *Handler) UpdateEvent(c *gin.Context) {

}

func (h *Handler) DeleteEvent(c *gin.Context) {
	var event EditRequest
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	index := -1
	for i, val := range h.Data[event.UserId] {
		if val.EventId == event.EventId {
			index = i
			break
		}
	}
	if index == -1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "event not found"})
		return
	}
	h.Data[event.UserId] = append(h.Data[event.UserId][:index], h.Data[event.UserId][index+1:]...)
	fmt.Println(len(h.Data[event.UserId]))
	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("event with id %d was succesfully deleted", event.EventId)})
}
