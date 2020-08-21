#!/usr/bin/env bash

cd $(dirname $0)/..

gcloud app deploy --quiet \
  default \
  with-trace \
  with-profiler \
  with-profiler-and-trace
