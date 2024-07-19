package utils

import (
	"regexp"
	"strings"
)

type IndexGroups map[string]string

// ParseIndexGroups
// The specification is a space-separated list of index entries. Normally, each entry is just a single character,
// but you may also specify multiple characters.  For instance, the entry "The" will link to all files and
// folders starting with "The".
//
// You may also create an entry using a group of index characters in parentheses. For instance, the entry
// "A-E(ABCDE)" will display as "A-E" and link to all files and folders starting with either
// A, B, C, D or E.  This may be useful for grouping less-frequently used characters (such and X, Y and Z), or
// for grouping accented characters (such as A, \u00C0 and \u00C1)
//
// Files and folders that are not covered by an index entry will be placed under the index entry "#".

var indexGroupsRx = regexp.MustCompile(`(.+)\((.+)\)`)

func ParseIndexGroups(spec string) IndexGroups {
	parsed := make(IndexGroups)
	split := strings.Split(spec, " ")
	for _, g := range split {
		sub := indexGroupsRx.FindStringSubmatch(g)
		if len(sub) > 0 {
			for _, c := range sub[2] {
				parsed[string(c)] = sub[1]
			}
		} else {
			parsed[g] = g
		}
	}
	return parsed
}
