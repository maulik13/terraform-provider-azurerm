package cluster

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterClient struct {
	Client *resourcemanager.Client
}

func NewClusterClientWithBaseURI(api environments.Api) (*ClusterClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "cluster", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ClusterClient: %+v", err)
	}

	return &ClusterClient{
		Client: client,
	}, nil
}
