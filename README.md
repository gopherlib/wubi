# 简介

汉字转五笔编码，五笔编码转汉字，支持自定义码表，默认集成了86版、98版、新世纪版本码表。

# 使用示例

```go
import "github.com/gopherlib/wubi"
```

以下代码使用86版五笔作为示例，其它98版、新世纪版使用方法类似。

## 86版汉字转五笔码
```go
c := wubi.New86()
codes := c.GetCode('在')
fmt.Println(codes) 

// [d dhf dhfd]
```

## 86版多个汉字转五笔码
```go
c := wubi.New86()
codes := c.GetCodes("干一行，爱一行，一行行，行行行，一行不行，行行不行")
fmt.Println(codes) 

// [[fggh] [g ggl ggll] [tf tfhh] [，] [ep epd epdc] [g ggl ggll] [tf tfhh] [，] [g ggl ggll] [tf tfhh] [tf tfhh] [，] [tf tfhh] [tf tfhh] [tf tfhh] [，] [g ggl ggll] [tf tfhh] [i gi gii] [tf tfhh] [，] [tf tfhh] [tf tfhh] [i gi gii] [tf tfhh]]
```

## 86版五笔码转汉字

```go
c := wubi.New86()
chars := c.GetChar("d")
fmt.Println(chars) 

// [在 石]
```

## 86版多个五笔码转汉字

```go
c := wubi.New86()
chars := c.GetChars([]string{"d", "r"})
fmt.Println(chars) 

// [[在 石] [白 的]]
```
