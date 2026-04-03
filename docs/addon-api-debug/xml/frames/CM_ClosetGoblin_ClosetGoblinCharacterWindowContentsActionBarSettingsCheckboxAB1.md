# Frame ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB1

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB1
- Raw Name: $parentCheckboxAB1
- Type: Button
- Parent: ClosetGoblinCharacterWindowContentsActionBarSettings
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
<Button inherits="EA_Button_DefaultCheckBox" layer="overlay" name="...">
  <Anchors>
    <Anchor point="bottom" relativePoint="top" relativeTo="$parentLabelAB1">
      <AbsPoint x="0" y="0"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.H..."/>
  </EventHandlers>
</Button>
```

## Attributes

- inherits: EA_Button_DefaultCheckBox
- layer: overlay
- name: $parentCheckboxAB1

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.HotbarChangeToggled1 |
