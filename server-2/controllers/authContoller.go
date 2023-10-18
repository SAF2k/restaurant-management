package controllers

import (
	"errors"
	"fmt"
	"restaurant-management/server-2/database"
	"restaurant-management/server-2/models"
	"restaurant-management/server-2/utils"
	"restaurant-management/server-2/utils/jwt"
	"restaurant-management/server-2/utils/password"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// Login service logs in a user
func Login(ctx *fiber.Ctx) error {
	b := new(models.LoginDTO)
	fmt.Println(b)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid email or password")
	}

	u := &models.UserResponse{}

	err := userCollection.FindOne(ctx.Context(), bson.M{"email": b.Email}).Decode(u)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not connect to database")
	}

	fmt.Println(u)

	if err := password.Verify(u.Password, b.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	t := jwt.Generate(&jwt.TokenPayload{
		ID: u.ID,
	})

	return ctx.JSON(&models.AuthResponse{
		User: u,
		Auth: &models.AccessResponse{
			Token: t,
		},
	})
}

// Signup service creates a user
func Signup(ctx *fiber.Ctx) error {
	b := new(models.SignupDTO)

	fmt.Println(b)
	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	// Check if email already exists
	err := userCollection.FindOne(ctx.Context(), bson.M{"email": b.Email}).Decode(&models.UserResponse{})

	// If email already exists, return
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return fiber.NewError(fiber.StatusInternalServerError, "Email already exists")
	}

	user := &models.UserResponse{
		ID:       primitive.NewObjectID(),
		Name:     b.Name,
		Email:    b.Email,
		Password: password.Generate(b.Password),
	}

	fmt.Println(user)

	// Create a user, if error return
	_, err = userCollection.InsertOne(ctx.Context(), user)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create user")
	}

	// generate access token
	t := jwt.Generate(&jwt.TokenPayload{
		ID: user.ID,
	})

	return ctx.JSON(&models.AuthResponse{
		User: &models.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Auth: &models.AccessResponse{
			Token: t,
		},
	})
}
