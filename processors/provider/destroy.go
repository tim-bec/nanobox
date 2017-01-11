package provider

import (
	"fmt"

	"github.com/jcelliott/lumber"

	"github.com/nanobox-io/nanobox/processors/provider/bridge"
	"github.com/nanobox-io/nanobox/util/locker"
	"github.com/nanobox-io/nanobox/util/provider"
)

// Destroy destroys the provider
func Destroy() error {
	locker.GlobalLock()
	defer locker.GlobalUnlock()

	if provider.BridgeRequired() {

		// remove the network bridge
		if err := bridge.Teardown(); err != nil {
			return fmt.Errorf("failed to teardown network bridge: %s", err.Error())
		}
	}

	// destroy the provider
	if err := provider.Destroy(); err != nil {
		lumber.Error("provider:Destroy:provider.Destroy(): %s", err.Error())
		return fmt.Errorf("failed to destroy the provider: %s", err.Error())
	}

	return nil
}
