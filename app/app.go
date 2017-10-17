package perimeter

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/v0rt3x/perimeter/config"
	"github.com/v0rt3x/perimeter/consul"
	"github.com/v0rt3x/perimeter/web/api"
	"github.com/v0rt3x/perimeter/web/ui"
	"github.com/v0rt3x/perimeter/web/ws"
)

func Run(c *config.PerimeterConfig) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Static("/static", "static")
	router.SetFuncMap(template.FuncMap{
		"attr": func(s string) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})
	router.LoadHTMLGlob("templates/*.html")

	wsGroup := router.Group("/stream")
	for _, route := range ws.GetRoutes() {
		wsGroup.Handle(route.Method, route.Path, route.Handler)
	}

	apiGroup := router.Group("/api")
	for _, route := range api.GetRoutes() {
		apiGroup.Handle(route.Method, route.Path, route.Handler)
	}

	uiGroup := router.Group("/")
	for _, route := range ui.GetRoutes() {
		uiGroup.Handle(route.Method, route.Path, route.Handler)
	}

	var client, err = consul.NewClient(fmt.Sprintf("%s:%d", c.Perimeter.Consul.Host, c.Perimeter.Consul.Port))

	if err != nil {
		panic(err)
	}

	client.RegisterService(c.Perimeter.Server.Port, "perimeter-server", "perimeter-server", []string{"perimeter", "master"})

	client.GetAgents()
	defer client.UnRegisterService("perimeter-server")

	router.Run(fmt.Sprintf("%s:%d", c.Perimeter.Server.Host, c.Perimeter.Server.Port))
}
