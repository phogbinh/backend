package handler

import (
	"backend/database"
	. "backend/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
AddOrderItemToCartHandler is a function for gin to handle AddOrderItemToCart api
*/
func AddOrderItemToCartHandler(c *gin.Context) {
	var addToCart AddToCart

	productID, err := strconv.Atoi(c.Query("p_id"))
	if err != nil {
		c.Status(400)
	} else {
		addToCart.ProductID = productID
	}

	cartID, err := strconv.Atoi(c.Query("cart_id"))
	if err != nil {
		c.Status(400)
	} else {
		addToCart.CartID = cartID
	}

	quantity, err := strconv.Atoi(c.Query("quantity"))
	if err != nil {
		c.Status(400)
	} else {
		addToCart.Quantity = quantity
	}

	code := database.AddOrderItemToCart(&addToCart, database.SqlDb)

	c.Status(code)
}
