package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`^report.*\.pdf|xml$`)
	fmt.Println(re.Match([]byte(`report-847e5313-1068-4957-95b8-0cef1ab3bcba-BC-2020-03-17-084406-68-01[0].xml`)))
	fmt.Println(re.Match([]byte(`report-847e5313-1068-4957-95b8-0cef1ab3bcba-BC-2020-03-17-084406-68-01[0].pdf`)))

}