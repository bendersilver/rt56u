#!/usr/bin/env bash
svr="192.168.1.1"

path="/tmp/m3u/static"

yarn build

scp -Cr dist/* $package $svr:$path