package router

import "demo25/router/jztinfo"

type RouterGroup struct {
	Jztinfo jztinfo.JztinfoGroup
}

var RouterApp = new(RouterGroup)
