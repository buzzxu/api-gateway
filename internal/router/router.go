package router

import (
	"api-gateway/internal/core/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// 3595182301@qq.com
func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: true,
		DisableStackAll:   true,
		StackSize:         4 << 10,
	}))
	routers(e)
	return e
}

func routers(e *echo.Echo) {
	e.POST("/", api.Receive)
}
