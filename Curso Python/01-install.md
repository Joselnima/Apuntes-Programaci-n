# Instalación y Configuración de Python

> **Primera vez: descarga Python y prepárate para programar**

---

## ¿QUÉ necesito?

Antes de empezar a programar necesitas:

1. **Python** = el lenguaje (programa que ejecuta código)
2. **Editor** = VS Code, PyCharm, etc (donde escribes)
3. **Terminal** = PowerShell, cmd, bash (para ejecutar)

**Analogía:**
- 🛠️ Python = martillo (herramienta)
- 📝 Editor = papel para dibujar planos
- 💬 Terminal = voz para dar órdenes

---

## ¿PARA QÉ?

### Después de instalar podrás:
```python
print("¡Hola Mundo!")  # ✅ Tu primer programa funciona
```

Sin instalación: ❌ Nada funciona

Con instalación: ✅ Puedes programar

---

## ¿CÓMO?

### 3 Pasos
```
1. INSTALAR Python  → Descarga desde python.org
2. VERIFICAR        → Abre terminal, escribe "python --version"
3. PROBAR           → Crea archivo .py y ejecuta
```

---

## Parte I: Instalación de Python

### En Windows

#### Opción 1: Instalador Oficial (Recomendado)

1. **Descarga**:
   - Ve a [python.org](https://www.python.org/downloads/)
   - **IMPORTANTE**: Marca la casilla **"Add Python to PATH"** ✅ 
   - Descarga la última versión (Python 3.12 o superior)

2. **Ejecución del instalador**:
   ```
   - Haz doble clic en el instalador descargado
   - Selecciona "Install Now" (instalación estándar)
   - O personaliza si sabes qué haces
   ```

3. **Verificación**:
   Abre PowerShell (Windows + X → PowerShell):
   ```bash
   python --version
   # Debes ver: Python 3.x.x
   ```

#### Opción 2: Windows Package Manager (Windows 11+)

```powershell
winget install Python.Python.3.12
```

---

### En macOS

#### Opción 1: Installer Oficial

1. Descarga desde [python.org](https://www.python.org/downloads/)
2. Ejecuta el instalador `.pkg`
3. Sigue las instrucciones

#### Opción 2: Homebrew (Recomendado para macOS)

```bash
brew install python@3.12
```

#### Verificación:

```bash
python3 --version
# Nota: En macOS a veces usa python3 no python
```

---

### En Linux (Ubuntu/Debian)

```bash
sudo apt update
sudo apt install python3.12 python3.12-venv python3-pip
python3 --version
```

#### Otras distribuciones:

**Fedora/RHEL**:
```bash
sudo dnf install python3
```

**Arch Linux**:
```bash
sudo pacman -S python
```

---

## Parte II: Verificación de Instalación

### 1. Verificar Python

```bash
python --version
# O en macOS/Linux:
python3 --version
```

Deberías ver algo como: `Python 3.12.0`

### 2. Acceder a la Shell Interactiva REPL

```bash
python
# O en macOS/Linux:
python3
```

Verás algo así:
```
Python 3.12.0 (main, Oct 2 2023, 00:00:00)
[GCC 12.2.0] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>>
```

### 3. Tu Primer Comando Python

```python
>>> print("¡Hola Python!")
¡Hola Python!
>>> 2 + 2
4
>>> exit()  # Para salir
```

¡Felicidades! Python funciona.

---

## Parte III: Entornos Virtuales (IMPORTANTE)

Un **entorno virtual** es como una carpeta aislada para cada proyecto con sus propias librerías. Esto es **crucial para evitar conflictos**.

### ¿Por qué usar entornos virtuales?

Imagina 2 proyectos:
- Proyecto A necesita Django 3.0
- Proyecto B necesita Django 4.0

Sin entornos virtuales: **Conflicto.** Django no puede estar en dos versiones al mismo tiempo.

Con entornos virtuales: Cada proyecto tiene su propio "mundo" de paquetes.

### Creando tu Primer Entorno Virtual

#### 1. Navega a tu carpeta de proyecto:

```bash
cd ~/mis_proyectos/mi_proyecto
```

#### 2. Crea el entorno virtual:

```bash
# Windows
python -m venv venv

# macOS/Linux
python3 -m venv venv
```

Se crea una carpeta `venv/`. ¡No la toques directamente!

#### 3. Activa el entorno virtual:

**Windows (PowerShell)**:
```powershell
.\venv\Scripts\Activate.ps1
```

**Windows (CMD)**:
```cmd
venv\Scripts\activate.bat
```

**macOS/Linux (Bash/Zsh)**:
```bash
source venv/bin/activate
```

#### 4. Verifica que está activado:

Tu terminal debería mostrar algo como:
```
(venv) C:\Users\tu_usuario\mi_proyecto>
```

El `(venv)` al principio confirma que el entorno está activo.

#### 5. Desactivar (cuando termines):

```bash
deactivate
```

El `(venv)` desaparecerá.

---

## Parte IV: pip - Gestor de Paquetes Python

Python tiene una librería estándar, pero necesitarás librerías adicionales. **pip** es el gestor oficial.

### Verificar pip

```bash
pip --version
```

Deberías ver algo como: `pip 23.2.1 from ...`

### Instalar Paquetes

**¡IMPORTANTE!** Asegúrate de que tu entorno virtual esté **ACTIVADO** primero.

```bash
# Formato general:
pip install nombre_del_paquete

# Ejemplos:
pip install requests          # Para HTTP
pip install flask            # Para web
pip install numpy            # Para ciencia de datos
pip install beautifulsoup4   # Para web scraping
```

### Instalar Versión Específica

```bash
pip install requests==2.28.1
```

### Ver Paquetes Instalados

```bash
pip list
```

### Crear requirements.txt

Para compartir tu proyecto con otros:

```bash
pip freeze > requirements.txt
```

Esto crea un archivo con TODOS los paquetes. Otros pueden instalarlos:

```bash
pip install -r requirements.txt
```

---

## Parte V: Crear tu Primer Archivo Python

### 1. Crea un archivo llamado `hola.py`:

```python
# hola.py
print("¡Hola Mundo!")
nombre = input("¿Cuál es tu nombre? ")
print(f"¡Hola {nombre}!")
```

### 2. Ejecútalo:

**Windows**:
```powershell
python hola.py
```

**macOS/Linux**:
```bash
python3 hola.py
```

### 3. Interactúa:

```
¿Cuál es tu nombre? Juan
¡Hola Juan!
```

---

## Parte VI: IDEs Recomendados (Opcional)

### 1. VS Code (Gratuito, Ligero, Recomendado)

1. Descarga desde [code.visualstudio.com](https://code.visualstudio.com/)
2. Instala la extensión "Python" de Microsoft
3. Crea un archivo `test.py` y VS Code oferecerá instalar pylint/flake8

**Shortcut para ejecutar**: Ctrl+Shift+D → Ejecutar Python

### 2. PyCharm Community (Gratuito, Pesado, Poderoso)

1. Descarga desde [jetbrains.com/pycharm](https://www.jetbrains.com/pycharm/)
2. Crea un nuevo proyecto
3. PyCharm gestiona entornos automáticamente

### 3. Thonny (Excelente para Principiantes)

1. Descarga desde [thonny.org](https://thonny.org/)
2. Super simple: File → New → Escribe → Run
3. Perfecto para aprender sin distracciones

---

## Parte VII: Troubleshooting Común

### Error: "python is not recognized"

**Solución**: Python no está en PATH
- Reinstala Python
- **MUY IMPORTANTE**: Marca "Add Python to PATH"

```bash
# Verifica si está instalado:
where python    # Windows
which python    # macOS/Linux
```

### Error: "modulo XXX no encontrado"

```bash
pip install xxx
# Por ejemplo:
pip install requests
```

### Error: "Permission denied" en Linux/macOS

```bash
# Usa sudo (con cuidado)
sudo pip install paquete
# Mejor aún, usa un entorno virtual (ver arriba)
```

### Error: "Different python versions"

Asegúrate de usar el mismo Python:

```bash
python --version
python -m pip --version  # Debe hablar del mismo Python
```

---

## Parte VIII: Flujo Típico de Trabajo

```mermaid
1. Crea carpeta del proyecto
   ↓
2. python -m venv venv
   ↓
3. Activa: source venv/bin/activate (o activate.bat)
   ↓
4. Crea archivo: hola.py
   ↓
5. Instala librerías: pip install flask
   ↓
6. Ejecuta: python hola.py
   ↓
7. Cuando termines: deactivate
   ↓
8. Para volver: activa de nuevo
```

---

## Parte IX: Resumen Rápido

| Comando | Función | Cuándo |
|---------|---------|-------|
| `python --version` | Ver versión | Para verificar |
| `python` | REPL interactivo | Para experimentar |
| `python archivo.py` | Ejecutar programa | Siempre |
| `python -m venv venv` | Crear entorno | Inicio de proyecto |
| `source venv/bin/activate` | Activar entorno | Antes de trabajar |
| `deactivate` | Desactivar | Cuando terminas |
| `pip install xxx` | Instalar librería | Cuando la necesitas |
| `pip list` | Ver paquetes | Para revisar |
| `pip freeze > requirements.txt` | Guardar proyecto | Para compartir |

---

## ✅ Checklist de Instalación

- [ ] Python 3.12+ instalado
- [ ] `python --version` funciona
- [ ] Puedes entrar a REPL interactivo (`python`)
- [ ] Creaste un entorno virtual
- [ ] El entorno se activa/desactiva
- [ ] pip funciona (`pip --version`)
- [ ] Ejecutaste tu primer script
- [ ] (Opcional) IDE instalado

**¡Si checaste todo, estás listo!** 🚀

---

## 📚 Recursos Adicionales

- [Documentación oficial de Python](https://docs.python.org/3/)
- [Real Python Tutorials](https://realpython.com/)
- [Python Package Index (PyPI)](https://pypi.org/)
- [Stack Overflow tag: python](https://stackoverflow.com/questions/tagged/python)

---

**Siguiente paso**: Lee [estructura_basica_python.md](estructura_basica_python.md) para aprender cómo estructura un programa Python.
