package main

import "time"
import "fmt"
import "math/rand"

import . "algebra"


func printFeedFroward(title string, neuralNetwork NeuralNetwork, allInput []Vector) {

   fmt.Println(title)

   for _, input := range allInput {
      output := neuralNetwork.FeedForward(input)
      fmt.Println(input.ToString() + " => " + output.ToString())
   }
}

func main() {

   rand.Seed(time.Now().UnixNano())

   v := func(i ...float64) Vector { return CreateVector(i) }

   input := []Vector {
      v(0, 0, 1),
      v(1, 0, 1),
      v(0, 1, 1),
      v(1, 1, 1)}

   output := []Vector {
      v(0, 1),
      v(1, 0),
      v(1, 0),
      v(0, 1)}

   neuralNetwork := CreateNeuralNetwork(3, 5, 2)

   printFeedFroward("Before Training", neuralNetwork, input)

   for i := 0; i < 10000; i++ {
      r := rand.Intn(len(input))
      neuralNetwork.Train(input[r], output[r])
   }

   printFeedFroward("After Training", neuralNetwork, input)
}



