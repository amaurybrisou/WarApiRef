# Frame ClosetGoblinCharacterWindowContentsCheckboxUseItemRack

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowContentsCheckboxUseItemRack
- Raw Name: $parentCheckboxUseItemRack
- Type: Button
- Parent: ClosetGoblinCharacterWindowContents
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
    <Anchor point="right" relativePoint="left" relativeTo="$parentLabelUseItemRack">
      <AbsPoint x="0" y="0"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.U..."/>
  </EventHandlers>
</Button>
```

## Attributes

- inherits: EA_Button_DefaultCheckBox
- layer: overlay
- name: $parentCheckboxUseItemRack

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.UseItemRackToggled |
