package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

// se crea una cola y se verifica que se comporta como recien creada

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })

}

// se encolan elementos y se desencolan verificando que los primeros son correctos

func TestEncolaDesencola(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Desencolar()
	require.False(t, cola.EstaVacia())
	require.Equal(t, 2, cola.VerPrimero())
	cola.Encolar(3)
	cola.Encolar(4)
	cola.Encolar(5)
	cola.Desencolar()
	cola.Desencolar()
	require.Equal(t, 4, cola.VerPrimero())
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 10001; i++ {
		cola.Encolar(i)
	}
	require.Equal(t, 0, cola.VerPrimero())
	for i := 0; i < 10000; i++ {
		cola.Desencolar()
	}

	require.Equal(t, 10000, cola.VerPrimero())
	cola.Desencolar()

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

// se encolan elementos y luego se desencolan hasta vaciarla, verificando que se comporte como recien creada
func TestDesEncolarHastaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 10; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < 10; i++ {
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

// se encolan elementos enteros en una cola, y strings en otra, verificando que los primeros sean correctos
func TestEncolarDistintos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 51; i++ {
		cola.Encolar(i)
	}
	require.Equal(t, 0, cola.VerPrimero())

	cola_2 := TDACola.CrearColaEnlazada[string]()
	cola_2.Encolar("Hola")
	cola_2.Encolar("Como")
	cola_2.Encolar("Estas?")
	require.Equal(t, "Hola", cola_2.VerPrimero())
	cola_2.Desencolar()
	require.Equal(t, "Como", cola_2.VerPrimero())
}
