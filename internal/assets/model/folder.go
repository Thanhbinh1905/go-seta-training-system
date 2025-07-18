package model

import "time"

type Folder struct {
	FolderID  string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      string    `gorm:"not null"`
	OwnerID   string    `gorm:"type:uuid"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Notes        []Note        `gorm:"foreignKey:FolderID"`
	FolderShares []FolderShare `gorm:"foreignKey:FolderID"`
}
