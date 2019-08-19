// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

type notifier interface {
	notify(func(r result))
}

type parser interface {
	parse() error
	notifier
}

type iterator interface {
	each(func(result))
}

type analyser interface {
	analyse(result)
	iterator
}

type summarizer interface {
	summarize(iterator) error
}

func report(a analyser, p parser, s summarizer) error {
	p.notify(a.analyse)

	if err := p.parse(); err != nil {
		return err
	}

	return s.summarize(a)
}
