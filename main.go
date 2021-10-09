package main

import (
	"awesomeProject/Controller"
)


type IController= Controller.IController
func main() {
	var controller IController=Controller.NewController()
	controller.CreateTaxi(1, "Herzliya", true)
	controller.CreateTaxi(2,"Raanana",true)
	controller.CreateTaxi(3,"Netanya",true)
	go controller.Run()

	Controller.RunServer(controller)

}







