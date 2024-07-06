package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/helper"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
)

// (POST /estate)
func (s *Server) PostEstate(ctx echo.Context) error {
	var req generated.PostEstateRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := helper.ValidateStruct(req); err != nil {
		return err
	}

	output, err := s.Repository.StoreEstate(ctx.Request().Context(), repository.StoreEstateInput{
		Width:  req.Width,
		Length: req.Length,
	})

	return ctx.JSON(http.StatusOK, generated.PostEstateResponse{
		Id:      output.Id,
		Message: "Hello World",
	})
}
