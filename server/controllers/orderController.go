package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saf2k/restaurant-management/server/database"
	"github.com/saf2k/restaurant-management/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		result, err := orderCollection.Find(ctx, bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Orders not found",
			})
			return
		}
		var allOrders []bson.M

		if err = result.All(ctx, &allOrders); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Foods not found",
			})
			return
		}

		c.JSON(http.StatusOK, allOrders)
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderID := c.Param("order_id")

		filter := bson.M{"order_id": orderID}

		var order models.Order

		// Find and decode the order
		err := orderCollection.FindOne(ctx, filter).Decode(&order)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Order not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}

		// Set the order in the response
		c.JSON(http.StatusOK, order)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var order models.Order
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Check if Table_id is present
		if order.Table_id == nil || *order.Table_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Table_id is required",
			})
			return
		}

		// Check if the table with the specified Table_id exists
		tableId := *order.Table_id
		var existingTable models.Table
		if err := tableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&existingTable); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Table not found",
			})
			return
		}

		// Validate the order_status field
		if order.Order_status != nil {
			validStatus := []string{"OPEN", "CLOSE"} // Define valid order_status values
			isValid := false
			for _, status := range validStatus {
				if *order.Order_status == status {
					isValid = true
					break
				}
			}
			if !isValid {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid order_status value",
				})
				return
			}
		}

		// Set timestamps
		now := time.Now()
		order.Order_date = now
		order.Created_at = now
		order.Updated_at = now
		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		result, insertErr := orderCollection.InsertOne(ctx, order)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while creating a new order",
			})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Get order ID from URL parameter
		orderID := c.Param("order_id")

		// Parse JSON request into an order struct
		var updatedOrder models.Order
		if err := c.BindJSON(&updatedOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Check if the order with the specified Order_id exists
		var existingOrder models.Order
		if err := orderCollection.FindOne(ctx, bson.M{"order_id": orderID}).Decode(&existingOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Order not found",
			})
			return
		}

		// Validate the order_status field
		if updatedOrder.Order_status != nil {
			validStatus := []string{"OPEN", "CLOSE"} // Define valid order_status values
			isValid := false
			for _, status := range validStatus {
				if *updatedOrder.Order_status == status {
					isValid = true
					break
				}
			}
			if !isValid {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid order_status value",
				})
				return
			}
		}

		// Update the existing order fields
		existingOrder.Order_status = updatedOrder.Order_status
		existingOrder.Updated_at = time.Now()

		// Perform the update operation
		update := bson.M{
			"$set": bson.M{
				"order_status": existingOrder.Order_status,
				"updated_at":   existingOrder.Updated_at,
			},
		}
		filter := bson.M{"order_id": orderID}
		_, updateErr := orderCollection.UpdateOne(ctx, filter, update)
		if updateErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error updating order",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Order updated successfully",
		})
	}
}

func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		orderID := c.Param("order_id")

		// Check if the order with the specified Order_id exists

		var existingOrder models.Order
		if err := orderCollection.FindOne(ctx, bson.M{"order_id": orderID}).Decode(&existingOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Order not found",
			})
			return
		}

		// Delete the order
		filter := bson.M{"order_id": orderID}
		_, err := orderCollection.DeleteOne(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error deleting order",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Order deleted successfully",
		})
	}
}
