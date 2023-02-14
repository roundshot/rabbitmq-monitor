#!/bin/bash

WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE

mkdir -p var

module=agent
app=spiderQ
conf=cfg