package main

import (
	"fmt"
	"time"
)

var bandera int64

type Proceso struct {
	Id, i int64
}

type Procesos struct{
	SliceProcesos []Proceso
}

func (ps *Procesos) Agregar(p Proceso){
	ps.SliceProcesos = append(ps.SliceProcesos,p)


	for iterador:=0;iterador<len(ps.SliceProcesos);iterador=iterador+1{
		go ps.SliceProcesos[iterador].Hacer()
	}
}

func (ps *Procesos) Mostrar(){
	for i:=0;i < len(ps.SliceProcesos);i = i +1{
		go ps.SliceProcesos[i].Mostrar()
	}
}

func (p *Proceso) Hacer() {
	for {
		p.i = p.i + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Proceso) Mostrar() {
	for {
		fmt.Println(p.Id,": " ,p.i)
		time.Sleep(time.Millisecond * 500)

		if bandera == 2{
			return
		}
	}
}

func (ps *Procesos) Eliminar(b int64){
	var posicion int

	posicion = -1

	for i:=0;i < len(ps.SliceProcesos) && posicion == -1;i = i +1{
		if ps.SliceProcesos[i].Id == b{
			posicion = i
		}
	}

	if posicion != -1{
		if posicion == len(ps.SliceProcesos)-1{
			ps.SliceProcesos = append(ps.SliceProcesos[:posicion])
		} else {
			ps.SliceProcesos = append(ps.SliceProcesos[:posicion],ps.SliceProcesos[posicion+1:]...)
		}
	}
}

func main()  {
	var opcion, id, eliminar int64
	id = 0

	procesos:= Procesos{}

	for{
		fmt.Println("Menu \n1.- Agregar proceso\n2.- Mostrar procesos\n3.- Terminar proceso\n4.- Salir")
		fmt.Print("Opcion: ")
		fmt.Scanln(&opcion)

		if (opcion == 1){
			procesos.Agregar(Proceso{Id:id,i:0})
			id = id + 1

		} else if (opcion == 2){
			bandera = 0
			procesos.Mostrar()
			fmt.Scanln(&bandera)
		} else if (opcion == 3){
			fmt.Scanln(&eliminar)
			procesos.Eliminar(eliminar)
		} else if (opcion == 4){
			break
		}
	}
} 