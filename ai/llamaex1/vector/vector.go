package vector

// Package vector provides functions for learning about vectors and vector
// databases.
// This code was taken from:
// https://github.com/kabychow/go-cosinesimilarity
// https://github.com/quartercastle/vector
// https://github.com/gonum/gonum

import "math"

// Data represents data that can be vectorized.
type Data interface {
	Vector() []float32
}

// =============================================================================

// SimilarityResult represents the result of performaing a similarity check
// between two embeddings.
type SimilarityResult struct {
	Target     Data
	DataPoint  Data
	Similarity float32
	Percentage float32
}

// Similarity calculates the similarity between two vectors.
func Similarity(target Data, dataPoints ...Data) []SimilarityResult {
	results := make([]SimilarityResult, len(dataPoints))

	te := target.Vector()

	for i, dp := range dataPoints {
		similarity := CosineSimilarity(te, dp.Vector())

		results[i] = SimilarityResult{
			Target:     target,
			DataPoint:  dp,
			Similarity: similarity,
			Percentage: similarity * 100,
		}
	}

	return results
}

// CosineSimilarity takes two vectors and computes the similarity between
// them using a cosine algorithm.
func CosineSimilarity(x, y []float32) float32 {
	var sum, s1, s2 float64

	for i := 0; i < len(x); i++ {
		sum += float64(x[i] * y[i])
		s1 += float64(x[i] * x[i])
		s2 += float64(y[i] * y[i])
	}

	if s1 == 0 || s2 == 0 {
		return 0.0
	}

	return float32(sum / (math.Sqrt(s1) * math.Sqrt(s2)))
}

// =============================================================================

const (
	x = iota
	y
	z
)

// Add calculates the addition of two vectors.
func Add(a, b []float32) []float32 {
	dimA, dimB := len(a), len(b)

	if (dimA == 1 || dimA == 2 || dimA == 3) && dimB == 1 {
		a[x] += b[x]
		return a
	}

	if dimA == 2 && dimB == 2 {
		a[x], a[y] = a[x]+b[x], a[y]+b[y]
		return a
	}

	if dimA == 3 && dimB == 2 {
		a[x], a[y] = a[x]+b[x], a[y]+b[y]
		return a
	}

	if dimA == 3 && dimB == 3 {
		a[x], a[y], a[z] = a[x]+b[x], a[y]+b[y], a[z]+b[z]
		return a
	}

	if dimB > dimA {
		axpyUnitaryTo(a, 1, a, b[:dimA])
	} else {
		axpyUnitaryTo(a, 1, a, b)
	}

	return a
}

// Sub calculates the subtraction of two vectors.
func Sub(a, b []float32) []float32 {
	dimA, dimB := len(a), len(b)

	if (dimA == 1 || dimA == 2 || dimA == 3) && dimB == 1 {
		a[x] -= b[x]
		return a
	}

	if dimA == 2 && dimB == 2 {
		a[x], a[y] = a[x]-b[x], a[y]-b[y]
		return a
	}

	if dimA == 3 && dimB == 1 {
		a[x] -= b[x]
		return a
	}

	if dimA == 3 && dimB == 2 {
		a[x], a[y] = a[x]-b[x], a[y]-b[y]
		return a
	}

	if dimA == 3 && dimB == 3 {
		a[x], a[y], a[z] = a[x]-b[x], a[y]-b[y], a[z]-b[z]
		return a
	}

	if dimB > dimA {
		axpyUnitaryTo(a, -1, b[:dimA], a)
	} else {
		axpyUnitaryTo(a, -1, b, a)
	}

	return a
}

func axpyUnitaryTo(dst []float32, alpha float32, x, y []float32) {
	dim := len(y)
	for i, v := range x {
		if i == dim {
			return
		}
		dst[i] = alpha*v + y[i]
	}
}
