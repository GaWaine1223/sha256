package sha256

import (
	"fmt"
	"testing"
)

func TestHashwithDifficulty(t *testing.T) {
	sum, nonce := HashwithDifficulty([]byte("hello world"), 3)
	fmt.Println("nonce = ", nonce)
	fmt.Printf("%x", sum)
	// Output: a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447
}
