# Frame ClosetGoblinCharacterWindowContentsEquipmentShowCloak

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowContentsEquipmentShowCloak
- Raw Name: $parentShowCloak
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
    <Anchor point="topleft" relativePoint="topleft" relativeTo="$parentSlot13">
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
- name: $parentShowCloak

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.ShowCloak |
| OnMouseOver | ClosetGoblinCharacterWindow.ShowShowCloakOnly |
| OnMouseOverEnd | ClosetGoblinCharacterWindow.HideCloakOptions |
