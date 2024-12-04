package webServer

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//go:embed templates/*.html
var templatesFS embed.FS

//go:embed resource/*
var staticFs embed.FS

// Config is webServer config
type Config struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

// WebServer interface ...
type WebServer interface {
	Run() error
}

// webServer general struct
type webServer struct {
	config            Config
	router            *gin.Engine
	fibonacciPrevious int
	fibonacciCurrent  int
}

// New constructor from webServer struct
func New(config Config) (WebServer, error) {

	gin.SetMode(gin.ReleaseMode)

	ws := webServer{
		config:            config,
		router:            gin.Default(),
		fibonacciPrevious: 0,
		fibonacciCurrent:  0,
	}

	err := ws.registerTemplates()
	if err != nil {
		return nil, err
	}

	ws.registerStatic()

	ws.registerControllers()

	return &ws, nil
}

// Run method start web server
func (ws *webServer) Run() error {
	addr := fmt.Sprintf("%s:%d", ws.config.Address, ws.config.Port)
	err := ws.router.Run(addr)
	if err != nil {
		return err
	}
	return nil
}

// registerControllers register controllers in router
func (ws *webServer) registerControllers() {

	const (
		indexEndpoint = "/"
	)

	ws.router.GET(indexEndpoint, ws.indexController)

}

// indexController index controller
func (ws *webServer) indexController(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", ws.fibonacciCurrent)

	if ws.fibonacciCurrent == 0 {
		ws.fibonacciPrevious = 0
		ws.fibonacciCurrent = 1
		return
	}

	ws.fibonacciPrevious, ws.fibonacciCurrent = ws.fibonacciCurrent, ws.fibonacciPrevious+ws.fibonacciCurrent
}

// registerTemplates register templates
func (ws *webServer) registerTemplates() error {
	tmpl, err := template.ParseFS(templatesFS, "templates/*.html")
	if err != nil {
		return err
	}

	ws.router.SetHTMLTemplate(tmpl)

	return nil
}

// registerStatic register static route
func (ws *webServer) registerStatic() {
	const staticRoute = "/static"
	ws.router.StaticFS(staticRoute, http.FS(staticFs))
}
