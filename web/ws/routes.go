package ws

import "v0rt3x/perimeter/web"

var routes = web.Routes{
	web.Route{
		"GET",
		"/:channel",
		StreamChannel,
	},
}

func GetRoutes() web.Routes {
	return routes
}
