package slice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SliceClient struct {
	Client *resourcemanager.Client
}

func NewSliceClientWithBaseURI(api environments.Api) (*SliceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "slice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SliceClient: %+v", err)
	}

	return &SliceClient{
		Client: client,
	}, nil
}
