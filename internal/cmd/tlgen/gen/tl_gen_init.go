package gen

import (
	"sort"

	"github.com/dave/jennifer/jen"
)

var tlPackagePath = "github.com/bs9/spread_service_gogram/internal/encoding/tl"
var errorsPackagePath = "github.com/pkg/errors"

func (g *Generator) generateInit(file *jen.File, _ bool) {
	structs, enums := g.getAllConstructors()

	initFunc := jen.Func().Id("init").Params().Block(
		g.createInitStructs(structs...),
		jen.Line(),
		g.createInitEnums(enums...),
	)

	file.Add(initFunc)
}

func (*Generator) createInitStructs(itemNames ...string) jen.Code {
	sort.Strings(itemNames)

	structs := make([]jen.Code, len(itemNames))
	for i, item := range itemNames {
		structs[i] = jen.Op("&").Id(item).Block()
	}

	return jen.Qual(tlPackagePath, "RegisterObjects").Call(
		structs...,
	)
}

func (*Generator) createInitEnums(itemNames ...string) jen.Code {
	sort.Strings(itemNames)

	enums := make([]jen.Code, len(itemNames))
	for i, item := range itemNames {
		enums[i] = jen.Id(item)
	}

	return jen.Qual(tlPackagePath, "RegisterEnums").Call(
		enums...,
	)
}
