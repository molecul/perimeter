package ui

import (
	"v0rt3x/perimeter/web"
)

var routes = web.Routes{
	web.Route{
		"GET",
		"/",
		IndexView,
	},
	web.Route{
		"GET",
		"/flag/",
		FlagView,
	},
	web.Route{
		"GET",
		"/exploit/",
		ExploitView,
	},
	web.Route{
		"GET",
		"/agents/",
		AgentView,
	},
}

var links = web.Links{
	web.Link{
		"/",
		"Overview",
		"dashboard",
	},
	web.Link{
		"/flag/",
		"Flag Processor",
		"flag",
	},
	web.Link{
		"/exploit/",
		"Exploits",
		"polymer",
	},
	web.Link{
		"/agents/",
		"Remote Agents",
		"cast",
	},
}

func GetRoutes() web.Routes {
	return routes
}

func GetLinks() web.Links {
	return links
}
