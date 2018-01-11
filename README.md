# coredns-etcd
Enable etcd plugin for coredns. 

# How to use coredns-etcd?

First, the below env must be set. 

- DDOG_DOMAIN The Dns Domain which will be used.
- DDOG_UP_STREAM The Up Dns Server. If coredns doesn't have the record, coredns will search from this upstream server. Default the port is 53.
- DDOG_CONF_PATH The path corefile will be save. Default is '/'
- DDOG_ETCD_ENDPOINT The Etcd Cluster endpoint. e.g. xxxxx:2379

