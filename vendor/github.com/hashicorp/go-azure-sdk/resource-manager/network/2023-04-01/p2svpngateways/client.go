package p2svpngateways

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type P2sVpnGatewaysClient struct {
	Client *resourcemanager.Client
}

func NewP2sVpnGatewaysClientWithBaseURI(api environments.Api) (*P2sVpnGatewaysClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "p2svpngateways", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating P2sVpnGatewaysClient: %+v", err)
	}

	return &P2sVpnGatewaysClient{
		Client: client,
	}, nil
}
