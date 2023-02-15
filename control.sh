#!/bin/bash

WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE

mkdir -p var

module=agent
app=spiderQ
conf=cfg.json
pidfile=var/app.pid
logfile=var/app.log

function check_pid() {
    if [ -f $pidfile ];then
        pid=`cat $pidfile`
        if [ -n $pid ]; then
            running=`ps -p $pi