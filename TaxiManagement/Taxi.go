package TaxiManagement

import (
	"errors"
	"fmt"
	"time"
)

type ITaxis interface {//with that interface the controller talk with the taxis struct
	CreateTaxi(int, string, bool,ReportsFromTaxis) (bool,error)
	GetTaxis() []*Taxi
}

type ITaxi interface {//with that interface the taxi station talk with the taxi
	StarTotWork(chan string)
	AddTravel(*Travel)
	GetId()int
	IsAvailable() bool
	RegisterToStation()
	GetNextTravel() *Travel
}
type Taxis struct{
	taxis []*Taxi
}

type Taxi struct {
	Id             int
	CurrentAddress string
	Available      bool
	MyTravelQueue  Queue //type:*Travel
	TaxiStation    ReportsFromTaxis
}

func NewTaxis() *Taxis{
	return &Taxis{taxis: []*Taxi{}}
}
func (taxis *Taxis) CreateTaxi(id int,currentAddress string,available bool,taxiStation ReportsFromTaxis) (bool,error){
	if currentAddress==""{
		return false,errors.New("Must add current address")
	}
	newTaxi:=Taxi{Id:id,CurrentAddress:currentAddress,Available: available,TaxiStation:taxiStation }
	newTaxi.RegisterToStation()
	taxis.taxis=append(taxis.taxis,&newTaxi)
	return true,nil
}

func (taxis *Taxis) GetTaxis() []*Taxi{
	return taxis.taxis
}



func (taxi *Taxi) GetId()int{
	return taxi.Id
}

func (taxi *Taxi) IsAvailable() bool{
	return taxi.Available
}

func (taxi *Taxi) RegisterToStation(){
	taxi.TaxiStation.AddTaxiToStation(taxi)
}

func (taxi *Taxi) GetNextTravel() *Travel{
	return taxi.MyTravelQueue.Dequeue().(*Travel)
}

func (taxi *Taxi) AddTravel(travel *Travel){
	taxi.MyTravelQueue.Enqueue(travel)
}
func (taxi *Taxi) StarTotWork(channel chan string){
		for emptyQueue:=taxi.MyTravelQueue.IsEmpty();!emptyQueue;emptyQueue=taxi.MyTravelQueue.IsEmpty(){
			travel:=taxi.MyTravelQueue.Dequeue().(*Travel)

			taxi.TaxiStation.ProgressMessage(taxi.Id,fmt.Sprintf("Drive from : %s ----> %s for pick-up passenger: %s\n",taxi.CurrentAddress,travel.Source,travel.PassengerName))
			time.Sleep(2000*time.Millisecond)
			taxi.TaxiStation.ProgressMessage(taxi.Id,fmt.Sprintf("Arrived to %s and picked-up passenger : %s\n",travel.Source,travel.PassengerName))
			time.Sleep(2000*time.Millisecond)
			taxi.TaxiStation.ProgressMessage(taxi.Id,fmt.Sprintf("Reached the destination : %s and dropped off the passenger : %s\n",travel.Destination,travel.PassengerName))
			taxi.CurrentAddress=travel.Destination
		}
		channel<-fmt.Sprintf("Taxi number : %d has no travels \n\n",taxi.Id)
}
