package sys

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/wailsapp/wails"
)

// Stats .
type Stats struct {
	log *wails.CustomLogger
}

// CPUUsage .
type CPUUsage struct {
	Average int `json:"avg"`
}

// WailsInit .
func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")

	go func() {

		for {

			runtime.Events.Emit("usage", s.GetCPUUsage())
			time.Sleep(1 * time.Second)
		}
	}()

	return nil
}

// GetCPUUsage .
func (s *Stats) GetCPUUsage() *CPUUsage {
	percent, err := cpu.Percent(1*time.Second, false)
	initialStat, err1 := os.Stat("K://WordPressLogAnalysisTool//local_agent//backend//Testings//cd-cpustats//test.txt")

	if err1 != nil {
		s.log.Errorf("undefined filepath!: %s", err.Error())
		s.log.Errorf("Initial file stat: %s", initialStat.ModTime())
	}
	for {
		stat, err := os.Stat("K://WordPressLogAnalysisTool//local_agent//backend//Testings//cd-cpustats//test.txt")
		//s.log.Errorf("Updated file stat: %s", stat.ModTime())
		if err != nil {

			s.log.Errorf("reading file: %s", err.Error())
		}

		if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
			s.log.Errorf("Updated file stat: %s", stat.ModTime())
			break
		}
	}

	if err != nil {

		s.log.Errorf("unable to get cpu stats: %s", err.Error())
		return nil
	} else {
		a := int(math.Round(percent[0]))
		fmt.Print(a)
	}

	return &CPUUsage{
		// Average: int(math.Round(percent[0])),
		Average: 555,
	}
}
