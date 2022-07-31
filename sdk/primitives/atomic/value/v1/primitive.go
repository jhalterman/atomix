// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	valuev1 "github.com/atomix/runtime/api/atomix/runtime/atomic/value/v1"
	"github.com/atomix/runtime/sdk/pkg/runtime"
	"google.golang.org/grpc"
)

const (
	Name       = "AtomicValue"
	APIVersion = "v1"
)

var Type = runtime.NewType[valuev1.AtomicValueServer](Name, APIVersion, register, resolve)

func register(server *grpc.Server, delegate *runtime.Delegate[valuev1.AtomicValueServer]) {
	valuev1.RegisterAtomicValueServer(server, newAtomicValueServer(delegate))
}

func resolve(client runtime.Client) (valuev1.AtomicValueServer, bool) {
	if provider, ok := client.(AtomicValueProvider); ok {
		return provider.AtomicValue(), true
	}
	return nil, false
}

type AtomicValueProvider interface {
	AtomicValue() valuev1.AtomicValueServer
}