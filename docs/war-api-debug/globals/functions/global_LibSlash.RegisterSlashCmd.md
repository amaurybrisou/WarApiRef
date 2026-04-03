# LibSlash.RegisterSlashCmd

- Category: Global Function
- Confidence level: MEDIUM
- Confidence score: 55/100
- Seen in: 1 addons

## Confidence Assessment

- Level: MEDIUM

- Score: 55/100

- Rationale: Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition, appears in slash command registration patterns.

## Evidence Signals

- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Appears in slash command registration patterns: Observed in shared command registration flows.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin |
| Files seen in | ClosetGoblin.lua |
| Namespaces detected | LibSlash |
| Source kinds | lua_calls |
| Example locations | CM_ClosetGoblin: Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | yes |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
LibSlash.RegisterSlashCmd(slashName, handler)
```

## Description

Observed wiring slash commands through a shared command-registration table.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| slashName | Observed as a slash command token. | Observed values: "cg" |
| handler | Observed as a command handler callback. | Observed values: ClosetGoblin.OnSlashCommand |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: Initialize -> LibSlash.RegisterSlashCmd("cg", ClosetGoblin.OnSlashCommand)

## Affects

- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
