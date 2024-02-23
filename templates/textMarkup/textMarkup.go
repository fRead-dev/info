
type fRead_textMarkup_pObj struct {
    begin string
    end   string
}

type fRead_textMarkup_alignObj struct {
    left   fRead_textMarkup_pObj
    right  fRead_textMarkup_pObj
    center fRead_textMarkup_pObj
}
type fRead_textMarkup_styleObj struct {
    throughline fRead_textMarkup_pObj
    underline   fRead_textMarkup_pObj
    italic      fRead_textMarkup_pObj
    bold        fRead_textMarkup_pObj
}

type fRead_textMarkupObj struct {
    ver string

    align fRead_textMarkup_alignObj
    style fRead_textMarkup_styleObj

    del []string
}

var Obj_fRead_textMarkup = fRead_textMarkupObj{
    ver: "1.0",

    align.left:   fRead_textMarkup_pObj{"::BL", "::EL"},
    align.right:  fRead_textMarkup_pObj{"::BR", "::ER"},
    align.center: fRead_textMarkup_pObj{"::BC", "::EC"},

    style.throughline: fRead_textMarkup_pObj{"::BS", "::ES"},
    style.underline:   fRead_textMarkup_pObj{"::BU", "::EU"},
    style.italic:      fRead_textMarkup_pObj{"::BI", "::EI"},
    style.bold:        fRead_textMarkup_pObj{"::BB", "::EB"},

    del: []string{"\r", "\0", "\t"},
}