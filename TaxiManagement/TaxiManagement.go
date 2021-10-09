package TaxiManagement

import (
	"awesomeProject/Notifications"
	"awesomeProject/Utils"
	"errors"
	"fmt"
	"time"
)
type Queue=Utils.Queue

type Travel struct{
	PassengerName string `json:"name"`
	Source        string `json:"source"`
	Destination   string `json:"destination"`
}


type ReportsFromTaxis interface { //with that interface the Taxi talk with a taxi-station
	ReportAvailability(int)
	AddTaxiToStation(ITaxi)
	RemoveFromStation(int)
	ProgressMessage(int,string)
}

type ITaxiManagement interface {//with that interface the controller talk with a taxi-station
	CreateTravel(string,string,string) (bool,error)
	AddTravelToQueue(travel *Travel)
	SetTravelsToDrivers()
	OpenTaxiStation()
	GetNotifications() []Notifications.Message
	ClearNotifications()
}

type TaxiManagement struct{
	NCenter Notifications.INotification
	RegisteredTaxis map[int]ITaxi //Stores all taxis registered at the station
	travelQueue Queue //type: *Travel,  Stores all the travels that are waiting to be attached to a taxi
	taxiQueue Queue //type: ITaxi,   Stores all taxis available now
}

func NewTaxiManagemnt() *TaxiManagement{
	return &TaxiManagement{RegisteredTaxis: map[int]ITaxi{},travelQueue: Queue{},taxiQueue: Queue{},NCenter: Notifications.NewNotificationsCenter()}
}

func (manager * TaxiManagement) GetNotifications() []Notifications.Message{
	return manager.NCenter.GetAll()
}

func (manager * TaxiManagement) ClearNotifications() {
	manager.NCenter.Clear()
}

func (manger *TaxiManagement) ProgressMessage(taxiId int,message string){
	manger.NCenter.Add(taxiId,message)
}

func (manager *TaxiManagement) OpenTaxiStation(){
	for {
		manager.SetTravelsToDrivers()
		time.Sleep(6*time.Second)
	}
}

func (manager *TaxiManagement) AddTaxiToStation(taxi ITaxi){//check if it exist taxi with same id if not add the taxi to registered map and if also available add to taxi queue
	taxiId:=taxi.GetId()
	_,ok:=manager.RegisteredTaxis[taxiId]
	if !ok{
		manager.RegisteredTaxis[taxiId]=taxi
		if taxi.IsAvailable(){
			manager.AddTaxiToQueue(taxi)
		}
	}
}

func (manager *TaxiManagement) ReportAvailability(id int){
	t,ok:=manager.RegisteredTaxis[id]
	if ok {
		manager.AddTaxiToQueue(t)
	}
}

func (manager *TaxiManagement) RemoveFromStation(id int){
	delete(manager.RegisteredTaxis,id)
}


func(manager *TaxiManagement) AddTaxiToQueue(taxi ITaxi){
	manager.taxiQueue.Enqueue(taxi)
}

func(manager *TaxiManagement) CreateTravel(name,source,destination string) (bool,error){
	if name=="" || source=="" || destination==""{
			return false,errors.New("Missing details")
	}
	manager.AddTravelToQueue(&Travel{PassengerName: name, Source: source, Destination: destination})
	return true,nil
}

func(manager *TaxiManagement) AddTravelToQueue(travel *Travel){
	manager.travelQueue.Enqueue(travel)
}

func(manager *TaxiManagement) SetTravelsToDrivers(){
	if empty:=manager.taxiQueue.IsEmpty();empty==true{
		fmt.Println("NO Taxi in the station now")
		return
	}
	if empty:=manager.travelQueue.IsEmpty();empty==true{
		fmt.Println("NO waiting travels")
		return
	}
	for empty:=manager.travelQueue.IsEmpty();empty!=true;empty=manager.travelQueue.IsEmpty(){
		travel:=(manager.travelQueue.Dequeue()).(*Travel)
		taxi:=(manager.taxiQueue.Dequeue()).(ITaxi)
		for !taxi.IsAvailable(){//search the first available taxi,the unavailable taxis stay out queue(when they report availability they will enter the queue again)
			taxi=manager.taxiQueue.Dequeue().(ITaxi)
		}

		taxi.AddTravel(travel)
		manager.AddTaxiToQueue(taxi)
	}
	manager.RunAllTaxis()
}

func (manager *TaxiManagement) RunAllTaxis(){
	channel:=make(chan string)
	for _,v:=range manager.RegisteredTaxis{
		if v.IsAvailable(){
			go v.StarTotWork(channel)
		}
	}
	for i:=len(manager.RegisteredTaxis);i>0;i--{
		<-channel
	}
}

