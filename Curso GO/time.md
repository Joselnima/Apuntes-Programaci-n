# Biblioteca `time` en Go

La biblioteca `time` es la principal forma de trabajar con fechas, horas, duraciones y temporizadores en Go.

## Funciones y tipos principales

- `time.Now()` : obtiene la fecha y hora actual.
- `time.Date(year, month, day, hour, min, sec, nsec, loc)` : crea un valor `time.Time` específico.
- `time.Parse(layout, value)` : convierte una cadena en `time.Time` según el formato dado.
- `time.ParseInLocation(layout, value, loc)` : parsea usando una zona horaria específica.
- `time.ParseDuration(s)` : convierte una cadena como `"2h45m"` en `time.Duration`.
- `time.Unix(sec, nsec)` : crea un `time.Time` a partir de segundos y nanosegundos desde la época Unix.
- `t.Format(layout)` : convierte un `time.Time` a cadena.
- `t.Add(d)` : suma una duración al tiempo.
- `t.AddDate(years, months, days)` : suma años, meses y días.
- `t.Sub(u)` : obtiene la diferencia entre dos tiempos como `time.Duration`.
- `t.Before(u)` : comprueba si `t` es anterior a `u`.
- `t.After(u)` : comprueba si `t` es posterior a `u`.
- `t.Equal(u)` : comprueba si dos tiempos son iguales.
- `t.IsZero()` : comprueba si el tiempo es el valor cero.
- `t.UTC()` / `t.Local()` : convierte la hora al tiempo UTC o local.
- `t.In(loc)` : convierte el tiempo a otra zona horaria.
- `t.Weekday()` : día de la semana.
- `t.YearDay()` : día del año.
- `t.Location()` : devuelve la zona horaria.
- `t.Zone()` : devuelve el nombre y desplazamiento de la zona.
- `t.Unix()` : segundos desde la época Unix.
- `t.UnixNano()` : nanosegundos desde la época Unix.
- `t.UnixMilli()` : milisegundos desde la época Unix.
- `t.UnixMicro()` : microsegundos desde la época Unix.
- `time.Sleep(d)` : pausa la goroutine actual durante una duración.
- `time.AfterFunc(d, f)` : ejecuta una función tras la duración.
- `time.Tick(d)` : devuelve un canal que envía la hora cada duración (no recomendado para detenerse fácilmente).
- `time.NewTimer(d)` : crea un temporizador que dispara una vez.
- `time.NewTicker(d)` : crea un ticker que repite cada duración.
- `time.After(d)` : devuelve un canal que recibe tras la duración.
- `time.Until(t)` : duración hasta el tiempo `t` desde ahora.
- `time.Since(t)` : duración desde el tiempo `t` hasta ahora.
- `time.Duration` : tipo para intervalos de tiempo.
- Constantes: `time.Nanosecond`, `time.Microsecond`, `time.Millisecond`, `time.Second`, `time.Minute`, `time.Hour`.

## Ejemplos

### Hora actual y descomposición
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ahora := time.Now()
    fmt.Println("Ahora:", ahora)
    fmt.Println("Año:", ahora.Year())
    fmt.Println("Mes:", ahora.Month())
    fmt.Println("Día:", ahora.Day())
    fmt.Println("Hora:", ahora.Hour())
    fmt.Println("Minuto:", ahora.Minute())
    fmt.Println("Segundo:", ahora.Second())
    fmt.Println("Zona:", ahora.Location())
}
```

### Formato y parseo
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Estilo Go: Mon Jan 2 15:04:05 MST 2006
    layout := "2006-01-02 15:04:05"
    texto := "2026-04-04 18:30:00"

    t, err := time.Parse(layout, texto)
    if err != nil {
        panic(err)
    }

    fmt.Println("Parseado:", t)
    fmt.Println("Formateado:", t.Format("02/01/2006 03:04 PM"))
}
```

### Duraciones y temporizadores
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    duracion := 2 * time.Second
    fmt.Println("Duración:", duracion)

    fmt.Println("Dormir 2 segundos...")
    time.Sleep(duracion)

    timer := time.NewTimer(1 * time.Second)
    <-timer.C
    fmt.Println("Timer disparado")

    ticker := time.NewTicker(500 * time.Millisecond)
    defer ticker.Stop()

    for i := 0; i < 3; i++ {
        <-ticker.C
        fmt.Println("Tick", i+1)
    }
}
```

### Comparaciones y diferencias
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ahora := time.Now()
    mañana := ahora.AddDate(0, 0, 1)
    pasado := ahora.Add(-24 * time.Hour)

    fmt.Println("Ahora antes de mañana?", ahora.Before(mañana))
    fmt.Println("Ahora después de pasado?", ahora.After(pasado))
    fmt.Println("Diferencia:", mañana.Sub(ahora))
    fmt.Println("Hace:", time.Since(pasado))
}
```
