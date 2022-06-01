package controller

import (
	"assignment-2/database"
	"assignment-2/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()

	var createOrder models.CreateOrder

	var items []models.Item

	err := ctx.ShouldBindJSON(&createOrder)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order := models.Order{
		CustomerName: createOrder.CustomerName,
		OrderedAt:    createOrder.OrderedAt,
	}

	db.Create(&order)
	idOrder := order.ID

	for _, v := range createOrder.Items {
		item := models.Item{
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Quantity,
			OrderID:     order.ID,
		}

		items = append(items, item)
	}

	fmt.Println(items)

	result := db.Create(&items)
	log.Println(idOrder, result.RowsAffected)

	var data interface{}

	data = models.Order{
		ID:           items[0].OrderID,
		CustomerName: createOrder.CustomerName,
		Items:        items,
		OrderedAt:    time.Now(),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Succes Creating A New Order",
		"Payload": data,
	})
}

func GetAllOrder(ctx *gin.Context) {
	db := database.GetDB()

	orders := []models.Order{}

	err := db.Preload("Items").Find(&orders).Error

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Succes Get All Orders",
		"Payload": orders,
	})
}

func GetOrderByID(ctx *gin.Context) {
	db := database.GetDB()

	order := models.Order{}

	var id_order = ctx.Param("orderId")
	finalId, _ := strconv.Atoi(id_order)
	condition := false

	orders := []models.Order{}

	err := db.Preload("Items").Find(&orders).Error

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	errOrder := db.Where("id = ?", id_order).Find(&order).Error
	if errOrder != nil {
		fmt.Println(err.Error())
		return
	}

	for _, o := range orders {
		if o.ID == finalId {
			condition = true
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": http.StatusNotFound,
			"msg":   "NOT FOUND",
		})
		return
	}

	items := []models.Item{}

	err = db.Where("order_id = ?", id_order).Find(&items).Error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	response := models.CreateOrder{}
	response.CustomerName = order.CustomerName
	response.OrderedAt = order.OrderedAt
	response.Items = items

	message := fmt.Sprintf("Succes Get Order With ID %s", id_order)

	ctx.JSON(http.StatusOK, gin.H{
		"Message": message,
		"Payload": response,
	})
}

func UpdateOrder(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("orderId")
	finalId, _ := strconv.Atoi(id)

	var updateData = models.CreateOrder{}

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order := models.Order{
		ID:           finalId,
		CustomerName: updateData.CustomerName,
		OrderedAt:    updateData.OrderedAt,
		Items:        updateData.Items,
	}

	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order)

	message := fmt.Sprintf("Succes Update Order With ID %s", id)

	ctx.JSON(http.StatusOK, gin.H{
		"Message": message,
		"Payload": order,
	})
}

func DeleteOrder(c *gin.Context) {
	db := database.GetDB()

	id := c.Param("orderId")
	finalId, _ := strconv.Atoi(id)

	orders := models.Order{}
	items := models.Item{}

	err := db.Where("order_id = ?", finalId).Delete(&items).Error

	rows := db.Where("id = ?", finalId).Delete(&orders).RowsAffected

	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Data Not Found",
		})
		return
	}
	if err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Data Success For Delete",
	})
}
