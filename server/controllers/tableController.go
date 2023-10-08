package controller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saf2k/restaurant-management/server/database"
	"github.com/saf2k/restaurant-management/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Define pagination parameters based on query parameters
		page := c.DefaultQuery("page", "1")        // Default to page 1 if not specified
		perPage := c.DefaultQuery("perPage", "10") // Default to 10 items per page if not specified

		pageNum, err := strconv.Atoi(page)
		if err != nil || pageNum < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid page number",
			})
			return
		}

		perPageNum, err := strconv.Atoi(perPage)
		if err != nil || perPageNum < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid items per page",
			})
			return
		}

		// Calculate skip value for pagination
		skip := (pageNum - 1) * perPageNum

		// Define the filter (if needed)
		filter := bson.M{} // You can add filters here if necessary

		// Apply pagination and filter
		options := options.Find()
		options.SetSkip(int64(skip))
		options.SetLimit(int64(perPageNum))

		result, err := tableCollection.Find(ctx, filter, options)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error fetching tables",
			})
			return
		}

		var tables []bson.M

		if err = result.All(ctx, &tables); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error processing tables",
			})
			return
		}

		c.JSON(http.StatusOK, tables)
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		tableId := c.Param("table_id")
		var table models.Table

		// Create a context with a timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel() // Cancel the context when done.

		err := tableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&table)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Table not found",
			})
			return
		}
		c.JSON(http.StatusOK, table)
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table

		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the table
		if validationError := validate.Struct(table); validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Error()})
			return
		}

		// Set Created_at and Updated_at fields directly
		table.Created_at = time.Now()
		table.Updated_at = time.Now()

		// Generate a new ObjectID and set it as Table_id
		table.ID = primitive.NewObjectID()
		table.Table_id = table.ID.Hex()

		// Insert the table into the database
		if _, err := tableCollection.InsertOne(c, table); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating table"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Table created successfully"})
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get table ID from URL parameter
		tableID := c.Param("table_id")

		// Parse JSON request into a table struct
		var table models.Table
		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Define the filter to find the table by ID
		filter := bson.M{"table_id": tableID}

		// Create an update document with the $set operator
		update := bson.M{
			"$set": bson.M{
				"number_of_guests": table.Number_of_guests,
				"table_number":     table.Table_number,
				"updated_at":       time.Now(),
			},
		}

		// Perform the update operation
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second) // Adjust the timeout as needed
		defer cancel()

		result, err := tableCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}

		if result.ModifiedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Table not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Table updated successfully",
		})
	}
}

func DeleteTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		tableID := c.Param("table_id")
		filter := bson.M{"table_id": tableID}

		result, err := tableCollection.DeleteOne(c, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting table"})
			return
		}

		if result.DeletedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Table not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Table deleted successfully"})
	}
}

func GetTableByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		user_id := c.Param("user_id")
		defer cancel()

		filter := bson.M{"user_id": user_id}

		result, err := tableCollection.Find(ctx, filter)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Table not found",
			})
			return
		}

		var tables []bson.M

		if err = result.All(ctx, &tables); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error processing tables",
			})
			return
		}

		c.JSON(http.StatusOK, tables)
	}
}
