# Lua encoding caveats

- Category: conventions
- Confidence: HIGH

## Description

Implementation-validated encoding and tooling findings for WAR Lua addon files.

## Involved APIs

- [Text](../xml/element_types/element_Text.md) (HIGH 100/100) - XML Element Type

## Flow Diagram

```text
OnInitialize -> RegisterEventHandler
```

## Example Code

```lua
`HIGH`: The WAR engine Lua parser rejects .lua files that begin with a UTF-8 BOM (byte sequence 0xEF 0xBB 0xBF). The error message is: Error Reading Lua buffer <file>: [string ...]:1: '=' expected near '»'. The '»' character is the visual artefact of the BOM rendered in the error log.
```

## Evidence

- `HIGH`: The WAR engine Lua parser rejects .lua files that begin with a UTF-8 BOM (byte sequence 0xEF 0xBB 0xBF). The error message is: Error Reading Lua buffer <file>: [string ...]:1: '=' expected near '»'. The '»' character is the visual artefact of the BOM rendered in the error log.
- Guidance: Always save Lua addon files as UTF-8 without BOM. In PowerShell 5.1, do NOT use Set-Content -Encoding UTF8 (it adds a BOM). Use [System.IO.File]::WriteAllText(path, content, [System.Text.UTF8Encoding]::new($false)) instead. To verify: read the first 3 bytes — they must NOT be 239 187 191.
