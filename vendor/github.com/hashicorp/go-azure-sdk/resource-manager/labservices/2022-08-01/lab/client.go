package lab

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabClient struct {
	Client *resourcemanager.Client
}

func NewLabClientWithBaseURI(api environments.Api) (*LabClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "lab", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LabClient: %+v", err)
	}

	return &LabClient{
		Client: client,
	}, nil
}
