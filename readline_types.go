package wall_street

// enums.go
type UndoEnum int

const (
	undoDelete UndoEnum = iota
	undoInsert UndoEnum = iota
	undoBegin  UndoEnum = iota
	undoEnd    UndoEnum = iota
)

type UndoList struct {
	next  *UndoList
	start int // where the change took place
	end   int
	text  string // text to insert, if undoing a delete
	what  UndoEnum
}

// the current undo list for the RL_LINE_BUFFER
var rl_undo_list *UndoList

// data structure for mapping textual names to code addresses
type FunMap struct {
	name          string
	rlCommandFunc func() // rl_command_func_t
}
