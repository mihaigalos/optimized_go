# Dot Product

Comparison between pure go implementation (sequential, parallelized) and go+rust optimization.

Parallel implementaiton is 2.39x faster than the sequential one.

```
i5-4570T [1549.246, 3600,0000] Mhz

Sequencial, 50000000 elements, dot product
generating random maps..
19.572541915s
starting..
15.856622657s

Parallel, 4 cores, 50000000 elements, dot product
generating random maps..
19.614979378s
starting..
parallelizing..
6.633762924s
``` 