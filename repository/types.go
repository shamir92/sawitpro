// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type StoreEstateInput struct {
	Width  int
	Length int
}

type StoreEstateOutput struct {
	Id uuid.UUID
}
