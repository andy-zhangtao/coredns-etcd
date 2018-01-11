FROM 	coredns/coredns
LABEL 	maintainer=ztao8607@gmail.com
RUN 	apk --update add libc6-compat ca-certificates wget openssl tzdata && \
    	update-ca-certificates
# Set environment variables.
ENV \
		ETCD_VERSION=3.2.12

# Install etcdctl from repository.
RUN \
		cd /tmp && \
		wget --no-check-certificate https://github.com/coreos/etcd/releases/download/v${ETCD_VERSION}/etcd-v${ETCD_VERSION}-linux-amd64.tar.gz && \
		tar zxvf etcd-*-linux-amd64.tar.gz && \
	    cp etcd-*-linux-amd64/etcdctl /usr/local/bin/etcdctl && \
		rm -rf etcd-*-linux-amd64 && \
		chmod +x /usr/local/bin/etcdctl
ENV		ETCDCTL_API=2
RUN 	cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    	echo "Asia/Shanghai" > /etc/timezone
ADD 	start.sh /start.sh
ENTRYPOINT ["/start.sh"]
ADD     coredns-etcd /coredns-etcd
EXPOSE 	53/udp 53/tcp