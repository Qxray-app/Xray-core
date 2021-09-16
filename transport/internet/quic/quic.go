package quic

import (
	"github.com/qxray-app/xray-core/common"
	"github.com/qxray-app/xray-core/transport/internet"
)

//go:generate go run github.com/qxray-app/xray-core/common/errors/errorgen

// Here is some modification needs to be done before update quic vendor.
// * use bytespool in buffer_pool.go
// * set MaxReceivePacketSize to 1452 - 32 (16 bytes auth, 16 bytes head)
//
//

const protocolName = "quic"
const internalDomain = "quic.internal.example.com"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
