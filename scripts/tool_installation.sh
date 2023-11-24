#!/bin/bash

git clone https://github.com/vitwit/avail-monitor.git ~/avail-monitor

cd ~/avail-monitor

cp example.config.toml config.toml

sed -i "s/<sidecar endpoint>/$SIDECAR_ENDPOINT/g" config.toml

go build -o avail-monitor

echo "[Unit]
Description=Avail Monitor
After=network-online.target

[Service]
Type=simple
WorkingDirectory=$HOME/avail-monitor
ExecStart=$HOME/avail-monitor/avail-monitor 
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target" | sudo tee "/etc/systemd/system/avail-monitor.service"

sudo systemctl enable avail-monitor.service

sudo systemctl start avail-monitor.service