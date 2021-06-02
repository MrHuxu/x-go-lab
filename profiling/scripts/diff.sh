#!/bin/bash

# $1: filename 1
# $2: filename 2

./FlameGraph/difffolded.pl $1.folded $2.folded | ./FlameGraph/flamegraph.pl > diff_$1_$2.svg