package calculate

import (
	"github.com/gin-gonic/gin"
	"github.com/jonathanmorais/endoenginnering/cr-service/pkg/cr"
	"github.com/jonathanmorais/endoenginnering/isf-service/pkg/isf"
	"net/http"
)

type Request struct {
	Cb  float32 `json:"cb"`
	Isf float32 `json:"isf"`
	Bgl float32 `json:"bgl"`
}

type Response struct {
	Status           int     `json:"status"`
	InsulinDoseCarbo float32 `json:"insulin_dose"`     // insulin dose per carbohydrate
	InsulinDoseBgl   float32 `json:"insulin_dose_bgl"` // insulin dose per blood glucose level
	CarboNeeded      float32 `json:"carbo_needed"`     // carbohydrate needed to reach target bgl
}

func Calculate(c *gin.Context) {

	// push body to response struct
	var request Request
	err := c.BindJSON(&request)
	if err != nil {
		c.Err()
	}

	// calculate InsulinDoseCarbo
	var idc float32
	idc = request.Cb * cr.RatioCalc()

	// calculate InsulinDoseBgl
	var idb float32
	idb = request.Bgl / isf.InsuFactor()

	// calculate CarboNeeded
	var cbn float32
	cbn = cr.RatioCalc() / request.Isf * request.Bgl

	// get calculate values and push to body response
	var response Response
	response.Status = http.StatusOK
	response.InsulinDoseCarbo = idc
	response.InsulinDoseBgl = idb
	response.CarboNeeded = cbn

	c.JSON(http.StatusOK, response)

}
