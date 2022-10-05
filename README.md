# 概念

一个非常简单的`key-value`内存存储服务。

# API

`token`+`key`能唯一确认一条记录，`token`必须是大于等于32位的字母和数字组成，`key`没有要求。

`value`必须是是UTF8字符串，如果是二进制文件需要用base64等方式字符串化。

## 获取值

```
GET
/:token/:key
```

直接返回内容。

## 设定值

```
POST/PUT
/:token/:key

value
```

`value`直接作为body整个内容提交，成功会返回设定的`value`。