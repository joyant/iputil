package iputil

import (
    "encoding/hex"
    "strings"
)

type IPV4 [4]uint8

type IPV6 [8]uint16

func IsIPV4(s string) bool {
    last := 0
    dot := 0
    fragment := 0
    for k, v := range s {
        if fragment >= 4 { // 在已经判断出是否是IPV4的情况下，避免遍历多余的字符串
            return false
        }
        if v == '.' {
            if k <= last {
                return false
            }
            if len(s[last:k]) > 3 {
                return false
            }
            if n := str2int(s[last:k]);n > 255 || n < 0 {
                return false
            }
            dot ++
            last = k + 1
            fragment ++
        } else if k == len(s)-1 {
            if len(s[last:]) == 0 || len(s[last:]) > 3 {
                return false
            }
            if n := str2int(s[last:]);n > 255 || n < 0 {
                return false
            }
            fragment ++
        }
    }
    return dot == 3 && fragment == 4
}

func IsIPV6(s string) bool {
    fragment := 0
    doubleColon := 0
    last := 0
    for i, l := 0, len(s); i < l; {
        if fragment >= 8 { //一个非法的字符串可能非常长，在获取足够的信息能作出判断的情况下，就终止循环
            return false
        }
        if s[i] == ':' {
            if i > 0 && s[i-1] == ':' { // 判断出现":::"的情况
                return false
            }
            if len(s[last:i]) > 0 { // 遇到"a::"的情况，先记录前面的a，后处理"::"
                if !isIPV6Fragment(s[last:i]) {
                    return false
                }
                last = i+1
                fragment ++
            }
            if l > i+1 && s[i+1] == ':' {
                if doubleColon > 0 {
                    return false
                }
                doubleColon ++
                i += 2
                last = i
                continue
            }
        } else if i == l-1 {
            if !isIPV6Fragment(s[last:]) {
                return false
            }
            fragment ++
        }
        i ++
    }
    // "::"是合法的IPV6地址，表示8个0，"::1"和"1::"也是合法的，
    // "::"表示连续的0，所以包含"::"的情况下，fragment最大只可能是6
    if doubleColon > 0 {
        return fragment <= 6
    }
    return fragment == 8
}

func isIPV6Fragment(s string) bool {
    if len(s) > 4 {
        return false
    }
    for _, v := range s {
        if !isHex(byte(v)) {
            return false
        }
    }
    return true
}

func ipv6FragmentNumeric(s string) (n uint16, is bool) {
    if len(s) > 4 {
        return
    }
    for _, v := range s {
        i := hexNumeric(byte(v))
        if i <  0 {
            is = false
            return
        }
        n = n * 16 + uint16(i)
    }
    return n, true
}

func hexNumeric(b byte) int {
    if b >= '0' && b <= '9' {
        return int(b-'0')
    } else if b >= 'a' && b <= 'f' {
        return int(b-'a'+10)
    } else if b >= 'A' && b <= 'F' {
        return int(b-'A'+10)
    }
    return -1
}

func isHex(b byte) bool {
    return isNumeric(b) || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
}

func isNumeric(b byte) bool {
    return b >= '0' && b <= '9'
}

func IsIP(s string) bool {
    for _, v := range s {
        if v == ':' {
            return IsIPV6(s)
        }
    }
    return IsIPV4(s)
}

func String2IPV4(s string) (IPV4, bool) {
    if IsIPV4(s) {
        slice := strings.Split(s, ".")
        ip := IPV4{}
        for k, v := range slice {
            ip[k] = uint8(str2int(v))
        }
        return ip, true
    }
    return IPV4{}, false
}

func String2IPV6(s string) (ip IPV6, is bool) {
    doubleColon := 0
    doubleColonIndex := -1
    last := 0
    fragments := make([]uint16, 0)
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
                if numeric, ok := ipv6FragmentNumeric(s[last:i]); !ok {
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
            if numeric, ok := ipv6FragmentNumeric(s[last:]); !ok {
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

func str2int(s string) int {
    n := 0
    for _, v := range s {
        if v < '0' || v > '9' {
            return -1 // not numeric
        }
        n = n * 10 + int(v-'0')
    }
    return n
}

func IPV42IPV6(ip IPV4) string {
    builder := strings.Builder{}
    builder.WriteString("::")
    for k, v := range ip {
        builder.WriteString(hex.EncodeToString([]byte{v}))
        if k != len(ip)-1 {
            builder.WriteByte(':')
        }
    }
    return builder.String()
}

func StringIPV42IPV6(ip string) (string, bool) {
    if v4, ok := String2IPV4(ip); ok {
        return IPV42IPV6(v4), true
    }
    return "", false
}

