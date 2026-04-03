# Frame ClosetGoblinSetRowName

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinSetRowName
- Raw Name: $parentName
- Type: Label
- Parent: ClosetGoblinSetRow
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
    <AbsPoint x="275" y="24"/>
  </Size>
  <Anchors>
    <Anchor point="topleft" relativePoint="topleft">
      <AbsPoint x="5" y="5"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.O..."/>
    <EventHandler event="OnRButtonUp" function="ClosetGoblinCharacterWindow.O..."/>
  </EventHandlers>
</Label>
```

## Attributes

- font: font_default_text
- name: $parentName

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.OnClickSetRow |
| OnRButtonUp | ClosetGoblinCharacterWindow.OnSetRowContextMenu |
