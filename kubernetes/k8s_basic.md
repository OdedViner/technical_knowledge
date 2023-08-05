**Pod:**
Pod is the smallest and simplest unit in the Kubernetes object model. 
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
****************************************************************************************************
**Deployment:**
A Deployment is a high-level resource object used to manage the deployment and scaling of a set of identical pods. It provides a way to declare the desired state of the pods and ensures that the desired number of replicas (pods) are running at all times, handling updates and rollbacks when necessary.


```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app-deployment
  labels:
    app: my-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: my-app-container
        image: my_app_image:latest
        ports:
        - containerPort: 80
```


Explanation of the YAML file:
```
apiVersion: The API version of the Kubernetes resource being defined. In this case, it's apps/v1, which corresponds to the Deployment API.
kind: The type of resource being defined. Here, it's a Deployment.
metadata: Metadata associated with the deployment, including the name and labels to identify it.
spec: Specifies the desired state of the deployment.
  replicas: The number of identical pods (replicas) that should be running at any given time. In this example, it's set to 3.
  selector: Defines how the Deployment identifies the pods it manages. It uses the labels to match the pods to manage.
  template: Describes the pod template that will be used to create new replicas.
    metadata: Labels to identify the pods created from this template.
    spec: Specifies the pod's specification, including containers, volumes, and other settings.
      containers: Describes the containers within the pod.
        name: The name of the container.
        image: The container image to use for the pod.
        ports: Specifies which ports to expose from the container.
```

```
kubectl apply -f deployment.yaml
```
*******************************************************************************************************
**StatefulSet**
StatefulSet is another type of controller, like Deployment, but it is specifically designed to manage stateful applications. Unlike a Deployment, which is used for stateless applications, a StatefulSet provides guarantees about the identity and stable network identity of its pods. StatefulSets are particularly useful for applications that require stable network identities and persistent storage, such as databases or other stateful services.

```
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: my-database
spec:
  serviceName: my-database
  replicas: 3
  selector:
    matchLabels:
      app: my-database
  template:
    metadata:
      labels:
        app: my-database
    spec:
      containers:
      - name: database-container
        image: my_database_image:latest
        ports:
        - containerPort: 3306
        volumeMounts:
        - name: data-volume
          mountPath: /data
  volumeClaimTemplates:
  - metadata:
      name: data-volume
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 10Gi
```

```
apiVersion: The API version of the Kubernetes resource being defined
kind: The type of resource being defined. Here, it's a StatefulSet.
metadata: Metadata associated with the StatefulSet, including the name.
spec: Specifies the desired state of the StatefulSet.
  serviceName: The name of the Kubernetes Service that will be automatically created to manage the network identities of the pods.
  replicas: The number of identical pods (replicas) that should be running at any given time. In this example, it's set to 3.
  selector: Defines how the StatefulSet identifies the pods it manages. It uses the labels to match the pods to manage.
  template: Describes the pod template that will be used to create new replicas.
    metadata: Labels to identify the pods created from this template.
    spec: Specifies the pod's specification, including containers, volumes, and other settings.
      containers: Describes the containers within the pod.
        name: The name of the container.
        image: The container image to use for the pod.
        ports: Specifies which ports to expose from the container.
        volumeMounts: Mounts the persistent volume to the container's filesystem.
    volumeClaimTemplates: Specifies a template for the PersistentVolumeClaims (PVCs) that will be automatically created.
```

*******************************************************************************************************
**ServiceAccount**
ServiceAccount is an identity associated with a pod or a group of pods that allows them to interact with the Kubernetes API server and access resources in the cluster. ServiceAccounts are used for authentication and authorization purposes, enabling fine-grained control over what resources a pod can access and what actions it can perform within the cluster.

When a pod is created without explicitly specifying a ServiceAccount, Kubernetes automatically assigns a default ServiceAccount to the pod. This default ServiceAccount has limited permissions, typically sufficient for most basic pods to function but may not have access to certain resources or actions.

```
apiVersion: v1
kind: ServiceAccount
metadata:
  name: my-service-account
```

Once the pods are running with the specified ServiceAccount, you can apply Kubernetes RBAC (Role-Based Access Control) to grant specific permissions to that ServiceAccount. RBAC allows you to control who can perform specific actions on various resources within the cluster.
*******************************************************************************************************

**PVC**
A PersistentVolumeClaim is a resource that allows a pod to request persistent storage with specific characteristics from a storage class. 

```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
```
kind: The type of resource being defined. Here, it's a PersistentVolumeClaim.
metadata: Metadata associated with the PVC, including the name.
spec: Specifies the desired state of the PVC.
  accessModes: [ReadWriteOnce, ReadOnlyMany, and ReadWriteMany]
  resources: Specifies the requested storage resources for the PVC.
  requests: The amount of storage requested. In this example 10Gi


```
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  volumes:
  - name: my-storage
    persistentVolumeClaim:
      claimName: my-pvc
  containers:
  - name: my-container
    image: my_app_image:latest
    volumeMounts:
    - name: my-storage
      mountPath: /data
```
The container within the pod can then use the /data directory to read and write data to the persistent storage provided by the PVC.


*******************************************************************************************************
**PV**
