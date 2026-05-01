# Módulo 10 - Ciclos de vida en React

En React, los ciclos de vida describen cuándo se crea un componente, cuándo se actualiza y cuándo se destruye. Esto es útil para:

- cargar datos cuando el componente aparece en pantalla
- actualizar datos cuando cambian las `props` o el `state`
- limpiar recursos cuando el componente ya no se usa

React tiene dos formas principales de pensar el ciclo de vida:

1. Clases (`class components`) con métodos de ciclo de vida.
2. Componentes funcionales con el hook `useEffect`.

---

## 1. Ciclo de vida en componentes de clase

Un componente de clase pasa por tres fases principales:

- Montaje (`mount`): el componente aparece por primera vez.
- Actualización (`update`): el componente vuelve a renderizarse.
- Desmontaje (`unmount`): el componente se elimina.

### Ejemplo con `class`:

```jsx
import React from 'react';

class Contador extends React.Component {
  constructor(props) {
    super(props);
    this.state = { cuenta: 0 };
  }

  componentDidMount() {
    console.log('Contador montado');
    document.title = `Cuenta: ${this.state.cuenta}`;
  }

  componentDidUpdate(prevProps, prevState) {
    if (prevState.cuenta !== this.state.cuenta) {
      console.log('Contador actualizado');
      document.title = `Cuenta: ${this.state.cuenta}`;
    }
  }

  componentWillUnmount() {
    console.log('Contador desmontado');
  }

  incrementar = () => {
    this.setState({ cuenta: this.state.cuenta + 1 });
  };

  render() {
    return (
      <div>
        <h2>Contador de ciclo de vida</h2>
        <p>Cuenta: {this.state.cuenta}</p>
        <button onClick={this.incrementar}>Sumar</button>
      </div>
    );
  }
}

export default Contador;
```

- `constructor`: inicializa el estado.
- `render`: dibuja el componente.
- `componentDidMount`: corre una vez cuando el componente ya está en pantalla.
- `componentDidUpdate`: corre cada vez que cambia el estado o las props.
- `componentWillUnmount`: corre justo antes de que el componente desaparezca.

---

## 2. Ejemplo real: cargar datos y limpiar recursos

```jsx
import React from 'react';

class Perfil extends React.Component {
  state = { usuario: null, cargando: true };

  componentDidMount() {
    fetch('https://jsonplaceholder.typicode.com/users/1')
      .then(response => response.json())
      .then(data => this.setState({ usuario: data, cargando: false }))
      .catch(() => this.setState({ cargando: false }));
  }

  componentWillUnmount() {
    console.log('Perfil se va a desmontar y se limpia todo');
  }

  render() {
    const { usuario, cargando } = this.state;

    if (cargando) return <p>Cargando usuario...</p>;
    if (!usuario) return <p>No se encontró el usuario.</p>;

    return (
      <div>
        <h2>{usuario.name}</h2>
        <p>Email: {usuario.email}</p>
      </div>
    );
  }
}

export default Perfil;
```

---

## 3. Ciclo de vida en componentes funcionales con `useEffect`

En React moderno, se recomienda usar componentes funcionales. El hook `useEffect` reemplaza a los métodos de clase.

### Ejemplo equivalente con `useEffect`:

```jsx
import React, { useState, useEffect } from 'react';

function Contador() {
  const [cuenta, setCuenta] = useState(0);

  useEffect(() => {
    console.log('Contador montado o actualizado');
    document.title = `Cuenta: ${cuenta}`;
  }, [cuenta]);

  return (
    <div>
      <h2>Contador con useEffect</h2>
      <p>Cuenta: {cuenta}</p>
      <button onClick={() => setCuenta(cuenta + 1)}>Sumar</button>
    </div>
  );
}

export default Contador;
```

- Si pasas `[]`, el efecto corre solo una vez al montar.
- Si pasas `[cuenta]`, el efecto corre cada vez que `cuenta` cambia.
- Si no pasas segundo argumento, el efecto corre en cada render.

---

## 4. Limpieza en `useEffect`

Cuando necesitas limpiar un recurso, `useEffect` puede devolver una función de limpieza.

```jsx
import React, { useState, useEffect } from 'react';

function Temporizador() {
  const [segundos, setSegundos] = useState(0);

  useEffect(() => {
    const intervalo = setInterval(() => {
      setSegundos(prev => prev + 1);
    }, 1000);

    return () => {
      clearInterval(intervalo);
      console.log('Temporizador desmontado');
    };
  }, []);

  return <p>Segundos: {segundos}</p>;
}

export default Temporizador;
```

Esta limpieza se ejecuta cuando el componente se desmonta o antes de repetir el efecto.

---

## 6. `memo` y optimización de renders

Cuando un componente padre se renderiza, React también vuelve a renderizar a sus hijos por defecto. A veces ese render no cambia nada y podemos evitarlo usando memoización.

### `React.memo`

`React.memo` memoriza el resultado de un componente funcional y solo lo vuelve a renderizar si sus `props` cambian.

```jsx
import React from 'react';

const Tarjeta = React.memo(function Tarjeta({ texto }) {
  console.log('Tarjeta renderizada');
  return <div>{texto}</div>;
});

export default Tarjeta;
```

- Si el padre renderiza y `texto` no cambia, `Tarjeta` no se renderiza otra vez.
- Esto es útil cuando el componente es pesado o tiene lógica compleja.

### `useMemo` para valores derivados

`useMemo` memoriza el valor calculado entre renders.

```jsx
import React, { useState, useMemo } from 'react';

function CalculadoraLenta({ numero }) {
  const resultado = useMemo(() => {
    console.log('Calculando valor pesado');
    let total = 0;
    for (let i = 0; i < 100000000; i++) {
      total += numero;
    }
    return total;
  }, [numero]);

  return <p>Resultado: {resultado}</p>;
}

export default CalculadoraLenta;
```

- `useMemo` solo recalcula cuando `numero` cambia.
- No uses `useMemo` en todos lados; úsalo cuando el cálculo sea costoso.

### `useCallback` para funciones

`useCallback` memoriza funciones para que no cambien en cada render.

```jsx
import React, { useState, useCallback } from 'react';

function Botones({ onClick }) {
  console.log('Botones renderizados');
  return <button onClick={onClick}>Presionar</button>;
}

function Contenedor() {
  const [count, setCount] = useState(0);

  const manejarClick = useCallback(() => {
    setCount(prev => prev + 1);
  }, []);

  return (
    <div>
      <p>Contador: {count}</p>
      <Botones onClick={manejarClick} />
    </div>
  );
}

export default Contenedor;
```

- `useCallback` es útil cuando pasas funciones a componentes hijos que usan `React.memo`.
- Si la función cambia en cada render, el hijo volverá a renderizar aunque sus props parezcan iguales.

---

## 5. Resumen práctico

- `componentDidMount` = montar una sola vez.
- `componentDidUpdate` = reaccionar a cambios de estado o props.
- `componentWillUnmount` = limpiar antes de desaparecer.
- `useEffect(() => {...}, [])` = montaje.
- `useEffect(() => {...}, [dep])` = actualización basada en dependencia.
- `useEffect(() => { return () => {...}; }, [...])` = limpiar recursos.

Con esto puedes controlar mejor cuándo se ejecutan acciones y evitar efectos secundarios fuera de lugar.
