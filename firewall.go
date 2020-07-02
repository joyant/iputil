package iputil

import (
    "errors"
    "strings"
)

var RangeDelimiter = '-' //允许引用类库这修改此变量，以支持如~等分隔符

var (
    ErrNotIPV6 = errors.New("not ipv6")
    ErrNotIPV4 = errors.New("not ipv4")
)

type rangeFragment [2]uint16

type rangeIP [8]rangeFragment

type Firewall interface {
    LoadIP(ip string) error
    Match(ip string) bool
}

func NewFirewall() Firewall {
    return new(sliceFirewall)
}

type sliceFirewall struct {
    ips []rangeIP
}

func (f *sliceFirewall)LoadIP(ip string) error {
    if strings.Contains(ip, ":") {
        if p, ok := likeIPV6String2RangeIp(ip); ok {
            f.ips = append(f.ips, p)
            return nil
        }
        return ErrNotIPV6
    }
    if p, ok := likeIPV4String2RangeIp(ip); ok {
        f.ips = append(f.ips, p)
        return nil
    }
    return ErrNotIPV4
}

func (f *sliceFirewall)Match(ip string) bool {
    var v6 IPV6
    var ok bool
    if strings.Contains(ip, ":") {
        v6, ok = String2IPV6(ip)
        if !ok {
            return false
        }
    } else {
        var v4 IPV4
        v4, ok = String2IPV4(ip)
        if !ok {
            return false
        }
        for k, v := range v4 {
            v6[4+k] = uint16(v)
        }
    }
    for _, v := range f.ips {
        if ipv6Equals(v, v6) {
            return true
        }
    }
    return false
}

func ipv6Equals(a rangeIP, b IPV6) bool {
    for k := range a {
        r := a[k]
        i := b[k]
        if !(i >= r[0] && i <= r[1]) {
            return false
        }
    }
    return true
}

func likeIPV6String2RangeIp(s string) (ip rangeIP, is bool) {
    doubleColon := 0
    doubleColonIndex := -1
    last := 0
    fragments := make([]rangeFragment, 0)
    for i, l := 0, len(s); i < l; {
        if len(fragments) >= 8 {
            is = false
            return
        }
        if s[i] == ':' {
            if i > 0 && s[i-1] == ':' {
                is = false
                return
            }
            if len(s[last:i]) > 0 {
                if numeric, ok := ipv6RangeFragment(s[last:i]); !ok {
                    is = false
                    return
                } else {
                    fragments = append(fragments, numeric)
                }
                last = i+1
            }
            if l > i+1 && s[i+1] == ':' {
                if doubleColon > 0 {
                    is = false
                    return
                }
                doubleColon ++
                i += 2
                last = i
                doubleColonIndex = len(fragments)
                continue
            }
        } else if i == l-1 {
            if numeric, ok := ipv6RangeFragment(s[last:]); !ok {
                is = false
                return
            } else {
                fragments = append(fragments, numeric)
            }
        }
        i ++
    }
    if doubleColon > 0 {
        is = len(fragments) <= 6
        if is {
            copy(ip[:doubleColonIndex], fragments[:doubleColonIndex])
            index := 8-(len(fragments)-doubleColonIndex)
            if index == 8 {
                index --
            }
            copy(ip[index:], fragments[doubleColonIndex:])
        }
        return
    }
    is = len(fragments) == 8
    if is {
        copy(ip[:], fragments)
    }
    return
}

func likeIPV4String2RangeIp(s string) (ip rangeIP, is bool) {
    last := 0
    dot := 0
    fragments := make([]rangeFragment, 0)
    for k, v := range s {
        if len(fragments) >= 4 {
            return
        }
        if v == '.' {
            if k <= last {
                return
            }
            n, ok := ipv4RangeFragment(s[last:k])
            if !ok {
                return
            }
            dot ++
            last = k + 1
            fragments = append(fragments, n)
        } else if k == len(s)-1 {
            if len(s[last:]) == 0 || len(s[last:]) > 3 {
                return
            }
            n, ok := ipv4RangeFragment(s[last:])
            if !ok {
                return
            }
            fragments = append(fragments, n)
        }
    }
    if dot == 3 && len(fragments) == 4 {
        copy(ip[4:], fragments[:])
        is = true
    }
    return
}

func ipv6RangeFragment(s string) (r rangeFragment, is bool) {
    if s == "*" {
        r[1] = 0xFFFF
        is = true
    } else if strings.Contains(s, string(RangeDelimiter)) {
        rs := strings.Split(s, string(RangeDelimiter))
        if len(rs) == 2 {
            for i := 0; i < len(rs); i++ {
                if n, ok := ipv6FragmentNumeric(rs[i]); ok {
                    r[i] = n
                } else {
                    return
                }
            }
            is = true
        }
    } else if n, ok := ipv6FragmentNumeric(s); ok {
        return rangeFragment{n, n}, true
    }
    return
}

func ipv4RangeFragment(s string) (r rangeFragment, is bool) {
    if s == "*" {
        r[1] = 0xFF
        is = true
    } else if strings.Contains(s, string(RangeDelimiter)) {
        rs := strings.Split(s, string(RangeDelimiter))
        if len(rs) == 2 {
            for i := 0; i < len(rs); i++ {
                if n := str2int(rs[i]); n >= 0 && n <= 0xFF {
                    r[i] = uint16(n)
                } else {
                    return
                }
            }
            is = true
        }
    } else if n := str2int(s); n >= 0 && n <= 0xFF {
        r[0] = uint16(n)
        r[1] = r[0]
        is = true
    }
    return
}