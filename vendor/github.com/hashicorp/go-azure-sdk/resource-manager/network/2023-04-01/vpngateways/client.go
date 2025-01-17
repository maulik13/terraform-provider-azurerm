package vpngateways

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnGatewaysClient struct {
	Client *resourcemanager.Client
}

func NewVpnGatewaysClientWithBaseURI(api environments.Api) (*VpnGatewaysClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "vpngateways", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VpnGatewaysClient: %+v", err)
	}

	return &VpnGatewaysClient{
		Client: client,
	}, nil
}
