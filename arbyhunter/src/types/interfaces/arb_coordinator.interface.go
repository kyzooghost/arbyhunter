package types

import (
	"context"

	dtos "arbyhunter/src/types/dtos"
	models "arbyhunter/src/types/models"
)

type IArbCoordinator interface {
	LaunchNodeAdaptor(ctx context.Context, dto dtos.LaunchNodeAdaptorDTO) models.UserResponse
	AddPool(ctx context.Context, dto dtos.AddPoolDTO) models.UserResponse
	HealthCheck(ctx context.Context) models.UserResponse
}
