# panic_recover
middlewares for web application framework.
this library provide at logging,writeErrorHttpResponse when panic occured.


# install

```
go get github.com/mappymappy/panic_recover
```

# usage

```
n := ghost.CreateEmptyGhost()
n.AddMiddleware(recover.Default())
```

# customize

panic_recover
```
recover := recover.Custom(yourWriter,yourLogger,yourErrorHandler)
n.AddMiddleware(recover)
```





