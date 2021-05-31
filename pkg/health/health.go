package health

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"myApi/pkg/commons"
	"net/http"
	"time"
)

func Check(c echo.Context) error {

	m := make(map[string]string)

	sqlDB, err := commons.DB.DB()
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := sqlDB.PingContext(ctx); err != nil {
		m["database"] = err.Error()
	}

	if len(m) > 0 {
		log.Errorf("health check failed %s", m)
		return c.JSON(http.StatusInternalServerError, m)
	}

	return nil
}
