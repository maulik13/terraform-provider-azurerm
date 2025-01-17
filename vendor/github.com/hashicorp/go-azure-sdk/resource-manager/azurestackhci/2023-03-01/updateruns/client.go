package updateruns

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRunsClient struct {
	Client *resourcemanager.Client
}

func NewUpdateRunsClientWithBaseURI(api environments.Api) (*UpdateRunsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "updateruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UpdateRunsClient: %+v", err)
	}

	return &UpdateRunsClient{
		Client: client,
	}, nil
}
