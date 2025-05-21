package monitor

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type GPUStats struct {
	Name        string
	Temperature int
	MemUsed     float64
	MemTotal    float64
	Utilization int
	ClockCore   int
	ClockMemory int
	ClockMax    int
	FanSpeed    int
	PowerDraw   float64
	PowerLimit  float64
}

func GetGPUInfo() (GPUStats, error) {
	out, err := exec.Command("nvidia-smi",
		"--query-gpu=name,temperature.gpu,utilization.gpu,memory.used,memory.total,"+
			"clocks.gr,clocks.mem,clocks.max.gr,fan.speed,power.draw,enforced.power.limit",
		"--format=csv,noheader,nounits").Output()
	if err != nil {
		return GPUStats{}, err
	}

	parts := strings.Split(strings.TrimSpace(string(out)), ", ")
	if len(parts) < 5 {
		return GPUStats{}, fmt.Errorf("dados incompletos do nvidia-smi")
	}

	temp, _ := strconv.Atoi(parts[1])
	util, _ := strconv.Atoi(parts[2])
	memUsed, _ := strconv.ParseFloat(parts[3], 64)
	memTotal, _ := strconv.ParseFloat(parts[4], 64)
	clockCore, _ := strconv.Atoi(parts[5])
	clockMem, _ := strconv.Atoi(parts[6])
	clockMax, _ := strconv.Atoi(parts[7])
	fanSpeed, _ := strconv.Atoi(parts[8])
	powerDraw, _ := strconv.ParseFloat(parts[9], 64)
	powerLimit, _ := strconv.ParseFloat(parts[10], 64)

	return GPUStats{
		Name:        parts[0],
		Temperature: temp,
		MemUsed:     memUsed,
		MemTotal:    memTotal,
		Utilization: util,
		ClockCore:   clockCore,
		ClockMemory: clockMem,
		ClockMax:    clockMax,
		FanSpeed:    fanSpeed,
		PowerDraw:   powerDraw,
		PowerLimit:  powerLimit,
	}, nil
}
