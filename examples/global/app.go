package global

import "time"

type app struct {
	Name    string
	Build   string
	Version string
	Date    time.Time

	// 启动时间
	LaunchTime time.Time
	Uptime     time.Duration

	Env string

	Host string
	Port string

	//网站根路径
	WEBROOT string
}

var App = app{}
