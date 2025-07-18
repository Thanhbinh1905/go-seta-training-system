package model

type TeamMember struct {
	TeamID string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID string `gorm:"type:uuid;primaryKey"`
}
