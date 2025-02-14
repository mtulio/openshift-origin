package installer

import (
	"fmt"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	e2e "k8s.io/kubernetes/test/e2e/framework"
)

// test Load Balancer created by installer to serve kube API
var _ = g.Describe("[sig-installer][Suite:openshift/installer] Manage LoadBalancer", func() {
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

	g.It("must have required health check configuration", func() {
		g.It("for kube API server", func() {
			// TODO:
			// - validate if the Load Balancer created to kube API server have the required health check configuration
			// - skip when not AWS(?)
			g.Skip("TODO: test not implemented")
		})
		g.It("for Machine Config Server", func() {
			// TODO:
			// - validate if the Load Balancer created to kube API server have the required health check configuration
			// - skip when not AWS(?)
			g.Skip("TODO: test not implemented")
		})
	})
})
