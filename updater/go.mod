module github.com/couriourc/supervisord-plus/updater

go 1.23.2

require (
	github.com/reactivex/rxgo/v2 v2.5.0
	github.com/sanbornm/go-selfupdate v0.0.0-20230714125711-e1c03e3d6ac7
)

require (
	github.com/cenkalti/backoff/v4 v4.0.0 // indirect
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/kr/binarydist v0.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.1.0 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/teivah/onecontext v0.0.0-20200513185103-40f981bfd775 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace (
	github.com/couriourc/supervisord-plus/config => ../config
	github.com/couriourc/supervisord-plus/util => ../util
)
