# Security Policy

## Supported Versions

| Version | Supported |
| ------- | --------- |
| 2.x     | yes       |
| 1.x     | no        |

Fixes are released from the latest 2.x version. Older major versions get no
updates.

## Reporting a Vulnerability

Do not open a public issue for a security problem.

Report it privately here:
[Report a vulnerability](https://github.com/DavidBelicza/TextRank/security/advisories/new)

Include the affected version, what an attacker can do, and code that shows the
problem when you have it.

You get an answer within 14 days. When the report is accepted, a fix is released
in a new 2.x version and the report is credited in the release notes unless you
ask otherwise.

## Scope

This is a text ranking library without network or file access. The realistic
risks are crashes or extreme memory and CPU use caused by crafted input text.
Reports of that kind are in scope.
