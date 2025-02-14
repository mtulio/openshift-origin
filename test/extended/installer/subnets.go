package installer

import (
	"fmt"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	e2e "k8s.io/kubernetes/test/e2e/framework"
)

var _ = g.Describe("[sig-installer][Suite:openshift/installer] Subnets ", func() {
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

	// validate subnets
	g.It("resources must have the required tags", func() {
		//   TODO:
		//   - when managed VPC (must have cluster tags set to owner across all subnets in the VPC)
		//   - when unmanaged VPC (must have cluster tags set to shared across config selected subnets in the VPC)
		g.Skip("TODO: test not implemented")
	})

	// validate managed subnets
	g.It("managed", func() {
		g.It("public resources must be attached to a valid route table [provider:aws]", func() {
			//   TODO:
			//   - Availability (regular) and Local zones' subnets must be in route table with default route to Internet Gateway
			//   - Wavelength zones' subnets must be in route table with default route to Carrier Gateway (this test is a must as won't fail on install)
			g.Skip("TODO: test not implemented")
		})

		g.It("edge resources must exist in required scheme [provider:aws]", func() {
			//   TODO:
			//   - Installer must create public and private subnets when edge zone is provided (managed edge subnet)
			g.Skip("TODO: test not implemented")
		})
	})

})
