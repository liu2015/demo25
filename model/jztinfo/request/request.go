package request

import (
	"demo25/model/request"
	"time"
)

type Findjztinfo struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Name string `json:"name" form:"name" `
	request.PageInfo
}
