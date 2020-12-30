# KNative docker images v0.13.0

## install
```shell script
git clone https://github.com/kiraqjx/tool.git
git checkout master
cd knative/v0.13.0

# install serving
kubectl apply --filename https://github.com/knative/serving/releases/download/v0.13.0/serving-crds.yaml
kubectl apply --filename serving-core/serving-core.yaml

# install istio net
kubectl apply --filename net-istio/serving-istio.yaml
```
