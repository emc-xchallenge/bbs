package format

import "github.com/gogo/protobuf/proto"

type Version byte

const (
	V0 Version = 0
)

//go:generate counterfeiter . Versioner
type Versioner interface {
	MigrateFromVersion(v Version) error
	Validate() error
	Version() Version
}

//go:generate counterfeiter . ProtoVersioner
type ProtoVersioner interface {
	proto.Message
	Versioner
}
