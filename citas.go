package main

import (
	"fmt"
	"time"
	)

type Node struct {

  nombre1 string
	nombre2 string
	apellido1 string
	apellido2 string
	caracterid string
	numeroid int
	horallegada int
	minutollegada int
	sintomas string
	eps string

	}



func (n *Node) String() string {



	return fmt.Sprint("\nNombre: ", n.nombre1, " " ,n.nombre2, " " , n.apellido1, " " ,n.apellido2, "\nIdentificacion: " ,n.caracterid, n.numeroid, "\nHora llegada: " , n.horallegada, ":" , n.minutollegada, "\nSintomas: ", n.sintomas, "\neps: ", n.eps)


}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

// Push adds a node to the queue.
func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}





func main() {



  var nombre1 string
	var nombre2 string
	var nombre3 string
	var nombre4 string
	var codigoletra string
	var codigonuemro int
	var sintomas string
	var eps string
	var cola string
	var epscod [5]int
	var contador=0
	var codigoeps=0


	epscod[0]=0
	epscod[1]=0
	epscod[2]=0
	epscod[3]=0
	epscod[4]=0


		i := 1

		s := NewQueue(1)

	     for i <= 2 {



			fmt.Println("ingresa nobre: ")

	    fmt.Scanf("%v %v %v %v %v",&nombre1,&nombre2,&nombre3,&nombre4)

			fmt.Println("ingrese codigo: ")

			fmt.Scanf("%v %v %v",&codigoletra,&codigonuemro)


			fmt.Println("ingrese sintomas: ")

			fmt.Scanf("%v %v",&sintomas)

		codigoeps=0

		for codigoeps > 5 || codigoeps < 1{

			fmt.Println("ingrese cododigo de Eps a la cual pertenece: \n1. compensar\n2. sanitas\n3. sura\n4. salud total\n5. coomeva\n")
			fmt.Scanf("%v %v",&codigoeps)


		}

  	eps=""

		switch (codigoeps) {
    case 1:
				epscod[0]=epscod[0]+1
				eps="compensar"
    case 2:
				epscod[1]=epscod[1]+1
				eps="sanitas"
    case 3:
				epscod[2]=epscod[2]+1
				eps="sura"
		case 4:
				epscod[3]=epscod[3]+1
				eps="salud total"
		case 5:
				epscod[4]=epscod[4]+1
		 		eps="coomeva"
    }

 			fmt.Println("\n\nRegistro eps\tPersonas atendidas \n \n1. compensar\t",epscod[0],"\n2. sanitas\t",epscod[1],"\n3. sura\t\t",epscod[2],"\n4. salud total\t",epscod[3],"\n5. coomeva\t",epscod[4],"\n")


	    fmt.Println("\ncola\n")

		now := time.Now()

   	s.Push(&Node{nombre1,nombre2,nombre3,nombre4,codigoletra,codigonuemro,now.Hour(),now.Minute(),sintomas,eps})
		contador=contador+1;


		fmt.Println("desea ingresar un nuevo Paciente S/N: ")



		fmt.Scanf("%v %v\n",&cola)

		if cola=="n" || cola=="N"{

			i=4

			for  contador > 0 {

				fmt.Println(s.Pop())
				contador = contador-1

		}

	}

}

fmt.Println("\n\nRegistro eps\tPersonas atendidas \n \n1. compensar\t",epscod[0],"\n2. sanitas\t",epscod[1],"\n3. sura\t\t",epscod[2],"\n4. salud total\t",epscod[3],"\n5. coomeva\t",epscod[4],"\n")

}
