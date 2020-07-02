package iputil

import (
    "reflect"
    "testing"
)

func TestIpv6RangeFragment(t *testing.T)  {
    // true
    got, ok := ipv6RangeFragment("FFFF")
    if !ok {
        t.Errorf("1 expected true got %t", ok)
    }
    expected := rangeFragment{0xffff, 0xffff}
    if !reflect.DeepEqual(got, expected) {
        t.Errorf("1 expected %v got %v", expected, got)
    }

    got2, ok2 := ipv6RangeFragment("0-FFFF")
    if !ok2 {
        t.Errorf("2 expected true got %t", ok2)
    }
    expected2 := rangeFragment{0, 0xffff}
    if !reflect.DeepEqual(got2, expected2) {
        t.Errorf("2 expected %v got %v", expected2, got2)
    }

    got3, ok3 := ipv6RangeFragment("*")
    if !ok3 {
        t.Errorf("3 expected true got %t", ok3)
    }
    expected3 := rangeFragment{0, 0xffff}
    if !reflect.DeepEqual(got3, expected3) {
        t.Errorf("3 expected %v got %v", expected3, got3)
    }

    got4, ok4 := ipv6RangeFragment("1-1008")
    if !ok4 {
        t.Errorf("4 expected true got %t", ok4)
    }
    expected4 := rangeFragment{1, 4104}
    if !reflect.DeepEqual(got4, expected4) {
        t.Errorf("4 expected %v got %v", expected4, got4)
    }

    //false

    _, ok5 := ipv6RangeFragment("1-*")
    if ok5 {
        t.Errorf("5 expected false got %t", ok5)
    }
    _, ok6 := ipv6RangeFragment("1-fffff")
    if ok6 {
        t.Errorf("6 expected false got %t", ok6)
    }
    _, ok7 := ipv6RangeFragment("1~10")
    if ok7 {
        t.Errorf("7 expected false got %t", ok7)
    }
}

func TestIpv4RangeFragment(t *testing.T)  {
    // true
    got, ok := ipv4RangeFragment("192")
    expected := rangeFragment{192, 192}
    if !ok {
        t.Errorf("1 expected true got %t", ok)
    }
    if !reflect.DeepEqual(got, expected) {
        t.Errorf("1 expected %v got %v", expected, got)
    }

    got2, ok2 := ipv4RangeFragment("*")
    expected2 := rangeFragment{0, 255}
    if !ok2 {
        t.Errorf("2 expected true got %t", ok2)
    }
    if !reflect.DeepEqual(got2, expected2) {
        t.Errorf("2 expected %v got %v", expected2, got2)
    }

    got3, ok3 := ipv4RangeFragment("1-108")
    expected3 := rangeFragment{1, 108}
    if !ok3 {
        t.Errorf("3 expected true got %t", ok3)
    }
    if !reflect.DeepEqual(got3, expected3) {
        t.Errorf("3 expected %v got %v", expected3, got3)
    }

    // false

    _, ok4 := ipv4RangeFragment("1-256")
    if ok4 {
        t.Errorf("4 expected true got %t", ok4)
    }
    _, ok5 := ipv4RangeFragment("1-*")
    if ok5 {
        t.Errorf("5 expected true got %t", ok5)
    }
    _, ok6 := ipv4RangeFragment("278")
    if ok6 {
        t.Errorf("6 expected true got %t", ok6)
    }
}

func TestLikeIPV4String2RangeIp(t *testing.T)  {
    // true
    got, ok := likeIPV4String2RangeIp("192.168.0.1")
    if !ok {
       t.Errorf("1 expected true got %t", ok)
    }
    expected := rangeIP{{}, {}, {}, {}, {192, 192}, {168, 168}, {0, 0}, {1, 1}}
    if !reflect.DeepEqual(got, expected) {
       t.Errorf("1 expected %v got %v", expected, got)
    }

    got2, ok2 := likeIPV4String2RangeIp("192.168.0-10.1")
    if !ok2 {
        t.Errorf("2 expected true got %t", ok2)
    }
    expected2 := rangeIP{{}, {}, {}, {}, {192, 192}, {168, 168}, {0, 10}, {1, 1}}
    if !reflect.DeepEqual(got2, expected2) {
        t.Errorf("2 expected %v got %v", expected2, got2)
    }

    got3, ok3 := likeIPV4String2RangeIp("*.168.0-10.1")
    if !ok3 {
        t.Errorf("3 expected true got %t", ok3)
    }
    expected3 := rangeIP{{}, {}, {}, {}, {0, 255}, {168, 168}, {0, 10}, {1, 1}}
    if !reflect.DeepEqual(got3, expected3) {
        t.Errorf("3 expected %v got %v", expected3, got3)
    }

    // false

    _, ok4 := likeIPV4String2RangeIp("*.168.0-10.355")
    if ok4 {
        t.Errorf("4 expected false got %t", ok4)
    }
    _, ok5 := likeIPV4String2RangeIp("256.168.0-10.0")
    if ok5 {
        t.Errorf("5 expected false got %t", ok5)
    }
    _, ok6 := likeIPV4String2RangeIp("256.168")
    if ok6 {
        t.Errorf("6 expected false got %t", ok6)
    }
    _, ok7 := likeIPV4String2RangeIp("256.168.0.-1")
    if ok7 {
        t.Errorf("7 expected false got %t", ok7)
    }
    _, ok8 := likeIPV4String2RangeIp("256.168.0.1.0")
    if ok8 {
        t.Errorf("8 expected false got %t", ok8)
    }
}

func TestLikeIPV6String2RangeIp(t *testing.T) {
    // true
    got0, ok0 := likeIPV6String2RangeIp("*:*:*::")
    if !ok0 {
        t.Errorf("1 expected true got %t", ok0)
    }
    expected0 := rangeIP{{0, 0xffff}, {0, 0xffff}, {0, 0xffff}}
    if !reflect.DeepEqual(got0, expected0) {
        t.Errorf("0 expected %v got %v", expected0, got0)
    }

    got, ok := likeIPV6String2RangeIp("::")
    if !ok {
        t.Errorf("1 expected true got %t", ok)
    }
    expected := rangeIP{}
    if !reflect.DeepEqual(got, expected) {
        t.Errorf("1 expected %v got %v", expected, got)
    }

    got2, ok2 := likeIPV6String2RangeIp("::1")
    if !ok2 {
        t.Errorf("2 expected true got %t", ok2)
    }
    expected2 := rangeIP{{}, {}, {}, {}, {}, {}, {}, {1, 1}}
    if !reflect.DeepEqual(got2, expected2) {
        t.Errorf("2 expected %v got %v", expected2, got2)
    }

    got3, ok3 := likeIPV6String2RangeIp("a:b::1")
    if !ok3 {
        t.Errorf("3 expected true got %t", ok3)
    }
    expected3 := rangeIP{{10, 10}, {11, 11}, {}, {}, {}, {}, {}, {1, 1}}
    if !reflect.DeepEqual(got3, expected3) {
        t.Errorf("3 expected %v got %v", expected3, got3)
    }

    got4, ok4 := likeIPV6String2RangeIp("1::1008:a:b:c:11")
    if !ok4 {
        t.Errorf("4 expected true got %t", ok4)
    }
    expected4 := rangeIP{{1, 1}, {}, {}, {4104, 4104}, {10, 10}, {11, 11}, {12, 12}, {17, 17}}
    if !reflect.DeepEqual(got4, expected4) {
        t.Errorf("4 expected %v got %v", expected4, got4)
    }

    got5, ok5 := likeIPV6String2RangeIp("1::*:a:b:c:11")
    if !ok5 {
        t.Errorf("5 expected true got %t", ok5)
    }
    expected5 := rangeIP{{1, 1}, {}, {}, {0, 0xffff}, {10, 10}, {11, 11}, {12, 12}, {17, 17}}
    if !reflect.DeepEqual(got5, expected5) {
        t.Errorf("5 expected %v got %v", expected5, got5)
    }

    got6, ok6 := likeIPV6String2RangeIp("a:b:c:1-ffff::")
    if !ok6 {
        t.Errorf("6 expected true got %t", ok6)
    }
    expected6 := rangeIP{{10, 10}, {11, 11}, {12, 12}, {1, 0xffff}}
    if !reflect.DeepEqual(got6, expected6) {
        t.Errorf("6 expected %v got %v", expected6, got6)
    }

    got7, ok7 := likeIPV6String2RangeIp("a:b:c:1-ffff:1-2:1-2005::")
    if !ok7 {
        t.Errorf("7 expected true got %t", ok7)
    }
    expected7 := rangeIP{{10, 10}, {11, 11}, {12, 12}, {1, 0xffff}, {1, 2}, {1, 8197}}
    if !reflect.DeepEqual(got7, expected7) {
        t.Errorf("7 expected %v got %v", expected7, got7)
    }

    // false

    _, ok8 := likeIPV6String2RangeIp("a:b:c:1-ffff:1-2:1-2005:::")
    if ok8 {
        t.Errorf("8 expected false got %t", ok7)
    }
    _, ok9 := likeIPV6String2RangeIp("*")
    if ok9 {
        t.Errorf("9 expected false got %t", ok9)
    }
    _, ok10 := likeIPV6String2RangeIp("*:*:*:")
    if ok10 {
        t.Errorf("10 expected false got %t", ok10)
    }
    _, ok11 := likeIPV6String2RangeIp(":::")
    if ok11 {
        t.Errorf("11 expected false got %t", ok11)
    }
}

func TestIpv6Equals(t *testing.T)  {
    // true
    ok := ipv6Equals(rangeIP{}, IPV6{})
    if !ok {
        t.Errorf("1 expected true got %t", ok)
    }
    ok2 := ipv6Equals(rangeIP{{1, 2}}, IPV6{1})
    if !ok2 {
        t.Errorf("2 expected true got %t", ok2)
    }
    ok3 := ipv6Equals(rangeIP{{1, 2}, {}, {}, {}, {}, {}, {}, {0,1008}}, IPV6{1, 0, 0, 0, 0, 0, 0, 1008})
    if !ok3 {
        t.Errorf("3 expected true got %t", ok3)
    }

    // false

    ok4 := ipv6Equals(rangeIP{{1, 2}, {}, {}, {}, {}, {}, {}, {0,1008}}, IPV6{1, 0, 0, 0, 0, 0, 0, 1009})
    if ok4 {
        t.Errorf("4 expected false got %t", ok4)
    }
    ok5 := ipv6Equals(rangeIP{{1, 2}, {}, {}, {}, {}, {}, {}, {0,1008}}, IPV6{1, 0, 0, 0, 1, 0, 0, 1008})
    if ok5 {
        t.Errorf("5 expected false got %t", ok5)
    }
}

func TestNewFirewall(t *testing.T) {
    f := NewFirewall()
    err := f.LoadIP("192.168.0.*")
    if err != nil {
        t.Errorf("1 expected nil got %v", err)
    }
    err = f.LoadIP("a:b:c:e:f:1:2:3")
    if err != nil {
        t.Errorf("2 expected nil got %v", err)
    }
    err = f.LoadIP("a:b:c:e:f:1:*:1")
    if err != nil {
        t.Errorf("3 expected nil got %v", err)
    }
    err = f.LoadIP("a:b:c:e:f:1:*:1-1008")
    if err != nil {
        t.Errorf("4 expected nil got %v", err)
    }
    if !f.Match("192.168.0.100") {
        t.Errorf("10 expected true got false")
    }
    if !f.Match("a:b:c:e:f:1:2:3") {
        t.Errorf("11 expected true got false")
    }
    if !f.Match("a:b:c:e:f:1:ffff:1") {
        t.Errorf("12 expected true got false")
    }
    if !f.Match("a:b:c:e:f:1:1:1007") {
        t.Errorf("13 expected true got false")
    }
    err = f.LoadIP("a:b:c:d:*::")
    if err != nil {
        t.Errorf("5 expected nil got %v", err)
    }
    if !f.Match("a:b:c:d:ffff:0:0:0") {
        t.Errorf("14 expected true got false")
    }
}