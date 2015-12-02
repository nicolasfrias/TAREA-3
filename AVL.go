package AVL
import ("fmt")

type  nodo struct {
  valor int
  altura int
  Left *nodo
  Right *nodo
}
type AVL struct{
  raiz *nodo
}
func NewAVL() *AVL{
  return New(AVL)
}
func(n *nodo) RotacionIZ()*nodo{
  node:=n.Right
  n.Right=node.Left
  node.Left=n
  n.altura= n.max(n.left.Altura(),n.right.Altura())+1
  node.altura= n.max(n.left.Altura(),n.right.Altura())+1
  return node
}
func(n*nodo) RotacionIZDER() *nodo{
  n.Left = n.RotacionIZ()
  n=n.RostacionDER()
  return n
}
func (n*nodo) RotacionDER() *nodo{
  node:=n.Left
  n.Left=node.Right
  node.Right=n
  n.altura= n.max(n.left.Altura(),n.right.Altura())+1
  node.altura= n.max(n.left.Altura(),n.right.Altura())+1
  return node
}
func (n*nodo) RotacionDERIZ()*nodo{
  n.Right=n.RotacionDER()
  n=n.RotacionIZ()
  return n
}

func (n*nodo)Altura()int{
  if raiz != nil{
    return n.altura
  }
  return 0
}
func max(a,b int)int{
  if a>b{
    return a
  }
  return b
}
func (a*AVL)Insertar(x int) {
return a.root.insertar(x)
}
func (n*nodo) insertar(x int){
  
if n = nil{
  n=&nodo{x,0,nil,nil}
  n.altura=n.max(n.left.Altura(),n.right.Altura())+1
  return
  }
if x<n.valor{
      n.Left =n.insertar(x)
      if n.left.Altura()-n.right.Altura()==2{
        if x < n.Left.valor{
          n=RotacionDER()
        }else{
          n=RotacionIZDER()
        }
      }
    }
  if x>n.valor{
        n.Right =n.insertar(x)
        if n.right.Altura()-n.left.Altura()==2{
          if x < n.Right.valor{
            n=RotacionIz()
          }else{
            n=RotacionDERIZ()
          }
        }
}
n.altura= max(n.left.Altura(),n.right.Altura())+1
}
func (a*AVL)Buscar(x int)bool{
return a.raiz.buscar(x)  
}
func (n *nodo)buscar(x int)bool{
    if n=nil{
      false
    }
if  x> n.valor{
  n.Right.buscar(x)
}
if  x< n.valor{
  n.Left.buscar(x)
}
if  x == n.valor{
  return true
}
}
