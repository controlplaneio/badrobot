# example-cp-repo

<!-- toc -->

<!-- use markdown-toc to generate a table of contents here https://github.com/jonschlinkert/markdown-toc -->

<!-- tocstop -->

## Heading


## Appendix


### Sequence Diagram Source Example

```
@startuml
title CP Theme
'skinparam handwritten true
skinparam {
    ArrowColor Black
    NoteColor Black
    NoteBackgroundColor White
    LifeLineBorderColor Black
    LifeLineColor Black
    ParticipantBorderColor Black
    ParticipantBackgroundColor Black
    ParticipantFontColor White
    defaultFontStyle Bold
}

== 1. title ==

"Dev Machine"->Github: commit and push
Github->Jenkins: call webhook,\ntrigger build

Jenkins->"Build Slave": automated trigger:\ncommit

== 2a. image scan ==

Jenkins->"Build Slave": automated trigger:\nimage scan
@enduml
```
