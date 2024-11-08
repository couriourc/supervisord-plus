module github.com/couriourc/supervisord-plus/config

go 1.16

require (
	github.com/couriourc/supervisord-plus/util v0.0.0-00010101000000-000000000000
	github.com/ochinchina/go-ini v1.0.1
	github.com/sirupsen/logrus v1.8.1
)

replace github.com/couriourc/supervisord-plus/util => ../util
