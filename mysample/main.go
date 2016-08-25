package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"reflect"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/google/flatbuffers/mysample/MyGame/Sample"
)

func main() {
	bin := flag.String("bin", "monster.bin", "flatbuffer monster binary file")
	check(*bin)
}

func check(bin string) error {
	buf, err := ioutil.ReadFile(bin)
	if err != nil {
		return err
	}

	monster := &Sample.Monster{}
	flatbuffers.GetRootAs(buf, 0, monster)

	if got := monster.Hp(); 80 != got {
		return Fail("hp", 80, got)
	}

	// default
	if got := monster.Mana(); 150 != got {
		return Fail("mana", 150, got)
	}

	if got := monster.Name(); !bytes.Equal([]byte("MyMonster"), got) {
		return Fail("name", "MyMonster", got)
	}

	// initialize a Vec3 from Pos()
	vec := new(Sample.Vec3)
	vec = monster.Pos(vec)
	if vec == nil {
		return errors.New("vec3 initialization failed")
	}

	// check that new allocs equal given ones:
	vec2 := monster.Pos(nil)
	if !reflect.DeepEqual(vec, vec2) {
		return errors.New("fresh allocation failed")
	}

	// verify the properties of the Vec3
	if got := vec.X(); float32(1.0) != got {
		return Fail("Pos.X", float32(1.0), got)
	}

	if got := vec.Y(); float32(2.0) != got {
		return Fail("Pos.Y", float32(2.0), got)
	}

	if got := vec.Z(); float32(3.0) != got {
		return Fail("Pos.Z", float32(3.0), got)
	}

	inventorySlice := monster.InventoryBytes()
	if len(inventorySlice) != monster.InventoryLength() {
		return Fail("len(monster.InventoryBytes) != monster.InventoryLength", len(inventorySlice), monster.InventoryLength())
	}

	if got := monster.InventoryLength(); 5 != got {
		return Fail("monster.InventoryLength", 5, got)
	}

	invsum := 0
	l := monster.InventoryLength()
	for i := 0; i < l; i++ {
		v := monster.Inventory(i)
		if v != inventorySlice[i] {
			return Fail("monster inventory slice[i] != Inventory(i)", v, inventorySlice[i])
		}
		invsum += int(v)
	}
	if invsum != 10 {
		return Fail("monster inventory sum", 10, invsum)
	}

	return nil
}

// Fail makes an error with message for when expectations differ from reality.
func Fail(name string, want, got interface{}) error {
	return fmt.Errorf("bad %s: want %#v got %#v", name, want, got)
}
