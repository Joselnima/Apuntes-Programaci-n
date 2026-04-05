# Manejo de fechas en Go

Go no tiene un paquete separado llamado `date`; la fecha se maneja con el paquete `time`, especialmente a través del tipo `time.Time`.

## Operaciones de fecha importantes

- `time.Now()` : devuelve la fecha y hora actuales.
- `time.Date(year, month, day, hour, min, sec, nsec, loc)` : crea una fecha y hora específicas.
- `time.Unix(sec, nsec)` : crea un `time.Time` desde la época Unix.
- `time.Parse(layout, value)` : parsea texto a fecha.
- `time.ParseInLocation(layout, value, loc)` : parsea texto usando una zona horaria.
- `time.ParseDuration(s)` : parsea una duración a `time.Duration`.
- `time.LoadLocation(name)` : carga una zona horaria por nombre.
- `time.FixedZone(name, offset)` : crea una zona horaria fija.
- `time.Since(t)` : duración desde `t` hasta ahora.
- `time.Until(t)` : duración desde ahora hasta `t`.
- `time.Sleep(d)` : pausa la goroutine durante `d`.
- `t.Year()` : obtiene el año.
- `t.Month()` : obtiene el mes.
- `t.Day()` : obtiene el día del mes.
- `t.Weekday()` : obtiene el día de la semana.
- `t.YearDay()` : obtiene el día del año.
- `t.Hour()` : obtiene la hora.
- `t.Minute()` : obtiene los minutos.
- `t.Second()` : obtiene los segundos.
- `t.Nanosecond()` : obtiene los nanosegundos.
- `t.Location()` : devuelve la zona horaria.
- `t.Zone()` : devuelve nombre y desplazamiento de la zona.
- `t.AddDate(years, months, days)` : avanza o retrocede fechas completas.
- `t.Add(d)` : suma una duración.
- `t.Sub(u)` : diferencia entre fechas.
- `t.Before(u)`, `t.After(u)`, `t.Equal(u)` : compara fechas.
- `t.IsZero()` : comprueba si el tiempo es cero.
- `t.Format(layout)` : da salida en un formato de texto.
- `t.Truncate(d)` : redondea hacia abajo a una duración.
- `t.Round(d)` : redondea al múltiplo de una duración.
- `t.UTC()` : convierte a UTC.
- `t.Local()` : convierte a la zona local.
- `t.In(loc)` : convierte a otra zona horaria.
- `t.Unix()` : segundos desde la época Unix.
- `t.UnixNano()` : nanosegundos desde la época Unix.
- `t.UnixMilli()` : milisegundos desde la época Unix.
- `t.UnixMicro()` : microsegundos desde la época Unix.

## Ejemplos

### Crear fechas y extraer componentes
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fecha := time.Date(2026, time.April, 4, 15, 30, 0, 0, time.UTC)
    fmt.Println("Fecha:", fecha)
    fmt.Println("Año:", fecha.Year())
    fmt.Println("Mes:", fecha.Month())
    fmt.Println("Día:", fecha.Day())
    fmt.Println("Día de la semana:", fecha.Weekday())
    fmt.Println("Día del año:", fecha.YearDay())
}
```

### Avanzar o retroceder fechas
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    hoy := time.Now()
    dentroDeUnaSemana := hoy.AddDate(0, 0, 7)
    haceUnMes := hoy.AddDate(0, -1, 0)

    fmt.Println("Hoy:", hoy.Format("2006-01-02"))
    fmt.Println("Dentro de una semana:", dentroDeUnaSemana.Format("2006-01-02"))
    fmt.Println("Hace un mes:", haceUnMes.Format("2006-01-02"))
}
```

### Parsear fechas y formato personalizado
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    layout := "2006-01-02"
    texto := "2026-04-04"

    fecha, err := time.Parse(layout, texto)
    if err != nil {
        panic(err)
    }

    fmt.Println("Fecha parseada:", fecha)
    fmt.Println("Formato largo:", fecha.Format("Monday, 02 January 2006"))
}
```

### Zonas horarias y conversión
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    utc := time.Now().UTC()
    fmt.Println("UTC:", utc)

    loc, err := time.LoadLocation("America/New_York")
    if err != nil {
        panic(err)
    }

    nuevaYork := utc.In(loc)
    fmt.Println("New York:", nuevaYork)
    fmt.Println("Local:", utc.Local())
}
```

### Crear desde Unix y obtener marcas de tiempo
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    desdeUnix := time.Unix(1712370000, 0)
    fmt.Println("Fecha desde Unix:", desdeUnix)
    fmt.Println("Unix segundos:", desdeUnix.Unix())
    fmt.Println("Unix milisegundos:", desdeUnix.UnixMilli())
    fmt.Println("Unix microsegundos:", desdeUnix.UnixMicro())
    fmt.Println("Unix nanosegundos:", desdeUnix.UnixNano())
}
```

### Parsear con zona horaria y duraciones
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    loc, err := time.LoadLocation("Asia/Tokyo")
    if err != nil {
        panic(err)
    }

    texto := "2026-04-04 12:00:00"
    layout := "2006-01-02 15:04:05"
    fecha, err := time.ParseInLocation(layout, texto, loc)
    if err != nil {
        panic(err)
    }

    duracion, err := time.ParseDuration("72h30m")
    if err != nil {
        panic(err)
    }

    fmt.Println("Fecha en Tokyo:", fecha)
    fmt.Println("Añadiendo duración:", fecha.Add(duracion))
}
```

### Comparar y redondear fechas
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    hoy := time.Now()
    otro := hoy.Add(48 * time.Hour)

    fmt.Println("hoy antes de otro?", hoy.Before(otro))
    fmt.Println("otro después de hoy?", otro.After(hoy))
    fmt.Println("igual?", hoy.Equal(otro))

    truncado := hoy.Truncate(24 * time.Hour)
    redondeado := hoy.Round(time.Hour)
    fmt.Println("Truncado a día:", truncado)
    fmt.Println("Redondeado a hora:", redondeado)
}
```
