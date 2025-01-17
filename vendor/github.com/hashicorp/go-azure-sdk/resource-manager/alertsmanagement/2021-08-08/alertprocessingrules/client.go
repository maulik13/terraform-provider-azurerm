package alertprocessingrules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertProcessingRulesClient struct {
	Client *resourcemanager.Client
}

func NewAlertProcessingRulesClientWithBaseURI(api environments.Api) (*AlertProcessingRulesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "alertprocessingrules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AlertProcessingRulesClient: %+v", err)
	}

	return &AlertProcessingRulesClient{
		Client: client,
	}, nil
}
