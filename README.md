# xk6-g0

Write k6 tests in golang

The xk6-g0 extension allows writing k6 tests in the go language.

**script.go**

```go
package main

import "net/http"

func Default() {
  http.Get("https://test.k6.io")
}
```

**run**
```bash
./k6 run script.go
```

Although k6's officially supported scripting language is JavaScript, support for other languages appears from time to time. In this blog post, you can read an understandable and clear explanation of why k6 officially supports only JavaScript language: [Why k6 does not support multiple scripting languages?](https://k6.io/blog/why-k6-does-not-introduce-multiple-scripting-languages/)

xk6-g0 is an experiment to use the go programming language as a full-fledged script language in k6 tests with the support of the community. Since k6 extensions (including xk6-g0) are made in the go language, every xk6-g0 user is also a potential contributor to xk6-g0 development. If the community really wants to use the go programming language to write k6 tests, hopefully they will be committed enough to contribute to xk6-g0 development. Otherwise, xk6-g0 remains an interesting experiment.

When using xk6-g0, the tests are executed by a built-in go interpreter ([yaegi](https://github.com/traefik/yaegi)), so there is no need for a compilation or build phase. It is true that the speed of interpreted execution does not reach the speed of compiled code, but it has many advantages. On the other hand, even with JavaScript support, the interpreter performs the tests.

Accepting the [benchmark measurement](https://github.com/d5/tengo#benchmark) made with [tengo](https://github.com/d5/tengo)'s developer, the JavaScript ([goja](https://github.com/dop251/goja)) interpreter calculates Fibonacci numbers twice as fast as the go interpreter ([yaegi](https://github.com/traefik/yaegi)). Well, it is relatively rare to count fibonacci numbers in tests, so we are not far off with the approximation that there is no double multiplier in execution speed. (someday a more accurate measurement would be useful)

## Usage

The go script should be put in the `main` package. The following lifecycle callback functions and configuration object can be exported by the script:
  - Setup: corresponds to the [setup function](https://k6.io/docs/using-k6/test-lifecycle/#setup-and-teardown-stages) of the JavaScript API
  - TearDown: corresponds to the [teardown function](https://k6.io/docs/using-k6/test-lifecycle/#setup-and-teardown-stages) of the JavaScript API
  - Default: corresponds to the [default export function](https://k6.io/docs/using-k6/test-lifecycle/#the-vu-stage) of the JavaScript API
  - HandleSummary: corresponds to the [handleSummary function](handlesummary) of the JavaScript
  - Options: corresponds to the JavaScript API [options object](https://k6.io/docs/using-k6/k6-options/reference/)

The return values of the lifecycle callback functions are optional, they can also be defined without a return value. The function parameters are also optional, but their order is fixed.

The script is executed similarly to the JavaScript language:

```bash
./k6 run scripts/simple/script.go
```

### Setup

`Setup` function corresponds to the [setup function](https://k6.io/docs/using-k6/test-lifecycle/#setup-and-teardown-stages) of the JavaScript API.

```go
func Setup() (interface{}, error)
```

The return values are optional, i.e. in addition to the above form, the following can also be used:

```go
func Setup() interface{}
func Setup() error
func Setup()
```

Optionally, the following parameters can also be used:
  - [context.Context](https://pkg.go.dev/context) execution context
  - [assert.Assertions](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions) provides assertion methods
  - [require.Assertions](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions) provides assertion methods

```go
func Setup(context.Context, assert.Assertions) (interface{}, error)
```

### Teardown

`TearDown` function corresponds to the [teardown function](https://k6.io/docs/using-k6/test-lifecycle/#setup-and-teardown-stages) of the JavaScript API.

```go
func Teardown(data interface{}) error
```

The parameter and the return value are optional, i.e. in addition to the above form, the following can also be used:

```go
func Teardown(data interface{})
func Teardown() error
func Teardown()
```

Optionally, the following parameters can also be used:
  - [context.Context](https://pkg.go.dev/context) execution context
  - [assert.Assertions](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions) provides assertion methods
  - [require.Assertions](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions) provides assertion methods

```go
func Teardown(ctx context.Context, assert assert.Assertions, data interface{}) error
```

### Default

`Default` function corresponds to the [default export function](https://k6.io/docs/using-k6/test-lifecycle/#the-vu-stage) of the JavaScript API.

```go
func Default(data interface{}) error
```

The parameter and the return value are optional, i.e. in addition to the above form, the following can also be used:

```go
func Default(data interface{})
func Default() error
func Default()
```

Optionally, the following parameters can also be used:
  - [context.Context](https://pkg.go.dev/context) execution context
  - [assert.Assertions](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions) provides assertion methods
  - [require.Assertions](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions) provides assertion methods

```go
func Default(ctx context.Context, assert assert.Assertions, data interface{}) error
```

### HandleSummary

`HandleSummary` function corresponds to the [handleSummary function](handlesummary) of the JavaScript.

```go
func HandleSummary(data map[string]interface{}) (map[string]interface{}, error)
```

The error return value is optional, i.e. in addition to the above form, the following can also be used:

```go
func HandleSummary(data map[string]interface{}) map[string]interface{}
```

### Options

`Options` variable corresponds to the JavaScript API [options object](https://k6.io/docs/using-k6/k6-options/reference/).

```go
var Options map[string]interface{}
```

## Roadmap

The xk6-g0 is currently in Proof of Concept status. The further fate of the development depends on the community's feedback on the usefulness of the concept.

**Is it useful to support the go language (yaegi interpreter) in k6 tests?**
You can vote here: https://github.com/szkiba/xk6-g0/discussions/1

## API

**The primary API design consideration: don't have an API at all.**

There are many popular packages for the go programming language, xk6-g0 tries to implement the necessary functionality by integrating and supporting these packages without its own API. This approach has many advantages, such as:

  - the test writer does not need to learn a new API
  - test scripts can be tested using standard go testing

In addition to the go standard library, the following third-party packages can be used:

  - https://github.com/go-resty/resty
  - https://github.com/sirupsen/logrus
  - https://github.com/stretchr/testify
  - https://github.com/PuerkitoBio/goquery
  - https://github.com/tidwall/gjson
  - https://github.com/PaesslerAG/jsonpath
  - https://github.com/santhosh-tekuri/jsonschema/v5
  - https://github.com/brianvoe/gofakeit/v6


### Checks

The `Default` function's optional [assert.Assertions](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions) or [require.Assertions](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions) parameters can be used to define the [k6 checks](https://k6.io/docs/using-k6/checks/). The name of the check will be the message parameter of the corresponding assertion function.

Of course, metrics are also created from the checks defined in this way.

```go
package main

import (
  "net/http"

  "github.com/stretchr/testify/assert"
)

func Default(assert *assert.Assertions) {
  res, err := http.Get("https://httpbin.test.k6.io/get")

  assert.NoError(err, "got response without error")
  assert.Equal(http.StatusOK, res.StatusCode, "status code was 200")
  assert.Equal("application/json", res.Header.Get("Content-Type"), "content type was application/json")
}
```

**output**

```plain

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: -
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)


     ✓ got response without error
     ✓ status code was 200
     ✓ content type was application/json

     checks.....................: 100.00% ✓ 3        ✗ 0
     data_received..............: 6.0 kB  14 kB/s
     data_sent..................: 457 B   1.1 kB/s
     http_req_blocked...........: avg=302.7ms  min=302.7ms  med=302.7ms  max=302.7ms  p(90)=302.7ms  p(95)=302.7ms 
     http_req_connecting........: avg=124.32ms min=124.32ms med=124.32ms max=124.32ms p(90)=124.32ms p(95)=124.32ms
     http_req_duration..........: avg=126.6ms  min=126.6ms  med=126.6ms  max=126.6ms  p(90)=126.6ms  p(95)=126.6ms 
     http_req_receiving.........: avg=355.31µs min=355.31µs med=355.31µs max=355.31µs p(90)=355.31µs p(95)=355.31µs
     http_req_sending...........: avg=54.2µs   min=54.2µs   med=54.2µs   max=54.2µs   p(90)=54.2µs   p(95)=54.2µs  
     http_req_tls_handshaking...: avg=151.39ms min=151.39ms med=151.39ms max=151.39ms p(90)=151.39ms p(95)=151.39ms
     http_req_waiting...........: avg=126.19ms min=126.19ms med=126.19ms max=126.19ms p(90)=126.19ms p(95)=126.19ms
     http_reqs..................: 1       2.326236/s
     iteration_duration.........: avg=429.68ms min=429.68ms med=429.68ms max=429.68ms p(90)=429.68ms p(95)=429.68ms
     iterations.................: 1       2.326236/s
```

### HTTP client

From the http package of the standard library, metrics are created on the use of the following:

- http.DefaultClient
- http.Get, http.Head, http.Post, http.PostForm

In addition, the https://github.com/go-resty/resty HTTP client can also be used, metrics are generated from its use.

```go
package main

import "github.com/go-resty/resty/v2"

func Default() error {
  _, err := client.R().Get("https://httpbin.test.k6.io/get")

  return err
}

var client *resty.Client

func init() {
  client = resty.New()
}
```

### HTML

HTML documents can be parsed and manipulated using the popular [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery) package, which brings a syntax and a set of features similar to jQuery to the Go language.

```go
package main

import (
  "github.com/PuerkitoBio/goquery"
  "github.com/sirupsen/logrus"
)

func Default() error {
  doc, err := goquery.NewDocument("https://test.k6.io")
  if err != nil {
    return err
  }

  logrus.Info(doc.Find("h1.title span.text-blue").Text())

  return nil
}
```

### JSON

The [gjson](https://github.com/tidwall/gjson) and [jsonpath](github.com/PaesslerAG/jsonpath) packages can be used to query JSON documents.

**gjson**
```go
package main

import (
  "net/http"

  "github.com/go-resty/resty/v2"
  "github.com/stretchr/testify/require"
  "github.com/tidwall/gjson"
)

func Default(require *require.Assertions) {
  res, err := resty.New().R().Get("https://httpbin.test.k6.io/get")

  require.NoError(err, "request success")
  require.Equal(http.StatusOK, res.StatusCode(), "status code 200")

  body := res.Body()

  val := gjson.GetBytes(body, "headers.Host").Str

  require.Equal("httpbin.test.k6.io", val, "headers.Host value OK")
}
```

**jsonpath**
```go
package main

import (
  "net/http"

  "github.com/PaesslerAG/jsonpath"
  "github.com/go-resty/resty/v2"
  "github.com/stretchr/testify/require"
)

func Default(require *require.Assertions) {
  body := make(map[string]interface{})
  res, err := resty.New().R().SetResult(&body).Get("https://httpbin.test.k6.io/get")

  require.NoError(err, "request success")
  require.Equal(http.StatusOK, res.StatusCode(), "status code 200")

  val, err := jsonpath.Get("$.headers.Host", body)

  require.NoError(err, "$.headers.Host no error")
  require.Equal("httpbin.test.k6.io", val, "$.headers.Host value OK")
}
```

### Logging

The https://github.com/sirupsen/logrus package can be used for logging in the test script.

```go
package main

import "github.com/sirupsen/logrus"

func Setup() interface{} {
  logrus.Info("Setup")

  return map[string]interface{}{
    "foo": "bar",
  }
}

func Default(data interface{}) {
  logrus.Info("Default", data)
}

func Teardown(data interface{}) {
  logrus.Info("Teardown", data)
}

func init() {
  logrus.Info("init")
}
```

### Context

The first parameter of the `Default` function is optionally a [context.Context](https://pkg.go.dev/context). This can be used to perform context aware operations and to access various context variables.

The usual k6 variables (eg `__VU`, `__ENV`, `__ITER`) and the variables of the `k6/execution` module can be accessed using the `Value` function of the context parameter.

```go
package main

import (
  "context"

  "github.com/sirupsen/logrus"
)

func Default(ctx context.Context) {
  vu := ctx.Value("__VU").(int64)
  env := ctx.Value("__ENV").(map[string]string)
  iter := ctx.Value("__ITER").(int64)

  logrus.Info(vu)
  logrus.Info(iter)
  logrus.Info(env["PATH"])
  logrus.Info(ctx.Value("execution.scenario.name"))
}
```

## Download

You can download pre-built k6 binaries from [Releases](https://github.com/szkiba/xk6-g0/releases/) page. Check [Packages](https://github.com/szkiba/xk6-g0/pkgs/container/xk6-g0) page for pre-built k6 Docker images.

## Build

You can build the k6 binary on various platforms, each with its requirements. The following shows how to build k6 binary with this extension on GNU/Linux distributions.

### Prerequisites

You must have the latest Go version installed to build the k6 binary. The latest version should match [k6](https://github.com/grafana/k6#build-from-source) and [xk6](https://github.com/grafana/xk6#requirements).

- [Git](https://git-scm.com/) for cloning the project
- [xk6](https://github.com/grafana/xk6) for building k6 binary with extensions

### Install and build the latest tagged version

1. Install `xk6`:

   ```shell
   go install go.k6.io/xk6/cmd/xk6@latest
   ```

2. Build the binary:

   ```shell
   xk6 build --with github.com/szkiba/xk6-g0@latest
   ```

> **Note**
> You can always use the latest version of k6 to build the extension, but the earliest version of k6 that supports extensions via xk6 is v0.43.1. The xk6 is constantly evolving, so some APIs may not be backward compatible.

### Build for development

If you want to add a feature or make a fix, clone the project and build it using the following commands. The xk6 will force the build to use the local clone instead of fetching the latest version from the repository. This process enables you to update the code and test it locally.

```bash
git clone git@github.com:szkiba/xk6-g0.git && cd xk6-g0
xk6 build --with github.com/szkiba/xk6-g0@latest=.
```

## Docker

You can also use pre-built k6 image within a Docker container. In order to do that, you will need to execute something like the following:

**Linux**

```plain
docker run -v $(pwd):/work -it --rm ghcr.io/szkiba/xk6-g0:latest run /work/scripts/simple/script.go
```

**Windows**

```plain
docker run -v %cd%:/work -it --rm ghcr.io/szkiba/xk6-g0:latest run /work/scripts/simple/script.go
```

## Example scripts

There are many examples in the [scripts](https://github.com/szkiba/xk6-g0/tree/master/scripts) directory that show how to use various features of the extension.

## Extending xk6-g0

xk6-g0 allows you to install additional packages in addition to the built-in go packages without changing the xk6-g0 source code. For this, for example, a function must be registered from the init() function of a custom k6 extension, which can be used to make additional packages available.

Check [xk6-g0-figure](https://github.com/szkiba/xk6-g0-figure) as an example addon.
