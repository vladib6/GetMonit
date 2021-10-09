package Notifications

type Message struct {
	TaxiId int `json:"taxiId"`
	Message string `json:"message"`
}

type INotification interface {
	Add(int,string)
	GetAll()[]Message
	Clear()
}
func NewNotificationsCenter() *NotificationsCenter{
	return &NotificationsCenter{notifications: []Message{}}
}
type NotificationsCenter struct {
	notifications []Message
}

func (n *NotificationsCenter) Add(taxiId int,message string){
	n.notifications=append(n.notifications,Message{
		TaxiId:  taxiId,
		Message: message,
	})
}

func (n *NotificationsCenter) GetAll() []Message{
	return n.notifications
}

func (n *NotificationsCenter) Clear() {
	 n.notifications=n.notifications[:0]
}