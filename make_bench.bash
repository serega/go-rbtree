gotgo -o temp_intset.go --package-name=main sortedset.got int
sed -n '1h;1!H;${;g;s/func testTypes.*f(arg0).*}//g;p;}' temp_intset.go > intset.go
rm temp_intset.go
gotgo -o inttree.go --package-name=main tree.got int

6g -o bench.6 bench.go inttree.go intset.go
6l -o bench bench.6
