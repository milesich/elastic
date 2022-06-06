// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"testing"
)

func TestIndicesGetMappingURL(t *testing.T) {
	client := setupTestClientAndCreateIndex(t)

	tests := []struct {
		Indices  []string
		Expected string
	}{
		{
			[]string{},
			"/_all/_mapping",
		},
		{
			[]string{"twitter"},
			"/twitter/_mapping",
		},
		{
			[]string{"store-1", "store-2"},
			"/store-1%2Cstore-2/_mapping",
		},
	}

	for _, test := range tests {
		path, _, err := client.GetMapping().Index(test.Indices...).buildURL()
		if err != nil {
			t.Fatal(err)
		}
		if path != test.Expected {
			t.Errorf("expected %q; got: %q", test.Expected, path)
		}
	}
}
