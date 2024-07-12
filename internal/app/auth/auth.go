package auth

import (
	"context"
	"errors"

	"github.com/voltgizerz/POS-restaurant/internal/constants"
)

type Auth struct {
	UserID       int64
	Username     string
	RoleID       int64
	IsUserActive bool
	RequestID    string
}

func GetUserLoginFromCtx(ctx context.Context) (*Auth, error) {
	userID, ok := ctx.Value(constants.CTXKeyUserID).(float64)
	if !ok {
		return nil, errors.New("user ID not found in context")
	}

	username, ok := ctx.Value(constants.CTXKeyUsername).(string)
	if !ok {
		return nil, errors.New("username not found in context")
	}

	roleID, ok := ctx.Value(constants.CTXKeyRoleID).(float64)
	if !ok {
		return nil, errors.New("role ID not found in context")
	}

	isUserActive, ok := ctx.Value(constants.CTXKeyIsActive).(bool)
	if !ok {
		return nil, errors.New("active status not found in context")
	}

	requestID, ok := ctx.Value(constants.CTXKeyRequestID).(string)
	if !ok {
		return nil, errors.New("request ID not found in context")
	}

	// Create a new User struct with the retrieved values
	user := &Auth{
		UserID:       int64(userID),
		Username:     username,
		RoleID:       int64(roleID),
		IsUserActive: isUserActive,
		RequestID:    requestID,
	}

	return user, nil
}
