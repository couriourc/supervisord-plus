module github.com/couriourc/supervisord-plus/config

go 1.23.2

require (
	github.com/couriourc/supervisord-plus/updater v0.0.0-00010101000000-000000000000
	github.com/couriourc/supervisord-plus/util v0.0.0-00010101000000-000000000000
	github.com/ochinchina/go-ini v1.0.1
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/kr/binarydist v0.1.0 // indirect
	github.com/sanbornm/go-selfupdate v0.0.0-20230714125711-e1c03e3d6ac7 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
)

replace (
	github.com/couriourc/supervisord-plus/updater => ../updater
	github.com/couriourc/supervisord-plus/util => ../util
)
