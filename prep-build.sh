#!/bin/bash

which go &>/dev/null
if [ $? -ne 0 -o -z "$GOROOT" ]
then
  echo "Failed to find GOROOT or the go executable - make sure it is set"
  exit 1
fi 

# Is this a valid workaround ??
# export GOROOT_BOOTSTRAP=$GOROOT

export GOARCH=386

echo "Preparing builds for Linux(386)"
export GOOS=linux
cd $GOROOT/src; ./make.bash --no-clean

echo "Preparing builds for OS X/Darwin(386)"
export GOOS=darwin
cd $GOROOT/src; ./make.bash --no-clean

echo "Preparing builds for Windows(386)"
export GOOS=windows
cd $GOROOT/src; ./make.bash --no-clean

# cd $GOROOT/src; ./all.bash
