**go buid:**  
go mod init main.go   
go get k8s.io/client-go/kubernetes  
go get k8s.io/client-go/tools/clientcmd 
go get k8s.io/client-go/rest  
go build main.go

**Docker:**  
docker build -t k8s:0.1.0 .  
docker tag k8s:0.1.0 quay.io/oviner/k8s-go:0.1.0  
docker push quay.io/oviner/k8s-go:0.1.0   


**k8s cmd:**    
$ minikube start --disk-size=10g --extra-disks=1 --driver kvm2
$ kubectl create deployment k8s-go --image quay.io/oviner/k8s-go:0.1.0 --dry-run=client -o yaml > k8s-go.yaml  
$ kubectl create -f k8s-go.yaml  

```
$ kubectl create role poddepl --resource pod,deployments --verb list
```
**Create a Role (poddepl):**    

kubectl create role poddepl: 
This part of the command instructs Kubernetes to create a Role resource named poddepl.

--resource pods,deployments: 
This specifies the resources to which the Role will grant permissions. 
In this case, it grants permissions to pods and deployments.

--verb list: 
This specifies the actions that are allowed on the specified resources. 
The list verb allows the service account associated with this Role to list (view) pods and deployments.


```
$ kubectl create rolebinding poddepl --role poddepl --serviceaccount default:default
```

**Create a RoleBinding (poddepl):**   

kubectl create rolebinding poddepl:   
This part of the command instructs Kubernetes to create a RoleBinding resource named poddepl.

--role poddepl: 
This specifies the Role (created in the previous command) to bind to this RoleBinding. 
In other words, it associates the permissions defined in the poddepl Role with this RoleBinding.

--serviceaccount default:default: 
This specifies the service account to which the RoleBinding is bound. 
In this case, it binds the Role poddepl to the default service account in the default namespace.

```
$ kubectl logs k8s-go-6767b78c75-rd64p
error stat /home/oviner/ClusterPath/auth/kubeconfig: no such file or directory building config from flags
pod in default ns:
k8s-go-6767b78c75-rd64p
deployment in default ns::
k8s-go
```