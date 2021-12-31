package orders

import (
	"context"
	"fmt"
	"log"
	"strconv"

	clearCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/clear-cart"
	cartEntity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
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
	productStruct := ProductStruct{}
	for _, val := range cartRes.Products {
		strs := fmt.Sprintf("%v", val["product_id"])
		// change string to objectid
		objId, err := primitive.ObjectIDFromHex(strs)
		if err != nil {
			panic(err)
		}
		productStruct.ProductIds = append(productStruct.ProductIds, objId)
	}

	// 2. find all products with the following product_ids from cart
	cursor, err := productCollection.Find(context.TODO(), bson.M{
		"_id": bson.M{
			"$in": productStruct.ProductIds,
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
	orderCartEntity := entity.Cart{}
	orderCartEntity.ID = cartRes.ID
	orderCartEntity.AuditLog = entity.AuditLog(cartRes.AuditLog)
	orderCartEntity.UserId = cartRes.UserId
	for _, product := range products {
		for _, val := range cartRes.Products {
			// convert product_id string to mongoid
			productID := fmt.Sprintf("%v", val["product_id"])
			objID, err := primitive.ObjectIDFromHex(productID)
			if err != nil {
				panic(err)
			}
			// compare if same product ids, then assign quantity
			if objID == product.ID {
				strs := fmt.Sprintf("%v", val["quantity"])
				value, err := strconv.ParseFloat(strs, 32)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(value)
				product.Quantity = value
			}
		}
		orderCartEntity.Products = append(orderCartEntity.Products, product)
	}

	// add initial order_status
	orderStatus := entity.OrderStatus{}
	orderStatus.Title = "Order Placed (COD)"
	orderStatus.Label = "order_placed_cod"

	// 4. insert request to orders db
	result, err := orderCollection.InsertOne(context.TODO(), bson.M{
		"cart":             orderCartEntity,
		"delivery_address": order.DeliveryAddress,
		"order_status":     orderStatus,
		"audit_log":        order.AuditLog,
	})
	if err != nil {
		panic(err)
	}

	ce := cartEntity.ProductRequest{}
	ce.UserId = au.UserId
	// 5. empty / clear user's cart
	clearCart.ClearService(ce)

	// return order id string
	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex()
}
