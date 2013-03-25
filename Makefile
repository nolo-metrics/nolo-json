SRC=lex.go parse.go
TESTS=lex_test.go parse_test.go
ALL=$(SRC) $(TESTS)

test: $(ALL)
	go test

tags: $(ALL)
	ctags -R .

clean:
	rm tags
