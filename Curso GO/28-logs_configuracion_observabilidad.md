# Logs, Configuración y Observabilidad en Go

> **Cómo registrar eventos, configurar aplicaciones y monitorear su comportamiento en producción**

---

## ¿Por qué existen los logs?

Imagina que tu aplicación falla en producción a las 3 AM. El cliente dice: "El servidor no responde".

**Sin logs:**
- ❓ ¿Qué pasó? No sé.
- ❓ ¿Cuándo pasó? No sé cuándo falló.
- ❓ ¿Por qué pasó? No tengo ni idea.
- 😱 Estás perdido.

**Con logs:**
- ✅ `2026-04-05 03:15:23 ERROR database connection timeout`
- ✅ `2026-04-05 03:15:24 ERROR query failed: context deadline exceeded`
- ✅ Sabes exactamente qué, cuándo y por qué.

**Los logs son tu "caja negra" para entender qué hizo la aplicación.**

---

## Parte I: Logging Estructurado vs Logging Desordenado

### ❌ Logging Desordenado (Malo)

```go
fmt.Println("Usuario creado")
fmt.Println("Error al conectar BD")
fmt.Println("Request: /api/usuarios")
```

**Problemas:**
- No tiene estructura
- Imposible filtrar por tipo de evento
- Sin timestamps
- Sin niveles de severidad (info, error, warning)
- Imposible parsear automáticamente

---

### ✅ Logging Estructurado (Bien)

```go
log.Printf("level=info user_id=123 action=create event=usuario_registrado")
log.Printf("level=error component=database error=connection_timeout")
log.Printf("level=info method=GET path=/api/usuarios status=200")
```

**Ventajas:**
- ✅ Estructura clara: `key=value`
- ✅ Fácil de filtrar: `grep "level=error"`
- ✅ Fácil de parsear: Herramientas pueden leerlo
- ✅ Fácil de buscar en logs centralizados

---

## Parte II: Logging con Slog (Standard de Go 1.21+)

Go 1.21 incluyó **slog**, la librería estándar de logging estructurado.

### Logging básico

```go
package main

import (
    "log/slog"
    "os"
)

func main() {
    // Logger por defecto al stdout
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    
    // Info
    logger.Info("aplicación iniciada", 
        "version", "1.0.0",
        "puerto", 8080,
    )
    
    // Warning
    logger.Warn("conexión lenta", 
        "duracion_ms", 5000,
        "esperado_ms", 1000,
    )
    
    // Error
    logger.Error("fallo al guardar usuario", 
        "user_id", 123,
        "error", "duplicate key",
    )
}
```

**Output:**
```
time=2026-04-05T15:30:00.000Z level=INFO msg="aplicación iniciada" version=1.0.0 puerto=8080
time=2026-04-05T15:30:05.000Z level=WARN msg="conexión lenta" duracion_ms=5000 esperado_ms=1000
time=2026-04-05T15:30:10.000Z level=ERROR msg="fallo al guardar usuario" user_id=123 error="duplicate key"
```

---

### Niveles de logging

```go
logger.Debug("debug info")    // Muy detallado, solo desarrollo
logger.Info("evento normal")   // Información general
logger.Warn("advertencia")     // Algo raro pero continuamos
logger.Error("error")          // Error, pero recuperable
```

---

### Logging JSON (para producción)

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

logger.Info("usuario registrado", 
    "user_id", 456,
    "email", "juan@mail.com",
)
```

**Output:**
```json
{"time":"2026-04-05T15:30:00Z","level":"INFO","msg":"usuario registrado","user_id":456,"email":"juan@mail.com"}
```

**Ventaja:** JSON es fácil de parsear en herramientas centralizadas (CloudWatch, DataDog, ELK).

---

### Logger global

```go
import "log/slog"

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func registrarUsuario(email string) error {
    logger.Info("iniciando registro", "email", email)
    
    err := db.Save(Usuario{Email: email})
    if err != nil {
        logger.Error("fallo registro", "email", email, "error", err)
        return err
    }
    
    logger.Info("usuario registrado", "email", email)
    return nil
}
```

---

## Parte III: Contextualización de Logs

### Agregar contexto a cada request

```go
import (
    "context"
    "log/slog"
)

// Agregar request ID al contexto
func middlewareConContextoLog(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Generar request ID único
        requestID := generateRequestID()
        
        // Agregar al contexto
        ctx := context.WithValue(r.Context(), "requestID", requestID)
        ctx = context.WithValue(ctx, "method", r.Method)
        ctx = context.WithValue(ctx, "path", r.URL.Path)
        
        logger.Info("request iniciado", 
            "requestID", requestID,
            "method", r.Method,
            "path", r.URL.Path,
        )
        
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func handleGetPerfil(w http.ResponseWriter, r *http.Request) {
    requestID := r.Context().Value("requestID")
    
    logger.Info("obtener perfil", "requestID", requestID)
    // ...
}
```

---

## Parte IV: Configuración de Aplicaciones

### El problema: Hardcodear valores

```go
// ❌ MALO: Valores pegados en código
const DBHost = "localhost"
const DBPort = 5432
const ApiKey = "sk_live_123456"
```

**Problemas:**
- Cambiar configuración = recompilar
- Secretos en el código fuente
- Diferente config para dev/staging/prod

---

### ✅ La solución: Variables de entorno

```go
import "os"

func main() {
    // Leer variables de entorno
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    apiKey := os.Getenv("API_KEY")
    
    if dbHost == "" {
        log.Fatal("DB_HOST no definido")
    }
}
```

**Uso:**
```bash
export DB_HOST=localhost
export DB_PORT=5432
export API_KEY=sk_live_123
go run main.go
```

---

### Mejora: Struct de configuración

```go
type Config struct {
    // Database
    DBHost     string
    DBPort     int
    DBUser     string
    DBPassword string
    
    // Server
    ServerPort int
    
    // Seguridad
    JWTSecret string
    
    // Comportamiento
    LogLevel string
}

func cargarConfig() *Config {
    return &Config{
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     atoi(os.Getenv("DB_PORT")),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        ServerPort: atoi(os.Getenv("SERVER_PORT")),
        JWTSecret:  os.Getenv("JWT_SECRET"),
        LogLevel:   os.Getenv("LOG_LEVEL"),
    }
}

func main() {
    config := cargarConfig()
    logger.Info("configuración cargada", "puerto", config.ServerPort)
}
```

---

### Con valores por defecto

```go
func cargarConfig() *Config {
    return &Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnvInt("DB_PORT", 5432),
        ServerPort: getEnvInt("SERVER_PORT", 8080),
        LogLevel:   getEnv("LOG_LEVEL", "info"),
    }
}

func getEnv(key, defaultValue string) string {
    valor := os.Getenv(key)
    if valor == "" {
        return defaultValue
    }
    return valor
}

func getEnvInt(key string, defaultValue int) int {
    valor := os.Getenv(key)
    if valor == "" {
        return defaultValue
    }
    num, err := strconv.Atoi(valor)
    if err != nil {
        return defaultValue
    }
    return num
}
```

---

### Librería: viper (para archivos .env)

```go
import "github.com/spf13/viper"

func cargarConfigConViper() {
    viper.SetConfigFile(".env")
    viper.SetConfigType("env")
    viper.AutomaticEnv()
    
    // Valores por defecto
    viper.SetDefault("port", 8080)
    viper.SetDefault("log_level", "info")
    
    err := viper.ReadInConfig()
    if err != nil {
        log.Println("usando ENV vars o defaults")
    }
    
    // Acceder
    puerto := viper.GetInt("port")
    logLevel := viper.GetString("log_level")
}
```

**.env file:**
```
DB_HOST=localhost
DB_PORT=5432
SERVER_PORT=8080
LOG_LEVEL=debug
```

---

## Parte V: Observabilidad - Métricas

### ¿Qué es observabilidad?

**Observabilidad** = Capacidad de entender qué está haciendo tu aplicación sin cambiar el código.

Se compone de 3 pilares:
1. **Logs** - Eventos que sucedieron
2. **Métricas** - Números medibles (latencia, errores, etc)
3. **Trazas distribuidas** - Seguimiento de requests entre servicios

---

### Recolectar métricas simples

```go
import (
    "time"
)

type Metricas struct {
    RequestsTotal     int64
    RequestsErrores   int64
    DuracionPromedio  time.Duration
}

var metricas = &Metricas{}

func middlewareMetricas(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        inicio := time.Now()
        
        // Wrappear response writer para capturar status code
        rw := &responseWriter{ResponseWriter: w}
        
        next.ServeHTTP(rw, r)
        
        duracion := time.Since(inicio)
        
        // Actualizar métricas
        metricas.RequestsTotal++
        if rw.statusCode >= 400 {
            metricas.RequestsErrores++
        }
        metricas.DuracionPromedio = (metricas.DuracionPromedio + duracion) / 2
        
        logger.Info("request completado",
            "duracion_ms", duracion.Milliseconds(),
            "status", rw.statusCode,
        )
    })
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}
```

---

### Endpoint para exponer métricas

```go
func handleMetricas(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "requests_total": metricas.RequestsTotal,
        "requests_errores": metricas.RequestsErrores,
        "duracion_promedio_ms": metricas.DuracionPromedio.Milliseconds(),
    })
}

mux := http.NewServeMux()
mux.HandleFunc("/metrics", handleMetricas)
```

---

### Prometheus (Estándar en industria)

```bash
go get github.com/prometheus/client_golang
```

```go
import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// Definir métricas
var (
    requestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total requests",
        },
        []string{"method", "path", "status"},
    )
    
    requestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "http_request_duration_seconds",
            Help: "Request duration",
        },
        []string{"method", "path"},
    )
)

func init() {
    prometheus.MustRegister(requestsTotal)
    prometheus.MustRegister(requestDuration)
}

// Usar en handler
func handleAPI(w http.ResponseWriter, r *http.Request) {
    inicio := time.Now()
    
    // ... procesar request ...
    
    duracion := time.Since(inicio)
    requestsTotal.WithLabelValues(r.Method, r.URL.Path, "200").Inc()
    requestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duracion.Seconds())
}

// Exponer métricas
mux := http.NewServeMux()
mux.Handle("/metrics", promhttp.Handler())
```

---

## Parte VI: Health Checks

### Endpoint de salud

```go
type HealthStatus struct {
    Status   string    `json:"status"`   // "healthy" o "unhealthy"
    Uptime   int64     `json:"uptime"`   // Segundos
    Database string    `json:"database"` // "ok" o "error"
    Version  string    `json:"version"`
}

var startTime = time.Now()

func handleHealth(w http.ResponseWriter, r *http.Request) {
    // Verificar BD
    dbStatus := "ok"
    err := db.Ping()
    if err != nil {
        dbStatus = "error"
    }
    
    status := "healthy"
    if dbStatus != "ok" {
        status = "unhealthy"
    }
    
    health := HealthStatus{
        Status:   status,
        Uptime:   int64(time.Since(startTime).Seconds()),
        Database: dbStatus,
        Version:  "1.0.0",
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(health)
}

mux := http.NewServeMux()
mux.HandleFunc("/health", handleHealth)
```

**Uso en Kubernetes:**
```yaml
livenessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 10
```

---

## Parte VII: Trazas Distribuidas (Básico)

Para seguimiento de requests entre servicios:

```bash
go get go.opentelemetry.io/otel
```

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("miapp")

func handleRegistro(w http.ResponseWriter, r *http.Request) {
    ctx, span := tracer.Start(r.Context(), "registrar_usuario")
    defer span.End()
    
    // Guardar en BD
    ctx2, spanDB := tracer.Start(ctx, "guardar_bd")
    err := db.WithContext(ctx2).Create(&usuario).Error
    spanDB.End()
    
    if err != nil {
        span.RecordError(err)
        http.Error(w, "error", http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusCreated)
}
```

---

## Parte VIII: Anti-patrones y Errores Comunes

### ❌ Loguear información sensible

```go
// MALO: Expone contraseña en logs
logger.Info("login intento", "email", email, "password", password)

// BIEN: Loguear solo evento
logger.Info("login intento", "email", email)
```

---

### ❌ Demasiados logs

```go
// MALO: Un log por cada línea
for i := 0; i < 1000000; i++ {
    logger.Debug("procesando item", "i", i)  // ¡Un millón de logs!
}

// BIEN: Log del progreso
logger.Info("procesando items", "total", 1000000)
// ... después
logger.Info("procesamiento completado")
```

---

### ❌ Hardcodear configuración

```go
// MALO
const DatabaseURL = "postgresql://user:pass@localhost:5432/db"

// BIEN
databaseURL := os.Getenv("DATABASE_URL")
```

---

### ❌ Ignorar errores de configuración

```go
// MALO
config := cargarConfig()  // Si falla, no lo sabes

// BIEN
config, err := cargarConfig()
if err != nil {
    log.Fatalf("error cargando config: %v", err)
}
```

---

## Parte IX: Best Practices - Resumen

### 1. Logging estructurado con slog

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("evento", "key1", value1, "key2", value2)
```

---

### 2. Niveles de log apropiados

```go
logger.Debug("para desarrolladores")      // Modo development
logger.Info("eventos normales")             // Siempre loguear
logger.Warn("situación rara")              // Revisar luego
logger.Error("error recuperable")          // Muy importante
```

---

### 3. Contexto en logs

```go
logger.Info("acción", 
    "requestID", requestID,
    "userID", userID,
    "action", "create_usuario",
)
```

---

### 4. Configuración por variables de entorno

```go
config := &Config{
    DBHost: os.Getenv("DB_HOST"),
    Port:   getEnvInt("PORT", 8080),
}
```

---

### 5. Health checks activos

```go
func handleHealth(w http.ResponseWriter, r *http.Request) {
    if !isHealthy() {
        w.WriteHeader(http.StatusServiceUnavailable)
        return
    }
    w.WriteHeader(http.StatusOK)
}
```

---

### 6. Métricas básicas

```go
requestsTotal.Inc()
requestDuration.Observe(duracion.Seconds())
```

---

### 7. No loguear ni exponer secretos

```go
// ❌ NUNCA
logger.Info("api key", "key", apiKey)

// ✅ BIEN
logger.Info("conectando a api externa")
```

---

### 8. Logs JSON en producción

```go
if os.Getenv("ENVIRONMENT") == "production" {
    logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
} else {
    logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
}
```

---

### 9. Agregar contexto a cada request

```go
func middlewareContext(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestID := generateID()
        ctx := context.WithValue(r.Context(), "requestID", requestID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

---

### 10. Monitorear regularmente

```bash
# Ver logs de hoy
grep "level=ERROR" app.log | tail -100

# Ver métrica de requests por segundo
curl http://localhost:8080/metrics | grep http_requests_total
```

---

## Ejemplo Completo: API con Logs, Config y Observabilidad

```go
package main

import (
    "encoding/json"
    "log/slog"
    "net/http"
    "os"
    "time"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type Config struct {
    Port     int
    LogLevel string
}

var metricas struct {
    requests int64
    errores  int64
}

func middlewareLogging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        inicio := time.Now()
        
        logger.Info("request iniciado", 
            "method", r.Method,
            "path", r.URL.Path,
        )
        
        next.ServeHTTP(w, r)
        
        duracion := time.Since(inicio)
        logger.Info("request completado",
            "method", r.Method,
            "path", r.URL.Path,
            "duracion_ms", duracion.Milliseconds(),
        )
        
        metricas.requests++
    })
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status": "healthy",
        "uptime": time.Now().Format(time.RFC3339),
    })
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "requests": metricas.requests,
        "errores":  metricas.errores,
    })
}

func main() {
    config := &Config{
        Port:     getEnvInt("PORT", 8080),
        LogLevel: os.Getenv("LOG_LEVEL"),
    }
    
    logger.Info("iniciando aplicación", "puerto", config.Port)
    
    mux := http.NewServeMux()
    mux.Handle("/api", middlewareLogging(http.HandlerFunc(handleAPI)))
    mux.HandleFunc("/health", handleHealth)
    mux.HandleFunc("/metrics", handleMetrics)
    
    addr := ":" + strconv.Itoa(config.Port)
    logger.Info("servidor escuchando", "addr", addr)
    
    err := http.ListenAndServe(addr, mux)
    if err != nil {
        logger.Error("server error", "error", err)
    }
}

func getEnvInt(key string, defaultValue int) int {
    valor := os.Getenv(key)
    if valor == "" {
        return defaultValue
    }
    num, _ := strconv.Atoi(valor)
    return num
}
```

---

## Resumen: Checklist de Observabilidad

- [ ] ¿Logs estructurados en JSON?
- [ ] ¿Contexto (requestID) en cada log?
- [ ] ¿Configuración en variables de entorno?
- [ ] ¿Health endpoint (/health)?
- [ ] ¿Métricas básicas (requests, errores)?
- [ ] ¿Logs rotativos para no saturar disco?
- [ ] ¿Centralizador de logs (CloudWatch, ELK)?
- [ ] ¿Alertas en errores frecuentes?
- [ ] ¿Acceso a métricas (/metrics)?
- [ ] ¿Sin información sensible en logs?

**Si cumples estos 10 puntos, tu app es observable en producción.** 📊
