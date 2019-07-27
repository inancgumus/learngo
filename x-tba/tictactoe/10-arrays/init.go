package main

// initCells initialize the played cells to empty
func initCells() {
	for i := range cells {
		cells[i] = emptyCell
	}
}
