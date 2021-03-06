#-------------------------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
#-------------------------------------------------------------------------------------------------------------

# When we update the base image version (which we do manually, prompted by Dependabot
# notifying us of a new Go version), make sure our new base images is listed at:
# https://hub.docker.com/_/golang
FROM golang:1.15.4-buster

# This Dockerfile adds a non-root user with sudo access. Use the "remoteUser"
# property in devcontainer.json to use it. On Linux, the container user's GID/UIDs
# will be updated to match your local UID/GID (when using the dockerFile property).
# See https://aka.ms/vscode-remote/containers/non-root-user for details.
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Configure apt, install packages and tools
# hadolint ignore=DL3003,DL4006
RUN apt-get update \
    && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends apt-utils=1.8.2.1 dialog=1.3-20190211-1 2>&1 \
    # Need zip and unzip for the Hoverfly installation
    && apt-get -y install --no-install-recommends zip=3.0-11+b1 unzip=6.0-23+deb10u1 \
    #
    # Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
    && apt-get -y install --no-install-recommends \
        git=1:2.20.1-2+deb10u3 \
        openssh-client=1:7.9p1-10+deb10u2 \
        less=487-0.1+b1 \
        iproute2=4.20.0-2 \
        procps=2:3.3.15-2 \
        lsb-release=10.2019051400 \
    # Install Zsh
    && apt-get -y install --no-install-recommends zsh=5.7.1-1 \
    && wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | zsh || true \
    #
    # Build Go tools w/module support
    && mkdir -p /tmp/gotools \
    && cd /tmp/gotools \
    && GOPATH=/tmp/gotools GO111MODULE=on go get -v golang.org/x/tools/gopls@v0.5.3 2>&1 \
    && GOPATH=/tmp/gotools GO111MODULE=on go get -v \
        github.com/mdempsky/gocode@latest \
        github.com/sqs/goreturns@latest \
        # gopkgs is still needed even with gopls: https://github.com/microsoft/vscode-go/issues/3050#issuecomment-592263369
        github.com/uudashr/gopkgs/v2/cmd/gopkgs@v2.1.2  \
        github.com/ramya-rao-a/go-outline@1.0.0  \
        github.com/acroca/go-symbols@v0.1.1  \
        github.com/rogpeppe/godef@v1.1.2  \
        github.com/fatih/gomodifytags@v1.13.0  \
        github.com/go-delve/delve/cmd/dlv@v1.5.0 2>&1 \
        github.com/golangci/golangci-lint/cmd/golangci-lint@v1.32.2 \
    #
    # Build gocode-gomod
    && GOPATH=/tmp/gotools go get -x -d github.com/stamblerre/gocode 2>&1 \
    && GOPATH=/tmp/gotools go build -o gocode-gomod github.com/stamblerre/gocode \
    #
    # Install Go tools
    && mv /tmp/gotools/bin/* /usr/local/bin/ \
    && mv gocode-gomod /usr/local/bin/ \
    #
    # Create a non-root user to use if preferred - see https://aka.ms/vscode-remote/containers/non-root-user.
    && groupadd --gid $USER_GID $USERNAME \
    && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME \
    # [Optional] Add sudo support
    && apt-get install -y --no-install-recommends sudo=1.8.27-1+deb10u2 \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME \
    #
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/* /tmp/gotools

# Update this to "on" or "off" as appropriate
# Updated from "auto" to "on" as we are using Go modules for this project
# See https://dev.to/maelvls/why-is-go111module-everywhere-and-everything-about-go-modules-24k
ENV GO111MODULE=on
