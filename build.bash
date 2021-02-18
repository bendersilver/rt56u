#!/usr/bin/env bash

package="go-ipTv"
svr="192.168.1.1"

path="/tmp/m3u"

GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags "-s -w" -o $package

scp -C $package $svr:$path
rm $package