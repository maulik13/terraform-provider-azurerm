package updatesummaries

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSummariesClient struct {
	Client *resourcemanager.Client
}

func NewUpdateSummariesClientWithBaseURI(api environments.Api) (*UpdateSummariesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "updatesummaries", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UpdateSummariesClient: %+v", err)
	}

	return &UpdateSummariesClient{
		Client: client,
	}, nil
}
