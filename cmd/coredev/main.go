package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var serveDir string
var reloadNeeded atomic.Bool
var resetIdleTimer = make(chan struct{}, 1)

const defaultIndexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Go WASM App</title>
</head>
<body>
    <h1>Go WASM App</h1>
    <script src="wasm_exec.js"></script>
    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
        });

        setInterval(() => {
            fetch('/reload-check').then(res => res.text()).then(flag => {
                if (flag.trim() === 'reload') {
                    console.log('🔄 Reloading page...');
                    window.location.reload();
                }
            });
        }, 1000);
    </script>
</body>
</html>`

func main() {
	port := flag.String("port", "8090", "Port to run the server on")
	dir := flag.String("dir", "static", "Directory to serve")
	noWatch := flag.Bool("no-watch", false, "Disable watching for file changes")
	noBuild := flag.Bool("no-build", false, "Disable rebuild at startup")
	buildRelease := flag.Bool("build-release", false, "Build once without watching or live-reload")
	flag.Parse()

	serveDir = *dir
	ensureWasmExec()
	ensureFileExists(filepath.Join(serveDir, "index.html"), defaultIndexHTML)

	if *buildRelease {
		rebuild(serveDir)
		fmt.Println("✅ Release build completed. Files are in:", serveDir)
		return
	}

	if *noBuild && *noWatch {
		log.Println("⚠️ Warning: both --no-build and --no-watch enabled; make sure main.wasm exists")
	}

	if !*noBuild {
		rebuild(serveDir)
	}

	if !*noWatch {
		go watchAndRebuild()
		go idleTimeoutChecker()
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/reload-check", func(w http.ResponseWriter, r *http.Request) {
		if reloadNeeded.Load() {
			reloadNeeded.Store(false)
			w.Write([]byte("reload"))
		} else {
			w.Write([]byte("noop"))
		}
	})

	r.Handle("/*", http.FileServer(http.Dir(*dir)))

	log.Printf("🚀 Golid dev server running at http://localhost:%s (serving from %s)", *port, *dir)
	log.Fatal(http.ListenAndServe(":"+*port, r))
}

func watchAndRebuild() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if strings.HasPrefix(path, "./cmd") {
				return filepath.SkipDir
			}
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&(fsnotify.Write|fsnotify.Create) != 0 && strings.HasSuffix(event.Name, ".go") {
				resetIdleTimer <- struct{}{}
				rebuild(serveDir)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Watcher error:", err)
		}
	}
}

func idleTimeoutChecker() {
	timer := time.NewTimer(5 * time.Minute)
	for {
		select {
		case <-resetIdleTimer:
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(5 * time.Minute)
		case <-timer.C:
			fmt.Print("\n⚠️  No changes detected for 5 minutes!")
			os.Exit(0)
		}
	}
}

func rebuild(dir string) {
	cmd := exec.Command("go", "build", "-o", filepath.Join(dir, "main.wasm"), "./app")
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()
	if err := cmd.Run(); err != nil {
		log.Println("❌ Build failed:", err)
	} else {
		log.Println("✅ Build succeeded")
		reloadNeeded.Store(true)
	}
}

func ensureFileExists(filename, content string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
			log.Fatalf("❌ Failed to create %s: %v", filename, err)
		}
		log.Printf("✅ Created %s", filename)
	}
}

func ensureWasmExec() {
	destPath := filepath.Join(serveDir, "wasm_exec.js")
	if _, err := os.Stat(destPath); err == nil {
		return
	}

	out, err := exec.Command("go", "env", "GOROOT").Output()
	if err != nil {
		log.Fatalf("❌ Failed to get GOROOT: %v", err)
	}
	goroot := strings.TrimSpace(string(out))

	paths := []string{
		filepath.Join(goroot, "lib", "wasm", "wasm_exec.js"),
		filepath.Join(goroot, "misc", "wasm", "wasm_exec.js"),
	}

	var wasmPath string
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			wasmPath = path
			break
		}
	}
	if wasmPath == "" {
		log.Fatal("❌ Could not find wasm_exec.js in known GOROOT paths")
	}

	input, err := os.ReadFile(wasmPath)
	if err != nil {
		log.Fatalf("❌ Failed to read wasm_exec.js: %v", err)
	}
	if err := os.WriteFile(destPath, input, 0644); err != nil {
		log.Fatalf("❌ Failed to write wasm_exec.js to %s: %v", destPath, err)
	}

	log.Println("✅ Copied wasm_exec.js from:", wasmPath)
}
