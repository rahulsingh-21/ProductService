package router

import (
	middleware "ProductService/Middleware"
	orderhandler "ProductService/OrderHandler"
	producthandler "ProductService/ProductHandler"

	"github.com/gin-gonic/gin"
)

func RouteDispatcher() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	productApis := router.Group("/apis/productservice")
	orderApis := router.Group("/apis/orderservice")

	productApis.POST("/insertcatalogue", producthandler.InsertProductCatalogue)
	productApis.GET("/getcatalogue", producthandler.GetCatalogue)
	productApis.GET("/getallcatalogue", producthandler.GetAllCatalogue)
	productApis.GET("/getallbycategory", producthandler.GetAllCatalogueByCategory)
	orderApis.POST("/placeorder", orderhandler.PlaceOrder)
	orderApis.PATCH("/updateorder", orderhandler.UpdateOrderState)
	// orderApis.POST("/user/registration", handler.Registration)
	// {
	// }
	return router
}
