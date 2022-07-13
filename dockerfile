# syntax=docker/dockerfile:1
# usage
# pwsh:
# $env:DOCKER_BUILDKIT=1; docker build --file Dockerfile --output bin . -t clone
# bash:
# DOCKER_BUILDKIT=1 docker build --file Dockerfile --output bin . -t clone

FROM golang:1.18 AS builder

RUN apt update
RUN apt install -y zip

COPY . /work
WORKDIR /work

RUN go mod download

RUN GOOS="windows" GOARCH="amd64" go build -o ./bin/windows_amd64/clone-all-the-repos.exe cmd/clone-all-the-repos.go
RUN GOOS="windows" GOARCH="arm64" go build -o ./bin/windows_arm64/clone-all-the-repos.exe cmd/clone-all-the-repos.go

RUN GOOS="linux" GOARCH="amd64" go build -o ./bin/linux_amd64/clone-all-the-repos cmd/clone-all-the-repos.go
RUN GOOS="linux" GOARCH="arm64" go build -o ./bin/linux_arm64/clone-all-the-repos cmd/clone-all-the-repos.go

RUN GOOS="darwin" GOARCH="amd64" go build -o ./bin/darwin_amd64/clone-all-the-repos cmd/clone-all-the-repos.go
RUN GOOS="darwin" GOARCH="arm64" go build -o ./bin/darwin_arm64/clone-all-the-repos cmd/clone-all-the-repos.go 

RUN zip -j clone-all-the-repos_windows_amd64.zip ./bin/windows_amd64/clone-all-the-repos.exe
RUN zip -j clone-all-the-repos_windows_arm64.zip ./bin/windows_arm64/clone-all-the-repos.exe

RUN zip -j clone-all-the-repos_linux_amd64.zip ./bin/linux_amd64/clone-all-the-repos
RUN zip -j clone-all-the-repos_linux_arm64.zip ./bin/linux_arm64/clone-all-the-repos

RUN zip -j clone-all-the-repos_darwin_amd64.zip ./bin/darwin_amd64/clone-all-the-repos
RUN zip -j clone-all-the-repos_darwin_arm64.zip ./bin/darwin_arm64/clone-all-the-repos

FROM scratch AS export-stage
COPY --from=builder /work/clone-all-the-repos_windows_amd64.zip .
COPY --from=builder /work/clone-all-the-repos_windows_arm64.zip .

COPY --from=builder /work/clone-all-the-repos_linux_amd64.zip .
COPY --from=builder /work/clone-all-the-repos_linux_arm64.zip .

COPY --from=builder /work/clone-all-the-repos_darwin_amd64.zip .
COPY --from=builder /work/clone-all-the-repos_darwin_arm64.zip .
