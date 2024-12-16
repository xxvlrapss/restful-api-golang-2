package rest

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/xxvlrapss/go_restorant_app.git/internal/model/constant"
	"github.com/xxvlrapss/go_restorant_app.git/internal/tracing"
	"github.com/xxvlrapss/go_restorant_app.git/internal/usecase/resto"
)

/*
Note:
Echo also provides JWT middleware out of the box, but in this example
we will be making ourselves a custom middleware
ref: https://echo.labstack.com/middleware/jwt/
*/

func GetAuthMiddleware(restoUsecase resto.Usecase) *authMiddleware {
	return &authMiddleware{
		restoUsecase: restoUsecase,
	}
}

type authMiddleware struct {
	restoUsecase resto.Usecase
}

func (am *authMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.CreateSpan(c.Request().Context(), "CheckAuth")
		defer span.End()

		sessionData, err := GetSessionData(c.Request())
		if err != nil {
			return &echo.HTTPError{
				Code:     401,
				Message:  err.Error(),
				Internal: err,
			}
		}

		userID, err := am.restoUsecase.CheckSession(ctx, sessionData)
		if err != nil {
			return &echo.HTTPError{
				Code:     401,
				Message:  err.Error(),
				Internal: err,
			}
		}

		authContext := context.WithValue(ctx, constant.AuthContextKey, userID)
		c.SetRequest(c.Request().WithContext(authContext))

		if err := next(c); err != nil {
			return err
		}

		return nil
	}
}