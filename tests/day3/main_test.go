package day1 // if you are doing black-box testing; use 'package day1' for white-box testing

import (
	"reflect"
	"testing"

	"github.com/burtenshaw/advent/src/day3"
)

func TestGetSchematicParts(t *testing.T) {
    tests := []struct {
        name  string
        schematic  day3.Schematic
        wantNumbers  []day3.SchematicNumber
        wantGearCandidates day3.GearCandidates
    }{
        {
            name: "Test 1: Basic schematic",
            schematic: day3.Schematic{Contents: "123*456\n789*0"},
            wantNumbers: []day3.SchematicNumber{
                day3.NewSchematicNumber(123),
                day3.NewSchematicNumber(456),
                day3.NewSchematicNumber(789),
                day3.NewSchematicNumber(0),
            },
            wantGearCandidates: day3.GearCandidates{
                day3.Coordinates{X: 3, Y: 0}: []int{123},
                day3.Coordinates{X: 7, Y: 1}: []int{789},
            },
        },
        // Add more test cases as needed.
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotNumbers, gotGearCandidates := tt.schematic.GetSchematicParts()
            if !reflect.DeepEqual(gotNumbers, tt.wantNumbers) {
                t.Errorf("GetSchematicParts() gotNumbers = %v, want %v", gotNumbers, tt.wantNumbers)
            }
            if !reflect.DeepEqual(gotGearCandidates, tt.wantGearCandidates) {
                t.Errorf("GetSchematicParts() gotGearCandidates = %v, want %v", gotGearCandidates, tt.wantGearCandidates)
            }
        })
    }
}