package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/iakozlov/crime-app-gateway/internal/domain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

var (
	ErrTimout     = errors.New("the request execution timeout has expired")
	ErrInvalidJWT = errors.New("invalid JWT token")
	jwtKey        = "user"
)

type CrimeAnalysisService interface {
	CrimeAnalysis(ctx context.Context, request domain.CrimeAnalysisRequest, username string) (*domain.CrimeAnalysisResponse, error)
}

type UserHistoryService interface {
	History(ctx context.Context, request domain.UserHistoryRequest) (*domain.UserHistoryResponse, error)
	AddHistory(ctx context.Context, request domain.UserHistoryItem, username string) error
}

type CrimeAppHandler struct {
	crimeAnalysisService CrimeAnalysisService
	userHistoryService   UserHistoryService
	log                  *logrus.Logger
}

func NewCrimeAnalysisHandler(analysisService CrimeAnalysisService, historyService UserHistoryService, log *logrus.Logger) *CrimeAppHandler {
	return &CrimeAppHandler{
		crimeAnalysisService: analysisService,
		userHistoryService:   historyService,
		log:                  log,
	}
}

func (h CrimeAppHandler) InitRoutes(e *echo.Echo, timeout time.Duration, jwtSecret string) {
	users := e.Group(
		"crime",
		middleware.RequestID(),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		}),
		middleware.Recover(),
		middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Skipper:      middleware.DefaultSkipper,
			ErrorMessage: ErrTimout.Error(),
			OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
				c.Error(err)
			},
			Timeout: timeout * time.Second,
		}),
		middleware.JWTWithConfig(middleware.JWTConfig{
			Claims:     &jwtCrimeClaims{},
			SigningKey: []byte(jwtSecret),
		}),
	)
	users.POST("/analysis", h.GetCrimeAnalysisHandler)
	users.POST("/history", h.GetUserHistory)
}

// GetCrimeAnalysisHandler godoc
// @Summary      represents crime analysis
// @Description  get info about crime analysis at some point
// @Tags         analysis
// @Accept       json
// @Produce      json
// @Param        id   body      domain.CrimeAnalysisRequest  true  "CrimeAnalysisInfo"
// @Success      200  {object}  string
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Failure default {object} echo.HTTPError
// @Router       /crime/analysis [post]
// @Security BearerAuth
func (h CrimeAppHandler) GetCrimeAnalysisHandler(c echo.Context) error {
	ctx := c.Request().Context()

	token, ok := c.Get(jwtKey).(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidJWT.Error())
	}
	crimeClaims, ok := token.Claims.(*jwtCrimeClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidJWT.Error())
	}

	request := domain.CrimeAnalysisRequest{}
	if err := c.Bind(&request); err != nil {
		h.log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	username := crimeClaims.Login

	response, err := h.crimeAnalysisService.CrimeAnalysis(ctx, request, username)
	if err != nil {
		h.log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"analysis": response,
	})
}

// GetUserHistory  godoc
// @Summary      represents user requests history
// @Description  get user requests history
// @Tags         history
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Failure default {object} echo.HTTPError
// @Router       /crime/history [post]
// @Security BearerAuth
func (h CrimeAppHandler) GetUserHistory(c echo.Context) error {
	ctx := c.Request().Context()

	token, ok := c.Get(jwtKey).(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidJWT.Error())
	}
	crimeClaims, ok := token.Claims.(*jwtCrimeClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidJWT.Error())
	}

	request := domain.UserHistoryRequest{UserName: crimeClaims.Login}

	response, err := h.userHistoryService.History(ctx, request)
	if err != nil {
		h.log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"history": response,
	})
}

type jwtCrimeClaims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}
