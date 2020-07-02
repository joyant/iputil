package iputil

import (
    "reflect"
    "testing"
)

func TestStr2int(t *testing.T)  {
    got := str2int("123")
    expected := 123
    if expected != got {
        t.Errorf("1 expected %d got %d", expected, got)
    }
    got2 := str2int("123 ")
    expected2 := -1
    if expected2 != got2 {
        t.Errorf("2 expected %d got %d", expected2, got2)
    }
    got3 := str2int("abc")
    expected3 := -1
    if expected3 != got3 {
        t.Errorf("3 expected %d got %d", expected3, got3)
    }
    got4 := str2int("123.4")
    expected4 := -1
    if expected4 != got4 {
        t.Errorf("4 expected %d got %d", expected4, got4)
    }
    got5 := str2int("012345")
    expected5 := 12345
    if expected5 != got5 {
        t.Errorf("5 expected %d got %d", expected5, got5)
    }
}

func TestIsIPV4(t *testing.T) {
    got := IsIPV4("255.255.255.1")
    expected := true
    if got != expected {
        t.Errorf("1 expected %t got %t", expected, got)
    }
    got2 := IsIPV4("255.255.255.")
    expected2 := false
    if got2 != expected2 {
       t.Errorf("2 expected %t got %t", expected2, got2)
    }
    got3 := IsIPV4(".255.255.255")
    expected3 := false
    if got3 != expected3 {
        t.Errorf("3 expected %t got %t", expected3, got3)
    }
    got4 := IsIPV4("255.255.255.1 ")
    expected4 := false
    if got4 != expected4 {
        t.Errorf("4 expected %t got %t", expected4, got4)
    }
    got5 := IsIPV4("1.1.1.0")
    expected5 := true
    if got5 != expected5 {
        t.Errorf("5 expected %t got %t", expected5, got5)
    }
}

func TestIsIPV6(t *testing.T) {
    // expected true
    got := IsIPV6("::1")
    if got != true {
        t.Errorf("0 expected true got %t", got)
    }
    got2 := IsIPV6("a:b:c:d::e")
    if got2 != true {
        t.Errorf("2 expected true got %t", got2)
    }
    got3 := IsIPV6("1::")
    if got3 != true {
        t.Errorf("3 expected true got %t", got3)
    }
    got4 := IsIPV6("::FFFF")
    if got4 != true {
        t.Errorf("4 expected true got %t", got4)
    }
    got5 := IsIPV6("ffff::ABCD")
    if got5 != true {
        t.Errorf("5 expected true got %t", got5)
    }
    got6 := IsIPV6("ABCD:EF01:2345:6789:ABCD:EF01:2345:6789")
    if got6 != true {
        t.Errorf("6 expected true got %t", got6)
    }
    got7 := IsIPV6("FEDC:BA98:7654:3210:FEDC:BA98:7654:3210")
    if got7 != true {
        t.Errorf("7 expected true got %t", got7)
    }
    got8 := IsIPV6("1080:0:0:0:8:800:200C:417A")
    if got8 != true {
        t.Errorf("8 expected true got %t", got8)
    }
    got9 := IsIPV6("1080:0:0:0:8:800:200C:417A")
    if got9 != true {
        t.Errorf("9 expected true got %t", got9)
    }
    got10 := IsIPV6("1080::8:800:200C:417A")
    if got10 != true {
        t.Errorf("10 expected true got %t", got10)
    }
    got11 := IsIPV6("FF01:0:0:0:0:0:0:101")
    if got11 != true {
        t.Errorf("11 expected true got %t", got11)
    }
    got12 := IsIPV6("FF01::101")
    if got12 != true {
        t.Errorf("12 expected true got %t", got12)
    }
    got13 := IsIPV6("::")
    if got13 != true {
        t.Errorf("13 expected true got %t", got13)
    }

    // expected false

    got14 := IsIPV6(":::")
    if got14 != false {
        t.Errorf("14 expected false got %t", got14)
    }
    got15 := IsIPV6("abcde::1")
    if got15 != false {
        t.Errorf("15 expected false got %t", got15)
    }
    got16 := IsIPV6("FEDC:BA98:7654:3210:FEDC:BA98:7654:3210:1")
    if got16 != false {
        t.Errorf("16 expected false got %t", got16)
    }
    got17 := IsIPV6("FEDC:BA98:7654:3210:FEDC:BA98:7654:3210:18:9:0:8:23:23:")
    if got17 != false {
        t.Errorf("17 expected false got %t", got17)
    }
    got18 := IsIPV6("192.168.0.1")
    if got18 != false {
        t.Errorf("18 expected false got %t", got18)
    }
    got19 := IsIPV6("1008:124::1::2")
    if got19 != false {
        t.Errorf("19 expected false got %t", got19)
    }
}

func TestString2IPV4(t *testing.T) {
    got, ok := String2IPV4("127.0.0.1")
    expected := IPV4{127, 0, 0, 1}
    if !ok {
        t.Errorf("1 expected true got %t", ok)
    }
    if !reflect.DeepEqual(got, expected) {
        t.Errorf("1 expected %v got %v", expected, got)
    }

    got2, ok2 := String2IPV4("192.168.100.1")
    expected2 := IPV4{192, 168, 100, 1}
    if !ok2 {
        t.Errorf("2 expected true got %t", ok2)
    }
    if !reflect.DeepEqual(got2, expected2) {
        t.Errorf("2 expected %v got %v", expected2, got2)
    }

    got3, ok3 := String2IPV4("192.168.100.1.2")
    expected3 := IPV4{192, 168, 100, 1}
    if ok3 {
        t.Errorf("3 expected true got %t", ok3)
    }
    if reflect.DeepEqual(got3, expected3) {
        t.Errorf("3 expected %v got %v", expected3, got3)
    }
}

func TestString2IPV6(t *testing.T) {
    got, ok := String2IPV6("::")
    expected := IPV6{}
    if !ok {
        t.Errorf("1 expected true got %v", ok)
    }
    if !reflect.DeepEqual(got, expected) {
        t.Errorf("1 expected %v got %v", expected, got)
    }

    got2, ok2 := String2IPV6("::1")
    expected2 := IPV6{0, 0, 0, 0, 0, 0, 0, 1}
    if !ok2 {
        t.Errorf("2 expected true got %v", ok2)
    }
    if !reflect.DeepEqual(got2, expected2) {
        t.Errorf("2 expected %v got %v", expected2, got2)
    }

    got3, ok3 := String2IPV6("1::")
    expected3 := IPV6{1, 0, 0, 0, 0, 0, 0, 0}
    if !ok2 {
        t.Errorf("3 expected true got %v", ok3)
    }
    if !reflect.DeepEqual(got3, expected3) {
        t.Errorf("3 expected %v got %v", expected3, got3)
    }

    got4, ok4 := String2IPV6("1::1")
    expected4 := IPV6{1, 0, 0, 0, 0, 0, 0, 1}
    if !ok4 {
        t.Errorf("4 expected true got %v", ok4)
    }
    if !reflect.DeepEqual(got4, expected4) {
        t.Errorf("4 expected %v got %v", expected4, got4)
    }

    got5, ok5 := String2IPV6("1:2:3:4:5:6:7:8")
    expected5 := IPV6{1, 2, 3, 4, 5, 6, 7, 8}
    if !ok5 {
        t.Errorf("5 expected true got %v", ok5)
    }
    if !reflect.DeepEqual(got5, expected5) {
        t.Errorf("5 expected %v got %v", expected5, got5)
    }

    got6, ok6 := String2IPV6("1:2::7:8")
    expected6 := IPV6{1, 2, 0, 0, 0, 0, 7, 8}
    if !ok6 {
        t.Errorf("6 expected true got %v", ok6)
    }
    if !reflect.DeepEqual(got6, expected6) {
        t.Errorf("6 expected %v got %v", expected6, got6)
    }

    got7, ok7 := String2IPV6("1:2::")
    expected7 := IPV6{1, 2, 0, 0, 0, 0, 0, 0}
    if !ok7 {
        t.Errorf("7 expected true got %v", ok7)
    }
    if !reflect.DeepEqual(got7, expected7) {
        t.Errorf("7 expected %v got %v", expected7, got7)
    }

    got8, ok8 := String2IPV6("::1:2:3:4:5:6")
    expected8 := IPV6{0, 0, 1, 2, 3, 4, 5, 6}
    if !ok8 {
        t.Errorf("8 expected true got %v", ok8)
    }
    if !reflect.DeepEqual(got8, expected8) {
        t.Errorf("8 expected %v got %v", expected8, got8)
    }

    got9, ok9 := String2IPV6("87:4B:2B:34::1")
    expected9 := IPV6{135, 75, 43, 52, 0, 0, 0, 1}
    if !ok9 {
        t.Errorf("9 expected true got %v", ok9)
    }
    if !reflect.DeepEqual(got9, expected9) {
        t.Errorf("9 expected %v got %v", expected9, got9)
    }
}

func TestIsIP(t *testing.T) {
    if !IsIP("87:4B:2B:34::1") {
        t.Errorf("expected true got false")
    }
    if !IsIP("::1") {
        t.Errorf("expected true got false")
    }
    if !IsIP("192.168.0.1") {
        t.Errorf("expected true got false")
    }
    if IsIP(":::") {
        t.Errorf("expected false got true")
    }
}

func TestIPV42IPV6(t *testing.T) {
    got := IPV42IPV6(IPV4{192, 168, 0, 1})
    expected := "::c0:a8:00:01"
    if got != expected {
        t.Errorf("expected %v got %v", expected, got)
    }
}

func TestStringIPV42IPV6(t *testing.T) {
    if v6, ok := StringIPV42IPV6("192.168.0.1"); !ok {
        t.Errorf("expected true got false")
    } else if v6 != "::c0:a8:00:01" {
        t.Errorf("expected %s got %s", "::c0:a8:00:01", v6)
    }
}

func TestIpv6FragmentNumeric(t *testing.T)  {
    n, ok := ipv6FragmentNumeric("ffff")
    if !ok {
        t.Errorf("expected true got %t", ok)
    }
    expected := uint16(0xffff)
    if n != expected {
        t.Errorf("expected %d got %d", expected, n)
    }
    n, ok = ipv6FragmentNumeric("fffff")
    if ok {
        t.Errorf("expected false got %t", ok)
    }
}
