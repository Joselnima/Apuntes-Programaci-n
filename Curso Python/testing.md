# Testing en Python

> **Escribir código para verificar que tu código funciona - automatizado**

---

## ¿QUÉ es un Test?

Un **test** = un programa que verifica tu código:

```python
# Tu código
def sumar(a, b):
    return a + b

# Test (verificador)
test_resultado = sumar(2, 3)
if test_resultado == 5:
    print("✅ PASA")
else:
    print("❌ FALLA")
```

**Analogía:**
- 🔬 Test = experimento científico
- 📋 Resultado esperado = predicción
- ✅/❌ = comprobamos si acertamos

---

## ¿PARA QÉ sirven?

### 1. Encontrar bugs temprano
```python
# Sin test: bug llega a usuario
# Con test: notas el error ANTES

def restar(a, b):
    return a + b  # ¡ERROR! (suma en vez de restar)

# Test lo descubre:
assert restar(5, 3) == 2  # ❌ FALLA: 8 != 2
```

### 2. Refactorizar sin miedo
```python
# Cambio el código pero los tests garantizan:
# "Todavía funcionaaa igual"

def sumar_lists(numeros):
    return sum(numeros)

# Múltiples tests garantizan que funciona
test_sumar_lists([1, 2, 3])  # ✅
test_sumar_lists([])         # ✅
```

### 3. Documentación ejecutable
```python
# El test MUESTRA cómo usar el código
# Mejor que comentarios (que pueden estar desactualizados)
```

### 4. Cambios seguros
```python
# Alguien modifica tu código
# Pero 100 tests dicen: "Todo funciona ✅"
```

---

## ¿CÓMO funcionan?

### 3 Pasos de un Test

```
1. ARRANCAR  → Prepare datos iniciales
2. EJECUTAR  → Corra la función a testear
3. VERIFICAR → Compare resultado con esperado
```

**Código:**
```python
# Paso 1: Arrancar
a = 2
b = 3

# Paso 2: Ejecutar
resultado = sumar(a, b)

# Paso 3: Verificar
assert resultado == 5, "¡Error! No sumó correctamente"
# Si passa: continúa
# Si falla: 💥 CRASH
```

---

## Parte I: unittest - Testing Estándar

Python incluye `unittest`:

```python
# calculadora.py
def sumar(a, b):
    return a + b

def dividir(a, b):
    if b == 0:
        raise ValueError("No se puede dividir por cero")
    return a / b

# test_calculadora.py
import unittest
from calculadora import sumar, dividir

class TestCalculadora(unittest.TestCase):
    def test_suma_positivos(self):
        self.assertEqual(sumar(2, 3), 5)
    
    def test_suma_negativos(self):
        self.assertEqual(sumar(-2, -3), -5)
    
    def test_suma_mixtos(self):
        self.assertEqual(sumar(5, -3), 2)
    
    def test_suma_cero(self):
        self.assertEqual(sumar(0, 0), 0)
    
    def test_dividir_válido(self):
        self.assertEqual(dividir(10, 2), 5.0)
    
    def test_dividir_por_cero(self):
        with self.assertRaises(ValueError):
            dividir(10, 0)

if __name__ == "__main__":
    unittest.main()
```

Ejecutar:
```bash
python test_calculadora.py
```

### Métodos usados

```python
self.assertEqual(a, b)          # a == b
self.assertNotEqual(a, b)       # a != b
self.assertTrue(x)              # x es True
self.assertFalse(x)             # x es False
self.assertIs(a, b)             # a es b (identity)
self.assertIn(a, b)             # a está en b
self.assertRaises(Exception)    # Lanza excepción
self.assertGreater(a, b)        # a > b
self.assertLess(a, b)           # a < b
```

---

## Parte II: pytest - Framework Moderno

pytest es más simple y poderoso:

```bash
pip install pytest
```

```python
# test_calculadora_pytest.py
from calculadora import sumar, dividir
import pytest

def test_suma_positivos():
    assert sumar(2, 3) == 5

def test_suma_negativos():
    assert sumar(-2, -3) == -5

def test_dividir_por_cero():
    with pytest.raises(ValueError):
        dividir(10, 0)

def test_dividir_válido():
    assert dividir(10, 2) == 5.0
```

Ejecutar:
```bash
pytest test_calculadora_pytest.py
# O simplemente:
pytest
```

---

## Parte III: Fixtures en pytest

Datos reutilizables:

```python
import pytest
from usuario import Usuario

@pytest.fixture
def usuario_valido():
    \"\"\"Crea usuario de prueba\"\"\"
    return Usuario(nombre="Juan", email="juan@test.com")

@pytest.fixture
def usuarios_varios():
    \"\"\"Crea múltiples usuarios\"\"\"
    return [
        Usuario(nombre="Juan", email="juan@test.com"),
        Usuario(nombre="María", email="maria@test.com"),
        Usuario(nombre="Pedro", email="pedro@test.com"),
    ]

def test_usuario_válido(usuario_valido):
    assert usuario_valido.nombre == "Juan"
    assert usuario_valido.email == "juan@test.com"

def test_contar_usuarios(usuarios_varios):
    assert len(usuarios_varios) == 3
```

---

## Parte IV: Parametrización

Ejecutar mismo test con diferentes datos:

```python
import pytest

@pytest.mark.parametrize("entrada,esperado", [
    (2, 4),
    (3, 9),
    (5, 25),
    (0, 0),
    (-2, 4),
])
def test_cuadrado(entrada, esperado):
    assert entrada ** 2 == esperado

# Corre 5 tests automáticamente
```

---

## Parte V: Test Discovery (Descubrimiento)

pytest busca automáticamente:

```
proyecto/
├── src/
│   └── calculadora.py
├── tests/
│   ├── __init__.py
│   ├── test_calculadora.py
│   └── test_usuario.py
└── pytest.ini
```

Archivo `pytest.ini`:
```ini
[pytest]
testpaths = tests
python_files = test_*.py
python_classes = Test*
python_functions = test_*
```

Ejecutar:
```bash
pytest                  # Descubre y ejecuta todo
pytest tests/           # Solo carpeta tests
pytest -v               # Verbose (detallado)
pytest -k "suma"        # Solo tests con "suma" en nombre
```

---

## Parte VI: Mocking - Simular Dependencias

```python
from unittest.mock import Mock, patch
import requests

# Sin mock - haría petición real (lento, dependencia externa)
# ❌ def obtener_usuario_real(id):
#     response = requests.get(f"https://api.example.com/users/{id}")
#     return response.json()

# Con mock - simula la respuesta
@patch('requests.get')
def test_obtener_usuario(mock_get):
    # Configurar mock
    mock_get.return_value.json.return_value = {"id": 1, "nombre": "Juan"}
    
    # Llamar función
    from usuario import obtener_usuario
    resultado = obtener_usuario(1)
    
    # Verificar
    assert resultado["nombre"] == "Juan"
    mock_get.assert_called_once_with("https://api.example.com/users/1")
```

---

## Parte VII: Cobertura de Tests

Ver qué líneas está testeando:

```bash
pip install pytest-cov
pytest --cov=src --cov-report=html
```

Genera reporte HTML con líneas cubiertas.

```python
# coverage.py
def función_no_testeada():
    print("Esto no se ejecuta en tests")  # Aparecerá sin cobertura

def función_testeada():
    return True  # Esta sí se prueba
```

---

## Parte VIII: Arranque y Desmontaje (Setup/Teardown)

Ejecutar código antes/después de tests:

```python
import pytest

class TestBD:
    @pytest.fixture(autouse=True)
    def setup_teardown(self):
        # Setup - antes del test
        print("\\nConectando a BD...")
        self.conexión = conectar_bd()
        
        yield  # Test se ejecuta aquí
        
        # Teardown - después del test
        print("\\nCerrando conexión...")
        self.conexión.cerrar()
    
    def test_insertar(self):
        assert self.conexión.insertar("usuarios", {"nombre": "Juan"})
    
    def test_consultar(self):
        resultado = self.conexión.GET("usuarios")
        assert len(resultado) > 0
```

---

## Parte IX: Test Doubles (Dobles de Prueba)

### Mock vs Stub vs Spy

```python
from unittest.mock import Mock, MagicMock

# STUB - retorna valores fijos
class StubUsuarioDB:
    def obtener(self, id):
        return {"id": 1, "nombre": "Juan"}

# MOCK - verifica interacciones
mock_usuario = Mock()
mock_usuario.obtener.return_value = {"id": 1}
mock_usuario.obtener(1)
mock_usuario.obtener.assert_called_with(1)

# SPY - envuelve objeto real, registra llamadas
real_usuario = UsuarioDB()
spy = MagicMock(wraps=real_usuario)
```

---

## Parte X: Testing de Excepciones

```python
import pytest

def dividir(a, b):
    if b == 0:
        raise ValueError("Denominador no puede ser 0")
    return a / b

def test_excepción():
    with pytest.raises(ValueError) as excinfo:
        dividir(10, 0)
    
    assert "Denominador no puede ser 0" in str(excinfo.value)
```

---

## Parte XI: Estructura AAA

**Arrange** - Preparar datos
**Act** - Ejecutar función
**Assert** - Verificar resultado

```python
def test_calcular_descuento():
    # ARRANGE
    monto = 100
    porcentaje = 10
    
    # ACT
    descuento = aplicar_descuento(monto, porcentaje)
    
    # ASSERT
    assert descuento == 90
```

---

## Parte XII: Buenas Prácticas en Tests

```python
# ✅ BIEN - tests independientes
def test_operación_1():
    resultado = función()
    assert resultado == 5

def test_operación_2():
    resultado = función()
    assert resultado == 5

# ❌ MAL - tests que dependen del orden
def test_a():
    global estado
    estado = 5

def test_b():
    assert estado == 5  # Falla si test_a no corrió primero

# ✅ BIEN - nombres descriptivos
def test_dividir_números_positivos_retorna_cociente():
    assert dividir(10, 2) == 5

# ❌ MAL - nombre vago
def test_divide():
    assert dividir(10, 2) == 5

# ✅ BIEN - test una cosa
def test_sumar_dos_números():
    assert sumar(2, 3) == 5

# ❌ MAL - test múltiples cosas
def test_operaciones():
    assert sumar(2, 3) == 5
    assert restar(5, 2) == 3
    assert multiplicar(3, 4) == 12
```

---

## Parte XIII: CI/CD y Tests

Tests integrados en pipeline automático:

```yaml
# .github/workflows/test.yml
name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v2
    - uses: actions/setup-python@v2
      with:
        python-version: '3.9'
    
    - run: pip install -r requirements-dev.txt
    - run: pytest --cov
    - run: flake8 src/
```

---

## Parte XIV: Ejemplo Completo

```python
# usuario.py
class Usuario:
    def __init__(self, nombre, email):
        if not nombre or not email:
            raise ValueError("Nombre y email requeridos")
        self.nombre = nombre
        self.email = email
    
    def es_email_valido(self):
        return "@" in self.email and "." in self.email

# test_usuario.py
import pytest
from usuario import Usuario

class TestUsuario:
    @pytest.fixture
    def usuario(self):
        return Usuario("Juan", "juan@example.com")
    
    def test_crear_usuario_válido(self, usuario):
        assert usuario.nombre == "Juan"
        assert usuario.email == "juan@example.com"
    
    def test_crear_usuario_sin_nombre(self):
        with pytest.raises(ValueError):
            Usuario("", "email@test.com")
    
    def test_crear_usuario_sin_email(self):
        with pytest.raises(ValueError):
            Usuario("Juan", "")
    
    def test_email_válido(self, usuario):
        assert usuario.es_email_valido() == True
    
    def test_email_inválido(self):
        usuario = Usuario("Juan", "juantest")
        assert usuario.es_email_valido() == False
    
    @pytest.mark.parametrize("email,válido", [
        ("juan@example.com", True),
        ("maria@test.co.uk", True),
        ("pedro", False),
        ("@invalid.com", False),
    ])
    def test_validar_emails(self, email, válido):
        usuario = Usuario("Test", email)
        assert usuario.es_email_valido() == válido
```

Ejecutar:
```bash
pytest test_usuario.py -v
```

---

## Parte XV: Tips y Trucos

```python
# Saltar test
@pytest.mark.skip(reason="No implementado todavía")
def test_futuro():
    pass

# Marcar como esperado fallar
@pytest.mark.xfail
def test_bug_conocido():
    assert 1 == 2  # Falla pero no reporta error

# Solo tests con cierto marcador
@pytest.mark.slow
def test_operación_lenta():
    pass

# Ejecutar: pytest -m slow
```

---

## Resumen

| Concepto | Uso |
|----------|-----|
| unittest | Testing estándar (incluido) |
| pytest | Testing moderno (recomendado) |
| Fixtures | Datos reutilizables |
| Parametrización | Mismo test, múltiples datos |
| Mock | Simular dependencias |
| Coverage | Ver qué está testeado |
| AAA | Arrange-Act-Assert |
| CI/CD | Tests automáticos |

---

**Siguiente**: [context_managers.md](context_managers.md)
