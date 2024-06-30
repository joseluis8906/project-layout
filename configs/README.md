# Configurations
Install etcdctl part of the etcd package and yq. (both packages are hosted in homebrew)
Save configurations into etcd:
```bash
./put.sh <file.yml>
```

Retrieve configurations from etcd:
```bash
./get.sh <file.yml>
```

Delete configurations from etcd:
```bash
./del.sh <file.yml>
```
