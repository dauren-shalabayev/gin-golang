package main

import (
	"net/http"
	"rest-api-in-gin/cmd/internal/database"

	"github.com/gin-gonic/gin"
)

// CreateAttendee godoc
// @Summary Create a new attendee
// @Description Create a new attendee record
// @Tags attendees
// @Accept json
// @Produce json
// @Param attendee body map[string]interface{} true "Attendee object"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /attendees [post]
func (app *application) CreateAttendee(c *gin.Context) {
	var attendee database.Attendee

	if err := c.ShouldBindJSON(&attendee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := app.models.Attendees.Insert(attendee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"attendee": attendee})
}

// GetAttendees godoc
// @Summary Get all attendees
// @Description Retrieve a list of all attendees
// @Tags attendees
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /attendees [get]
func (app *application) GetAttendees(c *gin.Context) {
	attendees, err := app.models.Attendees.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"attendees": attendees})
}

// GetAttendee godoc
// @Summary Get attendee by ID
// @Description Retrieve a specific attendee by their ID
// @Tags attendees
// @Produce json
// @Param id path string true "Attendee ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /attendees/{id} [get]
func (app *application) GetAttendee(c *gin.Context) {
	id := c.Param("id")
	attendee, err := app.models.Attendees.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"attendee": attendee})
}

// UpdateAttendee godoc
// @Summary Update attendee by ID
// @Description Update an existing attendee's information
// @Tags attendees
// @Accept json
// @Produce json
// @Param id path string true "Attendee ID"
// @Param attendee body map[string]interface{} true "Updated attendee object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /attendees/{id} [put]
func (app *application) UpdateAttendee(c *gin.Context) {
	id := c.Param("id")
	var attendee database.Attendee
	if err := c.ShouldBindJSON(&attendee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := app.models.Attendees.Update(id, attendee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendee updated successfully"})
}

// DeleteAttendee godoc
// @Summary Delete attendee by ID
// @Description Delete an attendee by their ID
// @Tags attendees
// @Produce json
// @Param id path string true "Attendee ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /attendees/{id} [delete]
func (app *application) DeleteAttendee(c *gin.Context) {
	id := c.Param("id")
	if err := app.models.Attendees.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Attendee deleted successfully"})
}
