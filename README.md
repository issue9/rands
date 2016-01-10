rand [![Build Status](https://travis-ci.org/issue9/rand.svg?branch=master)](https://travis-ci.org/issue9/rand)
======

rand 为一个随机字符串生成工具。
```go
// 生成一个长度为[8,10)之间的随机字符串，包含小写与数字字符
str := rand.String(8, 10, Lower, Digit)
```

### 安装

```shell
go get github.com/issue9/rand
```


### 文档

[![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/issue9/rand)
[![GoDoc](https://godoc.org/github.com/issue9/rand?status.svg)](https://godoc.org/github.com/issue9/rand)


### 版权

本项目采用[MIT](http://opensource.org/licenses/MIT)开源授权许可证，完整的授权说明可在[LICENSE](LICENSE)文件中找到。
