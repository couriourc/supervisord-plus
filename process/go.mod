module github.com/couriourc/supervisord-plus/process

go 1.16

require (
	github.com/ochinchina/filechangemonitor v0.3.1
	github.com/prometheus/client_golang v1.10.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.8.1
)

replace (
	github.com/couriourc/supervisord-plus/config => ../config
	github.com/couriourc/supervisord-plus/events => ../events
	github.com/couriourc/supervisord-plus/logger => ../logger
	github.com/couriourc/supervisord-plus/signals => ../signals
)