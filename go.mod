module github.com/Crshi/Blog

go 1.14

require (
	github.com/astaxie/beego v1.12.1
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.54.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.4 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/ini.v1 v1.54.0 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace (
	github.com/Crshi/Blog/conf => ./conf
	github.com/Crshi/Blog/middleware => ./middleware
	github.com/Crshi/Blog/models => ./models
	github.com/Crshi/Blog/pkg/e => ./pkg/e
	github.com/Crshi/Blog/pkg/setting => ./pkg/setting
	github.com/Crshi/Blog/pkg/util => ./pkg/util
	github.com/Crshi/Blog/routers => ./routers
	github.com/Crshi/Blog/routers/api => ./routers/api
)
