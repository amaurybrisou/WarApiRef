# Frame ClosetGoblinActionBarPageSelectorDown

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinActionBarPageSelectorDown
- Raw Name: $parentDown
- Type: Button
- Parent: ClosetGoblinActionBarPageSelector
- Parent Type: Window
- Inherits: EA_Button_DefaultDown
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
<Button inherits="EA_Button_DefaultDown" name="...">
  <Anchors>
    <Anchor point="bottom" relativePoint="top" relativeTo="$parentCurrentPage">
      <AbsPoint x="0" y="-3"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.H..."/>
    <EventHandler event="OnMouseOver" function="ActionBars.OnMouseoverHotbarP..."/>
  </EventHandlers>
</Button>
```

## Attributes

- inherits: EA_Button_DefaultDown
- name: $parentDown

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.HotbarPageDownProxy |
| OnMouseOver | ActionBars.OnMouseoverHotbarPageDown |
