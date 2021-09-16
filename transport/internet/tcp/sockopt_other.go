// +build !linux,!freebsd

package tcp

import (
	"github.com/qxray-app/xray-core/common/net"
	"github.com/qxray-app/xray-core/transport/internet"
)

func GetOriginalDestination(conn internet.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
