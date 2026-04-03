# Frame ClosetGoblinSetRowTactics

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinSetRowTactics
- Raw Name: $parentTactics
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
<Label font="font_default_text" name="..." textalign="center">
  <Size>
    <AbsPoint x="70" y="24"/>
  </Size>
  <Anchors>
    <Anchor point="topright" relativePoint="topleft" relativeTo="$parentName">
      <AbsPoint x="0" y="0"/>
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
- name: $parentTactics
- textalign: center

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.OnClickSetRow |
| OnRButtonUp | ClosetGoblinCharacterWindow.OnSetRowContextMenu |
