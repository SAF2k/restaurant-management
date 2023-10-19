package controllers

import (
	"restaurant-management/server-2/database"
	"restaurant-management/server-2/models"
	"restaurant-management/server-2/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var storeCollection *mongo.Collection = database.OpenCollection(database.Client, "store")

func GetAllStores(ctx *fiber.Ctx) error {
	//Find all stores
	result, err := storeCollection.Find(ctx.Context(), bson.M{})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not connect to database")
	}

	//Create bson.M type var
	var stores []bson.M

	//Decode all stores
	if err := result.All(ctx.Context(), &stores); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find any store")
	}

	//Return stores
	return ctx.JSON(stores)
}

func GetStore(ctx *fiber.Ctx) error {
	//Get id from params
	storeId := ctx.Params("id")

	//Create Store struct
	var store models.Store

	//Find store and decode
	err := storeCollection.FindOne(ctx.Context(), bson.M{"store_id": storeId}).Decode(&store)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find store")
	}

	//Return store
	return ctx.JSON(store)
}

func CreateStore(ctx *fiber.Ctx) error {
	//Create new Store struct
	store := new(models.Store)

	//Parse body and validate
	if err := utils.ParseBodyAndValidate(ctx, store); err != nil {
		return err
	}

	//Create bson.M type var
	store.ID = primitive.NewObjectID()
	store.Store_id = store.ID.Hex()

	//Create created_at and updated_at
	store.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	store.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	//Insert store into database
	result, err := storeCollection.InsertOne(ctx.Context(), store)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create menu")
	}

	//Return result
	return ctx.JSON(result)
}

func UpdateStore(ctx *fiber.Ctx) error {
	//Get id from params
	storeId := ctx.Params("id")

	//Create new Store struct
	store := new(models.Store)

	//Parse body and validate
	if err := utils.ParseBodyAndValidate(ctx, store); err != nil {
		return err
	}

	//Create update var
	update := bson.M{
		"$set": bson.M{
			"name": store.Name,
		},
	}

	//Update store in database
	result, err := storeCollection.UpdateOne(ctx.Context(), bson.M{"store_id": storeId}, update)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update store")
	}

	//Return result
	return ctx.JSON(result)
}
