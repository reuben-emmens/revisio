---
layout: default
title: revisio
---

# revisio

`revisio` is a digital flashcard management CLI. It lets you create flashcards
from key/value pairs or import them in bulk from CSV, with a small, consistent
API designed for third-party extensibility.

- [Getting Started](getting-started.html) — install `revisio` and create your first flashcard.
- [Command Reference](commands.html) — flags and usage for every subcommand.
- [Source on GitHub](https://github.com/reuben-emmens/revisio) — issues, releases, and contribution guidelines.

For a quick project overview, see the [repository README](https://github.com/reuben-emmens/revisio#readme).

To render this site locally:

```sh
# 1. Install Ruby and native build tools
sudo apt-get install -y ruby-full build-essential zlib1g-dev

# 2. Install Bundler
gem install bundler

# 3. From the docs/ folder, install gems into a local vendor path
cd docs
bundle config set --local path 'vendor/bundle'
bundle install

# 4. Serve the site with live reload
bundle exec jekyll serve
```

Open [http://localhost:4000](http://localhost:4000) to view the site. Editing any file under `docs/` triggers an automatic rebuild