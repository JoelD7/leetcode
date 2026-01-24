package time_based_key_value_store

type TimeMap struct {
	table map[string]*Entry
}

type Entry struct {
	timestamp         []int          //timestamps set on the map
	valuesByTimestamp map[int]string //values by timestamp
}

func NewEntry(t int, val string) *Entry {
	entry := &Entry{
		timestamp:         make([]int, 0),
		valuesByTimestamp: make(map[int]string),
	}

	entry.timestamp = append(entry.timestamp, t)
	entry.valuesByTimestamp[t] = val

	return entry
}

func Constructor() TimeMap {
	return TimeMap{
		table: make(map[string]*Entry),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	curTable, ok := this.table[key]
	if !ok {
		this.table[key] = NewEntry(timestamp, value)
		return
	}

	curTable.timestamp = append(curTable.timestamp, timestamp)
	curTable.valuesByTimestamp[timestamp] = value
}

func (this *TimeMap) Get(key string, timestamp int) string {
	curTable, ok := this.table[key]
	if !ok {
		return ""
	}

	val, ok := curTable.valuesByTimestamp[timestamp]
	if ok {
		return val
	}

	lastAddedTimestamp := this.searchNextLower(timestamp, curTable)
	return curTable.valuesByTimestamp[lastAddedTimestamp]
}

func (this *TimeMap) searchNextLower(timestamp int, entry *Entry) int {
	for i := len(entry.timestamp) - 1; i >= 0; i-- {
		if entry.timestamp[i] < timestamp {
			return entry.timestamp[i]
		}
	}

	return -1
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
