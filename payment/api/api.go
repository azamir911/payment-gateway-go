package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"payment/data"

	//"github.com/gorilla/mux"
	//"github.com/rs/zerolog/log"
	"net/http"
	"payment/processor"
	"payment/service"
	"payment/validator"
	"strconv"
)

func Serve() {
	//r := mux.NewRouter()
	//r.Methods("GET").Path("/index").Handler(http.HandlerFunc(index2))
	//
	//fmt.Println("Listening on localhost:8080...")
	//err := http.ListenAndServe(":8080", r)
	//if err != nil {
	//	log.Fatal().Msgf("%v", err)
	//}
	//
	engine := gin.Default()
	engine.GET("/index", index)
	engine.GET("/payments", getAllInvoice)
	engine.GET("/payment/:id", getInvoice)
	engine.POST("/payment", postInvoice)
	engine.POST("/close", closeChannels)

	if err := engine.Run("localhost:8080"); err != nil {
		log.Logger.Panic().Msgf("Error starting service %v", err)
	}

}

//func index2(writer http.ResponseWriter, request *http.Request) {
//	writeResponse(writer, http.StatusOK, "Welcome to the Payments App!", nil)
//}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// writeResponse is a helper method that allows to write and HTTP status & response
func writeResponse(w http.ResponseWriter, status int, data interface{}, err error) {
	resp := Response{
		Data: data,
	}
	if err != nil {
		resp.Error = fmt.Sprint(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if status != http.StatusOK {
		w.WriteHeader(status)
	}
	err = json.NewEncoder(w).Encode(data)
	if err := err; err != nil {
		log.Logger.Err(err).Msgf("error encoding resp %v", resp)
		//fmt.Fprintf(w, "error encoding resp %v:%s", resp, err)
	}
}

func index(context *gin.Context) {
	writeResponse(context.Writer, http.StatusOK, "Welcome to the Payments App!", nil)
}

func getAllInvoice(context *gin.Context) {
	all := service.GetInstance().GetAll()
	writeResponse(context.Writer, http.StatusOK, all, nil)
}

func getInvoice(context *gin.Context) {
	id := context.Param("id")
	invoice, err := strconv.Atoi(id)
	if err != nil {
		writeResponse(context.Writer, http.StatusBadRequest, nil, err)
		return
	}
	transaction, err := service.GetInstance().Get(invoice)
	if err != nil {
		writeResponse(context.Writer, http.StatusNotFound, nil, err)
		return
	}
	writeResponse(context.Writer, http.StatusOK, transaction, nil)
}

func postInvoice(context *gin.Context) {
	decoder := json.NewDecoder(context.Request.Body)
	transaction := data.NewEmptyTransaction()
	if err := decoder.Decode(transaction); err != nil {
		writeResponse(context.Writer, http.StatusBadRequest, nil, err)
		return
	}
	service.GetInstance().Save(*transaction)
	writeResponse(context.Writer, http.StatusCreated, "Create new transaction", nil)
}

func closeChannels(context *gin.Context) {
	service.GetInstance().Close()
	validator.GetInstance().Close()
	processor.GetInstance().Close()
	writeResponse(context.Writer, http.StatusOK, "Channels closed", nil)
}
