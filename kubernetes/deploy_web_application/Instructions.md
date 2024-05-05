1.Deploy Minikube  
```
minikube start --disk-size=10g --extra-disks=1 --driver kvm2
```
2.Create Web deployment based on my image quay.io/oviner/web-ci:latest  
```
kubectl create -f web-app-deployment.yaml
```
3.Create service:  
```
kubectl create -f web-app-service.yaml
```
4.Get URL:  
```
minikube service my-web-app-service --url
```