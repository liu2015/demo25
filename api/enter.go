package api

import "demo25/api/jztinfo"

type ApiGroup struct {
	Jztinfos jztinfo.JztinfoGroup
}

var ApiApp = new(ApiGroup)
