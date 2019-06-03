# Runner auto register
Gitlab Runner use docker use to automate register for scale runner
# Manual Build
## Requirements
* URL (Go to the runners section of Gitlab and grab url  )
* TOKEN (Go to the runners section of Gitlab and grab registration token )
* go version v11 up
* docker-ce (option with use container)
* gitlab-runner (option with use local)

# Build
## Clone
```
git clone https://github.com/piyapan/runner-auto.git
cd ./runner-auto
```
## Build 
```
make build

```
## Build docker image
```
make build-docker
```
# Install
## Local
```
make run
```
OR
```
./runner-auto
```
## Docker 
```
docker pull subaruqui/gitlab-runner:latest
docker run -d --name runner -v /var/run/docker.sock:/var/run/docker.sock -e URL=http://gitlab.exmaple.com -e TOKEN=xxxx subaruqui/gitlab-runner:latest
```
## Docker Compose
```
   runnner:
      image: subaruqui/gitlab-runner:latest
      volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      environment:
      - URL=http://gitlab.example.com
      - TOKEN=xxxxx
```
Scale runner with docker compose
```
docker-compose scale runner=10
```
# Clean
Clean file execute
```
make clean
```

