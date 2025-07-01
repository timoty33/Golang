package Strings

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Exercise() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "\033[0;31mPrecisamos de pelo menos 2 argumentos.\033[m")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		stringSplitOld := strings.Split(scan.Text(), old)
		stringNew := strings.Join(stringSplitOld, new)

		fmt.Println(stringNew)
	}
}
