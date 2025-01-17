package virtualnetworks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualNetworksClient struct {
	Client *resourcemanager.Client
}

func NewVirtualNetworksClientWithBaseURI(api environments.Api) (*VirtualNetworksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "virtualnetworks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualNetworksClient: %+v", err)
	}

	return &VirtualNetworksClient{
		Client: client,
	}, nil
}
