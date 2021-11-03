# Poor Men Zettelkasten  
Inspired by my [blog post on how to build a bloat-free Zettelkasten](https://gsilvapt.me/posts/building-a-zettelkasten-the-simple-way/) 
I decided to write a simple CLI tool that contains all the commands I usually do manually.

This is a toy project. Feel free to use it, report bugs in the issues page, suggest features there, or even contribute 
with code. The whole concept is heavily inspired in [@rwxrob](https://github.com/rwxrob) `zet` scripts.

The development of this project will always go in-line with my needs of a zettelkasten.


# Installation & Configuration  
After installing the CLI with `go install github.com/gsilvapt/pmz`, you will have access to `pmz new|search` 
commands. Each command has its own set of flags, according to what makes sense. It leverages `cobra` framework, so 
there are helpers for all commands.


# Usage 
To maximize productivity, this project is configuration based, hence you can find a `.pmz.yaml.example` file with variables 
you will need to update to your own use case. Make a copy in your home directory and remove the `.example` suffix.

* `pmz new` adds a new note - it creates a directory with the current timestamp and a `README.md` file inside.
* `pmz search` allows searching for keywords in notes' titles. Afterwards, user can either open or request more to 
display file contents.

## Templates  
The program supports a custom template for a new note. It can have the format you want, as long as it contains the 
`{{.Title}}` variable. If in doubt of what that looks like, please check [the template that exists](./templates/new_note) 
in this repository. 

You are free to use it too, either by cloning this project or by copying it manually to file in your system. 

To use your custom template, please provide its full path in the configuration file. Check the 
[sample configuration file](./.pmz.yaml.example) for an example.


# CONTRIBUTING  
When I started building this, I wanted a way to keep a [Zettelkasten](https://zettelkasten.de/) simple and bloat free - 
let's keep it that way! Nowadays, there are too many options, full of features I don't use or need. Plus, being able to 
integrate with everyday tools to me is key (things like `grep` for searching, `pandoc` to generate reports, etc).

To contribute with code, please check the issues tab and find an issue to work on. Either a bug or a feature, please 
start by opening an issue. This will help keep things organised.


# LICENSE  
This project is licensed using Apache 2. For more information, read [the license](./LICENSE).
