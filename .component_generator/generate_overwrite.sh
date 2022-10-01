#!/bin/sh
# trunk-ignore(shellcheck/SC2016)
find configs/ -name '*.yaml' -print0 | xargs -0 -I {} /bin/bash -c 'name=$(basename ${1}); echo ${name%.*}; ./generate.py ${1} --out_dir=../${name%.*}' '_' '{}' \+
go fmt ../...
