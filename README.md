[![GitHub Release][release-img]][release]
[![Build Action][action-build-img]][action-build]
[![Release snapshot Action][action-release-snapshot-img]][action-release-snapshot]
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/aquasecurity/k8s-node-collector/blob/main/LICENSE)

# k8s-node-collector

k8s-Node-collector is an open source collector who collect Node information (fs and process data) and output in a table/json format.

## k8s-node-collector as job in k8s

- simple k8s cluster run following job

```
kubectl apply -f job.yaml
```

* Check k8s pod status

```
kubectl get pods 

NAME                                     READY   STATUS      RESTARTS   AGE
node-collector-ng2z7                          0/1     Completed   0          6m13s
```

* Check k8s pod audit output

```
kubectl logs node-collector-ng2z7
```

* json output

```json
{
  "apiVersion": "v1",
  "kind": "NodeInfo",
  "type": "master",
  "info": {
    "adminConfFileOwnership": [
      "root:root"
    ],
    "adminConfFilePermissions": [
      600
    ],
    "certificateAuthoritiesFileOwnership": [
      "root:root"
    ],
    "certificateAuthoritiesFilePermissions": [
      644
    ],
    "containerNetworkInterfaceFileOwnership": [
      "root:root",
      "root:root"
    ],
    "containerNetworkInterfaceFilePermissions": [
      700,
      775
    ],
    "controllerManagerConfFileOwnership": [
      "root:root"
    ],
    "controllerManagerConfFilePermissions": [
      600
    ],
    "etcdDataDirectoryOwnership": [
      "root:root"
    ],
    "etcdDataDirectoryPermissions": [
      700
    ],
    "kubeAPIServerSpecFileOwnership": [
      "root:root"
    ],
    "kubeAPIServerSpecFilePermission": [
      600
    ],
    "kubeControllerManagerSpecFileOwnership": [
      "root:root"
    ],
    "kubeControllerManagerSpecFilePermission": [
      600
    ],
    "kubeEtcdSpecFileOwnership": [
      "root:root"
    ],
    "kubeEtcdSpecFilePermission": [
      600
    ],
    "kubePKIDirectoryFileOwnership": [
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root",
      "root:root"
    ],
    "kubePKIKeyFilePermissions": [
      600,
      600,
      600,
      600,
      600,
      600,
      600,
      600,
      600,
      600,
      600
    ],
    "kubeSchedulerSpecFileOwnership": [
      "root:root"
    ],
    "kubeSchedulerSpecFilePermission": [
      600
    ],
    "kubeconfigFileExistsOwnership": [

    ],
    "kubeconfigFileExistsPermissions": [

    ],
    "kubeletAnonymousAuthArgumentSet": [

    ],
    "kubeletAuthorizationModeArgumentSet": [
      "Node",
      "RBAC"
    ],
    "kubeletClientCaFileArgumentSet": [
      "/etc/kubernetes/pki/ca.crt"
    ],
    "kubeletConfFileOwnership": [
      "root:root"
    ],
    "kubeletConfFilePermissions": [
      600
    ],
    "kubeletConfigYamlConfigurationFileOwnership": [
      "root:root"
    ],
    "kubeletConfigYamlConfigurationFilePermission": [
      644
    ],
    "kubeletEventQpsArgumentSet": [

    ],
    "kubeletHostnameOverrideArgumentSet": [

    ],
    "kubeletMakeIptablesUtilChainsArgumentSet": [

    ],
    "kubeletOnlyUseStrongCryptographic": [

    ],
    "kubeletProtectKernelDefaultsArgumentSet": [

    ],
    "kubeletReadOnlyPortArgumentSet": [

    ],
    "kubeletRotateCertificatesArgumentSet": [

    ],
    "kubeletRotateKubeletServerCertificateArgumentSet": [

    ],
    "kubeletServiceFileOwnership": [
      "root:root"
    ],
    "kubeletServiceFilePermissions": [
      644
    ],
    "kubeletStreamingConnectionIdleTimeoutArgumentSet": [

    ],
    "kubeletTlsCertFileTlsArgumentSet": [
      "/etc/kubernetes/pki/apiserver.crt"
    ],
    "kubeletTlsPrivateKeyFileArgumentSet": [
      "/etc/kubernetes/pki/apiserver.key"
    ],
    "kubernetesPKICertificateFilePermissions": [
      644,
      644,
      644,
      644,
      644,
      644,
      644,
      644,
      644,
      644
    ],
    "schedulerConfFileOwnership": [
      "root:root"
    ],
    "schedulerConfFilePermissions": [
      600
    ]
  }
}
```
* job cleanup
```
kubectl delete -f job.yaml
```

[release-img]: https://img.shields.io/github/release/aquasecurity/k8s-node-collector.svg?logo=github
[release]: https://github.com/aquasecurity/k8s-node-collector/releases
[action-build-img]: https://github.com/aquasecurity/k8s-node-collector/actions/workflows/build.yaml/badge.svg
[action-build]: https://github.com/aquasecurity/k8s-node-collector/actions/workflows/build.yaml
[action-release-snapshot-img]: https://github.com/aquasecurity/k8s-node-collector/actions/workflows/release-snapshot.yaml/badge.svg
[action-release-snapshot]: https://github.com/aquasecurity/k8s-node-collector/actions/workflows/release-snapshot.yaml
