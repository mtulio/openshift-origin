package installer

import (
	"fmt"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	e2e "k8s.io/kubernetes/test/e2e/framework"
)

var _ = g.Describe("[sig-installer][Suite:openshift/installer] Edge", func() {
	defer g.GinkgoRecover()

	var cfg *rest.Config
	var c *kubernetes.Clientset
	var err error

	g.BeforeEach(func() {
		cfg, err = e2e.LoadConfig()
		o.Expect(err).NotTo(o.HaveOccurred())
		c, err = e2e.LoadClientset()
		o.Expect(err).NotTo(o.HaveOccurred())
	})

	fmt.Printf("No-op: ", cfg, c)

	g.It("nodes must have taints [provider:aws]", func() {
		//   TODO:
		//   - All edge nodes must have NoSchedule taints applied to the node on label node-role.kubernetes.io/edge
		g.Skip("TODO: test not implemented")
	})

})
