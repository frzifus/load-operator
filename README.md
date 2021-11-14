# Load Operator

<img src="assets/k8s_cold.png" width="200" height="200" align="right" />

The intention of this project is to escape from the bad weather during
the winter season. ;)

Instead of getting wet and cold, it is used to explore the kubernetes
api and develop a deeper understanding of how to interact with it.

## Overview

This repository consists of three basic components:
1. [load-operator](kubebuilder/README.md) based on [Kubebuilder](https://github.com/kubernetes-sigs/kubebuilder) scaffolding
2. [load-operator](operator-sdk/README.md) based on [Operator-sdk](https://github.com/operator-framework/operator-sdk) scaffolding
3. [worker](worker/README.md) app to simulate some load and memory usage


NOTE: Both operator share the same properties in their CRD.
```yaml
properties:
  Name:
    type: string
  TargetLoad:
    type: integer
  TargetMemory:
    type: integer
```

### PRO TIP
Using this project in combination with distributed nodes in your
apartment will dramatically reduce your heating costs!
(depending on your node performance)
