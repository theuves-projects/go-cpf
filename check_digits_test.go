package cpf

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestReverseDigits(t *testing.T) {
  var (
    param    [9]int
    expected [9]int
    actual   [9]int
  )

  param = [9]int{1, 3, 2, 4, 6, 5, 7, 9, 8}
  expected = [9]int{8, 9, 7, 5, 6, 4, 2, 3, 1}
  actual = reverseDigits(param)

  assert.Equal(t, expected, actual, "should reverse digits")
}

func TestGetCheckDigits(t *testing.T) {
  var (
    param    [9]int
    expected [2]int
    actual   [2]int
  )

  param = [9]int{1, 1, 1, 4, 4, 4, 7, 7, 7}
  expected = [2]int{3, 5}
  actual = getCheckDigits(param)
  assert.Equal(t, expected, actual, "should be 3 and 5")

  param = [9]int{2, 1, 8, 4, 2, 5, 9, 8, 5}
  expected = getCheckDigits(param)
  actual = [2]int{3, 8}
  assert.Equal(t, expected, actual, "should be 3 and 8")
}
