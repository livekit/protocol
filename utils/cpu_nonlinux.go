//go:build !linux

package utils

func newPlatformCPUMonitor() (platformCPUMonitor, error) {
	return newOsstatCPUMonitor()
}
