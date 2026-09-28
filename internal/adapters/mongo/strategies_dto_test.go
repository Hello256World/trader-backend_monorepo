package mongo

import (
	"testing"

	"github.com/Hello256World/trader-backend_monorepo/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestFromStrategyCoreToDTO(t *testing.T) {
	t.Parallel()

	validHexID := "507f1f77bcf86cd799439011"

	tests := []struct {
		name        string
		input       *domain.Strategy
		expectErr   bool
		errContains string
		validate    func(t *testing.T, dto *StrategyDTO)
	}{
		{
			name: "success - valid strategy is converted correctly",
			input: &domain.Strategy{
				ID:          validHexID,
				Name:        "Momentum Strategy",
				Description: "A strategy based on price momentum",
			},
			expectErr: false,
			validate: func(t *testing.T, dto *StrategyDTO) {
				expectedID, err := bson.ObjectIDFromHex(validHexID)
				require.NoError(t, err)

				assert.Equal(t, expectedID, dto.ID)
				assert.Equal(t, "Momentum Strategy", dto.Name)
				assert.Equal(t, "A strategy based on price momentum", dto.Description)
			},
		},
		{
			name:        "error - nil input returns error",
			input:       nil,
			expectErr:   true,
			errContains: "Invalid input Strategy",
		},
		{
			name: "error - invalid ObjectID hex format",
			input: &domain.Strategy{
				ID:          "not-a-valid-object-id",
				Name:        "Invalid Strategy",
				Description: "This should fail",
			},
			expectErr:   true,
			errContains: "invalid Strategy ID",
		},
		{
			name: "error - empty ID string",
			input: &domain.Strategy{
				ID:          "",
				Name:        "Empty ID Strategy",
				Description: "This should also fail",
			},
			expectErr:   true,
			errContains: "invalid Strategy ID",
		},
		{
			name: "success - empty name and description are allowed",
			input: &domain.Strategy{
				ID:          validHexID,
				Name:        "",
				Description: "",
			},
			expectErr: false,
			validate: func(t *testing.T, dto *StrategyDTO) {
				assert.Empty(t, dto.Name)
				assert.Empty(t, dto.Description)
			},
		},
	}

	for _, tt := range tests {
		tt := tt // جلوگیری از مشکل capture متغیر loop در Go < 1.22
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dto, err := fromStrategyCoreToDTO(tt.input)

			if tt.expectErr {
				require.Error(t, err)
				assert.Nil(t, dto)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				return
			}

			require.NoError(t, err)
			require.NotNil(t, dto)

			if tt.validate != nil {
				tt.validate(t, dto)
			}
		})
	}
}

func TestFromStrategyDTOToCore(t *testing.T) {
	t.Parallel()

	knownID := bson.NewObjectID()

	tests := []struct {
		name  string
		input StrategyDTO
		want  domain.Strategy
	}{
		{
			name: "success - normal values are mapped correctly",
			input: StrategyDTO{
				ID:          knownID,
				Name:        "MyName",
				Description: "MyDescription",
			},
			want: domain.Strategy{
				ID:          knownID.Hex(),
				Name:        "MyName",
				Description: "MyDescription",
			},
		},
		{
			name: "edge case - empty name and description are preserved as-is",
			input: StrategyDTO{
				ID:          knownID,
				Name:        "",
				Description: "",
			},
			want: domain.Strategy{
				ID:          knownID.Hex(),
				Name:        "",
				Description: "",
			},
		},
		{
			name: "edge case - zero-value ObjectID is converted to its hex form",
			input: StrategyDTO{
				ID:          bson.ObjectID{},
				Name:        "Zero ID Strategy",
				Description: "Should still map fields correctly",
			},
			want: domain.Strategy{
				ID:          bson.ObjectID{}.Hex(),
				Name:        "Zero ID Strategy",
				Description: "Should still map fields correctly",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := fromStrategyDTOToCore(tt.input)

			assert.Equal(t, tt.want, got)
		})
	}
}
