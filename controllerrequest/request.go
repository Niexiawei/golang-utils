package controllerrequest

type BasePageParamsRequest struct {
	Page     *int `binding:"omitempty,numeric,min=1" json:"page" form:"page" label:"page"`
	PageSize *int `binding:"omitempty,numeric,min=1" json:"pageSize" form:"pageSize" label:"pageSize"`
}

func (b *BasePageParamsRequest) SetPageParams(page, pageSize *int) {
	if b.Page != nil {
		*page = *b.Page
	}
	if b.PageSize != nil {
		*pageSize = *b.PageSize
	}
}
