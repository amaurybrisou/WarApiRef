# SystemData.Events.UPDATE_PROCESSED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 17 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, Ace, Aura, BuffHead, DAoCBuff, Enemy, GCDsaver, LibWBToggler, LoyalPet, PeaceOut, RetAlert, Shinies, Swift Assist, TexturedButtons, TimeInQueue, TortallDPSCore, WarTriage, WoH-Reticle, Aura, BuffHead, DAoCBuff, Enemy, LibWBToggler, LoyalPet |
| Files seen in | `/workspace/Ace/AceAddon-3.0.lua:591`, `/workspace/Aura/Source/TargetInfoFix.lua:54`, `/workspace/BuffHead/Setup/ContainerDemo.lua:197`, `/workspace/BuffHead/Setup/ContainerDemo.lua:223`, `/workspace/BuffHead/Setup/LayoutDemo.lua:147`, `/workspace/BuffHead/Setup/LayoutDemo.lua:166`, `/workspace/BuffHead/Setup/SetupEffectCache.lua:373`, `/workspace/BuffHead/Setup/SetupEffectCache.lua:386` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | AceAddon_OnUpdate_DONOTTOUCH, ActionBarsUpdated, BuffHead.Setup.ContainerDemo.Disable, BuffHead.Setup.ContainerDemo.Enable, BuffHead.Setup.ContainerDemo.OnUpdate, BuffHead.Setup.Demo.Disable |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 48 |
| Global usage count | 48 |
| Local definition count | 0 |
| Documentation references | 4 |
| Initialization flow references | 12 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
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

## Description

Observed SystemData field used by 17 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- Ace
- Ace, Aura, BuffHead, DAoCBuff, Enemy, GCDsaver, LibWBToggler, LoyalPet, PeaceOut, RetAlert, Shinies, Swift Assist, TexturedButtons, TimeInQueue, TortallDPSCore, WarTriage, WoH-Reticle
- Aura
- BuffHead
- DAoCBuff
- Enemy
- LibWBToggler
- LoyalPet
- PeaceOut
- RetAlert
- Shinies
- Swift Assist
- TexturedButtons
- TimeInQueue
- TortallDPSCore
- WarTriage
- WoH-Reticle

## Related APIs

- [WindowGetAlpha](../../window_api/functions/window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetHandleInput](../../window_api/functions/window_WindowGetHandleInput.md) (HIGH 100/100) - Window Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [SystemData.Events.UPDATE_PROCESSED](../../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
- [WindowGetAlpha](../../window_api/functions/window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetHandleInput](../../window_api/functions/window_WindowGetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](../../window_api/functions/window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](../../window_api/functions/window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: AceAddon_OnUpdate_DONOTTOUCH, ActionBarsUpdated, BuffHead.Setup.ContainerDemo.Disable, BuffHead.Setup.ContainerDemo.Enable, BuffHead.Setup.ContainerDemo.OnUpdate, BuffHead.Setup.Demo.Disable
