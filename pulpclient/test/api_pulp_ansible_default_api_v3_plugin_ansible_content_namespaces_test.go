/*
Pulp 3 API

Testing PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package pulpclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/osbuild/pulp-client/pulpclient"
)

func Test_pulpclient_PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesAPIService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesCreate(context.Background(), distroBasePath).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesAPIService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesDelete(context.Background(), distroBasePath, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesAPIService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesList(context.Background(), distroBasePath).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesAPIService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesPartialUpdate(context.Background(), distroBasePath, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesAPIService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesAPI.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesRead(context.Background(), distroBasePath, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}