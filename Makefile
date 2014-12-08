TEST_COMMAND = go test -bench=. -covermode atomic -v

default: test

test:
	$(TEST_COMMAND)

watch:
	fswatch -o -r . | \
		while read; do \
			$(TEST_COMMAND) | \
				tail -n1 | \
				terminal-notifier -title 'Go test results' -group Go; \
		done