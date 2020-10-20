package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ambu/app/ent"
	"github.com/ambu/app/ent/statuscar"
)

// StatuscarController defines the struct for the statuscar controller
type StatuscarController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateStatuscar handles POST requests for adding statuscar entities
// @Summary Create statuscar
// @Description Create statuscar
// @ID create-statuscar
// @Accept   json
// @Produce  json
// @Param statuscar body ent.Statuscar true "Statuscar entity"
// @Success 200 {object} ent.Statuscar
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuscars [post]
func (ctl *StatuscarController) CreateStatuscar(c *gin.Context) {
	obj := ent.Statuscar{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "statuscar binding failed",
		})
		return
	}

	s, err := ctl.client.Statuscar.
		Create().
		SetStatusDetail(obj.StatusDetail).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, s)
}

// GetStatuscar handles GET requests to retrieve a statuscar entity
// @Summary Get a statuscar entity by ID
// @Description get statuscar by ID
// @ID get-statuscar
// @Produce  json
// @Param id path int true "Statuscar ID"
// @Success 200 {object} ent.Statuscar
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuscars/{id} [get]
func (ctl *StatuscarController) GetStatuscar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.Statuscar.
		Query().
		Where(statuscar.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListStatuscar handles request to get a list of statuscar entities
// @Summary List statuscar entities
// @Description list statuscar entities
// @ID list-statuscar
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Statuscar
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuscars [get]
func (ctl *StatuscarController) ListStatuscar(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	statuscars, err := ctl.client.Statuscar.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, statuscars)
}

// DeleteStatuscar handles DELETE requests to delete a statuscar entity
// @Summary Delete a statuscar entity by ID
// @Description get statuscar by ID
// @ID delete-statuscar
// @Produce  json
// @Param id path int true "Statuscar ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuscars/{id} [delete]
func (ctl *StatuscarController) DeleteStatuscar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Statuscar.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateStatuscar handles PUT requests to update a statuscar entity
// @Summary Update a statuscar entity by ID
// @Description update statuscar by ID
// @ID update-statuscar
// @Accept   json
// @Produce  json
// @Param id path int true "Statuscar ID"
// @Param statuscar body ent.Statuscar true "Statuscar entity"
// @Success 200 {object} ent.Statuscar
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuscars/{id} [put]
func (ctl *StatuscarController) UpdateStatuscar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Statuscar{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "statuscar binding failed",
		})
		return
	}
	obj.ID = int(id)
	s, err := ctl.client.Statuscar.
		UpdateOneID(int(id)).
		SetStatusDetail(obj.StatusDetail).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, s)
}

// NewStatuscarController creates and registers handles for the StatuscarController
func NewStatuscarController(router gin.IRouter, client *ent.Client) *StatuscarController {
	sc := &StatuscarController{
		client: client,
		router: router,
	}
	sc.register()
	return sc
}

// InitStatuscarController registers routes to the main engine
func (ctl *StatuscarController) register() {
	statuscars := ctl.router.Group("/statuscars")
	statuscars.GET("", ctl.ListStatuscar)
	// CRUD
	statuscars.POST("", ctl.CreateStatuscar)
	statuscars.GET(":id", ctl.GetStatuscar)
	statuscars.PUT(":id", ctl.UpdateStatuscar)
	statuscars.DELETE(":id", ctl.DeleteStatuscar)
}
