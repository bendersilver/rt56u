#!/usr/bin/env bash

if [ "$1" == "start" ]
then
    exec /sbin/start-stop-daemon -p /tmp/m3u/ipTv.pid -Sbmvx /tmp/m3u/go-ipTv
elif [ "$1" == "stop" ]
then
    exec /sbin/start-stop-daemon -Kvx /tmp/m3u/go-ipTv
else
    echo "requery parametr start or stop. "
fi