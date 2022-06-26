package dbgorm

import "github.com/google/uuid"

type UUIDField struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
}
