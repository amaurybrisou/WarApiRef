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

<!-- OBSERVATION:apa_lua_utf8_bom_parse_error (promoted:2026-04-14T05:15:34Z) -->
> Source: AdvancedPetAssist | Confidence: HIGH | Promoted: 2026-04-14

- `HIGH`: The WAR engine Lua parser rejects .lua files that begin with a UTF-8 BOM (byte sequence 0xEF 0xBB 0xBF). The error message is: Error Reading Lua buffer <file>: [string ...]:1: '=' expected near '»'. The '»' character is the visual artefact of the BOM rendered in the error log.
  - Guidance: Always save Lua addon files as UTF-8 without BOM. In PowerShell 5.1, do NOT use Set-Content -Encoding UTF8 (it adds a BOM). Use [System.IO.File]::WriteAllText(path, content, [System.Text.UTF8Encoding]::new($false)) instead. To verify: read the first 3 bytes — they must NOT be 239 187 191.
