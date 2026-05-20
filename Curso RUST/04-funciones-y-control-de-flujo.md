# 04 - Funciones y Control de Flujo

## Declarar Funciones

Una función es un bloque de código reutilizable. En Rust, declaras funciones con `fn`:

```rust
fn saludar() {
    println!("¡Hola!");
}

fn main() {
    saludar();
    saludar();
}
```

Output:
```
¡Hola!
¡Hola!
```

### Funciones con Parámetros

Los parámetros son valores que pasas a la función. **Debes especificar el tipo**:

```rust
fn saludar(nombre: &str) {
    println!("¡Hola, {}!", nombre);
}

fn main() {
    saludar("Ana");
    saludar("Carlos");
}
```

Output:
```
¡Hola, Ana!
¡Hola, Carlos!
```

Múltiples parámetros:

```rust
fn sumar(a: i32, b: i32) {
    let resultado = a + b;
    println!("{} + {} = {}", a, b, resultado);
}

fn main() {
    sumar(5, 3);      // 8
    sumar(100, 50);   // 150
}
```

### Funciones que Retornan Valores

Para retornar un valor, especifica el tipo después de `->`:

```rust
fn sumar(a: i32, b: i32) -> i32 {
    a + b  // Nota: sin punto y coma
}

fn main() {
    let resultado = sumar(5, 3);
    println!("El resultado es: {}", resultado);  // 8
}
```

**Importante**: En la última línea de una función, si omites el `;`, eso es lo que se retorna. Si incluyes `;`, retorna `()` (nada).

```rust
fn cinco() -> i32 {
    5  // ✅ Retorna 5
}

fn nada() {
    5;  // ✅ No retorna nada (la última línea termina con ;)
}
```

Uso explícito de `return`:

```rust
fn mayor(a: i32, b: i32) -> i32 {
    if a > b {
        return a;
    }
    b
}

fn main() {
    println!("{}", mayor(10, 5));  // 10
}
```

## Expresiones vs Declaraciones

- **Declaración**: Ejecuta una acción pero no retorna valor. Termina con `;`
- **Expresión**: Retorna un valor. NO termina con `;`

```rust
fn main() {
    // Declaración
    let x = 5;
    
    // Expresión (nota: sin ;)
    let y = {
        let x = 3;
        x + 1  // Retorna 4
    };
    
    println!("y = {}", y);  // 4
}
```

## Condicionales (if)

El condicional `if` ejecuta código solo si una condición es verdadera:

```rust
fn main() {
    let numero = 6;
    
    if numero > 5 {
        println!("El número es mayor que 5");
    }
}
```

### if-else

```rust
fn main() {
    let numero = 3;
    
    if numero > 5 {
        println!("Mayor que 5");
    } else {
        println!("Menor o igual a 5");
    }
}
```

### if-else if-else

```rust
fn main() {
    let numero = 6;
    
    if numero % 2 == 0 {
        println!("El número es par");
    } else if numero % 3 == 0 {
        println!("El número es divisible por 3");
    } else {
        println!("El número no es par ni divisible por 3");
    }
}
```

### if como Expresión

Como `if` retorna un valor, puedes usarlo para asignar:

```rust
fn main() {
    let condicion = true;
    let numero = if condicion { 5 } else { 6 };
    
    println!("El número es: {}", numero);  // 5
}
```

**Importante**: Los dos branches deben retornar el mismo tipo:

```rust
// ❌ Error: if retorna i32, else retorna &str
let numero = if condicion { 5 } else { "seis" };

// ✅ Correcto
let numero = if condicion { 5 } else { 6 };
```

## Bucles (Loops)

### Loop Infinito

`loop` ejecuta código indefinidamente hasta que encuentres `break`:

```rust
fn main() {
    let mut contador = 0;
    
    loop {
        contador += 1;
        println!("Contador: {}", contador);
        
        if contador == 3 {
            break;  // Sale del loop
        }
    }
    
    println!("Listo!");
}
```

Output:
```
Contador: 1
Contador: 2
Contador: 3
Listo!
```

Retornar valor de un loop:

```rust
fn main() {
    let mut contador = 0;
    
    let resultado = loop {
        contador += 1;
        
        if contador == 10 {
            break contador * 2;  // Retorna 20
        }
    };
    
    println!("Resultado: {}", resultado);
}
```

### While Loop

Ejecuta mientras una condición sea verdadera:

```rust
fn main() {
    let mut numero = 3;
    
    while numero != 0 {
        println!("{}!", numero);
        numero -= 1;
    }
    
    println!("¡Despegue!");
}
```

Output:
```
3!
2!
1!
¡Despegue!
```

### For Loop

Itera sobre una colección o rango:

```rust
fn main() {
    for numero in 1..4 {  // 1, 2, 3 (el 4 no se incluye)
        println!("El número es: {}", numero);
    }
}
```

Output:
```
El número es: 1
El número es: 2
El número es: 3
```

Con inclusión del final:

```rust
fn main() {
    for numero in 1..=4 {  // 1, 2, 3, 4
        println!("El número es: {}", numero);
    }
}
```

For inverso:

```rust
fn main() {
    for numero in (1..4).rev() {  // 3, 2, 1
        println!("El número es: {}", numero);
    }
}
```

## Ejemplo Práctico: Tabla de Multiplicar

```rust
fn main() {
    let tabla = 5;
    
    println!("Tabla del {}", tabla);
    
    for numero in 1..=10 {
        let resultado = tabla * numero;
        println!("{} × {} = {}", tabla, numero, resultado);
    }
}
```

Output:
```
Tabla del 5
5 × 1 = 5
5 × 2 = 10
...
5 × 10 = 50
```

## Funciones Complejas

```rust
fn es_par(numero: i32) -> bool {
    numero % 2 == 0
}

fn cantidad_pares(inicio: i32, fin: i32) -> i32 {
    let mut contador = 0;
    
    for numero in inicio..=fin {
        if es_par(numero) {
            contador += 1;
        }
    }
    
    contador
}

fn main() {
    let pares = cantidad_pares(1, 10);
    println!("Números pares del 1 al 10: {}", pares);  // 5
}
```

## Scope (Alcance)

Las variables solo existen dentro de su scope:

```rust
fn main() {
    let x = 5;
    
    {  // Nuevo scope
        let y = 10;
        println!("x = {}, y = {}", x, y);  // ✅ Ambas existen aquí
    }
    
    println!("x = {}", x);  // ✅ x existe
    // println!("y = {}", y);  // ❌ Error: y no existe aquí
}
```

## Resumen: Palabras Clave de Control

| Palabra | Uso |
|---------|-----|
| `if` | Condicional |
| `else` | Alternativa a if |
| `else if` | Múltiples condiciones |
| `loop` | Bucle infinito |
| `while` | Bucle condicional |
| `for` | Iteración sobre rango/colección |
| `break` | Salir de un bucle |
| `continue` | Ir a siguiente iteración |
| `return` | Salir de función retornando valor |

## Siguiente Paso

Ahora conoces las bases para escribir programas interactivos. El siguiente tema es fundamental en Rust: el sistema de **propiedad (ownership)**, que lo hace único y seguro.

---

**Consejo**: Rust detecta la mayoría de errores en tiempo de compilación. Si algo no compila, lee el error: es muy descriptivo.
