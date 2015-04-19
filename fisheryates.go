package fisheryates

import (
    "math/rand"
    "time"
)

func init(){
  // set the seed
  rand.Seed(time.Now().UnixNano())
}

func Shuffle(list []interface{}){
  // get the length of the list
  n := len(list)

  // shuffle
  for i := n - 1; i > 0; i-- {
    j := rand.Intn(i + 1)
    list[i], list[j] = list[j], list[i]
  }
}

func ShuffleInt(list []int){
  // get the length of the list
  n := len(list)

  // shuffle
  for i := n - 1; i > 0; i-- {
    j := rand.Intn(i + 1)
    list[i], list[j] = list[j], list[i]
  }
}

func ShuffleString(list []string){
  // get the length of the list
  n := len(list)

  // shuffle
  for i := n - 1; i > 0; i-- {
    j := rand.Intn(i + 1)
    list[i], list[j] = list[j], list[i]
  }
}
