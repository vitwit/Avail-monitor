# Avail-monitor

![](./docs/logo.jpg)

**Avail-monitor** is a monitoring tool to monitor the network health of Avail. It uses Prometheus to export the metrics and Grafana dashboard to monitor the health of those metrics.

Note:- This tool requires substrate sidecar api endpoints to be accesible. Please run the sidecar before deploying this tool as the sidecar endpoint is required in the `config.toml` file. If the sidecar is not set up you can set it up following the instructions present [here](https://github.com/paritytech/substrate-api-sidecar#npm-package-installation-and-usage)

# Installation

* [Click here](./INSTRUCTIONS.md) for the tool installation instructions.


# Features

* [Click here](./docs/metric-desc.md) to know about the metrics being exported.
