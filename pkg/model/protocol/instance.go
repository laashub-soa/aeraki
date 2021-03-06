// Copyright Aeraki Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protocol

import (
	"strings"
)

// Instance defines network protocols for ports
type Instance string

const (
	// DUBBO declares that the port carries dubbo traffic.
	Dubbo Instance = "Dubbo"
	// Thrift declares that the port carries Thrift traffic.
	Thrift Instance = "Thrift"
	// Mongo declares that the port carries MongoDB traffic.
	Mongo Instance = "Mongo"
	// Redis declares that the port carries Redis traffic.
	Redis Instance = "Redis"
	// MySQL declares that the port carries MySQL traffic.
	MySQL Instance = "MySQL"
	// Unsupported - value to signify that the protocol is unsupported.
	Unsupported Instance = "UnsupportedProtocol"
)

// Parse from string ignoring case
func Parse(s string) Instance {
	switch strings.ToLower(s) {
	case "dubbo":
		return Dubbo
	case "thrift":
		return Thrift
	case "mongo":
		return Mongo
	case "redis":
		return Redis
	case "mysql":
		return MySQL
	}

	return Unsupported
}

// IsDubbo is true for protocols that use Dubbo as transport protocol
func (i Instance) IsDubbo() bool {
	switch i {
	case Dubbo:
		return true
	default:
		return false
	}
}

// IsThrift is true for protocols that use Thrift as transport protocol
func (i Instance) IsThrift() bool {
	switch i {
	case Thrift:
		return true
	default:
		return false
	}
}

func (i Instance) IsUnsupported() bool {
	return i == Unsupported
}

func (i Instance) ToString() string {
	return string(i)
}

func GetLayer7ProtocolFromPortName(name string) Instance {
	s := strings.Split(name, "-")
	if len(s) > 1 {
		return Parse(s[1])
	}
	return Unsupported
}

/*func (i Instance) String() string {
	return string(i)
}*/
