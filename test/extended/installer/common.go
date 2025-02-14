package installer

import (
	"context"
	"strings"

	exutil "github.com/openshift/origin/test/extended/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type nodeConfig struct {
	name         string
	instanceName string
	role         string
	zone         string
}

type installedClusterConfig struct {
	infraID       string
	cloudProvider string
	nodes         map[string]nodeConfig
	installConfig *string

	// flags for variants
	// isByoVPC is True when network infrastructure is provided by user, such as platform.aws.subnets.
	isByoVPC bool
	// hasEdge is True when there is machinePool config for edge or edge subets (BYO VPC)
	hasEdge bool
}

var cluster *installedClusterConfig

func newInstalledClusterConfig() (*installedClusterConfig, error) {
	cluster = &installedClusterConfig{}

	// TODO
	// check if required variables are set
	// extracts the installConfig from cluster

	// extracts the infra object from cluster

	// evaluate flags

	return cluster, nil
}

// getNodesByRole returns a map of nodes by their role.
func getNodesByRole(oc *exutil.CLI) (map[string]nodeConfig, error) {

	//kubeNodes, err := oc.AsAdmin().KubeClient().CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{LabelSelector: "node-role.kubernetes.io/control-plane="})
	kubeNodes, err := oc.AsAdmin().KubeClient().CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	nodes := make(map[string]nodeConfig)

	for _, node := range kubeNodes.Items {
		role := ""
		for _, label := range node.Labels {
			// discover the node-role label value
			if strings.HasPrefix(label, "node-role.kubernetes.io/") {
				role = strings.Split(label, "node-role.kubernetes.io/")[1]
			}
			// TODO(mtulio): discover the instance name
		}
		nodes[node.Name] = nodeConfig{
			name:         node.Name,
			instanceName: "TBD",
			role:         role,
			zone:         "TBD",
		}
	}

	return nodes, nil
}
