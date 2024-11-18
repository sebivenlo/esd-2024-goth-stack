#import "@preview/diatypst:0.1.0": *
#set text(font: "Fira Code")
#show: slides.with(
  title: "The GOTH Stack", // Required
  subtitle: "Refreshingly simple web applications",
  authors: ("Mirelle Bogdanski, Joel Eckhardt"),

  // Optional Styling
  ratio: 16/9,
  layout: "medium",
  title-color: blue.darken(75%),
  footer: true,
  counter: true,
)

#outline()

= First Section

== Overview

*What is the GOTH Stack about?*
- KISS approach to web development
- Server side rendering
- State is managed by the server

== Benefits of GOTH over React & co.
- Node sucks tbh (node_modules)
- Performance

== Preparing the workshop
- Google sebivenlo, search for "goth" on the GitHub page

= Components of the GOTH Stack
== GoLang
- Neat and simple language developed at Google by Rob Pike, Ken Thompson (Unix) and Robert Griesemer (V8)
- AOT compiled into a single binary, with all dependencies bundled
- Fast compile times, very performant
- Deliberately small feature set, easy to read and get productive in
// Include a picture of main.go, demonstrating how easy it is to read go
== GoLang (continued)
- Big companies moved at least part of their infrastructure from Node to Go (Uber)
== Templ
- Just a much nicer way to write HTML templates for Go
- Write HTML Blocks as Go functions, inline Go code etc.
- Go comes with its own templating, but dev experience with templ is much nicer
== HTMX
- Focus of todays' workshop
- A different way to think about web applications
- Make state management easy again
- A bit of a meme, but definitely used in 
== Extensions
- Air - hot server reloading (comparable to Vite)
- TailwindCSS - Could almost replace the T in Goth tbh - just a no-brainer to add
- SQLite - is extremely performant and easy to use, fits well into the KISS theme
- JS / React - in case you need more interactivity, feel free to use as much as you want or need. Just leave the state management to HTMX

// / *Term*: Definition