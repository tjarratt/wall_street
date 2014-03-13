package wall_street

import (
	"container/ring"
	"errors"
	"time"
)

const historyMaxEntries = 512 // should be enough for anyone

type HistEntry struct {
	Line      string
	Timestamp int64
	Data      interface{} // histdata_t #ifdef to void* or char*
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
	theHistory     *ring.Ring
	currentHistory **ring.Ring
	oldestHistory  **ring.Ring
)

// TODO: replace me with a constructor that wraps these package vars
func init() {
	ResetHistory()
}

/* Return the current history array.  The caller has to be careful, since this
   is the actual array of data, and could be bashed or made corrupt easily.
   The array is terminated with a NULL pointer. */
func HistoryList() (history []*HistEntry) {
	theHistory.Do(func(h interface{}) {
		history = append(history, h.(*HistEntry))
	})

	return
}

// addition: just resets the world for test purposes
func ResetHistory() {
	theHistory = ring.New(0)
	oldestHistory = &theHistory
	return
}

func CurrentHistory() *HistEntry {
	if theHistory.Len() == 0 {
		return nil
	}

	entry, ok := (*currentHistory).Value.(*HistEntry)
	if !ok {
		return nil
	} else {
		return entry
	}
}

/* Back up history_offset to the previous history entry, and return
   a pointer to that entry.  If there is no previous entry then return
   a NULL pointer. */
func PreviousHistory() *HistEntry {
	prev := (*currentHistory).Prev()
	currentHistory = &prev
	return (*currentHistory).Value.(*HistEntry)
}

/* Move history_offset forward to the next history entry, and return
   a pointer to that entry.  If there is no next entry then return a
   nil pointer. */
func NextHistory() *HistEntry {
	next := (*currentHistory).Next()
	currentHistory = &next
	return (*currentHistory).Value.(*HistEntry)
}

/* Return the history entry which is logically at OFFSET in the history array.
   OFFSET is relative to history_base. */
func HistoryGet(offset int) *HistEntry {
	return currentHistory.Move(offset).Value.(*HistEntry)
}

func HistoryGetTime(history *HistEntry) time.Time {
	return time.Unix(history.Timestamp, 0)
}

func AddHistory(input string) {
	// FIXME: should remove the oldest entry
	if theHistory.Len() == historyMaxEntries {
		theHistory.Unlink(1)
	}

	newEntry := new(HistEntry)
	newEntry.Line = input
	newEntry.Timestamp = time.Now().Unix()

	if theHistory.Len() == 0 {
		theHistory = ring.New(1)
		theHistory.Value = newEntry
		currentHistory = &theHistory
	} else {
		newLink := ring.New(1)
		newLink.Value = newEntry
		(*currentHistory).Link(newLink)
		currentHistory = &newLink
	}
}

// NOT IMPLEMENTED
// func AddHistoryTime(input string)
// func FreeHistoryEntry(history HistEntry)
// func CopyHistoryEntry(history HistEntry)

/* Make the history entry at WHICH have LINE and DATA.*/
// nb: `which` offset is now relative to current history **BREAKING CHANGE**
func ReplaceHistoryEntry(which int, line string, data interface{}) (err error) {
	history := currentHistory.Move(which).Value.(*HistEntry)
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
// nb: `which` offset is now relative to current history **BREAKING CHANGE**
func ReplaceHistoryData(which int, oldData, newData interface{}) (err error) {
	if which < -2 || which >= theHistory.Len() {
		err = errors.New("invalid history offset")
		return
	}

	if which >= 0 {
		history := theHistory.Move(which).Value.(*HistEntry)
		if history.Data == oldData {
			history.Data = newData
		}

	} else if which == -1 {
		theHistory.Do(func(value interface{}) {
			entry, ok := value.(*HistEntry)
			if !ok {
				return
			}

			if entry.Data == oldData {
				entry.Data = newData
			}
		})
	} else if which == -2 {
		for i := theHistory.Len(); i >= 0; i++ {
			entry := theHistory.Move(-1).Value.(*HistEntry)
			if entry.Data == oldData {
				entry.Data = newData
				break
			}
		}
	}

	return
}

// Remove history element WHICH from the history
func RemoveHistory(which int) {
	toBeRemoved := currentHistory.Move(which)
	prev := toBeRemoved.Prev()
	next := toBeRemoved.Next()

	prev.Link(next)
	currentHistory = &prev
	return
}
