package particles

import (
	"log"
	"math"
	"math/rand"
)

type Coffee struct {
  ParticleSystem
}

var dirs = [][]int{
  {-1, -1},
  {-1, 0},
  {-1, 1},

  {0, -1},
  {0, 1},

  {1, 0},
  {1, 1},
  {1, -1},
}


func countParticles(row, col int, counts[][]int) int {
  count := 0

  for _, dir := range dirs {
    r := row + dir[0]
    c := col + dir[1]
    if r < 0 || r >= len(counts) || c < 0 || c >= len(counts[0]) {
      continue
    }
    count = counts[row + dir[0]][col+dir[1]]
  }

  return count
}

func normalize(row, col int, counts[][]int) {

  if countParticles(row, col, counts) > 4 {
    counts[row][col] = 0
  }
}

func reset(p *Particle, params *ParticleParams) {
  p.Lifetime = int64(math.Floor(float64(params.MaxLife) * rand.Float64()))
  p.Speed = params.MaxSpeed * rand.Float64()

  maxX := math.Floor(float64(params.X) / 2)
  x := math.Max(-maxX, math.Min(rand.NormFloat64() * params.Scale, maxX))

  p.X = x + maxX
  p.Y = 0
}

func nextPos(p *Particle, deltaMs int64) {
  p.Lifetime -= deltaMs

  if p.Lifetime <= 0 {
    return
  }

  percent := (float64(deltaMs) / 2000.0)
  p.Y += p.Speed * percent
}

func NewCoffee(width, height int, scale float64) Coffee {
  if width % 2 == 0 {
    log.Fatal("width must be odd number")
  }

  ascii := func(row, col int, counts[][]int) string {
    count := counts[row][col]

    if count == 0 {
       return " "
    }
    if count < 4 {
      return "░"
    }
    if count < 6 {
       return "▒"
    }
    if count < 9 {
       return "▓"
    }

    return "█"
  }


  return Coffee{
    ParticleSystem: NewParticleSystem(
      ParticleParams{
        MaxLife: 7000,
        MaxSpeed: 1.5,
        ParticleCount: 60,

        reset: reset,
        ascii: ascii,
        nextPosition: nextPos,
        X: width,
        Y: height,
        Scale: scale,
      },
    ),
  }
}
