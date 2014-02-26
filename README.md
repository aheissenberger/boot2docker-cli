# boot2docker-cli

This is the Go port of boot2docker (https://github.com/boot2docker/boot2docker)
management script. It is intended to replace the shell script eventually. It is
currently usable but since it is under active development, frequent changes and
bugs are expected. USE AT YOUR OWN RISK. 

The Go port will produce a single binary without extra dependencies for the
following platforms:

- Linux/386
- Linux/amd64
- OS X (Darwin)/386
- OS X (Darwin)/amd64
- Windows/386
- Windows/amd64


## Installation

Once you have your `$GOPATH` properly setup, run

    go get github.com/boot2docker/boot2docker-cli/boot2docker


The binary will be available at `$GOPATH/bin/boot2docker`.

If you don't want to install golang you can use the `Dockerfile` to create the
binary for any supported target platform.

1. Build the image with the Go toolchain: `docker build -t boot2docker-golang .`
2. Choose the right settings for your plattform
  * Windows: `GOOS=windows GOARCH=amd64`
  * OS X: `GOOS=darwin GOARCH=amd64`
  * Linux: `GOOS=linux GOARCH=amd64`
3. Build the binaries
```
docker run -e GOOS=darwin -e GOARCH=amd64 --name boot2docker-cli boot2docker-golang
docker cp boot2docker-cli:/data/boot2docker-cli .
docker rm boot2docker-cli
```

The binary `boot2docker` will be in your current folder. Please do not forget to
rename the binary on Windows to `boot2docker.exe`


## Usage

To initialize a new boot2docker VM, run

    boot2docker init

Then you can start the VM by

    boot2docker up

To stop the VM, run

    boot2docker down

And finally if you don't need the VM anymore, run

    boot2docker delete

to remove it completely. 


## Configuration

The `boot2docker` binary reads configuration from the environment. Currently you can
configure the following options:

- `BOOT2DOCKER_VBM` path to the VirtualBox management utility, default is
  `VBoxManage`.
- `BOOT2DOCKER_SSH` path to the `ssh` client utility, default is `ssh`.
- `BOOT2DOCKER_VM` name of the boot2docker virtual machine in VirtualBox,
  default is `boot2docker-vm`.
- `BOOT2DOCKER_DIR` path to the boot2docker config directory, default is
  `$HOME/.boot2docker`.
- `BOOT2DOCKER_ISO` path to the boot2docker ISO image, default is
  `$BOOT2DOCKER_DIR/boot2docker.iso`.
- `BOOT2DOCKER_DISK` path to the boot2docker disk image, default is
  `$BOOT2DOCKER_DIR/boot2docker.vmdk`.
- `BOOT2DOCKER_DISKSIZE` boot2docker disk image size in MB, default is `20000`.
- `BOOT2DOCKER_MEMORY` boot2docker VM memory size in MB, default is `1024`.
- `BOOT2DOCKER_SSH_PORT` port on the host forwarding to port 22 in boot2docker
  VM, default is `2022`.
- `BOOT2DOCKER_DOCKER_PORT` port on the host forwarding to port 4243 in
  boot2docker VM, default is `4243`.


You can put custom options into your shell, e.g.

    export BOOT2DOCKER_VBM=VBoxManage
    export BOOT2DOCKER_SSH=ssh
    export BOOT2DOCKER_VM=boot2docker-vm
    export BOOT2DOCKER_DIR=$HOME/.boot2docker
    export BOOT2DOCKER_ISO=$BOOT2DOCKER_DIR/boot2docker.iso
    export BOOT2DOCKER_DISK=$BOOT2DOCKER_DIR/boot2docker.vmdk
    export BOOT2DOCKER_DISKSIZE=20000
    export BOOT2DOCKER_MEMORY=1024
    export BOOT2DOCKER_SSH_PORT=2022
    export BOOT2DOCKER_DOCKER_PORT=4243

**What is the development process**

We are implementing the same process as [Docker merge approval](https://github.com/dotcloud/docker/blob/master/CONTRIBUTING.md#merge-approval), so all commits need to be done via pull requests, and will need 3 or more LGTMs.
