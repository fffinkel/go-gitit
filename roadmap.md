# Roadmap

# Ethos

  - Fast
  - Easy to hack on
  - Pluggable
  - Single binary deployment - plugins are baked in

# Name

We need a new name!  Frew and I were brainstorming, and "kele" (Hawiian for "fast") came to mind

# Fetching Pages From Git

We need to be able to fetch pages from Git - rendering them as plain text is fine for now.

# Page Template

We need a nice-looking (ish) page template that wraps the wiki pages' contents.  It should probably look a lot like
gitit's template for familiarity's sake.

# Rendering

We need to be able to render pages as markdown - we're not going to worry about other formats supported by pandoc.

# Login

We need to support single-sign on (SSO) techniques.

# Support All Other Routes

Any route that's marked as "NYI" (for not yet implemented) needs to be, well, implemented.

# Search

Search should be implemented using fulltext search (FTS) so we can benefit from things like stemming and relevance.

At the very least, this functionality should be available via a plugin.
# Plugins

Plugins should be loaded by importing the plugins' packages with a blank identifier in something like `config.go`:

```go
import (
	_ "github.com/go-gitit/my-plugin"
)
```

The API for the plugins and how plugins are loaded (although it'll most certainly happen via `init`) needs to be determined.
