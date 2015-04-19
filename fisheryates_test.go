package fisheryates

import "testing"

func TestFib(t *testing.T) {
  // initialize the list
  intlist := []int{1,2,3,4,5}
  stringlist := []string{"1","2","3","4","5"}

  // shuffle the list
  ShuffleInt(intlist)
  ShuffleString(stringlist)
}
