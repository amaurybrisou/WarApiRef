# EA_ChatWindow.InsertText

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, Calling, DetauntHelper, MapPin, NerfedButtons, SOR, ZonePOP |
| Files seen in | AdvancedRenownTrainingImportExport.lua, CallingList.lua, NerfedTalks.lua, NerfedUtils.lua, SORObjs.lua, SORPager.lua, SORRealms.lua, Source/Nerfed.lua |
| Namespaces detected | EA_ChatWindow |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: ExportToLink, Calling: OnListEntryClick, DetauntHelper: HandleNerfEditCommand, MapPin: LButtonUp, MapPin: RButtonUp, MapPin: SendLink |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 19 |
| Global usage count | 19 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
EA_ChatWindow.InsertText(arg1)
```

## Description

Observed as a global function across 7 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: BOInfo[bar].name..L " ("..SORS.OBJBoon[SORS.Info[BOInfo[bar].id].objtype]..L ") ", CreateHyperLink(AddLink,L "WayPoints: "..towstring(GetStringFromTable("ZoneNames",MapPin.Pins.WayPoint[1].Zone))..L " ["..towstring(MapCordsX)..L "K,"..towstring(MapCordsY)..L "K]",{30,255,100},{}), CreateHyperLink(DataText,FormatTitle(PinData.Title)..L ": "..towstring(GetStringFromTable("ZoneNames",PinData.Zone))..L " ["..towstring(MapCordsX)..L "K,"..towstring(MapCordsY)..L "K]",{150,200,255},{}) |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdvancedRenownTrainer
- Calling
- DetauntHelper
- MapPin
- NerfedButtons
- SOR
- ZonePOP

## Examples

- AdvancedRenownTrainer: ExportToLink -> EA_ChatWindow.InsertText(hl)
- Calling: OnListEntryClick -> EA_ChatWindow.InsertText(braggingText)
- DetauntHelper: HandleNerfEditCommand -> EA_ChatWindow.InsertText(str)
- MapPin: LButtonUp -> EA_ChatWindow.InsertText(CreateHyperLink(DataText,FormatTitle(PinData.Title)..L ": "..towstring(GetStringFromTable("ZoneNames",PinData.Zone))..L " ["..towstring(MapCordsX)..L "K,"..towstring(MapCordsY)..L "K]",{150,200,255},{}))
- MapPin: LButtonUp -> EA_ChatWindow.InsertText(CreateHyperLink(AddLink,L "WayPoints: "..towstring(GetStringFromTable("ZoneNames",MapPin.Pins.WayPoint[1].Zone))..L " ["..towstring(MapCordsX)..L "K,"..towstring(MapCordsY)..L "K]",{30,255,100},{}))
- MapPin: RButtonUp -> EA_ChatWindow.InsertText(CreateHyperLink(AddLink,L "WayPoints: "..towstring(GetStringFromTable("ZoneNames",MapPin.Pins.WayPoint[1].Zone))..L " ["..towstring(MapCordsX)..L "K,"..towstring(MapCordsY)..L "K]",{30,255,100},{}))

## Used With

- [DialogManager.MakeTextEntryDialog](global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [Icons.GetCareerIconIDFromCareerLine](global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [EA_Window_WorldMap.ShowZone](global_EA_Window_WorldMap.ShowZone.md) (HIGH 90/100) - Global Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
