// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"strings"

	"giautm.dev/awesome/ent/schema/pulid"
	"giautm.dev/awesome/ent/todo"
)

// prefixMap maps PULID prefixes to table names.
var prefixMap = map[pulid.ID]string{
	"TD": todo.Table,
}

// WithPrefixedULID sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithPrefixedULID() NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = func(ctx context.Context, id pulid.ID) (string, error) {
			idx := strings.IndexRune(string(id), '_')
			if idx == -1 {
				return "", fmt.Errorf("pulid: incorrect id format")
			}

			prefix := id[:idx]
			typ := prefixMap[prefix]
			if typ == "" {
				return "", fmt.Errorf("pulid: could not map prefix '%s' to a type", prefix)
			}
			return typ, nil
		}
	}
}
