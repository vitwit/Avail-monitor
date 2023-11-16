## Default port customization

### Grafana:-
To change the default port of your `Grafana` server please edit the following files: `/usr/share/grafana/conf/defaults.ini` and `/etc/grafana/grafana.ini`. Search for `3000` and replace it with your custom port. Restart the systemd service after making your changes.
```
sudo systemctl restart grafana-server
```
### Prometheus:-
To change the default port of `Prometheus` process add this flag `--web.listen-address="0.0.0.0:<port>"` in `/lib/systemd/system/prometheus.service` file. Ex -
```
[Service]
Type=simple
ExecStart=$HOME/go/bin/prometheus --config.file=$HOME/prometheus.yml --web.listen-address="0.0.0.0:5000"
Restart=always
RestartSec=3
LimitNOFILE=4096
```

For the changes to take effect reload the file and restart the process.
```
sudo systemctl daemon-reload
sudo systemctl restart prometheus.service
```

