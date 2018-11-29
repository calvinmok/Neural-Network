package main

import "math"
import "math/rand"

import . "algebra"


func Sigmoid(v float64) float64 {
   return 1.0 / (1.0 + math.Exp(-v))
}

func SigmoidDerivative(v float64) float64 {
   return v * (1.0 - v)
}


type NeuralNetwork struct {
   numInput int
   numHidden int
   numOutput int
   inputToHidden Matrix
   hiddenToOutput Matrix
}

func initMatrixCell(_, _ int) float64 { return rand.Float64() }

func CreateNeuralNetwork(input, hidden, output int) NeuralNetwork {
   result := NeuralNetwork { numInput: input, numHidden: hidden, numOutput: output }
   result.inputToHidden = CreateMatrix(hidden, input).Init(initMatrixCell)
   result.hiddenToOutput = CreateMatrix(output, hidden).Init(initMatrixCell)
   return result
}

func (me NeuralNetwork) FeedForward(input Vector) Vector {
   hidden := input.Dot(me.inputToHidden).Map(Sigmoid)
   return hidden.Dot(me.hiddenToOutput).Map(Sigmoid)
}

func (me NeuralNetwork) Train(input Vector, expect Vector) {

   hidden := input.Dot(me.inputToHidden).Map(Sigmoid)
   output := hidden.Dot(me.hiddenToOutput).Map(Sigmoid)

   difference := expect.Subtract(output).MultipleF64(2).Multiple(output.Map(SigmoidDerivative))

   for h := 0; h < me.numHidden; h++ {
      d := 0.0
      for o := 0; o < me.numOutput; o++ {
         d += difference.Get(o) * me.hiddenToOutput.Get(o, h)
      }

      d *= SigmoidDerivative(hidden.Get(h))

      for i := 0; i < me.numInput; i++ {
         me.inputToHidden.CellAdd(h, i, d * input.Get(i))
      }
   }

   for o := 0; o < me.numOutput; o++ {
      d := difference.Get(o)

      for h := 0; h < me.numHidden; h++ {
         me.hiddenToOutput.CellAdd(o, h, d * hidden.Get(h))
      }
   }
}



