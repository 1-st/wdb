package constant

const (
	CMD_ADD_DSP     = "[word] 添加一个单词"
	CMD_ADDP_DSP    = "[phrase](explain)添加一个词组，在词组后的括号内写上释义,如果已经存在,修改释义"
	CMD_DEL_DSP     = "[word/phrase] 删除单词或者词组"
	CMD_FIND_DSP    = "[word/phrase] 寻找记录的单词或者词组"
	CMD_HELP_DSP    = "打印帮助"
	CMD_LIST_DSP    = "打印记录在数据库中的单词和词组"
	CMD_REVIEW_DSP  = "[N] 复习单词,后接复习的数量(默认为30)"
	CMD_REVIEWP_DSP = "[N] 复习词组,后接复习的数量(默认为30)"
	CMD_OK_DSP      = "[word/phrase] 将单词或者词组设置为已经掌握，以后不再复习,已经掌握的会重新投入复习"
)
