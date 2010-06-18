gotgo -o int_set.go --package-name=main sortedset.got int
gotgo -o int_tree.go --package-name=main tree.got int
6g -o bench.6 bench.go inttree.go intset.go
6l -o bench bench.6
