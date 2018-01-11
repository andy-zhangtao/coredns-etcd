#!/usr/bin/env sh
/coredns-etcd
if[ $? -ne 0];
nohup /coredns  $* >> /dns.log & tail -f /dns.log then
exit 1
fi
