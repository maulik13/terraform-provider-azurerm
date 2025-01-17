package applicationgatewayprivateendpointconnections

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGatewayPrivateEndpointConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewApplicationGatewayPrivateEndpointConnectionsClientWithBaseURI(api environments.Api) (*ApplicationGatewayPrivateEndpointConnectionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "applicationgatewayprivateendpointconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationGatewayPrivateEndpointConnectionsClient: %+v", err)
	}

	return &ApplicationGatewayPrivateEndpointConnectionsClient{
		Client: client,
	}, nil
}
