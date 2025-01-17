package dnsforwardingrulesets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DnsForwardingRulesetsClient struct {
	Client *resourcemanager.Client
}

func NewDnsForwardingRulesetsClientWithBaseURI(api environments.Api) (*DnsForwardingRulesetsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "dnsforwardingrulesets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DnsForwardingRulesetsClient: %+v", err)
	}

	return &DnsForwardingRulesetsClient{
		Client: client,
	}, nil
}
