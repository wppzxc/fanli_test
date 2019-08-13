# fanli_test

## download mingw 
```
https://codeload.github.com/go-vgo/Mingw/zip/master
```

## set env
```
Path = D:\<your mingw parent path>\mingw\bin
```
## build
```
go build -o fanli.exe
```

## run
```
fanli.exe --config ./example/config.yaml --v=0
```
## history api
```
old:
  token:
    GetTokenUrlPrefix = "http://v2.yituike.com/admin/Weixinm/login?access_token=098f6bcd4621d373cade4e832627b4f6&openid=os_Ph0j7eHfGJhowF8E-_kJc1fiM&invite_id=208"
    //os_Ph0j7eHfGJhowF8E-_kJc1fiM is the id from wx
  items:
    GetProcessUrl    = "http://v2.yituike.com/fans/fans/proxy_goods?state=1&page=1&limit=10"
    GetPremonitorUrl = "http://v2.yituike.com/fans/fans/proxy_goods?state=2&page=1&limit=10"

new:
  token:
    http://v2.yituike.com/index/login/login
  items:
    GetProcessUrl    = "http://v2.yituike.com/admin/goods/goods_store?limit=100&page=1&state=1"
    GetPremonitorUrl = "http://v2.yituike.com/admin/goods/goods_store?limit=100&page=1&state=2"
```