# towstring

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 9 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, NPC Item Sale Price, PartyCast, Soloq, TidyChat, TidyRoll, TimeToDie, ZCurse_Profiler |
| Files seen in | `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:1017`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:111`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:1138`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:1723`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:1756`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:1915`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:1952`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:3343` |
| Namespaces detected | towstring |
| Source kinds | lua_calls |
| Example locations | InfoScroller: InfoScroller.About, InfoScroller: InfoScroller.Test, InfoScroller: LIBGUI_Button:SetText, InfoScroller: LIBGUI_Combobox:Add, InfoScroller: LIBGUI_Combobox:Select, InfoScroller: LIBGUI_Label:SetText |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 166 |
| Global usage count | 166 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | no |
| Slash command presence | no |
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
towstring(arg1)
```

## Description

Observed as a global function across 9 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: "", "%%^", "%^" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- InfoScroller
- NPC Item Sale Price
- PartyCast
- Soloq
- TidyChat
- TidyRoll
- TimeToDie
- ZCurse_Profiler
- minesweep

## Examples

- InfoScroller: InfoScroller.About -> towstring(CreateHyperLink(L "0",L "Sullemunk",{255,255,40},{}))
- InfoScroller: InfoScroller.About -> towstring(version)
- InfoScroller: InfoScroller.Test -> towstring(CreateHyperLink(L "0",L "Test-Friend",{255,75,75},{}))
- InfoScroller: InfoScroller.Test -> towstring(CreateHyperLink(L "0",L "Test-Enemy",{75,75,255},{}))
- InfoScroller: InfoScroller.Test -> towstring(CreateHyperLink(L "0",L "Butcher's Pass",{75,255,75},{}))
- InfoScroller: InfoScroller.Test -> towstring(CreateHyperLink(L "0",L " 2000 Renown",{230,50,250},{}))

## Related APIs

- none

## Used With

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](../../window_api/functions/window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 98/100) - Window Function
- [GetIconData](global_GetIconData.md) (HIGH 71/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
