package model

import "time"

type Team struct {
	TeamID    string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TeamName  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Members  []TeamMember  `gorm:"foreignKey:TeamID"`
	Managers []TeamManager `gorm:"foreignKey:TeamID"`
}
