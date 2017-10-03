package react

const testVersion = 5

// Spreadsheet manages the creation of cells.
type Spreadsheet struct {
}

// New creates a Spreadsheet.
func New() *Spreadsheet {
	return new(Spreadsheet)
}

// CreateInput creates an input cell linked into the reactor
// with the given initial value.
func (s *Spreadsheet) CreateInput(value int) InputCell {
	return &SpreadsheetCell{value: value}
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (s *Spreadsheet) CreateCompute1(c Cell, callback func(int) int) ComputeCell {
	compute := SpreadsheetCell{computeFunc1: callback}
	compute.ObserveCells(c)
	return &compute
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (s *Spreadsheet) CreateCompute2(c1 Cell, c2 Cell, callback func(int, int) int) ComputeCell {
	compute := SpreadsheetCell{computeFunc2: callback}
	compute.ObserveCells(c1, c2)
	return &compute
}

// SpreadsheetCanceler manages registered auxiliary callbacks so they can be deleted.
type SpreadsheetCanceler struct {
	cell  *SpreadsheetCell
	index int
}

// Cancel removes the callback.
func (sc SpreadsheetCanceler) Cancel() {
	sc.cell.RemoveCallback(sc.index)
}

// SpreadsheetCell has a changeable value, changing the value triggers updates to
// other cells.
type SpreadsheetCell struct {
	value      int
	observedBy []*SpreadsheetCell
	observing  []Cell

	computeFunc1 func(int) int
	computeFunc2 func(int, int) int
	callbacks    []func(int)
}

// RegisterComputeCell adds references to a compute cell to the parent cell.
func (sc *SpreadsheetCell) RegisterComputeCell(computeCell *SpreadsheetCell) {
	sc.observedBy = append(sc.observedBy, computeCell)
	sc.recalculateAll()
}

// SetValue sets the value of the cell.
func (sc *SpreadsheetCell) SetValue(value int) {
	sc.value = value
	sc.recalculateAll()
}

// Value returns the cell's data (whether static or computed).
func (sc *SpreadsheetCell) Value() int {
	return sc.value
}

// ObserveCells registers one or more cell for notification upon change.
func (sc *SpreadsheetCell) ObserveCells(cells ...Cell) {
	for _, cell := range cells {
		sc.observing = append(sc.observing, cell)
		cell.(*SpreadsheetCell).RegisterComputeCell(sc)
	}
}

func (sc *SpreadsheetCell) recalculateAll() {
	for _, observer := range sc.observedBy {
		observer.recalculate(sc)
	}
}

func (sc *SpreadsheetCell) recalculate(caller *SpreadsheetCell) {
	original := sc.Value()
	switch {
	case sc.computeFunc2 != nil && len(sc.observing) > 1:
		value := sc.computeFunc2(sc.observing[0].Value(), sc.observing[1].Value())
		sc.SetValue(value)
	case sc.computeFunc1 != nil:
		value := sc.computeFunc1(sc.observing[0].Value())
		sc.SetValue(value)
	}
	// Run auxiliary callbacks if computed value changed
	if sc.Value() != original {
		for _, callback := range sc.callbacks {
			if callback != nil {
				callback(sc.Value())
			}
		}
	}
}

// AddCallback registers and auxiliary callback which will be called with the
// computed value after it changes.
func (sc *SpreadsheetCell) AddCallback(callback func(int)) Canceler {
	sc.callbacks = append(sc.callbacks, callback)
	return SpreadsheetCanceler{cell: sc, index: len(sc.callbacks) - 1}
}

// RemoveCallback deletes a registered auxiliary callback.
func (sc *SpreadsheetCell) RemoveCallback(index int) {
	// TODO Re-build splice without this callback
	sc.callbacks[index] = nil
}
