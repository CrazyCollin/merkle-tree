package merkletree

import "fmt"

type Node struct {
	Tree *MerkleTree

	Parent *Node
	Left   *Node
	Right  *Node

	Hash []byte

	isLeaf bool
	dup    bool

	C Content
}

// 遍历当前节点的子树，直到遇见叶子节点
func (n *Node) verifyNode() ([]byte, error) {
	if n.isLeaf {
		return n.C.CalculateHash()
	}
	rightBytes, err := n.Right.verifyNode()
	if err != nil {
		return nil, err
	}
	leftBytes, err := n.Left.verifyNode()
	if err != nil {
		return nil, err
	}

	strategy := n.Tree.hashStrategy()
	_, err = strategy.Write(append(leftBytes, rightBytes...))
	if err != nil {
		return nil, err
	}
	return strategy.Sum(nil), nil
}

func (n *Node) getNodeHash() ([]byte, error) {

	if n.isLeaf {
		return n.C.CalculateHash()
	}

	strategy := n.Tree.hashStrategy()

	_, err := strategy.Write(append(n.Left.Hash, n.Right.Hash...))
	if err != nil {
		return nil, err
	}
	return strategy.Sum(nil), nil
}

func (n *Node) String() string {
	return fmt.Sprintf("%")
}
