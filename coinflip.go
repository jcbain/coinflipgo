package main

import (
        "fmt"
        "math/rand"
        "time"
        "reflect"
)

// simulate flipping a fair coin
func FlipCoin() string {
        rand.Seed(time.Now().UnixNano())

        if flipint := rand.Intn(2); flipint == 0 {
               return "tails"
        }
        return "heads"
}

// flip until condition is met
func FlipCondition() int {
        flips    := [2]string{FlipCoin(), FlipCoin()}
        previous := 0
        current  := 1
        s        := flips[previous: current+1]
        cond     := []string{"heads","heads"}
        numflips := 2

        for !reflect.DeepEqual(s, cond) {
                s[previous] = s[current]
                s[current] = FlipCoin()
                numflips +=1
        }
        return numflips
}

// mean function
func mean(vals[]int) float64 {
        tot := 0.0
        for _, v := range vals{
                tot += float64(v)
        }
        return tot/float64(len(vals))
}

// run multiple trials
func FlippityFlip(numtrials int) []int {
        vals := make([]int, numtrials)
        for i := 0; i < numtrials; i++ {
                vals[i] = FlipCondition()
        }
        return vals
}

func main() {
        trials := FlippityFlip(10000)
        fmt.Println(mean(trials))
}
