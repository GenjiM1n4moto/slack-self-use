compare: # Compare with last two versions
	- make run
	- diff -u `find ./deleted-users/* | sort -r | head -n 2 | sort`
run:
	go run .
