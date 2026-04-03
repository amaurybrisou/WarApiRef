# Frame ClosetGoblinCharacterWindowContentsSortTacticsButton

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowContentsSortTacticsButton
- Raw Name: $parentSortTacticsButton
- Type: Button
- Parent: ClosetGoblinCharacterWindowContents
- Parent Type: Window
- Inherits: EA_Button_ListSort
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
<Button id="2" inherits="EA_Button_ListSort" name="...">
  <Size>
    <AbsPoint x="90" y="40"/>
  </Size>
  <Anchors>
    <Anchor point="topright" relativePoint="topleft" relativeTo="$parentSortNameButton">
      <AbsPoint x="0" y="0"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.O..."/>
  </EventHandlers>
</Button>
```

## Attributes

- id: 2
- inherits: EA_Button_ListSort
- name: $parentSortTacticsButton

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.OnClickSortTacticsButton |
