#  Go-MiniWiki: Standard Library File-System Architecture

> A study in building a lightweight, highly concurrent flat-file content engine entirely from scratch using Go's optimised standard library, avoiding all third-party framework dependencies.

---

###  Key Engineering Logic

* **Dependency-Free Architecture:** Built strictly using Go's native primitives (`net/http`, `html/template`, and `os`) to explore the runtime's built-in performance boundaries.
* **Concurrent I/O Boundaries:** Leverages Go's implicit file-handling concurrency models to read, write, and serve markdown/wiki files safely under simulated traffic load.
* **XSS-Safe Server-Side Rendering:** Utilises native `html/template` compilation to render dynamic web views with native, compile-time security context enforcement.

###  Technical Stack
* **Language:** Go (Golang) 1.22+
* **Routing & Handlers:** Native `net/http` standard library
* **Storage:** Local flat-file system architecture (OS file-descriptor read/write loops)
* **Views:** Native Go HTML Templates

---
*For a full breakdown of this project's mechanics, visit my technical showcase at [Waymakers Workshoppe](https://netlify.app).*
