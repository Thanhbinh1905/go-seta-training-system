package handler

import (
	"net/http"

	"github.com/Thanhbinh1905/seta-training-system/internal/team/dto"
	"github.com/Thanhbinh1905/seta-training-system/internal/team/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TeamHandler struct {
	service service.TeamService
}

func NewTeamHandler(s service.TeamService) *TeamHandler {
	return &TeamHandler{service: s}
}

func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var req dto.CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.service.CreateTeam(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team"})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func (h *TeamHandler) GetTeamByID(c *gin.Context) {
	teamID, err := uuid.Parse(c.Param("team_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid team ID"})
		return
	}

	team, err := h.service.GetTeamByID(c.Request.Context(), teamID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "team not found"})
		return
	}

	c.JSON(http.StatusOK, team)
}

func (h *TeamHandler) GetTeamsByUserID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	teams, err := h.service.GetTeamsByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch teams"})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func (h *TeamHandler) AddMember(c *gin.Context) {
	teamID, _ := uuid.Parse(c.Param("team_id"))
	userID, _ := uuid.Parse(c.Param("user_id"))

	err := h.service.AddMember(c.Request.Context(), teamID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add member"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *TeamHandler) AddManager(c *gin.Context) {
	teamID, _ := uuid.Parse(c.Param("team_id"))
	userID, _ := uuid.Parse(c.Param("user_id"))

	err := h.service.AddManager(c.Request.Context(), teamID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add manager"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *TeamHandler) RemoveMember(c *gin.Context) {
	teamID, _ := uuid.Parse(c.Param("team_id"))
	userID, _ := uuid.Parse(c.Param("user_id"))

	err := h.service.RemoveMember(c.Request.Context(), teamID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove member"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *TeamHandler) RemoveManager(c *gin.Context) {
	teamID, _ := uuid.Parse(c.Param("team_id"))
	userID, _ := uuid.Parse(c.Param("user_id"))

	err := h.service.RemoveManager(c.Request.Context(), teamID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove manager"})
		return
	}
	c.Status(http.StatusNoContent)
}
