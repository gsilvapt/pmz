# Poor Man Zettelkasten  
Inspired by my [blog post on how to build a bloat-free Zettelkasten](https://gsilvapt.me/posts/building-a-zettelkasten-the-simple-way/) 
I decided to write a simple CLI tool that contains all the commands I usually do manually.

This is a toy project, although it's production ready. Feel free to use it, report bugs in the issues page, suggest 
features in the same page, or even contribute with code.

# USAGE  
After installing the CLI with `go install github.com/gsilvapt/pmz`, you will have access to `pmz new|search|update|save` 
commands. Each command has its own set of flags, according to what makes sense. 

To maximize productivity, this project is configuration based, hence you can find a `.pmz.yaml.example` file with variables 
you will need to update to your own use case. Make a copy in your home directory and remove the `.example`.

* `pmz new` adds a new note - it creates a directory with the current timestamp and a `README.md` file inside.
* `pmz search` allows searching for notes. It is integrated with grep (unix power), so it requires `grep` to be 
installed. This command also contains a set of flags that make sense to it, so feel free to use the `--help` tag for 
more.
* `pmz update` if a git repository is configured (GitHub, GitLab, or whatever else, as long as you provide a token that 
can pull that repository), you can pull changes to your local repository.
* `pmz save` if a git repository is configured (GitHub, GitLab, or whatever else, as long as you provide a token that 
can push that repository), you can push your changes to the remote repository.


# CONTRIBUTING  
When I started building this, I wanted a way to keep a [Zettelkasten](https://zettelkasten.de/) simple and bloat free - 
let's keep it that way! Nowadays, there are too many options, full of features I don't use or need. Plus, being able to 
integrate with everyday tools to me is key (things like `grep` for searching, `pandoc` to generate reports, etc).

To contribute with code, please check the issues tab and find an issue to work on. Either a bug or a feature, please 
start by opening an issue. This will help keep things organised.

# LICENSE  
This project is licensed using GNU GPLv3. For more, read [the license](./LICENSE).
