package repository

import (
	"context"

	"github.com/Thanhbinh1905/seta-training-system/internal/team/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamRepository interface {
	CreateTeam(ctx context.Context, team *model.Team) error
	GetTeamByID(ctx context.Context, teamID uuid.UUID) (*model.Team, error)
	GetTeamsByUserID(ctx context.Context, userID uuid.UUID) ([]model.Team, error)
	RemoveMemberFromTeam(ctx context.Context, teamID uuid.UUID, userID uuid.UUID) error
	RemoveManagerFromTeam(ctx context.Context, teamID uuid.UUID, userID uuid.UUID) error
	AddMemberToTeam(ctx context.Context, teamID uuid.UUID, userID uuid.UUID) error
	AddManagerToTeam(ctx context.Context, teamID uuid.UUID, userID uuid.UUID) error
}

type teamRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) TeamRepository {
	return &teamRepository{db: db}
}

func (r *teamRepository) CreateTeam(ctx context.Context, team *model.Team) error {
	return r.db.Create(team).WithContext(ctx).Error
}

func (r *teamRepository) GetTeamByID(ctx context.Context, teamID uuid.UUID) (*model.Team, error) {
	var team model.Team
	err := r.db.WithContext(ctx).Preload("Members").Preload("Managers").First(&team, "team_id = ?", teamID).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *teamRepository) GetTeamsByUserID(ctx context.Context, userID uuid.UUID) ([]model.Team, error) {
	var team []model.Team
	err := r.db.WithContext(ctx).Preload("Members").Preload("Managers").Joins("team_members", "team_members.team_id = team.team_id").Joins("team_managers", "team_managers.team_id = team.team_id").
		Where("team_members.user_id = ? OR team_managers.user_id = ?", userID, userID).
		Find(&team).Error
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (r *teamRepository) RemoveMemberFromTeam(ctx context.Context, teamID uuid.UUID, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("team_id = ? AND user_id = ?", teamID, userID).Delete(&model.TeamMember{}).Error
}

func (r *teamRepository) RemoveManagerFromTeam(ctx context.Context, teamID uuid.UUID, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("team_id = ? AND user_id = ?", teamID, userID).Delete(&model.TeamManager{}).Error
}

func (r *teamRepository) AddMemberToTeam(ctx context.Context, teamID uuid.UUID, userID uuid.UUID) error {
	member := model.TeamMember{
		TeamID: teamID.String(),
		UserID: userID.String(),
	}
	return r.db.WithContext(ctx).Create(&member).Error
}

func (r *teamRepository) AddManagerToTeam(ctx context.Context, teamID uuid.UUID, userID uuid.UUID) error {
	manager := model.TeamManager{
		TeamID: teamID.String(),
		UserID: userID.String(),
	}

	return r.db.WithContext(ctx).Create(&manager).Error
}
