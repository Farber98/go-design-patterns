package observer

import "fmt"

/*
OBSERVER | PUB/SUB:
- Idea is to subscibe to some event that will trigger some behavior on many subscribed types.
- We uncouple an event from it's possible handlers.

OBJECTIVE:
- Useful to achieve many actions that are triggered on one event.
- Specially useful when you don't know how many actions are performed after an event in advance or there is a possibility that the number of actions is going to grow in the near future.
- Provide an event-driven architecture where one event can trigger one or more actions.
- Uncouple the actions that are performed from the event that triggers them.
- Provide more than one event that triggers the same action.

EXAMPLE: Notifier.
- Make a Publisher struct which is the one that triggers an event so it must accept new observers and remove them if necessary.
- When the Publisher struct is triggered, it must notify all its observers of the new event with the data associated.

ACCEPTANCE CRITERIA:
- We must have a Publisher with a NotiifyObservers method that accepts a message as an argument and triggers Notif method on every observer subscribed.
- We must have a method to add new Subscribers to the Publisher.
- We must have a method to remove new Subscribers from the Publisher.

*/

type Observer interface {
	Notify(string)
}

type Publisher struct {
	ObserversList []Observer
}

func (p *Publisher) AddObserver(o Observer) {
	p.ObserversList = append(p.ObserversList, o)
}
func (p *Publisher) RemoveObserver(o Observer) {
	var indexToRemove int

	for i, observer := range p.ObserversList {
		if observer == o {
			indexToRemove = i
			break
		}
	}
	p.ObserversList = append(p.ObserversList[:indexToRemove], p.ObserversList[indexToRemove+1:]...)
}
func (p *Publisher) NotifyObservers(m string) {
	fmt.Printf("Publisher received message %s to notify observers\n", m)
	for _, observer := range p.ObserversList {
		observer.Notify(m)
	}
}
