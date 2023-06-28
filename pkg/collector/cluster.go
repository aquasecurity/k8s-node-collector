package collector

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type Cluster struct {
	clientSet     *kubernetes.Clientset
	cConfig       clientcmd.ClientConfig
	restMapper    meta.RESTMapper
	dynamicClient dynamic.Interface
}

type Platform struct {
	Name    string
	Version string
}

func NewCluster(clientSet *kubernetes.Clientset, clientConfig clientcmd.ClientConfig, restMApper meta.RESTMapper, dynamicClient dynamic.Interface) *Cluster {
	return &Cluster{clientSet: clientSet, cConfig: clientConfig, restMapper: restMApper, dynamicClient: dynamicClient}
}

func GetCluster() (*Cluster, error) {
	cf := genericclioptions.NewConfigFlags(true)
	rest.SetDefaultWarningHandler(rest.NoWarnings{})
	clientConfig := cf.ToRawKubeConfigLoader()
	rc, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	restMapper, err := cf.ToRESTMapper()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(rc)
	if err != nil {
		return nil, err
	}
	k8sDynamicClient, err := dynamic.NewForConfig(rc)
	if err != nil {
		return nil, err
	}
	return NewCluster(clientset, clientConfig, restMapper, k8sDynamicClient), nil
}

func (cluster *Cluster) Platfrom() (Platform, error) {
	v := cluster.getOpenShiftVersion(context.Background())
	if len(v) != 0 {
		return Platform{Name: "ocp", Version: majorVersion(v)}, nil
	}
	version, err := cluster.clientSet.ServerVersion()
	if err != nil {
		return Platform{}, err
	}
	return getPlatformInfoFromVersion(version.GitVersion), nil
}

func getPlatformInfoFromVersion(s string) Platform {
	versionRe := regexp.MustCompile(`v(\d+\.\d+)\.\d+[-+](\w+)(?:[.\-])\w+`)
	subs := versionRe.FindStringSubmatch(s)
	if len(subs) < 3 {
		return Platform{
			Name:    "k8s",
			Version: majorVersion(s),
		}
	}
	return Platform{
		Name:    subs[2],
		Version: subs[1],
	}
}

func (cluster *Cluster) getOpenShiftVersion(ctx context.Context) string {
	gvr, err := cluster.restMapper.ResourceFor(schema.GroupVersionResource{Resource: "clusterversions"})
	if err != nil {
		return ""
	}
	dclient := cluster.getDynamicClient(gvr)
	resources, err := dclient.List(ctx, v1.ListOptions{})
	if err != nil {
		return ""
	}
	var version string
	for _, resource := range resources.Items {
		version, _, _ = unstructured.NestedString(resource.Object, []string{"status", "desired", "version"}...)

	}
	return version
}

func (cluster *Cluster) getDynamicClient(gvr schema.GroupVersionResource) dynamic.ResourceInterface {
	return cluster.dynamicClient.Resource(gvr).Namespace("")
}

func majorVersion(semanticVersion string) string {
	versionRe := regexp.MustCompile(`v(\d+\.\d+)\.\d+`)
	version := semanticVersion
	if !strings.HasPrefix(semanticVersion, "v") {
		version = fmt.Sprintf("v%s", semanticVersion)
	}
	subs := versionRe.FindStringSubmatch(version)
	return subs[1]
}
