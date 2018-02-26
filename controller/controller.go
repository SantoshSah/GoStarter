package controller

import (
	"GoStarter/config"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

// Controller implemented router of project
type Controller struct {
	Config config.Main
	Gin    *gin.Engine
}

// MyCustomClaims includes fields from jwt token.
type MyCustomClaims struct {
	UserID   int64  `json:"id"`
	Email    string `json:"email"`
	Language string `json:"language"`
	jwt.StandardClaims
}

func maxRequestsAtOnce(n int) gin.HandlerFunc {
	s := make(chan struct{}, n)
	return func(c *gin.Context) {
		s <- struct{}{}
		defer func() { <-s }()
		c.Next()
	}
}

func initGin() *gin.Engine {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(maxRequestsAtOnce(50))
	g.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge:         50 * time.Second,
	}))

	return g
}

func New(config config.Main) *Controller {
	return &Controller{
		Config: config,
		Gin:    initGin(),
	}
}

// Start initialize endpoints and launch http server.
func (c *Controller) Start() error {
	c.Gin.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
	//c.Gin.GET("/api/v1/login", c.login)

	// CMS endpoints. For access to this endpoints need access_token in header
	admin := c.Gin.Group("")
	admin.Use(c.checkAccessToken)
	admin.POST("/api/v1/signup", c.signUp)
	/*
		admin.POST("/api/v1/signin", c.signIn)

		// We can't use just a.Gin.Serve here because Gin will say / collides with /v1 stuff.
		c.Gin.Use(c.checkCookie)
		c.Gin.Use(static.Serve("/", static.LocalFile("./dist", true)))

		private := c.Gin.Group("")
		private.Use(c.authentication)
		private.POST("/api/v1/logout", c.logout)
	*/
	return c.Gin.Run(":" + c.Config.Port)
}

func (c *Controller) checkAccessToken(ctx *gin.Context) {
	accessToken := ctx.Request.Header.Get("Authorization")
	if accessToken != "Bearer "+c.Config.Auth.AccessToken {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "access_token is not valid",
		})
		ctx.Abort()
	}
}
