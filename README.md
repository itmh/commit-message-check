# Commit message hook [![Build Status](https://travis-ci.org/leominov/commit-msg.svg?branch=master)](https://travis-ci.org/itmh/commit-msg) [![Release](https://img.shields.io/github/release/leominov/commit-msg.svg)](https://github.com/itmh/commit-msg/releases/latest)

Простая утилита для проверки коммит-сообщения на соответствие некоторым нормам:

* Сообщение должно состоять из заголовка и содержимого, отделённого пустой строкой;
* Заголовок должен содержать номер задачи (например: DIRI55-100, OPS-1000, BAM-50000);
* Заголовок должен начинаться с заглавной буквы;
* Заголовок должен быть в длину не более 50 символов;
* Заголовок не должен заканчиваться точкой;
* Содержимое должно начинаться с заглавной буквы;
* Содержимое должно заканчиваться точкой.

# Установка

## Linux, Mac OS, FreeBSD

### Локально

Скопировать утилиту в каталог `.git/hooks/` необходимого репозитория, переименовать в `commit-msg`.

### Глобально

```
$ mkdir -p ~/.git_template/hooks && git config --global init.templatedir '~/.git_template'
```

Скопировать утилиту в каталог `~/.git_template/hooks/`, переименовать в `commit-msg`.

## Windows

Предполагается, что утилита сохранена в каталог загрузок.

### Локально

В необходимом репозитории в каталоге `.git/hooks/` создать файл `commit-msg` с содержимым:

```
#!/bin/sh

$HOME/Downloads/commit-msg-Windows-x86_64.exe "$1"
```
