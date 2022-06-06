// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"testing"
)

func TestIndicesGetFieldMappingURL(t *testing.T) {
	client := setupTestClientAndCreateIndex(t)

	tests := []struct {
		Indices  []string
		Fields   []string
		Expected string
	}{
		{
			[]string{},
			[]string{},
			"/_all/_mapping/field/%2A",
		},
		{
			[]string{},
			[]string{"message"},
			"/_all/_mapping/field/message",
		},
		{
			[]string{"twitter"},
			[]string{"*.id"},
			"/twitter/_mapping/field/%2A.id",
		},
		{
			[]string{"store-1", "store-2"},
			[]string{"message", "*.id"},
			"/store-1%2Cstore-2/_mapping/field/message%2C%2A.id",
		},
	}

	for _, test := range tests {
		path, _, err := client.GetFieldMapping().Index(test.Indices...).Field(test.Fields...).buildURL()
		if err != nil {
			t.Fatal(err)
		}
		if path != test.Expected {
			t.Errorf("expected %q; got: %q", test.Expected, path)
		}
	}
}
