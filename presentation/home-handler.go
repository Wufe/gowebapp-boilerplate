package presentation

import (
	"github.com/kataras/iris"
	"github.com/wufe/boilerplateprj/services"
)

// GetAlexaHandlers performs registration of alexa api endpoint
func GetHomeHandlers(party iris.Party) func(services.HomeService) {
	return func(homeService services.HomeService) {
		party.Get("/test", func(ctx iris.Context) {
			status := homeService.GetStatus()
			ctx.WriteString(status)
		})
	}
}
