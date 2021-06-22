# merge-intervals

## Purpose of this project

Learning stuff.

## FAQ

*Does this sample project implement streaming data?* No, not yet.

## How to use

### Docker version

 1. `docker build -t mergectl`
 2. `cat sample.json | docker run -t mergectl`

### go run

Using stdin:

`cat sample.json | go run ./cmd/mergectl`

Using file:

`go run ./cmd/mergectl sample.json`
