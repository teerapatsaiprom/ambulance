package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ambu/app/ent"
	"github.com/ambu/app/ent/car"
	"github.com/gin-gonic/gin"
)

// CarController defines the struct for the car controller
type CarController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateCar handles POST requests for adding car entities
// @Summary Create car
// @Description Create car
// @ID create-car
// @Accept   json
// @Produce  json
// @Param car body ent.Car true "Car entity"
// @Success 200 {object} ent.Car
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cars [post]
func (ctl *CarController) CreateCar(c *gin.Context) {
	obj := ent.Car{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "car binding failed",
		})
		return
	}

	cn, err := ctl.client.Car.
		Create().
		SetCarNo(obj.CarNo).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, cn)
}

// GetCar handles GET requests to retrieve a car entity
// @Summary Get a car entity by ID
// @Description get car by ID
// @ID get-car
// @Produce  json
// @Param id path int true "Car ID"
// @Success 200 {object} ent.Car
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cars/{id} [get]
func (ctl *CarController) GetCar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cn, err := ctl.client.Car.
		Query().
		Where(car.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cn)
}

// ListCar handles request to get a list of car entities
// @Summary List car entities
// @Description list car entities
// @ID list-car
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Car
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cars [get]
func (ctl *CarController) ListCar(c *gin.Context) {
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

	cars, err := ctl.client.Car.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cars)
}

// DeleteCar handles DELETE requests to delete a car entity
// @Summary Delete a car entity by ID
// @Description get car by ID
// @ID delete-car
// @Produce  json
// @Param id path int true "Car ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cars/{id} [delete]
func (ctl *CarController) DeleteCar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Car.
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

// UpdateCar handles PUT requests to update a car entity
// @Summary Update a car entity by ID
// @Description update car by ID
// @ID update-car
// @Accept   json
// @Produce  json
// @Param id path int true "Car ID"
// @Param car body ent.Car true "Car entity"
// @Success 200 {object} ent.Car
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cars/{id} [put]
func (ctl *CarController) UpdateCar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Car{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "car binding failed",
		})
		return
	}
	obj.ID = int(id)
	cn, err := ctl.client.Car.
		UpdateOneID(int(id)).
		SetCarNo(obj.CarNo).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, cn)
}

// NewCarController creates and registers handles for the CarController
func NewCarController(router gin.IRouter, client *ent.Client) *CarController {
	cn := &CarController{
		client: client,
		router: router,
	}
	cn.register()
	return cn
}

// InitCarController registers routes to the main engine
func (ctl *CarController) register() {
	cars := ctl.router.Group("/cars")
	cars.GET("", ctl.ListCar)
	// CRUD
	cars.POST("", ctl.CreateCar)
	cars.GET(":id", ctl.GetCar)
	cars.PUT(":id", ctl.UpdateCar)
	cars.DELETE(":id", ctl.DeleteCar)
}
