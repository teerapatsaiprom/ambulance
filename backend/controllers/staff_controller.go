package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ambu/app/ent"
	"github.com/ambu/app/ent/staff"
)

// StaffController defines the struct for the staff controller
type StaffController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateStaff handles POST requests for adding staff entities
// @Summary Create staff
// @Description Create staff
// @ID create-staff
// @Accept   json
// @Produce  json
// @Param Staff body ent.Staff true "Staff entity"
// @Success 200 {object} ent.Staff
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /staffs [post]
func (ctl *StaffController) CreateStaff(c *gin.Context) {
	obj := ent.Staff{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "staff binding failed",
		})
		return
	}

	st, err := ctl.client.Staff.
		Create().
		SetStaffEmail(obj.StaffEmail).
		SetStaffName(obj.StaffName).
		SetStaffPassword(obj.StaffPassword).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, st)
}

// GetStaff handles GET requests to retrieve a staff entity
// @Summary Get a staff entity by ID
// @Description get staff by ID
// @ID get-staff
// @Produce  json
// @Param id path int true "Staff ID"
// @Success 200 {object} ent.Staff
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /staffs/{id} [get]
func (ctl *StaffController) GetStaff(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	st, err := ctl.client.Staff.
		Query().
		Where(staff.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, st)
}

// ListStaff handles request to get a list of staff entities
// @Summary List staff entities
// @Description list staff entities
// @ID list-staff
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Staff
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /staffs [get]
func (ctl *StaffController) ListStaff(c *gin.Context) {
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

	staffs, err := ctl.client.Staff.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, staffs)
}

// DeleteStaff handles DELETE requests to delete a staff entity
// @Summary Delete a staff entity by ID
// @Description get staff by ID
// @ID delete-staff
// @Produce  json
// @Param id path int true "Staff ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Staffs/{id} [delete]
func (ctl *StaffController) DeleteStaff(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Staff.
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

// UpdateStaff handles PUT requests to update a staff entity
// @Summary Update a staff entity by ID
// @Description update staff by ID
// @ID update-staff
// @Accept   json
// @Produce  json
// @Param id path int true "Staff ID"
// @Param staff body ent.Staff true "Staff entity"
// @Success 200 {object} ent.Staff
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /staffs/{id} [put]
func (ctl *StaffController) UpdateStaff(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Staff{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "staff binding failed",
		})
		return
	}
	obj.ID = int(id)
	s, err := ctl.client.Staff.
		UpdateOneID(int(id)).
		SetStaffEmail(obj.StaffEmail).
		SetStaffName(obj.StaffName).
		SetStaffPassword(obj.StaffPassword).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, s)
}

// NewStaffController creates and registers handles for the StaffController
func NewStaffController(router gin.IRouter, client *ent.Client) *StaffController {
	sc := &StaffController{
		client: client,
		router: router,
	}
	sc.register()
	return sc
}

// InitStaffController registers routes to the main engine
func (ctl *StaffController) register() {
	staffs := ctl.router.Group("/staffs")
	staffs.GET("", ctl.ListStaff)
	// CRUD
	staffs.POST("", ctl.CreateStaff)
	staffs.GET(":id", ctl.GetStaff)
	staffs.PUT(":id", ctl.UpdateStaff)
	staffs.DELETE(":id", ctl.DeleteStaff)
}
