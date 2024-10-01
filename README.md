# qules-logger
**QulesLogger** is a simple linux system monitoring tool written in Go. primarily created for monitoring smaller servers, home PCs, and devices like raspberry pi, but it is now extended to support full fledged system level monitoring on linux systems. QulesLogger might also supports custom time series data collection in the future. which can be used to collect sensor, or application output/performance data.


Currently, it gathers following metrics under the monitoring interval given in the configuration file.

- System Info
- CPU Info
- Memory/SWAP info
- CPU/Memory usage graphs for past hour (default) or custom time range



**API documentation**
