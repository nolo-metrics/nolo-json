BINS=bin/nolo-gmetric
PLUGINS=plugins/load
TESTS=test/gmetric/check-exec test/gmetric/empty-params test/gmetric/multiple test/gmetric/single
FAKE_PLUGINS=test/fake/plugins/failing-return test/fake/plugins/inline-metadata test/fake/plugins/multiple test/fake/plugins/not-executable test/fake/plugins/shared-metadata test/fake/plugins/single
ALL=$(BINS) $(TESTS) $(PLUGINS) $(FAKE_PLUGINS)

test: $(ALL)
	test/gmetric/check-exec
	test/gmetric/empty-params
	test/gmetric/multiple
	test/gmetric/single

tags: $(ALL)
	ctags -R .

clean:
	rm tags
