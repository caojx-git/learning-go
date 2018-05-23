/**
 * 使用测试包
 */
package even

import "testing"

func TestEven(t *testing.T) {
	if !Even(2) {
		t.Log("2 should be even")
		t.Fail()
	}
}

func Even(i int) bool {
	return i%2 == 0
}
