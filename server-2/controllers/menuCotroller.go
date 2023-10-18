package controllers

import (
	"restaurant-management/server-2/database"

	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")
