# panic_recover
middlewares for web application framework.
provide at logging,writeErrorHttpResponse when panic occured.
this library designed for myframework [ghost](https://github.com/mappymappy/ghost)

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





