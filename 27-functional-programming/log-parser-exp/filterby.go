package main

type filterFunc func(result) (include bool)

func filterBy(results []result, filterer filterFunc) []result {
	out := results[:0]

	for _, r := range results {
		if !filterer(r) {
			continue
		}

		out = append(out, r)
	}

	return out
}
