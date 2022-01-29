package carts

import (
	"context"
	"encoding/json"
	"log"

	entity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	getCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/get"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "carts")
var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

type ProductStruct struct {
	ProductIds []primitive.ObjectID
}

type Params struct {
	ProductIds []primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
}

// 1. get all products with the item_ids from cart
// 2. Find all products with the following product_ids from cart
// 2. check if the orig product has different values with the cart product
// 3. return update the outdated item
func ValidateService(userId primitive.ObjectID) []entity.ProductResponse {

	// 1. get all products with the item_ids from cart
	cartRes := getCart.GetService(userId)
	productStruct := ProductStruct{}
	if len(cartRes.Products) <= 0 {
		// log.Fatal("No cart items.")
		return nil
	}
	for _, val := range cartRes.Products {
		productStruct.ProductIds = append(productStruct.ProductIds, val.ID)
	}

	// 2. Find all products with the following product_ids from cart
	cursor2, err := productCollection.Find(context.TODO(), bson.M{
		"_id": bson.M{
			"$in": productStruct.ProductIds,
		},
	})
	if err != nil {
		return nil
	}

	// 2.5. find all products
	products := []entity.ProductResponse{}
	var prodResult []bson.M
	if err = cursor2.All(context.TODO(), &prodResult); err != nil {
		log.Fatal(err)
	}
	// unmarshal result to products struct
	jsonData, err := json.MarshalIndent(prodResult, "", "    ")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &products)

	// 3. check if the orig product has different values with the cart product
	productsToUpdate := []entity.ProductResponse{}
	for _, cartProduct := range cartRes.Products {
		for _, origProduct := range products {
			if cartProduct.AuditLog.UpdatedAt != origProduct.AuditLog.UpdatedAt {
				productsToUpdate = append(productsToUpdate, cartProduct)
			}
		}
	}

	return productsToUpdate
}
