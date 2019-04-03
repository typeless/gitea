// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package integrations

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testRepoCommitsSearch(t *testing.T, query, commit string) {
	prepareTestEnv(t)

	session := loginUser(t, "user2")

	// Request repository commits page
	req := NewRequestf(t, "GET", "/user2/commits_search_test/commits/branch/master/search?q=%s", url.QueryEscape(query))
	resp := session.MakeRequest(t, req, http.StatusOK)

	doc := NewHTMLParser(t, resp.Body)
	sel := doc.doc.Find("#commits-table tbody tr td.sha a")
	assert.EqualValues(t, commit, strings.TrimSpace(sel.Text()))
}

func TestRepoCommitsSearch(t *testing.T) {
	testRepoCommitsSearch(t, "author:alice", "6e8eabd9a7")
	testRepoCommitsSearch(t, "committer:Tom", "58e97d1a24")
	testRepoCommitsSearch(t, "author:bob commit-4", "58e97d1a24")
	testRepoCommitsSearch(t, "author:bob commit after:2019-03-03", "58e97d1a24")
	testRepoCommitsSearch(t, "committer:alice commit before:2019-03-02", "6e8eabd9a7")
}
