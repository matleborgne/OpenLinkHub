package templates

import (
	"OpenLinkHub/src/config"
	"OpenLinkHub/src/dashboard"
	"OpenLinkHub/src/devices"
	"OpenLinkHub/src/logger"
	"OpenLinkHub/src/rgb"
	"OpenLinkHub/src/scheduler"
	"OpenLinkHub/src/temperatures"
	"OpenLinkHub/src/version"
	"html/template"
)

var (
	templates *template.Template
)

type Web struct {
	Title             string
	Tpl               *template.Template
	Devices           map[string]*devices.Device
	Configuration     config.Configuration
	Device            interface{}
	Lcd               interface{}
	TemperatureProbes interface{}
	Temperatures      map[string]temperatures.TemperatureProfileData
	Scheduler         *scheduler.Scheduler
	Rgb               map[string]rgb.Profile
	SystemInfo        interface{}
	CpuTemp           string
	GpuTemp           string
	Page              string
	StorageTemp       []temperatures.StorageTemperatures
	BuildInfo         *version.BuildInfo
	Dashboard         dashboard.Dashboard
}

// Init will parse all templates
func Init() {
	tpl, err := template.ParseFiles(
		"web/docs.html",
		"web/index.html",
		"web/lsh.html",
		"web/cc.html",
		"web/ccxt.html",
		"web/elite.html",
		"web/lncore.html",
		"web/lnpro.html",
		"web/cpro.html",
		"web/xc7.html",
		"web/memory.html",
		"web/k65pm.html",
		"web/k65plus.html",
		"web/k65plusW.html",
		"web/k70core.html",
		"web/k70pro.html",
		"web/k55core.html",
		"web/k100.html",
		"web/k100air.html",
		"web/k100airW.html",
		"web/st100.html",
		"web/mm700.html",
		"web/lt100.html",
		"web/psuhid.html",
		"web/katarpro.html",
		"web/ironclaw.html",
		"web/ironclawW.html",
		"web/ironclawWU.html",
		"web/slipstream.html",
		"web/nightsabreW.html",
		"web/nightsabreWU.html",
		"web/scimitar.html",
		"web/scimitarWU.html",
		"web/scimitarW.html",
		"web/rgb.html",
		"web/temperature.html",
		"web/scheduler.html",
		"web/navigation.html",
		"web/footer.html",
		"web/header.html",
		"web/404.html",
	)
	if err != nil {
		logger.Log(logger.Fields{"error": err}).Fatal("Failed to load templates")
	}

	templates = tpl
}

// GetTemplate will return a list of all templates
func GetTemplate() *template.Template {
	return templates
}
