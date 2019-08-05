package main

type groupFunc func(result) (key string)

func groupBy(results []result, keyer groupFunc) []result {
	grouped := make(map[string]result, len(results))

	for _, cur := range results {
		key := keyer(cur)
		grouped[key] = cur.add(grouped[key])
	}

	out := results[:0]
	for _, r := range grouped {
		out = append(out, r)
	}

	return out
}
