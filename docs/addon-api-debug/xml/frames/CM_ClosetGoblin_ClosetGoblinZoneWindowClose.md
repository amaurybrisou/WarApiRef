# Frame ClosetGoblinZoneWindowClose

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinZoneWindowClose
- Raw Name: $parentClose
- Type: Button
- Parent: ClosetGoblinZoneWindow
- Parent Type: Window
- Inherits: EA_Button_DefaultWindowClose
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
<Button inherits="EA_Button_DefaultWindowClose" name="...">
  <Anchors>
    <Anchor point="topright" relativePoint="topright" relativeTo="$parentTitleBar">
      <AbsPoint x="0" y="0"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinZoneWindow.Hide"/>
  </EventHandlers>
</Button>
```

## Attributes

- inherits: EA_Button_DefaultWindowClose
- name: $parentClose

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinZoneWindow.Hide |
