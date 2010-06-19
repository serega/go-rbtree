package rbtree


import "testing"
import "fmt"
import "container/vector"
import "sort"
import "rand"


func NewIntTree() *RBTree {
    intComp := func(i interface{}, j interface{}) int {
        return i.(int) - j.(int);
    }
    
    return NewTree(intComp)
}


func MakeIntTree() *RBTree {
    tree := NewIntTree()
    tree.Insert(5)
    tree.Insert(10)
    tree.Insert(7)
    tree.Insert(0)
    tree.Insert(3)
    tree.Insert(20)
    tree.Insert(15)
    tree.Insert(2)
    return tree
}

func verifyElements(tree *RBTree, a []int, t *testing.T) {
    c := make([]int, len(a))
    copy(c, a)
    sort.SortInts(c)
    
    index := 0
    tree.For(func(e interface{}) bool {
        i := e.(int)
        if i != c[index] {
            t.Errorf("Expected %d, Actual %d", c[index], i)
            return false
        }
        index++
        return true
    });
}


func verifyTreeWithRandomData(a []int, t *testing.T) {
    s := NewIntTree();
    for i,v := range a {
        s.Insert(v)
        VerifyRBTreeProperties(s, t)
        if s.Size() != i + 1 {
             t.Errorf("Expected size %d, actual %d", i + 1, s.Size())   
        }
        if !s.Contains(v) {
            t.Errorf("Contains %d returned false", v)
        }
        s.Insert(v)
        if s.Size() != i + 1 {
            t.Errorf("Expected size %d, actual %d", i + 1, s.Size())   
        }
        
        verifyElements(s, a[0:i+1], t)
    }
    
   
    for i,v := range a {
        s.Remove(v)
        VerifyRBTreeProperties(s, t)
        if s.Size() != len(a) - (i + 1) {
             t.Errorf("Expected size %d, actual %d", len(a) - (i + 1), s.Size())   
        }
        if s.Contains(v) {
            t.Errorf("Contains %d returned true", v)
        }
        verifyElements(s, a[i + 1:], t)
    }
}


/**
    From Introduction to Algorithms, 2/e
    Red-Black Tree has the following properties
   1. Every Node is either red or black
   2. The root is black
   3. Every leaf (NIL) is black
   4. If a node is red, then both its children are black
   5. For each node, all paths from the node to descendantat leaves
      contain the same number of black nodes
*/
func VerifyRBTreeProperties(tree *RBTree, t *testing.T) {
    if tree.root == NIL {
        return
    }
    if tree.root.color != BLACK {
        t.Errorf("(1) root color is not BLACK")
    }
    leafs := new(vector.Vector)
    inorderTreeWalk(tree.root, func(n *Node) { 
        if n.Leaf() {
            leafs.Push(n);
        }
    });
    var numBlacks int = -1
    leafs.Do(func(i interface{}) {
        n := i.(*Node)
        if n.Left != NIL { // NIL is always black
            t.Errorf("left of leaf is not NIL")
        }
        if n.Right != NIL { // NIL is always black
            t.Errorf("right of leaf is not NIL")
        }
        
        var blacks int = 0
        for n != tree.root {
            if n.color == BLACK {
                blacks += 1
            } else {
                if n.Parent.color == RED { // property 4 failed
                    t.Errorf("(4) two consecutive RED nodes %d %d", n.Value.(int), n.Parent.Value.(int))
                }
            }
            n = n.Parent
        }
        if numBlacks == -1 {
            numBlacks = blacks
        } else {
            if numBlacks != blacks {
                t.Errorf("(5) number of blacks differ. %d vs. %d", numBlacks, blacks)
            }
        }
    });
}


func TestTree(t *testing.T) {
    InitRBTree()
    tree := MakeIntTree()
    VerifyRBTreeProperties(tree, t)
    
    tree.Foreach(func(elem interface{}) { fmt.Printf("%d, ", elem.(int))})
    fmt.Printf("\n")
    
    min := tree.First().(int)
    if min != 0 {
        t.Errorf("Expected 0 actual %d", min)
    }
    
    max := tree.Last().(int)
    if max != 20 {
        t.Errorf("Expected 20 actual %d", max)
    }    
    
    size := tree.Size()
    if size != 8 {
        t.Errorf("Size(). Expected 8 actual %d", size)
    }
    
    
    has_3 := tree.Contains(3)
    if !has_3 {
        t.Errorf("Contains(3) returned false")
    }
    
    tree.Remove(3)
    has_3 = tree.Contains(3)
    if has_3 {
        t.Errorf("Contains(3) returned true")
    }
}


func TestEquals(t *testing.T) {
    InitRBTree()
    a := rand.Perm(1000)
    tree1 := NewIntTree()
    tree2 := NewIntTree()
    for _,i := range a {
        tree1.Insert(i)
        tree2.Insert(i)
    }
    if !tree1.equals(tree2) {
        t.Errorf("equals returned false")
    }
    
    tree2.Remove(tree2.Last())
    if  tree1.equals(tree2) {
        t.Errorf("equals returned true")
    }
}

func TestRandomTree(t *testing.T) {
    InitRBTree()
    a := rand.Perm(1000)
    verifyTreeWithRandomData(a, t)
}

