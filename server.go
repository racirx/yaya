package yaya

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	OldErr string
	New    string
	NewErr string
	Answer string
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
		var err error
		var oldF, newF float64
		msg := new(message)

		oldVal := c.PostForm("old")
		newVal := c.PostForm("new")
		if oldF, err = strconv.ParseFloat(oldVal, 64); err != nil {
			msg.OldErr = "Value must be a number"
		}

		if newF, err = strconv.ParseFloat(newVal, 64); err != nil {
			msg.NewErr = "Value must be a number"
		}

		if msg.OldErr != "" || msg.NewErr != "" {
			c.HTML(http.StatusBadRequest, "index.tmpl", msg)
			return
		}

		diff := oldF - newF
		perc := diff / oldF
		answ := fmt.Sprintf("%.2f", perc*100)
		msg.Answer = answ
		msg.Old = oldVal
		msg.New = newVal

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
