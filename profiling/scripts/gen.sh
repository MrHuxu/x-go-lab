#!/bin/bash

go tool pprof -raw -output=./$2.perf "$1/debug/pprof/profile?seconds=$3"

./FlameGraph/stackcollapse-go.pl $2.perf > $2.folded
./FlameGraph/flamegraph.pl $2.folded > $2.svg