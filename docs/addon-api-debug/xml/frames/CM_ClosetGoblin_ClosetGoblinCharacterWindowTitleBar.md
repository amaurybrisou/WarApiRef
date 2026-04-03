# Frame ClosetGoblinCharacterWindowTitleBar

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowTitleBar
- Raw Name: $parentTitleBar
- Type: Window
- Parent: ClosetGoblinCharacterWindow
- Parent Type: Window
- Inherits: EA_TitleBar_Default
- Template: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.xml:0`

## Children

- none

## Child Element Types

- none

## Structural Child Types

- Anchors

## Composition Pattern

```xml
<Window inherits="EA_TitleBar_Default" layer="background" name="...">
  <Anchors>
    <Anchor point="topleft" relativePoint="bottomleft" relativeTo="$parentBackground">
      <AbsPoint x="0" y="5"/>
    </Anchor>
    <Anchor point="topright" relativePoint="bottomright" relativeTo="$parentBackground">
      <AbsPoint x="0" y="5"/>
    </Anchor>
  </Anchors>
</Window>
```

## Attributes

- inherits: EA_TitleBar_Default
- layer: background
- name: $parentTitleBar

## Handlers

| Event | Function |
| --- | --- |
