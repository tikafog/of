// Code generated by entc, DO NOT EDIT.

package media

import (
	"github.com/google/uuid"
	"github.com/tikafog/of/dbc/media/announce"
	"github.com/tikafog/of/dbc/media/channel"
	"github.com/tikafog/of/dbc/media/discovery"
	"github.com/tikafog/of/dbc/media/informationv1"
	"github.com/tikafog/of/dbc/media/page"
	"github.com/tikafog/of/dbc/media/schema"
	"github.com/tikafog/of/dbc/media/toplist"
	"github.com/tikafog/of/dbc/media/version"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	announceMixin := schema.Announce{}.Mixin()
	announceMixinFields0 := announceMixin[0].Fields()
	_ = announceMixinFields0
	announceMixinFields1 := announceMixin[1].Fields()
	_ = announceMixinFields1
	announceFields := schema.Announce{}.Fields()
	_ = announceFields
	// announceDescCreatedUnix is the schema descriptor for created_unix field.
	announceDescCreatedUnix := announceMixinFields1[0].Descriptor()
	// announce.DefaultCreatedUnix holds the default value on creation for the created_unix field.
	announce.DefaultCreatedUnix = announceDescCreatedUnix.Default.(int64)
	// announceDescUpdatedUnix is the schema descriptor for updated_unix field.
	announceDescUpdatedUnix := announceMixinFields1[1].Descriptor()
	// announce.DefaultUpdatedUnix holds the default value on creation for the updated_unix field.
	announce.DefaultUpdatedUnix = announceDescUpdatedUnix.Default.(int64)
	// announceDescAnnounceNo is the schema descriptor for announce_no field.
	announceDescAnnounceNo := announceFields[0].Descriptor()
	// announce.DefaultAnnounceNo holds the default value on creation for the announce_no field.
	announce.DefaultAnnounceNo = announceDescAnnounceNo.Default.(string)
	// announceDescTitle is the schema descriptor for title field.
	announceDescTitle := announceFields[1].Descriptor()
	// announce.DefaultTitle holds the default value on creation for the title field.
	announce.DefaultTitle = announceDescTitle.Default.(string)
	// announceDescContent is the schema descriptor for content field.
	announceDescContent := announceFields[3].Descriptor()
	// announce.DefaultContent holds the default value on creation for the content field.
	announce.DefaultContent = announceDescContent.Default.(string)
	// announceDescLink is the schema descriptor for link field.
	announceDescLink := announceFields[4].Descriptor()
	// announce.DefaultLink holds the default value on creation for the link field.
	announce.DefaultLink = announceDescLink.Default.(string)
	// announceDescID is the schema descriptor for id field.
	announceDescID := announceMixinFields0[0].Descriptor()
	// announce.DefaultID holds the default value on creation for the id field.
	announce.DefaultID = announceDescID.Default.(func() uuid.UUID)
	channelMixin := schema.Channel{}.Mixin()
	channelMixinFields0 := channelMixin[0].Fields()
	_ = channelMixinFields0
	channelMixinFields1 := channelMixin[1].Fields()
	_ = channelMixinFields1
	channelFields := schema.Channel{}.Fields()
	_ = channelFields
	// channelDescCreatedUnix is the schema descriptor for created_unix field.
	channelDescCreatedUnix := channelMixinFields1[0].Descriptor()
	// channel.DefaultCreatedUnix holds the default value on creation for the created_unix field.
	channel.DefaultCreatedUnix = channelDescCreatedUnix.Default.(int64)
	// channelDescUpdatedUnix is the schema descriptor for updated_unix field.
	channelDescUpdatedUnix := channelMixinFields1[1].Descriptor()
	// channel.DefaultUpdatedUnix holds the default value on creation for the updated_unix field.
	channel.DefaultUpdatedUnix = channelDescUpdatedUnix.Default.(int64)
	// channelDescID is the schema descriptor for id field.
	channelDescID := channelMixinFields0[0].Descriptor()
	// channel.DefaultID holds the default value on creation for the id field.
	channel.DefaultID = channelDescID.Default.(func() uuid.UUID)
	discoveryMixin := schema.Discovery{}.Mixin()
	discoveryMixinFields0 := discoveryMixin[0].Fields()
	_ = discoveryMixinFields0
	discoveryMixinFields1 := discoveryMixin[1].Fields()
	_ = discoveryMixinFields1
	discoveryFields := schema.Discovery{}.Fields()
	_ = discoveryFields
	// discoveryDescCreatedUnix is the schema descriptor for created_unix field.
	discoveryDescCreatedUnix := discoveryMixinFields1[0].Descriptor()
	// discovery.DefaultCreatedUnix holds the default value on creation for the created_unix field.
	discovery.DefaultCreatedUnix = discoveryDescCreatedUnix.Default.(int64)
	// discoveryDescUpdatedUnix is the schema descriptor for updated_unix field.
	discoveryDescUpdatedUnix := discoveryMixinFields1[1].Descriptor()
	// discovery.DefaultUpdatedUnix holds the default value on creation for the updated_unix field.
	discovery.DefaultUpdatedUnix = discoveryDescUpdatedUnix.Default.(int64)
	// discoveryDescDate is the schema descriptor for date field.
	discoveryDescDate := discoveryFields[0].Descriptor()
	// discovery.DefaultDate holds the default value on creation for the date field.
	discovery.DefaultDate = discoveryDescDate.Default.(string)
	// discoveryDescRid is the schema descriptor for rid field.
	discoveryDescRid := discoveryFields[1].Descriptor()
	// discovery.DefaultRid holds the default value on creation for the rid field.
	discovery.DefaultRid = discoveryDescRid.Default.(string)
	// discoveryDescTitle is the schema descriptor for title field.
	discoveryDescTitle := discoveryFields[2].Descriptor()
	// discovery.DefaultTitle holds the default value on creation for the title field.
	discovery.DefaultTitle = discoveryDescTitle.Default.(string)
	// discoveryDescDetail is the schema descriptor for detail field.
	discoveryDescDetail := discoveryFields[3].Descriptor()
	// discovery.DefaultDetail holds the default value on creation for the detail field.
	discovery.DefaultDetail = discoveryDescDetail.Default.(string)
	// discoveryDescID is the schema descriptor for id field.
	discoveryDescID := discoveryMixinFields0[0].Descriptor()
	// discovery.DefaultID holds the default value on creation for the id field.
	discovery.DefaultID = discoveryDescID.Default.(func() uuid.UUID)
	informationv1Mixin := schema.InformationV1{}.Mixin()
	informationv1MixinFields0 := informationv1Mixin[0].Fields()
	_ = informationv1MixinFields0
	informationv1MixinFields1 := informationv1Mixin[1].Fields()
	_ = informationv1MixinFields1
	informationv1Fields := schema.InformationV1{}.Fields()
	_ = informationv1Fields
	// informationv1DescCreatedUnix is the schema descriptor for created_unix field.
	informationv1DescCreatedUnix := informationv1MixinFields1[0].Descriptor()
	// informationv1.DefaultCreatedUnix holds the default value on creation for the created_unix field.
	informationv1.DefaultCreatedUnix = informationv1DescCreatedUnix.Default.(int64)
	// informationv1DescUpdatedUnix is the schema descriptor for updated_unix field.
	informationv1DescUpdatedUnix := informationv1MixinFields1[1].Descriptor()
	// informationv1.DefaultUpdatedUnix holds the default value on creation for the updated_unix field.
	informationv1.DefaultUpdatedUnix = informationv1DescUpdatedUnix.Default.(int64)
	// informationv1DescRoot is the schema descriptor for root field.
	informationv1DescRoot := informationv1Fields[1].Descriptor()
	// informationv1.DefaultRoot holds the default value on creation for the root field.
	informationv1.DefaultRoot = informationv1DescRoot.Default.(string)
	// informationv1DescThumb is the schema descriptor for thumb field.
	informationv1DescThumb := informationv1Fields[2].Descriptor()
	// informationv1.DefaultThumb holds the default value on creation for the thumb field.
	informationv1.DefaultThumb = informationv1DescThumb.Default.(string)
	// informationv1DescThumbPath is the schema descriptor for thumb_path field.
	informationv1DescThumbPath := informationv1Fields[3].Descriptor()
	// informationv1.DefaultThumbPath holds the default value on creation for the thumb_path field.
	informationv1.DefaultThumbPath = informationv1DescThumbPath.Default.(string)
	// informationv1DescPoster is the schema descriptor for poster field.
	informationv1DescPoster := informationv1Fields[4].Descriptor()
	// informationv1.DefaultPoster holds the default value on creation for the poster field.
	informationv1.DefaultPoster = informationv1DescPoster.Default.(string)
	// informationv1DescPosterPath is the schema descriptor for poster_path field.
	informationv1DescPosterPath := informationv1Fields[5].Descriptor()
	// informationv1.DefaultPosterPath holds the default value on creation for the poster_path field.
	informationv1.DefaultPosterPath = informationv1DescPosterPath.Default.(string)
	// informationv1DescMedia is the schema descriptor for media field.
	informationv1DescMedia := informationv1Fields[6].Descriptor()
	// informationv1.DefaultMedia holds the default value on creation for the media field.
	informationv1.DefaultMedia = informationv1DescMedia.Default.(string)
	// informationv1DescMediaPath is the schema descriptor for media_path field.
	informationv1DescMediaPath := informationv1Fields[7].Descriptor()
	// informationv1.DefaultMediaPath holds the default value on creation for the media_path field.
	informationv1.DefaultMediaPath = informationv1DescMediaPath.Default.(string)
	// informationv1DescMediaIndex is the schema descriptor for media_index field.
	informationv1DescMediaIndex := informationv1Fields[8].Descriptor()
	// informationv1.DefaultMediaIndex holds the default value on creation for the media_index field.
	informationv1.DefaultMediaIndex = informationv1DescMediaIndex.Default.(string)
	// informationv1DescFrames is the schema descriptor for frames field.
	informationv1DescFrames := informationv1Fields[9].Descriptor()
	// informationv1.DefaultFrames holds the default value on creation for the frames field.
	informationv1.DefaultFrames = informationv1DescFrames.Default.(string)
	// informationv1DescFramesPath is the schema descriptor for frames_path field.
	informationv1DescFramesPath := informationv1Fields[10].Descriptor()
	// informationv1.DefaultFramesPath holds the default value on creation for the frames_path field.
	informationv1.DefaultFramesPath = informationv1DescFramesPath.Default.(string)
	// informationv1DescTitle is the schema descriptor for title field.
	informationv1DescTitle := informationv1Fields[12].Descriptor()
	// informationv1.DefaultTitle holds the default value on creation for the title field.
	informationv1.DefaultTitle = informationv1DescTitle.Default.(string)
	// informationv1DescVideoNo is the schema descriptor for video_no field.
	informationv1DescVideoNo := informationv1Fields[13].Descriptor()
	// informationv1.DefaultVideoNo holds the default value on creation for the video_no field.
	informationv1.DefaultVideoNo = informationv1DescVideoNo.Default.(string)
	// informationv1DescIntro is the schema descriptor for intro field.
	informationv1DescIntro := informationv1Fields[14].Descriptor()
	// informationv1.DefaultIntro holds the default value on creation for the intro field.
	informationv1.DefaultIntro = informationv1DescIntro.Default.(string)
	// informationv1DescDirector is the schema descriptor for director field.
	informationv1DescDirector := informationv1Fields[17].Descriptor()
	// informationv1.DefaultDirector holds the default value on creation for the director field.
	informationv1.DefaultDirector = informationv1DescDirector.Default.(string)
	// informationv1DescSystematics is the schema descriptor for systematics field.
	informationv1DescSystematics := informationv1Fields[18].Descriptor()
	// informationv1.DefaultSystematics holds the default value on creation for the systematics field.
	informationv1.DefaultSystematics = informationv1DescSystematics.Default.(string)
	// informationv1DescProducer is the schema descriptor for producer field.
	informationv1DescProducer := informationv1Fields[19].Descriptor()
	// informationv1.DefaultProducer holds the default value on creation for the producer field.
	informationv1.DefaultProducer = informationv1DescProducer.Default.(string)
	// informationv1DescPublisher is the schema descriptor for publisher field.
	informationv1DescPublisher := informationv1Fields[20].Descriptor()
	// informationv1.DefaultPublisher holds the default value on creation for the publisher field.
	informationv1.DefaultPublisher = informationv1DescPublisher.Default.(string)
	// informationv1DescSortType is the schema descriptor for sort_type field.
	informationv1DescSortType := informationv1Fields[21].Descriptor()
	// informationv1.DefaultSortType holds the default value on creation for the sort_type field.
	informationv1.DefaultSortType = informationv1DescSortType.Default.(string)
	// informationv1DescCaption is the schema descriptor for caption field.
	informationv1DescCaption := informationv1Fields[22].Descriptor()
	// informationv1.DefaultCaption holds the default value on creation for the caption field.
	informationv1.DefaultCaption = informationv1DescCaption.Default.(string)
	// informationv1DescGroup is the schema descriptor for group field.
	informationv1DescGroup := informationv1Fields[23].Descriptor()
	// informationv1.DefaultGroup holds the default value on creation for the group field.
	informationv1.DefaultGroup = informationv1DescGroup.Default.(string)
	// informationv1DescIndex is the schema descriptor for index field.
	informationv1DescIndex := informationv1Fields[24].Descriptor()
	// informationv1.DefaultIndex holds the default value on creation for the index field.
	informationv1.DefaultIndex = informationv1DescIndex.Default.(string)
	// informationv1DescReleaseDate is the schema descriptor for release_date field.
	informationv1DescReleaseDate := informationv1Fields[25].Descriptor()
	// informationv1.DefaultReleaseDate holds the default value on creation for the release_date field.
	informationv1.DefaultReleaseDate = informationv1DescReleaseDate.Default.(string)
	// informationv1DescFormat is the schema descriptor for format field.
	informationv1DescFormat := informationv1Fields[26].Descriptor()
	// informationv1.DefaultFormat holds the default value on creation for the format field.
	informationv1.DefaultFormat = informationv1DescFormat.Default.(string)
	// informationv1DescSeries is the schema descriptor for series field.
	informationv1DescSeries := informationv1Fields[27].Descriptor()
	// informationv1.DefaultSeries holds the default value on creation for the series field.
	informationv1.DefaultSeries = informationv1DescSeries.Default.(string)
	// informationv1DescLength is the schema descriptor for length field.
	informationv1DescLength := informationv1Fields[29].Descriptor()
	// informationv1.DefaultLength holds the default value on creation for the length field.
	informationv1.DefaultLength = informationv1DescLength.Default.(string)
	// informationv1DescUncensored is the schema descriptor for uncensored field.
	informationv1DescUncensored := informationv1Fields[31].Descriptor()
	// informationv1.DefaultUncensored holds the default value on creation for the uncensored field.
	informationv1.DefaultUncensored = informationv1DescUncensored.Default.(string)
	// informationv1DescSeason is the schema descriptor for season field.
	informationv1DescSeason := informationv1Fields[32].Descriptor()
	// informationv1.DefaultSeason holds the default value on creation for the season field.
	informationv1.DefaultSeason = informationv1DescSeason.Default.(string)
	// informationv1DescTotalEpisode is the schema descriptor for total_episode field.
	informationv1DescTotalEpisode := informationv1Fields[33].Descriptor()
	// informationv1.DefaultTotalEpisode holds the default value on creation for the total_episode field.
	informationv1.DefaultTotalEpisode = informationv1DescTotalEpisode.Default.(string)
	// informationv1DescEpisode is the schema descriptor for episode field.
	informationv1DescEpisode := informationv1Fields[34].Descriptor()
	// informationv1.DefaultEpisode holds the default value on creation for the episode field.
	informationv1.DefaultEpisode = informationv1DescEpisode.Default.(string)
	// informationv1DescLanguage is the schema descriptor for language field.
	informationv1DescLanguage := informationv1Fields[35].Descriptor()
	// informationv1.DefaultLanguage holds the default value on creation for the language field.
	informationv1.DefaultLanguage = informationv1DescLanguage.Default.(string)
	// informationv1DescSharpness is the schema descriptor for sharpness field.
	informationv1DescSharpness := informationv1Fields[36].Descriptor()
	// informationv1.DefaultSharpness holds the default value on creation for the sharpness field.
	informationv1.DefaultSharpness = informationv1DescSharpness.Default.(string)
	// informationv1DescWatermark is the schema descriptor for watermark field.
	informationv1DescWatermark := informationv1Fields[37].Descriptor()
	// informationv1.DefaultWatermark holds the default value on creation for the watermark field.
	informationv1.DefaultWatermark = informationv1DescWatermark.Default.(bool)
	// informationv1DescTotalBlocks is the schema descriptor for total_blocks field.
	informationv1DescTotalBlocks := informationv1Fields[40].Descriptor()
	// informationv1.DefaultTotalBlocks holds the default value on creation for the total_blocks field.
	informationv1.DefaultTotalBlocks = informationv1DescTotalBlocks.Default.(int)
	// informationv1DescID is the schema descriptor for id field.
	informationv1DescID := informationv1MixinFields0[0].Descriptor()
	// informationv1.DefaultID holds the default value on creation for the id field.
	informationv1.DefaultID = informationv1DescID.Default.(func() uuid.UUID)
	pageMixin := schema.Page{}.Mixin()
	pageMixinFields0 := pageMixin[0].Fields()
	_ = pageMixinFields0
	pageMixinFields1 := pageMixin[1].Fields()
	_ = pageMixinFields1
	pageFields := schema.Page{}.Fields()
	_ = pageFields
	// pageDescCreatedUnix is the schema descriptor for created_unix field.
	pageDescCreatedUnix := pageMixinFields1[0].Descriptor()
	// page.DefaultCreatedUnix holds the default value on creation for the created_unix field.
	page.DefaultCreatedUnix = pageDescCreatedUnix.Default.(int64)
	// pageDescUpdatedUnix is the schema descriptor for updated_unix field.
	pageDescUpdatedUnix := pageMixinFields1[1].Descriptor()
	// page.DefaultUpdatedUnix holds the default value on creation for the updated_unix field.
	page.DefaultUpdatedUnix = pageDescUpdatedUnix.Default.(int64)
	// pageDescTitle is the schema descriptor for title field.
	pageDescTitle := pageFields[1].Descriptor()
	// page.DefaultTitle holds the default value on creation for the title field.
	page.DefaultTitle = pageDescTitle.Default.(string)
	// pageDescFeaturedIndex is the schema descriptor for featured_index field.
	pageDescFeaturedIndex := pageFields[2].Descriptor()
	// page.DefaultFeaturedIndex holds the default value on creation for the featured_index field.
	page.DefaultFeaturedIndex = pageDescFeaturedIndex.Default.(int8)
	// pageDescFeaturedContent is the schema descriptor for featured_content field.
	pageDescFeaturedContent := pageFields[3].Descriptor()
	// page.DefaultFeaturedContent holds the default value on creation for the featured_content field.
	page.DefaultFeaturedContent = pageDescFeaturedContent.Default.(string)
	// pageDescRecommend is the schema descriptor for recommend field.
	pageDescRecommend := pageFields[4].Descriptor()
	// page.DefaultRecommend holds the default value on creation for the recommend field.
	page.DefaultRecommend = pageDescRecommend.Default.(int64)
	// pageDescID is the schema descriptor for id field.
	pageDescID := pageMixinFields0[0].Descriptor()
	// page.DefaultID holds the default value on creation for the id field.
	page.DefaultID = pageDescID.Default.(func() uuid.UUID)
	toplistMixin := schema.TopList{}.Mixin()
	toplistMixinFields0 := toplistMixin[0].Fields()
	_ = toplistMixinFields0
	toplistMixinFields1 := toplistMixin[1].Fields()
	_ = toplistMixinFields1
	toplistFields := schema.TopList{}.Fields()
	_ = toplistFields
	// toplistDescCreatedUnix is the schema descriptor for created_unix field.
	toplistDescCreatedUnix := toplistMixinFields1[0].Descriptor()
	// toplist.DefaultCreatedUnix holds the default value on creation for the created_unix field.
	toplist.DefaultCreatedUnix = toplistDescCreatedUnix.Default.(int64)
	// toplistDescUpdatedUnix is the schema descriptor for updated_unix field.
	toplistDescUpdatedUnix := toplistMixinFields1[1].Descriptor()
	// toplist.DefaultUpdatedUnix holds the default value on creation for the updated_unix field.
	toplist.DefaultUpdatedUnix = toplistDescUpdatedUnix.Default.(int64)
	// toplistDescTitle is the schema descriptor for title field.
	toplistDescTitle := toplistFields[2].Descriptor()
	// toplist.DefaultTitle holds the default value on creation for the title field.
	toplist.DefaultTitle = toplistDescTitle.Default.(string)
	// toplistDescIntro is the schema descriptor for intro field.
	toplistDescIntro := toplistFields[3].Descriptor()
	// toplist.DefaultIntro holds the default value on creation for the intro field.
	toplist.DefaultIntro = toplistDescIntro.Default.(string)
	// toplistDescID is the schema descriptor for id field.
	toplistDescID := toplistMixinFields0[0].Descriptor()
	// toplist.DefaultID holds the default value on creation for the id field.
	toplist.DefaultID = toplistDescID.Default.(func() uuid.UUID)
	versionFields := schema.Version{}.Fields()
	_ = versionFields
	// versionDescCurrent is the schema descriptor for Current field.
	versionDescCurrent := versionFields[0].Descriptor()
	// version.DefaultCurrent holds the default value on creation for the Current field.
	version.DefaultCurrent = versionDescCurrent.Default.(int)
	// version.CurrentValidator is a validator for the "Current" field. It is called by the builders before save.
	version.CurrentValidator = versionDescCurrent.Validators[0].(func(int) error)
	// versionDescLast is the schema descriptor for Last field.
	versionDescLast := versionFields[1].Descriptor()
	// version.DefaultLast holds the default value on creation for the Last field.
	version.DefaultLast = versionDescLast.Default.(int)
	// version.LastValidator is a validator for the "Last" field. It is called by the builders before save.
	version.LastValidator = versionDescLast.Validators[0].(func(int) error)
}