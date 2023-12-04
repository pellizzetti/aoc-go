package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"aocgen/pkg/aoc"
	"aocgen/pkg/gen"
	"aocgen/pkg/years"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var year, day int
var cpuProfile, memProfile bool
var cpuProfileRate int

var onStopProfiling func()
var profilingOnce sync.Once

const cpuProfileFile = "cpuprofile.prof"
const memProfileFile = "memprofile.prof"

var benchCmd = &cobra.Command{
	Use:   "bench",
	Short: "Run benchmarks for a given puzzle or whole year",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if year == 0 {
			year = time.Now().Year()
		}

		benchArgRegex := fmt.Sprintf("^Benchmark%d", year)
		if day > 0 {
			benchArgRegex += fmt.Sprintf("Day%s", gen.FormatDay(day))
		}

		cmdArgs := fmt.Sprintf("go test -benchmem -run=^$ -bench %s aocgen/pkg/year%d", benchArgRegex, year)
		c := exec.Command("bash", "-c", cmdArgs)
		out, err := c.Output()
		if err != nil {
			logrus.Error(err)
		}
		fmt.Println(string(out))
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build generated code",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		years.RegisterYears()

		gen.InitializeYearsPackages()

		years := aoc.Years()
		for _, y := range years {
			gen.InitializePackage(y)
			gen.NewBenchmarks(y)
		}
	},
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a puzzle",
	Long:  "Generate a puzzle from year and day inputs",
	Args: func(cmd *cobra.Command, args []string) error {
		if year <= 0 {
			year = time.Now().Year()
		}
		if day <= 0 {
			if time.Now().Month() == 12 {
				day = time.Now().Day()
			} else if time.Now().Month() == 11 && time.Now().Day() == 30 {
				day = 1
			} else {
				return errors.New("invalid day")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		gen.InitializePackage(year)
		gen.NewInputFile(year, day)
		gen.NewPuzzleFile(year, day)
		gen.InitializePackage(year)
		gen.InitializeYearsPackages()
		gen.NewBenchmarks(year)
	},
}

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Display puzzle input for a given puzzle",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(string(gen.WebInput(year, day)))
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all years or list all puzzles in a year",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		years.RegisterYears()

		if year != 0 {
			puzzles := aoc.Puzzles(year)
			keys := make([]int, 0)
			keysStrings := make([]string, 0)

			for k := range puzzles {
				keys = append(keys, k)
			}
			sort.Ints(keys)
			for k := range keys {
				keysStrings = append(keysStrings, strconv.Itoa(keys[k]))
			}

			fmt.Printf("%d puzzles completed or in progress:\n", len(keys))
			fmt.Println(strings.Join(keysStrings, ", "))
			return
		}

		years := aoc.Years()
		var yearsStrings []string
		for y := range years {
			yearsStrings = append(yearsStrings, strconv.Itoa(years[y]))
		}

		fmt.Printf("%d years completed or in progress:\n", len(years))
		fmt.Println(strings.Join(yearsStrings, ", "))
	},
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a puzzle and its input",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if year <= 0 {
			return errors.New("invalid year")
		}
		if day <= 0 {
			return errors.New("invalid day")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		gen.RemovePuzzle(year, day)
		gen.RemovePuzzleInput(year, day)
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a puzzle",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if year <= 0 {
			logrus.Fatal("invalid year")
		}

		years.RegisterYears()

		if day > 0 {
			runDay(year, day)
			return
		}

		runYear(year)
	},
}

func Execute() {
	rootCmd.PersistentFlags().IntVarP(&year, "year", "y", 0, "year input")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", 0, "day input")
	rootCmd.PersistentFlags().BoolVar(&cpuProfile, "cpu-profile", false, fmt.Sprintf("write cpu profile to file (%s)", cpuProfileFile))
	rootCmd.PersistentFlags().BoolVar(&memProfile, "mem-profile", false, fmt.Sprintf("write memory profile to file (%s)", memProfileFile))
	rootCmd.PersistentFlags().IntVar(&cpuProfileRate, "cpu-profile-rate", 1000000, "sets the CPU profiling rate to hz samples per second")

	rootCmd.AddCommand(benchCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(inputCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(runCmd)

	defer stopProfiling()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "AOC is a tool to support completing Advent of Code puzzles",
	Long:  "AOC supports generating puzzle data, including inputs directly from the website, and benchmarking answers",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		onStopProfiling = initProfiling(cpuProfile, memProfile)
	},
}

func runYear(year int) {
	puzzles := aoc.Puzzles(year)
	for i := 1; i <= len(puzzles); i++ {
		runDay(year, i)
	}
}

func runDay(year, day int) {
	aoc.Run(year, day, aoc.NewPuzzle(year, day), aoc.Input(year, day))
}

func initProfiling(cpuProfile, memProfile bool) (stop func()) {
	var doOnStop []func()
	stop = func() {
		for _, d := range doOnStop {
			if d != nil {
				d()
			}
		}
	}

	if cpuProfile {
		logrus.Info("CPU profile enabled")
		f, err := os.Create(cpuProfileFile)
		if err != nil {
			logrus.Error("Could not create CPU profile file")
			return stop
		}
		runtime.SetCPUProfileRate(cpuProfileRate)
		err = pprof.StartCPUProfile(f)
		if err != nil {
			logrus.Error("Could not start CPU profiling")
			return stop
		}
		// Add function to stop cpu profiling to doOnStop list
		doOnStop = append(doOnStop, func() {
			pprof.StopCPUProfile()
			_ = f.Close()
			logrus.Info("CPU profile stopped")
		})
	}

	if memProfile {
		logrus.Info("Memory profile enabled")
		f, err := os.Create(memProfileFile)
		if err != nil {
			logrus.Error("Could not create memory profile file")
			return stop
		}
		// Add function to stop memory profiling to doOnStop list
		doOnStop = append(doOnStop, func() {
			_ = pprof.WriteHeapProfile(f)
			_ = f.Close()
			logrus.Info("Memory profile stopped")
		})
	}

	return stop
}

func stopProfiling() {
	if onStopProfiling != nil {
		profilingOnce.Do(onStopProfiling)
	}
}

func main() {
	Execute()
}
