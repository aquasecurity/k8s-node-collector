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
    "adminConfFileOwnership": {
      "values": [
        "root:root"
      ]
    },
    "adminConfFilePermissions": {
      "values": [
        600
      ]
    },
    "certificateAuthoritiesFileOwnership": {
      "values": [
        "root:root"
      ]
    },
    "certificateAuthoritiesFilePermissions": {
      "values": [
        644
      ]
    },
    "containerNetworkInterfaceFileOwnership": {
      "values": [
        "root:root",
        "root:root"
      ]
    },
    "containerNetworkInterfaceFilePermissions": {
      "values": [
        700,
        775
      ]
    },
    "controllerManagerConfFileOwnership": {
      "values": [
        "root:root"
      ]
    },
    "controllerManagerConfFilePermissions": {
      "values": [
        600
      ]
    },
    "etcdDataDirectoryOwnership": {
      "values": [
        "root:root"
      ]
    },
    "etcdDataDirectoryPermissions": {
      "values": [
        700
      ]
    },
    "kubeAPIServerSpecFileOwnership": {
      "values": [
        "root:root"
      ]
    },
    "kubeAPIServerSpecFilePermission": {
      "values": [
        600
      ]
    },
    "kubeControllerManagerSpecFileOwnership": {
      "values": [
        "root:root"
      ]
    },
    "kubeControllerManagerSpecFilePermission": {
      "values": [
        600
      ]
    },
    "kubeEtcdSpecFileOwnership": {
      "values": [
        "root:root"
      ]
    },
    "kubeEtcdSpecFilePermission": {
      "values": [
        600
      ]
    },
    "kubePKIDirectoryFileOwnership": {
      "values": [
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
      ]
    },
    "kubePKIKeyFilePermissions": {
      "values": [
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
      ]
    },
    "kubeSchedulerSpecFileOwnership": {
      "values": [
        "root:root"
      ]
    },
    "kubeSchedulerSpecFilePermission": {
      "values": [
        600
      ]
    },
    "kubeconfigFileExistsOwnership": {
      "values": [

      ]
    },
    "kubeconfigFileExistsPermissions": {
      "values": [

      ]
    },
    "kubeletAnonymousAuthArgumentSet": {
      "values": [

      ]
    },
    "kubeletAuthorizationModeArgumentSet": {
      "values": [
        "Node",
        "RBAC"
      ]
    },
    "kubeletClientCaFileArgumentSet": {
      "values": [
        "/etc/kubernetes/pki/ca.crt"
      ]
    },
    "kubeletConfFileOwnership": {
      "values": [
        "root:root"
      ]
    },
    "kubeletConfFilePermissions": {
      "values": [
        600
      ]
    },
    "kubeletConfigYamlConfigurationFileOwnership": {
      "values": [
        "root:root"
      ]
    },
    "kubeletConfigYamlConfigurationFilePermission": {
      "values": [
        644
      ]
    },
    "kubeletEventQpsArgumentSet": {
      "values": [

      ]
    },
    "kubeletHostnameOverrideArgumentSet": {
      "values": [

      ]
    },
    "kubeletMakeIptablesUtilChainsArgumentSet": {
      "values": [

      ]
    },
    "kubeletOnlyUseStrongCryptographic": {
      "values": [

      ]
    },
    "kubeletProtectKernelDefaultsArgumentSet": {
      "values": [

      ]
    },
    "kubeletReadOnlyPortArgumentSet": {
      "values": [

      ]
    },
    "kubeletRotateCertificatesArgumentSet": {
      "values": [

      ]
    },
    "kubeletRotateKubeletServerCertificateArgumentSet": {
      "values": [

      ]
    },
    "kubeletServiceFileOwnership": {
      "values": [
        "root:root"
      ]
    },
    "kubeletServiceFilePermissions": {
      "values": [
        644
      ]
    },
    "kubeletStreamingConnectionIdleTimeoutArgumentSet": {
      "values": [

      ]
    },
    "kubeletTlsCertFileTlsArgumentSet": {
      "values": [
        "/etc/kubernetes/pki/apiserver.crt"
      ]
    },
    "kubeletTlsPrivateKeyFileArgumentSet": {
      "values": [
        "/etc/kubernetes/pki/apiserver.key"
      ]
    },
    "kubernetesPKICertificateFilePermissions": {
      "values": [
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
      ]
    },
    "schedulerConfFileOwnership": {
      "values": [
        "root:root"
      ]
    },
    "schedulerConfFilePermissions": {
      "values": [
        600
      ]
    }
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
