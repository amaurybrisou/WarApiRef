# Frame ClosetGoblinZoneRowZone

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinZoneRowZone
- Raw Name: $parentZone
- Type: Label
- Parent: ClosetGoblinZoneRow
- Parent Type: Window
- Inherits: none
- Template: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.xml:0`

## Children

- none

## Child Element Types

- none

## Structural Child Types

- Anchors
- EventHandlers
- Size

## Composition Pattern

```xml
<Label font="font_default_text" name="...">
  <Size>
    <AbsPoint x="120" y="24"/>
  </Size>
  <Anchors>
    <Anchor point="topleft" relativePoint="topleft">
      <AbsPoint x="5" y="5"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinZoneWindow.OnClic..."/>
    <EventHandler event="OnRButtonUp" function="ClosetGoblinZoneWindow.OnSetZ..."/>
  </EventHandlers>
</Label>
```

## Attributes

- font: font_default_text
- name: $parentZone

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinZoneWindow.OnClickZoneRow |
| OnRButtonUp | ClosetGoblinZoneWindow.OnSetZoneRowContextMenu |
