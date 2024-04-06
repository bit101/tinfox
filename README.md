# tinfox project creator

A rewrite of https://github.com/bit101/tinpig in Go.

tinfox is a simple command line utility for quickly creating projects of any kind. Projects are created from templates which are collections of folders and files. The folders and files can contain special tokens that can be replaced with other values when the project is created. You can use an existing temple as-is, modify a template, or create your own from scratch.

tinfox is currently somewhat more simple and more minimal than tinpig. This is based on years of personal use, seeing what features I regularly used and what I did not use.

## Requirements

Go 1.22.1

## Installation

```
go install github.com/bit101/tinfox@latest
```

## First Use

The first time you use tinfox, a `tinfox` directory will be created in the OS standard configuration location. This will contain a `templates` directory with one sample template, and a `config` file.

You can add additional templates to the `templates` directory, which will then be available for use.

Various settings can be applied in the `config` file, including the location of your templates, color theme values, and verbosity. These will be covered further on in this file.

## Use

Simply type `tinfox` on the command line. You will be shown a list of installed templates. Choose one and you'll be walked through the steps to create a project based on that template. The process will include:

1. Choosing a location for the new project. This must be a valid path on your file system. No other file or folder can exist at that location.
2. Depending on the template, you may be presented with one or more tokens to provide values for. These may or may not include default values. Enter a value for each token.
3. If project creation is successful, the path to the new project will be displayed along with any instructions included in the template.

## Configuration

The default `config` file looks like this:

```
{
  "templatesDir": "/home/keith/.config/tinfox/templates",
  "invalidPathChars": "‘“!#$%\u0026+^\u003c=\u003e` ",
  "headerColor": "boldgreen",
  "instructionColor": "yellow",
  "errorColor": "boldred",
  "defaultValueColor": "blue",
  "verbose": true
}
```

The `templatesDir` value will be slightly different based on your OS and user name of course. It defaults to the `templates` folder that is in the same location as the `config` file. But you can change it to point anywhere on your file system. You may likely want to keep your templates in a version controlled or backed up location.

The `invalidPathChars` value determines which characters will not be accepted when prompting the user for a path. Note that `<`, `=` and `>` have been rendered in unicode. Also note that the space character is included. If you want to be able to specify path names with spaces, you can remove that, but it may very well cause problems. 

The next four values specify colors that will be used in prompts. These are pretty obvious, but can be changed to suit your tastes or terminal theme.

The `verbose` value determines how much information is shown while prompting your for values and setting up your project. Expert users may be comfortable with setting this to false.

## Commands and Flags

`tinfox -h` or `tinpig --help` displays general help data.

`tinfox version` displays the current version of tinfox.

`tinfox list` displays all available templates and descriptions as a view-only list.

## Templates



## Differences from tinpig

- tinpig has an interactive configuration function that asks for a user name and email address. These values can be used in templates, but mostly are not and can be set up as tokens in templates in the case they are needed. They have been removed in tinfox. There is no configuration function. A default config is created and it can be edited manually. tinpig also had a reset command to reset the configuration. This does not exist in tinfox.

- tinpig allows you to specify a template and path on the command line. In general it's just easier to choose a template from a list and enter the path when prompted. So this did not carry over to tinfox. These could be useful in scripting or setting up shortcuts though, so it may come back.

- tinpig allows you to create a project in an existing directory, even if it was not empty. It would give a warning and would not overwrite any existing files. tinfox just disallows the use of an existing directory.

- tinpig allows for setting a temporary template directory when calling the command. tinfox does not have that functionality, but may have more advanced template management in the future.

- tinpig does not have the `verbose` feature added in tinfox.

- tinpig came with more built-in templates. tinfox only has one sample template.

### Template Differences

tinfox uses almost the exact same template format as tinpig. The only differences:

- tinpig uses a `tinpig.json` manifest file whereas tinfox uses `tinfox.json`
- tinpig has `preMessage` and `postMessage`. tinfox currently only has `postMessage`. But `preMessage` will probably be added.
- tinpig tokens have an `isPath` property. If a token is a path, it will be checked against invalid path chars when entered. tinfox does not do this (yet).
- Special tokens:
    - `TINPIG_USER_NAME` and `TINPIG_USER_EMAIL` do not exist in tinfox as the user name and email have been removed from config.
    - `TINPIG_PROJECT_PATH` has become `TINFOX_PROJECT_PATH`
    - `TINPIG_PROJECT_DIR` has become `TINFOX_PROJECT_DIR`


## TODO
- Check path for `isPath` tokens.
- Implement `preMessage`.
- Allow user to enter template on command line.
- Allow user to enter project path on command line.
- Allow use of alternate template directory.
- Additional template management features (template categories maybe).
- A verbose flag that overrides the 
- Include more default templates.
