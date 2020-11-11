package main

import (
	"fmt"
	"time"
)

var parar int64

type Procesos struct{
	SliceProcesos []Proceso
}

func (ps *Procesos) AgregarProceso(p Proceso){
	ps.SliceProcesos = append(ps.SliceProcesos,p)


	for iterador:=0;iterador<len(ps.SliceProcesos);iterador=iterador+1{
		go ps.SliceProcesos[iterador].HacerProceso()
	}
}

func (ps *Procesos) MostrarProcesos(){
	for i:=0;i < len(ps.SliceProcesos);i = i +1{
		go ps.SliceProcesos[i].MostrarProceso()
	}
}


func (ps *Procesos) EliminarProceso(b int64){
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


type Proceso struct {
	Id int64
	I int64
}

func (p *Proceso) HacerProceso() {
	for {
		p.I = p.I + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Proceso) MostrarProceso() {
	for {
		fmt.Println(p.Id,": " ,p.I)
		time.Sleep(time.Millisecond * 500)

		if parar == 2{
			return
		}
	}
}


func main()  {
	var opcion int64
	var id int64
	var datoEliminar int64
	id = 0

	procesos:= Procesos{}

	for{
		
		fmt.Println("Menu \n1.- Agregar proceso\n2.- Mostrar procesos\n3.- Terminar proceso\n4.- Salir")
		fmt.Print("Opcion: ")
		fmt.Scanln(&opcion)

		if (opcion == 1){
			procesos.AgregarProceso(Proceso{Id:id,I:0})
			id = id + 1

		} else if (opcion == 2){
			parar = 0
			procesos.MostrarProcesos()
			fmt.Scanln(&parar)
		} else if (opcion == 3){
			fmt.Scanln(&datoEliminar)
			procesos.EliminarProceso(datoEliminar)
		} else if (opcion == 4){
			break
		}
	}
} 