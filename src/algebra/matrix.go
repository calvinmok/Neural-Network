package algebra

import "fmt"


type Matrix struct {
   numRow int
   numColumn int
   array []float64
}

func CreateMatrix(numRow, numColumn int) Matrix {
   return Matrix { 
      numRow: numRow, 
      numColumn: numColumn,
      array: make([]float64, numRow * numColumn)}
}

func (me Matrix) Init(callback func(int, int) float64) Matrix {
   for r := 0; r < me.numRow; r++ {
      for c := 0; c < me.numColumn; c++ {
         me.Set(r, c, callback(r, c))
      }
   }
   return me
}

func (me Matrix) Get(r, c int) float64 {
   return me.array[r * me.numColumn + c]
}

func (me Matrix) Set(r, c int, value float64) {
   me.array[r * me.numColumn + c] = value
}

func (me Matrix) Row(r int) Vector {
   start := r * me.numColumn
   end := (r + 1) * me.numColumn
   return CreateVector(me.array[start:end])
}

func (me Matrix) CellAdd(r, c int, value float64) {
   me.Set(r, c, me.Get(r, c) + value)
}

func (me Matrix) Dot(vector Vector) Vector {
   result := CreateVectorWithSize(me.numRow)
   
   for r := 0; r < me.numRow; r++ {
      sum := 0.0
      for c := 0; c < me.numColumn; c++ {
         sum += me.Get(r, c) * vector.Get(c)
      }
      result.Set(r, sum)
   }

   return result
}

func (me Matrix) ToString() string {
   result := "["
   for r := 0; r < me.numRow; r++ {
      if r > 0 { result += "|" }
      for c := 0; c < me.numColumn; c++ {
         if c > 0 { result += "," }
         result += fmt.Sprintf("%.4f", me.Get(r, c))
      }
   }
   return result + "]"
}





