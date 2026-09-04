package repolog

type Ledger struct {
	Repositories []Repository `json:"repositories"`
}

type Repository struct {
	Name    string   `json:"name"`
	SrcUrl  string   `json:"srcUrl"`
	DstUrl  string   `json:"dstUrl"`
	Bundles []Bundle `json:"bundles"`
}

type Bundle struct {
	Order    int      `json:"order"`
	Filesize int64    `json:"fizesize"`
	Sha256   string   `json:"sha256"`
	Branches []Branch `json:"branches"`
}

type Branch struct {
	Name    string `json:"name"`
	LastRev string `json:"lastRev"`
	CurRev  string `json:"curRev"`
}

func New() Ledger {
	return Ledger{}
}
