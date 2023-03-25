package orderhandler

import (
	"net/http"
	"time"

	helper "ProductService/Helper"
	input "ProductService/Input"
	model "ProductService/Models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PlaceOrder(ctx *gin.Context) {
	var orderInput input.OrderInput
	if ctx.BindJSON(&orderInput) != nil {
		helper.StatusBadRequest(ctx, "Invalid Input")
		return
	}
	m := make(map[input.CatalougeInput]int64)
	for i := 0; i < len(orderInput); i++ {
		m[orderInput[i]]++
	}
	var totalPremiumOrders, totalPrice int
	totalPremiumOrders, totalPrice = 0, 0
	for order, count := range m {
		if count > 10 {
			helper.StatusBadRequest(ctx, "Order count greater than 10")
			return
		}
		if order.Category == "Premium" {
			totalPremiumOrders += int(count)
		}
		totalPrice += int(order.SalePrice) * int(count)
		err := model.UpdateCatalouge(order.Id, int(count))
		if err != nil {
			helper.StatusBadRequest(ctx, "Not able to Update Product Catalouge")
			return
		}
	}
	if totalPremiumOrders >= 3 {
		totalPrice -= totalPrice / 10
	}
	var orderdetails model.OrderDetails
	orderdetails.OrderNo = uuid.New().String()
	orderdetails.OrderStatus = "Placed"
	orderdetails.OrderPrice = int64(totalPrice)
	err := model.InsertOrderDetails(&orderdetails)
	if err != nil {
		helper.StatusBadRequest(ctx, "Not able to Insert Order Details")
		return
	}
	ctx.JSON(http.StatusOK, orderdetails)
}

func UpdateOrderState(ctx *gin.Context) {
	orderState := ctx.DefaultQuery("order_status", "Placed")
	orderNo := ctx.DefaultQuery("order_no", "")
	var DispatchDate time.Time
	if orderState == "Dispatched" {
		DispatchDate = time.Now()
	}
	err := model.UpdateOrderState(orderNo, orderState, DispatchDate)
	if err != nil {
		helper.StatusBadRequest(ctx, "Not able to Update Product Catalouge")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
