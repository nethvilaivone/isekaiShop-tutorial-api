package custom 

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	EchoRequest interface {
		Bind(obj any) error
	}
	customEchoRequest struct {
		ctx      echo.Context
		validate validator.Validate
	}
)

var (
	once               sync.Once
	validationInstance *validator.Validate
)

func NewCustomEchoResquest(echorequest echo.Context) EchoRequest {
	once.Do(func() {
		validationInstance = validator.New()
	})
	return &customEchoRequest{
		ctx:      echorequest,
		validate: *validationInstance,
	}

}

func (ctx *customEchoRequest) Bind(obj any) error {
	
	if err := ctx.Bind(obj); err != nil {

	}

	if err := ctx.validate.Struct(obj); err != nil {
		return err
	}

	return nil
}
