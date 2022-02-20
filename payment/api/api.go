package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment/processor"
	"payment/service"
	"payment/validator"
	"strconv"
)

func Serve() {

	engine := gin.Default()
	engine.GET("/index", index)
	engine.GET("/payments", getAllInvoice)
	engine.GET("/payment/:id", getInvoice)
	engine.POST("/payment", postInvoice)
	engine.POST("/close", closeChannels)

	engine.Run("localhost:8080")

}

func index(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Welcome to the Payments App!")
}

func getAllInvoice(context *gin.Context) {
	all := service.GetInstance().GetAll()
	msg := fmt.Sprintf("%v", all)
	context.IndentedJSON(http.StatusOK, msg)
}

func getInvoice(context *gin.Context) {
	id := context.Param("id")
	invoice, _ := strconv.Atoi(id)
	transaction, _ := service.GetInstance().Get(invoice)
	msg := fmt.Sprintf("%v", transaction)
	context.IndentedJSON(http.StatusOK, msg)
}

func postInvoice(context *gin.Context) {
	context.Body
	json.N
	context.IndentedJSON(http.StatusCreated, "Create new invoice")
}

func closeChannels(context *gin.Context) {
	service.GetInstance().Close()
	validator.GetInstance().Close()
	processor.GetInstance().Close()
	context.IndentedJSON(http.StatusCreated, "Channels closed")
}
