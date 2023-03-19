# k8soperator

K8s operator demo



## Source

- [k8soperator_controller.go](https://github.com/craftslab/k8soperator/blob/main/controllers/k8soperator_controller.go)
- [k8soperator_types.go](https://github.com/craftslab/k8soperator/blob/main/api/v1/k8soperator_types.go)
- [_v1_k8soperator.yaml](https://github.com/craftslab/k8soperator/blob/main/config/samples/_v1_k8soperator.yaml)



## Init

```bash
kubebuilder init \
  --project-name=k8soperator \
  --owner=craftslab \
  --repo=github.com/craftslab/k8soperator

kubebuilder create api --version v1 --kind K8sOperator
```



## Deploy

### CRD

```bash
make manifests
make install
make uninstall

kubectl get crd
````



### CR

```bash
kubectl apply -f config/samples/_v1_k8soperator.yaml

kubectl get k8soperator
```



### Controller

```bash
make docker-build IMG=craftslab/k8soperator:latest
kind load docker-image IMG=craftslab/k8soperator:latest --name dev

make deploy IMG=craftslab/k8soperator:latest
make undeploy
```



## Test

```bash
make run

kubectl get service
kubectl get deployment
kubectl get pod
```



## Reference

- [developing-kubernetes-operators-in-go](https://tonybai.com/2022/08/15/developing-kubernetes-operators-in-go-part1/)
- [k8sapp](https://github.com/craftslab/k8sapp)
- [kind-deploy](https://gist.github.com/craftslab/d53fe048813b5bd05f24118959e9c246)
