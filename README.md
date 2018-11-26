# gin-mod

This is a project for go module studying with gin.

## How to start

- download project with git

```
git clone github.com/songjiayang/gin-mod
cd gin-mod
```

- run code

```
go run gin-mod.go
```

Then you will see the project auto download all dependence look like:

```
go: finding github.com/gin-contrib/sse latest
go: finding github.com/ugorji/go/codec latest
go: finding github.com/golang/protobuf/proto latest

[GIN-debug] [WARNING] Now Gin requires Go 1.6 or later and Go 1.7 will be required soon.

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

- list all required packages

```
$ go list -m all

gin-mod
github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7
github.com/gin-gonic/gin v1.3.0
github.com/golang/protobuf v1.2.0
github.com/mattn/go-isatty v0.0.4
github.com/ugorji/go/codec v0.0.0-20181120210156-7d13b37dbec6
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
gopkg.in/go-playground/validator.v8 v8.18.2
gopkg.in/yaml.v2 v2.2.1
```

- update required package version

```
$ go mod edit -require=github.com/gin-gonic/gin@v1.1.4
$ go list -m all | grep gin

github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7
github.com/gin-gonic/gin v1.1.4
```

- require module with local

```
module github.com/songjiayang/gin-mod/web

require (
	github.com/songjiayang/go-mod/lib v0.0.0
)

replace github.com/songjiayang/go-mod/lib => ../lib
```

## reference:

- [初窥 Go module](https://tonybai.com/2018/07/15/hello-go-module)
- [wiki/Modules](https://github.com/golang/go/wiki/Modules)