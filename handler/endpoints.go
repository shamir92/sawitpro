package handler

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/helper"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetHello(ctx echo.Context, params generated.GetHelloParams) error {
	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

// (POST /estate)
func (s *Server) PostEstate(ctx echo.Context) error {
	var req generated.PostEstateRequest

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := helper.ValidateStruct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	output, err := s.Repository.StoreEstate(ctx.Request().Context(), repository.StoreEstateInput{
		Width:  req.Width,
		Length: req.Length,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, generated.PostEstateResponse{
		Id:      output.Id.String(),
		Message: "Hello World",
	})
}

func (s *Server) GetEstateIdDronePlan(ctx echo.Context, id string, params generated.GetEstateIdDronePlanParams) error {
	// TODO: Validate that the estate exists
	estate, err := s.Repository.GetEstateByID(ctx.Request().Context(), repository.GetEstateByIDInput{
		Id: uuid.MustParse(id),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}

	if estate.Id == uuid.Nil {
		return echo.NewHTTPError(http.StatusNotFound, "estate not found")
	}

	// TODO: Return the drone plan
	output, err := s.Repository.FetchEstateTrees(ctx.Request().Context(), uuid.MustParse(id))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "estate tree failed to be fetched")
	}

	// TODO: Return the drone plan
	if params.MaxDistance != nil {
		restPoint := helper.CalculateDroneWithMaxDistance(estate.Width, estate.Length, *params.MaxDistance, output)

		return ctx.JSON(http.StatusOK, generated.GetEstateIDDronePlanResponse{
			Distance: *params.MaxDistance,
			Rest: &struct {
				X int `json:"x"`
				Y int `json:"y"`
			}{
				X: int(restPoint.FinalRow),
				Y: int(restPoint.FinalCol),
			},
		})
	}
	return ctx.JSON(http.StatusOK, generated.GetEstateIDDronePlanResponse{
		Distance: helper.CalculateDroneDistance(estate.Width, estate.Length, output),
	})
}
func (s *Server) GetEstateIdStats(ctx echo.Context, id string) error {
	// TODO: Validate that the estate exists
	estate, err := s.Repository.GetEstateByID(ctx.Request().Context(), repository.GetEstateByIDInput{
		Id: uuid.MustParse(id),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if estate.Id == uuid.Nil {
		return echo.NewHTTPError(http.StatusNotFound, "Estate not found")
	}

	output, err := s.Repository.FetchEstateTrees(ctx.Request().Context(), uuid.MustParse(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "estate tree failed to be fetched")
	}

	var treeHeights []int
	for _, tree := range output {
		treeHeights = append(treeHeights, tree.Height)
	}

	// TODO: Return the stats
	return ctx.JSON(http.StatusOK, generated.GetEstateIdStatsResponse{
		Max:    slices.Max(treeHeights),
		Min:    slices.Min(treeHeights),
		Count:  len(treeHeights),
		Median: int(helper.Median(treeHeights)),
	})
}

func (s *Server) PostEstateIdTree(ctx echo.Context, id string) error {

	var req generated.PostEstateTreeRequest

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := helper.ValidateStruct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO: Validate that the estate exists
	estate, err := s.Repository.GetEstateByID(ctx.Request().Context(), repository.GetEstateByIDInput{
		Id: uuid.MustParse(id),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if estate.Id == uuid.Nil {
		return echo.NewHTTPError(http.StatusNotFound, "Estate not found")
	}

	// TODO: Validate that the tree does not exist
	output, err := s.Repository.StoreEstateIdTree(ctx.Request().Context(), repository.StoreEstateIdTreeInput{
		EstateId: uuid.MustParse(id),
		Height:   req.Height,
		X:        req.X,
		Y:        req.Y,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, generated.PostEstateResponse{
		Id:      output.Id.String(),
		Message: "Hello World",
	})
}
