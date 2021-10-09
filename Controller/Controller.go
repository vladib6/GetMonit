package Controller

import (
	"awesomeProject/Notifications"
	"awesomeProject/TaxiManagement"
	"errors"
	"net/http"
)

type ITaxiManagement= TaxiManagement.ITaxiManagement
type ITaxis= TaxiManagement.ITaxis
func NewController() *Controller{
	return &Controller{
		taxiStation:TaxiManagement.NewTaxiManagemnt(),
		taxis: TaxiManagement.NewTaxis()}
}

type IController interface {
	CreateTravel(string,string,string) (bool,error )
	CreateTaxi(int,string,bool)(bool,error)
	Run()
	ServeHTTP(w http.ResponseWriter,r *http.Request)
	GetNewNotifications() []Notifications.Message
	ClearNotifications()
}

type Controller struct{
	taxiStation ITaxiManagement
	taxis 		ITaxis
}


func (controller *Controller) ServeHTTP(w http.ResponseWriter,r *http.Request){
	switch s:=r.URL.String();s {
	case Travel:
		controller.CreateTravelReq(w,r)
		return

	case GetNotifications:
		controller.GetNotificationsReq(w,r)
		return

	case Clear:
		controller.Clear(w,r)
	}
}

func (controller *Controller) GetNewNotifications() []Notifications.Message{
	return controller.taxiStation.GetNotifications()
}

func (controller *Controller) 	ClearNotifications() {
	controller.taxiStation.ClearNotifications()
}

func (controller *Controller) CreateTravel(name string,source string,destination string)(bool,error){
	return controller.taxiStation.CreateTravel(name,source,destination)
}

func (controller *Controller) Run(){
	controller.taxiStation.OpenTaxiStation()
}

func (controller *Controller) CreateTaxi(id int,currentAddress string,available bool) (bool,error){
	s,ok:=controller.taxiStation.(*TaxiManagement.TaxiManagement)
	if !ok {
		return false,errors.New("Create Taxi Failed ,no suitable taxi management ")
	}
	return controller.taxis.CreateTaxi(id,currentAddress,available,s)
}

