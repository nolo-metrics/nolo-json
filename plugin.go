package nolo

type Metric struct {
	Identifier string
	Value      string
	Metadata   map[string]string
}

func (m Metric) ToMap() map[string]string {
	result := make(map[string]string)
	for k, v := range m.Metadata {
		result[k] = v
	}
	result["value"] = m.Value
	result["identifier"] = m.Identifier
	return result
}

type Plugin struct {
	Identifier string
	Metrics    []Metric
}

func (p Plugin) ToMap() map[string][]map[string]string {
	result := make(map[string][]map[string]string)
	length := len(p.Metrics)
	values := make([]map[string]string, length)
	for i, m := range p.Metrics {
	 	values[i] = m.ToMap()
	}
	result[p.Identifier] = values
	return result
}