/*
Pulp 3 API

Testing ExportersPulpAPIService

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

func Test_pulpclient_ExportersPulpAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExportersPulpAPIService ExportersCorePulpCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ExportersPulpAPI.ExportersCorePulpCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersPulpAPIService ExportersCorePulpDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpExporterHref string

		resp, httpRes, err := apiClient.ExportersPulpAPI.ExportersCorePulpDelete(context.Background(), pulpExporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersPulpAPIService ExportersCorePulpList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ExportersPulpAPI.ExportersCorePulpList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersPulpAPIService ExportersCorePulpPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpExporterHref string

		resp, httpRes, err := apiClient.ExportersPulpAPI.ExportersCorePulpPartialUpdate(context.Background(), pulpExporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersPulpAPIService ExportersCorePulpRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpExporterHref string

		resp, httpRes, err := apiClient.ExportersPulpAPI.ExportersCorePulpRead(context.Background(), pulpExporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExportersPulpAPIService ExportersCorePulpUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpExporterHref string

		resp, httpRes, err := apiClient.ExportersPulpAPI.ExportersCorePulpUpdate(context.Background(), pulpExporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
