package orders

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	clearCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/clear-cart"
	cartEntity "gitlab.com/JacobDCruz/supplier-portal/src/carts/entity"
	getCart "gitlab.com/JacobDCruz/supplier-portal/src/carts/get"
	database "gitlab.com/JacobDCruz/supplier-portal/src/config"
	addSellerOrder "gitlab.com/JacobDCruz/supplier-portal/src/orders/add-seller-order"
	entity "gitlab.com/JacobDCruz/supplier-portal/src/orders/entity"
	productUpdate "gitlab.com/JacobDCruz/supplier-portal/src/products/update"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")
var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

type ProductStruct struct {
	ProductIds []primitive.ObjectID
}

type SelectedProduct struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name             string             `json:"name" validate:"required"`
	Quantity         float64            `json:"quantity"`
	Status           string             `json:"status" validate:"required"`
	Slug             string             `json:"slug" validate:"required"`
	Description      string             `json:"description" validate:"required"`
	Tags             []string           `json:"tags"`
	Category         string             `json:"category" validate:"required"`
	SalesInformation interface{}        `json:"sales_information" bson:"sales_information" validate:"required"`
	Variants         interface{}        `json:"variants" bson:"variants"`
	ThumbnailImage   string             `json:"thumbnail_image" bson:"thumbnail_image" validate:"required"`
	OriginalImage    string             `json:"original_image" bson:"original_image" validate:"required"`
	AuditLog         interface{}        `json:"audit_log" bson:"audit_log"`
}

// create order id and return
func PlaceOrderService(order entity.PlaceOrder, au entity.Auth) string {
	var totalAmount = 0
	var subTotalAmount = 0

	// 1. find cart and get product ids
	cartRes := getCart.GetService(au.UserId)
	productStruct := ProductStruct{}
	orderCart := entity.Cart{}
	orderCart.ID = cartRes.ID
	orderCart.UserId = cartRes.UserId
	orderCart.AuditLog = entity.AuditLog(cartRes.AuditLog)

	for _, val := range cartRes.Products {
		productStruct.ProductIds = append(productStruct.ProductIds, val.ID)
	}

	// 2. Find all products with the following product_ids from cart
	cursor, err := productCollection.Find(context.TODO(), bson.M{
		"_id": bson.M{
			"$in": productStruct.ProductIds,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	products := []entity.Product{}
	var prodResult []bson.M
	if err = cursor.All(context.TODO(), &prodResult); err != nil {
		log.Fatal(err)
	}
	// unmarshal result to products struct
	jsonData, err := json.MarshalIndent(prodResult, "", "    ")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &products)

	// 3. insert product details to cart
	// include selected quantity and selected variant
	fmt.Println(products)
	for productKey, product := range products {
		for _, cartProduct := range cartRes.Products {
			if cartProduct.ID == product.ID {
				for variantKey, _ := range product.Variants {
					for _, cartVariant := range cartProduct.Variants {
						products[productKey].Variants[variantKey].ID = cartVariant.ID
						products[productKey].Variants[variantKey].Name = cartVariant.Name
						products[productKey].Variants[variantKey].Option = cartVariant.Option
					}
				}
				product.Quantity = cartProduct.Quantity
			}
		}
		// compute subtotal
		totalProductAmount := int(product.SalesInformation.Price) * int(product.Quantity)
		subTotalAmount = subTotalAmount + totalProductAmount
		// append
		orderCart.Products = append(orderCart.Products, product)
	}

	// add total amount
	// subtotal + shipping amount
	totalAmount = totalAmount + subTotalAmount + order.ShippingAmount

	// add initial order_status
	orderStatus := entity.OrderStatus{}
	orderStatus.Title = "Order Placed (COD)"
	orderStatus.Label = "order_placed_cod"

	// 4. insert request to orders db
	result, err := orderCollection.InsertOne(context.Background(), bson.M{
		"cart": bson.M{
			"_id":       orderCart.ID,
			"user_id":   orderCart.UserId,
			"products":  orderCart.Products,
			"audit_log": orderCart.AuditLog,
		},
		"order_id":         order.OrderId,
		"delivery_address": order.DeliveryAddress,
		"note":             order.Note,
		"subtotal_amount":  subTotalAmount,
		"total_amount":     totalAmount,
		"payment_method":   order.PaymentMethod,
		"shipping_courier": order.ShippingCourier,
		"shipping_amount":  order.ShippingAmount,
		"order_status":     orderStatus,
		"audit_log":        order.AuditLog,
	})
	if err != nil {
		panic(err)
	}

	// 5. empty / clear user's cart
	ce := cartEntity.AddToCart{}
	ce.UserId = au.UserId
	clearCart.ClearService(ce)

	// 6. update product's stock
	// loop here and call updatestock for every product update
	oid := result.InsertedID.(primitive.ObjectID)
	for _, product := range orderCart.Products {
		productUpdate.UpdateStock(product.ID, product.Quantity)

		// 7. on place order, add each ordered product to the order admin list
		// add order id on each product
		// on admin account, seller can update the status of the product to out of stock
		sellerRequest := entity.SellerOrder{}
		sellerRequest.ID = oid
		sellerRequest.OrderId = order.OrderId
		sellerRequest.Product = product
		sellerRequest.Quantity = int(product.Quantity)
		sellerRequest.DeliveryAddress = order.DeliveryAddress
		sellerRequest.OrderStatus = orderStatus
		sellerRequest.AuditLog = order.AuditLog
		addSellerOrder.AddService(sellerRequest)
	}

	// return order id string
	return oid.Hex()
}
