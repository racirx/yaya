package yaya

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ServerConfig struct {
	Port string
}

type Config struct {
	ServerConfig ServerConfig `json:"server"`
}

type Server struct {
	engine *gin.Engine
	Config *Config
}

type message struct {
	Old    string
	New    string
	Err    string
	Answer string
}

type Form struct {
	Old float64 `form:"old" binding:"required,number"`
	New float64 `form:"new" binding:"required,number"`
}

func (s *Server) Initialize(conf *Config) {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Yaya Calculator",
		})
	})

	r.POST("/", func(c *gin.Context) {
		msg := new(message)
		form := new(Form)

		if err := c.ShouldBind(form); err != nil {
			log.Printf("post route: %v\n", err)
			msg.Err = "Values must be a number"
			c.HTML(http.StatusBadRequest, "index.tmpl", msg)
			return
		}

		diff := form.Old - form.New
		perc := diff / form.Old
		answ := fmt.Sprintf("%.2f", perc*100)
		msg.Answer = answ
		msg.Old = fmt.Sprintf("%.2f", form.Old)
		msg.New = fmt.Sprintf("%.2f", form.New)

		log.Printf("answer: %s\n", msg.Answer)

		c.HTML(http.StatusOK, "index.tmpl", msg)
		return
	})

	*s = Server{
		engine: r,
		Config: conf,
	}
}

func (s *Server) Run() {
	if err := s.engine.Run(fmt.Sprintf(":%s", s.Config.ServerConfig.Port)); err != nil {
		log.Fatal(fmt.Sprintf("server: %v\n", err))
	}
}
