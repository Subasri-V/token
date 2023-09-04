package interfaces

import "token-eg/models"

type ICustomer interface {
	CreateCustomer(customer *models.UserSignIn) (*models.DBResponse, error)
}
