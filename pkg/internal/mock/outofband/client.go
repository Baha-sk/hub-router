/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package outofband

import "github.com/hyperledger/aries-framework-go/pkg/client/outofband"

// MockClient is a mock out-of-band client used in tests.
type MockClient struct {
	CreateInvitationErr error
}

// CreateInvitation creates a mock outofband invitation.
func (c *MockClient) CreateInvitation([]string, ...outofband.MessageOption) (*outofband.Invitation, error) {
	if c.CreateInvitationErr != nil {
		return nil, c.CreateInvitationErr
	}

	return &outofband.Invitation{}, nil
}
