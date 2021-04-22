package event_store

type UserEventStoreNats struct {
}

func NewUserEventStoreNats() UserEventStore {
	return UserEventStoreNats{}
}

func (store UserEventStoreNats) OnUserCreated(f func(UserCreatedMessage)) {

}

func (store UserEventStoreNats) OnUserDeleted() {

}

func (store UserEventStoreNats) OnUserUpdated() {

}

func (store UserEventStoreNats) PublishUserCreated() {

}

func (store UserEventStoreNats) PublishUserUpdated() {

}

func (store UserEventStoreNats) PublishUserDeleted() {

}
