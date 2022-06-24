package vdfparser

import (
	"errors"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"text/scanner"

	"github.com/BenLubar/vpk"
)

type VdfNode map[string]interface{}

func parseNode(s *scanner.Scanner, first bool) (*VdfNode, error) {
	result := &VdfNode{}
	expectKey := true
	curKey := ""

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tt := s.TokenText()
		switch tt {
		case "{":
			if expectKey {
				return nil, errors.New("{ unexpected at line " + strconv.Itoa(s.Position.Line))
			}
			nod, err := parseNode(s, false)
			if err != nil {
				return nil, err
			}
			(*result)[curKey] = nod
			expectKey = true
		case "}":
			return result, nil
		default:
			str, err := strconv.Unquote(tt)
			if err == nil {
				tt = str
			}
			if expectKey {
				curKey = tt
				expectKey = false
			} else {
				(*result)[curKey] = tt
				expectKey = true
			}
		}
	}
	if first {
		return result, nil
	}
	return nil, errors.New("unexpected EOF")
}

func parseVdf(vdf string) (*VdfNode, error) {
	var s scanner.Scanner
	s.Init(strings.NewReader(vdf))
	s.Mode = scanner.ScanComments | scanner.SkipComments | scanner.ScanStrings | scanner.ScanIdents
	return parseNode(&s, true)
}

func readVpkFileDescriptor(path string) (http.File, error) {
	opener := vpk.SingleVPK(path)
	file, err := vpk.Open(opener)
	if err != nil {
		return nil, err
	}
	for _, v := range file.Paths() {
		if strings.HasPrefix(v, "missions") {
			return file.Open(v)
		}
	}
	return nil, &ErrInvalidDescriptorFile
}

type VpkInfo struct {
	Num          string
	Name         string
	DisplayTitle string
	StartName    string
}

func ParseVpkMapFileInfo(path string) (info []VpkInfo, err error) {

	file, err := readVpkFileDescriptor(path)
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	vdf, err := parseVdf(string(bytes))
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			err = &ErrInvalidDescriptorFile
		}
	}()
	mission := (*vdf)["mission"]
	modes := (*mission.(*VdfNode))["modes"]
	name := (*mission.(*VdfNode))["FileName"].(string)
	displayTitle := (*mission.(*VdfNode))["DisplayTitle"].(string)
	coop := (*modes.(*VdfNode))["coop"]
	coopMap := coop.(*VdfNode)
	var vpkInfos []VpkInfo
	for num, stage := range *coopMap {
		infoMap := stage.(*VdfNode)
		mapName := (*infoMap)["Map"].(string)
		vpkInfos = append(vpkInfos, VpkInfo{
			Num:          num,
			Name:         name,
			DisplayTitle: displayTitle,
			StartName:    mapName,
		})
	}
	sort.Slice(vpkInfos, func(i, j int) bool {
		return vpkInfos[i].Num < vpkInfos[j].Num
	})
	return vpkInfos, nil
}
