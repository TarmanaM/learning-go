# Go Visibility Rules (Exported vs Unexported)

In Go, there are no `public`, `private`, or `protected` keywords like in other object-oriented languages. 
Instead, Go uses **capitalization** to determine the visibility of identifiers (variables, struct fields, functions, methods, types).

### 1. Exported (Public) Identifiers
If an identifier starts with an **uppercase letter** (e.g., `Nama`, `ID`, `PrintData`), it is **exported**. 
This means it can be accessed from outside the package where it was defined.

Example:
```go
package clean

type Biodata struct {
	ID        int8   // Exported: Can be accessed from the `main` package
	Nama      string // Exported
}
```

### 2. Unexported (Private) Identifiers
If an identifier starts with a **lowercase letter** (e.g., `alasan`, `nama`), it is **unexported**.
This means it is private and can **only be accessed from within the same package** where it is declared.

If you try to initialize or access an unexported field from another package, the compiler will throw an error: `cannot refer to unexported field`.

### Example Fix
**Before (Error):**
```go
// In clean package
type Biodata struct {
    nama string // unexported
}

// In main package
data := clean.Biodata{
    nama: "Thomas", // Error: cannot refer to unexported field nama
}
```

**After (Fixed):**
```go
// In clean package
type Biodata struct {
    Nama string // Exported (uppercase 'N')
}

// In main package
data := clean.Biodata{
    Nama: "Thomas", // Works perfectly!
}
```
