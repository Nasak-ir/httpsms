package services

import (
	"context"
	"testing"

	"github.com/NdoleStudio/httpsms/pkg/entities"
	"github.com/stretchr/testify/assert"
)

func TestBillingLimitsEnabled_DefaultsToFalse(t *testing.T) {
	t.Setenv("BILLING_LIMITS_ENABLED", "")
	assert.False(t, billingLimitsEnabled())
}

func TestBillingLimitsEnabled_RequiresExplicitTrue(t *testing.T) {
	t.Setenv("BILLING_LIMITS_ENABLED", "true")
	assert.True(t, billingLimitsEnabled())

	t.Setenv("BILLING_LIMITS_ENABLED", "invalid")
	assert.False(t, billingLimitsEnabled())
}

func TestBillingService_IsEntitledWithoutCreditLimits(t *testing.T) {
	t.Setenv("BILLING_LIMITS_ENABLED", "false")

	service := &BillingService{}
	assert.Nil(t, service.IsEntitledWithCount(context.Background(), entities.UserID("user-id"), 1_000_000))
}
