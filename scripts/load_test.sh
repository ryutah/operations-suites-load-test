#!/usr/bin/env bash

cd $(dirname $0)/..

set -e

mkdir -p .log

project_id=$(gcloud config get-value project)
declare -a services=("not-use-suites" "with-profiler" "with-trace" "with-both")

for service in ${services[@]}; do
  docker container run -i -w /root -v $(pwd)/.log:/root -it --rm peterevans/vegeta:6.8.1-vegeta12.8.3 sh -c \
    "echo 'GET https://${service}-dot-sandbox-hara.an.r.appspot.com/' | vegeta attack -rate=10 -duration=30s | vegeta report" > .log/${service}.txt
done
