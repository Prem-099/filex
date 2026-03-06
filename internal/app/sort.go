package app

import (
	"sort"
	"strings"
)

func (a *App) SortName() string {
	switch a.Sort{
	case SortName:
		return "name"
	case SortSize:
		return "size"
	case SortTime:
		return "time"
	}
	return ""
}

func (a *App) ApplySort(){
	switch a.Sort{
	case SortName:
		sort.Slice(a.Files, func(i,j int) bool {
			return strings.ToLower(a.Files[i].Name()) < strings.ToLower(a.Files[j].Name()) 
		})
	case SortSize:
		sort.Slice(a.Files, func(i,j int) bool {
			iInfo,_ := a.Files[i].Info()
			jInfo,_ := a.Files[j].Info()
			return iInfo.Size() > jInfo.Size()
		})
	case SortTime:
		sort.Slice(a.Files, func(i,j int) bool {
			iInfo,_ := a.Files[i].Info()
			jInfo,_ := a.Files[j].Info()
			return iInfo.ModTime().After(jInfo.ModTime())  
		})
	}
}

