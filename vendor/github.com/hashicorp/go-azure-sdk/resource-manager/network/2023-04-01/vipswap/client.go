package vipswap

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VipSwapClient struct {
	Client *resourcemanager.Client
}

func NewVipSwapClientWithBaseURI(api environments.Api) (*VipSwapClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "vipswap", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VipSwapClient: %+v", err)
	}

	return &VipSwapClient{
		Client: client,
	}, nil
}
