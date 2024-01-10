package auth

import "time"

type (
	Credential struct {
		Id           string `json:"id" `
		Employee_Id  string `json:"employee_id"`
		RoleCode     int    `json:"role_code"`
		AcessToken   string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		CeatedAt     string `json:"created_at"`
		UpdateAt     string `json:"update_at"`
	}

	UpdateRefreshTokenReq struct {
		EmployeeId   string    `json:"player_id"`
		AccessToken  string    `json:"access_token"`
		RefreshToken string    `json:"refresh_token"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	Auth struct {
		Id           string `json:"id"`
		EmployeeId   string `json:"player_id"`
		RefreshToken string `json:"refresh_token"`
	}
)
