package odict

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang/snappy"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/google/uuid"
	"github.com/odict/odict/go/models"
	"github.com/odict/odict/go/schema"
)

type xmlDefinitionGroup struct {
	XMLName     xml.Name `xml:"group"`
	Definitions []string `xml:"definition"`
	Description string   `xml:"description,attr"`
}

type xmlUsage struct {
	XMLName          xml.Name             `xml:"usage"`
	POS              string               `xml:"pos,attr"`
	DefinitionGroups []xmlDefinitionGroup `xml:"group"`
	Definitions      []string             `xml:"definition"`
}

type xmlEtymology struct {
	XMLName     xml.Name   `xml:"ety"`
	Description string     `xml:"description,attr"`
	Usages      []xmlUsage `xml:"usage"`
}

type xmlEntry struct {
	XMLName     xml.Name       `xml:"entry"`
	Term        string         `xml:"term,attr"`
	Etymologies []xmlEtymology `xml:"ety"`
}

type xmlDictionary struct {
	XMLName xml.Name   `xml:"dictionary"`
	Name    string     `xml:"name,attr"`
	Entries []xmlEntry `xml:"entry"`
}

func xmlToDictionary(xmlStr string) models.Dictionary {
	var dictionary models.Dictionary

	xml.Unmarshal([]byte(xmlStr), &dictionary)

	return dictionary
}

func getDefinitionsVectorFromUsage(builder *flatbuffers.Builder, usage models.Usage) flatbuffers.UOffsetT {
	definitions := usage.Definitions

	var defBuffer []flatbuffers.UOffsetT

	for idx := range definitions {
		defBuffer = append(defBuffer, builder.CreateString(definitions[idx]))
	}

	defCount := len(defBuffer)

	schema.GroupStartDefinitionsVector(builder, defCount)

	for i := defCount - 1; i >= 0; i-- {
		builder.PrependUOffsetT(defBuffer[i])
	}

	return builder.EndVector(defCount)
}

func getDefinitionsVectorFromGroup(builder *flatbuffers.Builder, group models.Group) flatbuffers.UOffsetT {
	definitions := group.Definitions

	var defBuffer []flatbuffers.UOffsetT

	for idx := range definitions {
		defBuffer = append(defBuffer, builder.CreateString(definitions[idx]))
	}

	defCount := len(defBuffer)

	schema.GroupStartDefinitionsVector(builder, defCount)

	for i := defCount - 1; i >= 0; i-- {
		builder.PrependUOffsetT(defBuffer[i])
	}

	return builder.EndVector(defCount)
}

func getGroupsVector(builder *flatbuffers.Builder, usage models.Usage) flatbuffers.UOffsetT {
	groups := usage.Groups

	var groupBuffer []flatbuffers.UOffsetT

	for idx := range groups {
		group := groups[idx]
		groupID := builder.CreateString(strconv.Itoa(idx))
		groupDescription := builder.CreateString(group.Description)
		groupDefinitions := getDefinitionsVectorFromGroup(builder, group)

		schema.GroupStart(builder)
		schema.GroupAddId(builder, groupID)
		schema.GroupAddDescription(builder, groupDescription)
		schema.EtymologyAddUsages(builder, groupDefinitions)

		groupBuffer = append(groupBuffer, schema.EtymologyEnd(builder))
	}

	groupCount := len(groupBuffer)

	schema.UsageStartGroupsVector(builder, groupCount)

	for i := groupCount - 1; i >= 0; i-- {
		builder.PrependUOffsetT(groupBuffer[i])
	}

	return builder.EndVector(groupCount)
}

func resolveSchemaPOS(pos models.PartOfSpeech) schema.POS {
	posMap := map[models.PartOfSpeech]schema.POS{
		"adjective":    schema.POSadj,
		"adj":          schema.POSadj,
		"adverb":       schema.POSadv,
		"adv":          schema.POSadv,
		"verb":         schema.POSverb,
		"v":            schema.POSverb,
		"n":            schema.POSnoun,
		"noun":         schema.POSnoun,
		"pronoun":      schema.POSpronoun,
		"pn":           schema.POSpronoun,
		"prep":         schema.POSprep,
		"preposition":  schema.POSprep,
		"conj":         schema.POSconj,
		"conjugation":  schema.POSconj,
		"intj":         schema.POSintj,
		"interjection": schema.POSintj,
		"prefix":       schema.POSprefix,
		"pre":          schema.POSprefix,
		"suffix":       schema.POSsuffix,
		"suf":          schema.POSsuffix,
		"particle":     schema.POSparticle,
		"part":         schema.POSparticle,
		"article":      schema.POSarticle,
		"art":          schema.POSarticle,
		"unknown":      schema.POSunknown,
	}

	if val, ok := posMap[pos]; ok {
		return val
	} else if len(strings.TrimSpace(string(pos))) == 0 {
		return schema.POSunknown
	} else {
		panic(fmt.Sprintf("Compilation error: invalid part-of-speech used: %s", pos))
	}
}

func getUsagesVector(builder *flatbuffers.Builder, ety models.Etymology) flatbuffers.UOffsetT {
	usages := ety.Usages

	var usageBuffer []flatbuffers.UOffsetT

	for key := range usages.Iterable {
		usage := usages.Get(key)
		usagePOS := resolveSchemaPOS(usage.POS)
		usageDefinitionGroups := getGroupsVector(builder, usage)
		usageDefinitions := getDefinitionsVectorFromUsage(builder, usage)

		schema.UsageStart(builder)
		schema.UsageAddPos(builder, usagePOS)
		schema.UsageAddGroups(builder, usageDefinitionGroups)
		schema.UsageAddDefinitions(builder, usageDefinitions)

		usageBuffer = append(usageBuffer, schema.UsageEnd(builder))
	}

	usageCount := len(usageBuffer)

	schema.EtymologyStartUsagesVector(builder, usageCount)

	for i := usageCount - 1; i >= 0; i-- {
		builder.PrependUOffsetT(usageBuffer[i])
	}

	return builder.EndVector(usageCount)
}

func getEtymologiesVector(builder *flatbuffers.Builder, entry models.Entry) flatbuffers.UOffsetT {
	etymologies := entry.Etymologies

	var etyBuffer []flatbuffers.UOffsetT

	for idx := range etymologies {
		ety := etymologies[idx]
		etyID := builder.CreateString(strconv.Itoa(idx))
		etyDescription := builder.CreateString(ety.Description)
		etyUsages := getUsagesVector(builder, ety)

		schema.EtymologyStart(builder)
		schema.EtymologyAddId(builder, etyID)
		schema.EtymologyAddDescription(builder, etyDescription)
		schema.EtymologyAddUsages(builder, etyUsages)

		etyBuffer = append(etyBuffer, schema.EtymologyEnd(builder))
	}

	etyCount := len(etyBuffer)

	schema.EntryStartEtymologiesVector(builder, etyCount)

	for i := etyCount - 1; i >= 0; i-- {
		builder.PrependUOffsetT(etyBuffer[i])
	}

	return builder.EndVector(etyCount)
}

func getEntriesVector(builder *flatbuffers.Builder, dictionary models.Dictionary) flatbuffers.UOffsetT {
	entries := dictionary.Entries

	var entryBuffer []flatbuffers.UOffsetT

	for key := range entries.Iterable {
		entry := entries.Get(key)
		entryTerm := builder.CreateString(entry.Term)
		entryEtymologies := getEtymologiesVector(builder, entry)

		schema.EntryStart(builder)
		schema.EntryAddTerm(builder, entryTerm)
		schema.EntryAddEtymologies(builder, entryEtymologies)

		entryBuffer = append(entryBuffer, schema.EntryEnd(builder))
	}

	entryCount := len(entryBuffer)

	schema.DictionaryStartEntriesVector(builder, entryCount)

	for i := entryCount - 1; i >= 0; i-- {
		builder.PrependUOffsetT(entryBuffer[i])
	}

	return builder.EndVector(entryCount)
}

func dictionaryToBytes(dictionary models.Dictionary) []byte {
	builder := flatbuffers.NewBuilder(1024)

	id := builder.CreateString(uuid.New().String())
	name := builder.CreateString(dictionary.Name)
	entries := getEntriesVector(builder, dictionary)

	schema.DictionaryStart(builder)
	schema.DictionaryAddId(builder, id)
	schema.DictionaryAddName(builder, name)
	schema.DictionaryAddEntries(builder, entries)

	builder.Finish(schema.DictionaryEnd(builder))

	return builder.FinishedBytes()
}

func createODictFile(outputPath string, dictionary models.Dictionary) {
	dictionaryBytes := dictionaryToBytes(dictionary)
	compressed := snappy.Encode(nil, dictionaryBytes)
	file, err := os.Create(outputPath)

	Check(err)

	defer file.Close()

	signature := []byte("ODICT")
	version := Uint16ToBytes(2)
	compressedSize := uint64(len(compressed))
	compressedSizeBytes := Uint64ToBytes(compressedSize)

	writer := bufio.NewWriter(file)

	sigBytes, sigErr := writer.Write(signature)
	versionBytes, versionErr := writer.Write(version)
	contentSizeBytes, contentCountErr := writer.Write(compressedSizeBytes)
	contentBytes, contentErr := writer.Write(compressed)
	total := sigBytes + versionBytes + contentSizeBytes + contentBytes

	Check(sigErr)
	Check(versionErr)
	Check(contentCountErr)
	Check(contentErr)

	Assert(sigBytes == 5, "Signature bytes do not equal 5")
	Assert(versionBytes == 2, "Version bytes do not equal 2")
	Assert(contentSizeBytes == 8, "Content byte count does not equal 8")
	Assert(contentBytes == int(compressedSize), "Content does not equal the computed byte count")

	writer.Flush()

	fmt.Printf("Wrote %d bytes to path: %s\n", total, outputPath)
}

// WriteDictionary generates an ODict binary file given
// a ODXML input file path
func WriteDictionary(xmlStr, outputPath string) {
	createODictFile(outputPath, xmlToDictionary(xmlStr))
}
