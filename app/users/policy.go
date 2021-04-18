package app

import "github.com/0x000def42/microshards-go-config/utils/access"

// Policy expressions
func (service UserService) AuthorizeList(actor access.Actor) bool {
	return isSuperActor(actor)
}

func (service UserService) AuthorizeCreate(actor access.Actor) bool {
	return isSuperActor(actor) || isGuest(actor)
}

func (service UserService) AuthorizeGetOne(actor access.Actor, id string) bool {
	return isSuperActor(actor) || isOwner(actor, id)
}

func (service UserService) AuthorizeUpdate(actor access.Actor, id string) bool {
	return isSuperActor(actor) || isOwner(actor, id)
}

func (service UserService) AuthorizeDelete(actor access.Actor, id string) bool {
	return isSuperActor(actor)
}

// Helper expressions
func isSuperActor(actor access.Actor) bool {
	return actor.Name() == access.ACTOR_SYSTEM ||
		actor.Name() == access.ACTOR_ADMIN_USER
}

func isGuest(actor access.Actor) bool {
	return actor.Name() == access.ACTOR_GUEST_USER
}

func isOwner(actor access.Actor, id string) bool {
	return actor.(access.PersistenceUserActor).Id == id
}
