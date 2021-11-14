# Load Operator (kubebuilder)

## Overview
Operator initialized using kubebuilder v3.2.0.
```bash
# init
kubebuilder init --domain klimlive.de --repo load-operator
# create api
kubebuilder create api --group work --version v1 --kind WorkDeployment
```

## Usage

### Installation of CRDs into the k8s cluster specified in kube_config.
Install CRDs:
```bash
make install
```

### Start Operator
Run a controller from your host.
```
make run
```

### Observation
In order to see how workers are started and stopped after creating/deleting a deployment, the following command can be started in a second term session.
```bash
kubectl get pods -n default -w
```

### Samples
Make usage of the given example specs:
```bash
# create worker deployment
kubectl apply -f config/samples/work_v1_workdeployment.yaml
# expect:
# workdeployment.work.klimlive.de/workdeployment-sample created

# delete worker deployment
kubectl delete -f config/samples/work_v1_workdeployment.yaml
# expect:
# workdeployment.work.klimlive.de "workdeployment-sample" deleted
```

## References
- [Book Kubebuilder](https://book.kubebuilder.io/)
