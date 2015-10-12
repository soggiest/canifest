#!/bin/bash

#
# Script used to control the build
# 

BIN=df_gen
DESTDIR=bin

which go &>/dev/null
if [ $? -ne 0 -o -z "$GOROOT" ]
then
  echo "Failed to find GOROOT or the go executable - make sure it is set"
  exit 1
fi 

export GOARCH=386
export GOPATH=`pwd`

mkdir -p $DESTDIR

# Ensure formatting is followed 
go fmt df_generator

echo "Building..."
go build -o $DESTDIR/linux/${BIN} df_generator

