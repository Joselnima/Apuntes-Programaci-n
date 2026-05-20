# 01 - Instalación y Configuración

## ¿Qué necesitas?

Antes de comenzar a programar en Rust, necesitas tener instalado:
1. El compilador de Rust
2. Cargo (el gestor de paquetes y herramienta de construcción)
3. Un editor de código (VS Code, IntelliJ, Vim, etc.)

## Instalación en Windows

### Opción 1: Usar rustup (Recomendado)

1. Descarga el instalador desde [rustup.rs](https://rustup.rs/)
2. Ejecuta el archivo `.exe` descargado
3. Presiona `1` y luego Enter para instalar con la configuración predeterminada
4. Espera a que se complete la instalación (puede tardar unos minutos)

### Opción 2: Usar Chocolatey

Si tienes Chocolatey instalado:
```powershell
choco install rust
```

### Opción 3: Usar Windows Package Manager

```powershell
winget install Rustlang.Rust.GNU
```

## Instalación en macOS y Linux

### Usar rustup

Abre tu terminal y ejecuta:

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

Luego, sigue las instrucciones en pantalla.

## Verificar la Instalación

Abre una nueva terminal (en Windows, PowerShell o Command Prompt) y ejecuta:

```bash
rustc --version
```

Deberías ver algo como:
```
rustc 1.90.0 (released 2025-09-18)
```

Verifica también que Cargo está instalado:

```bash
cargo --version
```

Deberías ver algo como:
```
cargo 1.90.0 (5fbf4a1ba 2025-09-18)
```

## Actualizar Rust

Para mantener Rust actualizado, ejecuta:

```bash
rustup update
```

## Desinstalar Rust

Si necesitas desinstalar Rust en el futuro, ejecuta:

```bash
rustup self uninstall
```

## Configurar tu Editor

### Visual Studio Code (Recomendado)

1. Instala VS Code desde [code.visualstudio.com](https://code.visualstudio.com/)
2. Abre VS Code y ve a "Extensiones" (Ctrl+Shift+X)
3. Busca e instala "rust-analyzer" de The Rust Programming Language
4. También puedes instalar "Codebase AI" u otra extensión para autocompletado

### Otras Opciones

- **IntelliJ IDEA**: Instala el plugin "Rust" desde el marketplace
- **Vim/Neovim**: Configura con rust.vim o rust-tools.nvim
- **Sublime Text**: Instala el paquete "Rust Enhanced"

## El Playground de Rust

Si no quieres instalar nada localmente, puedes usar el Playground de Rust en línea:
[https://play.rust-lang.org/](https://play.rust-lang.org/)

Es perfecto para experimentar rápidamente con código Rust sin necesidad de instalación.

## Crear tu Primer Proyecto

Una vez instalado Rust, crea tu primer proyecto:

```bash
cargo new hola_mundo
cd hola_mundo
```

Verás una estructura de carpetas:

```
hola_mundo/
├── Cargo.toml
└── src/
    └── main.rs
```

Ejecuta el proyecto:

```bash
cargo run
```

¡Felicidades! Ya has ejecutado tu primer programa en Rust. En el próximo capítulo aprenderás qué significa todo esto.

## Documentación Offline

Rust incluye documentación local. Accede a ella con:

```bash
rustup doc
```

Se abrirá tu navegador con la documentación completa de Rust instalada localmente.

## Siguiente Paso

Ya tienes todo listo. En el próximo capítulo, crearemos nuestro primer programa "Hola, mundo!" y entenderemos cómo funciona.

---

**Nota**: Los comandos mostrados arriba funcionan en Terminal/PowerShell. En Windows, asegúrate de tener PowerShell 7+ o usar Command Prompt. Si tienes problemas, consulta la [documentación oficial](https://doc.rust-lang.org/book/ch01-01-installation.html).
