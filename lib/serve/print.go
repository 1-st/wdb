package serve

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"wdb/lib/constant"
)

func PrintLogo() {
	color.Blue(`
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
	fmt.Fprintln(out, "	help						"+constant.CMD_HELP_DSP)
	fmt.Fprintln(out, "	find [word/phrase]			"+constant.CMD_FIND_DSP)
	fmt.Fprintln(out, "	add [word]					"+constant.CMD_ADD_DSP)
	fmt.Fprintln(out, `	addP [phrase]<[explain]>	`+constant.CMD_ADDP_DSP)
	fmt.Fprintln(out, "	del [word/phrase] 			"+constant.CMD_DEL_DSP)
	fmt.Fprintln(out, "	review [N] 					"+constant.CMD_REVIEW_DSP)
	fmt.Fprintln(out, "	reviewP [N] 				"+constant.CMD_REVIEWP_DSP)
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "	exit/quit/q					退出")
}
