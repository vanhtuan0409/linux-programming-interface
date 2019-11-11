build:
	go build -o bin/atomic .

test_append:
	rm tmp.txt || true
	./bin/atomic tmp.txt 1000000 & ./bin/atomic tmp.txt 1000000
	cat tmp.txt | wc

test_no_append:
	rm tmp.txt || true
	./bin/atomic -x tmp.txt 1000000 & ./bin/atomic -x tmp.txt 1000000
	cat tmp.txt | wc

test_buffer:
	rm tmp.txt || true
	strace ./bin/atomic -b tmp.txt 1000000
	cat tmp.txt | wc
