package battlescribe

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"io"
)

type Roster struct {
	XMLName             xml.Name `xml:"roster"`
	Text                string   `xml:",chardata"`
	Xsd                 string   `xml:"xsd,attr"`
	Xsi                 string   `xml:"xsi,attr"`
	BattleScribeVersion string   `xml:"battleScribeVersion,attr"`
	GameSystemId        string   `xml:"gameSystemId,attr"`
	GameSystemName      string   `xml:"gameSystemName,attr"`
	GameSystemRevision  string   `xml:"gameSystemRevision,attr"`
	ID                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	Xmlns               string   `xml:"xmlns,attr"`
	Costs               struct {
		Text string `xml:",chardata"`
		Cost []struct {
			Text   string `xml:",chardata"`
			TypeId string `xml:"typeId,attr"`
			Value  string `xml:"value,attr"`
			Name   string `xml:"name,attr"`
		} `xml:"cost"`
	} `xml:"costs"`
	CostLimits string `xml:"costLimits"`
	Forces     struct {
		Text  string `xml:",chardata"`
		Force struct {
			Text              string `xml:",chardata"`
			CatalogueId       string `xml:"catalogueId,attr"`
			CatalogueName     string `xml:"catalogueName,attr"`
			CatalogueRevision string `xml:"catalogueRevision,attr"`
			EntryId           string `xml:"entryId,attr"`
			ID                string `xml:"id,attr"`
			Name              string `xml:"name,attr"`
			Publications      struct {
				Text        string `xml:",chardata"`
				Publication []struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Name string `xml:"name,attr"`
				} `xml:"publication"`
			} `xml:"publications"`
			Categories struct {
				Text     string `xml:",chardata"`
				Category []struct {
					Text     string `xml:",chardata"`
					Primary  string `xml:"primary,attr"`
					EntryId  string `xml:"entryId,attr"`
					ID       string `xml:"id,attr"`
					Name     string `xml:"name,attr"`
					Rules    string `xml:"rules"`
					Profiles string `xml:"profiles"`
				} `xml:"category"`
			} `xml:"categories"`
			Forces     string `xml:"forces"`
			Selections struct {
				Text      string `xml:",chardata"`
				Selection []struct {
					Text          string `xml:",chardata"`
					Number        string `xml:"number,attr"`
					Type          string `xml:"type,attr"`
					EntryId       string `xml:"entryId,attr"`
					ID            string `xml:"id,attr"`
					Name          string `xml:"name,attr"`
					PublicationId string `xml:"publicationId,attr"`
					Page          string `xml:"page,attr"`
					Costs         struct {
						Text string `xml:",chardata"`
						Cost []struct {
							Text   string `xml:",chardata"`
							TypeId string `xml:"typeId,attr"`
							Value  string `xml:"value,attr"`
							Name   string `xml:"name,attr"`
						} `xml:"cost"`
					} `xml:"costs"`
					Categories struct {
						Text     string `xml:",chardata"`
						Category []struct {
							Text     string `xml:",chardata"`
							Primary  string `xml:"primary,attr"`
							EntryId  string `xml:"entryId,attr"`
							ID       string `xml:"id,attr"`
							Name     string `xml:"name,attr"`
							Rules    string `xml:"rules"`
							Profiles string `xml:"profiles"`
						} `xml:"category"`
					} `xml:"categories"`
					Selections struct {
						Text      string `xml:",chardata"`
						Selection []struct {
							Text         string `xml:",chardata"`
							Number       string `xml:"number,attr"`
							Type         string `xml:"type,attr"`
							EntryId      string `xml:"entryId,attr"`
							EntryGroupId string `xml:"entryGroupId,attr"`
							ID           string `xml:"id,attr"`
							Name         string `xml:"name,attr"`
							Page         string `xml:"page,attr"`
							Costs        struct {
								Text string `xml:",chardata"`
								Cost []struct {
									Text   string `xml:",chardata"`
									TypeId string `xml:"typeId,attr"`
									Value  string `xml:"value,attr"`
									Name   string `xml:"name,attr"`
								} `xml:"cost"`
							} `xml:"costs"`
							Categories struct {
								Text     string `xml:",chardata"`
								Category []struct {
									Text     string `xml:",chardata"`
									Primary  string `xml:"primary,attr"`
									EntryId  string `xml:"entryId,attr"`
									ID       string `xml:"id,attr"`
									Name     string `xml:"name,attr"`
									Rules    string `xml:"rules"`
									Profiles string `xml:"profiles"`
								} `xml:"category"`
							} `xml:"categories"`
							Selections struct {
								Text      string `xml:",chardata"`
								Selection []struct {
									Text         string `xml:",chardata"`
									Number       string `xml:"number,attr"`
									Type         string `xml:"type,attr"`
									EntryId      string `xml:"entryId,attr"`
									ID           string `xml:"id,attr"`
									Name         string `xml:"name,attr"`
									EntryGroupId string `xml:"entryGroupId,attr"`
									Costs        struct {
										Text string `xml:",chardata"`
										Cost []struct {
											Text   string `xml:",chardata"`
											TypeId string `xml:"typeId,attr"`
											Value  string `xml:"value,attr"`
											Name   string `xml:"name,attr"`
										} `xml:"cost"`
									} `xml:"costs"`
									Categories string `xml:"categories"`
									Selections string `xml:"selections"`
									Rules      string `xml:"rules"`
									Profiles   struct {
										Text    string `xml:",chardata"`
										Profile struct {
											Text            string `xml:",chardata"`
											TypeId          string `xml:"typeId,attr"`
											TypeName        string `xml:"typeName,attr"`
											Hidden          string `xml:"hidden,attr"`
											PublicationId   string `xml:"publicationId,attr"`
											Page            string `xml:"page,attr"`
											ID              string `xml:"id,attr"`
											Name            string `xml:"name,attr"`
											Characteristics struct {
												Text           string `xml:",chardata"`
												Characteristic []struct {
													Text   string `xml:",chardata"`
													TypeId string `xml:"typeId,attr"`
													Name   string `xml:"name,attr"`
												} `xml:"characteristic"`
											} `xml:"characteristics"`
											Modifiers      string `xml:"modifiers"`
											ModifierGroups string `xml:"modifierGroups"`
										} `xml:"profile"`
									} `xml:"profiles"`
								} `xml:"selection"`
							} `xml:"selections"`
							Rules struct {
								Text string `xml:",chardata"`
								Rule struct {
									Text           string `xml:",chardata"`
									Hidden         string `xml:"hidden,attr"`
									ID             string `xml:"id,attr"`
									Name           string `xml:"name,attr"`
									Description    string `xml:"description"`
									Modifiers      string `xml:"modifiers"`
									ModifierGroups string `xml:"modifierGroups"`
								} `xml:"rule"`
							} `xml:"rules"`
							Profiles struct {
								Text    string `xml:",chardata"`
								Profile struct {
									Text            string `xml:",chardata"`
									TypeId          string `xml:"typeId,attr"`
									TypeName        string `xml:"typeName,attr"`
									Hidden          string `xml:"hidden,attr"`
									ID              string `xml:"id,attr"`
									Name            string `xml:"name,attr"`
									Characteristics struct {
										Text           string `xml:",chardata"`
										Characteristic []struct {
											Text   string `xml:",chardata"`
											TypeId string `xml:"typeId,attr"`
											Name   string `xml:"name,attr"`
										} `xml:"characteristic"`
									} `xml:"characteristics"`
									Modifiers      string `xml:"modifiers"`
									ModifierGroups string `xml:"modifierGroups"`
								} `xml:"profile"`
							} `xml:"profiles"`
						} `xml:"selection"`
					} `xml:"selections"`
					Rules struct {
						Text string `xml:",chardata"`
						Rule []struct {
							Text           string `xml:",chardata"`
							Hidden         string `xml:"hidden,attr"`
							PublicationId  string `xml:"publicationId,attr"`
							ID             string `xml:"id,attr"`
							Name           string `xml:"name,attr"`
							Page           string `xml:"page,attr"`
							Description    string `xml:"description"`
							Modifiers      string `xml:"modifiers"`
							ModifierGroups string `xml:"modifierGroups"`
						} `xml:"rule"`
					} `xml:"rules"`
					Profiles struct {
						Text    string `xml:",chardata"`
						Profile []struct {
							Text            string `xml:",chardata"`
							TypeId          string `xml:"typeId,attr"`
							TypeName        string `xml:"typeName,attr"`
							Hidden          string `xml:"hidden,attr"`
							ID              string `xml:"id,attr"`
							Name            string `xml:"name,attr"`
							Characteristics struct {
								Text           string `xml:",chardata"`
								Characteristic []struct {
									Text   string `xml:",chardata"`
									TypeId string `xml:"typeId,attr"`
									Name   string `xml:"name,attr"`
								} `xml:"characteristic"`
							} `xml:"characteristics"`
							Modifiers      string `xml:"modifiers"`
							ModifierGroups string `xml:"modifierGroups"`
						} `xml:"profile"`
					} `xml:"profiles"`
				} `xml:"selection"`
			} `xml:"selections"`
			Rules    string `xml:"rules"`
			Profiles string `xml:"profiles"`
		} `xml:"force"`
	} `xml:"forces"`
	Tags string `xml:"tags"`
}

var (
	ErrCouldNotOpenROSZFile            = errors.New("could not open rosz file at path specified")
	ErrCouldNotOpenRosterFileInArchive = errors.New("could not open roster file in archive")
	ErrCouldNotReadRosterBytes         = errors.New("could not read roster bytes")
	ErrCouldNotParseRosterData         = errors.New("could not parse roster data")
)

// ReadRoszFile reads the rosz file in the supplied path, unzips it, and parses the roster inside
func ReadRoszFile(roszPath string) (Roster, error) {
	var (
		r   Roster
		err error
	)

	// open the roster archive file
	archive, err := zip.OpenReader(roszPath)
	if err != nil {
		return r, ErrCouldNotOpenROSZFile
	}
	defer archive.Close()

	if r, err = ReadRoszArchive(archive); err != nil {
		return r, err
	}

	return r, nil
}

// ReadRoszArchive parses the ros file in the supplied archive and returns a Roster
func ReadRoszArchive(archive *zip.ReadCloser) (Roster, error) {
	var r Roster

	for _, f := range archive.File {
		println(f.Name)
		fileInArchive, err := f.Open()
		if err != nil {
			return r, ErrCouldNotOpenRosterFileInArchive
		}
		defer fileInArchive.Close()

		strBuf := bytes.NewBufferString("")
		_, err = io.Copy(strBuf, fileInArchive)
		if err != nil {
			return r, ErrCouldNotReadRosterBytes
		}

		rosterBytes := strBuf.Bytes()

		err = xml.Unmarshal(rosterBytes, &r)
		if err != nil {
			return r, ErrCouldNotParseRosterData
		}
	}

	return r, nil
}
