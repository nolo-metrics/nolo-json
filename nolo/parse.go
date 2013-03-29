package nolo

// fixme: option values like key=2.3.4
func Parse(name, input string) *Plugin {
	plugin := new(Plugin)
	plugin.Identifier = name

	l := Lex(name, input)

	current_option_id := ""

Loop:
	for {
		item := l.nextItem()
		switch item.typ {
		case itemIdentifier:
			metric := new(Metric)
			metric.Identifier = item.val
			plugin.Metrics = append(plugin.Metrics, *metric)
		case itemValue:
			metric := &plugin.Metrics[len(plugin.Metrics) - 1]
			metric.Value = item.val
		case itemOptionIdentifier:
			current_option_id = item.val
		case itemOptionValue:
			metric := &plugin.Metrics[len(plugin.Metrics) - 1]
			if metric.Metadata == nil {
				metric.Metadata = make(map[string]string)
			}
			metric.Metadata[current_option_id] = item.val
		case itemEOF, itemError:
			break Loop
		}
	}

	return plugin
}