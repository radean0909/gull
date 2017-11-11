# gull

Convex Hull implementation in golang 😺

Not much too it really.

## Development

Install [dep](https://github.com/golang/dep)
```
brew install dep
```

Install dependencies
```
dep ensure
```

## Tests

Run the tests:
```
go test ./...
```

## Running the example

```
go run main.go
```

You should see a load of `pngs` in `results/`. They should look a little like this:

![](http://i66.tinypic.com/1z4bvnm.png)

The example is generating a two random sets of points, the first is in a circular distribution and the next is a small set of noise. Giving the effect above 👍
