package platform_external

import (
	"context"
	"fmt"
	"time"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"

	configv1 "github.com/openshift/api/config/v1"
	exutil "github.com/openshift/origin/test/extended/util"
	e2eskipper "k8s.io/kubernetes/test/e2e/framework/skipper"
	admissionapi "k8s.io/pod-security-admission/api"
)

func skipIfNotPlatformExternal(ctx context.Context, oc *exutil.CLI) {
	infra, err := oc.AdminConfigClient().ConfigV1().Infrastructures().Get(context.Background(), "cluster", metav1.GetOptions{})
	o.Expect(err).NotTo(o.HaveOccurred())

	if infra.Status.PlatformStatus.Type != configv1.ExternalPlatformType {
		e2eskipper.Skipf("this test is currently broken on platforms different than external")
	}
}

var _ = g.Describe("[sig-arch][Feature:PlatformExternal] Infrastructure", func() {
	defer g.GinkgoRecover()
	oc := exutil.NewCLIWithPodSecurityLevel("externalinfra", admissionapi.LevelRestricted)

	const (
		// should be higher than the pod start time, so first pod is still not ready when a seconds one starts with a big delay (might take a long time to pull image)
		podBecomesReadyTimeout = 5*time.Minute + 10*time.Second
	)

	g.It(fmt.Sprintf("Platform type must be external"), func() {
		ctx := context.TODO()
		skipIfNotPlatformExternal(ctx, oc)
	})

	// TODOs:
	// ProviderName must not be empty

	// ProviderName must be valid

	// > When CCM is enabled
	// providerId must be set on nodes when CCM is enabled
	// providerId must not be set when CCM is not enabled

	// nodes must be initialized when CCM is enabled

	// CCM pods must be running

	// CCM deployment must have required fields for OpenShift deployment??
	// Must be able to deploy a service load balancer
})
