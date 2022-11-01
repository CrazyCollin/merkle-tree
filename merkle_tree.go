package merkletree

import "hash"

type Content interface {
	CalculateHash() ([]byte, error)
	Equals(another Content) (bool, error)
}

type MerkleTree struct {
	Root         *Node
	merkleRoot   []byte
	Leafs        []*Node
	hashStrategy func() hash.Hash
}
