# Algorand Documentation Website

## Setup

- `pip install -r requirements.txt`

## Contributing
For more information about how to contribute documentation see the mkdocs documentation:
http://www.mkdocs.org/user-guide/writing-your-docs/

#### Navigation Bar
Because we have so many generated documents, we aren't able to use the default mkdocs.yml navigation bar, instead we're using [MkDocs Awesome Pages Plugin](https://github.com/lukasgeiter/mkdocs-awesome-pages-plugin).

This allows you to do a few things, but the most important thing to know is how to set the order of pages in the navigation bar. Each docs directory may contain an optional **.pages** file, for example [here is the one in the root directory](docs/.pages). You can specify directories or markdownfiles. If the navigation bar should have a different title than the directory, the **.pages** file can also include a title, [as shown here](docs/Reference-Docs/teal/.pages).

## Testing

```
mkdocs serve
```

## Deploying

```
mkdocs gh-deploy
```

## Generated Docs

There are some tools to help with this in the **scripts** directory.

- goal (go cli, spf/cobra)
- algokey (go cli, spf/cobra)
- kmd (swagger, ???)
- algod (swagger, ???)


