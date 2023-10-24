package controllers

import (
	"fmt"
	"restaurant-management/server-2/database"
	"restaurant-management/server-2/models"
	"restaurant-management/server-2/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, " table")

func GetAllTables(ctx *fiber.Ctx) error {
	//
	storeId := ctx.Params("s_id")

	var tables []models.Table

	result, err := tableCollection.Find(ctx.Context(), bson.M{"store_id": storeId})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "No table found")
	}

	if err := result.All(ctx.Context(), &tables); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not find any table")
	}
	fmt.Println(tables)

	return ctx.JSON(tables)
}

func GetTable(ctx *fiber.Ctx) error {
	//Collect table id and store id from params
	tableId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Create table model
	var table models.Table

	//Find table by table id and store id
	err := tableCollection.FindOne(ctx.Context(), bson.M{"store_id": storeId, "table_id": tableId}).Decode(&table)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Table not found")
	}

	fmt.Println(table)
	//Return table
	return ctx.JSON(table)
}

func CreateTable(ctx *fiber.Ctx) error {
	//
	storeId := ctx.Params("s_id")

	table := new(models.Table)

	fmt.Println(storeId)
	//Parse body and validate
	if err := utils.ParseBodyAndValidate(ctx, table); err != nil {
		return err
	}

	table.ID = primitive.NewObjectID()
	table.Table_id = table.ID.Hex()

	//Set created_at and updated_at
	table.Created_at = time.Now()
	table.Updated_at = time.Now()

	//Set store id
	table.Store_id = storeId

	//Insert table model
	_, err := tableCollection.InsertOne(ctx.Context(), table)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create table")
	}

	return ctx.JSON(table)
}

func UpdateTable(ctx *fiber.Ctx) error {
	//Collect table id and store id from params
	tableId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Create table model
	table := new(models.Table)

	//Find table by table id and store id
	err := tableCollection.FindOne(ctx.Context(), bson.M{"store_id": storeId, "table_id": tableId}).Decode(&table)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Table not found")
	}

	//Parse body and validate
	if err := utils.ParseBodyAndValidate(ctx, table); err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"number_of_guests": table.Number_of_guests,
			"table_number":     table.Table_number,
			"updated_at":       time.Now(),
		},
	}

	fmt.Println(update)

	//Find table by table id and store id
	result, err := tableCollection.UpdateOne(ctx.Context(), bson.M{"store_id": storeId, "table_id": tableId}, update)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update table")
	}

	//Return table
	return ctx.JSON(result)
}

func DeleteTable(ctx *fiber.Ctx) error {
	//Collect table id and store id from params
	tableId := ctx.Params("id")
	storeId := ctx.Params("s_id")

	//Find table by table id and store id
	result, err := tableCollection.DeleteOne(ctx.Context(), bson.M{"store_id": storeId, "table_id": tableId})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not delete table")
	}

	//Return table
	return ctx.JSON(result)
}
