package constant

const (
	CMD_ADD_DSP     = "[word/phrase] | [explain]添加一个单词或者词组,如果是词组，在|后写上解释"
	CMD_RM_DSP     = "[word/phrase] 删除单词或者词组"
	CMD_HELP_DSP    = "打印帮助"
	CMD_LIST_DSP    = "[word/phrase] 打印记录在数据库中的单词和词组,不加参数打印全部"
	CMD_REVIEW_DSP  = "[N] 复习单词,后接复习的数量(默认为30)"
	CMD_REVIEWP_DSP = "[N] 复习词组,后接复习的数量(默认为30)"
	CMD_OK_DSP      = "[word/phrase] 将单词或者词组设置为已经掌握，以后不再复习,已经掌握的会重新投入复习"
)
