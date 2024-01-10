package auth

import (
	"time"

	"gitlab.com/chalermphanEur/modules/users"
)

type (
	UserloginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	ProfileIntercepter struct {
		*users.UserProfile
		Credential *Credential `json:"credential"`
	}

	CredentialRes struct {
		Id           string    `json:"_id"`
		EmployeeId   string    `json:"employee_id"`
		RoleCode     int       `json:"role_code"`
		AccessToken  string    `json:"access_token"`
		RefreshToken string    `json:"refresh_token"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
	LogOutRequest struct {
		EmployeeId string `json:"employee_id"`
	}
)
