package controller

import (
	"net/http"
	"math/rand"
	"github.com/labstack/echo"
)

//PaymentStatusResponse is to construct response object for payment status
type PaymentStatusResponse struct {
	Status string `json:"status"`
	Code int `json:"code"`
}

var paymentStatuses = []string{
    "fail",
    "success",
}

//PaymentStatus return payment status
func PaymentStatus(c echo.Context) (err error) {
	index := rand.Intn(2)
	response := PaymentStatusResponse{
		Status: paymentStatuses[index],
		Code: index,
	}
	return c.JSON(http.StatusOK, response)
}
