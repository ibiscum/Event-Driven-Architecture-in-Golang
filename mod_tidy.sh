#!/bin/bash

# This will find every directory with a go.mod file and run go mod tidy on it.

root=`pwd`

for d in `find . -type d -print`; do
	cd "$root/$d";

	found=false;
	for file in `find . -maxdepth 1 -type f -name 'go.mod' -print`; do
		found=true;
		break;
	done;

        if $found
        then
                echo "Has go.mod file: $d";
		go mod tidy
        fi
done;
