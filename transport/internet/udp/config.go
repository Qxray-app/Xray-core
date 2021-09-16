package udp

import (
	"github.com/qxray-app/xray-core/common"
	"github.com/qxray-app/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
