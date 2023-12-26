package testsuites

import (
	"strings"
)

func inExternalConformanceSuite(name string) bool {
	if strings.Contains(name, "[Suite:k8s]") &&
		strings.Contains(name, "[Conformance]") &&
		strings.Contains(name, "[Feature:CCM]") {
		return true
	}
	if strings.Contains(name, "[Suite:openshift/conformance") &&
		strings.Contains(name, "[Feature:PlatformExternal]") {
		return true
	}

	return false
}

func inExternalSuite(name string) bool {
	if strings.Contains(name, "[Suite:k8s]") &&
		strings.Contains(name, "[Feature:CCM]") {
		return true
	}
	if strings.Contains(name, "[Feature:PlatformExternal]") {
		return true
	}

	return false
}
