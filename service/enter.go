package service

import "demo25/service/jztinfo"

type ServiceGroup struct {
	Jztinfos jztinfo.JztinfoService
}

var ServiceApp = new(ServiceGroup)
