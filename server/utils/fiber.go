package utils

import (
	"math"

	"github.com/SAF2k/restaurant-management/database"
	"github.com/SAF2k/restaurant-management/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
)

// // Create a new User type
// var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var userCollection *gorm.DB = database.OpenTable(database.Client, "user")

// ParseBody is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBody(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

// ParseBodyAndValidate is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ParseBody(ctx, body); err != nil {
		return err
	}

	return Validate(body)
}

// GetUser is helper function for getting authenticated user's id
func GetUser(c *fiber.Ctx) *string {
	// Get user from context
	user := c.Locals("user").(*models.UserResponse)

	// Create user id
	id := user.ID.Hex()

	// id := "1"

	// Find user by id
	// err := userCollection.FindOne(c.Context(), bson.M{"_id": id}).Decode(&user)
	err := userCollection.First(&user, bson.M{"_id": id}).Error
	if err != nil {
		return nil
	}

	// Return user id
	return &id
}

// ToFixed is helper function for rounding the float64 number
func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

// Helps to round the float64 number used in ToFixed function
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
