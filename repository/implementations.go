package repository

import "context"

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) StoreEstate(ctx context.Context, input StoreEstateInput) (output StoreEstateOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "INSERT INTO estate (width, length) VALUES ($1, $2) RETURNING id", input.Width, input.Length).Scan(&output.Id)
	ir err != nil {
		return nil, err.New("failed to store estate")
	}

	return output, nil
}
