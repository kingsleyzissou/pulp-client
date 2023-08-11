/*
Pulp 3 API

Testing ContentPackageIndicesAPIService

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

func Test_pulpclient_ContentPackageIndicesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentPackageIndicesAPIService ContentDebPackageIndicesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentPackageIndicesAPI.ContentDebPackageIndicesCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackageIndicesAPIService ContentDebPackageIndicesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentPackageIndicesAPI.ContentDebPackageIndicesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackageIndicesAPIService ContentDebPackageIndicesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debPackageIndexHref string

		resp, httpRes, err := apiClient.ContentPackageIndicesAPI.ContentDebPackageIndicesRead(context.Background(), debPackageIndexHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
