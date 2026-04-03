# Frame ClosetGoblinCharacterWindowContentsEquipmentShowHelm

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowContentsEquipmentShowHelm
- Raw Name: $parentShowHelm
- Type: Button
- Parent: ClosetGoblinCharacterWindowContentsEquipment
- Parent Type: Window
- Inherits: EA_Button_DefaultCheckBox
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
<Button inherits="EA_Button_DefaultCheckBox" layer="popup" name="...">
  <Anchors>
    <Anchor point="topleft" relativePoint="topleft" relativeTo="$parentSlot9">
      <AbsPoint x="0" y="0"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnMouseOver" function="ClosetGoblinCharacterWindow.S..."/>
    <EventHandler event="OnMouseOverEnd" function="ClosetGoblinCharacterWindow.H..."/>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.S..."/>
  </EventHandlers>
</Button>
```

## Attributes

- inherits: EA_Button_DefaultCheckBox
- layer: popup
- name: $parentShowHelm

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.ShowHelm |
| OnMouseOver | ClosetGoblinCharacterWindow.ShowShowHelmOnly |
| OnMouseOverEnd | ClosetGoblinCharacterWindow.HideShowHelm |
