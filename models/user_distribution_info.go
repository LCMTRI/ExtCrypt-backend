package models

import "github.com/google/uuid"

type UserDistributionInfo struct {
	ID            uuid.UUID `gorm:"primaryKey" json:"_id"`
	SourceName    string    `json:"source_name"`
	NumOfKey      int       `json:"num_of_key"`
	SourceOwnerId int       `json:"src_owner_id"`
}
