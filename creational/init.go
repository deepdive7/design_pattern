package creational

// SF ...
var SF *ShapeFactory

//CF ...
var CF *ColorFactory

// King is only one king obj
var singleKing *king

//KingReady ...
var KingReady chan int
var kingReqSeq int
var persons map[string]Person

func init() {
	SF = NewShapeFactory()
	CF = NewColorFactory()
	KingReady = make(chan int)
	kingReqSeq = 0
	persons = make(map[string]Person, 10)
	persons["english"] = english{"red", "blue", 0}
	persons["chinese"] = chinese{"yellow", "blackwhite", 1}
}
