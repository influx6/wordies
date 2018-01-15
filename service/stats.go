package service

import (
	"sync"
)

// FreshStat embodies generated stats from the
// word and letter counters.
type FreshStat struct {
	TotalLetters int
	TotalWords   int
	Words        map[int][]string
	Letters      map[int][]string
}

// StatMetric presents a simple interface to notify of a finished event.
type StatMetric interface {
	Update(FreshStat) error
}

// StatMetrics defines a slice of StatMetrics instances.
type StatMetrics []StatMetric

// Notify sends calls to all Notify instances within it'self.
func (n StatMetrics) Update(fs FreshStat) error {
	for _, elem := range n {
		if err := elem.Update(fs); err != nil {
			return err
		}
	}
	return nil
}

// Top5WordLetterStat implements the StatMetric interface, it
// generates the top 5 letters and words metric from the
// instance of FreshStat it receives when runned.
// It's stats method always returns it last successfully generated
// stats.
type Top5WordLetterStat struct {
	mu   sync.RWMutex
	last Top5Stat
}

// Update recalculates the top5 letters and words
// from both the received argument.
func (sn *Top5WordLetterStat) Update(fr FreshStat) error {
	var stat Top5Stat
	stat.Count = fr.TotalWords

	if ltfreq := sortFrequencies(fr.Letters); len(ltfreq) != 0 {
		for _, freq := range ltfreq {
			if len(stat.Letters) >= 5 {
				break
			}

			stat.Letters = append(stat.Letters, fr.Letters[freq]...)
		}

		if len(stat.Letters) > 5 {
			stat.Letters = stat.Letters[:5]
		}
	}

	if wdfreq := sortFrequencies(fr.Words); len(wdfreq) != 0 {
		for _, freq := range wdfreq {
			if len(stat.Words) >= 5 {
				break
			}

			stat.Words = append(stat.Words, fr.Words[freq]...)
		}

		if len(stat.Words) > 5 {
			stat.Words = stat.Words[:5]
		}
	}

	sn.mu.Lock()
	sn.last = stat
	sn.mu.Unlock()
	return nil
}

// Top5Stat embodies data stored by Top5WordLetterStat to
// contain metrics on total words, top 5 words and letters.
type Top5Stat struct {
	Count   int      `json:"count"`
	Letters []string `json:"top_5_letters"`
	Words   []string `json:"top_5_words"`
}

// Stat returns the latest generated stat since last call of it's
// .Update method.
func (sn *Top5WordLetterStat) Stat() Top5Stat {
	sn.mu.RLock()
	defer sn.mu.RUnlock()
	return sn.last
}
