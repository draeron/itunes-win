package itunes

type ITUserPlaylistSpecialKind int

const (
	ITUserPlaylistSpecialKindNone ITUserPlaylistSpecialKind = iota
	ITUserPlaylistSpecialKindPurchasedMusic
	ITUserPlaylistSpecialKindPartyShuffle
	ITUserPlaylistSpecialKindPodcasts
	ITUserPlaylistSpecialKindFolder
	ITUserPlaylistSpecialKindVideos
	ITUserPlaylistSpecialKindMusic
	ITUserPlaylistSpecialKindMovies
	ITUserPlaylistSpecialKindTVShows
	ITUserPlaylistSpecialKindAudiobooks
)

func (i ITUserPlaylistSpecialKind) String() string {
	strs := []string{
		"ITUserPlaylistSpecialKindNone",
		"ITUserPlaylistSpecialKindPurchasedMusic",
		"ITUserPlaylistSpecialKindPartyShuffle",
		"ITUserPlaylistSpecialKindPodcasts",
		"ITUserPlaylistSpecialKindFolder",
		"ITUserPlaylistSpecialKindVideos",
		"ITUserPlaylistSpecialKindMusic",
		"ITUserPlaylistSpecialKindMovies",
		"ITUserPlaylistSpecialKindTVShows",
		"ITUserPlaylistSpecialKindAudiobooks",
	}
	return strs[i]
}

type ITRatingKind int

const (
	ITRatingKindUser ITRatingKind = iota
	ITRatingKindComputed
)

func (i ITRatingKind) String() string {
	strs := []string{
		"ITRatingKindUser",
		"ITRatingKindComputed",
	}
	return strs[i]
}

type ITTrackKind int

const (
	ITTrackKindUnknown ITTrackKind = iota
	ITTrackKindFile
	ITTrackKindCD
	ITTrackKindURL
	ITTrackKindDevice
	ITTrackKindSharedLibrary
)

func (i ITTrackKind) String() string {
	strs := []string{
		"ITTrackKindUnknown",
		"ITTrackKindFile",
		"ITTrackKindCD",
		"ITTrackKindURL",
		"ITTrackKindDevice",
		"ITTrackKindSharedLibrary",
	}
	return strs[i]
}

type ITVideoKind int

const (
	ITVideoKindNone ITVideoKind = iota
	ITVideoKindMovie
	ITVideoKindMusicVideo
	ITVideoKindTVShow
)

type ITPlaylistKind int

const (
	ITPlaylistKindUnknown ITPlaylistKind = iota
	ITPlaylistKindLibrary
	ITPlaylistKindUser
	ITPlaylistKindCD
	ITPlaylistKindDevice
	ITPlaylistKindRadioTuner
)

func (i ITPlaylistKind) String() string {
	strs := []string{
		"ITPlaylistKindUnknown",
		"ITPlaylistKindLibrary",
		"ITPlaylistKindUser",
		"ITPlaylistKindCD",
		"ITPlaylistKindDevice",
		"ITPlaylistKindRadioTuner",
	}
	return strs[i]
}

type ITSourceKind int

const (
	ITSourceKindUnknown ITSourceKind = iota
	ITSourceKindLibrary
	ITSourceKindIPod
	ITSourceKindAudioCD
	ITSourceKindMP3CD
	ITSourceKindDevice
	ITSourceKindRadioTuner
	ITSourceKindSharedLibrary
)

func (i ITSourceKind) String() string {
	strs := []string{
		"ITSourceKindUnknown",
		"ITSourceKindLibrary",
		"ITSourceKindIPod",
		"ITSourceKindAudioCD",
		"ITSourceKindMP3CD",
		"ITSourceKindDevice",
		"ITSourceKindRadioTuner",
		"ITSourceKindSharedLibrary",
	}
	return strs[i]
}
