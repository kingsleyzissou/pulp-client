/*
Pulp 3 API

Testing ContentReleaseArchitecturesAPIService

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

func Test_pulpclient_ContentReleaseArchitecturesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentReleaseArchitecturesAPIService ContentDebReleaseArchitecturesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentReleaseArchitecturesAPI.ContentDebReleaseArchitecturesCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentReleaseArchitecturesAPIService ContentDebReleaseArchitecturesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentReleaseArchitecturesAPI.ContentDebReleaseArchitecturesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentReleaseArchitecturesAPIService ContentDebReleaseArchitecturesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debReleaseArchitectureHref string

		resp, httpRes, err := apiClient.ContentReleaseArchitecturesAPI.ContentDebReleaseArchitecturesRead(context.Background(), debReleaseArchitectureHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
