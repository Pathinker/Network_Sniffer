FROM golang:1.25-bookworm

RUN apt update && apt install -y \
    iproute2 \
    iputils-ping \
    net-tools \
    tcpdump \
    tree \
    wget \
    vim \
    libpcap-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /sniffer