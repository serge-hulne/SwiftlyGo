# SwiftlyGo

> Build beautiful WebAssembly apps in Go, swiftly and declaratively.

---

## 🚀 Overview

**SwiftlyGo** is a minimalist UI framework inspired by SwiftUI, powered by Go and WebAssembly. It provides a declarative way to build interactive web applications using idiomatic Go code.

No HTML, no CSS, no JavaScript.
Just Go.

---

## 🧩 Features

* ✅ **Declarative UI composition**
* 🔁 **Reactive state with Observables**
* 🧱 **Composable widgets**: `Label`, `TextField`, `Button`
* 🎯 **Fluent modifiers**: `.Padding()`, `.Background()`, `.Center()`
* 📦 **Single binary output** (WASM)
* 🔄 **Hot reload dev server** with live-rebuild
* 🔧 **Zero-config startup**

---

## 📦 Installation

```bash
git clone https://github.com/yourname/swiftlygo
cd swiftlygo
go run ./cmd/coredev
```

Make sure you have **Go 1.21+** with WebAssembly support.

---

## ✍️ Example

```go
package main

import (
	"fmt"
	"gocore/core"
	"gocore/ui"
)

func OneWay() {
	source := core.NewObservable("")

	in := ui.NewTextField()
	in.BindTo(source)

	out := ui.NewTextField()

	btn := ui.NewButton("Copy")
	btn.OnClick(func() {
		fmt.Println("Clicked:", source.Get())
		out.SetText(source.Get())
	})

	w := ui.NewWindow()
	w.Add(in, out, btn)
	w.Run()
}
```

---

## 📁 Directory Layout

```
.
├── app/             # Your app code
├── core/            # SwiftlyGo core framework
├── cmd/coredev/     # Live-reloading dev server
├── static/          # HTML shell + wasm_exec.js
└── main.wasm        # Auto-generated WebAssembly binary
```

---

## 🧪 Roadmap

* [x] Fluent styling modifiers
* [x] Observables and bindings
* [x] Derived computed observables
* [x] Layout containers (e.g., VBox, HBox)
* [ ] Router & navigation
* [ ] Component lifecycle hooks
* [ ] Minimal CSS-in-Go system

---

## 👥 Credits

Created with ❤️ by [@serge-hulne](https://github.com/serge-hulne)

Inspired by SwiftUI, Svelte, and the simplicity of Go.

---

## 📝 License

GPL-3.0
