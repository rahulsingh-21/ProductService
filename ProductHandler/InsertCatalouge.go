package producthandler

import (
	"net/http"
	"time"

	helper "ProductService/Helper"
	input "ProductService/Input"
	model "ProductService/Models"

	"github.com/gin-gonic/gin"
)

func InsertProductCatalogue(ctx *gin.Context) {
	var inputCatalouge input.CatalougeInput
	if ctx.BindJSON(&inputCatalouge) != nil {
		helper.StatusBadRequest(ctx, "Invalid Input")
		return
	}
	var catalouge model.Catalouge

	catalouge.Id = inputCatalouge.Id
	catalouge.Title = inputCatalouge.Title
	catalouge.Description = inputCatalouge.Description
	catalouge.Category = inputCatalouge.Category
	catalouge.Brand = inputCatalouge.Brand
	catalouge.Availablility = inputCatalouge.Availablility
	catalouge.Price = inputCatalouge.Price
	catalouge.SalePrice = inputCatalouge.SalePrice
	catalouge.CreatedAt = time.Now()
	catalouge.IsActive = 1
	err := model.InsertCatalouge(&catalouge)
	if err != nil {
		helper.StatusBadRequest(ctx, "Not able to Insert Product Catalouge")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
