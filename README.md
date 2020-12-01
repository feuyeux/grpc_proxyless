## gRPC Proxyless

### build & push image
```sh
sh xds-server/build_image.sh
sh hello-world/build_server_image.sh
sh hello-world/build_client_image.sh
```

### deploy to k8s
```sh
cp ~/shop/istio-1.8.0/kbc ~/.kube/config
PATH=$PATH:~/shop/istio-1.8.0/bin
istioctl install
kubectl get svc -n istio-system
```

```sh
kubectl create ns xds-proxyless
```

```sh
sh xds-server/apply_deployments.sh
sh hello-world/apply_client_deployments.sh
sh hello-world/apply_server_deployments.sh
```

### clean up
```sh
sh xds-server/delete_deployments.sh
sh hello-world/delete_client_deployments.sh
sh hello-world/delete_server_deployments.sh
```

### proto
```sh
protoc -I${PWD}/api/proto hello.proto --go_out=paths=source_relative,plugins=grpc:${PWD}/internal/app/http/rpc
```