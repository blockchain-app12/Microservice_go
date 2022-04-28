package application

import(
	 "github.com/julienschmidt/httprouter"
	 "log"
	 "net/http"
	 "Golang-Microservice/controller"
)

func StartApplication() {
	router := httprouter.New()
	router.POST("/create", controller.Create())
	router.GET("/getall", controller.GetAll())
	router.GET("/getbyid", controller.GetItem())
	router.PATCH("/update/:item_id", controller.UpdateItem())
	router.DELETE("/deleteitem", controller.DeleteItem())
	
	
	log.Fatal(http.ListenAndServe(":8000", router))
}