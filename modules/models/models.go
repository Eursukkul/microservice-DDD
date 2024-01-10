package models

//Token based
//offset based
type (
	PaginateReq struct {
		Start string `query:"start" validate:"max=64"`
		Limit int    `query:"limit" validate:"required,min=2,max=10"` //จำนวนข้อมูลที่ต้องการ
	}

	PaginateRes struct {
		Data  any           `json:"data"`
		Limit int           `json:"limit"`
		Total int64         `json:"total"`
		First FirstPaginate `json:"first"`
		Next  NextPaginate  `json:"next"`
	}

	FirstPaginate struct {
		Href string `json:"href"`
	}

	NextPaginate struct {
		Start string `json:"start"`
		Href  string `json:"href"`
	}

)