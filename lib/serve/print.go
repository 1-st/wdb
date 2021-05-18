package serve

import (
	"fmt"
	"github.com/fatih/color"
	"io"
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

func PrintHelp(out io.Writer) {
	fmt.Fprintln(out, "指令列表:")
	fmt.Fprintln(out, "	list						打印记录在数据库中的单词和词组")
	fmt.Fprintln(out, "	help						打印帮助")
	fmt.Fprintln(out, "	find [word/phrase]			寻找记录的单词或者词组")
	fmt.Fprintln(out, "	add [word] [explain]		添加一个单词,如果已经存在则添加一项释义")
	fmt.Fprintln(out, `	addP [phrase]<[explain]>	添加一个词组，在词组后的尖括号内写上释义`)
	fmt.Fprintln(out, "	del [word/phrase] 			删除单词或者词组")
	fmt.Fprintln(out, "	delX [word][N] 				删除单词或者词组第N个解释")
	fmt.Fprintln(out, "	review [N] 					复习单词,后接复习的数量(默认为3)")
	fmt.Fprintln(out, "	reviewP [N] 				复习词组,后接复习的数量(默认为3)")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "	exit/quit/q					退出")
	//fmt.Println("")
	//fmt.Println("** 形似词使用最大子串寻找,近义词使用预训练word2vec模型,相似系数在config.xml中指定\n词典使用stardict词典")
}
