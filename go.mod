module github.com/couriourc/supervisord-plus

go 1.16

require (
	github.com/couriourc/supervisord-plus/config v0.0.0-20210503132557-74b0760cc12e
	github.com/couriourc/supervisord-plus/events v0.0.0-20210503132557-74b0760cc12e
	github.com/couriourc/supervisord-plus/faults v0.0.0-20210503132557-74b0760cc12e
	github.com/couriourc/supervisord-plus/logger v0.0.0-20210503132557-74b0760cc12e
	github.com/couriourc/supervisord-plus/process v0.0.0-20210503132557-74b0760cc12e
	github.com/couriourc/supervisord-plus/signals v0.0.0-20210503132557-74b0760cc12e
	github.com/couriourc/supervisord-plus/types v0.0.0-20210503132557-74b0760cc12e
	github.com/couriourc/supervisord-plus/util v0.0.0-20230902082938-c2cae38b7454
	github.com/couriourc/supervisord-plus/xmlrpcclient v0.0.0-20210503132557-74b0760cc12e
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/rpc v1.2.0
	github.com/jessevdk/go-flags v1.5.0
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/kardianos/service v1.2.0
	github.com/ochinchina/go-daemon v0.1.5
	github.com/ochinchina/go-ini v1.0.1
	github.com/ochinchina/go-reaper v0.0.0-20181016012355-6b11389e79fc
	github.com/ochinchina/gorilla-xmlrpc v0.0.0-20171012055324-ecf2fe693a2c
	github.com/prometheus/client_golang v1.10.0
	github.com/prometheus/common v0.23.0 // indirect
	github.com/rs/cors v1.11.1
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/sys v0.0.0-20210503080704-8803ae5d1324 // indirect
)

replace (
	github.com/couriourc/supervisord-plus/config => ./config
	github.com/couriourc/supervisord-plus/events => ./events
	github.com/couriourc/supervisord-plus/faults => ./faults
	github.com/couriourc/supervisord-plus/logger => ./logger
	github.com/couriourc/supervisord-plus/process => ./process
	github.com/couriourc/supervisord-plus/signals => ./signals
	github.com/couriourc/supervisord-plus/types => ./types
	github.com/couriourc/supervisord-plus/util => ./util
	github.com/couriourc/supervisord-plus/xmlrpcclient => ./xmlrpcclient
)
