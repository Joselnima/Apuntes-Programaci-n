# 13 - Concurrencia Segura

Rust permite escribir código paralelo sin race conditions o data races. ¡Esto es revolucionario!

## Threads (Hilos)

Un thread es un hilo de ejecución independiente:

### Crear Threads

```rust
use std::thread;

fn main() {
    let t1 = thread::spawn(|| {
        for i in 1..=3 {
            println!("Thread 1: {}", i);
        }
    });
    
    let t2 = thread::spawn(|| {
        for i in 1..=3 {
            println!("Thread 2: {}", i);
        }
    });
    
    // Esperar a que terminen
    t1.join().unwrap();
    t2.join().unwrap();
}
```

`spawn()` crea un nuevo thread con un closure.

### Transferir Propiedad

```rust
use std::thread;

fn main() {
    let v = vec![1, 2, 3];
    
    let t = thread::spawn(move || {
        println!("{:?}", v);
    });
    
    // v se movió al thread, no puedo usarlo aquí
    t.join().unwrap();
}
```

### Sleep (Espera)

```rust
use std::thread;
use std::time::Duration;

fn main() {
    let t = thread::spawn(|| {
        for i in 1..=3 {
            println!("{}", i);
            thread::sleep(Duration::from_secs(1));
        }
    });
    
    t.join().unwrap();
}
```

## Canales (Message Passing)

Los canales permiten comunicación entre threads:

```rust
use std::thread;
use std::sync::mpsc;

fn main() {
    let (tx, rx) = mpsc::channel();
    
    thread::spawn(move || {
        let valores = vec![1, 2, 3];
        for v in valores {
            tx.send(v).unwrap();
        }
    });
    
    for recibido in rx {
        println!("Recibí: {}", recibido);
    }
}
```

- `tx` = transmisor (sender)
- `rx` = receptor (receiver)
- `mpsc` = Multiple Producer, Single Consumer

### Múltiples Transmisores

```rust
use std::thread;
use std::sync::mpsc;

fn main() {
    let (tx, rx) = mpsc::channel();
    let tx2 = tx.clone();
    
    thread::spawn(move || {
        tx.send("Thread 1").unwrap();
    });
    
    thread::spawn(move || {
        tx2.send("Thread 2").unwrap();
    });
    
    for msg in rx {
        println!("{}", msg);
    }
}
```

## Mutex<T> (Exclusión Mutua)

Asegura que solo un thread acceda a un dato a la vez:

```rust
use std::sync::Mutex;
use std::thread;

fn main() {
    let contador = Mutex::new(0);
    let mut handles = vec![];
    
    for _ in 0..3 {
        let contador_clone = contador.clone();  // ❌ Necesitamos Arc
        
        handles.push(thread::spawn(move || {
            let mut num = contador_clone.lock().unwrap();
            *num += 1;
        }));
    }
    
    for handle in handles {
        handle.join().unwrap();
    }
    
    println!("Resultado: {}", *contador.lock().unwrap());
}
```

**Pero tenemos un problema**: `Mutex` no se puede clonar entre threads sin `Arc`.

## Arc<Mutex<T>> (Atomic Reference Counting + Mutex)

Combinan múltiples propietarios thread-safe + acceso seguro:

```rust
use std::sync::{Arc, Mutex};
use std::thread;

fn main() {
    let contador = Arc::new(Mutex::new(0));
    let mut handles = vec![];
    
    for _ in 0..10 {
        let contador_clone = Arc::clone(&contador);
        
        handles.push(thread::spawn(move || {
            let mut num = contador_clone.lock().unwrap();
            *num += 1;
        }));
    }
    
    for handle in handles {
        handle.join().unwrap();
    }
    
    println!("Resultado: {}", *contador.lock().unwrap());  // 10
}
```

## RwLock<T> (Read-Write Lock)

Permite múltiples lectores pero un solo escritor:

```rust
use std::sync::{Arc, RwLock};
use std::thread;

fn main() {
    let datos = Arc::new(RwLock::new(vec![1, 2, 3]));
    
    // Múltiples lectores simultáneamente
    for i in 0..3 {
        let datos_clone = Arc::clone(&datos);
        thread::spawn(move || {
            let lectura = datos_clone.read().unwrap();
            println!("Thread {} lee: {:?}", i, *lectura);
        });
    }
    
    thread::sleep(std::time::Duration::from_secs(1));
}
```

## Problema: Dead Locks

Los locks pueden causar bloqueos indefinidos:

```rust
// ❌ Cuidado: esto podría deadlock
let lock1 = Mutex::new(1);
let lock2 = Mutex::new(2);

// Thread A toma lock1, espera lock2
// Thread B toma lock2, espera lock1
// ¡Deadlock!
```

**Estrategia**: Siempre toma locks en el mismo orden.

## Ejemplo Completo: Pool de Threads

```rust
use std::sync::{mpsc, Arc, Mutex};
use std::thread;

struct ThreadPool {
    workers: Vec<thread::JoinHandle<()>>,
    sender: mpsc::Sender<Job>,
}

type Job = Box<dyn FnOnce() + Send + 'static>;

impl ThreadPool {
    fn new(tamaño: usize) -> Self {
        let (sender, receiver) = mpsc::channel();
        let receiver = Arc::new(Mutex::new(receiver));
        
        let workers = (0..tamaño)
            .map(|_| {
                let receiver = Arc::clone(&receiver);
                thread::spawn(move || {
                    while let Ok(job) = receiver.lock().unwrap().recv() {
                        job();
                    }
                })
            })
            .collect();
        
        ThreadPool { workers, sender }
    }
    
    fn ejecutar<F>(&self, f: F)
    where
        F: FnOnce() + Send + 'static,
    {
        self.sender.send(Box::new(f)).unwrap();
    }
}

fn main() {
    let pool = ThreadPool::new(4);
    
    for i in 0..10 {
        pool.ejecutar(move || {
            println!("Tarea {}", i);
        });
    }
}
```

## Resumen

| Herramienta | Uso |
|------------|-----|
| `thread::spawn` | Crear threads |
| `mpsc::channel` | Comunicar threads |
| `Mutex` | Acceso exclusivo |
| `Arc` | Compartir propiedad entre threads |
| `RwLock` | Múltiples lectores |

## Siguiente Paso

Concurrencia es un tema avanzado. El próximo capítulo enseña **módulos y paquetes** para organizar código grande.

---

**Consejo**: La concurrencia en Rust es segura. El compilador previene data races. No tienes que preocuparte por condiciones de carrera.
