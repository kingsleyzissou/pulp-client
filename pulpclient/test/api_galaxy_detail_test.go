/*
Pulp 3 API

Testing GalaxyDetailAPIService

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

func Test_pulpclient_GalaxyDetailAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GalaxyDetailAPIService GalaxyCollectionDetailGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionHref string

		resp, httpRes, err := apiClient.GalaxyDetailAPI.GalaxyCollectionDetailGet(context.Background(), ansibleCollectionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
