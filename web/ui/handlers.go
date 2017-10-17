package ui

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MenuLink struct {
	Icon       string
	Color      string
	Attributes string
	ToolTip    string
}

type Exploit struct {
	Id          int
	Name        string
	Type        string
	Enabled     bool
	Description string
}

type MenuLinks []MenuLink

type Exploits []Exploit

func IndexView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":  "Overview",
		"links":  GetLinks(),
		"path":   c.Request.RequestURI,
		"stream": "index",
	})
}

func AgentView(c *gin.Context) {
	c.HTML(http.StatusOK, "agent.html", gin.H{
		"title":  "Overview",
		"links":  GetLinks(),
		"path":   c.Request.RequestURI,
		"stream": "agents",
		"menu": MenuLinks{
			MenuLink{"add", "green", "", "Add Agent"},
			MenuLink{"remove", "red", "", "Delete Agent"},
		},
	})
}

func ExploitView(c *gin.Context) {
	c.HTML(http.StatusOK, "exploit.html", gin.H{
		"title":  "Overview",
		"links":  GetLinks(),
		"path":   c.Request.RequestURI,
		"stream": "exploit",
		"menu": MenuLinks{
			MenuLink{"add", "green", "onclick=\"$('#add_exploit').modal('open');\"", "Add Exploit"},
		},
		"exploits": Exploits{
			Exploit{
				1,
				"Test Exploit",
				"Python",
				true,
				"Description",
			},
			Exploit{
				2,
				"Test Exploit",
				"GoLang",
				false,
				"Description",
			},
			Exploit{
				3,
				"Test Exploit",
				"Bash",
				true,
				"Description",
			},
			Exploit{
				1,
				"Test Exploit",
				"Python",
				false,
				"Description",
			},
			Exploit{
				2,
				"Test Exploit",
				"GoLang",
				false,
				"Description",
			},
			Exploit{
				3,
				"Test Exploit",
				"Bash",
				false,
				"Description",
			},
		},
	})
}

func FlagView(c *gin.Context) {
	c.HTML(http.StatusOK, "flag.html", gin.H{
		"title":  "Overview",
		"links":  GetLinks(),
		"path":   c.Request.RequestURI,
		"stream": "flag",
	})
}
