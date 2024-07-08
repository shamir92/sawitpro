// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"

	"github.com/google/uuid"
)

type RepositoryInterface interface {
	StoreEstate(ctx context.Context, input StoreEstateInput) (output StoreEstateOutput, err error)
	StoreEstateIdTree(ctx context.Context, input StoreEstateIdTreeInput) (output StoreEstateIdTreeOutput, err error)
	GetEstateByID(ctx context.Context, input GetEstateByIDInput) (output Estate, err error)
	FetchEstateTrees(ctx context.Context, estateId uuid.UUID) (outputs []EstateTree, err error)
}
