## Rules of the Game of Life

1. Any live cell with 0-1 live neighbours dies through underpopulation.
2. Any live cell with 4+ live neighbours dies through overpopulation.
3. Any dead cell with 3 live neighbours is given life.


Neighbours are any adjacent cell, including diagonals. The centre cell itself is not included in any count for its own fate.


Note that two live neighbours sustains an already alive cell, but does not give life to a new one.

All changes to cells happen simultaneously, so a new cell coming to life does not affect its neighbours until the next timestep.


## Advancing the Timestep

There could be two approaches to the calculation of new cells - checking each *live* cell and its neighbours, or just checking *every* cell. They cannot be changed when checked, as all cells must be checked before any can be changed. Each cell check can be wholly independent, and make use of concurrency if computation time becomes an issue.

Overall the problem of duplication arises if, and only if, the game consists of a list of live cells and this is checked to verify neighbours. I.e. there is a list iterated upon and the number of neighbours is the "sum of all items in the list which are close to the cell". The other way is checking "are the cells close to me live".

Something not considered above is to have a list of "cells changed last timestamp." This could be then used to generate a list of "cells that may change this timestamp". Any duplication in this list could be sorted in a few ways, but the solution should be relatively scalable. The most suited way seems to be to have an indicator that the next time stamp has already been calculated. Either have three states: *"alive", "dead"* and *"unknown"*, or just two booleans to represent both *"next state"* and *"have calculated next state"*. Both would require a reset at the start/end of each time frame, and the calculation is simply skipped where applicable. This seems easier and faster than checking for duplicates within the list, especially when a cell in the list may be duplicated up to 8 times.

Both these ideas may be implemented, to provide a comparison. One "dumb" method where every single cell has all of its neighbours checked. Another "smart" method where only cells adjacent to cells which just changed are checked. A lot of this seems to me to be premature optimisation. Something that's well known is that the extra work required to maintain these optimisations can cost much more than they save...

## Game Components

### Board / Area

The Game of Life doesn't have a board as such, as it is supposed to exist on an infinite plane. This is harder to implement, and there are three options that immediately come to mind:
 1. Model everything and show everything. The shown area can zoom out to contain all of the live cells.
 2. Model everything, but only show a window. The area modelled would need to be infinite to be a true game, as "something" outside the window could potentially "return"
 3. Only model what's shown (and maybe a little around it). This is just to have a board of fixed size, and only consider cells within the bounds.

### Cells

Cells can be alive or dead, and can change with each timestep. As such they would need a current state and next state, or there needs to be two sets of cells, one for the current timestamp and one for the next.

### Screen / Display

How the cells are shown. The game is (almost) useless if it cannot be observed.


# Software Architecture

## Packages

It seems that the best way to do things in Go is to use small, distinct packages. I think this generally means subpackages, but will look into it more. If each distinct thing is its own file, it can be its own package.

## Interfaces

As an example, instead of a function printing out to stdout, I can pass in stdout (or its equivalent) via one of its interfaces. That can allow me to create mocks, or change the use of the function without changing it itself. This is an example of Single Responsibility, where the function to write something only writes, nothing more. Even if it only looks like it writes, it might also be formatting and deciding where to write to.

## Starting Somewhere

To try and have a fully concrete plan before writing a line of code is absurd, as I barely understand the language. If I can code a small discrete portion of code, I can get experience with both the language and my implementations.

I will start with the terminal drawing section, as that is vital to see the rest. I pick to draw it to the terminal because I have done it before, and thus know the commands etc.

## Navigating a terminal window