BINS=bin/nolo-gmetric
PLUGINS=plugins/load plugins/test/failing-return plugins/test/inline-metadata plugins/test/multiple plugins/test/not-executable plugins/test/shared-metadata plugins/test/single
TESTS=test/gmetric/empty-params test/gmetric/check-exec test/gmetric/single test/gmetric/multiple
ALL=$(BINS) $(TESTS) $(PLUGINS)

test: $(ALL)
	test/gmetric/empty-params
	test/gmetric/check-exec
	test/gmetric/single
	test/gmetric/multiple

tags: $(ALL)
	ctags -R .
