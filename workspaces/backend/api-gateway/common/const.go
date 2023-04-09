package common

import "github.com/google/uuid"

const CurrentUser = "user"

type Requester interface{
	GetUserId() uuid.UUID
	GetEmail() string
	GetRole() string
}