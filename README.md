# merge-intervals

## Purpose of this project

Learning stuff.

## FAQ

**Does this sample project implement streaming data?** No.

---

**How did we verify the "correctness" of this algorithm?** Run it through the test suite at LeetCode.

---

**O-Notation / Time-complexity?** Depends on the sorting algorithm used, assuming the sort happens in `O(n * log(n))` the sorting dominates the time complexity, although there is some in-efficiency by using an inner for-loop (which has been removed, an earlier version contained in-efficiencies, changes were made in g#3f52a15).

---

**O-Notation / Space-complexity?** Depends on the sorting. In-place `O(log n)`, the result slice, however, contains `n` entries in the worst case, therefore `O(n+n)` for input slice plus result slice, which can be simplified to `O(n)`

---

**Streaming/Large Datasets?** It is impossible to apply this exact solution to an unordered stream of data. It would, however, be possible to process chunks of data. This would include (1) receiving a chunk of intervals, (2) sorting, (3) merging, (4) storing intermediary result. Then any next chunk would need to be appended to the already computed chunk, and the process repeated (2) sorting, (3) merging, (4) storing.

Alternatively or additionally one can sort the new chunk before appending and check if the lowest interval has a higher first value than the highest interval second value in the already processed chunk. Then there is no need to append and sort them, one can simply process the new chunk in isolation and append it to the already processed interval afterwards.



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
