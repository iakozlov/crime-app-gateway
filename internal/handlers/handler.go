package handlers

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github/iakozlov/crime-app-gateway/internal/domain"
	"net/http"
	"time"
)

var (
	ErrTimout = errors.New("the request execution timeout has expired")
)

type CrimeAnalysisService interface {
	CrimeAnalysis(ctx context.Context, request domain.CrimeAnalysisRequest) (domain.CrimeAnalysisResponse, error)
}

type UserHistoryService interface {
	History(ctx context.Context, request domain.UserHistoryRequest) (*domain.UserHistoryResponse, error)
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

func (h CrimeAppHandler) InitRoutes(e *echo.Echo, timeout time.Duration) {
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
	)
	users.POST("/analysis", h.GetCrimeAnalysisHandler)
	//users.POST("/login", h.SignInHandler)
}

// SignInHandler godoc
// @Summary      provides signing in operation
// @Description  authorize user in crime-app microservices ecosystem
// @Tags         analysis
// @Accept       json
// @Produce      json
// @Param        id   body      domain.CrimeAnalysisRequest  true  "User's consisted of login and password"
// @Success      200  {object}  string
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Failure default {object} echo.HTTPError
// @Router       /crime/analysis [post]
func (h CrimeAppHandler) GetCrimeAnalysisHandler(c echo.Context) error {
	ctx := c.Request().Context()
	request := domain.CrimeAnalysisRequest{}
	if err := c.Bind(&request); err != nil {
		h.log.Error(err)
		//todo: сделать маппинг статус кодов в названия
		return echo.NewHTTPError(400, err)
	}

	response, err := h.crimeAnalysisService.CrimeAnalysis(ctx, request)
	if err != nil {
		h.log.Error(err)
		return echo.NewHTTPError(500, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"analysis": response,
	})
}
