package server

import (
	"net/http"
	"reservation-service/db"
	"time"

	"github.com/gin-gonic/gin"
)

type ReservationController struct{}

type getReservationRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type getReservationListRequest struct {
	Offset int32 `form:"offset"`
	Limit  int32 `form:"limit" binding:"required,min=1,max=20"`
}

type createReservationRequest struct {
	Station_id int64     `json:"station_id" db:"station_id"`
	User_id    int64     `json:"user_id" db:"user_id"`
	Start      time.Time `json:"start" db:"start"`
	End        time.Time `json:"end" db:"end"`
}

type updateReservationRequest struct {
	Station_id int64     `json:"station_id" db:"station_id"`
	User_id    int64     `json:"user_id" db:"user_id"`
	Start      time.Time `json:"start" db:"start"`
	End        time.Time `json:"end" db:"end"`
}

func (server *Server) GetByID(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getReservationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Execute query.
	result, err := server.store.GetByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) GetAll(ctx *gin.Context) {

	// Check if request has parameters offset and limit for pagination.
	var req getReservationListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := db.ListReservationParam{
		Offset: req.Offset,
		Limit:  req.Limit,
	}

	// Execute query.
	result, err := server.store.GetAll(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) Create(ctx *gin.Context) {

	// Check if request has all required fields in json body.
	var req createReservationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := db.CreateReservationParam{
		Station_id: req.Station_id,
		User_id:    req.User_id,
		Start:      req.Start,
		End:        req.End,
	}

	// Execute query.
	result, err := server.store.Create(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (server *Server) Update(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var reqID getReservationRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Check if request has all required fields in json body.
	var req updateReservationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := db.UpdateReservationParam{
		Station_id: req.Station_id,
		User_id:    req.User_id,
		Start:      req.Start,
		End:        req.End,
	}

	// Execute query.
	result, err := server.store.Update(ctx, arg, reqID.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (server *Server) Delete(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getReservationRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Execute query.
	if err := server.store.Delete(ctx, req.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
