---
layout: default
title: Command Reference
---

# Command Reference

## `revisio`

```
revisio [FLAGS] <SUBCOMMAND> ...
```

| Flag        | Description         |
| ----------- | -------------------- |
| `--verbose` | Log verbose output   |

## `revisio create`

```
revisio create [FLAGS] <KEY> <VALUE>
```

Create a flashcard.

| Flag           | Description                |
| -------------- | --------------------------- |
| `-k, --key`    | The key of the flashcard    |
| `-v, --value`  | The value of the flashcard  |

## `revisio read`

```
revisio read [FLAGS] <KEY> <VALUE>
```

Read a flashcard.

| Flag         | Description               |
| ------------ | -------------------------- |
| `-k, --key`  | The key of the flashcard   |

## `revisio version`

```
revisio version
```

Print version information. Takes no flags beyond the global `--verbose`.
