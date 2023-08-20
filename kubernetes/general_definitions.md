**ETCD in Kubernetes:**  
etcd is a distributed key-value store that is often used as the primary data store for Kubernetes clusters. 


**Kube-API Server:**  
It serves as the front-end for the Kubernetes control plane and is responsible for exposing the Kubernetes API, 
which allows users, administrators, and other components to interact with the cluster.


**Kube Controller Manager:**  
It is responsible for managing various controllers that regulate the state of the cluster, ensuring that the desired state matches the current state. 


**Kube Scheduler:**  
Its primary responsibility is to assign work (in the form of pods) to worker nodes within the cluster based on various factors and policies. 


**kubelet:**  
The kubelet is a critical component that ensures the proper execution of containers within Pods on each node in a Kubernetes cluster.


**kube proxy:**  
It is a network service that runs on each node in a Kubernetes cluster. 
Its primary responsibility is to maintain network rules on nodes and manage communication between Pods and Services. 


**Pod**   
Pod is the smallest deployable unit and represents a single instance of a running process in a cluster. 
Pods are the atomic units in which containers are deployed. While containers have become synonymous with 
the term "microservices," Pods are Kubernetes' way of running containers.


**Deployment:**  
a Deployment is a high-level resource that provides declarative updates to applications. 
It represents a desired state for a set of replica Pods and ensures that the actual state matches the desired state. 


**ReplicaSet:**  
A ReplicaSet is a Kubernetes resource used for ensuring that a specified number of replicas (identical copies) of a Pod are running at all times. 
It is part of the Kubernetes system for managing containerized applications and is often used for scaling and ensuring high availability of applications. 


**Service:**  
a Service is an abstract way to expose an application running in a set of Pods as a network service. 
Services allow you to decouple the internal workings of your application from the network, making it easy to access your 
application and ensuring that it remains reachable even as Pods are created or terminated. 


**Services Cluster IP**  
A Service with a Cluster IP type is a service that is only accessible from within the cluster. 
It provides a stable, internal IP address that can be used by other Pods and services within 
the same Kubernetes cluster to access the service. 


**Services - Loadbalancer:**  
A Service with a LoadBalancer type is a service that exposes an application to the external world, 
allowing external clients to access the service. It leverages a cloud provider's load balancer to distribute incoming traffic 
across the Pods that are part of the service.


**Namespace:**  
Namespace is a logical and virtual cluster within a physical Kubernetes cluster. 
Namespaces provide a way to divide a Kubernetes cluster into multiple isolated and distinct environments or projects, 
each with its own resources, policies, and objects. 


**Imperative vs Declarative**  
Imperative management: Directly running commands like kubectl create, kubectl apply, and terraform apply to specify each step in creating or modifying resources.
Declarative management: Using resource manifests (YAML or JSON files) to define the desired state of resources and letting
tools like kubectl or Terraform figure out how to achieve that state.


**manual scheduling:**  
manual scheduling refers to the process of manually selecting the node or specific worker machine where a particular Pod should be scheduled. 


**Labels and Selectors:**  
Labels and selectors are fundamental concepts in Kubernetes that allow you to categorize 
and select objects (such as Pods, Services, and Nodes) based on metadata. Labels are key-value pairs associated with Kubernetes objects, 
while selectors are used to filter and query objects based on these labels.

**Taints and Tolerations**  
Taints and Tolerations are mechanisms that allow you to influence the scheduling of Pods onto specific nodes in your cluster. 
A Taint is a property that you can set on a node, indicating that the node has some special requirement or restriction.
A Toleration is a property that you can set on a Pod, allowing it to tolerate (i.e., be scheduled onto) nodes with specific Taints.


**Node Selectors:**  
Node Selectors are a way to influence the scheduling of Pods onto specific nodes based on node labels.


**Node affinity:**  
Node affinity is a Kubernetes feature that allows you to influence the scheduling of Pods based on the affinity or anti-affinity they have with nodes. 
Node affinity rules help determine which nodes are eligible for scheduling Pods based on specific node labels, taints, and Pod affinity rules.


**Resource Requirements:**  
Resource requirements define the minimum amount of CPU and memory that a container needs to function correctly. 
These requirements are used by the Kubernetes scheduler to decide where to place the Pod in the cluster, 
ensuring that the requested resources are available on the selected node.


**Resource Limits:**  
Resource limits define the maximum amount of CPU and memory that a container is allowed to use. 
These limits are used by Kubernetes to protect the overall stability of the cluster.
Containers exceeding their limits are subject to throttling and potential termination.


**DaemonSet**  
A DaemonSet is a type of workload resource that ensures that a copy of a specific Pod runs on each node in a cluster. 
DaemonSets are typically used for tasks that need to run on all or a specific subset of nodes in the cluster, 
such as logging, monitoring, or networking agents. 


**Static Pods:**  
Static Pods are a type of Pod in Kubernetes that are not managed by the Kubernetes control plane but are instead directly managed 
by a kubelet on a specific node. Unlike regular Pods, which are scheduled and managed by the Kubernetes API server, 
Static Pods are defined and managed by configuration files on individual nodes. 


**rolling update**  
A rolling update is a strategy for updating or modifying an application by replacing existing instances (Pods) 
of the application with new ones gradually. This approach ensures that a certain number of Pods from the old version are 
running alongside a certain number of Pods from the new version during the update process.


**Rollbacks:**  
Rollbacks are the process of reverting an application to a previous version when issues or errors are encountered with a new deployment. 
Kubernetes provides the ability to roll back to a previous version of a Deployment or StatefulSet.


**ConfigMap:**  
A ConfigMap is a Kubernetes resource used to store configuration data in a key-value format. 
ConfigMaps allow you to decouple configuration details from container images, making it easier to manage and configure 
applications running in Kubernetes Pods. They are commonly used to store environment variables, 
configuration files, or any other configuration data that applications need.


**Secrets:**  
Secrets are a resource used to store sensitive information, such as passwords, API keys, and other confidential data. 
Secrets are designed to keep this data secure and allow it to be used by Pods or other resources 
running in the cluster without exposing the sensitive information in plain text. 


**Kubeconfig:**   
A Kubeconfig (Kubernetes configuration) is a file used to configure access to a Kubernetes cluster. 
It contains information such as the cluster's address, authentication credentials, 
and other configuration settings. Kubeconfigs are crucial for managing Kubernetes clusters, 
and they allow users, administrators, and automated processes to interact with the cluster.


**RBAC:**  
RBAC, or Role-Based Access Control, is a security mechanism in Kubernetes that enables administrators to control access 
to resources and actions within a cluster. RBAC allows you to define who (users, service accounts, or groups) 
can perform specific operations (create, read, update, delete) on Kubernetes resources (Pods, Services, ConfigMaps, etc.). 
RBAC is essential for managing cluster security and ensuring that users have the appropriate permissions to perform their tasks.


**ClusterRole:**  
ClusterRole is a resource that defines a set of rules for granting permissions at the cluster level. 
Unlike regular Role resources, which are scoped to a specific namespace, ClusterRole resources apply cluster-wide. 
ClusterRoles are a fundamental part of Kubernetes' Role-Based Access Control (RBAC) system and are used to control 
access to cluster-wide resources and actions.


***RoleBinding:**  
A RoleBinding is a resource that associates a specific Role with subjects (users, groups, or service accounts) within 
a single namespace. It grants permissions to those subjects within that namespace.


**ClusterRoleBinding:**   
A ClusterRoleBinding is similar to a RoleBinding, but it applies cluster-wide, granting permissions across all namespaces. 
It associates a ClusterRole with subjects.


**Service Account:**  
A Service Account is an identity used by Pods to interact with the Kubernetes API server and other cluster services securely. 
Service Accounts are associated with Pods and are used to control the permissions and access levels of those Pods when 
they interact with cluster resources. They provide a way to grant or restrict access to various parts of the 
Kubernetes API and other resources within a namespace.


**security contexts:**  
security contexts are configurations that allow you to define and control the security settings and permissions for Pods and containers.
Security contexts help you enforce security policies and harden the security of your containerized applications. 
They can be set at both the Pod and container levels.


**Persistent Volume**  
A Persistent Volume (PV) is a representation of a physical piece of storage in the cluster. 
PVs are cluster resources created by cluster administrators, and they abstract the details of the underlying storage system.
A PV has a unique name, a capacity (the amount of storage it provides), an access mode (e.g., ReadWriteOnce, ReadOnlyMany, ReadWriteMany), and other attributes.
PVs can be statically provisioned by administrators or dynamically provisioned by Storage Classes based on the storage requirements defined in PVCs.


**Persistent Volume Claim (PVC):**  
A Persistent Volume Claim (PVC) is a request for a specific amount and class of storage by a Pod. 
PVCs are created by users or applications and define the desired storage characteristics.
When a Pod needs persistent storage, it references a PVC in its configuration. 
The Kubernetes control plane then finds an available PV that matches the PVC's requirements and binds the PVC to the PV.
PVCs can request storage with specific access modes (e.g., ReadWriteOnce for single-node access or ReadWriteMany for multi-node access).

**Storage Classes:**  
A Storage Class is a resource definition that provides a way to describe different classes of storage, 
including the type of storage, performance characteristics, and other parameters.

Storage Classes enable dynamic provisioning of PVs based on the PVC's requirements. 
They are defined by cluster administrators and can be associated with a specific provisioner or storage solution.


**CNI:**  
CNI stands for Container Network Interface. It's a specification and standard for defining how network 
plugins interact with container runtimes like Docker, Kubernetes, and others. 
CNI plugins are responsible for creating and managing network connections for containers, 
enabling communication between containers on the same host and across different hosts in a cluster.


**Ingress:**  
An Ingress is an API object that manages external access to services within the cluster. 
It provides a way to configure rules for routing external traffic to different services and is commonly used for HTTP and HTTPS traffic. 
Ingress acts as a layer of abstraction between external clients and services, allowing you to define how traffic should be directed.
