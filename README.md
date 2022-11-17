# Go Gena

![Go](https://github.com/mmogylenko/go-gena/workflows/Go/badge.svg) ![Gosec](https://github.com/mmogylenko/go-gena/workflows/Gosec/badge.svg) ![containerscan](https://github.com/mmogylenko/go-gena/workflows/ContainerScan/badge.svg) [![GitHub tag](https://img.shields.io/github/tag/mmogylenko/go-gena.svg)](https://github.com/mmogylenko/go-gena/tags/)

![go-gena](https://user-images.githubusercontent.com/7536624/92551931-42128880-f214-11ea-8ebb-a71817168353.png)


**go-gena** is a simple load generator docker image with *REST* written on Go. Everyone writes [K8s HPA](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/) Check with amazing `Code.education Rocks!` output but no-one publishes load generator image itself. Let's get it fixed!

*Web interface*

![go_gena_ui](https://user-images.githubusercontent.com/7536624/202356166-23d2523b-c02d-4160-a80e-95b08ff7731d.png)

## Installation
**go-gena** is available both from [Docker Hub](https://hub.docker.com/r/mmogylenko/go-gena) and [Github Container Registry](https://github.com/users/mmogylenko/packages/container/go-gena/)

Example:
```bash
docker pull mmogylenko/go-gena
```
## Run

To start load generator run container and expose *8080* port to local machine

```bash
docker run -p 8080:8080/tcp go-gena:latest
```

Using `curl` or [httpie](https://github.com/httpie/httpie) make GET request to `localhost:8080/calculate`
![go-gena-calculate](https://user-images.githubusercontent.com/7536624/92655559-7f275b00-f2a6-11ea-9ce5-3ddadfd68783.png)
## Examples

Check [examples](examples) directory for a use-cases of `go-gena` usage

## Licensing

This project is licensed under the Apache V2 License. See [LICENSE](LICENSE) for more information.
