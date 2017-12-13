# Kubernetes之kubectl常用命令
### kubectl get/logs/exec/describe/apply/create/delete 简介
- # kubectl get
get命令用于获取集群的一个或一些resource信息。使用--help查看详细信息。kubectl可以列出集群所有resource的详细。resource包括集群节点、运行的pod，service等。

基本使用方法
基本使用格式：
```
kubectl get <resources type> [rescorces name] 
[-namespace=namespace]
```
- ## kubectl get po
例如获取pod信息，可以直接使用"kubectl get po“获取当前运行的所有pods的信息
```
[root@node1 ~]# kubectl get po
NAME                                    READY     STATUS    RESTARTS   AGE
default-http-backend-3138300093-r4fj4   1/1       Running   0          7d
lb-mmsql-2237738246-906w0               1/1       Running   1          7d
lb-myapp-3785436410-vvmz9               1/1       Running   0          7d
mmsql-3854113765-c9b6r                  1/1       Running   0          7d
myapp1-279444632-k85rt                  1/1       Running   0          2h
```
- ## kubectl get po -o wide
获取pod运行在哪个节点上的信息。
```
[root@node1 ~]# kubectl get po -o wide
NAME                                    READY     STATUS    RESTARTS   AGE       IP              NODE
default-http-backend-3138300093-r4fj4   1/1       Running   0          8d        10.233.75.20    node2
lb-mmsql-2237738246-906w0               1/1       Running   1          7d        192.168.34.20   node4
lb-myapp-3785436410-vvmz9               1/1       Running   0          8d        192.168.34.13   node2
mmsql-3854113765-c9b6r                  1/1       Running   0          7d        10.233.75.22    node2
myapp1-279444632-k85rt                  1/1       Running   0          4h        10.233.74.70    node4

```
- ## kubectl get namespace
获取namespace信息。
```
[root@node1 ~]# kubectl get namespace
NAME          STATUS    AGE
default       Active    8d
ekos-plugin   Active    8d
kube-public   Active    8d
kube-system   Active    8d

```

- #### 可以使用 kubectl get rc,kubectl get svc, kubectl get nodes等获取其他resource信息。
- #### 获取一些更具体的信息，可以通过使用选项“-o”
1.kubectl get po <podname> -o yaml 以yawl格式输出pod的详细信息。
2.kubectl get po <podname> -o json 以jison格式输出pod的详细信息。
3.另外还可以使用”-o=custom-columns=“定义直接获取指定内容的值。如前面使用json和ymal格式的输出中，metadata.labels.app的值可以使用如下命令获取。
- ## kubectl logs
logs命令用于显示pod运行中，容器内程序输出到标准输出的内容。跟docker的logs命令类似。如果要获得tail -f 的方式，也可以使用-f选项。
```
[root@node1 ~]# kubectl logs myapp1-279444632-k85rt
2017/12/12 06:47:05 [I] [asm_amd64.s:2197] http server Running on http://:8080
2017/12/12 06:49:04 [D] [server.go:2568] |    10.233.75.0| 200 |   8.679292ms|   match| GET      /     r:/
2017/12/12 06:49:14 [D] [server.go:2568] |    10.233.75.0| 200 |   1.303161ms|   match| GET      /user/profile   r:/user/profile
2017/12/12 06:49:14 [D] [server.go:2568] |    10.233.75.0| 404 |    422.765µs| nomatch| GET      /favicon.ico
2017/12/12 06:49:25 [D] [server.go:2568] |    10.233.75.0| 200 |    485.161µs|   match| GET      /user/signup   r:/user/signup
E1212 06:49:30.665715       1 default.go:164] SigB db.Query error: <nil>
E1212 06:49:30.683190       1 default.go:171] Sig db.Prepare error: <nil>
E1212 06:49:30.855906       1 default.go:174] Sig stmt.Exec error: <nil>
2017/12/12 06:49:30 [D] [server.go:2568] |    10.233.75.0| 302 | 477.963739ms|   match| POST     /user/signup   r:/user/signup
2017/12/12 06:49:30 [D] [server.go:2568] |    10.233.75.0| 200 |    540.171µs|   match| GET      /user/login   r:/user/login
2017/12/12 06:49:31 [D] [server.go:2568] |    10.233.75.0| 200 |    110.399µs|   match| POST     /user/login   r:/user/login
2017/12/12 06:49:32 [D] [server.go:2568] |    10.233.75.0| 200 |    766.764µs|   match| GET      /user/login   r:/user/login
2017/12/12 06:49:35 [D] [server.go:2568] |    10.233.75.0| 302 |  63.478296ms|   match| POST     /user/login   r:/user/login
E1212 06:49:35.839788       1 default.go:86] LogB db.Query error: <nil>
E1212 06:49:35.839867       1 default.go:92] LogB rows.Scan error: <nil>
2017/12/12 06:49:35 [D] [server.go:2568] |    10.233.75.0| 200 |    882.955µs|   match| GET      /user/profile   r:/user/profile

```
- ## kubectl exec
为在一个已经运行的容器中执行一条shell命令，如果一个pod容器中，有多个容器，需要使用-c选项指定容器。
```
[root@node1 ~]# kubectl exec myapp1-279444632-k85rt hostname
myapp1-279444632-k85rt
```
- ## kubetcl describe
用于获取resource的相关信息，类似于get。但是get会获得更详细的信息，describe获得的是resource集群相关的信息。而且describe不支持-o选项。可用于查询状态。若想要查询更详细的信则使用get，且若某个pod并不是在runnig状态，这时查询就应该用describe。
```
[root@node1 ~]# kubectl describe po myapp1-279444632-k85rt
Name:		myapp1-279444632-k85rt
Namespace:	default
Node:		node4/192.168.34.20
Start Time:	Tue, 12 Dec 2017 14:46:55 +0800
Labels:		ekos-application=myapp
		ekos-service=myapp1
		log-ignore=false
		pod-template-hash=279444632
		type=service
		version=8
Annotations:	kubernetes.io/created-by={"kind":"SerializedReference","apiVersion":"v1","reference":{"kind":"ReplicaSet","namespace":"default","name":"myapp1-279444632","uid":"3b17f55e-df08-11e7-bf62-001a4a18011e","...
Status:		Running
IP:		10.233.74.70
Controllers:	ReplicaSet/myapp1-279444632
Containers:
  myapp3:
    Container ID:	docker://1d3b49492f359d3562a2527f6882cd0404f17c6e571dee6cc7d5588c71bb965f
    Image:		registry.ekos.local/default/myapp:6
    Image ID:		docker-pullable://registry.ekos.local/default/myapp@sha256:397db60a8f3d06614cf3e6f8846482d432007f7f0e64e4bfa1e8a8768fa32af3
    Port:		
    State:		Running
      Started:		Tue, 12 Dec 2017 14:47:05 +0800
    Ready:		True
    Restart Count:	0
    Limits:
      cpu:	125m
      memory:	64M
    Requests:
      cpu:		62m
      memory:		32M
    Environment:	<none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-1fpnm (ro)
Conditions:
  Type		Status
  Initialized 	True 
  Ready 	True 
  PodScheduled 	True 
Volumes:
  default-token-1fpnm:
    Type:	Secret (a volume populated by a Secret)
    SecretName:	default-token-1fpnm
    Optional:	false
QoS Class:	Burstable
Node-Selectors:	<none>
Tolerations:	<none>
Events:		<none>

```
- ## kubectl apply
利用apply对资源进行配置。只更新改动过的文件，apply命令的使用方式同replace相同，不同的是，apply不会删除原有resource，然后创建新的。apply直接在原有resource的基础上进行更新。同时kubectl apply还会resource中添加一条注释，标记当前的apply。
```
$ kubectl apply -f myapp-deploy.yaml 
deployment "myapp-server" configured
```
- ## kubectl create 
通过文件或控制台输入创建资源。如果已经定义了相应resource的yaml或json文件，直接kubectl create -f filename即可创建文件内定义的resource。
类似用法如下：
```
kubectl create -f <yaml或json格式文件>
``` 
- ## kubectl delete
根据resource名或label删除resource。
```
kubectl delete -f <yaml或json格式文件>
kubectl delete <type> -l <label>
kubectl delete <type> <name> [name]
```