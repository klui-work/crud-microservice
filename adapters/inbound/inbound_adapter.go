package inbound

import (
	"crud/adapters/inbound/api/controller"
)

type InitialAdapter struct {
	InitialProductController controller.IProductController
}

func NewInitialAdapter(
	InitialProductController controller.IProductController,
) *InitialAdapter {
	return &InitialAdapter{
		InitialProductController: InitialProductController,
	}
}
