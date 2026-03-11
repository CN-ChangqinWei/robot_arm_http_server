package rest

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type ForwardService interface {
	GetTopics() (res []domain.Forward, err error)
	GetTopicInfo(topic string) (res domain.Forward, err error)
}

type ForwardHandler struct {
	Service ForwardService
}

func NewForwardHandler(e *echo.Echo, svc ForwardService) {
	handler := &ForwardHandler{
		Service: svc,
	}
	e.GET("/forward/topics", handler.GetTopics)
	e.GET("/forward/info/:topic", handler.GetTopicInfo)
}

func (f *ForwardHandler) GetTopics(c echo.Context) (err error) {

	res, err := f.Service.GetTopics()
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}
func (f *ForwardHandler) GetTopicInfo(c echo.Context) (err error) {
	topic := c.Param("topic")

	res, err := f.Service.GetTopicInfo(topic)
	if nil != err {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, res)

}
