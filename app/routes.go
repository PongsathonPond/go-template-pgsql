package app

import (
	"fmt"
	"net/http"
	"os"

	"idev-cms-service/config"
	"idev-cms-service/docs"
	"idev-cms-service/service/util/logs"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginLogRus "github.com/toorop/gin-logrus"
)

type Router struct {
	app       *App
	router    *gin.Engine
	appConfig *config.Config
	log       logs.Log
	logger    *logrus.Logger
}

type Config struct {
	AppConfig *config.Config
	Log       logs.Log
	Logger    *logrus.Logger
}

func New(rc *Config) *Router {
	return &Router{
		app:       newApp(rc.AppConfig, rc.Log),
		router:    gin.New(),
		appConfig: rc.AppConfig,
		log:       rc.Log,
		logger:    rc.Logger,
	}
}

func (r *Router) RegisterRoute() *Router {
	r.routerMiddleware()
	gin.SetMode(r.appConfig.GinMode)
	//
	//// middleware
	//midAuth := r.app.mid.Authentication()
	midAuthor := r.app.mid.Authorization()

	// api v1
	v1 := r.router.Group(r.appConfig.SwaggerInfoBasePath)
	{
		v1.POST("login", r.app.authen.Login)
		v1.POST("token/refresh", r.app.tokens.Refresh)

		all := v1.Group("all")
		{
			all.GET("menus", r.app.menus.All)
			all.GET("roles", r.app.roles.All)
		}
		contentPublic := v1.Group("content")
		{
			contentPublic.GET("", r.app.content.List)
		}
	}

	// api v1 with auth
	v1auth := r.router.Group(r.appConfig.SwaggerInfoBasePath)
	{
		v1auth.POST("logout", r.app.authen.Logout)
		v1auth.POST("token/revoke", r.app.tokens.Revoke)

		users := v1auth.Group("users")
		{
			users.GET("", r.app.users.List)
			users.POST("", r.app.users.Create)
			users.GET(":id", r.app.users.Read)
			users.PUT(":id", r.app.users.Update)
			users.DELETE(":id", r.app.users.Delete)
		}

		me := v1auth.Group("me")
		{
			me.GET("", r.app.users.GetMe)
			me.PUT("", r.app.users.UpdateMe)
			me.GET("permissions", r.app.roles.GetUserPermission)
		}

		roles := v1auth.Group("roles", midAuthor)
		{
			roles.GET("", r.app.roles.List)
			roles.POST("", r.app.roles.Create)
			roles.GET(":id", r.app.roles.Read)
			roles.PUT(":id", r.app.roles.Update)
			roles.DELETE(":id", r.app.roles.Delete)
		}

		content := v1auth.Group("content", midAuthor)
		{
			//content.GET("", r.app.content.List)
			content.POST("", r.app.content.Create)
			content.GET(":id", r.app.content.Read)
			content.PUT(":id", r.app.content.Update)
			content.DELETE(":id", r.app.content.Delete)
		}

		category := v1auth.Group("category", midAuthor)
		{
			category.GET("", r.app.category.List)
			category.POST("", r.app.category.Create)
			category.GET(":id", r.app.category.Read)
			category.PUT(":id", r.app.category.Update)
			category.DELETE(":id", r.app.category.Delete)
		}
	}

	r.routeSwagger()
	r.routeHealthCheck()
	return r
}

func (r *Router) Start() {
	if err := r.router.Run(); err != nil {
		r.log.Panic(err)
	}
}

func (r *Router) routerMiddleware() {
	appENV := os.Getenv("APP_ENV")
	if appENV == "local" {
		r.router.Use(r.ginLogWithConfig())
	} else {
		r.router.Use(ginLogRus.Logger(r.logger))
	}
	r.router.Use(r.corsMiddleware())
	r.router.Use(gin.Recovery())
	r.router.Use(limit.MaxAllowed(r.appConfig.AppMaxAllowed))
}

func (r *Router) routeSwagger() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = r.appConfig.AppName
	docs.SwaggerInfo.Description = "API Specification Document"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = r.appConfig.SwaggerInfoHost
	docs.SwaggerInfo.BasePath = r.appConfig.SwaggerInfoBasePath

	docPath := ginSwagger.URL(fmt.Sprintf("//%s/swagger/doc.json", r.appConfig.SwaggerInfoHost))
	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, docPath))
}

func (r *Router) routeHealthCheck() {
	r.router.GET("/system/health", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
}

func (r *Router) ginLogWithConfig() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{
			// "/swagger/index.html",
			"/swagger/swagger-ui.css",
			"/swagger/swagger-ui-standalone-preset.js",
			"/swagger/swagger-ui-bundle.js",
			"/swagger/doc.json",
			"/system/health",
			"/favicon.ico",
		},
	})
}

func (r *Router) corsMiddleware() gin.HandlerFunc {
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowCredentials = true
	corsConf.AddAllowHeaders("Authorization", "UserGroups-Agent")
	return cors.New(corsConf)
}
