# Frame ClosetGoblinZoneRowSet

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinZoneRowSet
- Raw Name: $parentSet
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
<Label font="font_default_text" name="..." textalign="left">
  <Size>
    <AbsPoint x="345" y="24"/>
  </Size>
  <Anchors>
    <Anchor point="topright" relativePoint="topleft" relativeTo="$parentZone">
      <AbsPoint x="0" y="0"/>
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
- name: $parentSet
- textalign: left

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinZoneWindow.OnClickZoneRow |
| OnRButtonUp | ClosetGoblinZoneWindow.OnSetZoneRowContextMenu |
