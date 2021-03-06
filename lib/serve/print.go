package serve

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"wdb/lib/constant"
)

func PrintLogo() {
	color.HiBlue(`
__/\\\______________/\\\__/\\\\\\\\\\\\_____/\\\\\\\\\\\\\___        
 _\/\\\_____________\/\\\_\/\\\////////\\\__\/\\\/////////\\\_       
  _\/\\\_____________\/\\\_\/\\\______\//\\\_\/\\\_______\/\\\_      
   _\//\\\____/\\\____/\\\__\/\\\_______\/\\\_\/\\\\\\\\\\\\\\__     
    __\//\\\__/\\\\\__/\\\___\/\\\_______\/\\\_\/\\\/////////\\\_    
     ___\//\\\/\\\/\\\/\\\____\/\\\_______\/\\\_\/\\\_______\/\\\_   
      ____\//\\\\\\//\\\\\_____\/\\\_______/\\\__\/\\\_______\/\\\_  
       _____\//\\\__\//\\\______\/\\\\\\\\\\\\/___\/\\\\\\\\\\\\\/__ 
        ______\///____\///_______\////////////_____\/////////////____
`)
}

func PrintHelp() {
	out := os.Stdout
	fmt.Fprintln(out, "指令列表: **按TAB键可以自动补全")
	fmt.Fprintln(out, "	list						"+constant.CMD_LIST_DSP)
	fmt.Fprintln(out, "	add					"+constant.CMD_ADD_DSP)
	fmt.Fprintln(out, "	rm 			"+constant.CMD_RM_DSP)
	fmt.Fprintln(out)
	fmt.Fprintln(out, "	review 					"+constant.CMD_REVIEW_DSP)
	fmt.Fprintln(out, "	reviewP 				"+constant.CMD_REVIEWP_DSP)
	fmt.Fprintln(out, "	print 					"+constant.CMD_PRINT_DSP)
	fmt.Fprintln(out)
	fmt.Fprintln(out, "	help						"+constant.CMD_HELP_DSP)
	fmt.Fprintln(out, "	exit/quit/q					退出")
}
