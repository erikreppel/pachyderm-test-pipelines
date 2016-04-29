### Usage:

1. [Start Pachyderm](https://github.com/pachyderm/pachyderm/blob/master/SETUP.md)
2. `$ kubectl get all` and note the IP that pachd is running on

    ```
$ kubectl get all
NAME                   DESIRED      CURRENT       AGE
etcd                   1            1             1h
pachd                  2            2             1h
rethink                1            1             1h
NAME                   CLUSTER-IP   EXTERNAL-IP   PORT(S)                        AGE
etcd                   10.0.0.133   <none>        2379/TCP,2380/TCP              1h
kubernetes             10.0.0.1     <none>        443/TCP                        8d
pachd                  10.0.0.167   nodes         650/TCP,750/TCP                1h
rethink                10.0.0.233   <none>        8080/TCP,28015/TCP,29015/TCP   1h
NAME                   READY        STATUS        RESTARTS                       AGE
etcd-aspfe             1/1          Running       0                              1h
k8s-etcd-127.0.0.1     1/1          Running       0                              8d
k8s-master-127.0.0.1   4/4          Running       0                              8d
k8s-proxy-127.0.0.1    1/1          Running       0                              8d
pachd-nan8l            1/1          Running       2                              1h
pachd-rsvsv            1/1          Running       2                              1h
rethink-kqu9r          1/1          Running       0                              1h
    ```

3. set enivronment variable
    ```
    $ export PACHD_PORT_650_TCP_ADDR=10.0.0.167
    ```

4. `$ go run create_numbers/create_numbers.go`

5. `$ pachctl create-pipeline -f pipeline.json`

6. `$ pachctl list-job` to see the jobs created by the pipeline (might take a second for them all to show up)
