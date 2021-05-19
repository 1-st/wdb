package constant

import "encoding/xml"

type Cat struct {
	XMLName xml.Name `xml:"at,omitempty" json:"at,omitempty"`
	String  string   `xml:",chardata" json:",omitempty"`
}

type Cdb struct {
	XMLName    xml.Name  `xml:"db,omitempty" json:"db,omitempty"`
	Cphrases   *Cphrases `xml:"phrases,omitempty" json:"phrases,omitempty"`
	Cwords     *Cwords   `xml:"words,omitempty" json:"words,omitempty"`
}

type Cphrases struct {
	XMLName xml.Name   `xml:"phrases,omitempty" json:"phrases,omitempty"`
	Cphrase []*Cphrase `xml:"phrase,omitempty" json:"phrase,omitempty"`
}

type Cexplains struct {
	XMLName xml.Name `xml:"explains,omitempty" json:"explains,omitempty"`
	String string `xml:",chardata" json:",omitempty"`
}

type Cok struct {
	XMLName xml.Name `xml:"ok,omitempty" json:"ok,omitempty"`
	String  string   `xml:",chardata" json:",omitempty"`
}

type Csentence struct {
	XMLName     xml.Name     `xml:"sentence,omitempty" json:"sentence,omitempty"`
	Cexplain    *Cexplain    `xml:"explain,omitempty" json:"explain,omitempty"`
	Cid         *Cid         `xml:"id,omitempty" json:"id,omitempty"`
	Ctranslated *Ctranslated `xml:"translated,omitempty" json:"translated,omitempty"`
}


type Cphrase struct {
	XMLName xml.Name     `xml:"phrase,omitempty" json:"phrase,omitempty"`
	Cexplains *Cexplains `xml:"explains,omitempty" json:"explains,omitempty"`
	Cid *Cid             `xml:"id,omitempty" json:"id,omitempty"`
	Cok *Cok             `xml:"ok,omitempty" json:"ok,omitempty"`
	Cviews *Cviews       `xml:"views,omitempty" json:"views,omitempty"`
}

type Csentences struct {
	XMLName   xml.Name     `xml:"sentences,omitempty" json:"sentences,omitempty"`
	Csentence []*Csentence `xml:"sentence,omitempty" json:"sentence,omitempty"`
}

type Ctranslated struct {
	XMLName xml.Name `xml:"translated,omitempty" json:"translated,omitempty"`
	String  string   `xml:",chardata" json:",omitempty"`
}

type Cviews struct {
	XMLName xml.Name `xml:"views,omitempty" json:"views,omitempty"`
	Cat     []*Cat   `xml:"at,omitempty" json:"at,omitempty"`
}

type Cword struct {
	XMLName   xml.Name `xml:"word,omitempty" json:"word,omitempty"`
	Cid       *Cid     `xml:"id,omitempty" json:"id,omitempty"`
	Cok       *Cok     `xml:"ok,omitempty" json:"ok,omitempty"`
	Cviews    *Cviews  `xml:"views,omitempty" json:"views,omitempty"`
}

type Cwords struct {
	XMLName xml.Name `xml:"words,omitempty" json:"words,omitempty"`
	Cword   []*Cword `xml:"word,omitempty" json:"word,omitempty"`
}

type Caffix struct {
	XMLName                 xml.Name                 `xml:"affix,omitempty" json:"affix,omitempty"`
	Ccommon_dash_prefixes   *Ccommon_dash_prefixes   `xml:"common-prefixes,omitempty" json:"common-prefixes,omitempty"`
	Csuffixes               *Csuffixes               `xml:"suffixes,omitempty" json:"suffixes,omitempty"`
	Cuncommon_dash_prefixes *Cuncommon_dash_prefixes `xml:"uncommon-prefixes,omitempty" json:"uncommon-prefixes,omitempty"`
}

type Cbody struct {
	XMLName  xml.Name  `xml:"body,omitempty" json:"body,omitempty"`
	Caffix   *Caffix   `xml:"affix,omitempty" json:"affix,omitempty"`
	Cconfig  *Cconfig  `xml:"config,omitempty" json:"config,omitempty"`
	Cescapes *Cescapes `xml:"escapes,omitempty" json:"escapes,omitempty"`
}

type Ccommon_dash_prefixes struct {
	XMLName xml.Name `xml:"common-prefixes,omitempty" json:"common-prefixes,omitempty"`
	Cid     []*Cid   `xml:"id,omitempty" json:"id,omitempty"`
}

type Cconfig struct {
	XMLName                                        xml.Name                                        `xml:"config,omitempty" json:"config,omitempty"`
	Cmax_dash_similar_dash_words                   *Cmax_dash_similar_dash_words                   `xml:"max-similar-words,omitempty" json:"max-similar-words,omitempty"`
	Cmax_dash_view_dash_record                     *Cmax_dash_view_dash_record                     `xml:"max-view-record,omitempty" json:"max-view-record,omitempty"`
	Csimilar_dash_word_dash_threshold_dash_diff    *Csimilar_dash_word_dash_threshold_dash_diff    `xml:"similar-word-threshold-diff,omitempty" json:"similar-word-threshold-diff,omitempty"`
	Csimilar_dash_word_dash_threshold_dash_network *Csimilar_dash_word_dash_threshold_dash_network `xml:"similar-word-threshold-network,omitempty" json:"similar-word-threshold-network,omitempty"`
}

type Cescapes struct {
	XMLName xml.Name `xml:"escapes,omitempty" json:"escapes,omitempty"`
	String  string   `xml:",chardata" json:",omitempty"`
}

type Cexplain struct {
	XMLName xml.Name `xml:"explain,omitempty" json:"explain,omitempty"`
	String  string   `xml:",chardata" json:",omitempty"`
}

type Cid struct {
	XMLName xml.Name `xml:"id,omitempty" json:"id,omitempty"`
	String  string   `xml:",chardata" json:",omitempty"`
}

type Cmax_dash_similar_dash_words struct {
	XMLName xml.Name `xml:"max-similar-words,omitempty" json:"max-similar-words,omitempty"`
	String  string   `xml:",chardata" json:",omitempty"`
}

type Cmax_dash_view_dash_record struct {
	XMLName xml.Name `xml:"max-view-record,omitempty" json:"max-view-record,omitempty"`
	String  string   `xml:",chardata" json:",omitempty"`
}

type Cprefix struct {
	XMLName  xml.Name  `xml:"prefix,omitempty" json:"prefix,omitempty"`
	Cexplain *Cexplain `xml:"explain,omitempty" json:"explain,omitempty"`
	Cid      []*Cid    `xml:"id,omitempty" json:"id,omitempty"`
}

type Csimilar_dash_word_dash_threshold_dash_diff struct {
	XMLName xml.Name `xml:"similar-word-threshold-diff,omitempty" json:"similar-word-threshold-diff,omitempty"`
	String  string   `xml:",chardata" json:",omitempty"`
}

type Csimilar_dash_word_dash_threshold_dash_network struct {
	XMLName xml.Name `xml:"similar-word-threshold-network,omitempty" json:"similar-word-threshold-network,omitempty"`
	String  string   `xml:",chardata" json:",omitempty"`
}

type Csuffixes struct {
	XMLName xml.Name `xml:"suffixes,omitempty" json:"suffixes,omitempty"`
	Cid     []*Cid   `xml:"id,omitempty" json:"id,omitempty"`
}

type Cuncommon_dash_prefixes struct {
	XMLName xml.Name   `xml:"uncommon-prefixes,omitempty" json:"uncommon-prefixes,omitempty"`
	Cprefix []*Cprefix `xml:"prefix,omitempty" json:"prefix,omitempty"`
}
