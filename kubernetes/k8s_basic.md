Pod: is the smallest and simplest unit in the Kubernetes object model. 
It represents a single instance of a running process in a cluster. 
A Pod can contain one or more containers, which are co-located and share the
same network namespace, allowing them to communicate with each other using localhost.

```
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
  - name: my-container
    image: nginx:latest
    ports:
    - containerPort: 80

```