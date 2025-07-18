package model

type NoteShare struct {
	NoteID string `gorm:"type:uuid;primaryKey"`
	UserID string `gorm:"type:uuid;primaryKey"`
	Access string `gorm:"check:access IN ('read','write');not null"`
}
