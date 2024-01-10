package users

import (
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		EmployeeId               string `json:"employee_id"`
		EmployeePositionId       string `json:"employee_position_id"`
		EmployeePositionLevel    string `json:"employee_level"`
		EmployeePositionName     string `json:"employee_name"`
		Title                    string `json:"title"`
		FirstName                string `json:"first_name"`
		LastName                 string `json:"last_name"`
		NameEn                   string `json:"name_en"`
		Email                    string `json:"email"`
		JobfieldId               string `json:"job_field_id"`
		JobfieldName             string `json:"job_field_name"`
		JobId                    string `json:"job_id"`
		JobName                  string `json:"job_name"`
		OrganizationAbbreviation string `json:"organization_abbreviation"`
		OrganizationUpperId      string `json:"organization_upper_id"`
		OrganizationUpperName    string `json:"organization_upper_name"`
	}

)
