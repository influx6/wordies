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
	Update(FreshStat)
}

// StatMetrics defines a slice of StatMetrics instances.
type StatMetrics []StatMetric

// Notify sends calls to all Notify instances within it'self.
func (n StatMetrics) Update(fs FreshStat) {
	for _, elem := range n {
		go elem.Update(fs)
	}
}

// Top5WordLetterStat implements the Notify interface, which
// it uses to initiates the underline process of generating
// new stats about the top 5 letters and words.
type Top5WordLetterStat struct {
	mu   sync.RWMutex
	last Top5Stat
}

// Notify recalculates the latests top5 letters and words
// from both the LetterCounter and the WordCounter.
func (sn *Top5WordLetterStat) Update(fr FreshStat) {
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
}

// Top5Stat embodies data collected over top5
type Top5Stat struct {
	Count   int      `json:"count"`
	Letters []string `json:"top_5_letters"`
	Words   []string `json:"top_5_words"`
}

// Stat returns the latests generated since last call of it's
// .Notify method.
// We do this to minimize expensive ops of calculating new
// top5 markers everytime by letter the service notify
// need of update.
func (sn *Top5WordLetterStat) Stat() Top5Stat {
	sn.mu.RLock()
	defer sn.mu.RUnlock()
	return sn.last
}
