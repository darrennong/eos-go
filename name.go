package pc

import (
	//"fmt"
)

type uint128 struct {
	Low uint64
	Up uint64
}

func (i uint128)LittleEndian()[]byte{
	b := make([]byte,16)
	_ = b[15] // early bounds check to guarantee safety of writes below
	for shift,x,n:=uint(0),0,i.Low;x<8;x++{
		b[x] = byte(n>>shift)
		shift+=8
	}
	for shift,x,n:=uint(0),8,i.Up;x<16;x++{
		b[x] = byte(n>>shift)
		shift+=8
	}
	return b
}

func StringToName(s string) (uint128, error) {
	sLen := uint(len(s))
	var low uint64
	for shift,i:=uint(4),uint(0); i < 10; i++ {
		if i < sLen {
			c := uint64(charToSymbol(s[i]))<<shift
			low |= c
		}else{
			return uint128{low,0},nil
		}
		shift += 6
	}
	var up uint64
	for shift,i:=uint(0),uint(10);i<20;i++{
		if(i<sLen){
			c := uint64(charToSymbol(s[i]))
			up |= c<<shift
		}else{
			return uint128{low,up},nil
		}
		shift += 6
	}
	return uint128{low,up},nil
}

func charToSymbol(c byte) byte {
	for idx,s:=range base32Alphabet{
		if c==s{
			return byte(idx)
		}
	}
	return 0
}

var base32Alphabet = []byte(".-0123456789abcdefghijklmnopqrstuvwxyz_:<>[]{}()`~")

func NameToString(in uint128) string {
	// ported from libraries/chain/name.cpp in eosio
	a := make([]byte,21)
	mask,tmp,i := uint64(0x3f),in.Low>>4, uint32(0)
	for ; tmp>0; i++ {
		idx:=tmp&mask
		//fmt.Println("idx",idx, len(base32Alphabet),string(a))
		c := base32Alphabet[idx]
		a[i] = c
		tmp >>= 6
	}
	tmp = in.Up
	for ; tmp>0; i++ {
		idx:=tmp&mask
		//fmt.Println("idx",idx, len(base32Alphabet),string(a))
		c := base32Alphabet[idx]
		a[i] = c
		tmp >>= 6
	}
	return string(a)
}

/**
     uint128_t name = 0;
      int i = 0;
      for ( ; str[i] && i < 20; ++i) {
          // NOTE: char_to_symbol() returns char type, and without this explicit
          // expansion to uint64 type, the compilation fails at the point of usage
          // of string_to_name(), where the usage requires constant (compile time) expression.
//           name |= (char_to_symbol(str[i]) & 0x3f) << (64 - 6 * (i + 1));
           uint128_t c = (char_to_symbol(str[i]) & 0x3f);
           c <<= (4 + (6 * i));
           name |= c;
       }

 */