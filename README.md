[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/aquasecurity/k8s-node-collector/blob/main/LICENSE)

# kube-node-collector

Kube-Node-collector is an open source collector who collect Node information (fs and process data) and output in a table/json format.

## Kube-node-collector as job in k8s

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
  "kind": "Nodeinfo",
  "type": "master",
  "info": {
    "AdminConfFileOwnership": [
      "root:root"
    ],
    "AdminConfFilePermissions": [
      600
    ],
    "CertificateAuthoritiesFileOwnership": [
      "root:root"
    ],
    "CertificateAuthoritiesFilePermissions": [
      644
    ],
    "ContainerNetworkInterfaceFileOwnership": [
      "root:root",
      "root:root"
    ],
    "ContainerNetworkInterfaceFilePermissions": [
      700,
      775
    ],
    "ControllerManagerConfFileOwnership": [
      "root:root"
    ],
    "ControllerManagerConfFilePermissions": [
      600
    ],
    "EtcdDataDirectoryOwnership": [
      "root:root"
    ],
    "EtcdDataDirectoryPermissions": [
      700
    ],
    "KubePKIDirectoryFileOwnership": [
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
    "KubePKIKeyFilePermissions": [
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
    "KubeconfigFileExistsPermissions": [
      ""
    ],
    "KubeletAnonymousAuthArgumentSet": [
      ""
    ],
    "KubeletAuthorizationModeArgumentSet": [
      "Node",
      "RBAC"
    ],
    "KubeletClientCaFileArgumentSet": [
      "/etc/kubernetes/pki/ca.crt"
    ],
    "KubeletEventQpsArgumentSet": [
      ""
    ],
    "KubeletHostnameOverrideArgumentSet": [
      ""
    ],
    "KubeletMakeIptablesUtilChainsArgumentSet": [
      ""
    ],
    "KubeletOnlyUseStrongCryptographic": [
      ""
    ],
    "KubeletProtectKernelDefaultsArgumentSet": [
      ""
    ],
    "KubeletReadOnlyPortArgumentSet": [
      ""
    ],
    "KubeletRotateCertificatesArgumentSet": [
      ""
    ],
    "KubeletRotateKubeletServerCertificateArgumentSet": [
      ""
    ],
    "KubeletStreamingConnectionIdleTimeoutArgumentSet": [
      ""
    ],
    "KubeletTlsCertFileTlsArgumentSet": [
      "/etc/kubernetes/pki/apiserver.crt"
    ],
    "KubeletTlsPrivateKeyFileArgumentSet": [
      "/etc/kubernetes/pki/apiserver.key"
    ],
    "KubernetesPKICertificateFilePermissions": [
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
    "SchedulerConfFileOwnership": [
      "root:root"
    ],
    "SchedulerConfFilePermissions": [
      600
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
    "kubeSchedulerSpecFileOwnership": [
      "root:root"
    ],
    "kubeSchedulerSpecFilePermission": [
      600
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
    "kubeletServiceFileOwnership": [
      "root:root"
    ],
    "kubeletServiceFilePermissions": [
      644
    ]
  }
}
```
* job cleanup
```
kubectl delete -f job.yaml
```

