#!/usr/bin/env sh
/coredns-etcd
if [ $? -ne 0 ]; then
exit 1
fi
/coredns  $*