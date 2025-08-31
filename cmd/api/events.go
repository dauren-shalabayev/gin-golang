package main

import (
	"net/http"
	"rest-api-in-gin/cmd/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateEvent godoc
// @Summary Create a new event
// @Description Create a new event with the provided information
// @Tags events
// @Accept json
// @Produce json
// @Param event body map[string]interface{} true "Event object"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /events [post]
func (app *application) CreateEvent(c *gin.Context) {
	var event database.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := app.models.Events.Insert(event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"event": event})
}

// GetEvents godoc
// @Summary Get all events
// @Description Retrieve a list of all events
// @Tags events
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /events [get]
func (app *application) GetEvents(c *gin.Context) {
	events, err := app.models.Events.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"events": events})
}

// GetEvent godoc
// @Summary Get event by ID
// @Description Retrieve a specific event by its ID
// @Tags events
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /events/{id} [get]
func (app *application) GetEvent(c *gin.Context) {
	id := c.Param("id")
	event, err := app.models.Events.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"event": event})
}

// UpdateEvent godoc
// @Summary Update event by ID
// @Description Update an existing event's information
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Param event body map[string]interface{} true "Updated event object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /events/{id} [put]
func (app *application) UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var event database.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := app.models.Events.Update(id, event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

// DeleteEvent godoc
// @Summary Delete event by ID
// @Description Delete an event by its ID
// @Tags events
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /events/{id} [delete]
func (app *application) DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	if err := app.models.Events.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

// AddAttendeeToEvent godoc
// @Summary Add attendee to event
// @Description Add a user as an attendee to a specific event
// @Tags events
// @Produce json
// @Param id path string true "Event ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /events/{id}/attendees/{user_id} [post]
func (app *application) AddAttendeeToEvent(c *gin.Context) {
	id := c.Param("id")
	userID := c.Param("user_id")

	// Convert string IDs to integers
	eventID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	attendee := database.Attendee{
		UserID:  userIDInt,
		EventID: eventID,
	}

	if err := app.models.Attendees.Insert(attendee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Attendee added to event successfully"})
}

// GetAttendeesForEvent godoc
// @Summary Get attendees for event
// @Description Retrieve all attendees for a specific event
// @Tags events
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /events/{id}/attendees [get]
func (app *application) GetAttendeesForEvent(c *gin.Context) {
	id := c.Param("id")

	eventID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	attendees, err := app.models.Attendees.GetByEventID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"attendees": attendees})
}
