package controller

import (
	"math/rand"
	"net/http"

	"github.com/labstack/echo"
)

//PaymentStatusResponse is to construct response object for payment status
type PaymentStatusResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

var paymentStatuses = []string{
	"fail",
	"success",
}

//PaymentStatus return payment status
func PaymentStatus(c echo.Context) (err error) {
	response := rand.Intn(2)
	return c.JSON(http.StatusOK, response)
}
