// Copyright 2020-2022 the Pinniped contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package fositestorage

import (
	"github.com/ory/fosite"

	"go.pinniped.dev/internal/constable"
	"go.pinniped.dev/internal/oidc/clientregistry"
	"go.pinniped.dev/internal/psession"
)

const (
	ErrInvalidRequestType     = constable.Error("requester must be of type fosite.Request")
	ErrInvalidClientType      = constable.Error("requester's client must be of type clientregistry.Client")
	ErrInvalidSessionType     = constable.Error("requester's session must be of type PinnipedSession")
	StorageRequestIDLabelName = "storage.pinniped.dev/request-id"
)

func ValidateAndExtractAuthorizeRequest(requester fosite.Requester) (*fosite.Request, error) {
	request, ok1 := requester.(*fosite.Request)
	if !ok1 {
		return nil, ErrInvalidRequestType
	}
	_, ok2 := request.Client.(*clientregistry.Client)
	if !ok2 {
		return nil, ErrInvalidClientType
	}
	_, ok3 := request.Session.(*psession.PinnipedSession)
	if !ok3 {
		return nil, ErrInvalidSessionType
	}

	return request, nil
}
