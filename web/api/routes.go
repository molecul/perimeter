package api

import "v0rt3x/perimeter/web"

var routes = web.Routes{
	web.Route{
		"POST",
		"/test",
		Test,
	},
}

func GetRoutes() web.Routes {
	return routes
}
