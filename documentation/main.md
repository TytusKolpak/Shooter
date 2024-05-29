# Project

This document will try to describe the ways the autor built it's structure and decides how it should be build.

## Standards of development

1. Always keep the pattern of each file uniform so that it is easy to maintain any file in the same fashion.
2. Keep the Update and Draw functions at the very top of the file.
3. Any external parameters which can be excluded from the structure of the entity should be keep tin the `parameters.go` file. If possible couple them in order which is uniform with their purpose - like all which regard player next to one another, then an empty line and then all which regard enemy.
4. Always try to put reusable code and complex code outside of the Update and Draw functions of its file for clarity of these main functions.
5. If possible try to order these helper functions according to their importance or alphabetically.
6. Should a reference to an instance of a structure be referenced - assign it to a variable to reduce the number of foreign field accesses.
7. If a complex element is created and is considered as non-self-explanatory do put a concise description of it here in the documentation.

## Structure

The entry point o this project is the `main.go` file in it's root directory. It initializes the game. All other logical elements are placed in the `entities` directory. Each file in this directory considers a single logical unit of the game. Each file contains unique logic applied to according element in the game.

Most of the files do follow a fixed convention of content. They first contain the function regarding the Update of subject element. This function dictates how it should behave. Then there is a Draw function which dictates how it should be displayed on the screen. After that Each file is different containing the helper functions used to make these first two ones concise and clear, so that it's easy to understand and manage what they do.

Most of the files are "main" file of the entity type they regard, but there is also a helper file called `parameters.go`. It is a standalone file which gathers all the single element parameters (constant values and variable values) which can be extracted from the all other places for ease of modification.

## Directories

### Documentation

This directory contains all the descriptions that the autor thought to be important to the wellbeing of the project. This includes explanations of project structure, logical ownership and Standards of development. If any more elements than simple text file be used in here like pictures or screenshots of the game, then they will all be contained here.

## Game entities

### Game

This is the most important file in the whole project.

### Enemy

All logic regarding enemy behavior and their display are contained in the `enemy.go` file.

### Player

To be done

### Projectile

To be done

### Image

To be done
