// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Engine string

// Enum values for Engine
const (
	EngineStandard Engine = "standard"
	EngineNeural   Engine = "neural"
)

// Values returns all known values for Engine. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Engine) Values() []Engine {
	return []Engine{
		"standard",
		"neural",
	}
}

type Gender string

// Values returns all known values for Gender. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Gender) Values() []Gender {
	return []Gender{
		"Female",
		"Male",
	}
}

type LanguageCode string

// Values returns all known values for LanguageCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LanguageCode) Values() []LanguageCode {
	return []LanguageCode{
		"arb",
		"cmn-CN",
		"cy-GB",
		"da-DK",
		"de-DE",
		"en-AU",
		"en-GB",
		"en-GB-WLS",
		"en-IN",
		"en-US",
		"es-ES",
		"es-MX",
		"es-US",
		"fr-CA",
		"fr-FR",
		"is-IS",
		"it-IT",
		"ja-JP",
		"hi-IN",
		"ko-KR",
		"nb-NO",
		"nl-NL",
		"pl-PL",
		"pt-BR",
		"pt-PT",
		"ro-RO",
		"ru-RU",
		"sv-SE",
		"tr-TR",
	}
}

type OutputFormat string

// Enum values for OutputFormat
const (
	OutputFormatJson      OutputFormat = "json"
	OutputFormatMp3       OutputFormat = "mp3"
	OutputFormatOggVorbis OutputFormat = "ogg_vorbis"
	OutputFormatPcm       OutputFormat = "pcm"
)

// Values returns all known values for OutputFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OutputFormat) Values() []OutputFormat {
	return []OutputFormat{
		"json",
		"mp3",
		"ogg_vorbis",
		"pcm",
	}
}

type SpeechMarkType string

// Enum values for SpeechMarkType
const (
	SpeechMarkTypeSentence SpeechMarkType = "sentence"
	SpeechMarkTypeSsml     SpeechMarkType = "ssml"
	SpeechMarkTypeViseme   SpeechMarkType = "viseme"
	SpeechMarkTypeWord     SpeechMarkType = "word"
)

// Values returns all known values for SpeechMarkType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SpeechMarkType) Values() []SpeechMarkType {
	return []SpeechMarkType{
		"sentence",
		"ssml",
		"viseme",
		"word",
	}
}

type TaskStatus string

// Enum values for TaskStatus
const (
	TaskStatusScheduled  TaskStatus = "scheduled"
	TaskStatusInProgress TaskStatus = "inProgress"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusFailed     TaskStatus = "failed"
)

// Values returns all known values for TaskStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TaskStatus) Values() []TaskStatus {
	return []TaskStatus{
		"scheduled",
		"inProgress",
		"completed",
		"failed",
	}
}

type TextType string

// Enum values for TextType
const (
	TextTypeSsml TextType = "ssml"
	TextTypeText TextType = "text"
)

// Values returns all known values for TextType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TextType) Values() []TextType {
	return []TextType{
		"ssml",
		"text",
	}
}

type VoiceId string

// Values returns all known values for VoiceId. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (VoiceId) Values() []VoiceId {
	return []VoiceId{
		"Aditi",
		"Amy",
		"Astrid",
		"Bianca",
		"Brian",
		"Camila",
		"Carla",
		"Carmen",
		"Celine",
		"Chantal",
		"Conchita",
		"Cristiano",
		"Dora",
		"Emma",
		"Enrique",
		"Ewa",
		"Filiz",
		"Geraint",
		"Giorgio",
		"Gwyneth",
		"Hans",
		"Ines",
		"Ivy",
		"Jacek",
		"Jan",
		"Joanna",
		"Joey",
		"Justin",
		"Karl",
		"Kendra",
		"Kevin",
		"Kimberly",
		"Lea",
		"Liv",
		"Lotte",
		"Lucia",
		"Lupe",
		"Mads",
		"Maja",
		"Marlene",
		"Mathieu",
		"Matthew",
		"Maxim",
		"Mia",
		"Miguel",
		"Mizuki",
		"Naja",
		"Nicole",
		"Olivia",
		"Penelope",
		"Raveena",
		"Ricardo",
		"Ruben",
		"Russell",
		"Salli",
		"Seoyeon",
		"Takumi",
		"Tatyana",
		"Vicki",
		"Vitoria",
		"Zeina",
		"Zhiyu",
	}
}
