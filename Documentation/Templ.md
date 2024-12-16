# Templ
## An HTML templating language for Go that has great developer tooling.

![templ](https://raw.githubusercontent.com/a-h/templ/main/ide-demo.gif)

**[Source (GitHub)](https://github.com/a-h/templ)**

**Templ** is an open-source, code-generation-based, type-safe HTML templating framework for Go. Its main goal is to eliminate the risk of runtime template errors by making use of Go’s compile-time checks and to provide a convenient DSL (Domain Specific Language) for writing dynamic user interfaces. [1](https://github.com/a-h/templ/blob/main/README.md) [2](https://github.com/a-h/templ/wiki)

---

## **History**

Templ was created by [Andrew Hood (GitHub: @a-h)](https://github.com/a-h) around 2019–2020 [1](https://github.com/a-h/templ/blob/main/README.md), [3](https://github.com/a-h/templ/discussions). He sought a templating solution that was both idiomatic to Go and type-safe, circumventing the runtime-only checks commonly found in `text/template` or `html/template`. 

Templ emerged as a side project and evolved through community contributions. Early commits date back to November 2019 [4](https://github.com/a-h/templ/commits?after=4d281b7c869e79d7f438a03c5b792851ffb90966+34). It’s licensed under the MIT license [1](https://github.com/a-h/templ/blob/main/README.md).

---

## **Key Characteristics**

1. **Type-Safe Templating**  
   - **Compile-Time Validation**: Templates are turned into Go code, catching type errors at build time [1](https://github.com/a-h/templ/blob/main/README.md), [2](https://github.com/a-h/templ/wiki).  
   - **No Reflection**: Templ does not rely on runtime reflection, improving reliability and performance [1](https://github.com/a-h/templ/blob/main/README.md).

2. **Code Generation**  
   - **Templ’s DSL**: Templ uses `.templ` files that resemble HTML with inline Go expressions. The CLI translates them into `.go` files [2](https://github.com/a-h/templ/wiki), [5](https://github.com/a-h/templ/tree/main/examples).  
   - **Minimal Overhead**: The generated code is idiomatic Go with no large runtime dependencies [1](https://github.com/a-h/templ/blob/main/README.md).

3. **Modular Architecture**  
   - **Partial Templates**: Allows for smaller reusable components.  
   - **Layouts**: Shared layouts or “master pages” can wrap child templates [1](https://github.com/a-h/templ/blob/main/README.md), [2](https://github.com/a-h/templ/wiki).

4. **Integration**  
   - **Standard Library**: Generated functions integrate easily with `net/http` [1](https://github.com/a-h/templ/blob/main/README.md), [2].  
   - **Go Web Frameworks**: Works with Gin, Echo, or any router because you just call the compiled function [2](https://github.com/a-h/templ/wiki).

5. **Open Source & Community-Driven**  
   - **MIT License**: Free to use in personal or commercial projects [1](https://github.com/a-h/templ/blob/main/README.md).  
   - **Community Contributions**: Pull requests, issue discussions, and examples on GitHub [3](https://github.com/a-h/templ/discussions).

---

## **How Templ Works**

1. **Write Templ DSL**:  
   A `.templ` file uses HTML-like syntax with embedded Go expressions:
   ```go
   Template MyTemplate(person Person) {
     div {
       h1("Hello, " + person.Name)
       if person.Age > 18 {
         p("You are an adult.")
       } else {
         p("You are a minor.")
       }
     }
   }
   ```

2. **Run Templ CLI**:  
   ```bash
   templ generate
   ```
   This generates `.go` files with strongly typed functions [2](https://github.com/a-h/templ/wiki), [5](https://github.com/a-h/templ/tree/main/examples).

3. **Import & Render**:  
   ```go
   func handler(w http.ResponseWriter, r *http.Request) {
       myTemplate := MyTemplate(Person{Name: "Alice", Age: 25})
       w.Write([]byte(myTemplate))
   }
   ```
   The rendered output is pure HTML [5](https://github.com/a-h/templ/tree/main/examples).

---

## **Pros**

1. **Compile-Time Safety**  
   - Prevents runtime template surprises via immediate feedback during compilation [1](https://github.com/a-h/templ/blob/main/README.md).

2. **Performance**  
   - No runtime parsing overhead; the output is directly compiled Go [1](https://github.com/a-h/templ/blob/main/README.md), [5](https://github.com/a-h/templ/tree/main/examples).

3. **Developer Experience**  
   - DSL is HTML-like with embedded Go logic, reminiscent of JSX or Razor pages [1](https://github.com/a-h/templ/blob/main/README.md), [2](https://github.com/a-h/templ/wiki).

4. **Flexible Integration**  
   - Easily integrates with `net/http` or any Go router or framework [2](https://github.com/a-h/templ/wiki), [3](https://github.com/a-h/templ/discussions).

5. **Lightweight**  
   - Everything compiles into your Go binary with no additional templating engines or runtime interpreters [1](https://github.com/a-h/templ/blob/main/README.md).

---

## **Cons**

1. **Learning Curve**  
   - Requires adopting the DSL and an extra build step via `templ generate` [1](https://github.com/a-h/templ/blob/main/README.md), [2](https://github.com/a-h/templ/wiki).

2. **Limited Ecosystem**  
   - Official `html/template` has broader adoption and community resources; Templ is still growing [1](https://github.com/a-h/templ/blob/main/README.md), [3](https://github.com/a-h/templ/discussions).

3. **Increased Build Complexity**  
   - A code generation phase means CI/CD pipelines must run `templ generate` before compiling [2](https://github.com/a-h/templ/wiki).

---

## **Use Cases**

- **Server-Side Web Rendering**: A type-safe alternative to `html/template` for building server-rendered UIs [1](https://github.com/a-h/templ/blob/main/README.md), [5].  
- **Email Template Generation**: Avoid runtime issues when generating HTML email content [1](https://github.com/a-h/templ/blob/main/README.md), [2](https://github.com/a-h/templ/wiki).  
- **Microservices / Small UIs**: Ideal for mini web interfaces in microservices [2](https://github.com/a-h/templ/wiki).  
- **Reusable UI Components**: Build partials and layouts that can be composed into larger pages [5](https://github.com/a-h/templ/tree/main/examples).

-----

## **Getting Started**

### **Installation**
```bash
go install github.com/a-h/templ/cmd/templ@latest
```
Installs the Templ CLI tool into `$GOPATH/bin` [2](https://github.com/a-h/templ/wiki).

### **Project Structure**
```
myproject/
├── main.go
└── templates/
    └── index.templ
```
A common folder layout for Templ-based projects [2](https://github.com/a-h/templ/wiki), [5](https://github.com/a-h/templ/tree/main/examples).

### **Write Your First Templ** (`index.templ`)
```go
Template Index(name string) {
  html {
    head {
      title("Welcome!")
    }
    body {
      h1("Hello " + name)
    }
  }
}
```
A basic example that uses Templ’s DSL [2](https://github.com/a-h/templ/wiki), [5](https://github.com/a-h/templ/tree/main/examples).

### **Generate & Build**
```bash
templ generate
go build .
```
Templ CLI creates `index_templ.go` that you can compile alongside your Go code [2](https://github.com/a-h/templ/wiki).

### **Use the Generated Code in `main.go`**
```go
package main

import (
    "log"
    "net/http"

    "myproject/templates" // The generated file
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        output := templates.Index("Alice")
        w.Write([]byte(output))
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```
A standard approach using Templ in a Go application [2](https://github.com/a-h/templ/wiki), [5](https://github.com/a-h/templ/tree/main/examples).

---

## **Learning Resources**

- **Official GitHub Repository**  
  [https://github.com/a-h/templ](https://github.com/a-h/templ)  
  *Read the README and explore the code.*  
  > **Source**: [1](https://github.com/a-h/templ/blob/main/README.md)

- **Templ Wiki / Docs**  
  [Templ Wiki](https://github.com/a-h/templ/wiki)  
  *Contains official guides, advanced usage, and examples.*  
  > **Source**: [2](https://github.com/a-h/templ/wiki)

- **Example Projects**  
  [Examples Folder](https://github.com/a-h/templ/tree/main/examples)  
  *Real-world code samples.*  
  > **Source**: [5](https://github.com/a-h/templ/tree/main/examples)

- **Community Articles / Tutorials**  
  - *Search Medium or dev.to for Templ tutorials.*  
  - *Community blog posts often linked in [Discussions](https://github.com/a-h/templ/discussions).*  
  > **Source**: [3](https://github.com/a-h/templ/discussions)

---

## **Media Resources**

- **Community Screencasts**  
  *Check the “Discussions” section or pinned issues on GitHub for user-created video tutorials or demos.*  

*(Templ doesn’t yet have official YouTube presence like Go does, but user-made content is emerging.)*  

---

## **Sources** 
1. **[Templ GitHub Repository - README](https://github.com/a-h/templ/blob/main/README.md)**  
   - Official README with an introduction, high-level overview, and basic instructions.  
2. **[Templ Wiki](https://github.com/a-h/templ/wiki)**  
   - Official documentation covering usage patterns, “getting started,” advanced configuration, partials, layouts, etc.  
3. **[Templ GitHub Discussions](https://github.com/a-h/templ/discussions)**  
   - Community Q&A, announcements, and knowledge sharing.  
4. **[Templ Initial Commits History](https://github.com/a-h/templ/commits?after=4d281b7c869e79d7f438a03c5b792851ffb90966+34)**  
   - Shows some of the earliest code, indicating the timeline (circa late 2019).  
5. **[Templ Examples Folder](https://github.com/a-h/templ/tree/main/examples)**  
   - Contains actual code samples for reference, from simple “Hello World” to more complex partial/layout usage.

---
