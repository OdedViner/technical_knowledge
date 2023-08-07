1.Deploy Minikube

minikube start --disk-size=10g --extra-disks=1 --driver kvm2

2.Create pv

$ kubectl create -f pv.yaml 
persistentvolume/my-pv created

3.Check pv status:

$ kubectl get pv
NAME    CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS   REASON   AGE
my-pv   1Gi        RWO            Retain           Available                                   15s

4.Create PVC:

$ kubectl create -f pvc.yaml 
persistentvolumeclaim/my-pvc created

5.Check pvc statsus:

$ kubectl get pvc 
NAME     STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
my-pvc   Bound    pvc-b46c5563-c3d8-46e4-b54d-e72b7c068bc6   1Gi        RWO            standard       17s

6.Create pod associate to pvc

$ kubectl create -f pod.yaml 
pod/my-pod created

7.Check Pod status:

$ kubectl get pod
NAME     READY   STATUS    RESTARTS   AGE
my-pod   1/1     Running   0          21s
