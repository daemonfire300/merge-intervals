# merge-intervals

## Purpose of this project

Learning stuff.

## FAQ

*Does this sample project implement streaming data?* No.

*How did we verify the "correctness" of this algorithm?* Run it through the test suite at LeetCode.

*O-Notation?* Depends on the sorting algorithm used, assuming the sort happens in `n * log(n)` the sorting dominates the time complexity, although there is some in-efficiency by using an inner for-loop (which has been removed, an earlier version contained in-efficiencies, changes were made in g#3f52a15).

## How to use

### Docker version

 1. `docker build -t mergectl`
 2. `cat sample.json | docker run -t mergectl`

There is also a public Docker Hub page for this project https://hub.docker.com/repository/docker/realjulius/merge-intervals.

### go run

Using stdin:

`cat sample.json | go run ./cmd/mergectl`

Using file:

`go run ./cmd/mergectl sample.json`
