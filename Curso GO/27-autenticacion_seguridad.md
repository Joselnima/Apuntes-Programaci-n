# Autenticación y Seguridad en Go - Guía Profesional

> **Cómo proteger aplicaciones, autenticar usuarios y evitar vulnerabilidades comunes en producción**

---

## ¿Por qué existe la seguridad?

Imagina que tu aplicación maneja datos de usuarios: nombres, emails, contraseñas, información bancaria.

**Sin seguridad:**
- ❌ Alguien accede a la BD y roba todas las contraseñas en texto plano
- ❌ Un atacante inyecta SQL malicioso en un formulario
- ❌ Un bot hace mil solicitudes por segundo derribando tu servidor
- ❌ Alguien falsifica un token y accede como admin
- ❌ La conexión HTTP se intercepta y leen los datos

**Con seguridad:**
- ✅ Contraseñas hasheadas con bcrypt (imposible recuperar)
- ✅ Queries preparadas contra SQL injection
- ✅ Rate limiting bloqueaía los bots
- ✅ JWT firmado criptográficamente (imposible falsificar)
- ✅ HTTPS encripta toda comunicación

La **seguridad no es opcional en producción**, es obligatoria.

---

## Parte I: Hashing de Contraseñas - El Fundamento

### El problema: Almacenar contraseñas

```go
// ❌ NUNCA hagas esto
usuario.Contraseña = "miContra123"    // Almacenan texto plano
```

Si alguien accede a la BD:
```
Juan | miContra123
María | qwerty
Carlos | password123
```

**Todos compromedidos.**

---

### ✅ La solución: bcrypt

`bcrypt` es una función de hashing que:
- 🔒 Es unidireccional (no se puede deshacer)
- 💪 Es lenta (dificulta ataques de fuerza bruta)
- 🧂 Tiene "salt" (imposible usar rainbow tables)

**Instalación:**
```bash
go get golang.org/x/crypto/bcrypt
```

---

### Hashear una contraseña

```go
package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    contraseña := "miContra123"
    
    // Hashear
    hash, err := bcrypt.GenerateFromPassword([]byte(contraseña), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Println("Hash:", string(hash))
    // Output: $2a$10$N9qo8uLOickgx2ZMRZoMye... (diferente cada vez)
}
```

**¿Qué pasó?**
- `bcrypt.GenerateFromPassword()` hashea la contraseña
- `bcrypt.DefaultCost` es el costo computacional (más alto = más lento = más seguro)
- El resultado es diferente cada vez (por el salt)

---

### Verificar una contraseña

```go
contraseña := "miContra123"
hashGuardado := "$2a$10$N9qo8uLOickgx2ZMRZoMye..."  // De la BD

// Verificar
err := bcrypt.CompareHashAndPassword([]byte(hashGuardado), []byte(contraseña))
if err != nil {
    fmt.Println("Contraseña incorrecta")
    return
}

fmt.Println("Contraseña correcta!")
```

**¿Cómo funciona?**
- `bcrypt.CompareHashAndPassword()` rehashea la contraseña con el salt del hash guardado
- Si el resultado coincide, la contraseña es correcta
- Si no coincide, es incorrecta

---

### Patrón: Registro de usuario

```go
const (
    costoBcrypt = bcrypt.DefaultCost  // O bcrypt.DefaultCost + 2 para más seguridad
)

func registrarUsuario(nombre, email, contraseña string) error {
    // 1. Validar
    if contraseña == "" || len(contraseña) < 8 {
        return fmt.Errorf("contraseña debe tener al menos 8 caracteres")
    }
    
    // 2. Hashear
    hash, err := bcrypt.GenerateFromPassword([]byte(contraseña), costoBcrypt)
    if err != nil {
        return fmt.Errorf("error al hashear: %w", err)
    }
    
    // 3. Guardar en BD
    usuario := Usuario{
        Nombre:     nombre,
        Email:      email,
        HashPass:   string(hash),  // Guardar el hash, NUNCA la contraseña
    }
    
    err = db.Create(&usuario).Error
    if err != nil {
        return fmt.Errorf("error al guardar: %w", err)
    }
    
    return nil
}
```

---

### Patrón: Login de usuario

```go
func login(email, contraseña string) (*Usuario, error) {
    // 1. Buscar usuario
    var usuario Usuario
    err := db.Where("email = ?", email).First(&usuario).Error
    if err != nil {
        return nil, fmt.Errorf("usuario no encontrado")
    }
    
    // 2. Verificar contraseña
    err = bcrypt.CompareHashAndPassword(
        []byte(usuario.HashPass),
        []byte(contraseña),
    )
    if err != nil {
        return nil, fmt.Errorf("contraseña incorrecta")
    }
    
    // 3. Retornar usuario
    return &usuario, nil
}
```

---

## Parte II: JWT - Token de Autenticación

### ¿Qué es JWT?

JWT (JSON Web Token) es una forma segura de:
1. **Verficar identidad** - Probar que eres quién dices ser
2. **Autorizar recursos** - Decir "este usuario puede acceder a X"
3. **Mantener estado sin sesión** - Sin guardar en servidor

**Estructura:**
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U
```

Tres partes separadas por `.`:
1. **Header** - Tipo de token y algoritmo
2. **Payload** - Datos (claims)
3. **Signature** - Firma criptográfica

---

### Generar JWT

```go
import "github.com/golang-jwt/jwt/v5"

const SecretoJWT = "tu-secreto-super-seguro-muy-largo"  // En producción: env variable

type Claims struct {
    UserID   uint
    Email    string
    jwt.RegisteredClaims
}

func generarToken(userID uint, email string) (string, error) {
    // 1. Crear claims (datos)
    claims := Claims{
        UserID: userID,
        Email:  email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),  // Expira en 24 horas
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
            Issuer:    "miapp",
        },
    }
    
    // 2. Crear token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    // 3. Firmar
    tokenString, err := token.SignedString([]byte(SecretoJWT))
    if err != nil {
        return "", err
    }
    
    return tokenString, nil
}
```

---

### Validar JWT

```go
func validarToken(tokenString string) (*Claims, error) {
    claims := &Claims{}
    
    // Parsear y validar
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        // Verificar algoritmo
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("algoritmo inesperado: %v", token.Header["alg"])
        }
        return []byte(SecretoJWT), nil
    })
    
    if err != nil {
        return nil, fmt.Errorf("error al parsear token: %w", err)
    }
    
    if !token.Valid {
        return nil, fmt.Errorf("token inválido")
    }
    
    return claims, nil
}
```

---

### Patrón: Middleware de autenticación

```go
func middlewareJWT(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 1. Obtener token del header
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "sin token", http.StatusUnauthorized)
            return
        }
        
        // Formato: "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "formato inválido", http.StatusUnauthorized)
            return
        }
        
        tokenString := parts[1]
        
        // 2. Validar token
        claims, err := validarToken(tokenString)
        if err != nil {
            http.Error(w, "token inválido", http.StatusUnauthorized)
            return
        }
        
        // 3. Guardar en contexto
        ctx := context.WithValue(r.Context(), "userID", claims.UserID)
        ctx = context.WithValue(ctx, "email", claims.Email)
        
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

---

### Uso en endpoint

```go
func handleGetPerfil(w http.ResponseWriter, r *http.Request) {
    // Obtener del contexto
    userID := r.Context().Value("userID").(uint)
    email := r.Context().Value("email").(string)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "userID": userID,
        "email":  email,
    })
}

func main() {
    mux := http.NewServeMux()
    
    // Proteger con middleware
    mux.Handle("/perfil", middlewareJWT(http.HandlerFunc(handleGetPerfil)))
    
    http.ListenAndServe(":8080", mux)
}
```

---

## Parte III: Validación de Entradas - Prevenir Ataques

### ❌ SQL Injection

```go
// MALO: Concatenar strings
email := "juan@mail.com' OR '1'='1"
query := fmt.Sprintf("SELECT * FROM usuarios WHERE email = '%s'", email)
// Query resultante: SELECT * FROM usuarios WHERE email = 'juan@mail.com' OR '1'='1'
// ¡Devuelve TODOS los usuarios!

// BIEN: Usar placeholders (parámetros preparados)
var usuario Usuario
db.Where("email = ?", email).First(&usuario)  // El ? es reemplazado de forma segura
```

---

### Validar Emails

```go
import "regexp"

func validarEmail(email string) bool {
    // Patrón simple
    pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    matched, _ := regexp.MatchString(pattern, email)
    return matched
}

// Uso
if !validarEmail(email) {
    fmt.Println("Email inválido")
}
```

---

### Validar Contraseña (Fuerza)

```go
func validarContraseña(p string) error {
    if len(p) < 8 {
        return fmt.Errorf("debe tener al menos 8 caracteres")
    }
    
    tieneNumero := false
    tieneMayuscula := false
    tieneMinuscula := false
    tieneespecial := false
    
    for _, c := range p {
        switch {
        case c >= '0' && c <= '9':
            tieneNumero = true
        case c >= 'A' && c <= 'Z':
            tieneMayuscula = true
        case c >= 'a' && c <= 'z':
            tieneMinuscula = true
        case string(c) == "!@#$%^&*()-_=+":
            tieneespecial = true
        }
    }
    
    if !tieneNumero {
        return fmt.Errorf("debe contener números")
    }
    if !tieneMayuscula {
        return fmt.Errorf("debe contener mayúsculas")
    }
    if !tieneMinuscula {
        return fmt.Errorf("debe contener minúsculas")
    }
    
    return nil
}
```

---

## Parte IV: Rate Limiting - Proteger contra Bots

### Limitador simple por IP

```go
import "sync"

type RateLimiter struct {
    mu       sync.Mutex
    visitantes map[string][]time.Time
}

func (rl *RateLimiter) Permitido(ip string, maxSolicitudes int, duracion time.Duration) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    ahora := time.Now()
    limiteT := ahora.Add(-duracion)
    
    // Limpiar solicitudes antiguas
    rl.visitantes[ip] = filtrarAntiguas(rl.visitantes[ip], limiteT)
    
    // Si menos solicitudes que el límite
    if len(rl.visitantes[ip]) < maxSolicitudes {
        rl.visitantes[ip] = append(rl.visitantes[ip], ahora)
        return true
    }
    
    return false
}

func filtrarAntiguas(times []time.Time, limite time.Time) []time.Time {
    var resultado []time.Time
    for _, t := range times {
        if t.After(limite) {
            resultado = append(resultado, t)
        }
    }
    return resultado
}
```

---

### Middleware de Rate Limiting

```go
var limitador = &RateLimiter{
    visitantes: make(map[string][]time.Time),
}

func middlewareRateLimit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr
        
        // 10 solicitudes por minuto
        if !limitador.Permitido(ip, 10, time.Minute) {
            http.Error(w, "demasiadas solicitudes", http.StatusTooManyRequests)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}
```

---

## Parte V: CORS - Control de Acceso

### Habilitar CORS correctamente

```go
func middlewareCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Permitir origen (en producción: desde variable de entorno)
        w.Header().Set("Access-Control-Allow-Origin", "https://tudominio.com")
        
        // Métodos permitidos
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        
        // Headers permitidos
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        // Credenciales
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        
        // Tiempo de caché (segundos)
        w.Header().Set("Access-Control-Max-Age", "3600")
        
        // Manejar preflight
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}
```

---

### ⚠️ NUNCA hagas esto

```go
// ❌ INSEGURO: Permitir todos los orígenes
w.Header().Set("Access-Control-Allow-Origin", "*")
w.Header().Set("Access-Control-Allow-Credentials", "true")
// Esto es contradictorio y peligroso
```

---

## Parte VI: HTTPS - Encriptación de Conexión

### Usar HTTPS

```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handleHome)
    
    // Iniciar con TLS/SSL
    // cert.pem y key.pem generados con: go run $(go env GOROOT)/src/crypto/tls/generate_cert.go -host localhost
    err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", mux)
    if err != nil {
        log.Fatal(err)
    }
}
```

**En producción:**
- Usa certificados de Let's Encrypt (gratuitos)
- Redirige HTTP → HTTPS
- Usa HSTS headers

---

## Parte VII: Encriptación de Datos Sensibles

### Encriptar datos

```go
import "crypto/aes"
import "crypto/cipher"
import "crypto/rand"

func encriptar(texto string, clave string) (string, error) {
    // Clave debe tener 16, 24 o 32 bytes
    if len(clave) != 32 {
        clave = strings.Repeat(clave, 32/len(clave)+1)[:32]
    }
    
    bloque, err := aes.NewCipher([]byte(clave))
    if err != nil {
        return "", err
    }
    
    // Generar IV
    iv := make([]byte, aes.BlockSize)
    _, err = rand.Read(iv)
    if err != nil {
        return "", err
    }
    
    // Encriptar
    stream := cipher.NewCFBEncrypter(bloque, iv)
    encrypted := make([]byte, len(texto))
    stream.XORKeyStream(encrypted, []byte(texto))
    
    // Combinar IV + encrypted
    resultado := append(iv, encrypted...)
    return base64.StdEncoding.EncodeToString(resultado), nil
}

func desencriptar(texto64 string, clave string) (string, error) {
    // Decodificar base64
    datos, err := base64.StdEncoding.DecodeString(texto64)
    if err != nil {
        return "", err
    }
    
    if len(clave) != 32 {
        clave = strings.Repeat(clave, 32/len(clave)+1)[:32]
    }
    
    bloque, err := aes.NewCipher([]byte(clave))
    if err != nil {
        return "", err
    }
    
    // Extraer IV
    iv := datos[:aes.BlockSize]
    encrypted := datos[aes.BlockSize:]
    
    // Desencriptar
    stream := cipher.NewCFBDecrypter(bloque, iv)
    decrypted := make([]byte, len(encrypted))
    stream.XORKeyStream(decrypted, encrypted)
    
    return string(decrypted), nil
}
```

---

## Parte VIII: Anti-patrones y Errores Comunes

### ❌ No guardes secretos en el código

```go
// MALO
const ApiSecret = "sk_live_51PgQR2IwVqZ9Z7z7z"

// BIEN
apiSecret := os.Getenv("API_SECRET")
```

---

### ❌ No ignores errores de seguridad

```go
// MALO
hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

// BIEN
hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
if err != nil {
    return fmt.Errorf("error al hashear: %w", err)
}
```

---

### ❌ No uses JWT sin HTTPS

```go
// MALO: Token viajando en texto plano
// http://localhost:8080/api/datos?token=asd123...

// BIEN: Token en header, sobre HTTPS
// https://api.tudominio.com/datos
// Authorization: Bearer asd123...
```

---

### ❌ No reutilices secretos

```go
// MALO: Mismo secreto para todo
const Secret = "el-mismo-secreto-para-jwt-bd-apis"

// BIEN: Secretos diferentes
jwtSecret := os.Getenv("JWT_SECRET")
dbPassword := os.Getenv("DB_PASSWORD")
apiKey := os.Getenv("EXTERNAL_API_KEY")
```

---

## Parte IX: Best Practices - Resumen

### 1. Passwords: Sempre hash con bcrypt

```go
hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
// Guardar hash, NUNCA la contraseña
```

---

### 2. Autenticación: Usa JWT para APIs

```go
// Generar en login
token, _ := generarToken(usuarioID, email)

// Validar en middleware
claims, _ := validarToken(token)
```

---

### 3. Validación: SIEMPRE valida entrada

```go
// ✅ Siempre
if err := validarEmail(email); err != nil {
    return err
}
if err := validarContraseña(pwd); err != nil {
    return err
}
```

---

### 4. Bases de datos: Parameterized queries

```go
// ✅ BIEN
db.Where("email = ?", email).First(&user)

// ❌ NUNCA
db.Where(fmt.Sprintf("email = '%s'", email)).First(&user)
```

---

### 5. Secretos: Variables de entorno

```go
jwtSecret := os.Getenv("JWT_SECRET")
if jwtSecret == "" {
    log.Fatal("JWT_SECRET no definido")
}
```

---

### 6. HTTPS: Siempre en producción

```go
// Producción
http.ListenAndServeTLS(":443", "cert.pem", "key.pem", handler)

// Desarrollo (si necesitas)
// http.ListenAndServe(":8080", handler)
```

---

### 7. Timeouts: Previene ataques de negación de servicio

```go
server := &http.Server{
    Addr:         ":8080",
    Handler:      mux,
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  60 * time.Second,
}
server.ListenAndServe()
```

---

### 8. Logging: Registra intentos fallidos

```go
// Login fallido
log.Printf("Login fallido para %s desde %s", email, ip)

// Acceso denegado
log.Printf("Acceso denegado a %s para %s", endpoint, userID)
```

---

### 9. Dependencias: Mantén actualizadas

```bash
go get -u
go mod tidy
```

---

### 10. Tests: Prueba la seguridad

```go
func TestHasheoDiferente(t *testing.T) {
    pwd := "miContra123"
    hash1, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
    hash2, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
    
    if string(hash1) == string(hash2) {
        t.Error("Hashes diferentes pero iguales (sin salt)")
    }
}
```

---

## Resumen Final: Checklist de Seguridad

- [ ] ¿Contraseñas hasheadas con bcrypt?
- [ ] ¿JWT para autenticación de APIs?
- [ ] ¿HTTPS con certificado válido?
- [ ] ¿Validación de entrada SIEMPRE?
- [ ] ¿Queries preparadas (no SQL injection)?
- [ ] ¿Rate limiting contra bots?
- [ ] ¿CORS configurado correctamente?
- [ ] ¿Secretos en variables de entorno?
- [ ] ¿Logging de eventos de seguridad?
- [ ] ¿Dependencias actualizadas?

**Si cumples estos 10 puntos, tu aplicación es segura en producción.** 🔒
