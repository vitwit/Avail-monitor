#!/bin/bash

set -e

cd $HOME

echo "------- checking for go, will be installed if not installed already -------"

command_exists () {
    type "$1" &> /dev/null ;
}

if command_exists go ; then
    echo "Golang is already installed"
else
    echo "-----Installing dependencies-----"
    sudo apt-get update
    sudo apt-get install build essentials jq -y

    wget https://dl.google.com/go/go1.20.3.linux-amd64.tar.gz
    tar -xvf go1.20.3.linux-amd64.tar.gz
    sudo mv go /usr/local

    echo "-------updating bashrc-------"
    export GOPATH=$HOME/go
    export GOROOT=/usr/local/go
    export GOBIN=$GOPATH/bin
    export PATH=$PATH:/usr/local/go/bin:$GOBIN
    
    echo "" >> ~/.bashrc
    echo 'export GOPATH=$HOME/go' >> ~/.bashrc
    echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
    echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
    echo 'export PATH=$PATH:/usr/local/go/bin:$GOBIN' >> ~/.bashrc

    source ~/.bashrc

    mkdir -p "$GOBIN"
fi

echo "-------Installing grafana-------"

sudo apt-get install -y adduser libfontconfig1 musl

wget https://dl.grafana.com/oss/release/grafana_10.2.1_amd64.deb

sudo dpkg -i grafana_10.2.1_amd64.deb

echo "-------Starting grafana server using systemd-------"

sudo systemctl daemon-reload

sudo systemctl start grafana-server

echo "-------Installing prometheus-------"

wget https://github.com/prometheus/prometheus/releases/download/v2.22.1/prometheus-2.22.1.linux-amd64.tar.gz
tar -xvf prometheus-2.22.1.linux-amd64.tar.gz

tar -xvf prometheus-2.22.1.linux-amd64.tar.gz

cp prometheus-2.22.1.linux-amd64/prometheus $GOBIN

cp prometheus-2.22.1.linux-amd64/prometheus.yml $HOME

echo "----- writing prometheus.yml-----"

echo "
- job_name: prometheus
    static_configs:
      - targets: ['localhost:9090']
    
- job_name: avail
    static_configs:
      - targets: ['localhost:1234']" >> "$HOME/prometheus.yml"

echo "------- Setting up prometheus system service -------"

echo "[Unit]
Description=Prometheus
After=network-online.target

[Service]
Type=simple
ExecStart=$HOME/go/bin/prometheus --config.file=$HOME/prometheus.yml
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target" | sudo tee "/lib/systemd/system/prometheus.service"

echo "------starting prometheus -----------"

sudo systemctl daemon-reload
sudo systemctl enable prometheus.service
sudo systemctl start prometheus.service

echo "------- Installation completed successfully -------"
