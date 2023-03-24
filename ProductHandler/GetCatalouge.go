package producthandler

import (
	"net/http"
	"strconv"

	helper "ProductService/Helper"
	model "ProductService/Models"
	response "ProductService/Response"

	"github.com/gin-gonic/gin"
)

func GetCatalogue(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "A_1")
	title := ctx.DefaultQuery("title", "title_1")
	availability := ctx.DefaultQuery("availability", "0")

	availabilityInt, err := strconv.Atoi(availability)
	if err != nil {
		helper.StatusBadRequest(ctx, "Invalid Input")
		return
	}
	var catalouge model.Catalouge
	err = model.GetCatalouge(&catalouge, id, title, availabilityInt)
	if err != nil {
		helper.StatusBadRequest(ctx, "Not able to Fetch Product Catalouge")
		return
	}

	ctx.JSON(http.StatusOK, response.Catalouge(catalouge))
}

func GetAllCatalogue(ctx *gin.Context) {
	var catalouge []model.Catalouge
	catalouge, err := model.GetAllCatalouge()
	if err != nil {
		helper.StatusBadRequest(ctx, "Not able to Fetch Product Catalouge")
		return
	}

	ctx.JSON(http.StatusOK, catalouge)
}

func GetAllCatalogueByCategory(ctx *gin.Context) {
	category := ctx.DefaultQuery("category", "Premium")
	var catalouge []model.Catalouge
	catalouge, err := model.GetAllCatalougeByCategory(category)
	if err != nil {
		helper.StatusBadRequest(ctx, "Not able to Fetch Product Catalouge")
		return
	}

	ctx.JSON(http.StatusOK, catalouge)
}
