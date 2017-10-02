package react

const testVersion = 5

// Spreadsheet manages the creation of cells.
type Spreadsheet struct {
	// TODO: keep a slice of input and compute cells
}

// New creates a Spreadsheet.
func New() *Spreadsheet {
	s := new(Spreadsheet)
	return s
}

// CreateInput creates an input cell linked into the reactor
// with the given initial value.
func (s *Spreadsheet) CreateInput(data int) InputCell {
	return &SpreadsheetCell{data: data}
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (s *Spreadsheet) CreateCompute1(c Cell, callback func(int) int) ComputeCell {
	input := c.(*SpreadsheetCell)
	compute := SpreadsheetCell{}
	compute.AddComputeFunc1(callback)
	compute.ObserveCell(input)
	return &compute
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (s *Spreadsheet) CreateCompute2(c1 Cell, c2 Cell, callback func(int, int) int) ComputeCell {
	compute := SpreadsheetCell{}
	compute.AddComputeFunc2(callback)
	compute.ObserveCell(c1)
	compute.ObserveCell(c2)
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
	data       int
	observedBy *SpreadsheetCell

	computeFunc1 func(int) int
	computeFunc2 func(int, int) int
	callbacks    []func(int)
	observing    [2]Cell
}

// RegisterComputeCell adds references to a compute cell to the parent cell.
func (sc *SpreadsheetCell) RegisterComputeCell(computeCell *SpreadsheetCell) {
	sc.observedBy = computeCell
	sc.recalculateAll()
}

// SetValue sets the value of the cell.
func (sc *SpreadsheetCell) SetValue(data int) {
	sc.data = data
	sc.recalculateAll()
}

// Value returns the cell's data (whether static or computed).
func (sc *SpreadsheetCell) Value() int {
	return sc.data
}

func (sc *SpreadsheetCell) recalculateAll() {
	if sc.observedBy != nil {
		sc.observedBy.recalculate()
	}
}

// ObserveCell registers a cell for notification upon change.
func (sc *SpreadsheetCell) ObserveCell(cell Cell) {
	if sc.observing[0] == nil {
		sc.observing[0] = cell
	} else {
		sc.observing[1] = cell
	}
	cell.(*SpreadsheetCell).RegisterComputeCell(sc)
}

func (sc *SpreadsheetCell) recalculate() {
	original := sc.Value()
	switch {
	case sc.computeFunc2 != nil && sc.observing[1] != nil:
		value := sc.computeFunc2(sc.observing[0].Value(), sc.observing[1].Value())
		sc.SetValue(value)
	case sc.computeFunc1 != nil && sc.observing[0] != nil:
		value := sc.computeFunc1(sc.observing[0].Value())
		sc.SetValue(value)
	default:
		// sc.observing is nil
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

// AddComputeFunc1 adds a single argument callback which will be called when the value changes.
func (sc *SpreadsheetCell) AddComputeFunc1(callback func(int) int) {
	sc.computeFunc1 = callback
}

// AddComputeFunc2 adds a two argument callback which will be called when the value changes.
func (sc *SpreadsheetCell) AddComputeFunc2(callback2 func(int, int) int) {
	sc.computeFunc2 = callback2
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
