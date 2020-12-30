# KNative docker images v0.19.0

## install
```shell script
git clone https://github.com/kiraqjx/tool.git
git checkout master
cd knative/v0.19.0

# install sering
kubectl apply --filename https://github.com/knative/serving/releases/download/v0.19.0/serving-crds.yaml
kubectl apply --filename serving-core/serving-core.yaml

# install istio net
kubectl apply --filename net-istio/release.yaml

# install eventing
kubectl apply --filename https://github.com/knative/eventing/releases/download/v0.19.0/eventing-crds.yaml
kubectl apply --filename event-core/eventing-core.yaml

# install channel (in-memory)
kubectl apply --filename channel/in-memory-channel.yaml

#install broker
kubectl apply --filename mtbroker/mt-channel-broker.yaml
```
