# Frame ClosetGoblinActionBarPageSelectorUp

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinActionBarPageSelectorUp
- Raw Name: $parentUp
- Type: Button
- Parent: ClosetGoblinActionBarPageSelector
- Parent Type: Window
- Inherits: EA_Button_DefaultUp
- Template: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.xml:0`

## Children

- none

## Child Element Types

- none

## Structural Child Types

- Anchors
- EventHandlers

## Composition Pattern

```xml
<Button inherits="EA_Button_DefaultUp" name="...">
  <Anchors>
    <Anchor point="topleft" relativePoint="topleft" relativeTo="$parent"/>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.H..."/>
    <EventHandler event="OnMouseOver" function="ActionBars.OnMouseoverHotbarP..."/>
  </EventHandlers>
</Button>
```

## Attributes

- inherits: EA_Button_DefaultUp
- name: $parentUp

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.HotbarPageUpProxy |
| OnMouseOver | ActionBars.OnMouseoverHotbarPageUp |
