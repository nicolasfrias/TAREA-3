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
func(a *AVL) RotacionIZ()*nodo{
  node:=a.Right
  raiz.Right=nodo.Left
  nodo.Left=raiz
  raiz.altura= a.max(a.AlturaIz(),a.AlturaDer())+1
  node.altura= a.max(a.AlturaIz(),a.AlturaDer())+1
  return node
}
func(a*AVL) RotacionIZDER() *nodo{
  raiz.Left = a.RotacionIZ()
  raiz=a.RostacionDER()
  return raiz
}
func (a *AVL) RotacionDER() *nodo{
  node:=raiz.Left
  raiz.Left=nodo.Right
  node.Right=raiz
  raiz.altura= a.max(a.AlturaIz(),a.AlturaDer())+1
  node.altura= a.max(a.AlturaIz(),a.AlturaDer())+1
  return node
}
func (a*AVL) RotacionDERIZ()*nodo{
  raiz.Right=a.RostacionDER()
  raiz=a.RotacionIZ()
  return raiz
}
func (a*AVL)AlturaDer()int{
  if raiz != nil{
    return raiz.Right.altura
  }
  return 0
}
func (a*AVL)AlturaIz()int{
  if raiz != nil{
    return raiz.Lefth.altura
  }
  return 0
}
func max(a,b int)int{
  if a>b{
    return a
  }
  return b
}
func (a*AVL)insertar(x int) {
if raiz = nil{
  raiz=&AVL{x,0,nil,nil}
  raiz.altura=max(a.AlturaIz(),a.AlturaDer())+1
  return
}
if x<raiz.valor{
      raiz.Left =a.insertar(x)
      if a.AlturaIz()-a.AlturaDer()==2{
        if x < raiz.Left.valor{
          raiz=RostacionDER()
        }else{
          raiz=RotacionIZDER()
        }
      }
    }
  if x>raiz.valor{
        raiz.Right =a.insertar(x)
        if a.AlturaDer()-a.AlturaIz()==2{
          if x < raiz.Right.valor{
            raiz=RotacionIz()
          }else{
            raiz=RotacionDERIZ()
          }
        }
}
raiz.altura= max(a.AlturaIz(),a.AlturaDer())+1
}
func (a*AVL)buscar(x int)bool{
    if raiz=nil{
      false
    }
if  x> raiz.valor{
  a.Right.buscar(x)
}
if  x< raiz.valor{
  a.Left.buscar(x)
}
if  x == raiz.valor{
  return true
}
}
