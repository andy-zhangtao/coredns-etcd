FROM    vikings/alpine:base
ADD     start.sh /start.sh
ADD     coredns-etcd /coredns-etcd
ENTRYPOINT ["/start.sh"]