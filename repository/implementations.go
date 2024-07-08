package repository

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
)

func (r *Repository) GetEstateByID(ctx context.Context, input GetEstateByIDInput) (output GetEstateByIDOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT  width, length FROM estate WHERE id = $1", input.Id).Scan(&output.Width, &output.Length)
	if err != nil {
		return
	}
	output.Id = input.Id
	return output, nil
}

func (r *Repository) StoreEstate(ctx context.Context, input StoreEstateInput) (output StoreEstateOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "INSERT INTO estate (width, length) VALUES ($1, $2) RETURNING id", input.Width, input.Length).Scan(&output.Id)
	if err != nil {
		log.Println(err)
		return StoreEstateOutput{}, errors.New("failed to store estate")
	}

	return output, nil
}

func (r *Repository) StoreEstateIdTree(ctx context.Context, input StoreEstateIdTreeInput) (output StoreEstateIdTreeOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "INSERT INTO estate_tree (estate_id, height, x, y) VALUES ($1, $2, $3, $4) RETURNING id", input.EstateId, input.Height, input.X, input.Y).Scan(&output.Id)
	if err != nil {
		log.Println(err)
		return StoreEstateIdTreeOutput{}, errors.New("failed to store tree")
	}
	return output, nil
}

func (r *Repository) FetchEstateTrees(ctx context.Context, estateId uuid.UUID) (outputs []EstateTree, err error) {
	row, err := r.Db.QueryContext(ctx, "SELECT id, height FROM estate_tree WHERE estate_id = $1", estateId)
	if err != nil {
		return
	}
	defer row.Close()

	for row.Next() {
		var output EstateTree
		row.Scan(&output.Id, &output.Height)
		outputs = append(outputs, output)
	}

	return outputs, nil
}
