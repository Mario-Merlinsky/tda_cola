package cola

const ERROR_COLA_VACIA = "La cola esta vacia"

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	return cola
}

func crearNodoCola[T any](dato T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = dato
	return nodo
}

// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil

}

// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje
// "La cola esta vacia".
func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() == true {
		panic(ERROR_COLA_VACIA)
	} else {
		return cola.primero.dato
	}
}

// Encolar agrega un nuevo elemento a la cola, al final de la misma.
func (cola *colaEnlazada[T]) Encolar(elem T) {
	nodo := crearNodoCola(elem)
	if cola.EstaVacia() == true {
		cola.primero = nodo
		cola.ultimo = nodo
	} else {
		cola.ultimo.prox = nodo
		cola.ultimo = nodo
	}
}

// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() == true {
		panic(ERROR_COLA_VACIA)
	} else {
		aux := cola.primero.dato
		cola.primero = cola.primero.prox
		return aux
	}
}
