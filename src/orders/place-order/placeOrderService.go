package orders

import (
	"context"
	"fmt"
	"log"
	"strconv"

	getCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/get"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/orders/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")
var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

type ProductStruct struct {
	ProductIds []primitive.ObjectID
}

// create order id and return
func PlaceOrderService(order entity.PlaceOrder, au entity.Auth) string {
	// 1. find cart and get product ids
	cartRes := getCart.GetService(au.UserId)
	prod := ProductStruct{}
	for _, val := range cartRes.Products {
		strs := fmt.Sprintf("%v", val["product_id"])
		// change string to objectid
		objId, err := primitive.ObjectIDFromHex(strs)
		if err != nil {
			panic(err)
		}
		prod.ProductIds = append(prod.ProductIds, objId)
	}

	// 2. find all products with the following product_ids from cart
	cursor, err := productCollection.Find(context.TODO(), bson.M{
		"_id": bson.M{
			"$in": prod.ProductIds,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	products := []entity.Product{}
	if err = cursor.All(context.TODO(), &products); err != nil {
		log.Fatal(err)
	}

	// 3. insert product details to cart
	cartEntity := entity.Cart{}
	cartEntity.ID = cartRes.ID
	cartEntity.AuditLog = entity.AuditLog(cartRes.AuditLog)
	cartEntity.UserId = cartRes.UserId
	// book, err := json.Marshal(cartEntity)
	// fmt.Println(book)
	for _, product := range products {
		// var m map[string]interface{}
		// product["quantity"] = 3232
		for _, val := range cartRes.Products {
			// book, err := json.Marshal(product)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			strs := fmt.Sprintf("%v", val["quantity"])
			value, err := strconv.ParseFloat(strs, 32)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(value)
			product.Quantity = value
			fmt.Println("im here ererrrrrr")
			// fmt.Println(string(book))
			// fmt.Printf("t1: %T\n", val["quantity"])
			// var decoded interface{}
			// test := json.NewDecoder(io.Reader("", []byte(fmt.Sprintf("%v", product))).Decode(&decoded)
			// fmt.Println(test)
			// json.Unmarshal([]byte(fmt.Sprintf("%v", product)), &decoded)
			// test := decoded["quantity"]["type"]
			// fmt.Println(1111111111111111)
			// fmt.Println(test)
			// fmt.Println(2222222222222222)
			// product["quantity"] = val["quantity"]
			// m["quantity"] = val["quantity"]
			// newData, err := json.Marshal(m)
			// fmt.Println(val)
			cartEntity.Products = append(cartEntity.Products, product)
		}
	}

	// 4. insert request to orders db
	result, err := orderCollection.InsertOne(context.TODO(), bson.M{
		"cart":             cartEntity,
		"delivery_address": order.DeliveryAddress,
		"audit_log":        order.AuditLog,
	})
	if err != nil {
		panic(err)
	}

	// return order id string
	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
