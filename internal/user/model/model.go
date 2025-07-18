package model

import (
	"time"

	assetsmodel "github.com/Thanhbinh1905/seta-training-system/internal/assets/model"
	teammodel "github.com/Thanhbinh1905/seta-training-system/internal/team/model"
)

type Role string

const (
	Manager Role = "manager"
	Member  Role = "member"
)

type User struct {
	UserID       string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username     string    `gorm:"not null"`
	Email        string    `gorm:"not null;uniqueIndex"`
	PasswordHash string    `gorm:"not null"`
	Role         Role      `gorm:"type:text;check:role IN ('manager','member');not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	TeamMemberships []teammodel.TeamMember    `gorm:"foreignKey:UserID"`
	TeamManager     []teammodel.TeamManager   `gorm:"foreignKey:UserID"`
	Folders         []assetsmodel.Folder      `gorm:"foreignKey:OwnerID"`
	Notes           []assetsmodel.Note        `gorm:"foreignKey:OwnerID"`
	FolderShares    []assetsmodel.FolderShare `gorm:"foreignKey:UserID"`
	NoteShares      []assetsmodel.NoteShare   `gorm:"foreignKey:UserID"`
}
