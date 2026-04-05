# Capítulo 27: Proyecto CRUD Completo - Construyendo una Biblioteca Digital

## La Analogía de la Biblioteca

Imagina que eres el bibliotecario de una biblioteca digital moderna. Tu trabajo es:

- 📚 **Catalogar libros**: Crear nuevos registros en el sistema
- 🔍 **Buscar libros**: Leer y mostrar información de los libros
- ✏️ **Actualizar catálogos**: Modificar información de libros existentes
- 🗑️ **Retirar libros**: Eliminar libros del catálogo cuando sea necesario

En el mundo del desarrollo web, esto se conoce como **CRUD** (Create, Read, Update, Delete). Es el patrón fundamental que usarás en casi todas las aplicaciones que construyas.

En este capítulo, construiremos juntos una **Biblioteca Digital** completa, integrando todo lo que hemos aprendido: formularios reactivos, servicios HTTP, routing, autenticación, guards, interceptores, y manejo de estado. Será como construir una biblioteca real, pero en el mundo digital.

---

## Arquitectura del Proyecto: Planificando Nuestra Biblioteca

Antes de escribir código, vamos a diseñar la arquitectura de nuestra aplicación. Una buena planificación es como el plano de una casa: te ahorra muchos problemas después.

### Estructura de la Aplicación

```
biblioteca-digital/
├── app/
│   ├── core/
│   │   ├── guards/
│   │   ├── interceptors/
│   │   └── services/
│   ├── features/
│   │   ├── auth/
│   │   └── books/
│   │       ├── components/
│   │       ├── services/
│   │       └── models/
│   ├── shared/
│   │   ├── components/
│   │   └── pipes/
│   └── app-routing.module.ts
├── assets/
└── environments/
```

### Modelo de Datos: El Libro

```typescript
// models/book.model.ts
export interface Book {
  id?: number;
  title: string;
  author: string;
  isbn: string;
  publishedYear: number;
  genre: string;
  description: string;
  coverImage?: string;
  availableCopies: number;
  totalCopies: number;
  createdAt?: Date;
  updatedAt?: Date;
}

export interface CreateBookRequest {
  title: string;
  author: string;
  isbn: string;
  publishedYear: number;
  genre: string;
  description: string;
  totalCopies: number;
}

export interface UpdateBookRequest extends Partial<CreateBookRequest> {
  availableCopies?: number;
}

// Tipos para filtros y búsqueda
export interface BookFilters {
  title?: string;
  author?: string;
  genre?: string;
  availableOnly?: boolean;
}

export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  limit: number;
  totalPages: number;
}
```

---

## Paso 1: Configuración del Proyecto y Servicios Base

### Configuración Inicial

```bash
# Crear el proyecto
ng new biblioteca-digital --routing --style=css

# Instalar dependencias necesarias
npm install @angular/material @angular/cdk @angular/platform-browser-dynamic
npm install @angular/forms @angular/common/http rxjs

# Para desarrollo: JSON Server como backend simulado
npm install -D json-server json-server-auth faker

# Instalar Angular Material
ng add @angular/material
```

### Variables de Entorno

```typescript
// environments/environment.ts
export const environment = {
  production: false,
  apiUrl: 'http://localhost:3000',
  jwtSecret: 'your-development-secret-key'
};

// environments/environment.prod.ts
export const environment = {
  production: true,
  apiUrl: 'https://api.biblioteca-digital.com',
  jwtSecret: 'your-production-secret-key'
};
```

### Servicio Base con Manejo de Errores

```typescript
// core/services/base-api.service.ts
import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpParams } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { environment } from '../../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class BaseApiService {

  constructor(protected http: HttpClient) {}

  protected get<T>(endpoint: string, params?: HttpParams): Observable<T> {
    return this.http.get<T>(`${environment.apiUrl}${endpoint}`, { params })
      .pipe(catchError(this.handleError));
  }

  protected post<T>(endpoint: string, data: any): Observable<T> {
    return this.http.post<T>(`${environment.apiUrl}${endpoint}`, data)
      .pipe(catchError(this.handleError));
  }

  protected put<T>(endpoint: string, data: any): Observable<T> {
    return this.http.put<T>(`${environment.apiUrl}${endpoint}`, data)
      .pipe(catchError(this.handleError));
  }

  protected delete<T>(endpoint: string): Observable<T> {
    return this.http.delete<T>(`${environment.apiUrl}${endpoint}`)
      .pipe(catchError(this.handleError));
  }

  private handleError = (error: HttpErrorResponse): Observable<never> => {
    let errorMessage = 'Ha ocurrido un error desconocido';

    if (error.error instanceof ErrorEvent) {
      // Error del lado del cliente
      errorMessage = `Error del cliente: ${error.error.message}`;
    } else {
      // Error del lado del servidor
      switch (error.status) {
        case 400:
          errorMessage = 'Datos inválidos. Por favor, verifica la información.';
          break;
        case 401:
          errorMessage = 'No autorizado. Tu sesión puede haber expirado.';
          break;
        case 403:
          errorMessage = 'No tienes permisos para realizar esta acción.';
          break;
        case 404:
          errorMessage = 'El recurso solicitado no existe.';
          break;
        case 409:
          errorMessage = 'Conflicto: El recurso ya existe o hay un conflicto de datos.';
          break;
        case 500:
          errorMessage = 'Error interno del servidor. Inténtalo más tarde.';
          break;
        default:
          errorMessage = `Error ${error.status}: ${error.message}`;
      }
    }

    console.error('API Error:', error);
    return throwError(() => new Error(errorMessage));
  };
}
```

---

## Paso 2: Servicio de Libros - El Corazón de Nuestra Biblioteca

```typescript
// features/books/services/book.service.ts
import { Injectable } from '@angular/core';
import { HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { BaseApiService } from '../../../core/services/base-api.service';
import {
  Book,
  CreateBookRequest,
  UpdateBookRequest,
  BookFilters,
  PaginatedResponse
} from '../models/book.model';

@Injectable({
  providedIn: 'root'
})
export class BookService extends BaseApiService {

  // Obtener todos los libros con filtros y paginación
  getBooks(filters?: BookFilters, page = 1, limit = 10): Observable<PaginatedResponse<Book>> {
    let params = new HttpParams()
      .set('page', page.toString())
      .set('limit', limit.toString());

    if (filters) {
      if (filters.title) params = params.set('title', filters.title);
      if (filters.author) params = params.set('author', filters.author);
      if (filters.genre) params = params.set('genre', filters.genre);
      if (filters.availableOnly) params = params.set('availableOnly', 'true');
    }

    return this.get<PaginatedResponse<Book>>('/books', params);
  }

  // Obtener un libro por ID
  getBookById(id: number): Observable<Book> {
    return this.get<Book>(`/books/${id}`);
  }

  // Crear un nuevo libro
  createBook(bookData: CreateBookRequest): Observable<Book> {
    return this.post<Book>('/books', bookData);
  }

  // Actualizar un libro existente
  updateBook(id: number, bookData: UpdateBookRequest): Observable<Book> {
    return this.put<Book>(`/books/${id}`, bookData);
  }

  // Eliminar un libro
  deleteBook(id: number): Observable<void> {
    return this.delete<void>(`/books/${id}`);
  }

  // Buscar libros por término
  searchBooks(query: string): Observable<Book[]> {
    const params = new HttpParams().set('q', query);
    return this.get<Book[]>('/books/search', params);
  }

  // Obtener géneros disponibles
  getGenres(): Observable<string[]> {
    return this.get<string[]>('/books/genres');
  }

  // Prestar un libro (disminuir copias disponibles)
  borrowBook(id: number): Observable<Book> {
    return this.put<Book>(`/books/${id}/borrow`, {});
  }

  // Devolver un libro (aumentar copias disponibles)
  returnBook(id: number): Observable<Book> {
    return this.put<Book>(`/books/${id}/return`, {});
  }

  // Obtener estadísticas de la biblioteca
  getLibraryStats(): Observable<{
    totalBooks: number;
    availableBooks: number;
    borrowedBooks: number;
    genres: { [key: string]: number };
  }> {
    return this.get('/books/stats');
  }
}
```

---

## Paso 3: Componentes de la Biblioteca - La Interfaz de Usuario

### Componente de Listado de Libros

```typescript
// features/books/components/book-list/book-list.component.ts
import { Component, OnInit, OnDestroy } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Subject, debounceTime, distinctUntilChanged, takeUntil } from 'rxjs';
import { Book, BookFilters, PaginatedResponse } from '../../models/book.model';
import { BookService } from '../../services/book.service';

@Component({
  selector: 'app-book-list',
  templateUrl: './book-list.component.html',
  styleUrls: ['./book-list.component.css']
})
export class BookListComponent implements OnInit, OnDestroy {
  books: Book[] = [];
  currentPage = 1;
  totalPages = 1;
  totalBooks = 0;
  isLoading = false;
  errorMessage = '';

  filtersForm: FormGroup;
  availableGenres: string[] = [];

  private destroy$ = new Subject<void>();

  constructor(
    private bookService: BookService,
    private fb: FormBuilder
  ) {
    this.filtersForm = this.fb.group({
      title: [''],
      author: [''],
      genre: [''],
      availableOnly: [false]
    });
  }

  ngOnInit() {
    this.loadGenres();
    this.loadBooks();
    this.setupFiltersSubscription();
  }

  ngOnDestroy() {
    this.destroy$.next();
    this.destroy$.complete();
  }

  loadBooks() {
    this.isLoading = true;
    this.errorMessage = '';

    const filters: BookFilters = this.filtersForm.value;

    this.bookService.getBooks(filters, this.currentPage)
      .pipe(takeUntil(this.destroy$))
      .subscribe({
        next: (response: PaginatedResponse<Book>) => {
          this.books = response.data;
          this.totalBooks = response.total;
          this.totalPages = response.totalPages;
          this.isLoading = false;
        },
        error: (error) => {
          this.errorMessage = error.message;
          this.isLoading = false;
        }
      });
  }

  loadGenres() {
    this.bookService.getGenres()
      .pipe(takeUntil(this.destroy$))
      .subscribe({
        next: (genres) => this.availableGenres = genres,
        error: (error) => console.error('Error loading genres:', error)
      });
  }

  setupFiltersSubscription() {
    this.filtersForm.valueChanges
      .pipe(
        debounceTime(300),
        distinctUntilChanged(),
        takeUntil(this.destroy$)
      )
      .subscribe(() => {
        this.currentPage = 1; // Reset to first page when filtering
        this.loadBooks();
      });
  }

  onPageChange(page: number) {
    this.currentPage = page;
    this.loadBooks();
  }

  onEditBook(book: Book) {
    // Navegar al formulario de edición
    // this.router.navigate(['/books/edit', book.id]);
  }

  onDeleteBook(book: Book) {
    if (confirm(`¿Estás seguro de que quieres eliminar "${book.title}"?`)) {
      this.bookService.deleteBook(book.id!)
        .pipe(takeUntil(this.destroy$))
        .subscribe({
          next: () => {
            this.loadBooks(); // Recargar la lista
          },
          error: (error) => {
            this.errorMessage = `Error al eliminar el libro: ${error.message}`;
          }
        });
    }
  }

  onBorrowBook(book: Book) {
    if (book.availableCopies > 0) {
      this.bookService.borrowBook(book.id!)
        .pipe(takeUntil(this.destroy$))
        .subscribe({
          next: (updatedBook) => {
            // Actualizar el libro en la lista
            const index = this.books.findIndex(b => b.id === book.id);
            if (index !== -1) {
              this.books[index] = updatedBook;
            }
          },
          error: (error) => {
            this.errorMessage = `Error al prestar el libro: ${error.message}`;
          }
        });
    }
  }

  clearFilters() {
    this.filtersForm.reset();
  }
}
```

```html
<!-- features/books/components/book-list/book-list.component.html -->
<div class="book-list-container">
  <div class="header">
    <h2>Catálogo de Libros</h2>
    <button
      mat-raised-button
      color="primary"
      routerLink="/books/create">
      <mat-icon>add</mat-icon>
      Agregar Libro
    </button>
  </div>

  <!-- Filtros -->
  <div class="filters-card">
    <form [formGroup]="filtersForm" class="filters-form">
      <mat-form-field appearance="outline">
        <mat-label>Título</mat-label>
        <input matInput formControlName="title" placeholder="Buscar por título">
        <mat-icon matSuffix>search</mat-icon>
      </mat-form-field>

      <mat-form-field appearance="outline">
        <mat-label>Autor</mat-label>
        <input matInput formControlName="author" placeholder="Buscar por autor">
      </mat-form-field>

      <mat-form-field appearance="outline">
        <mat-label>Género</mat-label>
        <mat-select formControlName="genre">
          <mat-option value="">Todos los géneros</mat-option>
          <mat-option *ngFor="let genre of availableGenres" [value]="genre">
            {{ genre }}
          </mat-option>
        </mat-select>
      </mat-form-field>

      <mat-checkbox formControlName="availableOnly">
        Solo libros disponibles
      </mat-checkbox>

      <button
        mat-stroked-button
        type="button"
        (click)="clearFilters()">
        <mat-icon>clear</mat-icon>
        Limpiar Filtros
      </button>
    </form>
  </div>

  <!-- Mensaje de error -->
  <div *ngIf="errorMessage" class="error-message">
    <mat-icon color="warn">error</mat-icon>
    {{ errorMessage }}
  </div>

  <!-- Loading -->
  <div *ngIf="isLoading" class="loading">
    <mat-spinner diameter="50"></mat-spinner>
    <p>Cargando libros...</p>
  </div>

  <!-- Lista de libros -->
  <div *ngIf="!isLoading && books.length > 0" class="books-grid">
    <mat-card *ngFor="let book of books" class="book-card">
      <mat-card-header>
        <mat-card-title>{{ book.title }}</mat-card-title>
        <mat-card-subtitle>{{ book.author }}</mat-card-subtitle>
      </mat-card-header>

      <img
        mat-card-image
        [src]="book.coverImage || '/assets/images/default-book.jpg'"
        [alt]="book.title"
        class="book-cover">

      <mat-card-content>
        <p class="book-description">{{ book.description | slice:0:150 }}...</p>
        <div class="book-details">
          <span class="genre-chip" mat-chip>{{ book.genre }}</span>
          <span class="year">{{ book.publishedYear }}</span>
        </div>
        <div class="availability">
          <mat-icon [color]="book.availableCopies > 0 ? 'primary' : 'warn'">
            {{ book.availableCopies > 0 ? 'check_circle' : 'cancel' }}
          </mat-icon>
          {{ book.availableCopies }} de {{ book.totalCopies }} disponibles
        </div>
      </mat-card-content>

      <mat-card-actions>
        <button mat-button routerLink="/books/{{ book.id }}">
          <mat-icon>visibility</mat-icon>
          Ver Detalles
        </button>

        <button mat-button color="primary" (click)="onEditBook(book)">
          <mat-icon>edit</mat-icon>
          Editar
        </button>

        <button
          mat-button
          color="warn"
          (click)="onDeleteBook(book)"
          [disabled]="book.availableCopies < book.totalCopies">
          <mat-icon>delete</mat-icon>
          Eliminar
        </button>

        <button
          mat-button
          color="accent"
          (click)="onBorrowBook(book)"
          [disabled]="book.availableCopies === 0">
          <mat-icon>shopping_cart</mat-icon>
          {{ book.availableCopies > 0 ? 'Prestar' : 'No disponible' }}
        </button>
      </mat-card-actions>
    </mat-card>
  </div>

  <!-- Sin resultados -->
  <div *ngIf="!isLoading && books.length === 0" class="no-results">
    <mat-icon class="no-results-icon">library_books</mat-icon>
    <h3>No se encontraron libros</h3>
    <p>Intenta ajustar los filtros de búsqueda</p>
  </div>

  <!-- Paginación -->
  <div *ngIf="!isLoading && totalPages > 1" class="pagination">
    <button
      mat-button
      [disabled]="currentPage === 1"
      (click)="onPageChange(currentPage - 1)">
      <mat-icon>chevron_left</mat-icon>
      Anterior
    </button>

    <span class="page-info">
      Página {{ currentPage }} de {{ totalPages }}
      ({{ totalBooks }} libros totales)
    </span>

    <button
      mat-button
      [disabled]="currentPage === totalPages"
      (click)="onPageChange(currentPage + 1)">
      Siguiente
      <mat-icon>chevron_right</mat-icon>
    </button>
  </div>
</div>
```

```css
/* features/books/components/book-list/book-list.component.css */
.book-list-container {
  padding: 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.filters-card {
  background: white;
  border-radius: 8px;
  padding: 24px;
  margin-bottom: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.filters-form {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
  align-items: end;
}

.books-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 24px;
  margin-bottom: 24px;
}

.book-card {
  height: fit-content;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.book-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

.book-cover {
  height: 200px;
  object-fit: cover;
  border-radius: 4px;
}

.book-description {
  margin: 16px 0;
  color: #666;
  line-height: 1.5;
}

.book-details {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.genre-chip {
  background: #e3f2fd;
  color: #1976d2;
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 12px;
  font-weight: 500;
}

.availability {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 500;
}

.mat-card-actions {
  padding-top: 16px;
  justify-content: space-between;
}

.no-results {
  text-align: center;
  padding: 48px;
  color: #666;
}

.no-results-icon {
  font-size: 64px;
  width: 64px;
  height: 64px;
  margin-bottom: 16px;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 16px;
  margin-top: 24px;
}

.page-info {
  color: #666;
  font-weight: 500;
}

.error-message {
  display: flex;
  align-items: center;
  gap: 8px;
  background: #ffebee;
  color: #c62828;
  padding: 12px 16px;
  border-radius: 4px;
  margin-bottom: 16px;
}

.loading {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 48px;
  gap: 16px;
}
```

### Componente de Formulario de Libro

```typescript
// features/books/components/book-form/book-form.component.ts
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { BookService } from '../../services/book.service';
import { Book, CreateBookRequest, UpdateBookRequest } from '../../models/book.model';

@Component({
  selector: 'app-book-form',
  templateUrl: './book-form.component.html',
  styleUrls: ['./book-form.component.css']
})
export class BookFormComponent implements OnInit {
  bookForm: FormGroup;
  isEditMode = false;
  isLoading = false;
  isSubmitting = false;
  errorMessage = '';
  successMessage = '';

  currentBook: Book | null = null;
  availableGenres: string[] = [
    'Ficción', 'No Ficción', 'Ciencia Ficción', 'Fantasía',
    'Misterio', 'Romance', 'Biografía', 'Historia',
    'Ciencia', 'Tecnología', 'Arte', 'Literatura Infantil'
  ];

  constructor(
    private fb: FormBuilder,
    private bookService: BookService,
    private route: ActivatedRoute,
    private router: Router
  ) {
    this.bookForm = this.createForm();
  }

  ngOnInit() {
    const id = this.route.snapshot.params['id'];
    if (id) {
      this.isEditMode = true;
      this.loadBook(+id);
    }
  }

  createForm(): FormGroup {
    return this.fb.group({
      title: ['', [Validators.required, Validators.minLength(2), Validators.maxLength(200)]],
      author: ['', [Validators.required, Validators.minLength(2), Validators.maxLength(100)]],
      isbn: ['', [Validators.required, Validators.pattern(/^(?:\d{10}|\d{13})$/)]],
      publishedYear: ['', [Validators.required, Validators.min(1000), Validators.max(new Date().getFullYear())]],
      genre: ['', Validators.required],
      description: ['', [Validators.required, Validators.minLength(10), Validators.maxLength(1000)]],
      totalCopies: [1, [Validators.required, Validators.min(1), Validators.max(1000)]],
      coverImage: ['']
    });
  }

  loadBook(id: number) {
    this.isLoading = true;
    this.bookService.getBookById(id).subscribe({
      next: (book) => {
        this.currentBook = book;
        this.bookForm.patchValue({
          title: book.title,
          author: book.author,
          isbn: book.isbn,
          publishedYear: book.publishedYear,
          genre: book.genre,
          description: book.description,
          totalCopies: book.totalCopies,
          coverImage: book.coverImage || ''
        });
        this.isLoading = false;
      },
      error: (error) => {
        this.errorMessage = `Error al cargar el libro: ${error.message}`;
        this.isLoading = false;
      }
    });
  }

  onSubmit() {
    if (this.bookForm.valid) {
      this.isSubmitting = true;
      this.errorMessage = '';
      this.successMessage = '';

      const formValue = this.bookForm.value;

      if (this.isEditMode && this.currentBook) {
        this.updateBook(this.currentBook.id!, formValue);
      } else {
        this.createBook(formValue);
      }
    } else {
      this.markFormGroupTouched();
    }
  }

  createBook(formValue: any) {
    const bookData: CreateBookRequest = {
      title: formValue.title,
      author: formValue.author,
      isbn: formValue.isbn,
      publishedYear: formValue.publishedYear,
      genre: formValue.genre,
      description: formValue.description,
      totalCopies: formValue.totalCopies
    };

    if (formValue.coverImage) {
      bookData.coverImage = formValue.coverImage;
    }

    this.bookService.createBook(bookData).subscribe({
      next: (book) => {
        this.successMessage = 'Libro creado exitosamente';
        this.isSubmitting = false;
        setTimeout(() => {
          this.router.navigate(['/books']);
        }, 1500);
      },
      error: (error) => {
        this.errorMessage = `Error al crear el libro: ${error.message}`;
        this.isSubmitting = false;
      }
    });
  }

  updateBook(id: number, formValue: any) {
    const bookData: UpdateBookRequest = {
      title: formValue.title,
      author: formValue.author,
      isbn: formValue.isbn,
      publishedYear: formValue.publishedYear,
      genre: formValue.genre,
      description: formValue.description,
      totalCopies: formValue.totalCopies
    };

    if (formValue.coverImage) {
      bookData.coverImage = formValue.coverImage;
    }

    this.bookService.updateBook(id, bookData).subscribe({
      next: (book) => {
        this.successMessage = 'Libro actualizado exitosamente';
        this.isSubmitting = false;
        setTimeout(() => {
          this.router.navigate(['/books']);
        }, 1500);
      },
      error: (error) => {
        this.errorMessage = `Error al actualizar el libro: ${error.message}`;
        this.isSubmitting = false;
      }
    });
  }

  onCancel() {
    this.router.navigate(['/books']);
  }

  onImageUpload(event: any) {
    const file = event.target.files[0];
    if (file) {
      // Aquí implementarías la lógica de subida de imagen
      // Por simplicidad, solo guardamos la URL
      const reader = new FileReader();
      reader.onload = (e) => {
        this.bookForm.patchValue({
          coverImage: e.target?.result as string
        });
      };
      reader.readAsDataURL(file);
    }
  }

  private markFormGroupTouched() {
    Object.keys(this.bookForm.controls).forEach(key => {
      this.bookForm.get(key)?.markAsTouched();
    });
  }

  // Getters para facilitar el acceso en el template
  get title() { return this.bookForm.get('title'); }
  get author() { return this.bookForm.get('author'); }
  get isbn() { return this.bookForm.get('isbn'); }
  get publishedYear() { return this.bookForm.get('publishedYear'); }
  get genre() { return this.bookForm.get('genre'); }
  get description() { return this.bookForm.get('description'); }
  get totalCopies() { return this.bookForm.get('totalCopies'); }
}
```

```html
<!-- features/books/components/book-form/book-form.component.html -->
<div class="book-form-container">
  <div class="form-header">
    <h2>{{ isEditMode ? 'Editar Libro' : 'Agregar Nuevo Libro' }}</h2>
    <button mat-button (click)="onCancel()">
      <mat-icon>arrow_back</mat-icon>
      Volver al Catálogo
    </button>
  </div>

  <!-- Loading -->
  <div *ngIf="isLoading" class="loading">
    <mat-spinner diameter="50"></mat-spinner>
    <p>Cargando libro...</p>
  </div>

  <!-- Mensajes -->
  <div *ngIf="errorMessage" class="message error-message">
    <mat-icon>error</mat-icon>
    {{ errorMessage }}
  </div>

  <div *ngIf="successMessage" class="message success-message">
    <mat-icon>check_circle</mat-icon>
    {{ successMessage }}
  </div>

  <!-- Formulario -->
  <form *ngIf="!isLoading" [formGroup]="bookForm" (ngSubmit)="onSubmit()" class="book-form">
    <div class="form-grid">
      <!-- Columna izquierda -->
      <div class="form-column">
        <mat-form-field appearance="outline" class="full-width">
          <mat-label>Título del Libro</mat-label>
          <input
            matInput
            formControlName="title"
            placeholder="Ej: Cien años de soledad"
            [class.error]="title?.invalid && title?.touched">
          <mat-error *ngIf="title?.errors?.['required']">El título es requerido</mat-error>
          <mat-error *ngIf="title?.errors?.['minlength']">Mínimo 2 caracteres</mat-error>
          <mat-error *ngIf="title?.errors?.['maxlength']">Máximo 200 caracteres</mat-error>
        </mat-form-field>

        <mat-form-field appearance="outline" class="full-width">
          <mat-label>Autor</mat-label>
          <input
            matInput
            formControlName="author"
            placeholder="Ej: Gabriel García Márquez"
            [class.error]="author?.invalid && author?.touched">
          <mat-error *ngIf="author?.errors?.['required']">El autor es requerido</mat-error>
          <mat-error *ngIf="author?.errors?.['minlength']">Mínimo 2 caracteres</mat-error>
          <mat-error *ngIf="author?.errors?.['maxlength']">Máximo 100 caracteres</mat-error>
        </mat-form-field>

        <mat-form-field appearance="outline" class="full-width">
          <mat-label>ISBN</mat-label>
          <input
            matInput
            formControlName="isbn"
            placeholder="Ej: 9788437604947"
            [class.error]="isbn?.invalid && isbn?.touched">
          <mat-hint>10 o 13 dígitos sin guiones</mat-hint>
          <mat-error *ngIf="isbn?.errors?.['required']">El ISBN es requerido</mat-error>
          <mat-error *ngIf="isbn?.errors?.['pattern']">ISBN inválido (solo números, 10 o 13 dígitos)</mat-error>
        </mat-form-field>

        <mat-form-field appearance="outline">
          <mat-label>Año de Publicación</mat-label>
          <input
            matInput
            type="number"
            formControlName="publishedYear"
            placeholder="Ej: 1967"
            [class.error]="publishedYear?.invalid && publishedYear?.touched">
          <mat-error *ngIf="publishedYear?.errors?.['required']">El año es requerido</mat-error>
          <mat-error *ngIf="publishedYear?.errors?.['min']">Año inválido</mat-error>
          <mat-error *ngIf="publishedYear?.errors?.['max']">El año no puede ser futuro</mat-error>
        </mat-form-field>

        <mat-form-field appearance="outline">
          <mat-label>Género</mat-label>
          <mat-select formControlName="genre" [class.error]="genre?.invalid && genre?.touched">
            <mat-option *ngFor="let genreOption of availableGenres" [value]="genreOption">
              {{ genreOption }}
            </mat-option>
          </mat-select>
          <mat-error *ngIf="genre?.errors?.['required']">El género es requerido</mat-error>
        </mat-form-field>
      </div>

      <!-- Columna derecha -->
      <div class="form-column">
        <mat-form-field appearance="outline">
          <mat-label>Número de Copias</mat-label>
          <input
            matInput
            type="number"
            formControlName="totalCopies"
            min="1"
            max="1000"
            [class.error]="totalCopies?.invalid && totalCopies?.touched">
          <mat-error *ngIf="totalCopies?.errors?.['required']">El número de copias es requerido</mat-error>
          <mat-error *ngIf="totalCopies?.errors?.['min']">Mínimo 1 copia</mat-error>
          <mat-error *ngIf="totalCopies?.errors?.['max']">Máximo 1000 copias</mat-error>
        </mat-form-field>

        <mat-form-field appearance="outline" class="full-width">
          <mat-label>Descripción</mat-label>
          <textarea
            matInput
            formControlName="description"
            rows="4"
            placeholder="Describe brevemente el contenido del libro..."
            [class.error]="description?.invalid && description?.touched">
          </textarea>
          <mat-hint>{{ description?.value?.length || 0 }}/1000 caracteres</mat-hint>
          <mat-error *ngIf="description?.errors?.['required']">La descripción es requerida</mat-error>
          <mat-error *ngIf="description?.errors?.['minlength']">Mínimo 10 caracteres</mat-error>
          <mat-error *ngIf="description?.errors?.['maxlength']">Máximo 1000 caracteres</mat-error>
        </mat-form-field>

        <div class="image-upload">
          <mat-label>Imagen de Portada (Opcional)</mat-label>
          <input
            type="file"
            accept="image/*"
            (change)="onImageUpload($event)"
            style="display: none"
            #fileInput>
          <button
            mat-stroked-button
            type="button"
            (click)="fileInput.click()">
            <mat-icon>cloud_upload</mat-icon>
            Subir Imagen
          </button>
          <div *ngIf="bookForm.get('coverImage')?.value" class="image-preview">
            <img [src]="bookForm.get('coverImage')?.value" alt="Vista previa">
          </div>
        </div>
      </div>
    </div>

    <!-- Botones de acción -->
    <div class="form-actions">
      <button
        mat-button
        type="button"
        (click)="onCancel()">
        Cancelar
      </button>

      <button
        mat-raised-button
        color="primary"
        type="submit"
        [disabled]="isSubmitting">
        <mat-icon *ngIf="isSubmitting">hourglass_empty</mat-icon>
        <mat-icon *ngIf="!isSubmitting">{{ isEditMode ? 'save' : 'add' }}</mat-icon>
        {{ isSubmitting ? 'Guardando...' : (isEditMode ? 'Actualizar Libro' : 'Crear Libro') }}
      </button>
    </div>
  </form>
</div>
```

```css
/* features/books/components/book-form/book-form.component.css */
.book-form-container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 24px;
}

.form-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.book-form {
  background: white;
  border-radius: 8px;
  padding: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
  margin-bottom: 24px;
}

.form-column {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.full-width {
  width: 100%;
}

.image-upload {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.image-preview {
  margin-top: 8px;
}

.image-preview img {
  max-width: 200px;
  max-height: 200px;
  object-fit: cover;
  border-radius: 4px;
  border: 1px solid #ddd;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 16px;
  padding-top: 24px;
  border-top: 1px solid #e0e0e0;
}

.message {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 16px;
  border-radius: 4px;
  margin-bottom: 16px;
}

.error-message {
  background: #ffebee;
  color: #c62828;
  border: 1px solid #ffcdd2;
}

.success-message {
  background: #e8f5e8;
  color: #2e7d32;
  border: 1px solid #c8e6c9;
}

.loading {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 48px;
  gap: 16px;
}

/* Responsive */
@media (max-width: 768px) {
  .form-grid {
    grid-template-columns: 1fr;
  }

  .form-header {
    flex-direction: column;
    gap: 16px;
    align-items: flex-start;
  }
}
```

### Componente de Detalles del Libro

```typescript
// features/books/components/book-detail/book-detail.component.ts
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { BookService } from '../../services/book.service';
import { Book } from '../../models/book.model';

@Component({
  selector: 'app-book-detail',
  templateUrl: './book-detail.component.html',
  styleUrls: ['./book-detail.component.css']
})
export class BookDetailComponent implements OnInit {
  book: Book | null = null;
  isLoading = false;
  errorMessage = '';

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private bookService: BookService
  ) {}

  ngOnInit() {
    const id = +this.route.snapshot.params['id'];
    this.loadBook(id);
  }

  loadBook(id: number) {
    this.isLoading = true;
    this.bookService.getBookById(id).subscribe({
      next: (book) => {
        this.book = book;
        this.isLoading = false;
      },
      error: (error) => {
        this.errorMessage = `Error al cargar el libro: ${error.message}`;
        this.isLoading = false;
      }
    });
  }

  onEdit() {
    if (this.book) {
      this.router.navigate(['/books/edit', this.book.id]);
    }
  }

  onDelete() {
    if (this.book && confirm(`¿Estás seguro de que quieres eliminar "${this.book.title}"?`)) {
      this.bookService.deleteBook(this.book.id!).subscribe({
        next: () => {
          this.router.navigate(['/books']);
        },
        error: (error) => {
          this.errorMessage = `Error al eliminar el libro: ${error.message}`;
        }
      });
    }
  }

  onBorrow() {
    if (this.book && this.book.availableCopies > 0) {
      this.bookService.borrowBook(this.book.id!).subscribe({
        next: (updatedBook) => {
          this.book = updatedBook;
        },
        error: (error) => {
          this.errorMessage = `Error al prestar el libro: ${error.message}`;
        }
      });
    }
  }

  onReturn() {
    if (this.book) {
      this.bookService.returnBook(this.book.id!).subscribe({
        next: (updatedBook) => {
          this.book = updatedBook;
        },
        error: (error) => {
          this.errorMessage = `Error al devolver el libro: ${error.message}`;
        }
      });
    }
  }
}
```

```html
<!-- features/books/components/book-detail/book-detail.component.html -->
<div class="book-detail-container" *ngIf="book">
  <div class="book-header">
    <button mat-button (click)="router.navigate(['/books'])">
      <mat-icon>arrow_back</mat-icon>
      Volver al Catálogo
    </button>
  </div>

  <div *ngIf="errorMessage" class="error-message">
    <mat-icon color="warn">error</mat-icon>
    {{ errorMessage }}
  </div>

  <div *ngIf="isLoading" class="loading">
    <mat-spinner diameter="50"></mat-spinner>
    <p>Cargando detalles del libro...</p>
  </div>

  <div *ngIf="!isLoading && book" class="book-detail">
    <div class="book-cover-section">
      <img
        [src]="book.coverImage || '/assets/images/default-book.jpg'"
        [alt]="book.title"
        class="book-cover-large">
    </div>

    <div class="book-info-section">
      <h1 class="book-title">{{ book.title }}</h1>
      <h2 class="book-author">por {{ book.author }}</h2>

      <div class="book-meta">
        <span class="meta-item">
          <mat-icon>calendar_today</mat-icon>
          {{ book.publishedYear }}
        </span>
        <span class="meta-item genre-chip">
          {{ book.genre }}
        </span>
        <span class="meta-item">
          <mat-icon>tag</mat-icon>
          ISBN: {{ book.isbn }}
        </span>
      </div>

      <div class="availability-section">
        <div class="availability-status" [class.available]="book.availableCopies > 0" [class.unavailable]="book.availableCopies === 0">
          <mat-icon>{{ book.availableCopies > 0 ? 'check_circle' : 'cancel' }}</mat-icon>
          <span>{{ book.availableCopies > 0 ? 'Disponible' : 'No disponible' }}</span>
        </div>
        <p class="availability-text">
          {{ book.availableCopies }} de {{ book.totalCopies }} copias disponibles
        </p>
      </div>

      <div class="book-description">
        <h3>Descripción</h3>
        <p>{{ book.description }}</p>
      </div>

      <div class="book-actions">
        <button mat-raised-button color="primary" (click)="onEdit()">
          <mat-icon>edit</mat-icon>
          Editar Libro
        </button>

        <button mat-stroked-button color="warn" (click)="onDelete()">
          <mat-icon>delete</mat-icon>
          Eliminar Libro
        </button>

        <button
          mat-raised-button
          color="accent"
          (click)="onBorrow()"
          [disabled]="book.availableCopies === 0">
          <mat-icon>shopping_cart</mat-icon>
          {{ book.availableCopies > 0 ? 'Prestar Libro' : 'No Disponible' }}
        </button>

        <button
          mat-stroked-button
          (click)="onReturn()"
          [disabled]="book.availableCopies === book.totalCopies">
          <mat-icon>undo</mat-icon>
          Devolver Libro
        </button>
      </div>
    </div>
  </div>
</div>
```

---

## Paso 4: Configuración de Rutas Protegidas

```typescript
// features/books/books-routing.module.ts
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from '../../core/guards/auth.guard';
import { BookListComponent } from './components/book-list/book-list.component';
import { BookFormComponent } from './components/book-form/book-form.component';
import { BookDetailComponent } from './components/book-detail/book-detail.component';

const routes: Routes = [
  {
    path: '',
    component: BookListComponent,
    canActivate: [AuthGuard]
  },
  {
    path: 'create',
    component: BookFormComponent,
    canActivate: [AuthGuard]
  },
  {
    path: 'edit/:id',
    component: BookFormComponent,
    canActivate: [AuthGuard]
  },
  {
    path: ':id',
    component: BookDetailComponent,
    canActivate: [AuthGuard]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BooksRoutingModule { }
```

---

## Paso 5: Backend Simulado con JSON Server

Para probar nuestra aplicación, vamos a crear un backend simulado usando JSON Server.

```json
// db.json
{
  "books": [
    {
      "id": 1,
      "title": "Cien años de soledad",
      "author": "Gabriel García Márquez",
      "isbn": "9788437604947",
      "publishedYear": 1967,
      "genre": "Ficción",
      "description": "Una de las obras más importantes del realismo mágico, narra la historia de la familia Buendía a través de siete generaciones en el pueblo ficticio de Macondo.",
      "totalCopies": 5,
      "availableCopies": 3,
      "createdAt": "2024-01-01T00:00:00.000Z",
      "updatedAt": "2024-01-01T00:00:00.000Z"
    },
    {
      "id": 2,
      "title": "1984",
      "author": "George Orwell",
      "isbn": "9780451524935",
      "publishedYear": 1949,
      "genre": "Ciencia Ficción",
      "description": "Una novela distópica que explora los peligros del totalitarismo, la vigilancia masiva y la manipulación de la verdad.",
      "totalCopies": 3,
      "availableCopies": 1,
      "createdAt": "2024-01-02T00:00:00.000Z",
      "updatedAt": "2024-01-02T00:00:00.000Z"
    }
  ],
  "users": [
    {
      "id": 1,
      "email": "admin@biblioteca.com",
      "password": "admin123",
      "name": "Administrador",
      "role": "admin"
    }
  ]
}
```

```typescript
// server.js - Servidor JSON Server con autenticación
const jsonServer = require('json-server');
const server = jsonServer.create();
const router = jsonServer.router('db.json');
const middlewares = jsonServer.defaults();
const jwt = require('jsonwebtoken');
const bcrypt = require('bcryptjs');

server.use(middlewares);
server.use(jsonServer.bodyParser);

// Middleware de autenticación
server.use('/api', (req, res, next) => {
  if (req.path === '/login' || req.path === '/register') {
    return next();
  }

  const authHeader = req.headers.authorization;
  if (!authHeader) {
    return res.status(401).json({ message: 'Token requerido' });
  }

  const token = authHeader.split(' ')[1];
  try {
    jwt.verify(token, 'your-secret-key');
    next();
  } catch (error) {
    res.status(403).json({ message: 'Token inválido' });
  }
});

// Endpoint de login
server.post('/api/login', (req, res) => {
  const { email, password } = req.body;
  const users = router.db.get('users').value();
  const user = users.find(u => u.email === email);

  if (!user || !bcrypt.compareSync(password, user.password)) {
    return res.status(401).json({ message: 'Credenciales incorrectas' });
  }

  const token = jwt.sign(
    { id: user.id, email: user.email, role: user.role },
    'your-secret-key',
    { expiresIn: '24h' }
  );

  res.json({ token, user: { id: user.id, email: user.email, name: user.name, role: user.role } });
});

// Endpoint de registro
server.post('/api/register', (req, res) => {
  const { name, email, password } = req.body;
  const users = router.db.get('users').value();

  if (users.find(u => u.email === email)) {
    return res.status(409).json({ message: 'El email ya está registrado' });
  }

  const hashedPassword = bcrypt.hashSync(password, 10);
  const newUser = {
    id: Date.now(),
    name,
    email,
    password: hashedPassword,
    role: 'user'
  };

  router.db.get('users').push(newUser).write();

  const token = jwt.sign(
    { id: newUser.id, email: newUser.email, role: newUser.role },
    'your-secret-key',
    { expiresIn: '24h' }
  );

  res.json({
    token,
    user: { id: newUser.id, email: newUser.email, name: newUser.name, role: newUser.role }
  });
});

// Endpoint para estadísticas
server.get('/api/books/stats', (req, res) => {
  const books = router.db.get('books').value();
  const stats = {
    totalBooks: books.length,
    availableBooks: books.reduce((sum, book) => sum + book.availableCopies, 0),
    borrowedBooks: books.reduce((sum, book) => sum + (book.totalCopies - book.availableCopies), 0),
    genres: books.reduce((acc, book) => {
      acc[book.genre] = (acc[book.genre] || 0) + 1;
      return acc;
    }, {})
  };
  res.json(stats);
});

// Endpoint para géneros
server.get('/api/books/genres', (req, res) => {
  const books = router.db.get('books').value();
  const genres = [...new Set(books.map(book => book.genre))];
  res.json(genres);
});

// Endpoint para búsqueda
server.get('/api/books/search', (req, res) => {
  const { q } = req.query;
  const books = router.db.get('books').value();
  const results = books.filter(book =>
    book.title.toLowerCase().includes(q.toLowerCase()) ||
    book.author.toLowerCase().includes(q.toLowerCase()) ||
    book.description.toLowerCase().includes(q.toLowerCase())
  );
  res.json(results);
});

// Endpoints para prestar/devolver libros
server.put('/api/books/:id/borrow', (req, res) => {
  const { id } = req.params;
  const book = router.db.get('books').find({ id: parseInt(id) }).value();

  if (!book) {
    return res.status(404).json({ message: 'Libro no encontrado' });
  }

  if (book.availableCopies <= 0) {
    return res.status(400).json({ message: 'No hay copias disponibles' });
  }

  router.db.get('books')
    .find({ id: parseInt(id) })
    .assign({ availableCopies: book.availableCopies - 1 })
    .write();

  res.json(router.db.get('books').find({ id: parseInt(id) }).value());
});

server.put('/api/books/:id/return', (req, res) => {
  const { id } = req.params;
  const book = router.db.get('books').find({ id: parseInt(id) }).value();

  if (!book) {
    return res.status(404).json({ message: 'Libro no encontrado' });
  }

  if (book.availableCopies >= book.totalCopies) {
    return res.status(400).json({ message: 'Todas las copias ya están disponibles' });
  }

  router.db.get('books')
    .find({ id: parseInt(id) })
    .assign({ availableCopies: book.availableCopies + 1 })
    .write();

  res.json(router.db.get('books').find({ id: parseInt(id) }).value());
});

server.use('/api', router);
server.listen(3000, () => {
  console.log('JSON Server is running on http://localhost:3000');
});
```

---

## Mejores Prácticas Implementadas

### 1. **Separación de Responsabilidades**
- **Servicios**: Manejan la lógica de negocio y comunicación HTTP
- **Componentes**: Gestionan la UI y estado local
- **Modelos**: Definen la estructura de datos
- **Guards**: Controlan el acceso a rutas

### 2. **Manejo de Estado Reactivo**
```typescript
// Usamos BehaviorSubject para estado reactivo
private booksSubject = new BehaviorSubject<Book[]>([]);
public books$ = this.booksSubject.asObservable();
```

### 3. **Validación Robusta**
```typescript
// Validaciones tanto en frontend como backend
this.bookForm = this.fb.group({
  title: ['', [Validators.required, Validators.minLength(2)]],
  isbn: ['', [Validators.required, Validators.pattern(/^(?:\d{10}|\d{13})$/)]]
});
```

### 4. **Manejo de Errores Global**
```typescript
// Interceptores para manejo centralizado de errores
@Injectable()
export class ErrorInterceptor implements HttpInterceptor {
  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return next.handle(request).pipe(
      catchError((error: HttpErrorResponse) => {
        // Lógica centralizada de manejo de errores
      })
    );
  }
}
```

### 5. **UX Mejorada**
- Loading states durante operaciones asíncronas
- Mensajes de éxito y error informativos
- Navegación intuitiva
- Responsive design

---

## Errores Comunes y Soluciones

### ❌ **Error 1: Mutar el Array Original**
```typescript
// MAL
this.books.push(newBook); // Muta el array original

// BIEN
this.books = [...this.books, newBook]; // Crea nuevo array
```

### ❌ **Error 2: No Manejar Estados de Loading**
```typescript
// MAL
deleteBook(id: number) {
  this.bookService.deleteBook(id).subscribe(() => {
    this.loadBooks(); // Usuario no sabe qué está pasando
  });
}

// BIEN
deleteBook(id: number) {
  this.isDeleting = true;
  this.bookService.deleteBook(id).subscribe({
    next: () => {
      this.loadBooks();
      this.showSuccess('Libro eliminado');
    },
    error: (error) => this.showError(error.message),
    complete: () => this.isDeleting = false
  });
}
```

### ❌ **Error 3: Validación Solo en Frontend**
```typescript
// MAL: Solo validación en frontend
if (this.bookForm.valid) {
  this.bookService.createBook(this.bookForm.value);
}

// BIEN: Validación en backend + manejo de errores
this.bookService.createBook(this.bookForm.value).subscribe({
  error: (error) => {
    if (error.status === 409) {
      this.bookForm.get('isbn')?.setErrors({ duplicate: true });
    }
  }
});
```

---

## Proyecto Final: Ejecutando la Biblioteca Digital

### Paso 1: Instalar Dependencias Adicionales
```bash
npm install bcryptjs jsonwebtoken cors
```

### Paso 2: Configurar el Servidor
```bash
# Crear archivo server.js con el código anterior
node server.js
```

### Paso 3: Ejecutar la Aplicación
```bash
# Terminal 1: Backend
npm run server

# Terminal 2: Frontend
ng serve
```

### Paso 4: Probar Todas las Funcionalidades
1. **Regístrate** como nuevo usuario
2. **Inicia sesión** con tus credenciales
3. **Agrega** varios libros al catálogo
4. **Busca y filtra** libros por diferentes criterios
5. **Edita** la información de un libro
6. **Presta y devuelve** libros
7. **Elimina** un libro del catálogo
8. **Cierra sesión** y verifica que las rutas estén protegidas

---

## Reflexiones Finales

¡Felicitaciones! Has construido una **Biblioteca Digital completa** que demuestra el poder del patrón CRUD en Angular. Esta aplicación incluye:

- 🔐 **Autenticación completa** con guards e interceptores
- 📚 **Gestión completa de libros** (CRUD operations)
- 🔍 **Búsqueda y filtrado avanzado**
- 📊 **Estadísticas y reportes**
- 🎨 **Interfaz moderna** con Angular Material
- 🛡️ **Manejo robusto de errores**
- 📱 **Diseño responsive**

Este proyecto es solo el comienzo. Puedes extenderlo con:

- **Sistema de préstamos** con fechas de devolución
- **Notificaciones** por email cuando un libro esté disponible
- **Reviews y ratings** de los usuarios
- **Categorización avanzada** con tags y subcategorías
- **API RESTful completa** con documentación
- **Tests unitarios** y de integración

El patrón CRUD que has dominado aquí se aplica a cualquier dominio: e-commerce, blogs, sistemas de gestión, redes sociales, etc. ¡Las posibilidades son infinitas!

En el próximo capítulo, exploraremos las mejores prácticas profesionales para llevar tus aplicaciones Angular al siguiente nivel. ¿Estás listo para convertirte en un desarrollador Angular profesional?

---

## 📚 Recursos Adicionales

- [Angular Material Documentation](https://material.angular.io/)
- [JSON Server Documentation](https://github.com/typicode/json-server)
- [RxJS Documentation](https://rxjs.dev/)
- [Angular Reactive Forms Guide](https://angular.io/guide/reactive-forms)

## 🎯 Checklist del Proyecto CRUD

- [ ] Arquitectura modular con feature modules
- [ ] Servicios HTTP con manejo de errores
- [ ] Formularios reactivos con validación
- [ ] Guards para protección de rutas
- [ ] Interceptores para autenticación automática
- [ ] Componentes reutilizables
- [ ] Estado reactivo con RxJS
- [ ] Backend simulado con JSON Server
- [ ] Interfaz responsive con Angular Material
- [ ] Manejo completo del ciclo de vida de datos
- [ ] Búsqueda, filtrado y paginación
- [ ] Tests básicos de funcionalidad

¡Has completado un proyecto CRUD completo! Este es un hito importante en tu viaje como desarrollador Angular. 🎉
