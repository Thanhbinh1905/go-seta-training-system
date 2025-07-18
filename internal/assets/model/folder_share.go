package model

type FolderShare struct {
	FolderID string `gorm:"type:uuid;primaryKey"`
	UserID   string `gorm:"type:uuid;primaryKey"`
	Access   string `gorm:"check:access IN ('read','write');not null"`
}
