# Módulo 21 - Writing Modes

En este módulo aprenderás sobre CSS Writing Modes, una característica avanzada que permite controlar la dirección del texto y el flujo de los elementos, incluyendo modos de escritura vertical, horizontal, y direccionales para diferentes idiomas.

## Introducción a CSS Writing Modes

### ¿Qué son los Writing Modes?

CSS Writing Modes definen cómo fluye el texto en un documento, permitiendo layouts verticales, horizontales, y direccionales. Esto es crucial para idiomas que se escriben de derecha a izquierda (RTL) o verticalmente.

```css
/* Modo de escritura horizontal (por defecto) */
.element {
  writing-mode: horizontal-tb;
}

/* Modo de escritura vertical */
.element {
  writing-mode: vertical-rl;
}

/* Modo de escritura vertical (izquierda a derecha) */
.element {
  writing-mode: vertical-lr;
}
```

## Propiedades de Writing Modes

### Writing-mode

```css
/* Horizontal top-to-bottom (por defecto) */
.horizontal-tb {
  writing-mode: horizontal-tb;
}

/* Vertical right-to-left */
.vertical-rl {
  writing-mode: vertical-rl;
}

/* Vertical left-to-right */
.vertical-lr {
  writing-mode: vertical-lr;
}
```

### Direction y Text-orientation

```css
/* Dirección del texto */
.rtl-text {
  direction: rtl;
}

.ltr-text {
  direction: ltr;
}

/* Orientación del texto en modos verticales */
.text-upright {
  text-orientation: upright;
}

.text-mixed {
  text-orientation: mixed;
}

.text-sideways {
  text-orientation: sideways;
}
```

## Ejemplo Práctico Básico

```html
<!DOCTYPE html>
<html lang="es" dir="ltr">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Writing Modes - Básico</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: 'Arial', sans-serif;
      background: #f5f5f7;
      color: #333;
      padding: 20px;
    }

    .container {
      max-width: 1200px;
      margin: 0 auto;
    }

    .header {
      text-align: center;
      margin-bottom: 40px;
    }

    .header h1 {
      font-size: 2.5rem;
      margin-bottom: 10px;
      color: #333;
    }

    .header p {
      font-size: 1.1rem;
      color: #666;
    }

    .demo-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
      gap: 30px;
      margin-bottom: 50px;
    }

    .demo-item {
      background: white;
      border-radius: 12px;
      padding: 25px;
      box-shadow: 0 4px 20px rgba(0,0,0,0.1);
      border: 1px solid #e0e0e0;
    }

    .demo-item h3 {
      margin-bottom: 20px;
      color: #333;
      font-size: 1.3rem;
      text-align: center;
    }

    .writing-mode-demo {
      height: 200px;
      border: 2px solid #667eea;
      border-radius: 8px;
      padding: 15px;
      margin-bottom: 15px;
      background: linear-gradient(45deg, #f8f9fa, #e9ecef);
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 16px;
      font-weight: bold;
      color: #333;
    }

    /* Horizontal top-to-bottom */
    .horizontal-tb {
      writing-mode: horizontal-tb;
    }

    /* Vertical right-to-left */
    .vertical-rl {
      writing-mode: vertical-rl;
      text-orientation: mixed;
    }

    /* Vertical left-to-right */
    .vertical-lr {
      writing-mode: vertical-lr;
      text-orientation: mixed;
    }

    .code-example {
      background: #2d3748;
      color: #e2e8f0;
      padding: 15px;
      border-radius: 6px;
      font-family: 'Monaco', 'Menlo', monospace;
      font-size: 14px;
      margin-top: 15px;
    }

    .direction-demo {
      margin-top: 50px;
      background: white;
      border-radius: 12px;
      padding: 30px;
      box-shadow: 0 4px 20px rgba(0,0,0,0.1);
    }

    .direction-demo h2 {
      text-align: center;
      margin-bottom: 30px;
      color: #333;
      font-size: 2rem;
    }

    .direction-examples {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 30px;
    }

    .direction-item {
      padding: 20px;
      border-radius: 8px;
      border: 2px solid #e0e0e0;
    }

    .ltr-example {
      direction: ltr;
      background: linear-gradient(45deg, #e3f2fd, #bbdefb);
    }

    .rtl-example {
      direction: rtl;
      background: linear-gradient(45deg, #fce4ec, #f8bbd9);
    }

    .direction-item h4 {
      margin-bottom: 15px;
      color: #333;
      font-size: 1.1rem;
    }

    .direction-item p {
      margin-bottom: 10px;
      line-height: 1.6;
      color: #666;
    }

    .toggle-direction {
      text-align: center;
      margin-top: 30px;
    }

    .toggle-btn {
      background: #667eea;
      color: white;
      border: none;
      padding: 12px 24px;
      border-radius: 6px;
      cursor: pointer;
      font-size: 16px;
      transition: background-color 0.2s ease;
    }

    .toggle-btn:hover {
      background: #5a67d8;
    }

    .text-orientation-demo {
      margin-top: 50px;
      background: white;
      border-radius: 12px;
      padding: 30px;
      box-shadow: 0 4px 20px rgba(0,0,0,0.1);
    }

    .text-orientation-demo h2 {
      text-align: center;
      margin-bottom: 30px;
      color: #333;
      font-size: 2rem;
    }

    .orientation-grid {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
      gap: 20px;
    }

    .orientation-item {
      height: 150px;
      border: 2px solid #667eea;
      border-radius: 8px;
      padding: 15px;
      background: linear-gradient(45deg, #f8f9fa, #e9ecef);
      writing-mode: vertical-rl;
      display: flex;
      align-items: center;
      justify-content: center;
      text-align: center;
    }

    .upright {
      text-orientation: upright;
    }

    .mixed {
      text-orientation: mixed;
    }

    .sideways {
      text-orientation: sideways;
    }

    .orientation-label {
      position: absolute;
      bottom: 10px;
      left: 10px;
      background: rgba(0,0,0,0.7);
      color: white;
      padding: 5px 10px;
      border-radius: 4px;
      font-size: 12px;
    }

    .orientation-item {
      position: relative;
    }
  </style>
</head>
<body>
  <div class="container">
    <header class="header">
      <h1>CSS Writing Modes</h1>
      <p>Controlando la dirección y orientación del texto</p>
    </header>

    <div class="demo-grid">
      <div class="demo-item">
        <h3>Horizontal TB</h3>
        <div class="writing-mode-demo horizontal-tb">
          Este texto fluye horizontalmente de izquierda a derecha y de arriba hacia abajo (comportamiento por defecto).
        </div>
        <div class="code-example">
writing-mode: horizontal-tb;
        </div>
      </div>

      <div class="demo-item">
        <h3>Vertical RL</h3>
        <div class="writing-mode-demo vertical-rl">
          Este texto fluye verticalmente de arriba hacia abajo y de derecha a izquierda.
        </div>
        <div class="code-example">
writing-mode: vertical-rl;<br>
text-orientation: mixed;
        </div>
      </div>

      <div class="demo-item">
        <h3>Vertical LR</h3>
        <div class="writing-mode-demo vertical-lr">
          Este texto fluye verticalmente de arriba hacia abajo y de izquierda a derecha.
        </div>
        <div class="code-example">
writing-mode: vertical-lr;<br>
text-orientation: mixed;
        </div>
      </div>
    </div>

    <div class="direction-demo">
      <h2>Dirección del Texto</h2>

      <div class="direction-examples">
        <div class="direction-item ltr-example">
          <h4>Left-to-Right (LTR)</h4>
          <p>Esta es la dirección por defecto para la mayoría de los idiomas occidentales. El texto fluye de izquierda a derecha.</p>
          <p>Hello World! Cómo estás hoy?</p>
        </div>

        <div class="direction-item rtl-example">
          <h4>Right-to-Left (RTL)</h4>
          <p>Esta dirección se usa para idiomas como el árabe y el hebreo. El texto fluye de derecha a izquierda.</p>
          <p>مرحبا بالعالم! كيف حالك اليوم؟</p>
        </div>
      </div>

      <div class="toggle-direction">
        <button class="toggle-btn" onclick="toggleDirection()">Cambiar Dirección del Documento</button>
      </div>
    </div>

    <div class="text-orientation-demo">
      <h2>Text Orientation</h2>

      <div class="orientation-grid">
        <div class="orientation-item upright">
          <div class="orientation-label">Upright</div>
          Texto vertical con caracteres erguidos
        </div>

        <div class="orientation-item mixed">
          <div class="orientation-label">Mixed</div>
          Texto vertical con orientación mixta
        </div>

        <div class="orientation-item sideways">
          <div class="orientation-label">Sideways</div>
          Texto vertical rotado 90 grados
        </div>
      </div>
    </div>
  </div>

  <script>
    function toggleDirection() {
      const html = document.documentElement;
      const currentDir = html.getAttribute('dir');
      const newDir = currentDir === 'ltr' ? 'rtl' : 'ltr';
      html.setAttribute('dir', newDir);

      const button = document.querySelector('.toggle-btn');
      button.textContent = `Cambiar Dirección del Documento (${newDir.toUpperCase()})`;
    }
  </script>
</body>
</html>
```

## Layouts con Writing Modes

### Diseño vertical para idiomas asiáticos

```css
/* Layout vertical para japonés/chino */
.vertical-layout {
  writing-mode: vertical-rl;
  text-orientation: mixed;
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.vertical-header {
  writing-mode: horizontal-tb;
  text-align: center;
  margin-bottom: 20px;
}

.vertical-content {
  flex: 1;
  padding: 20px;
}
```

### Sidebar vertical

```css
.vertical-sidebar {
  writing-mode: vertical-rl;
  text-orientation: mixed;
  width: 60px;
  height: 100vh;
  background: #333;
  color: white;
  padding: 10px;
}

.vertical-nav-item {
  margin-bottom: 20px;
  transform: rotate(180deg); /* Para leer correctamente */
}
```

## Ejemplo Complejo: Sitio Multilingüe

```html
<!DOCTYPE html>
<html lang="es" dir="ltr">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CSS Writing Modes - Multilingüe</title>
  <style>
    * {
      margin: 0;
      padding: 0;
      box-sizing: border-box;
    }

    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Noto Sans CJK JP', sans-serif;
      background: #1a1a1a;
      color: #e0e0e0;
      transition: all 0.3s ease;
    }

    .app {
      display: flex;
      min-height: 100vh;
    }

    .sidebar {
      width: 280px;
      background: linear-gradient(180deg, #2d3748 0%, #1a202c 100%);
      padding: 30px 20px;
      box-shadow: 2px 0 10px rgba(0,0,0,0.3);
    }

    .sidebar-header {
      margin-bottom: 30px;
    }

    .sidebar-title {
      font-size: 1.5rem;
      font-weight: bold;
      color: #4ecdc4;
      margin-bottom: 10px;
    }

    .sidebar-subtitle {
      color: #a0aec0;
      font-size: 0.9rem;
    }

    .language-selector {
      margin-bottom: 30px;
    }

    .language-btn {
      display: block;
      width: 100%;
      padding: 10px 15px;
      margin-bottom: 8px;
      background: rgba(78, 205, 196, 0.1);
      border: 1px solid rgba(78, 205, 196, 0.3);
      color: #4ecdc4;
      border-radius: 6px;
      cursor: pointer;
      transition: all 0.2s ease;
    }

    .language-btn:hover,
    .language-btn.active {
      background: rgba(78, 205, 196, 0.2);
      border-color: #4ecdc4;
    }

    .nav-item {
      display: block;
      color: #e2e8f0;
      text-decoration: none;
      padding: 12px 16px;
      margin-bottom: 8px;
      border-radius: 8px;
      transition: all 0.2s ease;
    }

    .nav-item:hover {
      background: rgba(78, 205, 196, 0.1);
      color: #4ecdc4;
    }

    .nav-item.active {
      background: rgba(78, 205, 196, 0.2);
      color: #4ecdc4;
    }

    .main-content {
      flex: 1;
      padding: 30px;
      overflow-y: auto;
    }

    .content-header {
      margin-bottom: 30px;
    }

    .page-title {
      font-size: 2rem;
      font-weight: bold;
      color: #e0e0e0;
      margin-bottom: 10px;
    }

    .page-subtitle {
      color: #a0aec0;
      font-size: 1.1rem;
    }

    .content-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 25px;
      margin-bottom: 50px;
    }

    .content-card {
      background: rgba(255,255,255,0.05);
      border-radius: 12px;
      padding: 25px;
      border: 1px solid rgba(255,255,255,0.1);
      transition: all 0.3s ease;
    }

    .content-card:hover {
      transform: translateY(-5px);
      border-color: rgba(78, 205, 196, 0.3);
    }

    .card-title {
      font-size: 1.3rem;
      font-weight: bold;
      margin-bottom: 15px;
      color: #e0e0e0;
    }

    .card-text {
      color: #a0aec0;
      line-height: 1.6;
    }

    /* Modos de escritura específicos por idioma */
    .writing-mode-ltr {
      writing-mode: horizontal-tb;
      direction: ltr;
    }

    .writing-mode-rtl {
      writing-mode: horizontal-tb;
      direction: rtl;
    }

    .writing-mode-vertical {
      writing-mode: vertical-rl;
      text-orientation: mixed;
    }

    /* Contenido específico por idioma */
    .content-es {
      font-family: 'Arial', sans-serif;
    }

    .content-ar {
      font-family: 'Arial', sans-serif;
      direction: rtl;
    }

    .content-ja {
      font-family: 'Noto Sans CJK JP', sans-serif;
      writing-mode: vertical-rl;
      text-orientation: mixed;
    }

    .content-zh {
      font-family: 'Noto Sans CJK SC', sans-serif;
      writing-mode: vertical-rl;
      text-orientation: mixed;
    }

    /* Layout vertical para idiomas asiáticos */
    .vertical-layout .main-content {
      writing-mode: vertical-rl;
      text-orientation: mixed;
    }

    .vertical-layout .content-grid {
      display: flex;
      flex-direction: column;
      gap: 25px;
    }

    .vertical-layout .content-card {
      writing-mode: horizontal-tb;
      text-orientation: upright;
    }

    /* Transiciones suaves */
    .app {
      transition: all 0.3s ease;
    }

    /* Responsive */
    @media (max-width: 768px) {
      .app {
        flex-direction: column;
      }

      .sidebar {
        width: 100%;
        order: -1;
      }

      .vertical-layout .main-content {
        writing-mode: horizontal-tb;
      }

      .vertical-layout .content-grid {
        display: grid;
        grid-template-columns: 1fr;
      }
    }
  </style>
</head>
<body>
  <div class="app" id="app">
    <aside class="sidebar">
      <div class="sidebar-header">
        <h1 class="sidebar-title">🌐 MultiLang</h1>
        <p class="sidebar-subtitle">Sitio multilingüe</p>
      </div>

      <div class="language-selector">
        <button class="language-btn active" data-lang="es">Español</button>
        <button class="language-btn" data-lang="ar">العربية</button>
        <button class="language-btn" data-lang="ja">日本語</button>
        <button class="language-btn" data-lang="zh">中文</button>
      </div>

      <nav>
        <a href="#" class="nav-item active">🏠 Inicio</a>
        <a href="#" class="nav-item">📖 Acerca de</a>
        <a href="#" class="nav-item">📞 Contacto</a>
        <a href="#" class="nav-item">⚙️ Configuración</a>
      </nav>
    </aside>

    <main class="main-content writing-mode-ltr" id="mainContent">
      <header class="content-header">
        <h1 class="page-title" id="pageTitle">Bienvenido</h1>
        <p class="page-subtitle" id="pageSubtitle">Explora diferentes modos de escritura</p>
      </header>

      <div class="content-grid" id="contentGrid">
        <div class="content-card content-es">
          <h3 class="card-title">Diseño Responsivo</h3>
          <p class="card-text">Los writing modes permiten crear interfaces que se adaptan automáticamente a diferentes direcciones de escritura y orientaciones de texto.</p>
        </div>

        <div class="content-card content-es">
          <h3 class="card-title">Internacionalización</h3>
          <p class="card-text">Esencial para crear aplicaciones web que funcionen correctamente en múltiples idiomas y culturas alrededor del mundo.</p>
        </div>

        <div class="content-card content-es">
          <h3 class="card-title">Compatibilidad</h3>
          <p class="card-text">Soportado en todos los navegadores modernos, permitiendo experiencias consistentes en diferentes plataformas.</p>
        </div>
      </div>
    </main>
  </div>

  <script>
    const languages = {
      es: {
        title: 'Bienvenido',
        subtitle: 'Explora diferentes modos de escritura',
        cards: [
          {
            title: 'Diseño Responsivo',
            text: 'Los writing modes permiten crear interfaces que se adaptan automáticamente a diferentes direcciones de escritura y orientaciones de texto.'
          },
          {
            title: 'Internacionalización',
            text: 'Esencial para crear aplicaciones web que funcionen correctamente en múltiples idiomas y culturas alrededor del mundo.'
          },
          {
            title: 'Compatibilidad',
            text: 'Soportado en todos los navegadores modernos, permitiendo experiencias consistentes en diferentes plataformas.'
          }
        ]
      },
      ar: {
        title: 'مرحباً',
        subtitle: 'استكشف أوضاع الكتابة المختلفة',
        cards: [
          {
            title: 'التصميم المتجاوب',
            text: 'تسمح أوضاع الكتابة بإنشاء واجهات تتكيف تلقائياً مع اتجاهات الكتابة واتجاهات النص المختلفة.'
          },
          {
            title: 'التدويل',
            text: 'أساسي لإنشاء تطبيقات الويب التي تعمل بشكل صحيح في لغات وثقافات متعددة حول العالم.'
          },
          {
            title: 'التوافق',
            text: 'مدعوم في جميع المتصفحات الحديثة، مما يسمح بتجارب متسقة على منصات مختلفة.'
          }
        ]
      },
      ja: {
        title: 'ようこそ',
        subtitle: 'さまざまな書き込みモードを探索する',
        cards: [
          {
            title: 'レスポンシブデザイン',
            text: 'ライティングモードにより、さまざまな書き込み方向やテキスト方向に自動的に適応するインターフェースを作成できます。'
          },
          {
            title: '国際化',
            text: '世界中の複数の言語や文化で正しく動作するWebアプリケーションを作成するために不可欠です。'
          },
          {
            title: '互換性',
            text: '最新のブラウザすべてでサポートされており、さまざまなプラットフォームで一貫したエクスペリエンスを提供します。'
          }
        ]
      },
      zh: {
        title: '欢迎',
        subtitle: '探索不同的书写模式',
        cards: [
          {
            title: '响应式设计',
            text: '书写模式允许创建自动适应不同书写方向和文本方向的界面。'
          },
          {
            title: '国际化',
            text: '对于创建在世界各地多种语言和文化中正常工作的Web应用程序至关重要。'
          },
          {
            title: '兼容性',
            text: '在所有现代浏览器中受支持，允许在不同平台上提供一致的体验。'
          }
        ]
      }
    };

    function setLanguage(lang) {
      const app = document.getElementById('app');
      const mainContent = document.getElementById('mainContent');
      const pageTitle = document.getElementById('pageTitle');
      const pageSubtitle = document.getElementById('pageSubtitle');
      const contentGrid = document.getElementById('contentGrid');

      // Actualizar botones de idioma
      document.querySelectorAll('.language-btn').forEach(btn => {
        btn.classList.remove('active');
      });
      document.querySelector(`[data-lang="${lang}"]`).classList.add('active');

      // Aplicar modo de escritura
      app.className = 'app';
      mainContent.className = 'main-content';

      if (lang === 'ar') {
        document.documentElement.setAttribute('dir', 'rtl');
        mainContent.classList.add('writing-mode-rtl');
      } else if (lang === 'ja' || lang === 'zh') {
        document.documentElement.setAttribute('dir', 'ltr');
        app.classList.add('vertical-layout');
      } else {
        document.documentElement.setAttribute('dir', 'ltr');
        mainContent.classList.add('writing-mode-ltr');
      }

      // Actualizar contenido
      const content = languages[lang];
      pageTitle.textContent = content.title;
      pageSubtitle.textContent = content.subtitle;

      contentGrid.innerHTML = content.cards.map(card => `
        <div class="content-card content-${lang}">
          <h3 class="card-title">${card.title}</h3>
          <p class="card-text">${card.text}</p>
        </div>
      `).join('');
    }

    // Event listeners
    document.querySelectorAll('.language-btn').forEach(btn => {
      btn.addEventListener('click', () => {
        setLanguage(btn.dataset.lang);
      });
    });

    // Inicializar con español
    setLanguage('es');
  </script>
</body>
</html>
```

## Técnicas Avanzadas

### Combinando con CSS Grid y Flexbox

```css
/* Grid con writing modes */
.grid-vertical {
  writing-mode: vertical-rl;
  display: grid;
  grid-template-columns: 1fr 2fr 1fr;
  gap: 20px;
}

/* Flexbox con dirección RTL */
.flex-rtl {
  direction: rtl;
  display: flex;
  justify-content: flex-start;
}
```

### Animaciones con writing modes

```css
.writing-transition {
  transition: writing-mode 0.5s ease;
}

.writing-transition:hover {
  writing-mode: vertical-rl;
}
```

## Casos de Uso Prácticos

### 1. Diseño de libros verticales

```css
.book-layout {
  writing-mode: vertical-rl;
  text-orientation: mixed;
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.book-page {
  writing-mode: horizontal-tb;
  padding: 40px;
}
```

### 2. Navegación vertical

```css
.vertical-nav {
  writing-mode: vertical-rl;
  text-orientation: mixed;
  display: flex;
  flex-direction: row;
}

.nav-item {
  margin-inline-end: 20px;
}
```

### 3. Interfaces multilingües

```css
/* Automático basado en el atributo lang */
[lang^="ar"] {
  direction: rtl;
}

[lang^="ja"] {
  writing-mode: vertical-rl;
  text-orientation: mixed;
}
```

## Mejores Prácticas

### 1. Usar writing modes apropiadamente

```css
/* ✅ Bueno: usar para idiomas verticales */
.japanese-content {
  writing-mode: vertical-rl;
  text-orientation: mixed;
}

/* ❌ Evitar: no usar solo por estilo */
.styled-vertical {
  writing-mode: vertical-rl; /* Solo si es necesario */
}
```

### 2. Considerar el rendimiento

```css
/* Writing modes pueden afectar el rendimiento */
.writing-mode-change {
  will-change: writing-mode;
}
```

### 3. Testing en múltiples idiomas

```css
/* Asegurarse de que funciona en diferentes direcciones */
.test-ltr { direction: ltr; }
.test-rtl { direction: rtl; }
.test-vertical { writing-mode: vertical-rl; }
```

## Compatibilidad

### Fallbacks

```css
.element {
  /* Fallback para navegadores antiguos */
  direction: ltr;

  /* Writing modes modernos */
  writing-mode: horizontal-tb;
}

@supports not (writing-mode: vertical-rl) {
  .fallback-layout {
    /* Layout alternativo */
    flex-direction: row;
  }
}
```

## Resumen

CSS Writing Modes revolucionan el diseño web internacional:

- ✅ **Writing-mode**: Control del flujo de texto (horizontal-tb, vertical-rl, vertical-lr)
- ✅ **Direction**: Control de dirección del texto (ltr, rtl)
- ✅ **Text-orientation**: Orientación de caracteres en modos verticales
- ✅ **Internacionalización**: Soporte nativo para múltiples idiomas
- ✅ **Layouts flexibles**: Combinación con Grid y Flexbox
- ✅ **Compatibilidad**: Excelente soporte en navegadores modernos

Los Writing Modes son esenciales para crear experiencias web verdaderamente globales que respeten las convenciones de escritura de diferentes culturas e idiomas.