package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"questionGame/service/userservice"
)

func (s Server) userLogin(c echo.Context) error {

	var req userservice.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	response, lErr := s.userSvc.Login(req)
	if lErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, lErr.Error())
	}

	return c.JSON(http.StatusOK, response)
}

func (s Server) userRegister(c echo.Context) error {

	var uReq userservice.RegisterRequest
	if err := c.Bind(&uReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := s.userSvc.Register(uReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}

func (s Server) userProfile(c echo.Context) error {
	authToken := c.Request().Header.Get("Authorization")
	claims, err := s.authSvc.ParseToken(authToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	response, err := s.userSvc.GetProfile(userservice.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)

}
