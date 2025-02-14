package installer

import (
	"fmt"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	e2e "k8s.io/kubernetes/test/e2e/framework"
)

// test if installer created Public IPs allocating from a custom Public IPv4 Pool when config is set.
var _ = g.Describe("[sig-installer][Suite:openshift/installer] IPs are allocated from the Public IPv4 Pool [provider:AWS]", func() {
	defer g.GinkgoRecover()

	var cfg *rest.Config
	var c *kubernetes.Clientset
	var err error
	var icfg *installedClusterConfig

	g.BeforeEach(func() {
		cfg, err = e2e.LoadConfig()
		o.Expect(err).NotTo(o.HaveOccurred())
		c, err = e2e.LoadClientset()
		o.Expect(err).NotTo(o.HaveOccurred())
		icfg, err = newInstalledClusterConfig()
		o.Expect(err).NotTo(o.HaveOccurred())
	})

	fmt.Printf("No-op: ", cfg, c, icfg)

	// TODO:
	// skip when:
	// - feature is not enabled in the install-config
	// - must check CI is not setting the config when lease isn't available
	// - private installations (should not happen as validation won't pass the config/pre-install)
	// test for each EIP allocated by the cluster:
	// - matches with Pool ID set in the config (can't use amazon provided pool)
	g.Skip("TODO: test not implemented")
})
