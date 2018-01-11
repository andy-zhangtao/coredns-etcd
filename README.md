# coredns-etcd
Enable etcd plugin for coredns.

# How to use coredns-etcd?

The below env must be set.

- DDOG_DOMAIN The Dns Domain which will be used.
- DDOG_UP_STREAM The Up Dns Server. If coredns doesn't have the record, coredns will search from this upstream server. Default the port is 53.
- DDOG_CONF_PATH The path corefile will be save. Default is '/'
- DDOG_ETCD_ENDPOINT The Etcd Cluster endpoint. e.g. xxxxx:2379

There are more info for you reference.

- DDOG_DOMAIN: 需要DNS解析的域名. 例如: mydomain.com
- DDOG_ETCD_ENDPOINT: Etcd地址,CoreDNS用来持久化数据. 例如: etcd.com:2379
- DDOG_UP_STREAM: 上游DNS地址, 支持三种配置方式:
   * 配置为/etc/resolv.conf
   * 配置为IP,例如: 10.0.0.1;10.0.0.2
   * 配置为IP+Port,例如: 10.0.0.1:54;10.0.0.2

以下变量若为空会使用默认值
  - DDOG_CONF_PATH: Corefile路径,默认为 /
