package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"

	"github.com/Mague/forex-bridge/models"
	"github.com/Mague/forex-bridge/payloads"
	"github.com/Mague/forex-bridge/repositories"
)

type Operation struct {
	ctx    *gin.Context
	router *gin.RouterGroup
}

func (this Operation) Load(router *gin.RouterGroup) {
	this.router = router
	operation := this.router.Group("/operations")
	{
		operation.GET("/", this.all)
		//operation.GET("/:id", this.byId)
		operation.POST("/add", this.add)
	}
}

// all godoc
// @Summary      List operations
// @Description  get operations
// @Tags         Operations
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Operation
// @Router       /operations [get]
func (this Operation) all(ctx *gin.Context) {
	repository := new(repositories.OperationRepository)
	operations := repository.All()

	ctx.JSON(http.StatusOK, operations)
}

// add godoc
// @Summary      Add an operation
// @Description  add by json operation
// @Tags         Operations
// @Accept       json
// @Produce      json
// @Param        operation  body      payloads.Operation  true  "Add operation"
// @Success      200      {object}  models.Operation
// @Failure      400      {object}  httputil.HTTPError
// @Failure      500      {object}  httputil.HTTPError
// @Router       /operations/add [post]
func (this Operation) add(ctx *gin.Context) {
	/* jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))
	*/
	var operationsPayload payloads.Operation
	if err := ctx.ShouldBindJSON(&operationsPayload); err != nil {
		fmt.Println(err)
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var operations []models.Operation
	for i := 0; i < len(operationsPayload); i++ {
		operations = append(operations, models.Operation{
			AccountId:   operationsPayload[i].AccountId,
			OrderNumber: operationsPayload[i].OrderNumber,
			Symbol:      operationsPayload[i].Symbol,
			OrderType:   operationsPayload[i].OrderType,
			Price:       operationsPayload[i].Price,
			StopLoss:    operationsPayload[i].StopLoss,
			TakeProfit:  operationsPayload[i].TakeProfit,
		})
	}

	repository := new(repositories.OperationRepository)
	result := repository.Create(operations)
	if result.Error != nil {
		fmt.Println(result.Error)
		httputil.NewError(ctx, http.StatusBadRequest, result.Error)
		return
	} else {

		ctx.JSON(http.StatusCreated, operations)
	}

}
