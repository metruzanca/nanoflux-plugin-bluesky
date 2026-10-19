// Command nanoflux-plugin-bluesky serves the Bluesky plugin over gRPC so
// nanoflux can load it as an external plugin.
//
// Build it and drop the binary into the host's plugins directory
// (NF_PLUGINS_DIR, ./plugins by default):
//
//	GOOS=linux CGO_ENABLED=0 go build -o ../nanoflux-plugin-bluesky .
package main

import (
	goplugin "github.com/hashicorp/go-plugin"

	bluesky "nanoflux-plugin-bluesky/bluesky"

	"github.com/metruzanca/nanoflux/pluginapi"
)

func main() {
	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: pluginapi.Handshake,
		Plugins:         pluginapi.PluginSet(bluesky.Plugin{}),
		GRPCServer:      goplugin.DefaultGRPCServer,
	})
}
