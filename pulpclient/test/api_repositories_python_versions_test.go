/*
Pulp 3 API

Testing RepositoriesPythonVersionsAPIService

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

func Test_pulpclient_RepositoriesPythonVersionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesPythonVersionsAPIService RepositoriesPythonPythonVersionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsDelete(context.Background(), pythonPythonRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesPythonVersionsAPIService RepositoriesPythonPythonVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsList(context.Background(), pythonPythonRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesPythonVersionsAPIService RepositoriesPythonPythonVersionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsRead(context.Background(), pythonPythonRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesPythonVersionsAPIService RepositoriesPythonPythonVersionsRepair", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pythonPythonRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesPythonVersionsAPI.RepositoriesPythonPythonVersionsRepair(context.Background(), pythonPythonRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}