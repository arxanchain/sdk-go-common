#!/bin/bash

set -e

for i in "crypto errors log rest structs utils"
do
	OUTPUT="$(goimports -l -e $i)"
	if [[ $OUTPUT ]]; then
		echo "The following files contain goimports errors"
		echo $OUTPUT
		echo "The goimports command must be run for these files"
		exit 1
	fi
done
