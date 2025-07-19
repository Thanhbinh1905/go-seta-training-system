package dto

type CreateTeamRequest struct {
	TeamName string `json:"teamName" binding:"required"`
	Managers []struct {
		ManagerID   string `json:"managerId"`
		ManagerName string `json:"managerName"`
	} `json:"managers"`

	Members []struct {
		MemberID   string `json:"memberId"`
		MemberName string `json:"memberName"`
	} `json:"members"`
}
