# automata
This a small template to start a celular automata using Raylib with Go.
It includes functions such as `GetNeighbor` and `GetCell`, a `Vec2` struct for positions, and a `Game` struct where you can access the grid of cells easily

## grid
The grid is a 2D array of pointers to the `Cell` struct.  
The cell struct has 2 values: `Color`: the color of the cell, `Alive`: wether the cell is or not alive, and `Position` a Vec2 that stores the cell position.

## game
The `Game` struct has:
- Width: The window width
- Height: The window height
- CellSize: The size of each cell: if Width and Height is 800 and CellSize is 10, there will be 80 cells.
- Grid: The grid of cells.
It also has the `GetNeighbor`, `GetNumberAliveNeighbors`, `Exists` and `GetCell` functions to easily make your rules for the automata.

## consts
The `main.go` file has these consts:
```go
const Width = 800
const Height = 800
const CellSize = 10

// Iterations Per Second
const Speed_IPSecond = 10
```
These are used to set the other values easily.

## how to use:
Click "Use this template"
Clone the new repository.
Edit the init cells in `main.go`. By default even cells are alive.
Edit `update.go` this is the stuff that will happen to the cells each iteration

## screenshots
![image](https://github.com/checkm4ted/automata/assets/146487129/dd7ac5d8-2a32-4773-bb16-60dfcda18d54)
![image](https://github.com/checkm4ted/automata/assets/146487129/c7016c48-3cbb-442d-8800-eab9d41d48a8)
![image](https://github.com/checkm4ted/automata/assets/146487129/bdeeb6ed-cb28-4dbb-9b5c-5dd42f8a0386)
![image](https://github.com/checkm4ted/automata/assets/146487129/20a7a5f4-bda5-4e72-b15d-f3d69f9aaea0)
