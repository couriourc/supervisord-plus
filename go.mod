module github.com/couriourc/supervisord-plus

go 1.23.2

toolchain go1.23.3

require (
	github.com/couriourc/supervisord-plus/config v0.0.0-00010101000000-000000000000
	github.com/couriourc/supervisord-plus/events v0.0.0-00010101000000-000000000000
	github.com/couriourc/supervisord-plus/faults v0.0.0-00010101000000-000000000000
	github.com/couriourc/supervisord-plus/logger v0.0.0-00010101000000-000000000000
	github.com/couriourc/supervisord-plus/process v0.0.0-00010101000000-000000000000
	github.com/couriourc/supervisord-plus/signals v0.0.0-00010101000000-000000000000
	github.com/couriourc/supervisord-plus/types v0.0.0-00010101000000-000000000000
	github.com/couriourc/supervisord-plus/updater v0.0.0-00010101000000-000000000000
	github.com/couriourc/supervisord-plus/util v0.0.0-20230902082938-c2cae38b7454
	github.com/couriourc/supervisord-plus/xmlrpcclient v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/rpc v1.2.0
	github.com/jessevdk/go-flags v1.5.0
	github.com/kardianos/service v1.2.0
	github.com/ochinchina/go-daemon v0.1.5
	github.com/ochinchina/go-ini v1.0.1
	github.com/ochinchina/go-reaper v0.0.0-20181016012355-6b11389e79fc
	github.com/ochinchina/gorilla-xmlrpc v0.0.0-20171012055324-ecf2fe693a2c
	github.com/prometheus/client_golang v1.10.0
	github.com/rs/cors v1.11.1
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/net v0.31.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/expr-lang/expr v1.16.9 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/kr/binarydist v0.1.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/ochinchina/filechangemonitor v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.23.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/reactivex/rxgo/v2 v2.5.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rogpeppe/go-charset v0.0.0-20190617161244-0dc95cdf6f31 // indirect
	github.com/sanbornm/go-selfupdate v0.0.0-20230714125711-e1c03e3d6ac7 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/teivah/onecontext v0.0.0-20200513185103-40f981bfd775 // indirect
	golang.org/x/sync v0.9.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/couriourc/supervisord-plus/config => ./config
	github.com/couriourc/supervisord-plus/events => ./events
	github.com/couriourc/supervisord-plus/faults => ./faults
	github.com/couriourc/supervisord-plus/logger => ./logger
	github.com/couriourc/supervisord-plus/process => ./process
	github.com/couriourc/supervisord-plus/signals => ./signals
	github.com/couriourc/supervisord-plus/types => ./types
	github.com/couriourc/supervisord-plus/updater => ./updater
	github.com/couriourc/supervisord-plus/util => ./util
	github.com/couriourc/supervisord-plus/xmlrpcclient => ./xmlrpcclient
)
