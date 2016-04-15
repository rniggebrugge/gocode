package main

import (
  "fmt"
)

func main(){
  pop := map[string]int{"tokyo":12, "amsterdam":1, "london":8, "paris":5,"lisbon":2}
  fmt.Println(len(pop), pop)
  delete(pop, "paris")
  fmt.Println(len(pop), pop)
  delete(pop, "rotterdam")
  fmt.Println(len(pop), pop)
  pop["tokyo"], pop["amsterdam"] = pop["amsterdam"], pop["tokyo"]
  fmt.Println(len(pop), pop)
  
  city := make(map[int]string, len(pop))
  for c, p := range pop {
    city[p]=c
  }
  fmt.Println("\n------- ----- ---\n",city)
}