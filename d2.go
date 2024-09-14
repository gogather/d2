package d2

import (
	"github.com/gogather/safemap/v2"
)

type D2 struct {
	md2 *safemap.SafeMap
}

func NewD2() *D2 {
	return &D2{
		md2: safemap.New(),
	}
}

func (md2 *D2) Sections() []string {
	return md2.md2.Keys()
}

func (md2 *D2) Get(section, key string) (value interface{}, exist bool) {
	sect, ok := md2.md2.Get(section)
	if !ok {
		return nil, ok
	}
	sectMap, _ := sect.(*safemap.SafeMap)
	return sectMap.Get(key)
}

func (md2 *D2) GetSection(section string) (value *safemap.SafeMap, exist bool) {
	sect, ok := md2.md2.Get(section)
	if !ok {
		return nil, ok
	}
	sectMap, ok := sect.(*safemap.SafeMap)
	return sectMap, ok
}

func (md2 *D2) Add(section, key string, value interface{}) {
	sect, ok := md2.md2.Get(section)
	if !ok {
		sect = safemap.New()
	}

	sectMap, _ := sect.(*safemap.SafeMap)

	sectMap.Put(key, value)

	md2.md2.Put(section, sectMap)
}

func (md2 *D2) RemoveKey(section, key string) {
	sectMap, ok := md2.md2.Get(section)
	if !ok {
		return
	}
	sect := sectMap.(*safemap.SafeMap)
	_, ok = sect.Get(key)
	if ok {
		sect.Remove(key)
	}
	if len(sect.GetMap()) <= 0 {
		md2.md2.Remove(section)
	} else {
		md2.md2.Put(section, sect)
	}
}

func (md2 *D2) RemoveSection(section string) {
	md2.md2.Remove(section)
}

func (md2 *D2) Clear() {
	md2.md2 = safemap.New()
}

func (md2 *D2) GetMapMap() map[string]map[string]interface{} {
	dumpResult := map[string]map[string]interface{}{}
	for section, lowerMap := range md2.md2.GetMap() {
		sectionMap := map[string]interface{}{}
		for key, value := range lowerMap.(*safemap.SafeMap).GetMap() {
			sectionMap[key] = value
		}
		dumpResult[section] = sectionMap
	}
	return dumpResult
}
