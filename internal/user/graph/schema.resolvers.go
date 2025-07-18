package graph

import (
	"context"
	"time"

	"github.com/Thanhbinh1905/seta-training-system/internal/user/graph/model"
	usermodel "github.com/Thanhbinh1905/seta-training-system/internal/user/model"
	"github.com/Thanhbinh1905/seta-training-system/pkg/apperror"
	"github.com/Thanhbinh1905/seta-training-system/pkg/jwt"
	"github.com/Thanhbinh1905/seta-training-system/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, username, email, password string, role model.Role) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Error("Bcrypt hashing failed", zap.Error(err), zap.String("email", email))
		return nil, apperror.Internal("could not create user")
	}

	user := &usermodel.User{
		UserID:       uuid.New().String(),
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         usermodel.Role(role),
		CreatedAt:    time.Now(),
	}

	if err := r.DB.Create(user).Error; err != nil {
		return nil, apperror.Conflict("email already in use")
	}

	logger.Log.Info("New user created", zap.String("user_id", user.UserID), zap.String("email", email))

	return &model.User{
		UserID:    user.UserID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      model.Role(user.Role),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, email, password string) (*model.AuthPayload, error) {
	var user usermodel.User

	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, apperror.Unauthorized("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, apperror.Unauthorized("invalid email or password")
	}

	token, err := jwt.GenerateJWT(user.UserID, string(user.Role), r.JWTSecret, 48*time.Hour)
	if err != nil {
		logger.Log.Error("JWT generation failed", zap.Error(err), zap.String("user_id", user.UserID))
		return nil, apperror.Internal("login failed")
	}

	logger.Log.Info("User logged in", zap.String("user_id", user.UserID), zap.String("email", email))

	return &model.AuthPayload{
		Token: token,
	}, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	// Nếu bạn implement revoke token, log ở đây cũng cần cập nhật
	logger.Log.Info("User logged out")
	return true, nil
}

func (r *queryResolver) FetchUsers(ctx context.Context) ([]*model.User, error) {
	userID, ok := ctx.Value("userID").(string)
	if !ok || userID == "" {
		return nil, apperror.Unauthorized("unauthorized access")
	}

	var users []usermodel.User
	if err := r.DB.Find(&users).Error; err != nil {
		logger.Log.Error("Failed to fetch users from DB", zap.Error(err), zap.String("requested_by", userID))
		return nil, apperror.Internal("could not fetch users")
	}

	logger.Log.Info("Fetched users", zap.Int("count", len(users)), zap.String("requested_by", userID))

	result := make([]*model.User, 0, len(users))
	for _, user := range users {
		result = append(result, &model.User{
			UserID:    user.UserID,
			Username:  user.Username,
			Email:     user.Email,
			Role:      model.Role(user.Role),
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		})
	}

	return result, nil
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
