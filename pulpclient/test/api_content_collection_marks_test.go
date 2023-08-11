/*
Pulp 3 API

Testing ContentCollectionMarksAPIService

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

func Test_pulpclient_ContentCollectionMarksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentCollectionMarksAPIService ContentAnsibleCollectionMarksCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentCollectionMarksAPI.ContentAnsibleCollectionMarksCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentCollectionMarksAPIService ContentAnsibleCollectionMarksList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentCollectionMarksAPI.ContentAnsibleCollectionMarksList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentCollectionMarksAPIService ContentAnsibleCollectionMarksRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ansibleCollectionVersionMarkHref string

		resp, httpRes, err := apiClient.ContentCollectionMarksAPI.ContentAnsibleCollectionMarksRead(context.Background(), ansibleCollectionVersionMarkHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
