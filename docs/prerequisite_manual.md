### Install Grafana for Ubuntu
Download the latest .deb file and extract it:

```sh
$ cd $HOME
$ sudo apt-get install -y adduser libfontconfig1 musl
$ wget https://dl.grafana.com/oss/release/grafana_10.2.1_amd64.deb
$ sudo dpkg -i grafana_10.2.1_amd64.deb
```

Start the grafana server
```
$ sudo -S systemctl daemon-reload

$ sudo -S systemctl start grafana-server

The default port that Grafana runs on is 3000. 
```
### Install Prometheus 

```sh
$ cd $HOME
$ wget https://github.com/prometheus/prometheus/releases/download/v2.22.1/prometheus-2.22.1.linux-amd64.tar.gz
$ tar -xvf prometheus-2.22.1.linux-amd64.tar.gz
$ sudo cp prometheus-2.22.1.linux-amd64/prometheus $GOBIN
$ sudo cp prometheus-2.22.1.linux-amd64/prometheus.yml $HOME
```

- Add the following in prometheus.yml using your editor of choice

```sh
 scrape_configs:

- job_name: prometheus
    static_configs:
      - targets: ['localhost:9090']
    
- job_name: avail
    static_configs:
      - targets: ['localhost:1234']
 
```
- Setup Prometheus System service
 ```
 sudo nano /lib/systemd/system/prometheus.service
 ```
 
 Copy-paste the following:
 
 ```
 [Unit]
Description=Prometheus
After=network-online.target

[Service]
Type=simple
ExecStart=/home/ubuntu/go/bin/prometheus --config.file=/home/ubuntu/prometheus.yml
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
 ```
For the purpose of this guide it is assumed the `user` is using `ubuntu` as OS. Please make the required changes according to your preferred OS.

```sh
$ sudo systemctl daemon-reload
$ sudo systemctl enable prometheus.service
$ sudo systemctl start prometheus.service
```

To use custom bind ports for the prerequisites please follow these [instructions](./custom_port.md) 

#### Clean-up (Optional)

```
$ rm grafana_10.22.1_amd64.deb prometheus-2.22.1.linux-amd64.tar.gz
```