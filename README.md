# fireload
---
[![Build Status](https://travis-ci.org/CloudCom/fireload.svg?branch=master)](https://travis-ci.org/CloudCom/fireload) [![Coverage Status](https://coveralls.io/repos/CloudCom/fireload/badge.svg?branch=master)](https://coveralls.io/r/CloudCom/fireload?branch=master)
---

A load balancer for multiple Firebase instances.

## Installation

```go
go get -u github.com/CloudCom/fireload
```

## Usage

Import fireload

```go
import "github.com/CloudCom/fireload"
```

### Namespaces

Creating a namespace:

```go
instance1 := fireload.NewNamespace(“https://foo.firebaseIO.com”)
instance2 := fireload.NewNamespace(“https://bar.firebaseIO.com”)
```

You can also set/get Metadata on namespaces:

```go
instance1.Metadata.Set(“awesome”, true)
instance1.Metadata.Set("secret", "very very secret")

instance2.Metadata.Set(“secret”, “my-awesome-secret”)

if secret, ok := instance2.Metdata.Get("awesome"); !ok {
  println("instance2 is not awesome")
}
```

### Pool

Creating a pool:

```go
pool, err := fireload.NewPool(instance1, instance2)
if err != nil {
  log.Fatal(err)
}
```

Pools allow you to specify the selection strategy used when retrieving a Namespace.
  * Random (default)
  * RoundRobin

```go
pool.SetStrategy(fireload.StrategyRoundRobin)
```

Getting a namespace and using it:

```go
namespace := pool.Next()
f := firego.New(namespace.Domain)
```