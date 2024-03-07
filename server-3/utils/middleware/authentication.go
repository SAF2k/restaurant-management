package middleware

import (
	"context"
	"fmt"
	"restaurant-management/server-2/database"
	"restaurant-management/server-2/models"
	"restaurant-management/server-2/utils/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// Auth is the authentication middleware
func Auth(c *fiber.Ctx) error {
	h := c.Get("Authorization")
	if h == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
	}

	// Split the header into chunks
	chunks := strings.Split(h, " ")

	// If header signature is not like `Bearer <token>`, then throw
	if len(chunks) < 2 {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid Authorization header")
	}

	// Verify the token which is in the chunks
	user, err := jwt.Verify(chunks[1])

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	// Convert the user ID to a MongoDB ObjectID
	userIdFilter := bson.M{"user_id": user.ID}
	userDocument := new(models.UserResponse)

	// Fetch the user document
	err = userCollection.FindOne(context.Background(), userIdFilter).Decode(&userDocument)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}
		fmt.Println("Error fetching user document:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong")
	}
	fmt.Println("User Doc", userDocument)

	// Check if the store ID matches what's expected
	if userDocument.Store_id != c.Params("s_id") {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	// You can store the user information in the context for further use
	c.Locals("USER", userDocument)

	return c.Next()
}

func GetUserAuth(c *fiber.Ctx) error {
	h := c.Get("Authorization")
	if h == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
	}

	// Split the header into chunks
	chunks := strings.Split(h, " ")

	// If header signature is not like `Bearer <token>`, then throw
	if len(chunks) < 2 {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid Authorization header")
	}

	// Verify the token which is in the chunks
	user, err := jwt.Verify(chunks[1])

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	// Convert the user ID to a MongoDB ObjectID
	userIdFilter := bson.M{"user_id": user.ID}
	userDocument := new(models.UserResponse)

	// Fetch the user document
	err = userCollection.FindOne(context.Background(), userIdFilter).Decode(&userDocument)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}
		fmt.Println("Error fetching user document:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong")
	}
	fmt.Println("User Doc", userDocument)
	// You can store the user information in the context for further use
	c.Locals("USER", userDocument)

	return c.Next()
}
