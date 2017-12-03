# go-dynaport

A Go library for getting free ports, dynamically. 

## Example

``` go
ports := dynaport.Get(3)
// ports is something like []int{10000, 10001, 10002}
```

Or

``` go
ports, err := dynaport.GetWithErr(3)
if err != nil {
  // handle err
}

// use ports...
```

Or 

```
ports := dynaport.GetS(3)
// ports is something like []string{"10000", "10001", "10002"}
```

## License

MIT

--- 

- [travisjeffery.com](http://travisjeffery.com)
- GitHub [@travisjeffery](https://github.com/travisjeffery)
- Twitter [@travisjeffery](https://twitter.com/travisjeffery)
- Medium [@travisjeffery](https://medium.com/@travisjeffery)


