rands
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fissue9%2Frands%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/issue9/rands/goto?ref=master)
[![license](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)
[![codecov](https://codecov.io/gh/issue9/rands/branch/master/graph/badge.svg)](https://codecov.io/gh/issue9/rands)
[![Go Reference](https://pkg.go.dev/badge/github.com/issue9/rands.svg)](https://pkg.go.dev/github.com/issue9/rands)
======

rands 为一个随机字符串生成工具。

```go
// 生成一个长度为 [8,10) 之间的随机字符串
str := rands.String(8, 10, []byte("1234567890abcdefg"))


// 生成一个带缓存功能的随机字符串生成器
r, err := rands.New(time.Now().Unix(), 100, 5, 7, []byte("asdfghijklmn"))
ctx,cancel := context.WithCancel(context.Background())
go r.Serve(ctx)
defer cancel()
str1 := r.String()
str2 := r.String()
```

安装
----

```shell
go get github.com/issue9/rands/v2
```

版权
----

本项目采用 [MIT](https://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
