package api

import (
	"api-gateway/internal/core"
	"fmt"
	"github.com/labstack/echo/v4"
)

func Receive(c echo.Context) error {
	var request core.Request
	if err := c.Bind(&request); err != nil {
		return c.JSON(200, core.ErrResult(core.NewError("1010", fmt.Sprintf("参数解析失败,详细信息: %s", err.Error()))))
	}
	result := core.Exec(&request)
	return c.JSON(200, result)
}
