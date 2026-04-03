# Frame ClosetGoblinCharacterWindowContentsSetList

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowContentsSetList
- Raw Name: $parentSetList
- Type: ListBox
- Parent: ClosetGoblinCharacterWindowContents
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
- ListData: populationfunction, table

## Composition Pattern

```xml
<ListBox layer="secondary" name="..." rowcount="100" rowdef="ClosetGoblinSetRow" rowspacing="1" scrollbar="EA_ScrollBar_DefaultVerticalC..." visiblerows="16">
  <Anchors>
    <Anchor point="bottomleft" relativePoint="topleft" relativeTo="$parentSortNameButton">
      <AbsPoint x="0" y="-10"/>
    </Anchor>
    <Anchor point="bottomright" relativePoint="bottomright" relativeTo="$parentSetListBackground">
      <AbsPoint x="0" y="-3"/>
    </Anchor>
  </Anchors>
  <ListData populationfunction="ClosetGoblinCharacterWindow.U..." table="ClosetGoblin.currentProfile.sets">
    <ListColumns>
      <ListColumn format="wstring" variable="name" windowname="Name"/>
      <ListColumn format="wstring" variable="tactics" windowname="Tactics"/>
    </ListColumns>
  </ListData>
</ListBox>
```

## Attributes

- layer: secondary
- name: $parentSetList
- rowcount: 100
- rowdef: ClosetGoblinSetRow
- rowspacing: 1
- scrollbar: EA_ScrollBar_DefaultVerticalChain
- visiblerows: 16

## Handlers

| Event | Function |
| --- | --- |
