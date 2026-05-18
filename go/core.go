// Package evcore is the gomobile-bound entry point for Everywhere's
// networking stack. It boots one of three upstream proxy cores —
// Xray, sing-box, or mihomo — that all share a single Go runtime
// when this module is bound as one xcframework.
//
// Each core owns its own TUN inbound, fed the iOS utun file
// descriptor obtained from NEPacketTunnelProvider. There is no
// separate userland tun→socks shim.
//
// The Swift side calls (in order):
//
//	SetResourcesPath(path)                   // optional, asset dir
//	StartCore(coreType, configContent,       // boots the proxy core
//	          tunFD, mtu)                    //   with TUN attached
//
// While running:
//
//	Suspend()                                // on NE sleep()
//	Resume()                                 // on NE wake()
//	UpdateDefaultInterface(name, index,      // on NWPathMonitor
//	                       expensive, ...)   //   update
//
// On teardown:
//
//	StopAll()
//
// The provided configuration must declare a TUN inbound for the
// active core; ConfigNormalizer on the Swift side handles that.
package evcore

import (
	"errors"
	"fmt"
	"sync"
)

const (
	CoreTypeXray    = "xray"
	CoreTypeSingBox = "singbox"
	CoreTypeMihomo  = "mihomo"
)

var (
	mu           sync.Mutex
	coreInstance coreRunner
)

type coreRunner interface {
	stop() error
	// suspend pauses non-essential activity in the running core
	// (URL-test probes, keepalives, new-connection handling) until
	// resume is called. Cores with no native pause hook (xray) leave
	// this as a no-op.
	suspend()
	resume()
	// updateDefaultInterface tells the core that the device's default
	// network interface has changed. name == "" / index == -1 means
	// the device currently has no usable path.
	updateDefaultInterface(name string, index int32, isExpensive, isConstrained bool)
}

func Version() string { return "Everywhere Core v0.2" }

// StartCore boots the chosen proxy core with TUN attached to the
// given iOS utun file descriptor. The FD lifetime stays with the
// caller — cores that need to own a copy dup it internally.
func StartCore(coreType, configContent string, tunFD, mtu int) error {
	mu.Lock()
	defer mu.Unlock()
	if coreInstance != nil {
		return errors.New("a core is already running")
	}
	if tunFD < 0 {
		return errors.New("invalid tun file descriptor")
	}
	if mtu <= 0 {
		mtu = 1500
	}
	var (
		r   coreRunner
		err error
	)
	switch coreType {
	case CoreTypeXray:
		r, err = startXray(configContent, tunFD, mtu)
	case CoreTypeSingBox:
		r, err = startSingBox(configContent, tunFD, mtu)
	case CoreTypeMihomo:
		r, err = startMihomo(configContent, tunFD, mtu)
	default:
		return fmt.Errorf("unknown core type: %s", coreType)
	}
	if err != nil {
		// gomobile's seq.ToRefNum boxes returned errors via a map keyed
		// by the value itself, which panics on non-comparable types
		// (e.g. mihomo's constant.ErrNotSafePath has a []string field).
		// Flatten to a plain *errorString so Swift gets the message
		// instead of the NE process crashing.
		return errors.New(err.Error())
	}
	coreInstance = r
	return nil
}

// Suspend pauses non-essential activity of the running core. Call it
// from NEPacketTunnelProvider.sleep() — without that hop, the Go cores
// keep firing scheduled work (sing-box URL-test probes, mihomo new-
// connection handling, wireguard keepalives) at full pace through iOS
// device-sleep windows. Returns nil when no core is running.
func Suspend() error {
	mu.Lock()
	defer mu.Unlock()
	if coreInstance == nil {
		return nil
	}
	coreInstance.suspend()
	return nil
}

// Resume reverses Suspend. Call it from NEPacketTunnelProvider.wake().
func Resume() error {
	mu.Lock()
	defer mu.Unlock()
	if coreInstance == nil {
		return nil
	}
	coreInstance.resume()
	return nil
}

// UpdateDefaultInterface tells the running core the device's default
// physical interface changed — drive it from an NWPathMonitor on the
// Swift side. Pass `index = -1` and `name = ""` when there is no
// usable path. Without these updates, outbound sockets pinned by Go
// to a stale path keep retransmitting after a WiFi↔cellular handoff.
func UpdateDefaultInterface(name string, index int32, isExpensive, isConstrained bool) error {
	mu.Lock()
	defer mu.Unlock()
	if coreInstance == nil {
		return nil
	}
	coreInstance.updateDefaultInterface(name, index, isExpensive, isConstrained)
	return nil
}

// StopAll halts the running core. Teardown is detached: the upstream
// libraries' close paths can each take seconds (Xray drains
// outbounds, sing-box has a 10s/service timeout, mihomo cleans up
// DNS/listeners), and we don't want the Network Extension to block
// on that — iOS terminates the NE process shortly after stopTunnel
// returns, which reclaims everything anyway. Errors from the
// detached stop are intentionally dropped.
func StopAll() error {
	mu.Lock()
	prev := coreInstance
	coreInstance = nil
	mu.Unlock()

	go func() {
		defer func() { _ = recover() }()
		if prev != nil {
			_ = prev.stop()
		}
	}()
	return nil
}
