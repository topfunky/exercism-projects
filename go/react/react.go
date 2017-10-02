package react

import "fmt"

const testVersion = 5

type Spreadsheet struct {
	// TODO: keep a slice of input and compute cells
}

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
	compute.AddCallback(callback)
	compute.ObserveCell(input)
	return &compute
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (s *Spreadsheet) CreateCompute2(c1 Cell, c2 Cell, callback func(int, int) int) ComputeCell {
	compute := SpreadsheetCell{}
	compute.AddCallback2(callback)
	compute.ObserveCell(c1)
	compute.ObserveCell(c2)
	return &compute
}

type SpreadsheetCanceler struct {
}

// Cancel removes the callback.
func (sc SpreadsheetCanceler) Cancel() {

}

// SpreadsheetInputCell has a changeable value, changing the value triggers updates to
// other cells.
type SpreadsheetCell struct {
	data       int
	observedBy *SpreadsheetCell

	callback1 func(int) int
	callback2 func(int, int) int
	observing [2]Cell
}

func (sc *SpreadsheetCell) RegisterComputeCell(computeCell *SpreadsheetCell) {
	sc.observedBy = computeCell
	sc.recalculateAll()
}

// SetValue sets the value of the cell.
func (sc *SpreadsheetCell) SetValue(data int) {
	fmt.Printf("Setting %d\n", data)
	sc.data = data
	sc.recalculateAll()
}

func (sc *SpreadsheetCell) Value() int {
	return sc.data
}

func (sc *SpreadsheetCell) recalculateAll() {
	if sc.observedBy != nil {
		fmt.Println("Recalculating...")
		sc.observedBy.recalculate()
	}
}

func (sc *SpreadsheetCell) ObserveCell(cell Cell) {
	if sc.observing[0] == nil {
		sc.observing[0] = cell
	} else {
		sc.observing[1] = cell
	}
	cell.(*SpreadsheetCell).RegisterComputeCell(sc)
}

func (sc *SpreadsheetCell) recalculate() {
	fmt.Printf("recalculate() %v\n", sc.observing)
	switch {
	case sc.callback2 != nil && sc.observing[1] != nil:
		fmt.Println("  -> Executing callback2")
		value := sc.callback2(sc.observing[0].Value(), sc.observing[1].Value())
		sc.SetValue(value)
	case sc.callback1 != nil && sc.observing[0] != nil:
		fmt.Println("  -> Executing callback1")
		value := sc.callback1(sc.observing[0].Value())
		sc.SetValue(value)
	default:
		fmt.Println("  -> sc.observing is nil")
	}
}

// AddCallback adds a single argument callback which will be called when the value changes.
// It returns a Canceler which can be used to remove the callback.
func (sc *SpreadsheetCell) AddCallback(callback func(int) int) Canceler {
	sc.callback1 = callback
	return new(SpreadsheetCanceler)
}

// AddCallback adds a two argument callback which will be called when the value changes.
// It returns a Canceler which can be used to remove the callback.
func (sc *SpreadsheetCell) AddCallback2(callback2 func(int, int) int) Canceler {
	sc.callback2 = callback2
	return new(SpreadsheetCanceler)
}
