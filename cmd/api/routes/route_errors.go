package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"myApi/pkg/commons/errors"
	"net/http"
	"strings"
)

func buildHTTPErrorHandlerFunc(e *echo.Echo) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {

		var logged bool

		he, ok := err.(*echo.HTTPError)
		if ok {
			h, found := he.Message.(*errors.E)
			if found {
				he.Message = h.Key
				if len(h.Details) > 0 {
					log.Error(h.Key, ", ", strings.Join(h.Details, ", "))
					logged = true
				}
			}
		}

		h, ok := err.(*errors.E)
		if ok {
			if len(h.Details) > 0 {
				log.Error(h.Key, ", ", strings.Join(h.Details, ", "))
				logged = true
			}
			if h.Client {
				err = echo.NewHTTPError(http.StatusBadRequest, h.Key)
			} else {
				err = echo.NewHTTPError(http.StatusInternalServerError, h.Key)
			}
		}

		if !logged {
			log.Error(err)
		}

		e.DefaultHTTPErrorHandler(err, c)
	}
}
