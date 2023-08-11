/*
Pulp 3 API

Testing RepositoriesAnsibleVersionsAPIService

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

func Test_pulpclient_RepositoriesAnsibleVersionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesAnsibleVersionsAPIService RepositoriesAnsibleAnsibleVersionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesAnsibleVersionsAPI.RepositoriesAnsibleAnsibleVersionsDelete(context.Background(), ansibleAnsibleRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAnsibleVersionsAPIService RepositoriesAnsibleAnsibleVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesAnsibleVersionsAPI.RepositoriesAnsibleAnsibleVersionsList(context.Background(), ansibleAnsibleRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAnsibleVersionsAPIService RepositoriesAnsibleAnsibleVersionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesAnsibleVersionsAPI.RepositoriesAnsibleAnsibleVersionsRead(context.Background(), ansibleAnsibleRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAnsibleVersionsAPIService RepositoriesAnsibleAnsibleVersionsRebuildMetadata", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesAnsibleVersionsAPI.RepositoriesAnsibleAnsibleVersionsRebuildMetadata(context.Background(), ansibleAnsibleRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAnsibleVersionsAPIService RepositoriesAnsibleAnsibleVersionsRepair", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleAnsibleRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesAnsibleVersionsAPI.RepositoriesAnsibleAnsibleVersionsRepair(context.Background(), ansibleAnsibleRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}