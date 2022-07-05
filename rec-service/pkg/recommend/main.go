package recommend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonathanmorais/endoenginnering/cr-service/pkg/cr"
	"github.com/jonathanmorais/endoenginnering/isf-service/pkg/isf"
)

type Request struct {
	Cb  float64 `json:"cb"`
	Isf float64 `json:"isf"`
	Bgl float64 `json:"bgl"`
}

type Response struct {
	Status           int     `json:"status"`
	InsulinDoseCarbo float64 `json:"insulin_dose"`     // insulin dose per carbohydrate
	InsulinDoseBgl   float64 `json:"insulin_dose_bgl"` // insulin dose per blood glucose level
	CarboNeeded      float64 `json:"carbo_needed"`     // carbohydrate needed to reach target bgl
}

func Recommend(c *gin.Context) Response {

	// push body to response struct
	var request Request
	err := c.BindJSON(&request)
	if err != nil {
		c.Err()
	}

	// calculate InsulinDoseCarbo
	var idc float64
	idc = request.Cb * cr.RatioCalc()

	// calculate InsulinDoseBgl
	var idb float64
	idb = request.Bgl / isf.InsuFactor()

	// calculate CarboNeeded
	var cbn float64
	cbn = cr.RatioCalc() / request.Isf * request.Bgl

	// get calculate values and push to body response
	var response Response
	response.Status = http.StatusOK
	response.InsulinDoseCarbo = idc
	response.InsulinDoseBgl = idb
	response.CarboNeeded = cbn

	//c.JSON(http.StatusOK, response)

	return response

}
