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

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoices")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		cursor, err := invoiceCollection.Find(ctx, primitive.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while fetching invoices",
			})
			return
		}

		var invoices []bson.M = make([]bson.M, 0)
		if err = cursor.All(ctx, &invoices); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while decoding invoices",
			})
			return
		}

		c.JSON(http.StatusOK, invoices)

	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var invoice models.Invoice
		invoice_id, _ := primitive.ObjectIDFromHex(c.Param("invoice_id"))
		if err := invoiceCollection.FindOne(ctx, bson.M{"_id": invoice_id}).Decode(&invoice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while fetching invoice",
			})
			return
		}

		c.JSON(http.StatusOK, invoice)
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var invoice models.Invoice
		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Check if order_id is present
		if invoice.Order_id == nil || *invoice.Order_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "order_id is required",
			})
			return
		}

		// Validate the payment_method field
		if invoice.Payment_method != nil {
			validMethods := []string{"CARD", "CASH"} // Define valid payment methods
			isValid := false
			for _, method := range validMethods {
				if *invoice.Payment_method == method {
					isValid = true
					break
				}
			}
			if !isValid {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid payment_method value",
				})
				return
			}
		}

		// Validate the payment_status field
		if invoice.Payment_status != nil {
			validStatus := []string{"PAID", "UNPAID"} // Define valid payment statuses
			isValid := false
			for _, status := range validStatus {
				if *invoice.Payment_status == status {
					isValid = true
					break
				}
			}
			if !isValid {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid payment_status value",
				})
				return
			}
		}

		// Calculate the payment due date as 1 hour from the current time
		invoice.Payment_due_date = time.Now().Add(time.Hour)

		// Set timestamps
		now := time.Now()
		invoice.Create_at = now
		invoice.Update_at = now
		invoice.ID = primitive.NewObjectID()
		invoice.Invoice_id = invoice.ID.Hex()

		result, insertErr := invoiceCollection.InsertOne(ctx, invoice)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while creating a new invoice",
			})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update a invoice",
		})
	}
}

func DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Delete a invoice",
		})
	}
}
