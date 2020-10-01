package permissions

type Permission string

const (
	USERREGISTRATION = "UserRegistration"
)

// EventsPermission is a map of all server side events which should be handled TODO: better explanation
var EventsPermission = map[string]string{
	"/api/v1/users/registration": USERREGISTRATION,
}

const (
	PUBLIC  = "PUBLIC"
	PRIVATE = "PRIVATE"
)

var permissions = map[string]string{
	"/api/v1/users/registration": PUBLIC,
	"/api/v1/users/getUsers": PRIVATE,
}

func GetPermission() map[string]string {
	return permissions
}
