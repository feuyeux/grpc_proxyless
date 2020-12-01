#!/usr/bin/env bash
SCRIPT_PATH="$(
  cd "$(dirname "$0")" >/dev/null 2>&1
  pwd -P
)/"
cd "$SCRIPT_PATH" || exit

docker build -t registry.cn-beijing.aliyuncs.com/asm_repo/hello-grpc-server:1.0.0 -f build/docker/Dockerfile.server .
docker push registry.cn-beijing.aliyuncs.com/asm_repo/hello-grpc-server:1.0.0