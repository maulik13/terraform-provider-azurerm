package privatezones

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateZonesClient struct {
	Client *resourcemanager.Client
}

func NewPrivateZonesClientWithBaseURI(api environments.Api) (*PrivateZonesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "privatezones", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateZonesClient: %+v", err)
	}

	return &PrivateZonesClient{
		Client: client,
	}, nil
}
