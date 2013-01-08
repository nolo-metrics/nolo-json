BINS=bin/nolo-gmetric
PLUGINS=plugins/load plugins/passenger
TESTS=test/gmetric/check-exec test/gmetric/empty-params test/gmetric/multiple test/gmetric/single test/plugins/load test/plugins/passenger
FAKE_PLUGINS=test/fake/plugins/failing-return test/fake/plugins/inline-metadata test/fake/plugins/multiple test/fake/plugins/not-executable test/fake/plugins/shared-metadata test/fake/plugins/single
FAKE_BINS=test/fake/bin/passenger-status test/fake/bin/passenger-memory-stats test/fake/bin/uptime
ALL=$(BINS) $(TESTS) $(PLUGINS) $(FAKE_PLUGINS) $(FAKE_BINS)

test: $(ALL)
	test/gmetric/check-exec
	test/gmetric/empty-params
	test/gmetric/multiple
	test/gmetric/single
	test/plugins/load
	test/plugins/passenger

tags: $(ALL)
	ctags -R .

clean:
	rm tags
