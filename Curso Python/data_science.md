# Data Science en Python

> **Analizar números, crear gráficos, encontrar patrones en datos**

---

## ¿QUÉ es Data Science?

**Data Science** = extraer información valiosa de datos:

```
1. Datos en bruto       → [1,2,3,1,2,5,4,3,2,1,3]
2. Análisis (¿promedio?) ↓ (promedio = 2.6)
3. Visualización       → Gráfico
4. Conclusión          → "Hay más 2s y 3s"
```

**Analogía:**
- 📊 Accountant = suma números (Excel)
- 🔬 Data Scientist = busca PATRONES (Python + libraries)

**Se necesita:**
- 🔢 NumPy = matemática rápida
- 📊 Pandas = manipular datos
- 📈 Matplotlib = dibujar gráficos
- 🤖 Scikit-learn = máquinas que aprenden

---

## ¿PARA QÉ sirve?

### 1. Encontrar tendencias
```python
# Datos de ventas 2020-2024
# ¿Crece cada año? ¿Temporadas?
# ¿A qué mes se vende más?
```

### 2. Predicciones
```python
# Histórico de clima
# Modelo predice: "Mañana 25°C"

# Histórico de precios
# Modelo predice: "ETH subirá a $4000"
```

### 3. Detectar anomalías
```python
# Monitorear tráfico de red
# Detectar: "¡ATAQUE! Tráfico anormal"

# Fraude en tarjetas
# Detectar: "Esta compra es sospechosa"
```

### 4. Agrupar información
```python
# 1000 clientes
# Agrupar en: "VIP", "Frecuente", "Casual"
# Estrategia diferente para cada grupo
```

---

## ¿CÓMO funcionan?

### 4 Etapas del Data Science

```
DATOS → LIMPIAR → ANALIZAR → VISUALIZAR → ACTUAR
  ↑
Datos crudos,
con errores,
incompletos

                          ↓
                    Tabla perfecta
                         ↓
                    NumPy/Pandas hace
                    cálculos
                         ↓
                    Matplotlib dibuja
                    gráficos
                         ↓
                    Tomas decisión
```

### 4 Librerías Principales

| Librería | Para qué | Analogía |
|----------|----------|----------|
| **NumPy** | Operaciones rápidas | Calculadora turbo |
| **Pandas** | Tablas de datos | Excel en Python |
| **Matplotlib** | Gráficos | Dibujante |
| **Scikit-learn** | Máquinas inteligentes | Robot que aprende |

---

## Parte I: NumPy - Arrays y Matemática

```bash
pip install numpy
```

```python
import numpy as np

# Crear arrays
arr = np.array([1, 2, 3, 4, 5])
print(arr.dtype)        # int64
print(arr.shape)        # (5,)

# Matrices
matriz = np.array([[1, 2, 3], [4, 5, 6]])
print(matriz.shape)     # (2, 3)

# Inicialización
ceros = np.zeros((2, 3))
unos = np.ones((3, 3))
rango = np.arange(0, 10, 2)        # 0 a 10, paso 2
linspace = np.linspace(0, 1, 5)    # 5 valores entre 0 y 1

# Operaciones
a = np.array([1, 2, 3])
b = np.array([4, 5, 6])

suma = a + b            # [5, 7, 9]
producto = a * b        # [4, 10, 18]
division = b / a        # [4.0, 2.5, 2.0]

# Funciones
print(np.sum(a))        # 6
print(np.mean(a))       # 2.0
print(np.std(a))        # 0.816...
print(np.max(a))        # 3
print(np.argmax(a))     # Índice del máximo
```

---

## Parte II: Pandas - Manipulación de Datos

```bash
pip install pandas
```

```python
import pandas as pd

# Series (1D)
series = pd.Series([10, 20, 30], index=['a', 'b', 'c'])
print(series['a'])      # 10

# DataFrame (2D - tabla)
df = pd.DataFrame({
    'nombre': ['Juan', 'María', 'Pedro'],
    'edad': [25, 30, 35],
    'salario': [3000, 3500, 4000]
})

print(df.head())        # Primeras filas
print(df.shape)         # (3, 3)
print(df.columns)       # Nombres de columnas

# Acceso
print(df['nombre'][0])  # Juan
print(df.loc[0])        # Fila 0 (por índice)
print(df.iloc[0, 1])    # [0, 1] (por posición)

# Operaciones
df['edad_años'] = df['edad']  # Nueva columna
df.drop('edad_años', axis=1)  # Eliminar columna

# Filtros
mayores_30 = df[df['edad'] > 30]
print(mayores_30)

# Estadísticas
print(df.describe())    # Media, std, min, max, etc
print(df['edad'].mean())

# Agrupación
por_rango = df.groupby(lambda x: 'joven' if df.loc[x, 'edad'] < 30 else 'mayor')['salario'].sum()
```

---

## Parte III: Lectura de Datos

```python
import pandas as pd

# CSV
df = pd.read_csv('datos.csv')
df.to_csv('salida.csv', index=False)

# Excel
df = pd.read_excel('datos.xlsx', sheet_name='Hoja1')
df.to_excel('salida.xlsx')

# JSON
df = pd.read_json('datos.json')
df.to_json('salida.json')

# SQL
import sqlite3
conexión = sqlite3.connect('base_datos.db')
df = pd.read_sql('SELECT * FROM usuarios', conexión)
```

---

## Parte IV: Matplotlib - Visualización

```bash
pip install matplotlib
```

```python
import matplotlib.pyplot as plt
import numpy as np

# Línea simple
x = np.linspace(0, 10, 100)
y = np.sin(x)

plt.plot(x, y, label='sin(x)', color='blue', linewidth=2)
plt.xlabel('X')
plt.ylabel('Y')
plt.title('Función Seno')
plt.legend()
plt.grid(True)
plt.show()

# Múltiples líneas
plt.plot(x, np.sin(x), label='sin(x)')
plt.plot(x, np.cos(x), label='cos(x)')
plt.legend()
plt.show()

# Scatter (puntos)
plt.scatter([1, 2, 3, 4], [1, 4, 2, 3], color='red', s=100)
plt.show()

# Histograma
datos = np.random.normal(100, 15, 1000)
plt.hist(datos, bins=30, color='green', alpha=0.7)
plt.show()

# Barras
categorías = ['A', 'B', 'C', 'D']
valores = [10, 20, 15, 30]
plt.bar(categorías, valores, color='orange')
plt.show()

# Subplots
fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(12, 5))
ax1.plot(x, np.sin(x))
ax1.set_title('Seno')
ax2.plot(x, np.cos(x))
ax2.set_title('Coseno')
plt.show()
```

---

## Parte V: Seaborn - Visualización Estadística

```bash
pip install seaborn
```

```python
import seaborn as sns
import pandas as pd
import matplotlib.pyplot as plt

# Datos de ejemplo
iris = sns.load_dataset('iris')

# Scatter con colores
sns.scatterplot(data=iris, x='sepal_length', y='sepal_width', 
                hue='species', s=100)
plt.show()

# Histograma con distribución
sns.histplot(data=iris, x='sepal_length', hue='species', kde=True)
plt.show()

# Matriz de correlación
correlación = iris.corr()
sns.heatmap(correlación, annot=True, cmap='coolwarm')
plt.show()

# Boxplot
sns.boxplot(data=iris, x='species', y='sepal_length')
plt.show()

# Pairplot (todos vs todos)
sns.pairplot(iris, hue='species')
plt.show()
```

---

## Parte VI: Scikit-learn - Machine Learning

```bash
pip install scikit-learn
```

### Regresión Lineal

```python
from sklearn.linear_model import LinearRegression
from sklearn.model_selection import train_test_split
import numpy as np

# Datos
X = np.array([[1], [2], [3], [4], [5]])
y = np.array([2, 4, 5, 4, 5])

# Dividir datos
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# Modelo
modelo = LinearRegression()
modelo.fit(X_train, y_train)

# Predicción
y_pred = modelo.predict(X_test)
print(modelo.score(X_test, y_test))  # R² score

# Coeficiente
print(modelo.coef_)
print(modelo.intercept_)
```

### Clasificación - Iris Dataset

```python
from sklearn.datasets import load_iris
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score, classification_report

# Cargar datos
iris = load_iris()
X = iris.data
y = iris.target

# Dividir
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2)

# Modelo
modelo = RandomForestClassifier(n_estimators=100)
modelo.fit(X_train, y_train)

# Evaluación
y_pred = modelo.predict(X_test)
print(f"Precisión: {accuracy_score(y_test, y_pred):.2%}")
print(classification_report(y_test, y_pred))
```

---

## Parte VII: Clustering - K-Means

```python
from sklearn.cluster import KMeans
import numpy as np
import matplotlib.pyplot as plt

# Datos
X = np.random.randn(100, 2)

# Clustering
kmeans = KMeans(n_clusters=3, random_state=42)
clusters = kmeans.fit_predict(X)

# Visualizar
plt.scatter(X[:, 0], X[:, 1], c=clusters, cmap='viridis', s=100)
plt.scatter(kmeans.cluster_centers_[:, 0], 
           kmeans.cluster_centers_[:, 1],
           c='red', marker='X', s=500)
plt.show()

print(f"Centroides: {kmeans.cluster_centers_}")
```

---

## Parte VIII: Normalización de Datos

```python
from sklearn.preprocessing import StandardScaler, MinMaxScaler
import numpy as np

datos = np.array([[1, 2], [3, 4], [5, 6]])

# StandardScaler (media 0, std 1)
scaler = StandardScaler()
datos_escalados = scaler.fit_transform(datos)

# MinMaxScaler (0 a 1)
scaler = MinMaxScaler()
datos_01 = scaler.fit_transform(datos)
```

---

## Parte IX: Validación Cruzada

```python
from sklearn.model_selection import cross_val_score
from sklearn.ensemble import RandomForestClassifier
from sklearn.datasets import load_iris

iris = load_iris()

modelo = RandomForestClassifier()

# 5-fold cross validation
scores = cross_val_score(modelo, iris.data, iris.target, cv=5)
print(f"Scores: {scores}")
print(f"Promedio: {scores.mean():.2f} (+/- {scores.std():.2f})")
```

---

## Parte X: Ejemplo Completo - Predicción de Precios

```python
import pandas as pd
import numpy as np
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestRegressor
from sklearn.metrics import mean_squared_error, r2_score
import matplotlib.pyplot as plt

# Datos de ejemplo
np.random.seed(42)
X = np.random.rand(100, 3) * 100
y = X[:, 0] * 2 + X[:, 1] * 3 + X[:, 2] * 0.5 + np.random.normal(0, 5, 100)

# Dividir
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2)

# Entrenar
modelo = RandomForestRegressor(n_estimators=100)
modelo.fit(X_train, y_train)

# Evaluar
y_pred = modelo.predict(X_test)
mse = mean_squared_error(y_test, y_pred)
r2 = r2_score(y_test, y_pred)

print(f"MSE: {mse:.2f}")
print(f"R²: {r2:.2f}")

# Importancia de características
importancia = modelo.feature_importances_
for i, imp in enumerate(importancia):
    print(f"Característica {i}: {imp:.2%}")

# Visualizar
plt.scatter(y_test, y_pred)
plt.xlabel('Real')
plt.ylabel('Predicho')
plt.show()
```

---

## Parte XI: TensorFlow - Deep Learning (Intro)

```bash
pip install tensorflow
```

```python
import tensorflow as tf
from tensorflow import keras
import numpy as np

# Datos
X_train = np.random.rand(1000, 10)
y_train = np.random.randint(0, 2, 1000)

# Modelo
modelo = keras.Sequential([
    keras.layers.Dense(64, activation='relu', input_shape=(10,)),
    keras.layers.Dense(32, activation='relu'),
    keras.layers.Dense(1, activation='sigmoid')
])

# Compilar
modelo.compile(optimizer='adam', 
              loss='binary_crossentropy',
              metrics=['accuracy'])

# Entrenar
modelo.fit(X_train, y_train, epochs=10, batch_size=32)

# Predicción
predicción = modelo.predict([[0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0]])
```

---

## Parte XII: Checklist Data Science

1. ✅ Cargar datos (pandas, numpy)
2. ✅ Exploración (describe, head, visualización)
3. ✅ Limpieza (valores nulos, outliers, tipo de datos)
4. ✅ Normalización (escalar características)
5. ✅ Dividir datos (train/test)
6. ✅ Entrenar modelo (scikit-learn, TensorFlow)
7. ✅ Evaluar (cross-validation, métricas)
8. ✅ Visualizar resultados
9. ✅ Iterar y mejorar

---

## Resumen

| Librería | Usa |
|----------|-----|
| NumPy | Arrays numéricos |
| Pandas | DataFrames |
| Matplotlib | Visualización básica |
| Seaborn | Visualización estadística |
| Scikit-learn | Machine learning |
| TensorFlow | Deep learning |
| SciPy | Computación científica |

---

**Fin del Curso Básico** 🎓

Para continuar: Machine Learning avanzado, NLP, Computer Vision, etc.
