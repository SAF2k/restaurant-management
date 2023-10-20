package controllers

import (
	"restaurant-management/server-2/database"
	"restaurant-management/server-2/models"
	"restaurant-management/server-2/utils"
	"restaurant-management/server-2/utils/jwt"
	"restaurant-management/server-2/utils/password"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// Login service logs in a user
func Login(ctx *fiber.Ctx) error {
	// Create a new LoginDTO type
	b := new(models.LoginDTO)

	// Parse body and validate

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	// create a user response
	user := &models.UserResponse{}

	// check if email exists
	err := userCollection.FindOne(ctx.Context(), bson.M{"email": b.Email}).Decode(user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not connect to database")
	}

	// compare passwords
	if err := password.Verify(user.Password, b.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	// generate access token
	t := jwt.Generate(&jwt.TokenPayload{
		ID: user.User_id,
	})

	// return user and access token
	return ctx.JSON(&models.AuthResponse{
		User: user,
		Auth: &models.AccessResponse{
			Token: t,
		},
	})
}

// Signup service creates a user
func Signup(ctx *fiber.Ctx) error {
	// Create a new SignupDTO type
	b := new(models.SignupDTO)

	// Parse body and validate
	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	// Check if email or store_id exists
	if err := userCollection.FindOne(ctx.Context(), bson.M{"email": b.Email}).Err(); err == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Email already exists")
	}

	// Check if email or store_id exists
	if err := userCollection.FindOne(ctx.Context(), bson.M{"store_id": b.Store_id}).Err(); err == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Store ID already exists")
	}

	// Check if email or store_id exists
	if err := storeCollection.FindOne(ctx.Context(), bson.M{"store_id": b.Store_id}).Err(); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Store ID Not found")
	}

	id_ := primitive.NewObjectID()
	user_id_ := id_.Hex()

	// Create a user model
	user := &models.UserResponse{
		ID:         id_,
		User_id:    user_id_,
		Name:       b.Name,
		Email:      b.Email,
		Store_id:   b.Store_id,
		Password:   password.Generate(b.Password),
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	// Create a user, if error return
	_, err := userCollection.InsertOne(ctx.Context(), user)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create user")
	}

	// generate access token
	t := jwt.Generate(&jwt.TokenPayload{
		ID: user.User_id,
	})

	// return user and access token
	return ctx.JSON(&models.AuthResponse{
		User: &models.UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Store_id: user.Store_id,
		},
		Auth: &models.AccessResponse{
			Token: t,
		},
	})
}
