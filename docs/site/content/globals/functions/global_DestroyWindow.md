# DestroyWindow

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 28 addons

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
| Addons seen in | Ace, Aura, AutoMark, BankWindowFix, BuffHead, DAoCBuff, Effigy, Enemy |
| Files seen in | `/workspace/Ace/LibGUI.lua:120`, `/workspace/Aura/Source/Aura.lua:359`, `/workspace/AutoMark/Source/AutoMark.lua:27`, `/workspace/BankWindowFix/Source/BankWindowFix.lua:15`, `/workspace/BuffHead/Container.lua:886`, `/workspace/DAoCBuff/Source/DAoCBuff.lua:767`, `/workspace/DAoCBuff/Source/DAoCBuffHeadFrames.lua:212`, `/workspace/Effigy/EffigySlashCommands.lua:53` |
| Namespaces detected | DestroyWindow |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Destroy, Aura: Aura:DeleteWindow, AutoMark: AutoMark.local.DestroyMarker, AutoMark: DestroyMarker, BankWindowFix: BankWindowFix.Initialize, BuffHead: BuffHeadContainer:Destroy |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 70 |
| Global usage count | 70 |
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
DestroyWindow(windowName)
```

## Description

Observed tearing down runtime-created windows.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "DAoCBuffMessageWindow", "FastInteractWindow", "GetStatsCompareLine"..i |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- Destroys a runtime-created window instance.

## Seen In

- Ace
- Aura
- AutoMark
- BankWindowFix
- BuffHead
- DAoCBuff
- Effigy
- Enemy
- FastInteract
- GCDsaver
- GetStats
- JunkDump
- LibWBToggler
- MapMonster
- Miracle Grow Remix
- PotionBar
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- TidyRoll
- WSCT
- WarTriage
- WoH-Reticle
- nRarity
- wbLeadHelper

## Examples

- Ace: LIBGUI_ELEMENT:Destroy -> DestroyWindow(self.name)
- Aura: Aura:DeleteWindow -> DestroyWindow(windowId)
- AutoMark: AutoMark.local.DestroyMarker -> DestroyWindow(marker.window_name)
- AutoMark: DestroyMarker -> DestroyWindow(marker.window_name)
- BankWindowFix: BankWindowFix.Initialize -> DestroyWindow(SLOTS_NAME)
- BuffHead: BuffHeadContainer:Destroy -> DestroyWindow(self:GetName())

## Related APIs

- none

## Used With

- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](../../window_api/functions/window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [ScrollWindowSetOffset](../../window_api/functions/window_ScrollWindowSetOffset.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function
- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [BankWindow](../tables/table_BankWindow.md) (HIGH 100/100) - Global Table
- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
