package getrecommendations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRecommendationsClient struct {
	Client *resourcemanager.Client
}

func NewGetRecommendationsClientWithBaseURI(api environments.Api) (*GetRecommendationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "getrecommendations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GetRecommendationsClient: %+v", err)
	}

	return &GetRecommendationsClient{
		Client: client,
	}, nil
}
