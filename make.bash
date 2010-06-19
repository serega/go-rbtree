gotgo -o tree.go --package-name=rbtree tree.got interface{}
gotgo -o temp_sortedset.go --package-name=rbtree sortedset.got interface{}
sed -n '1h;1!H;${;g;s/func testTypes.*f(arg0).*}//g;p;}' temp_sortedset.go > sortedset.go
rm temp_sortedset.go
make