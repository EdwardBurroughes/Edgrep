package edgrep

type RegexMetaObject struct {
	Pattern        string
	Exclude        bool
	Casesensitvity bool
}

type GrepFileOptions struct {
	Regexmeta     RegexMetaObject
	Filepath      string
	Recursiveflag bool
	Rootname      string
}
