package access

type Value string

const (
	DENY  Value = "deny"
	ALLOW Value = "allow"
)

// Setup actors

type ActorFactory struct {
}

func (factory *ActorFactory) System() SystemActor {
	return SystemActor{}
}

func (factory *ActorFactory) GuestUser() GuestUserActor {
	return GuestUserActor{}
}

func (factory *ActorFactory) RegistredUser(id string) RegistredUserActor {
	return RegistredUserActor{
		PersistenceUserActor{
			Id:   id,
			role: USER_ACTOR_ROLE_REGISTRED,
		},
	}
}

func (factory *ActorFactory) ConfirmedUser(id string) ConfirmedUserActor {
	return ConfirmedUserActor{
		PersistenceUserActor{
			Id:   id,
			role: USER_ACTOR_ROLE_CONFIRMED,
		},
	}
}

func (factory *ActorFactory) AdminUser(id string) AdminUserActor {
	return AdminUserActor{
		PersistenceUserActor{
			Id:   id,
			role: USER_ACTOR_ROLE_ADMIN,
		},
	}
}

const (
	ACTOR_SYSTEM           ActorName = "system"
	ACTOR_GUEST_USER       ActorName = "guest_user"
	ACTOR_PERSISTENCE_USER ActorName = "persistence_user"
	ACTOR_REGISTED_USER    ActorName = "registred_user"
	ACTOR_CONFIRMED_USER   ActorName = "confirmed_user"
	ACTOR_ADMIN_USER       ActorName = "admin_user"
)

type ActorName string

type Actor interface {
	Name() ActorName
}

type UserActorRole int

// Define consts for value of RoleType field
const (
	USER_ACTOR_ROLE_REGISTRED UserActorRole = 1
	USER_ACTOR_ROLE_CONFIRMED UserActorRole = 2
	USER_ACTOR_ROLE_ADMIN     UserActorRole = 3
)

type PersistenceUserActor struct {
	Id   string
	role UserActorRole
}

type SystemActor struct {
}
type GuestUserActor struct {
}
type RegistredUserActor struct {
	PersistenceUserActor
}
type ConfirmedUserActor struct {
	PersistenceUserActor
}
type AdminUserActor struct {
	PersistenceUserActor
}

func (actor PersistenceUserActor) Name() ActorName {
	return ACTOR_GUEST_USER
}

func (actor SystemActor) Name() ActorName {
	return ACTOR_SYSTEM
}

func (actor GuestUserActor) Name() ActorName {
	return ACTOR_GUEST_USER
}

func (actor RegistredUserActor) Name() ActorName {
	return ACTOR_REGISTED_USER
}
func (actor ConfirmedUserActor) Name() ActorName {
	return ACTOR_CONFIRMED_USER
}

func (actor AdminUserActor) Name() ActorName {
	return ACTOR_ADMIN_USER
}
