
# Go Map
A simple implementation of a hashset.
Written using generics.

## Usage
```go
mapTest := new(lib.Set[string]).New()
mapTest.Add("foo bar")
mapTest.Add("bar")
mapTest.Remove("bar")
mapTest.In("bar")//False
```

## Testing
```
$ cd lib
$ go test
```