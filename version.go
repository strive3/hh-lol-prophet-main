package hh_lol_prophet

import "github.com/strive3/hh-lol-prophet/global"

var (
	APPVersion = "0.2.3"
	Commit     = "dev"
	BuildTime  = ""
	BuildUser  = ""
)

func init() {
	global.SetAppInfo(global.AppInfo{
		Version:   APPVersion,
		Commit:    Commit,
		BuildUser: BuildUser,
		BuildTime: BuildTime,
	})
}
