// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

type byPrice list

func (l byPrice) Less(i, j int) bool { return l[i].sum() < l[j].sum() }
func (l byPrice) Len() int           { return len(l) }
func (l byPrice) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

type byName list

func (l byName) Less(i, j int) bool { return l[i].String() < l[j].String() }
func (l byName) Len() int           { return len(l) }
func (l byName) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
