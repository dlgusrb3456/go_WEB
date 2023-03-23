package WEB_UUID

import (
	"github.com/google/uuid"
)

func GetUUID() string {
	id := uuid.New()
	return id.String()
}
