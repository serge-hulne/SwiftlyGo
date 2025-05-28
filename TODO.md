# TODO

## To test:
- text input (read. write)
- text areas (read, write)
- Observables.
- Derived observables.

## 1. Derived Observables (Computed)

```
Enable reactive computed values:
greeting := core.Derive(func() string {
    return "Hello, " + name.Get()
})
```

## 2. Container Widgets

```
Add VBox, HBox, and Flex layout components:
layout := core.NewVBox(input, button, label)
win.Add(layout)
```

## 3. Component Lifecycle Hooks

```
Optionally:
type Component interface {
	Init()
	Destroy()
}
```

## 4. Event Bus

```
Basic publish-subscribe event system for app-wide communication:
core.Publish("user:logged-in", userID)
core.Subscribe("user:logged-in", func(data any) { ... })
```

## 5. Minimal Router

```
Useful for SPA-style apps with:
```


## 6. Misc:

⚙️ Derive(...).Map(...) chaining
🔀 Minimal routing (core.Route("home"))
♻️ Lifecycle (Init(), Dispose())
🧪 Built-in unit tests for widget behaviors