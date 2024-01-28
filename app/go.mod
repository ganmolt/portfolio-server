module portfolio-server

go 1.21.4

replace controllers => ./controllers

replace models/user => ./models/user

replace controllers/dbpkg => ./controllers/dbpkg

replace controllers/crypto => ./controllers/crypto

replace controllers/users => ./controllers/users

replace controllers/basicauth => ./controllers/basicauth

replace controllers/works => ./controllers/works

require (
	github.com/gin-gonic/gin v1.9.1
	gorm.io/driver/mysql v1.5.2 // indirect
	gorm.io/gorm v1.25.6 // indirect
)

require (
	controllers v0.0.0-00010101000000-000000000000
	controllers/users v0.0.0-00010101000000-000000000000
	controllers/works v0.0.0-00010101000000-000000000000
	github.com/gin-contrib/cors v1.5.0
)

require (
	controllers/basicauth v0.0.0-00010101000000-000000000000 // indirect
	controllers/crypto v0.0.0-00010101000000-000000000000 // indirect
	controllers/dbpkg v0.0.0-00010101000000-000000000000 // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
)

require (
	github.com/bytedance/sonic v1.10.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.15.5 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.5.0 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	models/user v0.0.0-00010101000000-000000000000 // indirect
)
