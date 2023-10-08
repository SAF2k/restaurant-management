package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID           primitive.ObjectID `bson:"_id"`
	Order_id     *string            `json:"order_id"`
	Table_id     *string            `json:"table_id"`
	Order_status *string            `json:"order_status" validate:"eq=OPEN|eq=CLOSE|eq="`
	Order_date   time.Time          `json:"order_date"`
	Created_at    time.Time          `json:"create_at"`
	Updated_at    time.Time          `json:"update_at"`
}

{
	"order_id": "1",
	"table_id": "1",
	"order_status": "OPEN",
}
