/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package internal

import (
	"github.com/shashidharatd/federation-dns/pkg/dnsprovider/providers/google/clouddns/internal/interfaces"
	dns "google.golang.org/api/dns/v1"
)

// Compile time check for interface adherence
var _ interfaces.ManagedZone = ManagedZone{}

type ManagedZone struct{ impl *dns.ManagedZone }

func (m ManagedZone) Name() string {
	return m.impl.Name
}

func (m ManagedZone) Id() uint64 {
	return m.impl.Id
}

func (m ManagedZone) DnsName() string {
	return m.impl.DnsName
}
