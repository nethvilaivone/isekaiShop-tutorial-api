package custom

import "github.com/labstack/echo/v4"

type ErrorMessage struct {
	Message string `json: "message"`
}


func CustomErrorMessage(ct echo.Context , statcode int , message string ) error {
	return  ct.JSON(statcode , ErrorMessage{Message: message})
}


