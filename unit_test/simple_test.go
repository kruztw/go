// ref: https://bingdoal.github.io/backend/2022/05/unit-test-on-golang/

package main

import "testing"

func TestAdd(t *testing.T) {
	ans := add(1, 2)
	if ans != 3 {
		t.Errorf("Ans isn't correct. ans: %d", ans)
	}
}
