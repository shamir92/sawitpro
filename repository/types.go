// This file contains types that are used in the repository layer.
package repository

import (
	"github.com/google/uuid"
)

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type GetEstateByIDInput struct {
	Id uuid.UUID
}

type GetEstateByIDOutput struct {
	Id     uuid.UUID
	Width  int
	Length int
}

type StoreEstateInput struct {
	Width  int
	Length int
}

type StoreEstateOutput struct {
	Id uuid.UUID
}

type StoreEstateIdTreeInput struct {
	EstateId uuid.UUID
	Height   int
	X        int
	Y        int
}

type StoreEstateIdTreeOutput struct {
	Id uuid.UUID
}

type FetchEstateStatsInput struct {
	EstateId uuid.UUID
}

type FetchEstateStatsOutput struct {
	Count  int
	Min    int
	Max    int
	Median int
}

type Estate struct {
	Id     uuid.UUID
	Width  int
	Length int
}

type EstateTree struct {
	Id     uuid.UUID
	Height int
	X      int
	Y      int
}
