package model

import "time"

type Note struct {
	NoteID    string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title     string `gorm:"not null"`
	Body      string
	FolderID  string    `gorm:"type:uuid"`
	OwnerID   string    `gorm:"type:uuid"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	NoteShares []NoteShare `gorm:"foreignKey:NoteID"`
}
