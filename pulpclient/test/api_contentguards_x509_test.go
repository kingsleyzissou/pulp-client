/*
Pulp 3 API

Testing ContentguardsX509APIService

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

func Test_pulpclient_ContentguardsX509APIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentguardsX509APIService ContentguardsCertguardX509Create", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentguardsX509API.ContentguardsCertguardX509Create(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsX509APIService ContentguardsCertguardX509Delete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certguardX509CertGuardHref string

		httpRes, err := apiClient.ContentguardsX509API.ContentguardsCertguardX509Delete(context.Background(), certguardX509CertGuardHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsX509APIService ContentguardsCertguardX509List", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentguardsX509API.ContentguardsCertguardX509List(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsX509APIService ContentguardsCertguardX509PartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certguardX509CertGuardHref string

		resp, httpRes, err := apiClient.ContentguardsX509API.ContentguardsCertguardX509PartialUpdate(context.Background(), certguardX509CertGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsX509APIService ContentguardsCertguardX509Read", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certguardX509CertGuardHref string

		resp, httpRes, err := apiClient.ContentguardsX509API.ContentguardsCertguardX509Read(context.Background(), certguardX509CertGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentguardsX509APIService ContentguardsCertguardX509Update", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var certguardX509CertGuardHref string

		resp, httpRes, err := apiClient.ContentguardsX509API.ContentguardsCertguardX509Update(context.Background(), certguardX509CertGuardHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}