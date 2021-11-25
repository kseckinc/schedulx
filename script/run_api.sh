#!/bin/bash
RUN_NAME="gf.ops.schedulx"
ps -ef|grep $RUN_NAME | grep "bin" |grep -v grep
if [ $? -ne 0 ]
then
  echo "start....."
  CURDIR=$(cd $(dirname $0); pwd)
  exec $CURDIR/bin/$RUN_NAME
else
  echo "running....."
fi