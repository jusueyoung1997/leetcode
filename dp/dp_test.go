package dp

import "testing"

func TestNthUglyNumber(t *testing.T) {
	if nthUglyNumber(10) != 12 {
		t.Errorf("expect 12 but get %d", nthUglyNumber(10))
	}
}

func TestMincostTickets(t *testing.T) {
	if mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}) != 17 {
		t.Errorf("expect 17 but get %d", mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
	}

	if mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}) != 11 {
		t.Errorf("expect 11 but get %d", mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	}

}

func TestLongestValidParentheses(t *testing.T) {
	if longestValidParentheses(")()())") != 4 {
		t.Errorf("expect 4 but get %d", longestValidParentheses(")()())"))
	}

	if longestValidParentheses("") != 0 {
		t.Errorf("expect 0 but get %d", longestValidParentheses(""))
	}

	if longestValidParentheses("(()") != 2 {
		t.Errorf("expect 2 but get %d", longestValidParentheses("(()"))
	}

	if longestValidParentheses("(()))))))((()))))())))()))))))))))()") != 6 {
		t.Errorf("expect 2 but get %d", longestValidParentheses("(()))))))((()))))())))()))))))))))()"))
	}
}

func TestIsMatch(t *testing.T) {
	if !isMatch("asdfssssvb", "a.dfs*vb") {
		t.Errorf("expect true but false")
	}

	if !isMatch("asdfssssvb", "a.dfs*.*vb") {
		t.Errorf("expect true but false")
	}

	if !isMatch("asdfssssvb", "a.dfs*b*v.") {
		t.Errorf("expect true but false")
	}

	if !isMatch("asdfssssvb", "a.dfs*b*v.*") {
		t.Errorf("expect true but false")
	}

	if isMatch("asdfssssvb", "a.dfs*b*v*") {
		t.Errorf("expect false but true")
	}

	if !isMatch("aaaaaaa", "a*") {
		t.Errorf("expect true but false")
	}

	if !isMatch("ab", ".*") {
		t.Errorf("expect true but false")
	}

	if !isMatch("aab", "c*a*b") {
		t.Errorf("expect true but false")
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	if longestCommonSubsequence("abcde", "ace") != 3 {
		t.Errorf("expect 3, but get %d", longestCommonSubsequence("abcde", "ace"))
	}

	if longestCommonSubsequence("abcdefsdas", "acefds") != 6 {
		t.Errorf("expect 6, but get %d", longestCommonSubsequence("abcde", "ace"))
	}
}

func TestLongestPalindromeSubseq(t *testing.T) {
	if longestPalindromeSubseq("bbbab") != 4 {
		t.Errorf("expect 4, but get %d", longestPalindromeSubseq("bbbab"))
	}

	if longestPalindromeSubseq("bbbaaabb") != 7 {
		t.Errorf("expect 7, but get %d", longestPalindromeSubseq("bbbaaabb"))
	}
}

func TestLongestIncreasingPath(t *testing.T) {
	if longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}) != 4 {
		t.Errorf("expect 4, but get %d", longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
	}

	if longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}) != 4 {
		t.Errorf("expect 4, but get %d", longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}))
	}
}
