# Avail Monitor setup

## Prerequisites

- Go 1.20+
- Prometheus 2.x+
- Grafana

## Installing Prerequisites

- Installlation script
- Manual Installation

Either of the above methods can be used to install the prerequisites

## Installation-script

- [click here](./scripts/prerequisites.sh) to view Installation script
- Installation script downloads and installs Prometheus & Grafana, downloads Golang if it's not already installed.
- Execute the script using the following command

```bash
curl -s -L https://raw.githubusercontent.com/vitwit/avail-monitor/avail-develop/scripts/prerequisites.sh | bash
```

## Manual-Installation

- To install prerequisites manually, please follow this [guide](./docs/prerequisite_manual.md)

# Install and configure the tool 

There are two ways of installing the tool:-

 - Installation script
 - Manual installation

## 1) Installation script

It clones and sets up the monitoring service as a systemd service. Please export the following `env` variable before executing the script as the variable will be used to initialize the `config.toml` file of the tool.

```
export SIDECAR_ENDPOINT="<sidecar url>" # Ex- export SIDECAR_ENDPOINT="https://sidecar.goldberg.avail.tools" 
```

You can find the tool installation script [here](./scripts/tool_installation.sh)

You can execute the script using the following one-liner
```bash
curl -s -L https://raw.githubusercontent.com/vitwit/avail-monitor/avail-develop/scripts/tool_installation.sh | bash
```

You can check the logs of the tool using:
```
journalctl -u avail-monitor.service -f
```

## 2) Manual Installation

```bash
git clone https://github.com/vitwit/avail-monitor.git
cd avail-monitor
cp example.config.toml config.toml
```
Open the `config.toml` and enter the sidecar url you want to connect to.

Build and run the monitoring binary
```bash
go build -o avail-monitor

# Generate systemd file for the tool
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
```
You can see all the metrics on http://localhost:1234/metrics




