package updates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatesClient struct {
	Client *resourcemanager.Client
}

func NewUpdatesClientWithBaseURI(api environments.Api) (*UpdatesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "updates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UpdatesClient: %+v", err)
	}

	return &UpdatesClient{
		Client: client,
	}, nil
}
