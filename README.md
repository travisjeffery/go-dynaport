# go-dynaport

A Go library for getting free ports, dynamically. 

## Example

``` go
ports := dynaport.MustGet(3)

jockoPort := ports[0]
serfPort := ports[1]
raftPort := ports[2]
```

Or

``` go
ports, err := dynaport.Get(3)
if err != nil {
  // handle err
}

// use ports...
```

## License

MIT

--- 

- [travisjeffery.com](http://travisjeffery.com)
- GitHub [@travisjeffery](https://github.com/travisjeffery)
- Twitter [@travisjeffery](https://twitter.com/travisjeffery)
- Medium [@travisjeffery](https://medium.com/@travisjeffery)


