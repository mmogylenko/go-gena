# Kubernetes Horizontal Pod Autoscaler Demo

- [Run & expose go-gena (via hpa-demo service)](#run---expose-go-gena--via-hpa-demo-service-)
- [Create Horizontal Pod Autoscaler](#create-horizontal-pod-autoscaler)
- [Generate Load](#generate-load)
- [Cleanup](#cleanup)

To demonstrate Horizontal Pod Autoscaler we will use a custom `go-gena` docker image.

## Run & expose go-gena (via hpa-demo service)
- create `hpa-demo` service:
```bash
➜  k8s-hpa-demo kubectl apply -f service.yml
service/hpa-demo created
➜  k8s-hpa-demo
```

- create `go-gena` deployment:
```bash
➜  k8s-hpa-demo kubectl apply -f deployment.yml
deployment.apps/go-gena created
➜  k8s-hpa-demo
```

- check Service and PODs:
```bash
➜  k8s-hpa-demo kubectl get svc
NAME       TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
hpa-demo   ClusterIP   10.111.158.188   <none>        80/TCP    3m3s
➜  k8s-hpa-demo kubectl get pods
NAME                      READY   STATUS    RESTARTS   AGE
go-gena-9b7cb9b86-vs59b   1/1     Running   0          116s
➜  k8s-hpa-demo
```

## Create Horizontal Pod Autoscaler

```bash
➜  k8s-hpa-demo kubectl autoscale deployment go-gena --cpu-percent=50 --min=1 --max=10
horizontalpodautoscaler.autoscaling/go-gena autoscaled
➜  k8s-hpa-demo
```
- check the current status of autoscaler by running:
```bash
➜  k8s-hpa-demo kubectl get hpa
NAME      REFERENCE            TARGETS   MINPODS   MAXPODS   REPLICAS   AGE
go-gena   Deployment/go-gena   0%/50%    1         10        1          38s
➜  k8s-hpa-demo
```

## Generate Load

Run `busybox` container and hit `hpa-demo` service:
```bash
➜  k8s-hpa-demo kubectl run -it --rm load-generator --image=busybox /bin/sh
If you don't see a command prompt, try pressing enter.
/ # while true; do wget -q -O- http://hpa-demo/api/calculate; done
{"message":249995994525.7887}
{"message":249995994525.7887}
{"message":249995994525.7887}
{"message":249995994525.7887}
{"message":249995994525.7887}
{"message":249995994525.7887}
{"message":249995994525.7887}

......
```

Wait a few minutes and check the CPU load by executing:
```bash
➜  k8s-hpa-demo kubectl get hpa
NAME      REFERENCE            TARGETS    MINPODS   MAXPODS   REPLICAS   AGE
go-gena   Deployment/go-gena   200%/50%   1         10        4          2m10s
```
Check pods:
```bash
➜  k8s-hpa-demo kubectl get pods
NAME                      READY   STATUS        RESTARTS   AGE
go-gena-9b7cb9b86-9phlx   1/1     Running       0          94s
go-gena-9b7cb9b86-ckltj   1/1     Running       0          33s
go-gena-9b7cb9b86-dkb8l   1/1     Running       0          33s
go-gena-9b7cb9b86-lf88b   1/1     Running       0          33s
```

Here, CPU consumption has increased to 200% of the request. As a result, the deployment was resized to 4 replicas


## Cleanup

- Stop `load-generator`
- Delete HPA/DEPLOYMENT/SERVICE
```yaml
➜  k8s-hpa-demo kubectl delete hpa go-gena
horizontalpodautoscaler.autoscaling "go-gena" deleted
➜  k8s-hpa-demo kubectl delete deployment go-gena
deployment.apps "go-gena" deleted
➜  k8s-hpa-demo kubectl delete service hpa-demo
service "hpa-demo" deleted
```
