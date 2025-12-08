#!/bin/bash
set -e

GO_VERSION="1.25.4"
GO_TARBALL="go${GO_VERSION}.linux-amd64.tar.gz"
GO_URL="https://dl.google.com/go/${GO_TARBALL}"

echo ">>> Updating package list..."
sudo apt update

echo ">>> Installing required dependencies..."
sudo apt install -y \
    iproute2 \
    iputils-ping \
    net-tools \
    tcpdump \
    tree \
    wget \
    vim \
    libpcap-dev

echo ">>> Downloading Go $GO_VERSION..."
wget -q "$GO_URL" -O "$GO_TARBALL"

echo ">>> Removing previous Go installation (if any)..."
sudo rm -rf /usr/local/go

echo ">>> Extracting Go $GO_VERSION..."
sudo tar -C /usr/local -xzf "$GO_TARBALL"

echo ">>> Removing downloaded tarball..."
rm -f "$GO_TARBALL"

echo ">>> Configuring Go environment variables..."
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc

source ~/.bashrc

echo ">>> Go $GO_VERSION installation completed successfully."
go version