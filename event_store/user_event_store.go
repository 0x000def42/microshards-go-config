package event_store

type UserEventStore interface {
	PublishUserCreated()
	PublishUserUpdated()
	PublishUserDeleted()
	OnUserCreated(f func(UserCreatedMessage))
	OnUserUpdated()
	OnUserDeleted()
}

type Message struct {
}

type UserCreatedMessage struct {
}
