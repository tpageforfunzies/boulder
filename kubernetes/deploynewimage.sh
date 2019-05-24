#!/bin/bash
# call it with bash ./kubernetes/deploynewimage.sh X.XX to build, push and update the deployment
if [ -z "$1" ]; then
    echo "gotta have a version number dog"
    exit 2
fi
echo "======building image======"
docker build -t tpageforfunzies/boulderarmhf:v"$1" .
echo "======pushing image======"
docker push tpageforfunzies/boulderarmhf:v"$1"
echo "======updating deployment to rollout new image======"
kubectl --record deployment.apps/bouldertracker set image deployment.v1.apps/bouldertracker bouldertracker-api=tpageforfunzies/boulderarmhf:v"$1"