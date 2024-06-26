package crt

import (
	lang "github.com/mt1976/crt/language"
	page "github.com/mt1976/crt/page"
	actn "github.com/mt1976/crt/page/actions"
	term "github.com/mt1976/crt/terminal"
)

func Sample() {
	vp := term.NewWithSize(80, 25)
	pg := page.NewPage(&vp, lang.New("Testing testing 123"))
	pg.AddMenuOption(1, "Option 1 Description", "Option 1 Action", "Option")
	pg.AddFieldValuePair("Field1", "Value1")
	pg.AddFieldValuePair("Field2", "Value2")
	pg.AddFieldValuePair("Field3", "Value3")
	pg.AddFieldValuePair("Field4", "Value4")
	pg.AddColumnsTitle("Column1", "Column2", "Column3", "Column4")
	pg.AddColumns("Value1", "Value2", "Value3", "Value4")
	pg.AddColumns("Value1", "Value2", "Value3", "Value4")
	pg.AddColumns("Value1", "Value2", "Value3", "Value4")
	pg.AddFieldValuePair("Field5", "Value5")
	pg.AddFieldValuePair("Field6", "Value6")
	pg.AddFieldValuePair("Field7", "Value7")
	pg.AddFieldValuePair("Field8", "Value8")
	pg.AddFieldValuePair("Field9", "Value9")
	pg.AddFieldValuePair("Field10", "Value10")
	pg.AddFieldValuePair("Field11", "Value11")
	pg.AddFieldValuePair("Field12", "Value12")
	pg.AddFieldValuePair("Field13", "Value13")
	pg.AddFieldValuePair("Field14", "Value14")
	pg.AddFieldValuePair("Field15", "Value15")
	pg.AddFieldValuePair("Field16", "Value16")
	pg.AddAction(actn.New("K"))
	pg.Display_Actions()
}
