package types

const UserContextKey ContextKey = "user"

type AuthenticatedUser struct {
	Email    string
	LoggedIn bool
}
