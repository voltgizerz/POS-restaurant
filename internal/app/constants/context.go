package constants

type ContextKey string

const (
	CTXKeyUserID    ContextKey = "userID"
	CTXKeyUsername  ContextKey = "username"
	CTXKeyRoleID    ContextKey = "roleID"
	CTXKeyIsActive  ContextKey = "isActive"
	CTXKeyRequestID ContextKey = "requestID"
)
