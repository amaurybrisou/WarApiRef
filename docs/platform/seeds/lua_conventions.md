# Lua Conventions Seed

## Runtime Baselines

- Lua 5.0-style constraints apply.
- Prefer defensive table iteration for sparse payloads.

## UI Text Baselines

- UI text should be wstring-safe (`L"..."`, `towstring(...)`).

## Addon Hygiene

- Keep lifecycle hooks symmetrical.
- Keep handler registration and deregistration explicit.
- Encapsulate direct engine calls where possible.

## Uncertainty Rule

- Do not convert inferred behavior into hard guarantees.
