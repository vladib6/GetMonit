package Controller

import (
	"awesomeProject/TaxiManagement"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (  //Path for http request
	Travel="/travel"
	GetNotifications="/getnotifications"
	Clear="/clear"
)

func  RunServer(c IController) {
	mux := http.NewServeMux()
	mux.Handle("/clear", c)
	mux.Handle("/travel", c)
	mux.Handle("/getnotifications", c)
	buildHandler := http.FileServer(http.Dir("UI/getmonit/build"))
	mux.Handle("/", buildHandler)

	srv := &http.Server{
		Handler:      mux,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server started on PORT 8080")
	log.Fatal(srv.ListenAndServe())
}

func (controller *Controller) CreateTravelReq(w http.ResponseWriter,r *http.Request){
	var requestBody TaxiManagement.Travel
	if err:=json.NewDecoder(r.Body).Decode(&requestBody);err!=nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error to create travel,data incorrect"))
		return
	}
	_, err := controller.CreateTravel(requestBody.PassengerName,requestBody.Source,requestBody.Destination)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error to create Travel in server"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Travel request was recorded ,We are looking for a free drive, This may take several minutes... "))

}

func  (controller *Controller) GetNotificationsReq(w http.ResponseWriter,r *http.Request){
	jsonbytes,err:=json.Marshal(controller.GetNewNotifications())
	if err!=nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error to get notifications from server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonbytes)
}

func  (controller *Controller) Clear(w http.ResponseWriter,r *http.Request){
	controller.ClearNotifications();
	w.WriteHeader(http.StatusOK)
}