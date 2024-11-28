module github.com/couriourc/supervisord-plus/config

go 1.23.2

require (
	github.com/couriourc/supervisord-plus/updater v0.0.0-00010101000000-000000000000
	github.com/couriourc/supervisord-plus/util v0.0.0-00010101000000-000000000000
	github.com/expr-lang/expr v1.16.9
	github.com/ochinchina/go-ini v1.0.1
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/cenkalti/backoff/v4 v4.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/kr/binarydist v0.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/reactivex/rxgo/v2 v2.5.0 // indirect
	github.com/sanbornm/go-selfupdate v0.0.0-20230714125711-e1c03e3d6ac7 // indirect
	github.com/stretchr/objx v0.1.0 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/teivah/onecontext v0.0.0-20200513185103-40f981bfd775 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
	golang.org/x/tools v0.1.10 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace (
	github.com/couriourc/supervisord-plus/updater => ../updater
	github.com/couriourc/supervisord-plus/util => ../util
)
