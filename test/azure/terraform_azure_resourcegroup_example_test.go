//go:build azure
// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"testing"

	"github.com/aztfmod/terratest-helper-caf/state"
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAzureResourceGroupExample(t *testing.T) {
	t.Parallel()

	// subscriptionID is overridden by the environment variable "ARM_SUBSCRIPTION_ID".
	subscriptionID := ""

	// Initialize terraform state.
	tfState := state.NewTerraformState(t, "launchpad")

	// Get the values fromt the state file.
	resourceGroups := tfState.GetResourceGroups()

	// Verify the resource group exists.
	for _, resourceGroup := range resourceGroups {
		name := resourceGroup.GetName()
		exists := azure.ResourceGroupExists(t, name, tfState.SubscriptionID)
		assert.True(t, exists, fmt.Sprintf("Resource group (%s) does not exist", name))
	}
}
