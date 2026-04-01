# Window Lifecycle

This page reconstructs the window-facing lifecycle from observed function calls, XML hooks, and lifecycle flow evidence.

## Creation APIs

- CreateWindow (HIGH)
- CreateWindowFromTemplate (HIGH)

## Visibility Hooks

- OnInitialize (HIGH)
- OnShown (HIGH)
- OnHidden (HIGH)

## Destruction APIs

- DestroyWindow (HIGH)
- DoesWindowExist (HIGH)

## Observed Flow

- manifest: Addon manifests declare file load order and bootstrap hooks before Lua runtime logic begins.
- initialize: Initialization hooks create windows, hydrate settings, and bind runtime callbacks.
- saved-variables: Saved-variable roots appear before normal runtime hooks and provide persistent addon state.
- xml: XML windows, templates, and handlers become available as the UI layer is created.
- runtime: Runtime logic pivots into event-driven updates wired through shared event registration APIs.
- update: Optional update hooks provide repeated processing after initialization has completed.
- shutdown: Shutdown hooks unregister commands or handlers and persist addon-owned state.
