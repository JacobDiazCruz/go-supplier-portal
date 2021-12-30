package orders

import (
	"context"
	"encoding/json"

	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/orders/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

type Param struct {
	id string
}

// statuses:
// 1. Order placed (COD Payment)
// 1. To Ship
// 2. To Receive
// 3. Completed
// 4. Canceled
func GetService(id string) entity.Order {
	order := entity.Order{}
	var result bson.M
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	// query
	query := bson.M{"_id": objID}
	if err2 := orderCollection.FindOne(context.TODO(), query).Decode(&result); err != nil {
		panic(err2)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &order)
	return order
}
