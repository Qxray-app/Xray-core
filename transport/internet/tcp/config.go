package tcp

import (
	"github.com/qxray-app/xray-core/common"
	"github.com/qxray-app/xray-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
