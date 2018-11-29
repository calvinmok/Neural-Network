package algebra

import "fmt"


type Vector struct {
   array []float64
}

func CreateVectorWithSize(size int) Vector {
   return Vector { array: make([]float64, size) }
}

func CreateVector(allValue []float64) Vector {
   result := CreateVectorWithSize(len(allValue))
   for i, value := range allValue { result.array[i] = value }
   return result
}

func (me Vector) Clone() Vector {
   return CreateVector(me.array)
}

func (me Vector) Size() int { 
   return len(me.array)
}

func (me Vector) Get(i int) float64 {
   return me.array[i]
}

func (me Vector) Set(i int, value float64) {
   me.array[i] = value
}

func (me Vector) Map(callback func(float64) float64) Vector {
   result := me.Clone()
   for i := 0; i < me.Size(); i++ {
      result.Set(i, callback(result.Get(i)))
   }
   return result
}

func (me Vector) Subtract(vector Vector) Vector {
   result := me.Clone()
   for i := 0; i < me.Size(); i++ {
      result.Set(i, me.Get(i) - vector.Get(i))
   }
   return result
}

func (me Vector) Multiple(vector Vector) Vector {
   result := me.Clone()
   for i := 0; i < me.Size(); i++ {
      result.Set(i, me.Get(i) * vector.Get(i))
   }
   return result
}

func (me Vector) MultipleF64(value float64) Vector {
   result := me.Clone()
   for i := 0; i < me.Size(); i++ {
      result.Set(i, me.Get(i) * value)
   }
   return result
}

func (me Vector) DotV(vector Vector) float64 {
   result := 0.0
   for i := 0; i < me.Size(); i++ {
      result += me.Get(i) * vector.Get(i)
   }
   return result
}

func (me Vector) Dot(matrix Matrix) Vector {
   result := CreateVectorWithSize(matrix.numRow)

   for r := 0; r < matrix.numRow; r++ {
      sum := 0.0
      for c := 0; c < matrix.numColumn; c++ {
         sum += matrix.Get(r, c) * me.Get(c)
      }
      result.Set(r, sum)
   }

   return result
}

func (me Vector) ToString() string {
   result := "["
   for i := 0; i < me.Size(); i++ {
      if i > 0 { result += "," }
      result += fmt.Sprintf("%.4f", me.Get(i))
   }
   return result + "]"
}






