package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ambu/app/ent"
	"github.com/ambu/app/ent/car"
	"github.com/ambu/app/ent/staff"
	"github.com/ambu/app/ent/statuscar"
	"github.com/ambu/app/ent/user"
	"github.com/gin-gonic/gin"
)

// PredicamentController defines the struct for the predicament controller
type PredicamentController struct {
	client *ent.Client
	router gin.IRouter
}

// Predicament defines the struct for the predicament
type Predicament struct {
	CarID       int
	StaffID     int
	StatuscarID int
	UserID      int
	AddedTime   string
}

// CreatePredicament handles POST requests for adding predicament entities
// @Summary Create predicament
// @Description Create predicament
// @ID create-predicament
// @Accept   json
// @Produce  json
// @Param predicament body Predicament true "Predicament entity"
// @Success 200 {object} ent.Predicament
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /predicaments [post]
func (ctl *PredicamentController) CreatePredicament(c *gin.Context) {
	obj := Predicament{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "predicament binding failed",
		})
		return
	}

	cn, err := ctl.client.Car.
		Query().
		Where(car.IDEQ(int(obj.CarID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "car not found",
		})
		return
	}

	st, err := ctl.client.Staff.
		Query().
		Where(staff.IDEQ(int(obj.StaffID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "staff diagnostic  not found",
		})
		return
	}

	s, err := ctl.client.Statuscar.
		Query().
		Where(statuscar.IDEQ(int(obj.StatuscarID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "statuscar diagnostic  not found",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.UserID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}
	times, err := time.Parse(time.RFC3339, obj.AddedTime)

	sa, err := ctl.client.Predicament.
		Create().
		SetCar(cn).
		SetStaff(st).
		SetStatuscar(s).
		SetUser(u).
		SetAddedTime(times).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, sa)
}

// ListPredicament handles request to get a list of predicament entities
// @Summary List predicament entities
// @Description list predicament entities
// @ID list-predicament
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Predicament
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /predicaments [get]
func (ctl *PredicamentController) ListPredicament(c *gin.Context) {
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

	predicaments, err := ctl.client.Predicament.
		Query().
		WithCar().
		WithStaff().
		WithStatuscar().
		WithUser().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, predicaments)
}

// DeletePredicament handles DELETE requests to delete a predicament entity
// @Summary Delete a predicament entity by ID
// @Description get predicament by ID
// @ID delete-predicament
// @Produce  json
// @Param id path int true "Predicament ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /predicaments/{id} [delete]
func (ctl *PredicamentController) DeletePredicament(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Predicament.
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

// NewPredicamentController creates and registers handles for the predicament controller
func NewPredicamentController(router gin.IRouter, client *ent.Client) *PredicamentController {
	drc := &PredicamentController{
		client: client,
		router: router,
	}
	drc.register()
	return drc
}

// InitPredicamentController registers routes to the main engine
func (ctl *PredicamentController) register() {
	predicaments := ctl.router.Group("/predicaments")

	predicaments.GET("", ctl.ListPredicament)

	// CRUD
	predicaments.POST("", ctl.CreatePredicament)
	predicaments.DELETE(":id", ctl.DeletePredicament)
}
