package wall_street

import (
	"errors"
	"time"
)

const historyMaxEntries = 512 // should be enough for anyone

type HistEntry struct {
	Line string
	Timestamp int64
	Data interface{} // histdata_t #ifdef to void* or char*
}

/*
  ** UNIMPLEMENTED **
*/
// get, set history state
// using_history
// history_total_bytes
// where_history
// history_set_pos

/*
  ** Implementation Ahoy **
*/
// global state is fun, no?
var (
	the_history []*HistEntry
	history_offset int
)

/* Return the current history array.  The caller has to be careful, since this
   is the actual array of data, and could be bashed or made corrupt easily.
   The array is terminated with a NULL pointer. */
func HistoryList() (history []*HistEntry) {
	return the_history
}

// FIXME: should probably return nil or an error if offset is bad
func CurrentHistory() *HistEntry {
	return the_history[history_offset]
}

/* Back up history_offset to the previous history entry, and return
   a pointer to that entry.  If there is no previous entry then return
   a NULL pointer. */
func PreviousHistory() *HistEntry {
	history_offset--  // FIXME: bounds checking or ring buffer
	return the_history[history_offset]
}

/* Move history_offset forward to the next history entry, and return
   a pointer to that entry.  If there is no next entry then return a
   NULL pointer. */
func NextHistory() *HistEntry {
	history_offset++ // FIXME: bounds checking or ring buffer
	return the_history[history_offset]
}

/* Return the history entry which is logically at OFFSET in the history array.
   OFFSET is relative to history_base. */
func HistoryGet(offset int) *HistEntry {
	// FIXME: wtf history_base ???
	return the_history[offset] // FIXME: bounds checking or ring buffer
}

func HistoryGetTime(history *HistEntry) time.Time {
	return time.Unix(history.Timestamp, 0)
}

func AddHistory(input string) {
	if len(the_history) == historyMaxEntries {
		the_history = the_history[1:]
	}

	newEntry := new(HistEntry)
	newEntry.Line = input
	newEntry.Timestamp = time.Now().Unix()
	the_history = append(the_history, newEntry)
}

// NOT IMPLEMENTED
// func AddHistoryTime(input string)
// func FreeHistoryEntry(history HistEntry)
// func CopyHistoryEntry(history HistEntry)

/* Make the history entry at WHICH have LINE and DATA.*/
func ReplaceHistoryEntry(which int, line string, data interface{}) (err error) {
	if which < 0 || which >= len(the_history) {
		err = errors.New("invalid history offset")
		return
	}

	history := the_history[which]
	history.Line = line
	history.Data = data
	return
}

// FIXME: avoid using either of these ridiculous functions if possible

/* Replace the DATA in the specified history entries, replacing OLD with
   NEW.  WHICH says which one(s) to replace:  WHICH == -1 means to replace
   all of the history entries where entry->data == OLD; WHICH == -2 means
   to replace the `newest' history entry where entry->data == OLD; and
   WHICH >= 0 means to replace that particular history entry's data, as
   long as it matches OLD. */
func ReplaceHistoryData(which int, oldData, newData interface{}) (err error) {
	if which < -2 || which >= len(the_history) {
		err = errors.New("invalid history offset")
		return
	}

	if which >= 0 {
		history := the_history[which]
		if history.Data == oldData {
			history.Data = newData
		}

	} else if which == -1 {
		for _, entry := range the_history {
			if entry.Data == oldData {
				entry.Data = newData
			}
		}
	} else if which == -2 {
		for i := len(the_history); i >= 0; i++ {
			if the_history[i].Data == oldData {
				the_history[i].Data = newData
				break
			}
		}
	}

	return
}
