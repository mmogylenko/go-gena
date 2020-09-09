# Go Gena!

![Go](https://github.com/mmogylenko/go-gena/workflows/Go/badge.svg) ![Gosec](https://github.com/mmogylenko/kuberhealthy-aws-iam-role-check/workflows/go-gena/badge.svg) [![GitHub tag](https://img.shields.io/github/tag/mmogylenko/go-gena.svg)](https://github.com/mmogylenko/go-gena/tags/)

![go-gena](https://user-images.githubusercontent.com/7536624/92551931-42128880-f214-11ea-8ebb-a71817168353.png)


**go-gena** is a simple load generator docker image with *REST* written on Go. Everyone writes [K8s HPA](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/) Check with amazing `Code.education Rocks!` output but no-one publishes load generator image itself. Let's get it fixed!

### Installation
**go-gena** is available both from [Docker Hub](https://hub.docker.com/r/mmogylenko/go-gena) and [Github Container Registry](https://github.com/users/mmogylenko/packages/container/go-gena/)

Example:
```bash
docker pull mmogylenko/go-gena
```
### Run

To start load generator run container and expose *8080* port to local machine

```bash
docker run -p 8080:8080/tcp go-gena:latest
```

Using `curl` or [httpie](https://github.com/httpie/httpie) make GET request to `localhost:8080/start`
![go-gena-start](https://user-images.githubusercontent.com/7536624/92623296-e3342a00-f27a-11ea-9130-57c03721fa97.png)
Stop load by making GET request to `localhost:8080/stop`
![gena-go-stop](https://user-images.githubusercontent.com/7536624/92623469-1d9dc700-f27b-11ea-8c46-1691d5659d85.png)
### Licensing

This project is licensed under the Apache V2 License. See [LICENSE](LICENSE) for more information.