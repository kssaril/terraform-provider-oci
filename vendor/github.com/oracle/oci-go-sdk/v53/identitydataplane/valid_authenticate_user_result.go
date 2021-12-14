// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"github.com/oracle/oci-go-sdk/v53/common"
)

// ValidAuthenticateUserResult The representation of ValidAuthenticateUserResult
type ValidAuthenticateUserResult struct {

	// The tenant name.
	TenantInput *string `mandatory:"true" json:"tenantInput"`

	// The user name.
	UserInput *string `mandatory:"true" json:"userInput"`

	// The resolved principal.
	ResolvedPrincipal *AuthenticationPrincipal `mandatory:"true" json:"resolvedPrincipal"`
}

func (m ValidAuthenticateUserResult) String() string {
	return common.PointerString(m)
}
