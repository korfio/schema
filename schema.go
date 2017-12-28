package schema

import "time"

type Text string
type Boolean bool
type Number struct {
	int
	float64
}
type Time time.Time
type DateTime time.Time
type Date time.Time

type LakeBodyOfWater struct {
	BodyOfWater
}

type TVClip struct {
	PartOfTVSeries TVSeries `json:"partOfTVSeries,omitempty"`
	Clip
}

type DrugPregnancyCategory struct {
	MedicalEnumeration
}

type GovernmentService struct {
	ServiceOperator Organization `json:"serviceOperator,omitempty"`
	Service
}

type GovernmentOrganization struct {
	Organization
}

type SpeakableSpecification struct {
	Xpath       XPathType       `json:"xpath,omitempty"`
	CssSelector CssSelectorType `json:"cssSelector,omitempty"`
	Intangible
}

type FDAnotEvaluated struct {
	DrugPregnancyCategory
}

type EducationEvent struct {
	Event
}

type TransferAction struct {
	FromLocation Place `json:"fromLocation,omitempty"`
	ToLocation   Place `json:"toLocation,omitempty"`
	Action
}

type EnrollingByInvitation struct {
	MedicalStudyStatus
}

type RightHandDriving struct {
	SteeringPositionValue
}

type MedicalSpecialty struct {
	Specialty
	MedicalEnumeration
}

type CivicStructure struct {
	OpeningHours Text `json:"openingHours,omitempty"`
	Place
}

type Dermatologic struct {
	MedicalSpecialty
}

type SteeringPositionValue struct {
	QualitativeValue
}

type GamePlayMode struct {
	Enumeration
}

type Seat struct {
	SeatingType interface{} `json:"seatingType,omitempty"` // Text, QualitativeValue
	SeatNumber  Text        `json:"seatNumber,omitempty"`
	SeatRow     Text        `json:"seatRow,omitempty"`
	SeatSection Text        `json:"seatSection,omitempty"`
	Intangible
}

type NoteDigitalDocument struct {
	DigitalDocument
}

type MusicAlbum struct {
	AlbumProductionType MusicAlbumProductionType `json:"albumProductionType,omitempty"`
	ByArtist            MusicGroup               `json:"byArtist,omitempty"`
	AlbumReleaseType    MusicAlbumReleaseType    `json:"albumReleaseType,omitempty"`
	AlbumRelease        MusicRelease             `json:"albumRelease,omitempty"`
	MusicPlaylist
}

type TaxiVehicleUsage struct {
	CarUsageType
}

type PreventionIndication struct {
	MedicalIndication
}

type CommentPermission struct {
	DigitalDocumentPermissionType
}

type EventRescheduled struct {
	EventStatusType
}

type SomeProducts struct {
	InventoryLevel QuantitativeValue `json:"inventoryLevel,omitempty"`
	Product
}

type MedicalTrial struct {
	TrialDesign MedicalTrialDesign `json:"trialDesign,omitempty"`
	Phase       Text               `json:"phase,omitempty"`
	MedicalStudy
}

type PayAction struct {
	Recipient interface{} `json:"recipient,omitempty"` // Person, Organization, ContactPoint, Audience
	Purpose   interface{} `json:"purpose,omitempty"`   // Thing, MedicalDevicePurpose
	TradeAction
}

type PaymentPastDue struct {
	PaymentStatusType
}

type SoftwareSourceCode struct {
	CodeRepository      URL                 `json:"codeRepository,omitempty"`
	ProgrammingLanguage interface{}         `json:"programmingLanguage,omitempty"` // ComputerLanguage, Text
	TargetProduct       SoftwareApplication `json:"targetProduct,omitempty"`
	Runtime             Text                `json:"runtime,omitempty"`
	SampleType          Text                `json:"sampleType,omitempty"`
	CodeSampleType      Text                `json:"codeSampleType,omitempty"`
	RuntimePlatform     Text                `json:"runtimePlatform,omitempty"`
	CreativeWork
}

type TripleBlindedTrial struct {
	MedicalTrialDesign
}

type SendAction struct {
	DeliveryMethod DeliveryMethod `json:"deliveryMethod,omitempty"`
	Recipient      interface{}    `json:"recipient,omitempty"` // Person, Organization, ContactPoint, Audience
	TransferAction
}

type PaymentService struct {
	FinancialProduct
}

type LeaveAction struct {
	Event Event `json:"event,omitempty"`
	InteractAction
}

type VideoGameSeries struct {
	Episodes           Episode            `json:"episodes,omitempty"`
	GameLocation       interface{}        `json:"gameLocation,omitempty"` // PostalAddress, Place, URL
	NumberOfEpisodes   Integer            `json:"numberOfEpisodes,omitempty"`
	Actors             Person             `json:"actors,omitempty"`
	Season             CreativeWorkSeason `json:"season,omitempty"`
	Trailer            VideoObject        `json:"trailer,omitempty"`
	CharacterAttribute Thing              `json:"characterAttribute,omitempty"`
	GamePlatform       interface{}        `json:"gamePlatform,omitempty"` // Thing, Text, URL
	ProductionCompany  Organization       `json:"productionCompany,omitempty"`
	Actor              Person             `json:"actor,omitempty"`
	Episode            Episode            `json:"episode,omitempty"`
	PlayMode           GamePlayMode       `json:"playMode,omitempty"`
	NumberOfSeasons    Integer            `json:"numberOfSeasons,omitempty"`
	Quest              Thing              `json:"quest,omitempty"`
	GameItem           Thing              `json:"gameItem,omitempty"`
	ContainsSeason     CreativeWorkSeason `json:"containsSeason,omitempty"`
	Director           Person             `json:"director,omitempty"`
	Directors          Person             `json:"directors,omitempty"`
	CheatCode          CreativeWork       `json:"cheatCode,omitempty"`
	Seasons            CreativeWorkSeason `json:"seasons,omitempty"`
	MusicBy            interface{}        `json:"musicBy,omitempty"` // MusicGroup, Person
	NumberOfPlayers    QuantitativeValue  `json:"numberOfPlayers,omitempty"`
	CreativeWorkSeries
}

type Winery struct {
	FoodEstablishment
}

type LegalValueLevel struct {
	Enumeration
}

type Library struct {
	LocalBusiness
}

type MedicalDevice struct {
	AdverseOutcome        MedicalEntity     `json:"adverseOutcome,omitempty"`
	Contraindication      interface{}       `json:"contraindication,omitempty"` // Text, MedicalContraindication
	Procedure             Text              `json:"procedure,omitempty"`
	PostOp                Text              `json:"postOp,omitempty"`
	Purpose               interface{}       `json:"purpose,omitempty"` // Thing, MedicalDevicePurpose
	Indication            MedicalIndication `json:"indication,omitempty"`
	PreOp                 Text              `json:"preOp,omitempty"`
	SeriousAdverseOutcome MedicalEntity     `json:"seriousAdverseOutcome,omitempty"`
	MedicalEntity
}

type Apartment struct {
	Occupancy     QuantitativeValue `json:"occupancy,omitempty"`
	NumberOfRooms interface{}       `json:"numberOfRooms,omitempty"` // QuantitativeValue, Number
	Accommodation
}

type CDFormat struct {
	MusicReleaseFormatType
}

type SpeechPathology struct {
	MedicalSpecialty
}

type MedicalIntangible struct {
	MedicalEntity
}

type MobileApplication struct {
	CarrierRequirements Text `json:"carrierRequirements,omitempty"`
	SoftwareApplication
}

type BikeStore struct {
	Store
}

type Game struct {
	GameItem           Thing             `json:"gameItem,omitempty"`
	Quest              Thing             `json:"quest,omitempty"`
	CharacterAttribute Thing             `json:"characterAttribute,omitempty"`
	GameLocation       interface{}       `json:"gameLocation,omitempty"` // PostalAddress, Place, URL
	NumberOfPlayers    QuantitativeValue `json:"numberOfPlayers,omitempty"`
	CreativeWork
}

type DaySpa struct {
	HealthAndBeautyBusiness
}

type KosherDiet struct {
	RestrictedDiet
}

type StructuredValue struct {
	Intangible
}

type SocialMediaPosting struct {
	SharedContent CreativeWork `json:"sharedContent,omitempty"`
	Article
}

type Lung struct {
	PhysicalExam
}

type ReadAction struct {
	ConsumeAction
}

type BackgroundNewsArticle struct {
	NewsArticle
}

type PublicationEvent struct {
	PublishedOn         BroadcastService `json:"publishedOn,omitempty"`
	PublishedBy         interface{}      `json:"publishedBy,omitempty"` // Organization, Person
	Free                Boolean          `json:"free,omitempty"`
	IsAccessibleForFree Boolean          `json:"isAccessibleForFree,omitempty"`
	Event
}

type CaseSeries struct {
	MedicalObservationalStudyDesign
}

type SoldOut struct {
	ItemAvailability
}

type BroadcastChannel struct {
	BroadcastFrequency   interface{}             `json:"broadcastFrequency,omitempty"` // Text, BroadcastFrequencySpecification
	BroadcastServiceTier Text                    `json:"broadcastServiceTier,omitempty"`
	Genre                interface{}             `json:"genre,omitempty"` // URL, Text
	InBroadcastLineup    CableOrSatelliteService `json:"inBroadcastLineup,omitempty"`
	BroadcastChannelId   Text                    `json:"broadcastChannelId,omitempty"`
	Intangible
}

type Season struct {
	CreativeWork
}

type MixtapeAlbum struct {
	MusicAlbumProductionType
}

type Thesis struct {
	InSupportOf Text `json:"inSupportOf,omitempty"`
	CreativeWork
}

type MusicComposition struct {
	Composer             interface{}      `json:"composer,omitempty"` // Organization, Person
	RecordedAs           MusicRecording   `json:"recordedAs,omitempty"`
	IncludedComposition  MusicComposition `json:"includedComposition,omitempty"`
	MusicCompositionForm Text             `json:"musicCompositionForm,omitempty"`
	MusicalKey           Text             `json:"musicalKey,omitempty"`
	Lyrics               CreativeWork     `json:"lyrics,omitempty"`
	Lyricist             Person           `json:"lyricist,omitempty"`
	MusicArrangement     MusicComposition `json:"musicArrangement,omitempty"`
	IswcCode             Text             `json:"iswcCode,omitempty"`
	FirstPerformance     Event            `json:"firstPerformance,omitempty"`
	CreativeWork
}

type Nerve struct {
	SourcedFrom BrainStructure      `json:"sourcedFrom,omitempty"`
	Branch      AnatomicalStructure `json:"branch,omitempty"`
	NerveMotor  Muscle              `json:"nerveMotor,omitempty"`
	SensoryUnit interface{}         `json:"sensoryUnit,omitempty"` // AnatomicalStructure, SuperficialAnatomy
	AnatomicalStructure
}

type BusTrip struct {
	BusNumber        Text        `json:"busNumber,omitempty"`
	Provider         interface{} `json:"provider,omitempty"` // Organization, Person
	ArrivalTime      DateTime    `json:"arrivalTime,omitempty"`
	ArrivalBusStop   interface{} `json:"arrivalBusStop,omitempty"` // BusStation, BusStop
	DepartureTime    DateTime    `json:"departureTime,omitempty"`
	DepartureBusStop interface{} `json:"departureBusStop,omitempty"` // BusStation, BusStop
	BusName          Text        `json:"busName,omitempty"`
	Intangible
}

type VeganDiet struct {
	RestrictedDiet
}

type Place struct {
	SpecialOpeningHoursSpecification OpeningHoursSpecification    `json:"specialOpeningHoursSpecification,omitempty"`
	GeospatiallyOverlaps             interface{}                  `json:"geospatiallyOverlaps,omitempty"` // GeospatialGeometry, Place
	Event                            Event                        `json:"event,omitempty"`
	FaxNumber                        Text                         `json:"faxNumber,omitempty"`
	ContainsPlace                    Place                        `json:"containsPlace,omitempty"`
	GeospatiallyWithin               interface{}                  `json:"geospatiallyWithin,omitempty"` // GeospatialGeometry, Place
	IsAccessibleForFree              Boolean                      `json:"isAccessibleForFree,omitempty"`
	IsicV4                           Text                         `json:"isicV4,omitempty"`
	ContainedIn                      Place                        `json:"containedIn,omitempty"`
	Photos                           interface{}                  `json:"photos,omitempty"` // Photograph, ImageObject
	Maps                             URL                          `json:"maps,omitempty"`
	AmenityFeature                   LocationFeatureSpecification `json:"amenityFeature,omitempty"`
	GeospatiallyCrosses              interface{}                  `json:"geospatiallyCrosses,omitempty"` // Place, GeospatialGeometry
	GeospatiallyCovers               interface{}                  `json:"geospatiallyCovers,omitempty"`  // Place, GeospatialGeometry
	Geo                              interface{}                  `json:"geo,omitempty"`                 // GeoCoordinates, GeoShape
	OpeningHoursSpecification        OpeningHoursSpecification    `json:"openingHoursSpecification,omitempty"`
	AggregateRating                  AggregateRating              `json:"aggregateRating,omitempty"`
	HasMap                           interface{}                  `json:"hasMap,omitempty"` // URL, Map
	PublicAccess                     Boolean                      `json:"publicAccess,omitempty"`
	GeospatiallyContains             interface{}                  `json:"geospatiallyContains,omitempty"` // GeospatialGeometry, Place
	AdditionalProperty               PropertyValue                `json:"additionalProperty,omitempty"`
	Telephone                        Text                         `json:"telephone,omitempty"`
	Logo                             interface{}                  `json:"logo,omitempty"` // URL, ImageObject
	Reviews                          Review                       `json:"reviews,omitempty"`
	Address                          interface{}                  `json:"address,omitempty"`              // PostalAddress, Text
	GeospatiallyDisjoint             interface{}                  `json:"geospatiallyDisjoint,omitempty"` // Place, GeospatialGeometry
	GeospatiallyEquals               interface{}                  `json:"geospatiallyEquals,omitempty"`   // GeospatialGeometry, Place
	GlobalLocationNumber             Text                         `json:"globalLocationNumber,omitempty"`
	Events                           Event                        `json:"events,omitempty"`
	SmokingAllowed                   Boolean                      `json:"smokingAllowed,omitempty"`
	GeospatiallyCoveredBy            interface{}                  `json:"geospatiallyCoveredBy,omitempty"`  // Place, GeospatialGeometry
	GeospatiallyIntersects           interface{}                  `json:"geospatiallyIntersects,omitempty"` // GeospatialGeometry, Place
	BranchCode                       Text                         `json:"branchCode,omitempty"`
	Map                              URL                          `json:"map,omitempty"`
	Review                           Review                       `json:"review,omitempty"`
	GeospatiallyTouches              interface{}                  `json:"geospatiallyTouches,omitempty"` // GeospatialGeometry, Place
	MaximumAttendeeCapacity          Integer                      `json:"maximumAttendeeCapacity,omitempty"`
	ContainedInPlace                 Place                        `json:"containedInPlace,omitempty"`
	Photo                            interface{}                  `json:"photo,omitempty"` // Photograph, ImageObject
	Thing
}

type PerformanceRole struct {
	CharacterName Text `json:"characterName,omitempty"`
	Role
}

type Synagogue struct {
	PlaceOfWorship
}

type ReservationConfirmed struct {
	ReservationStatusType
}

type PET struct {
	MedicalImagingTechnique
}

type DepartAction struct {
	MoveAction
}

type HardwareStore struct {
	Store
}

type HVACBusiness struct {
	HomeAndConstructionBusiness
}

type PaintAction struct {
	CreateAction
}

type ParkingFacility struct {
	CivicStructure
}

type ResultsNotAvailable struct {
	MedicalStudyStatus
}

type ReplyAction struct {
	ResultComment Comment `json:"resultComment,omitempty"`
	CommunicateAction
}

type RsvpAction struct {
	RsvpResponse             RsvpResponseType `json:"rsvpResponse,omitempty"`
	Comment                  Comment          `json:"comment,omitempty"`
	AdditionalNumberOfGuests Number           `json:"additionalNumberOfGuests,omitempty"`
	InformAction
}

type BedDetails struct {
	NumberOfBeds Number `json:"numberOfBeds,omitempty"`
	Intangible
}

type NewsArticle struct {
	Dateline     Text `json:"dateline,omitempty"`
	PrintSection Text `json:"printSection,omitempty"`
	PrintPage    Text `json:"printPage,omitempty"`
	PrintEdition Text `json:"printEdition,omitempty"`
	PrintColumn  Text `json:"printColumn,omitempty"`
	Article
}

type MedicalRiskFactor struct {
	IncreasesRiskOf MedicalEntity `json:"increasesRiskOf,omitempty"`
	MedicalEntity
}

type EvidenceLevelC struct {
	MedicalEvidenceLevel
}

type UpdateAction struct {
	Collection       Thing `json:"collection,omitempty"`
	TargetCollection Thing `json:"targetCollection,omitempty"`
	Action
}

type LiquorStore struct {
	Store
}

type Florist struct {
	Store
}

type SubwayStation struct {
	CivicStructure
}

type HealthClub struct {
	SportsActivityLocation
	HealthAndBeautyBusiness
}

type FollowAction struct {
	Followee interface{} `json:"followee,omitempty"` // Person, Organization
	InteractAction
}

type Longitudinal struct {
	MedicalObservationalStudyDesign
}

type Chapter struct {
	PageEnd    interface{} `json:"pageEnd,omitempty"`    //
	Pagination interface{} `json:"pagination,omitempty"` //
	PageStart  interface{} `json:"pageStart,omitempty"`  //
	CreativeWork
}

type QualitativeValue struct {
	Greater            QualitativeValue `json:"greater,omitempty"`
	Equal              QualitativeValue `json:"equal,omitempty"`
	AdditionalProperty PropertyValue    `json:"additionalProperty,omitempty"`
	GreaterOrEqual     QualitativeValue `json:"greaterOrEqual,omitempty"`
	Lesser             QualitativeValue `json:"lesser,omitempty"`
	LesserOrEqual      QualitativeValue `json:"lesserOrEqual,omitempty"`
	ValueReference     interface{}      `json:"valueReference,omitempty"` // Enumeration, PropertyValue, StructuredValue, QuantitativeValue, QualitativeValue
	NonEqual           QualitativeValue `json:"nonEqual,omitempty"`
	Enumeration
}

type MedicalImagingTechnique struct {
	MedicalEnumeration
}

type IgnoreAction struct {
	AssessAction
}

type BowlingAlley struct {
	SportsActivityLocation
}

type CreativeWorkSeries struct {
	EndDate   interface{} `json:"endDate,omitempty"` // Date, DateTime
	Issn      Text        `json:"issn,omitempty"`
	StartDate interface{} `json:"startDate,omitempty"` // DateTime, Date
	CreativeWork
}

type AutoBodyShop struct {
	AutomotiveBusiness
}

type UserReview struct {
	Review
}

type EntertainmentBusiness struct {
	LocalBusiness
}

type HowToSupply struct {
	EstimatedCost interface{} `json:"estimatedCost,omitempty"` // Text, MonetaryAmount
	HowToItem
}

type EndorsementRating struct {
	Rating
}

type TaxiReservation struct {
	PickupTime     DateTime    `json:"pickupTime,omitempty"`
	PartySize      interface{} `json:"partySize,omitempty"` // Integer, QuantitativeValue
	PickupLocation Place       `json:"pickupLocation,omitempty"`
	Reservation
}

type Vein struct {
	DrainsTo      Vessel              `json:"drainsTo,omitempty"`
	Tributary     AnatomicalStructure `json:"tributary,omitempty"`
	RegionDrained interface{}         `json:"regionDrained,omitempty"` // AnatomicalSystem, AnatomicalStructure
	Vessel
}

type Genetic struct {
	MedicalSpecialty
}

type EngineSpecification struct {
	EngineType         interface{}       `json:"engineType,omitempty"` // Text, URL, QualitativeValue
	EngineDisplacement QuantitativeValue `json:"engineDisplacement,omitempty"`
	Torque             QuantitativeValue `json:"torque,omitempty"`
	FuelType           interface{}       `json:"fuelType,omitempty"` // QualitativeValue, URL, Text
	EnginePower        QuantitativeValue `json:"enginePower,omitempty"`
	StructuredValue
}

type BeautySalon struct {
	HealthAndBeautyBusiness
}

type LocationFeatureSpecification struct {
	HoursAvailable OpeningHoursSpecification `json:"hoursAvailable,omitempty"`
	ValidThrough   DateTime                  `json:"validThrough,omitempty"`
	ValidFrom      DateTime                  `json:"validFrom,omitempty"`
	PropertyValue
}

type CompoundPriceSpecification struct {
	PriceComponent UnitPriceSpecification `json:"priceComponent,omitempty"`
	PriceSpecification
}

type RegisterAction struct {
	InteractAction
}

type AlignmentObject struct {
	TargetName           Text `json:"targetName,omitempty"`
	TargetDescription    Text `json:"targetDescription,omitempty"`
	TargetUrl            URL  `json:"targetUrl,omitempty"`
	AlignmentType        Text `json:"alignmentType,omitempty"`
	EducationalFramework Text `json:"educationalFramework,omitempty"`
	Intangible
}

type EventStatusType struct {
	Enumeration
}

type WPSideBar struct {
	WebPageElement
}

type CatholicChurch struct {
	PlaceOfWorship
}

type EPRelease struct {
	MusicAlbumReleaseType
}

type SearchResultsPage struct {
	WebPage
}

type CreateAction struct {
	Action
}

type RentalCarReservation struct {
	PickupTime      DateTime `json:"pickupTime,omitempty"`
	DropoffLocation Place    `json:"dropoffLocation,omitempty"`
	PickupLocation  Place    `json:"pickupLocation,omitempty"`
	DropoffTime     DateTime `json:"dropoffTime,omitempty"`
	Reservation
}

type VisualArtwork struct {
	Letterer       Person      `json:"letterer,omitempty"`
	Penciler       Person      `json:"penciler,omitempty"`
	Surface        interface{} `json:"surface,omitempty"` // URL, Text
	Depth          interface{} `json:"depth,omitempty"`   // QuantitativeValue, Distance
	Artform        interface{} `json:"artform,omitempty"` // URL, Text
	Width          interface{} `json:"width,omitempty"`   // Distance, QuantitativeValue
	Colorist       Person      `json:"colorist,omitempty"`
	Inker          Person      `json:"inker,omitempty"`
	ArtworkSurface interface{} `json:"artworkSurface,omitempty"` // Text, URL
	Artist         Person      `json:"artist,omitempty"`
	ArtMedium      interface{} `json:"artMedium,omitempty"`  // Text, URL
	Height         interface{} `json:"height,omitempty"`     // Distance, QuantitativeValue
	ArtEdition     interface{} `json:"artEdition,omitempty"` // Text, Integer
	CreativeWork
}

type TrackAction struct {
	DeliveryMethod DeliveryMethod `json:"deliveryMethod,omitempty"`
	FindAction
}

type Continent struct {
	Landform
}

type MovieClip struct {
	Clip
}

type MedicalResearcher struct {
	MedicalAudience
}

type DigitalFormat struct {
	MusicReleaseFormatType
}

type LeftHandDriving struct {
	SteeringPositionValue
}

type Neck struct {
	PhysicalExam
}

type MedicalClinic struct {
	MedicalSpecialty MedicalSpecialty `json:"medicalSpecialty,omitempty"`
	AvailableService interface{}      `json:"availableService,omitempty"` // MedicalTest, MedicalProcedure, MedicalTherapy
	MedicalBusiness
	MedicalOrganization
}

type PerformingArtsTheater struct {
	CivicStructure
}

type MedicalContraindication struct {
	MedicalEntity
}

type Country struct {
	AdministrativeArea
}

type Head struct {
	PhysicalExam
}

type DrugStrength struct {
	AvailableIn      AdministrativeArea  `json:"availableIn,omitempty"`
	ActiveIngredient Text                `json:"activeIngredient,omitempty"`
	StrengthUnit     Text                `json:"strengthUnit,omitempty"`
	StrengthValue    Number              `json:"strengthValue,omitempty"`
	MaximumIntake    MaximumDoseSchedule `json:"maximumIntake,omitempty"`
	MedicalIntangible
}

type OrderItem struct {
	OrderItemStatus OrderStatus    `json:"orderItemStatus,omitempty"`
	OrderQuantity   Number         `json:"orderQuantity,omitempty"`
	OrderDelivery   ParcelDelivery `json:"orderDelivery,omitempty"`
	OrderItemNumber Text           `json:"orderItemNumber,omitempty"`
	OrderedItem     interface{}    `json:"orderedItem,omitempty"` // OrderItem, Product
	Intangible
}

type Pathology struct {
	MedicalSpecialty
}

type Vessel struct {
	AnatomicalStructure
}

type ComicIssue struct {
	Inker        Person `json:"inker,omitempty"`
	Letterer     Person `json:"letterer,omitempty"`
	VariantCover Text   `json:"variantCover,omitempty"`
	Artist       Person `json:"artist,omitempty"`
	Colorist     Person `json:"colorist,omitempty"`
	Penciler     Person `json:"penciler,omitempty"`
	PublicationIssue
}

type Comment struct {
	ParentItem    Question `json:"parentItem,omitempty"`
	DownvoteCount Integer  `json:"downvoteCount,omitempty"`
	UpvoteCount   Integer  `json:"upvoteCount,omitempty"`
	CreativeWork
}

type HealthAndBeautyBusiness struct {
	LocalBusiness
}

type Schedule struct {
	EventSchedule   Duration    `json:"eventSchedule,omitempty"`
	ExceptDate      interface{} `json:"exceptDate,omitempty"`      // Date, DateTime
	RepeatFrequency interface{} `json:"repeatFrequency,omitempty"` // Duration, Text
	ByMonth         Integer     `json:"byMonth,omitempty"`
	RepeatCount     Integer     `json:"repeatCount,omitempty"`
	ByMonthDay      Integer     `json:"byMonthDay,omitempty"`
	ByDay           DayOfWeek   `json:"byDay,omitempty"`
	Intangible
}

type Sunday struct {
	DayOfWeek
}

type ProfessionalService struct {
	LocalBusiness
}

type InternetCafe struct {
	LocalBusiness
}

type VitalSign struct {
	MedicalSign
}

type BrainStructure struct {
	AnatomicalStructure
}

type Diagnostic struct {
	MedicalDevicePurpose
}

type ItemPage struct {
	WebPage
}

type LowCalorieDiet struct {
	RestrictedDiet
}

type MortgageLoan struct {
	LoanMortgageMandateAmount MonetaryAmount `json:"loanMortgageMandateAmount,omitempty"`
	DomiciledMortgage         Boolean        `json:"domiciledMortgage,omitempty"`
	LoanOrCredit
}

type AggregateOffer struct {
	LowPrice   interface{} `json:"lowPrice,omitempty"` // Text, Number
	Offers     Offer       `json:"offers,omitempty"`
	OfferCount Integer     `json:"offerCount,omitempty"`
	HighPrice  interface{} `json:"highPrice,omitempty"` // Text, Number
	Offer
}

type InfectiousDisease struct {
	InfectiousAgentClass InfectiousAgentClass `json:"infectiousAgentClass,omitempty"`
	TransmissionMethod   Text                 `json:"transmissionMethod,omitempty"`
	InfectiousAgent      Text                 `json:"infectiousAgent,omitempty"`
	MedicalCondition
}

type PalliativeProcedure struct {
	MedicalTherapy
	MedicalProcedure
}

type Radiography struct {
}

type MedicalProcedureType struct {
	MedicalEnumeration
}

type BarOrPub struct {
	FoodEstablishment
}

type PeopleAudience struct {
	RequiredMinAge  Integer          `json:"requiredMinAge,omitempty"`
	RequiredMaxAge  Integer          `json:"requiredMaxAge,omitempty"`
	RequiredGender  Text             `json:"requiredGender,omitempty"`
	HealthCondition MedicalCondition `json:"healthCondition,omitempty"`
	SuggestedMaxAge Number           `json:"suggestedMaxAge,omitempty"`
	SuggestedGender Text             `json:"suggestedGender,omitempty"`
	SuggestedMinAge Number           `json:"suggestedMinAge,omitempty"`
	Audience
}

type JewelryStore struct {
	Store
}

type Ear struct {
	PhysicalExam
}

type EmailMessage struct {
	Message
}

type PhysicalActivityCategory struct {
	Enumeration
}

type ComicSeries struct {
	Periodical
}

type Mountain struct {
	Landform
}

type Optometric struct {
	MedicalBusiness
	MedicalSpecialty
}

type OceanBodyOfWater struct {
	BodyOfWater
}

type ItemList struct {
	ItemListOrder   interface{} `json:"itemListOrder,omitempty"` // Text, ItemListOrderType
	NumberOfItems   Integer     `json:"numberOfItems,omitempty"`
	ItemListElement interface{} `json:"itemListElement,omitempty"` // ListItem, Text, Thing
	Intangible
}

type TakeAction struct {
	TransferAction
}

type UserBlocks struct {
	UserInteraction
}

type FDAcategoryA struct {
	DrugPregnancyCategory
}

type Hardcover struct {
	BookFormatType
}

type AutoWash struct {
	AutomotiveBusiness
}

type PrimaryCare struct {
	MedicalBusiness
	MedicalSpecialty
}

type UserPlusOnes struct {
	UserInteraction
}

type OutletStore struct {
	Store
}

type ShoeStore struct {
	Store
}

type HowToStep struct {
	ItemList
}

type Muscle struct {
	Origin       AnatomicalStructure `json:"origin,omitempty"`
	Action       Text                `json:"action,omitempty"`
	Nerve        Nerve               `json:"nerve,omitempty"`
	BloodSupply  Vessel              `json:"bloodSupply,omitempty"`
	MuscleAction Text                `json:"muscleAction,omitempty"`
	Insertion    AnatomicalStructure `json:"insertion,omitempty"`
	Antagonist   Muscle              `json:"antagonist,omitempty"`
	AnatomicalStructure
}

type DiscoverAction struct {
	FindAction
}

type City struct {
	AdministrativeArea
}

type DrawAction struct {
	CreateAction
}

type MovieSeries struct {
	Director          Person       `json:"director,omitempty"`
	Trailer           VideoObject  `json:"trailer,omitempty"`
	MusicBy           interface{}  `json:"musicBy,omitempty"` // MusicGroup, Person
	Actor             Person       `json:"actor,omitempty"`
	Directors         Person       `json:"directors,omitempty"`
	ProductionCompany Organization `json:"productionCompany,omitempty"`
	Actors            Person       `json:"actors,omitempty"`
	CreativeWorkSeries
}

type LegalForceStatus struct {
	Enumeration
}

type EventPostponed struct {
	EventStatusType
}

type CampingPitch struct {
	Accommodation
}

type MeetingRoom struct {
	Room
}

type HealthInsurancePlan struct {
	UsesHealthPlanIdStandard    interface{}         `json:"usesHealthPlanIdStandard,omitempty"` // URL, Text
	IncludesHealthPlanFormulary HealthPlanFormulary `json:"includesHealthPlanFormulary,omitempty"`
	HealthPlanDrugTier          Text                `json:"healthPlanDrugTier,omitempty"`
	BenefitsSummaryUrl          URL                 `json:"benefitsSummaryUrl,omitempty"`
	HealthPlanId                Text                `json:"healthPlanId,omitempty"`
	IncludesHealthPlanNetwork   HealthPlanNetwork   `json:"includesHealthPlanNetwork,omitempty"`
	HealthPlanMarketingUrl      URL                 `json:"healthPlanMarketingUrl,omitempty"`
	ContactPoint                ContactPoint        `json:"contactPoint,omitempty"`
	HealthPlanDrugOption        Text                `json:"healthPlanDrugOption,omitempty"`
	Intangible
}

type RefurbishedCondition struct {
	OfferItemCondition
}

type PublicationIssue struct {
	IssueNumber interface{} `json:"issueNumber,omitempty"` // Text, Integer
	CreativeWork
}

type ResultsAvailable struct {
	MedicalStudyStatus
}

type Barcode struct {
	ImageObject
}

type Audiobook struct {
	Duration interface{} `json:"duration,omitempty"` //
	ReadBy   Person      `json:"readBy,omitempty"`
	AudioObject
	Book
}

type AmusementPark struct {
	EntertainmentBusiness
}

type PreOrder struct {
	ItemAvailability
}

type GovernmentBuilding struct {
	CivicStructure
}

type Flight struct {
	FlightDistance          interface{}        `json:"flightDistance,omitempty"` // Distance, Text
	Aircraft                interface{}        `json:"aircraft,omitempty"`       // Vehicle, Text
	Provider                interface{}        `json:"provider,omitempty"`       // Organization, Person
	ArrivalTerminal         Text               `json:"arrivalTerminal,omitempty"`
	DepartureGate           Text               `json:"departureGate,omitempty"`
	WebCheckinTime          DateTime           `json:"webCheckinTime,omitempty"`
	DepartureTime           DateTime           `json:"departureTime,omitempty"`
	Carrier                 Organization       `json:"carrier,omitempty"`
	DepartureTerminal       Text               `json:"departureTerminal,omitempty"`
	DepartureAirport        Airport            `json:"departureAirport,omitempty"`
	MealService             Text               `json:"mealService,omitempty"`
	ArrivalGate             Text               `json:"arrivalGate,omitempty"`
	FlightNumber            Text               `json:"flightNumber,omitempty"`
	BoardingPolicy          BoardingPolicyType `json:"boardingPolicy,omitempty"`
	EstimatedFlightDuration interface{}        `json:"estimatedFlightDuration,omitempty"` // Duration, Text
	ArrivalAirport          Airport            `json:"arrivalAirport,omitempty"`
	ArrivalTime             DateTime           `json:"arrivalTime,omitempty"`
	Seller                  interface{}        `json:"seller,omitempty"` // Person, Organization
	Intangible
}

type InForce struct {
	LegalForceStatus
}

type UnofficialLegalValue struct {
	LegalValueLevel
}

type False struct {
	Boolean
}

type Brand struct {
	Review          Review          `json:"review,omitempty"`
	Logo            interface{}     `json:"logo,omitempty"` // URL, ImageObject
	AggregateRating AggregateRating `json:"aggregateRating,omitempty"`
	Intangible
}

type Integer struct {
	Number
}

type TrainReservation struct {
	Reservation
}

type FoodEstablishment struct {
	HasMenu             interface{} `json:"hasMenu,omitempty"` // Text, URL, Menu
	StarRating          Rating      `json:"starRating,omitempty"`
	AcceptsReservations interface{} `json:"acceptsReservations,omitempty"` // URL, Boolean, Text
	Menu                interface{} `json:"menu,omitempty"`                // Text, URL, Menu
	ServesCuisine       Text        `json:"servesCuisine,omitempty"`
	LocalBusiness
}

type Invoice struct {
	PaymentDueDate       DateTime      `json:"paymentDueDate,omitempty"`
	Customer             interface{}   `json:"customer,omitempty"` // Organization, Person
	Broker               interface{}   `json:"broker,omitempty"`   // Person, Organization
	PaymentDue           DateTime      `json:"paymentDue,omitempty"`
	MinimumPaymentDue    interface{}   `json:"minimumPaymentDue,omitempty"` // PriceSpecification, MonetaryAmount
	Provider             interface{}   `json:"provider,omitempty"`          // Organization, Person
	PaymentMethodId      Text          `json:"paymentMethodId,omitempty"`
	ConfirmationNumber   Text          `json:"confirmationNumber,omitempty"`
	ReferencesOrder      Order         `json:"referencesOrder,omitempty"`
	PaymentMethod        PaymentMethod `json:"paymentMethod,omitempty"`
	ScheduledPaymentDate Date          `json:"scheduledPaymentDate,omitempty"`
	AccountId            Text          `json:"accountId,omitempty"`
	TotalPaymentDue      interface{}   `json:"totalPaymentDue,omitempty"` // MonetaryAmount, PriceSpecification
	PaymentStatus        interface{}   `json:"paymentStatus,omitempty"`   // PaymentStatusType, Text
	BillingPeriod        Duration      `json:"billingPeriod,omitempty"`
	Intangible
}

type MotorizedBicycle struct {
	Vehicle
}

type CommunicateAction struct {
	Language  Language    `json:"language,omitempty"`
	Recipient interface{} `json:"recipient,omitempty"` // Person, Organization, ContactPoint, Audience
	About     Thing       `json:"about,omitempty"`
	InteractAction
}

type Otolaryngologic struct {
	MedicalBusiness
	MedicalSpecialty
}

type ApartmentComplex struct {
	Residence
}

type BedType struct {
	QualitativeValue
}

type AccountingService struct {
	FinancialService
}

type ParkingMap struct {
	MapCategoryType
}

type PhysicalExam struct {
	MedicalEnumeration
	MedicalProcedure
}

type OrderProcessing struct {
	OrderStatus
}

type DrugLegalStatus struct {
	ApplicableLocation AdministrativeArea `json:"applicableLocation,omitempty"`
	MedicalIntangible
}

type BusinessAudience struct {
	YearsInOperation  QuantitativeValue `json:"yearsInOperation,omitempty"`
	YearlyRevenue     QuantitativeValue `json:"yearlyRevenue,omitempty"`
	NumberOfEmployees QuantitativeValue `json:"numberOfEmployees,omitempty"`
	Audience
}

type CookAction struct {
	FoodEvent         FoodEvent   `json:"foodEvent,omitempty"`
	Recipe            Recipe      `json:"recipe,omitempty"`
	FoodEstablishment interface{} `json:"foodEstablishment,omitempty"` // FoodEstablishment, Place
	CreateAction
}

type PublicHolidays struct {
	DayOfWeek
}

type Painting struct {
	CreativeWork
}

type PawnShop struct {
	Store
}

type LodgingBusiness struct {
	CheckoutTime      DateTime                     `json:"checkoutTime,omitempty"`
	AmenityFeature    LocationFeatureSpecification `json:"amenityFeature,omitempty"`
	PetsAllowed       interface{}                  `json:"petsAllowed,omitempty"` // Text, Boolean
	Audience          Audience                     `json:"audience,omitempty"`
	StarRating        Rating                       `json:"starRating,omitempty"`
	AvailableLanguage interface{}                  `json:"availableLanguage,omitempty"` // Text, Language
	CheckinTime       DateTime                     `json:"checkinTime,omitempty"`
	LocalBusiness
}

type Rheumatologic struct {
	MedicalSpecialty
}

type Geriatric struct {
	MedicalBusiness
	MedicalSpecialty
}

type HealthPlanFormulary struct {
	HealthPlanCostSharing    Boolean `json:"healthPlanCostSharing,omitempty"`
	OffersPrescriptionByMail Boolean `json:"offersPrescriptionByMail,omitempty"`
	HealthPlanDrugTier       Text    `json:"healthPlanDrugTier,omitempty"`
	Intangible
}

type Terminated struct {
	MedicalStudyStatus
}

type Plumber struct {
	HomeAndConstructionBusiness
}

type LoanOrCredit struct {
	RecourseLoan       Boolean                `json:"recourseLoan,omitempty"`
	LoanRepaymentForm  RepaymentSpecification `json:"loanRepaymentForm,omitempty"`
	Currency           interface{}            `json:"currency,omitempty"` //
	LoanTerm           QuantitativeValue      `json:"loanTerm,omitempty"`
	RenegotiableLoan   Boolean                `json:"renegotiableLoan,omitempty"`
	RequiredCollateral interface{}            `json:"requiredCollateral,omitempty"` // Thing, Text
	LoanType           interface{}            `json:"loanType,omitempty"`           // Text, URL
	GracePeriod        Duration               `json:"gracePeriod,omitempty"`
	FinancialProduct
}

type DrugPrescriptionStatus struct {
	MedicalEnumeration
}

type CommunityHealth struct {
	MedicalBusiness
	MedicalSpecialty
}

type Tuesday struct {
	DayOfWeek
}

type Person struct {
	Awards               Text              `json:"awards,omitempty"`
	Spouse               Person            `json:"spouse,omitempty"`
	IsicV4               Text              `json:"isicV4,omitempty"`
	VatID                Text              `json:"vatID,omitempty"`
	Children             Person            `json:"children,omitempty"`
	FamilyName           Text              `json:"familyName,omitempty"`
	RelatedTo            Person            `json:"relatedTo,omitempty"`
	Award                Text              `json:"award,omitempty"`
	Follows              Person            `json:"follows,omitempty"`
	Knows                Person            `json:"knows,omitempty"`
	NetWorth             interface{}       `json:"netWorth,omitempty"` // MonetaryAmount, PriceSpecification
	Owns                 interface{}       `json:"owns,omitempty"`     // Product, OwnershipInfo
	Naics                Text              `json:"naics,omitempty"`
	HonorificSuffix      Text              `json:"honorificSuffix,omitempty"`
	Funder               interface{}       `json:"funder,omitempty"` // Person, Organization
	FaxNumber            Text              `json:"faxNumber,omitempty"`
	Height               interface{}       `json:"height,omitempty"` // Distance, QuantitativeValue
	Telephone            Text              `json:"telephone,omitempty"`
	BirthPlace           Place             `json:"birthPlace,omitempty"`
	HonorificPrefix      Text              `json:"honorificPrefix,omitempty"`
	Nationality          Country           `json:"nationality,omitempty"`
	Parents              Person            `json:"parents,omitempty"`
	WorksFor             Organization      `json:"worksFor,omitempty"`
	DeathDate            Date              `json:"deathDate,omitempty"`
	Address              interface{}       `json:"address,omitempty"`      // PostalAddress, Text
	WorkLocation         interface{}       `json:"workLocation,omitempty"` // Place, ContactPoint
	BirthDate            Date              `json:"birthDate,omitempty"`
	TaxID                Text              `json:"taxID,omitempty"`
	AlumniOf             interface{}       `json:"alumniOf,omitempty"` // EducationalOrganization, Organization
	PerformerIn          Event             `json:"performerIn,omitempty"`
	PublishingPrinciples interface{}       `json:"publishingPrinciples,omitempty"` // CreativeWork, URL
	MemberOf             interface{}       `json:"memberOf,omitempty"`             // Organization, ProgramMembership
	ContactPoints        ContactPoint      `json:"contactPoints,omitempty"`
	Weight               QuantitativeValue `json:"weight,omitempty"`
	DeathPlace           Place             `json:"deathPlace,omitempty"`
	MakesOffer           Offer             `json:"makesOffer,omitempty"`
	Gender               interface{}       `json:"gender,omitempty"`       // Text, GenderType
	HomeLocation         interface{}       `json:"homeLocation,omitempty"` // ContactPoint, Place
	Sibling              Person            `json:"sibling,omitempty"`
	Brand                interface{}       `json:"brand,omitempty"` // Organization, Brand
	GlobalLocationNumber Text              `json:"globalLocationNumber,omitempty"`
	Seeks                Demand            `json:"seeks,omitempty"`
	GivenName            Text              `json:"givenName,omitempty"`
	AdditionalName       Text              `json:"additionalName,omitempty"`
	JobTitle             Text              `json:"jobTitle,omitempty"`
	Colleague            interface{}       `json:"colleague,omitempty"` // URL, Person
	Parent               Person            `json:"parent,omitempty"`
	HasPOS               Place             `json:"hasPOS,omitempty"`
	Colleagues           Person            `json:"colleagues,omitempty"`
	Duns                 Text              `json:"duns,omitempty"`
	Affiliation          Organization      `json:"affiliation,omitempty"`
	Siblings             Person            `json:"siblings,omitempty"`
	HasOfferCatalog      OfferCatalog      `json:"hasOfferCatalog,omitempty"`
	Email                Text              `json:"email,omitempty"`
	Thing
}

type DietNutrition struct {
	MedicalBusiness
	MedicalSpecialty
}

type Episode struct {
	Trailer           VideoObject        `json:"trailer,omitempty"`
	EpisodeNumber     interface{}        `json:"episodeNumber,omitempty"` // Text, Integer
	Directors         Person             `json:"directors,omitempty"`
	MusicBy           interface{}        `json:"musicBy,omitempty"` // MusicGroup, Person
	Actor             Person             `json:"actor,omitempty"`
	PartOfSeason      CreativeWorkSeason `json:"partOfSeason,omitempty"`
	PartOfSeries      CreativeWorkSeries `json:"partOfSeries,omitempty"`
	Actors            Person             `json:"actors,omitempty"`
	Director          Person             `json:"director,omitempty"`
	ProductionCompany Organization       `json:"productionCompany,omitempty"`
	CreativeWork
}

type ZoneBoardingPolicy struct {
	BoardingPolicyType
}

type LiteraryEvent struct {
	Event
}

type RVPark struct {
	CivicStructure
}

type Ligament struct {
	AnatomicalStructure
}

type MusicAlbumProductionType struct {
	Enumeration
}

type PriceSpecification struct {
	PriceCurrency             Text               `json:"priceCurrency,omitempty"`
	ValidThrough              DateTime           `json:"validThrough,omitempty"`
	Price                     interface{}        `json:"price,omitempty"` // Text, Number
	EligibleQuantity          QuantitativeValue  `json:"eligibleQuantity,omitempty"`
	EligibleTransactionVolume PriceSpecification `json:"eligibleTransactionVolume,omitempty"`
	MaxPrice                  Number             `json:"maxPrice,omitempty"`
	MinPrice                  Number             `json:"minPrice,omitempty"`
	ValueAddedTaxIncluded     Boolean            `json:"valueAddedTaxIncluded,omitempty"`
	ValidFrom                 DateTime           `json:"validFrom,omitempty"`
	StructuredValue
}

type AppendAction struct {
	InsertAction
}

type Museum struct {
	CivicStructure
}

type HinduTemple struct {
	PlaceOfWorship
}

type Clip struct {
	MusicBy       interface{}        `json:"musicBy,omitempty"`    // MusicGroup, Person
	ClipNumber    interface{}        `json:"clipNumber,omitempty"` // Text, Integer
	Actors        Person             `json:"actors,omitempty"`
	Actor         Person             `json:"actor,omitempty"`
	PartOfSeason  CreativeWorkSeason `json:"partOfSeason,omitempty"`
	Director      Person             `json:"director,omitempty"`
	Directors     Person             `json:"directors,omitempty"`
	PartOfSeries  CreativeWorkSeries `json:"partOfSeries,omitempty"`
	PartOfEpisode Episode            `json:"partOfEpisode,omitempty"`
	CreativeWork
}

type PublicSwimmingPool struct {
	SportsActivityLocation
}

type DigitalDocumentPermissionType struct {
	Enumeration
}

type MedicalSignOrSymptom struct {
	Cause             MedicalCause   `json:"cause,omitempty"`
	PossibleTreatment MedicalTherapy `json:"possibleTreatment,omitempty"`
	MedicalCondition
}

type WinAction struct {
	Loser Person `json:"loser,omitempty"`
	AchieveAction
}

type ItemListUnordered struct {
	ItemListOrderType
}

type CreativeWork struct {
	Video                VideoObject        `json:"video,omitempty"`
	Encodings            MediaObject        `json:"encodings,omitempty"`
	CommentCount         Integer            `json:"commentCount,omitempty"`
	SourceOrganization   Organization       `json:"sourceOrganization,omitempty"`
	IsFamilyFriendly     Boolean            `json:"isFamilyFriendly,omitempty"`
	Audience             Audience           `json:"audience,omitempty"`
	Contributor          interface{}        `json:"contributor,omitempty"` // Person, Organization
	AggregateRating      AggregateRating    `json:"aggregateRating,omitempty"`
	IsPartOf             CreativeWork       `json:"isPartOf,omitempty"`
	TimeRequired         Duration           `json:"timeRequired,omitempty"`
	Citation             interface{}        `json:"citation,omitempty"` // CreativeWork, Text
	Provider             interface{}        `json:"provider,omitempty"` // Organization, Person
	TranslationOfWork    CreativeWork       `json:"translationOfWork,omitempty"`
	AccessibilitySummary Text               `json:"accessibilitySummary,omitempty"`
	Creator              interface{}        `json:"creator,omitempty"` // Organization, Person
	Offers               Offer              `json:"offers,omitempty"`
	Award                Text               `json:"award,omitempty"`
	DateModified         interface{}        `json:"dateModified,omitempty"` // DateTime, Date
	Expires              Date               `json:"expires,omitempty"`
	ContentLocation      Place              `json:"contentLocation,omitempty"`
	SchemaVersion        interface{}        `json:"schemaVersion,omitempty"` // URL, Text
	EducationalAlignment AlignmentObject    `json:"educationalAlignment,omitempty"`
	AccountablePerson    Person             `json:"accountablePerson,omitempty"`
	About                Thing              `json:"about,omitempty"`
	AccessModeSufficient Text               `json:"accessModeSufficient,omitempty"`
	ThumbnailUrl         URL                `json:"thumbnailUrl,omitempty"`
	Genre                interface{}        `json:"genre,omitempty"` // URL, Text
	Editor               Person             `json:"editor,omitempty"`
	ExampleOfWork        CreativeWork       `json:"exampleOfWork,omitempty"`
	WorkExample          CreativeWork       `json:"workExample,omitempty"`
	AssociatedMedia      MediaObject        `json:"associatedMedia,omitempty"`
	PublisherImprint     Organization       `json:"publisherImprint,omitempty"`
	AccessibilityHazard  Text               `json:"accessibilityHazard,omitempty"`
	AccessibilityAPI     Text               `json:"accessibilityAPI,omitempty"`
	Character            Person             `json:"character,omitempty"`
	Reviews              Review             `json:"reviews,omitempty"`
	Mentions             Thing              `json:"mentions,omitempty"`
	Text                 Text               `json:"text,omitempty"`
	WorkTranslation      CreativeWork       `json:"workTranslation,omitempty"`
	LocationCreated      Place              `json:"locationCreated,omitempty"`
	CopyrightHolder      interface{}        `json:"copyrightHolder,omitempty"` // Organization, Person
	DatePublished        Date               `json:"datePublished,omitempty"`
	HasPart              CreativeWork       `json:"hasPart,omitempty"`
	RecordedAt           Event              `json:"recordedAt,omitempty"`
	SpatialCoverage      Place              `json:"spatialCoverage,omitempty"`
	LearningResourceType Text               `json:"learningResourceType,omitempty"`
	IsBasedOn            interface{}        `json:"isBasedOn,omitempty"` // Product, CreativeWork, URL
	Material             interface{}        `json:"material,omitempty"`  // Text, URL, Product
	Funder               interface{}        `json:"funder,omitempty"`    // Person, Organization
	CopyrightYear        Number             `json:"copyrightYear,omitempty"`
	Audio                AudioObject        `json:"audio,omitempty"`
	TypicalAgeRange      Text               `json:"typicalAgeRange,omitempty"`
	EducationalUse       Text               `json:"educationalUse,omitempty"`
	Version              interface{}        `json:"version,omitempty"` // Text, Number
	Headline             Text               `json:"headline,omitempty"`
	TemporalCoverage     interface{}        `json:"temporalCoverage,omitempty"` // DateTime, URL, Text
	Author               interface{}        `json:"author,omitempty"`           // Organization, Person
	Keywords             Text               `json:"keywords,omitempty"`
	DiscussionUrl        URL                `json:"discussionUrl,omitempty"`
	IsAccessibleForFree  Boolean            `json:"isAccessibleForFree,omitempty"`
	PublishingPrinciples interface{}        `json:"publishingPrinciples,omitempty"` // CreativeWork, URL
	Translator           interface{}        `json:"translator,omitempty"`           // Organization, Person
	Position             interface{}        `json:"position,omitempty"`             // Integer, Text
	Publisher            interface{}        `json:"publisher,omitempty"`            // Organization, Person
	IsBasedOnUrl         interface{}        `json:"isBasedOnUrl,omitempty"`         // Product, CreativeWork, URL
	AccessibilityControl Text               `json:"accessibilityControl,omitempty"`
	DateCreated          interface{}        `json:"dateCreated,omitempty"` // Date, DateTime
	MainEntity           Thing              `json:"mainEntity,omitempty"`
	Comment              Comment            `json:"comment,omitempty"`
	Producer             interface{}        `json:"producer,omitempty"` // Organization, Person
	Awards               Text               `json:"awards,omitempty"`
	AccessibilityFeature Text               `json:"accessibilityFeature,omitempty"`
	Encoding             MediaObject        `json:"encoding,omitempty"`
	FileFormat           interface{}        `json:"fileFormat,omitempty"` // URL, Text
	License              interface{}        `json:"license,omitempty"`    // CreativeWork, URL
	ContentRating        Text               `json:"contentRating,omitempty"`
	ReleasedEvent        PublicationEvent   `json:"releasedEvent,omitempty"`
	Review               Review             `json:"review,omitempty"`
	InteractivityType    Text               `json:"interactivityType,omitempty"`
	AccessMode           Text               `json:"accessMode,omitempty"`
	Publication          PublicationEvent   `json:"publication,omitempty"`
	AlternativeHeadline  Text               `json:"alternativeHeadline,omitempty"`
	InteractionStatistic InteractionCounter `json:"interactionStatistic,omitempty"`
	ContentReferenceTime DateTime           `json:"contentReferenceTime,omitempty"`
	Thing
}

type BankOrCreditUnion struct {
	FinancialService
}

type Atlas struct {
	CreativeWork
}

type Ayurvedic struct {
	MedicineSystem
}

type MovieTheater struct {
	ScreenCount Number `json:"screenCount,omitempty"`
	CivicStructure
	EntertainmentBusiness
}

type MedicalStudyStatus struct {
	MedicalEnumeration
}

type MusicPlaylist struct {
	Tracks    MusicRecording `json:"tracks,omitempty"`
	Track     interface{}    `json:"track,omitempty"` // ItemList, MusicRecording
	NumTracks Integer        `json:"numTracks,omitempty"`
	CreativeWork
}

type OutOfStock struct {
	ItemAvailability
}

type Endocrine struct {
	MedicalSpecialty
}

type MovieRentalStore struct {
	Store
}

type MedicalGuidelineRecommendation struct {
	RecommendationStrength Text `json:"recommendationStrength,omitempty"`
	MedicalGuideline
}

type Ticket struct {
	IssuedBy      Organization `json:"issuedBy,omitempty"`
	DateIssued    DateTime     `json:"dateIssued,omitempty"`
	TicketedSeat  Seat         `json:"ticketedSeat,omitempty"`
	TicketToken   interface{}  `json:"ticketToken,omitempty"` // URL, Text
	TotalPrice    interface{}  `json:"totalPrice,omitempty"`  // Number, Text, PriceSpecification
	UnderName     interface{}  `json:"underName,omitempty"`   // Organization, Person
	TicketNumber  Text         `json:"ticketNumber,omitempty"`
	PriceCurrency Text         `json:"priceCurrency,omitempty"`
	Intangible
}

type GameServerStatus struct {
	Enumeration
}

type EventCancelled struct {
	EventStatusType
}

type HowToSection struct {
	Steps interface{} `json:"steps,omitempty"` // ItemList, CreativeWork, Text
	ItemList
}

type SpokenWordAlbum struct {
	MusicAlbumProductionType
}

type SkiResort struct {
	SportsActivityLocation
}

type RoofingContractor struct {
	HomeAndConstructionBusiness
}

type MiddleSchool struct {
	EducationalOrganization
}

type EvidenceLevelA struct {
	MedicalEvidenceLevel
}

type TouristAttraction struct {
	AvailableLanguage interface{} `json:"availableLanguage,omitempty"` // Text, Language
	TouristType       interface{} `json:"touristType,omitempty"`       // Text, Audience
	Place
}

type Patient struct {
	Diagnosis       MedicalCondition `json:"diagnosis,omitempty"`
	HealthCondition MedicalCondition `json:"healthCondition,omitempty"`
	Drug            Drug             `json:"drug,omitempty"`
	Person
	MedicalAudience
}

type SingleRelease struct {
	MusicAlbumReleaseType
}

type CourseInstance struct {
	CourseMode interface{} `json:"courseMode,omitempty"` // Text, URL
	Instructor Person      `json:"instructor,omitempty"`
	Event
}

type UserPlays struct {
	UserInteraction
}

type VegetarianDiet struct {
	RestrictedDiet
}

type MediaObject struct {
	RegionsAllowed       Place        `json:"regionsAllowed,omitempty"`
	PlayerType           Text         `json:"playerType,omitempty"`
	Height               interface{}  `json:"height,omitempty"` // Distance, QuantitativeValue
	UploadDate           Date         `json:"uploadDate,omitempty"`
	Bitrate              Text         `json:"bitrate,omitempty"`
	Width                interface{}  `json:"width,omitempty"` // Distance, QuantitativeValue
	RequiresSubscription Boolean      `json:"requiresSubscription,omitempty"`
	EncodingFormat       Text         `json:"encodingFormat,omitempty"`
	EncodesCreativeWork  CreativeWork `json:"encodesCreativeWork,omitempty"`
	ContentSize          Text         `json:"contentSize,omitempty"`
	ProductionCompany    Organization `json:"productionCompany,omitempty"`
	AssociatedArticle    NewsArticle  `json:"associatedArticle,omitempty"`
	ContentUrl           URL          `json:"contentUrl,omitempty"`
	EmbedUrl             URL          `json:"embedUrl,omitempty"`
	CreativeWork
}

type DVDFormat struct {
	MusicReleaseFormatType
}

type AdultEntertainment struct {
	EntertainmentBusiness
}

type Vehicle struct {
	NumberOfForwardGears        interface{}           `json:"numberOfForwardGears,omitempty"`   // Number, QuantitativeValue
	FuelType                    interface{}           `json:"fuelType,omitempty"`               // QualitativeValue, URL, Text
	VehicleSeatingCapacity      interface{}           `json:"vehicleSeatingCapacity,omitempty"` // Number, QuantitativeValue
	ProductionDate              Date                  `json:"productionDate,omitempty"`
	MileageFromOdometer         QuantitativeValue     `json:"mileageFromOdometer,omitempty"`
	VehicleInteriorColor        Text                  `json:"vehicleInteriorColor,omitempty"`
	WeightTotal                 QuantitativeValue     `json:"weightTotal,omitempty"`
	FuelCapacity                QuantitativeValue     `json:"fuelCapacity,omitempty"`
	NumberOfAxles               interface{}           `json:"numberOfAxles,omitempty"` // Number, QuantitativeValue
	NumberOfDoors               interface{}           `json:"numberOfDoors,omitempty"` // Number, QuantitativeValue
	ModelDate                   Date                  `json:"modelDate,omitempty"`
	FuelEfficiency              QuantitativeValue     `json:"fuelEfficiency,omitempty"`
	EmissionsCO2                Number                `json:"emissionsCO2,omitempty"`
	Wheelbase                   QuantitativeValue     `json:"wheelbase,omitempty"`
	DateVehicleFirstRegistered  Date                  `json:"dateVehicleFirstRegistered,omitempty"`
	AccelerationTime            QuantitativeValue     `json:"accelerationTime,omitempty"`
	Payload                     QuantitativeValue     `json:"payload,omitempty"`
	VehicleConfiguration        Text                  `json:"vehicleConfiguration,omitempty"`
	VehicleTransmission         interface{}           `json:"vehicleTransmission,omitempty"` // Text, URL, QualitativeValue
	Speed                       QuantitativeValue     `json:"speed,omitempty"`
	NumberOfPreviousOwners      interface{}           `json:"numberOfPreviousOwners,omitempty"` // Number, QuantitativeValue
	VehicleModelDate            Date                  `json:"vehicleModelDate,omitempty"`
	VehicleEngine               EngineSpecification   `json:"vehicleEngine,omitempty"`
	SteeringPosition            SteeringPositionValue `json:"steeringPosition,omitempty"`
	DriveWheelConfiguration     interface{}           `json:"driveWheelConfiguration,omitempty"` // Text, DriveWheelConfigurationValue
	KnownVehicleDamages         Text                  `json:"knownVehicleDamages,omitempty"`
	BodyType                    interface{}           `json:"bodyType,omitempty"` // Text, QualitativeValue, URL
	TongueWeight                QuantitativeValue     `json:"tongueWeight,omitempty"`
	TrailerWeight               QuantitativeValue     `json:"trailerWeight,omitempty"`
	SeatingCapacity             interface{}           `json:"seatingCapacity,omitempty"` // Number, QuantitativeValue
	VehicleInteriorType         Text                  `json:"vehicleInteriorType,omitempty"`
	StupidProperty              QuantitativeValue     `json:"stupidProperty,omitempty"`
	VehicleIdentificationNumber Text                  `json:"vehicleIdentificationNumber,omitempty"`
	CargoVolume                 QuantitativeValue     `json:"cargoVolume,omitempty"`
	FuelConsumption             QuantitativeValue     `json:"fuelConsumption,omitempty"`
	MeetsEmissionStandard       interface{}           `json:"meetsEmissionStandard,omitempty"` // QualitativeValue, Text, URL
	PurchaseDate                Date                  `json:"purchaseDate,omitempty"`
	NumberOfAirbags             interface{}           `json:"numberOfAirbags,omitempty"` // Number, Text
	Product
}

type NotYetRecruiting struct {
	MedicalStudyStatus
}

type ResumeAction struct {
	ControlAction
}

type OnlineFull struct {
	GameServerStatus
}

type ElementarySchool struct {
	EducationalOrganization
}

type SoftwareApplication struct {
	AvailableOnDevice      Text                `json:"availableOnDevice,omitempty"`
	FeatureList            interface{}         `json:"featureList,omitempty"`          // Text, URL
	SoftwareRequirements   interface{}         `json:"softwareRequirements,omitempty"` // Text, URL
	CountriesSupported     Text                `json:"countriesSupported,omitempty"`
	MemoryRequirements     interface{}         `json:"memoryRequirements,omitempty"` // Text, URL
	ReleaseNotes           interface{}         `json:"releaseNotes,omitempty"`       // URL, Text
	Permissions            Text                `json:"permissions,omitempty"`
	SoftwareAddOn          SoftwareApplication `json:"softwareAddOn,omitempty"`
	InstallUrl             URL                 `json:"installUrl,omitempty"`
	SupportingData         DataFeed            `json:"supportingData,omitempty"`
	OperatingSystem        Text                `json:"operatingSystem,omitempty"`
	StorageRequirements    interface{}         `json:"storageRequirements,omitempty"` // URL, Text
	Screenshot             interface{}         `json:"screenshot,omitempty"`          // URL, ImageObject
	SoftwareVersion        Text                `json:"softwareVersion,omitempty"`
	ApplicationSuite       Text                `json:"applicationSuite,omitempty"`
	FileSize               Text                `json:"fileSize,omitempty"`
	ApplicationCategory    interface{}         `json:"applicationCategory,omitempty"` // Text, URL
	SoftwareHelp           CreativeWork        `json:"softwareHelp,omitempty"`
	ProcessorRequirements  Text                `json:"processorRequirements,omitempty"`
	Device                 Text                `json:"device,omitempty"`
	Requirements           interface{}         `json:"requirements,omitempty"` // Text, URL
	CountriesNotSupported  Text                `json:"countriesNotSupported,omitempty"`
	ApplicationSubCategory interface{}         `json:"applicationSubCategory,omitempty"` // Text, URL
	DownloadUrl            URL                 `json:"downloadUrl,omitempty"`
	CreativeWork
}

type PlasticSurgery struct {
	MedicalBusiness
	MedicalSpecialty
}

type CollectionPage struct {
	WebPage
}

type BankAccount struct {
	AccountOverdraftLimit MonetaryAmount `json:"accountOverdraftLimit,omitempty"`
	BankAccountType       interface{}    `json:"bankAccountType,omitempty"` // Text, URL
	AccountMinimumInflow  MonetaryAmount `json:"accountMinimumInflow,omitempty"`
	FinancialProduct
}

type InteractionCounter struct {
	InteractionType      Action      `json:"interactionType,omitempty"`
	InteractionService   interface{} `json:"interactionService,omitempty"` // WebSite, SoftwareApplication
	UserInteractionCount Integer     `json:"userInteractionCount,omitempty"`
	StructuredValue
}

type ShoppingCenter struct {
	LocalBusiness
}

type QuantitativeValue struct {
	ValueReference     interface{}   `json:"valueReference,omitempty"` // Enumeration, PropertyValue, StructuredValue, QuantitativeValue, QualitativeValue
	UnitCode           interface{}   `json:"unitCode,omitempty"`       // URL, Text
	UnitText           Text          `json:"unitText,omitempty"`
	MinValue           Number        `json:"minValue,omitempty"`
	Value              interface{}   `json:"value,omitempty"` // StructuredValue, Number, Boolean, Text
	AdditionalProperty PropertyValue `json:"additionalProperty,omitempty"`
	MaxValue           Number        `json:"maxValue,omitempty"`
	StructuredValue
}

type MedicalEnumeration struct {
	Enumeration
}

type AllWheelDriveConfiguration struct {
	DriveWheelConfigurationValue
}

type AuthoritativeLegalValue struct {
	LegalValueLevel
}

type DeleteAction struct {
	UpdateAction
}

type ComicStory struct {
	Colorist Person `json:"colorist,omitempty"`
	Penciler Person `json:"penciler,omitempty"`
	Inker    Person `json:"inker,omitempty"`
	Letterer Person `json:"letterer,omitempty"`
	Artist   Person `json:"artist,omitempty"`
	CreativeWork
}

type MusicEvent struct {
	Event
}

type GroceryStore struct {
	Store
}

type CreditCard struct {
	MonthlyMinimumRepaymentAmount interface{} `json:"monthlyMinimumRepaymentAmount,omitempty"` // MonetaryAmount, Number
	PaymentCard
	LoanOrCredit
}

type DataFeedItem struct {
	DateDeleted  DateTime    `json:"dateDeleted,omitempty"`
	DateModified interface{} `json:"dateModified,omitempty"` // DateTime, Date
	DateCreated  interface{} `json:"dateCreated,omitempty"`  // Date, DateTime
	Item         Thing       `json:"item,omitempty"`
	Intangible
}

type Article struct {
	WordCount      Integer     `json:"wordCount,omitempty"`
	ArticleSection Text        `json:"articleSection,omitempty"`
	Speakable      interface{} `json:"speakable,omitempty"` // SpeakableSpecification, URL
	ArticleBody    Text        `json:"articleBody,omitempty"`
	CreativeWork
}

type Physiotherapy struct {
	MedicalBusiness
	MedicalSpecialty
}

type GeoCoordinates struct {
	Elevation      interface{} `json:"elevation,omitempty"`      // Text, Number
	Address        interface{} `json:"address,omitempty"`        // PostalAddress, Text
	AddressCountry interface{} `json:"addressCountry,omitempty"` // Country, Text
	Longitude      interface{} `json:"longitude,omitempty"`      // Number, Text
	Latitude       interface{} `json:"latitude,omitempty"`       // Number, Text
	PostalCode     Text        `json:"postalCode,omitempty"`
	StructuredValue
}

type Language struct {
	Intangible
}

type TypeAndQuantityNode struct {
	TypeOfGood       interface{}      `json:"typeOfGood,omitempty"` // Service, Product
	BusinessFunction BusinessFunction `json:"businessFunction,omitempty"`
	UnitCode         interface{}      `json:"unitCode,omitempty"` // URL, Text
	UnitText         Text             `json:"unitText,omitempty"`
	AmountOfThisGood Number           `json:"amountOfThisGood,omitempty"`
	StructuredValue
}

type MusicStore struct {
	Store
}

type BuddhistTemple struct {
	PlaceOfWorship
}

type RandomizedTrial struct {
	MedicalTrialDesign
}

type ImageObject struct {
	ExifData             interface{} `json:"exifData,omitempty"` // Text, PropertyValue
	Thumbnail            ImageObject `json:"thumbnail,omitempty"`
	RepresentativeOfPage Boolean     `json:"representativeOfPage,omitempty"`
	Caption              Text        `json:"caption,omitempty"`
	MediaObject
}

type ItemAvailability struct {
	Enumeration
}

type RemixAlbum struct {
	MusicAlbumProductionType
}

type DonateAction struct {
	Recipient interface{} `json:"recipient,omitempty"` // Person, Organization, ContactPoint, Audience
	TradeAction
}

type PaymentDeclined struct {
	PaymentStatusType
}

type FDAcategoryD struct {
	DrugPregnancyCategory
}

type Motorcycle struct {
	Vehicle
}

type HealthPlanCostSharingSpecification struct {
	HealthPlanPharmacyCategory  Text               `json:"healthPlanPharmacyCategory,omitempty"`
	HealthPlanCoinsuranceRate   Number             `json:"healthPlanCoinsuranceRate,omitempty"`
	HealthPlanCoinsuranceOption Text               `json:"healthPlanCoinsuranceOption,omitempty"`
	HealthPlanCopay             PriceSpecification `json:"healthPlanCopay,omitempty"`
	HealthPlanCopayOption       Text               `json:"healthPlanCopayOption,omitempty"`
	Intangible
}

type PlaceOfWorship struct {
	CivicStructure
}

type AssessAction struct {
	Action
}

type AutomatedTeller struct {
	FinancialService
}

type MedicalDevicePurpose struct {
	MedicalEnumeration
}

type TextDigitalDocument struct {
	DigitalDocument
}

type Canal struct {
	BodyOfWater
}

type RadioSeries struct {
	MusicBy           interface{}        `json:"musicBy,omitempty"` // MusicGroup, Person
	Directors         Person             `json:"directors,omitempty"`
	ContainsSeason    CreativeWorkSeason `json:"containsSeason,omitempty"`
	Episode           Episode            `json:"episode,omitempty"`
	NumberOfEpisodes  Integer            `json:"numberOfEpisodes,omitempty"`
	Seasons           CreativeWorkSeason `json:"seasons,omitempty"`
	Actors            Person             `json:"actors,omitempty"`
	ProductionCompany Organization       `json:"productionCompany,omitempty"`
	Season            CreativeWorkSeason `json:"season,omitempty"`
	Actor             Person             `json:"actor,omitempty"`
	Episodes          Episode            `json:"episodes,omitempty"`
	Director          Person             `json:"director,omitempty"`
	NumberOfSeasons   Integer            `json:"numberOfSeasons,omitempty"`
	Trailer           VideoObject        `json:"trailer,omitempty"`
	CreativeWorkSeries
}

type RentAction struct {
	RealEstateAgent RealEstateAgent `json:"realEstateAgent,omitempty"`
	Landlord        interface{}     `json:"landlord,omitempty"` // Organization, Person
	TradeAction
}

type MedicalObservationalStudyDesign struct {
	MedicalEnumeration
}

type PaymentChargeSpecification struct {
	AppliesToDeliveryMethod DeliveryMethod `json:"appliesToDeliveryMethod,omitempty"`
	AppliesToPaymentMethod  PaymentMethod  `json:"appliesToPaymentMethod,omitempty"`
	PriceSpecification
}

type BodyOfWater struct {
	Landform
}

type OfflinePermanently struct {
	GameServerStatus
}

type ClaimReview struct {
	ClaimReviewed Text `json:"claimReviewed,omitempty"`
	Review
}

type MedicalCode struct {
	CodingSystem Text `json:"codingSystem,omitempty"`
	MedicalIntangible
	CategoryCode
}

type AMRadioChannel struct {
	RadioChannel
}

type RadiationTherapy struct {
	MedicalTherapy
}

type LifestyleModification struct {
	MedicalEntity
}

type InStoreOnly struct {
	ItemAvailability
}

type MedicalCause struct {
	CauseOf MedicalEntity `json:"causeOf,omitempty"`
	MedicalEntity
}

type FlightReservation struct {
	PassengerPriorityStatus interface{} `json:"passengerPriorityStatus,omitempty"` // Text, QualitativeValue
	SecurityScreening       Text        `json:"securityScreening,omitempty"`
	PassengerSequenceNumber Text        `json:"passengerSequenceNumber,omitempty"`
	BoardingGroup           Text        `json:"boardingGroup,omitempty"`
	Reservation
}

type PaymentComplete struct {
	PaymentStatusType
}

type PhysicalActivity struct {
	Pathophysiology   Text        `json:"pathophysiology,omitempty"`
	AssociatedAnatomy interface{} `json:"associatedAnatomy,omitempty"` // AnatomicalSystem, SuperficialAnatomy, AnatomicalStructure
	Epidemiology      Text        `json:"epidemiology,omitempty"`
	LifestyleModification
}

type RsvpResponseYes struct {
	RsvpResponseType
}

type SingleCenterTrial struct {
	MedicalTrialDesign
}

type Registry struct {
	MedicalObservationalStudyDesign
}

type WantAction struct {
	ReactAction
}

type Emergency struct {
	MedicalBusiness
	MedicalSpecialty
}

type MedicalSign struct {
	IdentifyingTest MedicalTest  `json:"identifyingTest,omitempty"`
	IdentifyingExam PhysicalExam `json:"identifyingExam,omitempty"`
	MedicalSignOrSymptom
}

type DigitalDocument struct {
	HasDigitalDocumentPermission DigitalDocumentPermission `json:"hasDigitalDocumentPermission,omitempty"`
	CreativeWork
}

type UserTweets struct {
	UserInteraction
}

type ExerciseGym struct {
	SportsActivityLocation
}

type DrugCostCategory struct {
	MedicalEnumeration
}

type EducationalOrganization struct {
	Alumni Person `json:"alumni,omitempty"`
	Organization
}

type LowFatDiet struct {
	RestrictedDiet
}

type PharmacySpecialty struct {
	MedicalSpecialty
}

type MulticellularParasite struct {
	InfectiousAgentClass
}

type Preschool struct {
	EducationalOrganization
}

type GovernmentPermit struct {
	Permit
}

type AboutPage struct {
	WebPage
}

type BroadcastRelease struct {
	MusicAlbumReleaseType
}

type Answer struct {
	Comment
}

type ViewAction struct {
	ConsumeAction
}

type Completed struct {
	MedicalStudyStatus
}

type MusculoskeletalExam struct {
	PhysicalExam
}

type PerformingGroup struct {
	Organization
}

type CarUsageType struct {
	QualitativeValue
}

type OfficeEquipmentStore struct {
	Store
}

type PublicToilet struct {
	CivicStructure
}

type SportingGoodsStore struct {
	Store
}

type CableOrSatelliteService struct {
	Service
}

type InfectiousAgentClass struct {
	MedicalEnumeration
}

type PrependAction struct {
	InsertAction
}

type OrderStatus struct {
	Enumeration
}

type Locksmith struct {
	HomeAndConstructionBusiness
}

type Order struct {
	Discount           interface{}    `json:"discount,omitempty"` // Number, Text
	OrderDelivery      ParcelDelivery `json:"orderDelivery,omitempty"`
	DiscountCurrency   Text           `json:"discountCurrency,omitempty"`
	PaymentMethodId    Text           `json:"paymentMethodId,omitempty"`
	Seller             interface{}    `json:"seller,omitempty"`   // Person, Organization
	Customer           interface{}    `json:"customer,omitempty"` // Organization, Person
	PartOfInvoice      Invoice        `json:"partOfInvoice,omitempty"`
	PaymentMethod      PaymentMethod  `json:"paymentMethod,omitempty"`
	PaymentDue         DateTime       `json:"paymentDue,omitempty"`
	Broker             interface{}    `json:"broker,omitempty"` // Person, Organization
	IsGift             Boolean        `json:"isGift,omitempty"`
	BillingAddress     PostalAddress  `json:"billingAddress,omitempty"`
	DiscountCode       Text           `json:"discountCode,omitempty"`
	OrderDate          DateTime       `json:"orderDate,omitempty"`
	PaymentDueDate     DateTime       `json:"paymentDueDate,omitempty"`
	OrderStatus        OrderStatus    `json:"orderStatus,omitempty"`
	OrderedItem        interface{}    `json:"orderedItem,omitempty"` // OrderItem, Product
	OrderNumber        Text           `json:"orderNumber,omitempty"`
	Merchant           interface{}    `json:"merchant,omitempty"` // Organization, Person
	ConfirmationNumber Text           `json:"confirmationNumber,omitempty"`
	AcceptedOffer      Offer          `json:"acceptedOffer,omitempty"`
	PaymentUrl         URL            `json:"paymentUrl,omitempty"`
	Intangible
}

type HobbyShop struct {
	Store
}

type ContactPoint struct {
	ContactOption     ContactPointOption        `json:"contactOption,omitempty"`
	ProductSupported  interface{}               `json:"productSupported,omitempty"` // Product, Text
	AreaServed        interface{}               `json:"areaServed,omitempty"`       // GeoShape, Text, Place, AdministrativeArea
	HoursAvailable    OpeningHoursSpecification `json:"hoursAvailable,omitempty"`
	ServiceArea       interface{}               `json:"serviceArea,omitempty"` // Place, GeoShape, AdministrativeArea
	FaxNumber         Text                      `json:"faxNumber,omitempty"`
	AvailableLanguage interface{}               `json:"availableLanguage,omitempty"` // Text, Language
	Email             Text                      `json:"email,omitempty"`
	ContactType       Text                      `json:"contactType,omitempty"`
	Telephone         Text                      `json:"telephone,omitempty"`
	StructuredValue
}

type UserInteraction struct {
	Event
}

type GraphicNovel struct {
	BookFormatType
}

type MapCategoryType struct {
	Enumeration
}

type True struct {
	Boolean
}

type HowTo struct {
	Steps         interface{} `json:"steps,omitempty"`  // ItemList, CreativeWork, Text
	Supply        interface{} `json:"supply,omitempty"` // HowToSupply, Text
	TotalTime     Duration    `json:"totalTime,omitempty"`
	Tool          interface{} `json:"tool,omitempty"` // HowToTool, Text
	PerformTime   Duration    `json:"performTime,omitempty"`
	PrepTime      Duration    `json:"prepTime,omitempty"`
	EstimatedCost interface{} `json:"estimatedCost,omitempty"` // Text, MonetaryAmount
	Yield         interface{} `json:"yield,omitempty"`         // QuantitativeValue, Text
	CreativeWork
}

type ParentAudience struct {
	ChildMaxAge Number `json:"childMaxAge,omitempty"`
	ChildMinAge Number `json:"childMinAge,omitempty"`
	PeopleAudience
}

type FDAcategoryC struct {
	DrugPregnancyCategory
}

type Festival struct {
	Event
}

type PresentationDigitalDocument struct {
	DigitalDocument
}

type Accommodation struct {
	FloorSize      QuantitativeValue            `json:"floorSize,omitempty"`
	AmenityFeature LocationFeatureSpecification `json:"amenityFeature,omitempty"`
	NumberOfRooms  interface{}                  `json:"numberOfRooms,omitempty"` // QuantitativeValue, Number
	PetsAllowed    interface{}                  `json:"petsAllowed,omitempty"`   // Text, Boolean
	PermittedUsage Text                         `json:"permittedUsage,omitempty"`
	Place
}

type DoseSchedule struct {
	TargetPopulation Text        `json:"targetPopulation,omitempty"`
	Frequency        Text        `json:"frequency,omitempty"`
	DoseValue        interface{} `json:"doseValue,omitempty"` // QualitativeValue, Number
	DoseUnit         Text        `json:"doseUnit,omitempty"`
	MedicalIntangible
}

type Renal struct {
	MedicalSpecialty
}

type Blog struct {
	BlogPost  BlogPosting `json:"blogPost,omitempty"`
	BlogPosts BlogPosting `json:"blogPosts,omitempty"`
	Issn      Text        `json:"issn,omitempty"`
	CreativeWork
}

type Review struct {
	ItemReviewed Thing  `json:"itemReviewed,omitempty"`
	ReviewBody   Text   `json:"reviewBody,omitempty"`
	ReviewRating Rating `json:"reviewRating,omitempty"`
	CreativeWork
}

type WearAction struct {
	UseAction
}

type Nursing struct {
	MedicalBusiness
	MedicalSpecialty
}

type Motel struct {
	LodgingBusiness
}

type ConfirmAction struct {
	InformAction
}

type TraditionalChinese struct {
	MedicineSystem
}

type FrontWheelDriveConfiguration struct {
	DriveWheelConfigurationValue
}

type Permit struct {
	ValidFor       Duration           `json:"validFor,omitempty"`
	ValidIn        AdministrativeArea `json:"validIn,omitempty"`
	IssuedThrough  Service            `json:"issuedThrough,omitempty"`
	ValidUntil     Date               `json:"validUntil,omitempty"`
	ValidFrom      DateTime           `json:"validFrom,omitempty"`
	PermitAudience Audience           `json:"permitAudience,omitempty"`
	IssuedBy       Organization       `json:"issuedBy,omitempty"`
	Intangible
}

type HomeGoodsStore struct {
	Store
}

type Obstetric struct {
	MedicalBusiness
	MedicalSpecialty
}

type AerobicActivity struct {
	PhysicalActivityCategory
}

type OrganizationRole struct {
	NumberedPosition Number `json:"numberedPosition,omitempty"`
	Role
}

type Distance struct {
	Quantity
}

type WebApplication struct {
	BrowserRequirements Text `json:"browserRequirements,omitempty"`
	SoftwareApplication
}

type ReturnAction struct {
	Recipient interface{} `json:"recipient,omitempty"` // Person, Organization, ContactPoint, Audience
	TransferAction
}

type Action struct {
	Participant  interface{}      `json:"participant,omitempty"` // Person, Organization
	ActionStatus ActionStatusType `json:"actionStatus,omitempty"`
	EndTime      DateTime         `json:"endTime,omitempty"`
	Error        Thing            `json:"error,omitempty"`
	Instrument   Thing            `json:"instrument,omitempty"`
	Location     interface{}      `json:"location,omitempty"` // Place, PostalAddress, Text
	StartTime    DateTime         `json:"startTime,omitempty"`
	Result       Thing            `json:"result,omitempty"`
	Agent        interface{}      `json:"agent,omitempty"` // Organization, Person
	Target       EntryPoint       `json:"target,omitempty"`
	Object       Thing            `json:"object,omitempty"`
	Thing
}

type MedicalSymptom struct {
	MedicalSignOrSymptom
}

type Courthouse struct {
	GovernmentBuilding
}

type ReportageNewsArticle struct {
	NewsArticle
}

type HowToTip struct {
	ListItem
}

type CafeOrCoffeeShop struct {
	FoodEstablishment
}

type Pulmonary struct {
	MedicalSpecialty
}

type OrderDelivered struct {
	OrderStatus
}

type ContactPointOption struct {
	Enumeration
}

type ReviewAction struct {
	ResultReview Review `json:"resultReview,omitempty"`
	AssessAction
}

type WPFooter struct {
	WebPageElement
}

type MedicalTest struct {
	NormalRange    interface{}      `json:"normalRange,omitempty"` // Text, MedicalEnumeration
	AffectedBy     Drug             `json:"affectedBy,omitempty"`
	UsesDevice     MedicalDevice    `json:"usesDevice,omitempty"`
	UsedToDiagnose MedicalCondition `json:"usedToDiagnose,omitempty"`
	SignDetected   MedicalSign      `json:"signDetected,omitempty"`
	MedicalEntity
}

type MenuSection struct {
	HasMenuItem    MenuItem    `json:"hasMenuItem,omitempty"`
	HasMenuSection MenuSection `json:"hasMenuSection,omitempty"`
	CreativeWork
}

type MusicVenue struct {
	CivicStructure
}

type FoodEstablishmentReservation struct {
	PartySize interface{} `json:"partySize,omitempty"` // Integer, QuantitativeValue
	EndTime   DateTime    `json:"endTime,omitempty"`
	StartTime DateTime    `json:"startTime,omitempty"`
	Reservation
}

type OrganizeAction struct {
	Action
}

type Friday struct {
	DayOfWeek
}

type Role struct {
	StartDate     interface{} `json:"startDate,omitempty"`     // DateTime, Date
	RoleName      interface{} `json:"roleName,omitempty"`      // Text, URL
	NamedPosition interface{} `json:"namedPosition,omitempty"` // URL, Text
	EndDate       interface{} `json:"endDate,omitempty"`       // Date, DateTime
	Intangible
}

type Saturday struct {
	DayOfWeek
}

type TieAction struct {
	AchieveAction
}

type Virus struct {
	InfectiousAgentClass
}

type PreOrderAction struct {
	TradeAction
}

type DefenceEstablishment struct {
	GovernmentBuilding
}

type BreadcrumbList struct {
	ItemList
}

type ArtGallery struct {
	EntertainmentBusiness
}

type Aquarium struct {
	CivicStructure
}

type LaboratoryScience struct {
	MedicalSpecialty
}

type TradeAction struct {
	PriceSpecification PriceSpecification `json:"priceSpecification,omitempty"`
	Price              interface{}        `json:"price,omitempty"` // Text, Number
	Action
}

type Abdomen struct {
	PhysicalExam
}

type EmergencyService struct {
	LocalBusiness
}

type Menu struct {
	HasMenuSection MenuSection `json:"hasMenuSection,omitempty"`
	HasMenuItem    MenuItem    `json:"hasMenuItem,omitempty"`
	CreativeWork
}

type DamagedCondition struct {
	OfferItemCondition
}

type Male struct {
	GenderType
}

type PathologyTest struct {
	TissueSample Text `json:"tissueSample,omitempty"`
	MedicalTest
}

type Bakery struct {
	FoodEstablishment
}

type CurrencyConversionService struct {
	FinancialProduct
}

type XPathType struct {
	Text
}

type Volcano struct {
	Landform
}

type CoOp struct {
	GamePlayMode
}

type OrderProblem struct {
	OrderStatus
}

type TennisComplex struct {
	SportsActivityLocation
}

type SportsTeam struct {
	Coach   Person `json:"coach,omitempty"`
	Athlete Person `json:"athlete,omitempty"`
	SportsOrganization
}

type BusinessEntityType struct {
	Enumeration
}

type UserCheckins struct {
	UserInteraction
}

type MedicalScholarlyArticle struct {
	PublicationType Text `json:"publicationType,omitempty"`
	ScholarlyArticle
}

type Message struct {
	DateReceived      DateTime     `json:"dateReceived,omitempty"`
	ToRecipient       interface{}  `json:"toRecipient,omitempty"` // ContactPoint, Audience, Organization, Person
	DateRead          DateTime     `json:"dateRead,omitempty"`
	Sender            interface{}  `json:"sender,omitempty"`    // Person, Audience, Organization
	Recipient         interface{}  `json:"recipient,omitempty"` // Person, Organization, ContactPoint, Audience
	DateSent          DateTime     `json:"dateSent,omitempty"`
	CcRecipient       interface{}  `json:"ccRecipient,omitempty"`  // Organization, ContactPoint, Person
	BccRecipient      interface{}  `json:"bccRecipient,omitempty"` // Person, Organization, ContactPoint
	MessageAttachment CreativeWork `json:"messageAttachment,omitempty"`
	CreativeWork
}

type ReservationPending struct {
	ReservationStatusType
}

type GiveAction struct {
	Recipient interface{} `json:"recipient,omitempty"` // Person, Organization, ContactPoint, Audience
	TransferAction
}

type Store struct {
	LocalBusiness
}

type Hotel struct {
	LodgingBusiness
}

type ChildrensEvent struct {
	Event
}

type ExhibitionEvent struct {
	Event
}

type NightClub struct {
	EntertainmentBusiness
}

type EventVenue struct {
	CivicStructure
}

type Embassy struct {
	GovernmentBuilding
}

type StrengthTraining struct {
	PhysicalActivityCategory
}

type MotorcycleRepair struct {
	AutomotiveBusiness
}

type WritePermission struct {
	DigitalDocumentPermissionType
}

type Throat struct {
	PhysicalExam
}

type Attorney struct {
	LegalService
}

type BusOrCoach struct {
	AcrissCode Text              `json:"acrissCode,omitempty"`
	RoofLoad   QuantitativeValue `json:"roofLoad,omitempty"`
	Vehicle
}

type ReceiveAction struct {
	DeliveryMethod DeliveryMethod `json:"deliveryMethod,omitempty"`
	Sender         interface{}    `json:"sender,omitempty"` // Person, Audience, Organization
	TransferAction
}

type MovingCompany struct {
	HomeAndConstructionBusiness
}

type PrescriptionOnly struct {
	DrugPrescriptionStatus
}

type InsertAction struct {
	ToLocation Place `json:"toLocation,omitempty"`
	AddAction
}

type Newspaper struct {
	Periodical
}

type MedicalTrialDesign struct {
	Enumeration
	MedicalEnumeration
}

type URL struct {
	Text
}

type HowToTool struct {
	HowToItem
}

type AskAction struct {
	Question Question `json:"question,omitempty"`
	CommunicateAction
}

type Homeopathic struct {
	MedicineSystem
}

type Thursday struct {
	DayOfWeek
}

type BroadcastEvent struct {
	IsLiveBroadcast  Boolean `json:"isLiveBroadcast,omitempty"`
	BroadcastOfEvent Event   `json:"broadcastOfEvent,omitempty"`
	VideoFormat      Text    `json:"videoFormat,omitempty"`
	PublicationEvent
}

type GenderType struct {
	Enumeration
}

type BefriendAction struct {
	InteractAction
}

type BusStop struct {
	CivicStructure
}

type Sculpture struct {
	CreativeWork
}

type Photograph struct {
	CreativeWork
}

type EmploymentAgency struct {
	LocalBusiness
}

type Campground struct {
	LodgingBusiness
	CivicStructure
}

type HealthPlanNetwork struct {
	HealthPlanCostSharing Boolean `json:"healthPlanCostSharing,omitempty"`
	HealthPlanNetworkTier Text    `json:"healthPlanNetworkTier,omitempty"`
	HealthPlanNetworkId   Text    `json:"healthPlanNetworkId,omitempty"`
	Intangible
}

type Balance struct {
	PhysicalActivityCategory
}

type AnatomicalStructure struct {
	RelatedCondition          MedicalCondition    `json:"relatedCondition,omitempty"`
	SubStructure              AnatomicalStructure `json:"subStructure,omitempty"`
	Diagram                   ImageObject         `json:"diagram,omitempty"`
	RelatedTherapy            MedicalTherapy      `json:"relatedTherapy,omitempty"`
	PartOfSystem              AnatomicalSystem    `json:"partOfSystem,omitempty"`
	Function                  Text                `json:"function,omitempty"`
	BodyLocation              Text                `json:"bodyLocation,omitempty"`
	AssociatedPathophysiology Text                `json:"associatedPathophysiology,omitempty"`
	ConnectedTo               AnatomicalStructure `json:"connectedTo,omitempty"`
	MedicalEntity
}

type AssignAction struct {
	AllocateAction
}

type SiteNavigationElement struct {
	WebPageElement
}

type CrossSectional struct {
	MedicalObservationalStudyDesign
}

type MonetaryAmount struct {
	MinValue     Number      `json:"minValue,omitempty"`
	ValidFrom    DateTime    `json:"validFrom,omitempty"`
	ValidThrough DateTime    `json:"validThrough,omitempty"`
	Value        interface{} `json:"value,omitempty"` // StructuredValue, Number, Boolean, Text
	MaxValue     Number      `json:"maxValue,omitempty"`
	StructuredValue
}

type DigitalDocumentPermission struct {
	Grantee        interface{}                   `json:"grantee,omitempty"` // Audience, ContactPoint, Organization, Person
	PermissionType DigitalDocumentPermissionType `json:"permissionType,omitempty"`
	Intangible
}

type Psychiatric struct {
	MedicalBusiness
	MedicalSpecialty
}

type CssSelectorType struct {
	Text
}

type TreatmentIndication struct {
	MedicalIndication
}

type PaymentAutomaticallyApplied struct {
	PaymentStatusType
}

type SurgicalProcedure struct {
	MedicalProcedure
}

type SportsActivityLocation struct {
	LocalBusiness
}

type CategoryCodeSet struct {
	HasCategoryCode CategoryCode `json:"hasCategoryCode,omitempty"`
	CreativeWork
}

type EventReservation struct {
	Reservation
}

type FoodEvent struct {
	Event
}

type ServiceChannel struct {
	ProcessingTime       Duration      `json:"processingTime,omitempty"`
	ServiceSmsNumber     ContactPoint  `json:"serviceSmsNumber,omitempty"`
	ProvidesService      Service       `json:"providesService,omitempty"`
	ServicePostalAddress PostalAddress `json:"servicePostalAddress,omitempty"`
	ServiceLocation      Place         `json:"serviceLocation,omitempty"`
	AvailableLanguage    interface{}   `json:"availableLanguage,omitempty"` // Text, Language
	ServicePhone         ContactPoint  `json:"servicePhone,omitempty"`
	ServiceUrl           URL           `json:"serviceUrl,omitempty"`
	Intangible
}

type HalalDiet struct {
	RestrictedDiet
}

type ProgramMembership struct {
	MembershipNumber    Text         `json:"membershipNumber,omitempty"`
	Member              interface{}  `json:"member,omitempty"`  // Organization, Person
	Members             interface{}  `json:"members,omitempty"` // Person, Organization
	HostingOrganization Organization `json:"hostingOrganization,omitempty"`
	ProgramName         Text         `json:"programName,omitempty"`
	Intangible
}

type MedicalTherapy struct {
	DuplicateTherapy      MedicalTherapy `json:"duplicateTherapy,omitempty"`
	Contraindication      interface{}    `json:"contraindication,omitempty"` // Text, MedicalContraindication
	SeriousAdverseOutcome MedicalEntity  `json:"seriousAdverseOutcome,omitempty"`
	TherapeuticProcedure
}

type DepartmentStore struct {
	Store
}

type LockerDelivery struct {
	DeliveryMethod
}

type HinduDiet struct {
	RestrictedDiet
}

type MedicalEvidenceLevel struct {
	MedicalEnumeration
}

type Podiatric struct {
	MedicalBusiness
	MedicalSpecialty
}

type LendAction struct {
	Borrower Person `json:"borrower,omitempty"`
	TransferAction
}

type EventSeries struct {
	Event
}

type Physician struct {
	AvailableService    interface{}      `json:"availableService,omitempty"` // MedicalTest, MedicalProcedure, MedicalTherapy
	MedicalSpecialty    MedicalSpecialty `json:"medicalSpecialty,omitempty"`
	HospitalAffiliation Hospital         `json:"hospitalAffiliation,omitempty"`
	MedicalBusiness
}

type Organization struct {
	GlobalLocationNumber     Text              `json:"globalLocationNumber,omitempty"`
	Members                  interface{}       `json:"members,omitempty"` // Person, Organization
	AggregateRating          AggregateRating   `json:"aggregateRating,omitempty"`
	Telephone                Text              `json:"telephone,omitempty"`
	HasPOS                   Place             `json:"hasPOS,omitempty"`
	Naics                    Text              `json:"naics,omitempty"`
	Founder                  Person            `json:"founder,omitempty"`
	EthicsPolicy             interface{}       `json:"ethicsPolicy,omitempty"` // URL, CreativeWork
	Employees                Person            `json:"employees,omitempty"`
	Email                    Text              `json:"email,omitempty"`
	UnnamedSourcesPolicy     interface{}       `json:"unnamedSourcesPolicy,omitempty"` // URL, CreativeWork
	IsicV4                   Text              `json:"isicV4,omitempty"`
	ServiceArea              interface{}       `json:"serviceArea,omitempty"`              // Place, GeoShape, AdministrativeArea
	Funder                   interface{}       `json:"funder,omitempty"`                   // Person, Organization
	ActionableFeedbackPolicy interface{}       `json:"actionableFeedbackPolicy,omitempty"` // URL, CreativeWork
	MakesOffer               Offer             `json:"makesOffer,omitempty"`
	ContactPoints            ContactPoint      `json:"contactPoints,omitempty"`
	FoundingLocation         Place             `json:"foundingLocation,omitempty"`
	NumberOfEmployees        QuantitativeValue `json:"numberOfEmployees,omitempty"`
	AreaServed               interface{}       `json:"areaServed,omitempty"` // GeoShape, Text, Place, AdministrativeArea
	Location                 interface{}       `json:"location,omitempty"`   // Place, PostalAddress, Text
	DissolutionDate          Date              `json:"dissolutionDate,omitempty"`
	MemberOf                 interface{}       `json:"memberOf,omitempty"`             // Organization, ProgramMembership
	Owns                     interface{}       `json:"owns,omitempty"`                 // Product, OwnershipInfo
	PublishingPrinciples     interface{}       `json:"publishingPrinciples,omitempty"` // CreativeWork, URL
	DiversityPolicy          interface{}       `json:"diversityPolicy,omitempty"`      // CreativeWork, URL
	VatID                    Text              `json:"vatID,omitempty"`
	FaxNumber                Text              `json:"faxNumber,omitempty"`
	LegalName                Text              `json:"legalName,omitempty"`
	FoundingDate             Date              `json:"foundingDate,omitempty"`
	SubOrganization          Organization      `json:"subOrganization,omitempty"`
	Events                   Event             `json:"events,omitempty"`
	ParentOrganization       Organization      `json:"parentOrganization,omitempty"`
	Reviews                  Review            `json:"reviews,omitempty"`
	Event                    Event             `json:"event,omitempty"`
	Address                  interface{}       `json:"address,omitempty"` // PostalAddress, Text
	Award                    Text              `json:"award,omitempty"`
	Member                   interface{}       `json:"member,omitempty"` // Organization, Person
	Awards                   Text              `json:"awards,omitempty"`
	TaxID                    Text              `json:"taxID,omitempty"`
	Seeks                    Demand            `json:"seeks,omitempty"`
	Alumni                   Person            `json:"alumni,omitempty"`
	LeiCode                  Text              `json:"leiCode,omitempty"`
	Department               Organization      `json:"department,omitempty"`
	Review                   Review            `json:"review,omitempty"`
	CorrectionsPolicy        interface{}       `json:"correctionsPolicy,omitempty"` // URL, CreativeWork
	Duns                     Text              `json:"duns,omitempty"`
	Logo                     interface{}       `json:"logo,omitempty"` // URL, ImageObject
	Founders                 Person            `json:"founders,omitempty"`
	HasOfferCatalog          OfferCatalog      `json:"hasOfferCatalog,omitempty"`
	Employee                 Person            `json:"employee,omitempty"`
	Brand                    interface{}       `json:"brand,omitempty"` // Organization, Brand
	Thing
}

type Bridge struct {
	CivicStructure
}

type DJMixAlbum struct {
	MusicAlbumProductionType
}

type AgreeAction struct {
	ReactAction
}

type Dentistry struct {
	MedicalSpecialty
}

type ConvenienceStore struct {
	Store
}

type Observational struct {
	MedicalObservationalStudyDesign
}

type Chiropractic struct {
	MedicineSystem
}

type PoliceStation struct {
	EmergencyService
	CivicStructure
}

type ItemListOrderType struct {
	Enumeration
}

type BloodTest struct {
	MedicalTest
}

type Room struct {
	Accommodation
}

type Question struct {
	SuggestedAnswer Answer  `json:"suggestedAnswer,omitempty"`
	AnswerCount     Integer `json:"answerCount,omitempty"`
	UpvoteCount     Integer `json:"upvoteCount,omitempty"`
	DownvoteCount   Integer `json:"downvoteCount,omitempty"`
	AcceptedAnswer  Answer  `json:"acceptedAnswer,omitempty"`
	CreativeWork
}

type DDxElement struct {
	Diagnosis          MedicalCondition     `json:"diagnosis,omitempty"`
	DistinguishingSign MedicalSignOrSymptom `json:"distinguishingSign,omitempty"`
	MedicalIntangible
}

type Fungus struct {
	InfectiousAgentClass
}

type Dentist struct {
	MedicalBusiness
}

type StupidType struct {
	StupidProperty QuantitativeValue `json:"stupidProperty,omitempty"`
	Thing
}

type UnRegisterAction struct {
	InteractAction
}

type BookStore struct {
	Store
}

type ListenAction struct {
	ConsumeAction
}

type EntryPoint struct {
	ActionApplication SoftwareApplication `json:"actionApplication,omitempty"`
	EncodingType      Text                `json:"encodingType,omitempty"`
	Application       SoftwareApplication `json:"application,omitempty"`
	HttpMethod        Text                `json:"httpMethod,omitempty"`
	UrlTemplate       Text                `json:"urlTemplate,omitempty"`
	ContentType       Text                `json:"contentType,omitempty"`
	ActionPlatform    interface{}         `json:"actionPlatform,omitempty"` // Text, URL
	Intangible
}

type Anesthesia struct {
	MedicalSpecialty
}

type XRay struct {
	MedicalImagingTechnique
}

type Surgical struct {
	MedicalSpecialty
}

type LegislativeBuilding struct {
	GovernmentBuilding
}

type GasStation struct {
	AutomotiveBusiness
}

type Infectious struct {
	MedicalSpecialty
}

type TravelAction struct {
	Distance Distance `json:"distance,omitempty"`
	MoveAction
}

type TollFree struct {
	ContactPointOption
}

type ComputerStore struct {
	Store
}

type MedicalRiskCalculator struct {
	MedicalRiskEstimator
}

type Product struct {
	Mpn                       Text               `json:"mpn,omitempty"`
	ProductionDate            Date               `json:"productionDate,omitempty"`
	IsConsumableFor           Product            `json:"isConsumableFor,omitempty"`
	Gtin12                    Text               `json:"gtin12,omitempty"`
	Offers                    Offer              `json:"offers,omitempty"`
	Award                     Text               `json:"award,omitempty"`
	Brand                     interface{}        `json:"brand,omitempty"` // Organization, Brand
	Audience                  Audience           `json:"audience,omitempty"`
	Width                     interface{}        `json:"width,omitempty"` // Distance, QuantitativeValue
	Gtin8                     Text               `json:"gtin8,omitempty"`
	Material                  interface{}        `json:"material,omitempty"` // Text, URL, Product
	ReleaseDate               Date               `json:"releaseDate,omitempty"`
	Model                     interface{}        `json:"model,omitempty"` // ProductModel, Text
	Awards                    Text               `json:"awards,omitempty"`
	Weight                    QuantitativeValue  `json:"weight,omitempty"`
	AggregateRating           AggregateRating    `json:"aggregateRating,omitempty"`
	Color                     Text               `json:"color,omitempty"`
	Sku                       Text               `json:"sku,omitempty"`
	Review                    Review             `json:"review,omitempty"`
	IsSimilarTo               interface{}        `json:"isSimilarTo,omitempty"` // Product, Service
	IsAccessoryOrSparePartFor Product            `json:"isAccessoryOrSparePartFor,omitempty"`
	IsRelatedTo               interface{}        `json:"isRelatedTo,omitempty"` // Service, Product
	PurchaseDate              Date               `json:"purchaseDate,omitempty"`
	Gtin13                    Text               `json:"gtin13,omitempty"`
	ItemCondition             OfferItemCondition `json:"itemCondition,omitempty"`
	Reviews                   Review             `json:"reviews,omitempty"`
	ProductID                 Text               `json:"productID,omitempty"`
	AdditionalProperty        PropertyValue      `json:"additionalProperty,omitempty"`
	Gtin14                    Text               `json:"gtin14,omitempty"`
	Logo                      interface{}        `json:"logo,omitempty"`   // URL, ImageObject
	Height                    interface{}        `json:"height,omitempty"` // Distance, QuantitativeValue
	Depth                     interface{}        `json:"depth,omitempty"`  // QuantitativeValue, Distance
	Thing
}

type MusicRecording struct {
	InAlbum     MusicAlbum       `json:"inAlbum,omitempty"`
	InPlaylist  MusicPlaylist    `json:"inPlaylist,omitempty"`
	ByArtist    MusicGroup       `json:"byArtist,omitempty"`
	RecordingOf MusicComposition `json:"recordingOf,omitempty"`
	IsrcCode    Text             `json:"isrcCode,omitempty"`
	CreativeWork
}

type AdvertiserContentArticle struct {
	Article
}

type DiabeticDiet struct {
	RestrictedDiet
}

type InsuranceAgency struct {
	FinancialService
}

type OrderInTransit struct {
	OrderStatus
}

type RentalVehicleUsage struct {
	CarUsageType
}

type TrainStation struct {
	CivicStructure
}

type WebSite struct {
	Issn Text `json:"issn,omitempty"`
	CreativeWork
}

type GeoCircle struct {
	GeoMidpoint GeoCoordinates `json:"geoMidpoint,omitempty"`
	GeoRadius   interface{}    `json:"geoRadius,omitempty"` // Distance, Number, Text
	GeoShape
}

type ToyStore struct {
	Store
}

type UnitPriceSpecification struct {
	UnitText          Text              `json:"unitText,omitempty"`
	ReferenceQuantity QuantitativeValue `json:"referenceQuantity,omitempty"`
	UnitCode          interface{}       `json:"unitCode,omitempty"` // URL, Text
	PriceType         Text              `json:"priceType,omitempty"`
	BillingIncrement  Number            `json:"billingIncrement,omitempty"`
	PriceSpecification
}

type InternationalTrial struct {
	MedicalTrialDesign
}

type ReservationHold struct {
	ReservationStatusType
}

type LoseAction struct {
	Winner Person `json:"winner,omitempty"`
	AchieveAction
}

type Intangible struct {
	Thing
}

type ReimbursementCap struct {
	DrugCostCategory
}

type BroadcastService struct {
	BroadcastDisplayName Text             `json:"broadcastDisplayName,omitempty"`
	VideoFormat          Text             `json:"videoFormat,omitempty"`
	BroadcastTimezone    Text             `json:"broadcastTimezone,omitempty"`
	BroadcastFrequency   interface{}      `json:"broadcastFrequency,omitempty"` // Text, BroadcastFrequencySpecification
	Area                 Place            `json:"area,omitempty"`
	BroadcastAffiliateOf Organization     `json:"broadcastAffiliateOf,omitempty"`
	ParentService        BroadcastService `json:"parentService,omitempty"`
	HasBroadcastChannel  BroadcastChannel `json:"hasBroadcastChannel,omitempty"`
	Broadcaster          Organization     `json:"broadcaster,omitempty"`
	Service
}

type RsvpResponseMaybe struct {
	RsvpResponseType
}

type AllocateAction struct {
	Purpose interface{} `json:"purpose,omitempty"` // Thing, MedicalDevicePurpose
	OrganizeAction
}

type ExchangeRateSpecification struct {
	ExchangeRateSpread  interface{}            `json:"exchangeRateSpread,omitempty"` // Number, MonetaryAmount
	CurrentExchangeRate UnitPriceSpecification `json:"currentExchangeRate,omitempty"`
	Currency            interface{}            `json:"currency,omitempty"` //
	StructuredValue
}

type Genitourinary struct {
	PhysicalExam
}

type LibrarySystem struct {
	Organization
}

type GeneralContractor struct {
	HomeAndConstructionBusiness
}

type ApprovedIndication struct {
	MedicalIndication
}

type TaxiService struct {
	Service
}

type Oncologic struct {
	MedicalBusiness
	MedicalSpecialty
}

type NoninvasiveProcedure struct {
	MedicalProcedureType
}

type MarryAction struct {
	InteractAction
}

type LimitedAvailability struct {
	ItemAvailability
}

type HotelRoom struct {
	Occupancy QuantitativeValue `json:"occupancy,omitempty"`
	Room
}

type MedicalWebPage struct {
	Aspect Text `json:"aspect,omitempty"`
	WebPage
}

type ReservationCancelled struct {
	ReservationStatusType
}

type SearchAction struct {
	Query Text `json:"query,omitempty"`
	Action
}

type TherapeuticProcedure struct {
	Drug           Drug              `json:"drug,omitempty"`
	Indication     MedicalIndication `json:"indication,omitempty"`
	DoseSchedule   DoseSchedule      `json:"doseSchedule,omitempty"`
	AdverseOutcome MedicalEntity     `json:"adverseOutcome,omitempty"`
	MedicalProcedure
}

type LodgingReservation struct {
	NumAdults              interface{} `json:"numAdults,omitempty"`       // QuantitativeValue, Integer
	NumChildren            interface{} `json:"numChildren,omitempty"`     // Integer, QuantitativeValue
	LodgingUnitType        interface{} `json:"lodgingUnitType,omitempty"` // QualitativeValue, Text
	CheckoutTime           DateTime    `json:"checkoutTime,omitempty"`
	CheckinTime            DateTime    `json:"checkinTime,omitempty"`
	LodgingUnitDescription Text        `json:"lodgingUnitDescription,omitempty"`
	Reservation
}

type EndorseAction struct {
	Endorsee interface{} `json:"endorsee,omitempty"` // Person, Organization
	ReactAction
}

type SatiricalArticle struct {
	Article
}

type SeaBodyOfWater struct {
	BodyOfWater
}

type ParcelService struct {
	DeliveryMethod
}

type MedicalOrganization struct {
	IsAcceptingNewPatients Boolean          `json:"isAcceptingNewPatients,omitempty"`
	HealthPlanNetworkId    Text             `json:"healthPlanNetworkId,omitempty"`
	MedicalSpecialty       MedicalSpecialty `json:"medicalSpecialty,omitempty"`
	Organization
}

type StudioAlbum struct {
	MusicAlbumProductionType
}

type PaymentStatusType struct {
	Enumeration
}

type Collection struct {
	CreativeWork
}

type DanceEvent struct {
	Event
}

type PaymentMethod struct {
	Enumeration
}

type Appearance struct {
	PhysicalExam
}

type Property struct {
	DomainIncludes Class       `json:"domainIncludes,omitempty"`
	InverseOf      Property    `json:"inverseOf,omitempty"`
	SupersededBy   interface{} `json:"supersededBy,omitempty"` // Property, Enumeration, Class
	RangeIncludes  Class       `json:"rangeIncludes,omitempty"`
	Category       interface{} `json:"category,omitempty"` //
	Intangible
}

type HomeAndConstructionBusiness struct {
	LocalBusiness
}

type CheckInAction struct {
	CommunicateAction
}

type HowToItem struct {
	RequiredQuantity interface{} `json:"requiredQuantity,omitempty"` // Number, QuantitativeValue, Text
	ListItem
}

type SuspendAction struct {
	ControlAction
}

type Offer struct {
	ItemCondition             OfferItemCondition  `json:"itemCondition,omitempty"`
	AcceptedPaymentMethod     interface{}         `json:"acceptedPaymentMethod,omitempty"` // LoanOrCredit, PaymentMethod
	Seller                    interface{}         `json:"seller,omitempty"`                // Person, Organization
	ValidThrough              DateTime            `json:"validThrough,omitempty"`
	AvailabilityStarts        DateTime            `json:"availabilityStarts,omitempty"`
	Reviews                   Review              `json:"reviews,omitempty"`
	PriceCurrency             Text                `json:"priceCurrency,omitempty"`
	IncludesObject            TypeAndQuantityNode `json:"includesObject,omitempty"`
	ItemOffered               interface{}         `json:"itemOffered,omitempty"` // Product, Service
	AggregateRating           AggregateRating     `json:"aggregateRating,omitempty"`
	EligibleTransactionVolume PriceSpecification  `json:"eligibleTransactionVolume,omitempty"`
	AddOn                     Offer               `json:"addOn,omitempty"`
	AvailabilityEnds          DateTime            `json:"availabilityEnds,omitempty"`
	AdvanceBookingRequirement QuantitativeValue   `json:"advanceBookingRequirement,omitempty"`
	OfferedBy                 interface{}         `json:"offeredBy,omitempty"` // Organization, Person
	Availability              ItemAvailability    `json:"availability,omitempty"`
	EligibleDuration          QuantitativeValue   `json:"eligibleDuration,omitempty"`
	PriceSpecification        PriceSpecification  `json:"priceSpecification,omitempty"`
	DeliveryLeadTime          QuantitativeValue   `json:"deliveryLeadTime,omitempty"`
	Gtin13                    Text                `json:"gtin13,omitempty"`
	ValidFrom                 DateTime            `json:"validFrom,omitempty"`
	Gtin14                    Text                `json:"gtin14,omitempty"`
	AreaServed                interface{}         `json:"areaServed,omitempty"` // GeoShape, Text, Place, AdministrativeArea
	Sku                       Text                `json:"sku,omitempty"`
	Mpn                       Text                `json:"mpn,omitempty"`
	SerialNumber              Text                `json:"serialNumber,omitempty"`
	BusinessFunction          BusinessFunction    `json:"businessFunction,omitempty"`
	EligibleCustomerType      BusinessEntityType  `json:"eligibleCustomerType,omitempty"`
	InventoryLevel            QuantitativeValue   `json:"inventoryLevel,omitempty"`
	EligibleRegion            interface{}         `json:"eligibleRegion,omitempty"` // GeoShape, Text, Place
	Gtin12                    Text                `json:"gtin12,omitempty"`
	IneligibleRegion          interface{}         `json:"ineligibleRegion,omitempty"` // Place, GeoShape, Text
	AvailableAtOrFrom         Place               `json:"availableAtOrFrom,omitempty"`
	Warranty                  WarrantyPromise     `json:"warranty,omitempty"`
	Gtin8                     Text                `json:"gtin8,omitempty"`
	AvailableDeliveryMethod   DeliveryMethod      `json:"availableDeliveryMethod,omitempty"`
	Review                    Review              `json:"review,omitempty"`
	EligibleQuantity          QuantitativeValue   `json:"eligibleQuantity,omitempty"`
	Price                     interface{}         `json:"price,omitempty"` // Text, Number
	PriceValidUntil           Date                `json:"priceValidUntil,omitempty"`
	Intangible
}

type LymphaticVessel struct {
	RunsTo         Vessel      `json:"runsTo,omitempty"`
	OriginatesFrom Vessel      `json:"originatesFrom,omitempty"`
	RegionDrained  interface{} `json:"regionDrained,omitempty"` // AnatomicalSystem, AnatomicalStructure
	Vessel
}

type MultiPlayer struct {
	GamePlayMode
}

type Church struct {
	PlaceOfWorship
}

type TipAction struct {
	Recipient interface{} `json:"recipient,omitempty"` // Person, Organization, ContactPoint, Audience
	TradeAction
}

type EmployeeRole struct {
	BaseSalary     interface{} `json:"baseSalary,omitempty"` // PriceSpecification, MonetaryAmount, Number
	SalaryCurrency Text        `json:"salaryCurrency,omitempty"`
	OrganizationRole
}

type DiagnosticProcedure struct {
	MedicalProcedure
}

type WPHeader struct {
	WebPageElement
}

type Drug struct {
	ActiveIngredient              Text                  `json:"activeIngredient,omitempty"`
	Overdosage                    Text                  `json:"overdosage,omitempty"`
	PrescriptionStatus            interface{}           `json:"prescriptionStatus,omitempty"` // DrugPrescriptionStatus, Text
	Warning                       interface{}           `json:"warning,omitempty"`            // URL, Text
	MaximumIntake                 MaximumDoseSchedule   `json:"maximumIntake,omitempty"`
	MechanismOfAction             Text                  `json:"mechanismOfAction,omitempty"`
	DrugUnit                      Text                  `json:"drugUnit,omitempty"`
	BreastfeedingWarning          Text                  `json:"breastfeedingWarning,omitempty"`
	FoodWarning                   Text                  `json:"foodWarning,omitempty"`
	NonProprietaryName            Text                  `json:"nonProprietaryName,omitempty"`
	Rxcui                         Text                  `json:"rxcui,omitempty"`
	Cost                          DrugCost              `json:"cost,omitempty"`
	ClincalPharmacology           Text                  `json:"clincalPharmacology,omitempty"`
	AvailableStrength             DrugStrength          `json:"availableStrength,omitempty"`
	ClinicalPharmacology          Text                  `json:"clinicalPharmacology,omitempty"`
	PregnancyWarning              Text                  `json:"pregnancyWarning,omitempty"`
	IncludedInHealthInsurancePlan HealthInsurancePlan   `json:"includedInHealthInsurancePlan,omitempty"`
	LabelDetails                  URL                   `json:"labelDetails,omitempty"`
	IsAvailableGenerically        Boolean               `json:"isAvailableGenerically,omitempty"`
	DosageForm                    Text                  `json:"dosageForm,omitempty"`
	AlcoholWarning                Text                  `json:"alcoholWarning,omitempty"`
	PrescribingInfo               URL                   `json:"prescribingInfo,omitempty"`
	ProprietaryName               Text                  `json:"proprietaryName,omitempty"`
	Manufacturer                  interface{}           `json:"manufacturer,omitempty"` //
	InteractingDrug               Drug                  `json:"interactingDrug,omitempty"`
	DrugClass                     DrugClass             `json:"drugClass,omitempty"`
	DoseSchedule                  DoseSchedule          `json:"doseSchedule,omitempty"`
	RelatedDrug                   Drug                  `json:"relatedDrug,omitempty"`
	LegalStatus                   interface{}           `json:"legalStatus,omitempty"` // DrugLegalStatus, Text, MedicalEnumeration
	AdministrationRoute           Text                  `json:"administrationRoute,omitempty"`
	IsProprietary                 Boolean               `json:"isProprietary,omitempty"`
	PregnancyCategory             DrugPregnancyCategory `json:"pregnancyCategory,omitempty"`
	Substance
}

type JobPosting struct {
	EstimatedSalary        interface{}  `json:"estimatedSalary,omitempty"` // PriceSpecification, Number, MonetaryAmount
	EducationRequirements  Text         `json:"educationRequirements,omitempty"`
	EmploymentType         Text         `json:"employmentType,omitempty"`
	WorkHours              Text         `json:"workHours,omitempty"`
	SalaryCurrency         Text         `json:"salaryCurrency,omitempty"`
	Industry               Text         `json:"industry,omitempty"`
	Title                  Text         `json:"title,omitempty"`
	Responsibilities       Text         `json:"responsibilities,omitempty"`
	OccupationalCategory   Text         `json:"occupationalCategory,omitempty"`
	Skills                 Text         `json:"skills,omitempty"`
	DatePosted             Date         `json:"datePosted,omitempty"`
	Qualifications         Text         `json:"qualifications,omitempty"`
	ValidThrough           DateTime     `json:"validThrough,omitempty"`
	SpecialCommitments     Text         `json:"specialCommitments,omitempty"`
	ExperienceRequirements Text         `json:"experienceRequirements,omitempty"`
	Incentives             Text         `json:"incentives,omitempty"`
	JobLocation            Place        `json:"jobLocation,omitempty"`
	Benefits               Text         `json:"benefits,omitempty"`
	JobBenefits            Text         `json:"jobBenefits,omitempty"`
	BaseSalary             interface{}  `json:"baseSalary,omitempty"` // PriceSpecification, MonetaryAmount, Number
	IncentiveCompensation  Text         `json:"incentiveCompensation,omitempty"`
	HiringOrganization     Organization `json:"hiringOrganization,omitempty"`
	Intangible
}

type BookmarkAction struct {
	OrganizeAction
}

type RespiratoryTherapy struct {
	MedicalTherapy
	MedicalSpecialty
}

type HairSalon struct {
	HealthAndBeautyBusiness
}

type DataFeed struct {
	DataFeedElement interface{} `json:"dataFeedElement,omitempty"` // Text, Thing, DataFeedItem
	Dataset
}

type Suite struct {
	NumberOfRooms interface{}       `json:"numberOfRooms,omitempty"` // QuantitativeValue, Number
	Occupancy     QuantitativeValue `json:"occupancy,omitempty"`
	Accommodation
}

type GroupBoardingPolicy struct {
	BoardingPolicyType
}

type EatAction struct {
	ConsumeAction
}

type ReviewNewsArticle struct {
	CriticReview
	NewsArticle
}

type CheckAction struct {
	FindAction
}

type CheckOutAction struct {
	CommunicateAction
}

type DemoAlbum struct {
	MusicAlbumProductionType
}

type Park struct {
	CivicStructure
}

type Conversation struct {
	CreativeWork
}

type Class struct {
	SupersededBy interface{} `json:"supersededBy,omitempty"` // Property, Enumeration, Class
	Category     interface{} `json:"category,omitempty"`     //
	Intangible
}

type CardiovascularExam struct {
	PhysicalExam
}

type DiagnosticLab struct {
	AvailableTest MedicalTest `json:"availableTest,omitempty"`
	MedicalOrganization
}

type MaximumDoseSchedule struct {
	DoseSchedule
}

type LeisureTimeActivity struct {
	PhysicalActivityCategory
}

type SubscribeAction struct {
	InteractAction
}

type OrderPickupAvailable struct {
	OrderStatus
}

type TravelAgency struct {
	LocalBusiness
}

type MedicalAudience struct {
	MedicalEnumeration
	PeopleAudience
	Audience
}

type Withdrawn struct {
	MedicalStudyStatus
}

type GovernmentOffice struct {
	LocalBusiness
}

type HighSchool struct {
	EducationalOrganization
}

type ScheduleAction struct {
	PlanAction
}

type Researcher struct {
	Audience
}

type Hospital struct {
	AvailableService interface{}      `json:"availableService,omitempty"` // MedicalTest, MedicalProcedure, MedicalTherapy
	MedicalSpecialty MedicalSpecialty `json:"medicalSpecialty,omitempty"`
	MedicalOrganization
	EmergencyService
	CivicStructure
}

type UserPageVisits struct {
	UserInteraction
}

type Distillery struct {
	FoodEstablishment
}

type VideoObject struct {
	VideoFrameSize Text        `json:"videoFrameSize,omitempty"`
	Directors      Person      `json:"directors,omitempty"`
	Thumbnail      ImageObject `json:"thumbnail,omitempty"`
	Caption        Text        `json:"caption,omitempty"`
	VideoQuality   Text        `json:"videoQuality,omitempty"`
	Actor          Person      `json:"actor,omitempty"`
	MusicBy        interface{} `json:"musicBy,omitempty"` // MusicGroup, Person
	Director       Person      `json:"director,omitempty"`
	Actors         Person      `json:"actors,omitempty"`
	Transcript     Text        `json:"transcript,omitempty"`
	MediaObject
}

type EducationalAudience struct {
	EducationalRole Text `json:"educationalRole,omitempty"`
	Audience
}

type PlanAction struct {
	ScheduledTime DateTime `json:"scheduledTime,omitempty"`
	OrganizeAction
}

type CityHall struct {
	GovernmentBuilding
}

type ImagingTest struct {
	ImagingTechnique MedicalImagingTechnique `json:"imagingTechnique,omitempty"`
	MedicalTest
}

type CancelAction struct {
	PlanAction
}

type VideoGameClip struct {
	Clip
}

type PhysicalTherapy struct {
	MedicalTherapy
}

type Paperback struct {
	BookFormatType
}

type ElectronicsStore struct {
	Store
}

type Protozoa struct {
	InfectiousAgentClass
}

type OrderAction struct {
	DeliveryMethod DeliveryMethod `json:"deliveryMethod,omitempty"`
	TradeAction
}

type OfferItemCondition struct {
	Enumeration
}

type Hematologic struct {
	MedicalSpecialty
}

type DoubleBlindedTrial struct {
	MedicalTrialDesign
}

type PotentialActionStatus struct {
	ActionStatusType
}

type ConsumeAction struct {
	ExpectsAcceptanceOf Offer `json:"expectsAcceptanceOf,omitempty"`
	Action
}

type DiscussionForumPosting struct {
	SocialMediaPosting
}

type InformAction struct {
	Event Event `json:"event,omitempty"`
	CommunicateAction
}

type MedicalGuidelineContraindication struct {
	MedicalGuideline
}

type MotorcycleDealer struct {
	AutomotiveBusiness
}

type Reservoir struct {
	BodyOfWater
}

type Nose struct {
	PhysicalExam
}

type WarrantyPromise struct {
	DurationOfWarranty QuantitativeValue `json:"durationOfWarranty,omitempty"`
	WarrantyScope      WarrantyScope     `json:"warrantyScope,omitempty"`
	StructuredValue
}

type AutoRepair struct {
	AutomotiveBusiness
}

type InvestmentFund struct {
	InvestmentOrDeposit
}

type EmployerReview struct {
	Review
}

type DrivingSchoolVehicleUsage struct {
	CarUsageType
}

type DepositAccount struct {
	InvestmentOrDeposit
	BankAccount
}

type LowLactoseDiet struct {
	RestrictedDiet
}

type Mosque struct {
	PlaceOfWorship
}

type QAPage struct {
	WebPage
}

type PhotographAction struct {
	CreateAction
}

type Residence struct {
	Place
}

type PreSale struct {
	ItemAvailability
}

type WatchAction struct {
	ConsumeAction
}

type Cemetery struct {
	CivicStructure
}

type ItemListOrderDescending struct {
	ItemListOrderType
}

type DislikeAction struct {
	ReactAction
}

type ShareAction struct {
	CommunicateAction
}

type Airline struct {
	IataCode       Text               `json:"iataCode,omitempty"`
	BoardingPolicy BoardingPolicyType `json:"boardingPolicy,omitempty"`
	Organization
}

type AggregateRating struct {
	ItemReviewed Thing   `json:"itemReviewed,omitempty"`
	RatingCount  Integer `json:"ratingCount,omitempty"`
	ReviewCount  Integer `json:"reviewCount,omitempty"`
	Rating
}

type ActivateAction struct {
	ControlAction
}

type VinylFormat struct {
	MusicReleaseFormatType
}

type Clinician struct {
	MedicalAudience
}

type Table struct {
	WebPageElement
}

type TrainTrip struct {
	ArrivalPlatform   Text         `json:"arrivalPlatform,omitempty"`
	DeparturePlatform Text         `json:"departurePlatform,omitempty"`
	Provider          interface{}  `json:"provider,omitempty"` // Organization, Person
	TrainName         Text         `json:"trainName,omitempty"`
	TrainNumber       Text         `json:"trainNumber,omitempty"`
	ArrivalStation    TrainStation `json:"arrivalStation,omitempty"`
	DepartureStation  TrainStation `json:"departureStation,omitempty"`
	DepartureTime     DateTime     `json:"departureTime,omitempty"`
	ArrivalTime       DateTime     `json:"arrivalTime,omitempty"`
	Intangible
}

type ArriveAction struct {
	MoveAction
}

type DownloadAction struct {
	TransferAction
}

type Notary struct {
	LegalService
}

type ExerciseAction struct {
	Diet                   Diet                   `json:"diet,omitempty"`
	SportsEvent            SportsEvent            `json:"sportsEvent,omitempty"`
	ExerciseType           Text                   `json:"exerciseType,omitempty"`
	ExerciseRelatedDiet    Diet                   `json:"exerciseRelatedDiet,omitempty"`
	FromLocation           Place                  `json:"fromLocation,omitempty"`
	ToLocation             Place                  `json:"toLocation,omitempty"`
	Opponent               Person                 `json:"opponent,omitempty"`
	Distance               Distance               `json:"distance,omitempty"`
	ExercisePlan           ExercisePlan           `json:"exercisePlan,omitempty"`
	SportsActivityLocation SportsActivityLocation `json:"sportsActivityLocation,omitempty"`
	Course                 Place                  `json:"course,omitempty"`
	SportsTeam             SportsTeam             `json:"sportsTeam,omitempty"`
	ExerciseCourse         Place                  `json:"exerciseCourse,omitempty"`
	PlayAction
}

type AuthorizeAction struct {
	Recipient interface{} `json:"recipient,omitempty"` // Person, Organization, ContactPoint, Audience
	AllocateAction
}

type CassetteFormat struct {
	MusicReleaseFormatType
}

type DayOfWeek struct {
	Enumeration
}

type Mass struct {
	Quantity
}

type OccupationalTherapy struct {
	MedicalTherapy
	MedicalSpecialty
}

type Cardiovascular struct {
	MedicalSpecialty
}

type DisagreeAction struct {
	ReactAction
}

type MusicVideoObject struct {
	MediaObject
}

type JoinAction struct {
	Event Event `json:"event,omitempty"`
	InteractAction
}

type WebPageElement struct {
	CreativeWork
}

type PaymentDue struct {
	PaymentStatusType
}

type NutritionInformation struct {
	SaturatedFatContent   Mass   `json:"saturatedFatContent,omitempty"`
	SugarContent          Mass   `json:"sugarContent,omitempty"`
	TransFatContent       Mass   `json:"transFatContent,omitempty"`
	CholesterolContent    Mass   `json:"cholesterolContent,omitempty"`
	FiberContent          Mass   `json:"fiberContent,omitempty"`
	CarbohydrateContent   Mass   `json:"carbohydrateContent,omitempty"`
	ServingSize           Text   `json:"servingSize,omitempty"`
	ProteinContent        Mass   `json:"proteinContent,omitempty"`
	UnsaturatedFatContent Mass   `json:"unsaturatedFatContent,omitempty"`
	FatContent            Mass   `json:"fatContent,omitempty"`
	Calories              Energy `json:"calories,omitempty"`
	SodiumContent         Mass   `json:"sodiumContent,omitempty"`
	StructuredValue
}

type FindAction struct {
	Action
}

type LandmarksOrHistoricalBuildings struct {
	Place
}

type Osteopathic struct {
	MedicineSystem
}

type MedicalBusiness struct {
	LocalBusiness
}

type ReservationStatusType struct {
	Enumeration
}

type RsvpResponseNo struct {
	RsvpResponseType
}

type EventScheduled struct {
	EventStatusType
}

type ActiveActionStatus struct {
	ActionStatusType
}

type MedicalStudy struct {
	Population      Text               `json:"population,omitempty"`
	StudySubject    MedicalEntity      `json:"studySubject,omitempty"`
	StudyLocation   AdministrativeArea `json:"studyLocation,omitempty"`
	Outcome         interface{}        `json:"outcome,omitempty"` // MedicalEntity, Text
	HealthCondition MedicalCondition   `json:"healthCondition,omitempty"`
	Sponsor         Organization       `json:"sponsor,omitempty"`
	Status          interface{}        `json:"status,omitempty"` // EventStatusType, MedicalStudyStatus, Text
	MedicalEntity
}

type TVSeries struct {
	CountryOfOrigin   Country            `json:"countryOfOrigin,omitempty"`
	NumberOfEpisodes  Integer            `json:"numberOfEpisodes,omitempty"`
	Trailer           VideoObject        `json:"trailer,omitempty"`
	Actor             Person             `json:"actor,omitempty"`
	Directors         Person             `json:"directors,omitempty"`
	Episode           Episode            `json:"episode,omitempty"`
	MusicBy           interface{}        `json:"musicBy,omitempty"` // MusicGroup, Person
	ProductionCompany Organization       `json:"productionCompany,omitempty"`
	Seasons           CreativeWorkSeason `json:"seasons,omitempty"`
	Season            CreativeWorkSeason `json:"season,omitempty"`
	ContainsSeason    CreativeWorkSeason `json:"containsSeason,omitempty"`
	NumberOfSeasons   Integer            `json:"numberOfSeasons,omitempty"`
	Episodes          Episode            `json:"episodes,omitempty"`
	Actors            Person             `json:"actors,omitempty"`
	Director          Person             `json:"director,omitempty"`
	CreativeWork
	CreativeWorkSeries
}

type InvestmentOrDeposit struct {
	FinancialProduct
}

type InStock struct {
	ItemAvailability
}

type MusicRelease struct {
	RecordLabel        Organization           `json:"recordLabel,omitempty"`
	MusicReleaseFormat MusicReleaseFormatType `json:"musicReleaseFormat,omitempty"`
	CreditedTo         interface{}            `json:"creditedTo,omitempty"` // Organization, Person
	CatalogNumber      Text                   `json:"catalogNumber,omitempty"`
	ReleaseOf          MusicAlbum             `json:"releaseOf,omitempty"`
	MusicPlaylist
}

type Musculoskeletal struct {
	MedicalSpecialty
}

type Bacteria struct {
	InfectiousAgentClass
}

type Urologic struct {
	MedicalSpecialty
}

type Enumeration struct {
	SupersededBy interface{} `json:"supersededBy,omitempty"` // Property, Enumeration, Class
	Intangible
}

type ProfilePage struct {
	WebPage
}

type DietarySupplement struct {
	LegalStatus         interface{}             `json:"legalStatus,omitempty"` // DrugLegalStatus, Text, MedicalEnumeration
	Background          Text                    `json:"background,omitempty"`
	TargetPopulation    Text                    `json:"targetPopulation,omitempty"`
	MechanismOfAction   Text                    `json:"mechanismOfAction,omitempty"`
	Manufacturer        interface{}             `json:"manufacturer,omitempty"` //
	NonProprietaryName  Text                    `json:"nonProprietaryName,omitempty"`
	ActiveIngredient    Text                    `json:"activeIngredient,omitempty"`
	IsProprietary       Boolean                 `json:"isProprietary,omitempty"`
	ProprietaryName     Text                    `json:"proprietaryName,omitempty"`
	SafetyConsideration Text                    `json:"safetyConsideration,omitempty"`
	RecommendedIntake   RecommendedDoseSchedule `json:"recommendedIntake,omitempty"`
	MaximumIntake       MaximumDoseSchedule     `json:"maximumIntake,omitempty"`
	Substance
}

type Female struct {
	GenderType
}

type SuperficialAnatomy struct {
	RelatedAnatomy            interface{}      `json:"relatedAnatomy,omitempty"` // AnatomicalSystem, AnatomicalStructure
	AssociatedPathophysiology Text             `json:"associatedPathophysiology,omitempty"`
	RelatedTherapy            MedicalTherapy   `json:"relatedTherapy,omitempty"`
	Significance              Text             `json:"significance,omitempty"`
	RelatedCondition          MedicalCondition `json:"relatedCondition,omitempty"`
	MedicalEntity
}

type RadioEpisode struct {
	Episode
}

type LinkRole struct {
	InLanguage       interface{} `json:"inLanguage,omitempty"` //
	LinkRelationship Text        `json:"linkRelationship,omitempty"`
	Role
}

type FinancialService struct {
	FeesAndCommissionsSpecification interface{} `json:"feesAndCommissionsSpecification,omitempty"` // Text, URL
	LocalBusiness
}

type SportsClub struct {
	SportsActivityLocation
}

type DeliveryEvent struct {
	AvailableThrough  DateTime       `json:"availableThrough,omitempty"`
	HasDeliveryMethod DeliveryMethod `json:"hasDeliveryMethod,omitempty"`
	AvailableFrom     DateTime       `json:"availableFrom,omitempty"`
	AccessCode        Text           `json:"accessCode,omitempty"`
	Event
}

type VoteAction struct {
	Candidate Person `json:"candidate,omitempty"`
	ChooseAction
}

type Joint struct {
	FunctionalClass   interface{} `json:"functionalClass,omitempty"` // MedicalEntity, Text
	BiomechnicalClass Text        `json:"biomechnicalClass,omitempty"`
	StructuralClass   Text        `json:"structuralClass,omitempty"`
	AnatomicalStructure
}

type AnatomicalSystem struct {
	RelatedStructure          AnatomicalStructure `json:"relatedStructure,omitempty"`
	RelatedTherapy            MedicalTherapy      `json:"relatedTherapy,omitempty"`
	AssociatedPathophysiology Text                `json:"associatedPathophysiology,omitempty"`
	RelatedCondition          MedicalCondition    `json:"relatedCondition,omitempty"`
	ComprisedOf               interface{}         `json:"comprisedOf,omitempty"` // AnatomicalStructure, AnatomicalSystem
	MedicalEntity
}

type OfflineTemporarily struct {
	GameServerStatus
}

type SellAction struct {
	Buyer           Person          `json:"buyer,omitempty"`
	WarrantyPromise WarrantyPromise `json:"warrantyPromise,omitempty"`
	TradeAction
}

type PetStore struct {
	Store
}

type ChooseAction struct {
	ActionOption interface{} `json:"actionOption,omitempty"` // Thing, Text
	Option       interface{} `json:"option,omitempty"`       // Text, Thing
	AssessAction
}

type OTC struct {
	DrugPrescriptionStatus
}

type IndividualProduct struct {
	SerialNumber Text `json:"serialNumber,omitempty"`
	Product
}

type SaleEvent struct {
	Event
}

type GlutenFreeDiet struct {
	RestrictedDiet
}

type BroadcastFrequencySpecification struct {
	BroadcastFrequencyValue interface{} `json:"broadcastFrequencyValue,omitempty"` // QuantitativeValue, Number
	Intangible
}

type TVSeason struct {
	CountryOfOrigin Country  `json:"countryOfOrigin,omitempty"`
	PartOfTVSeries  TVSeries `json:"partOfTVSeries,omitempty"`
	CreativeWorkSeason
	CreativeWork
}

type NewCondition struct {
	OfferItemCondition
}

type DataCatalog struct {
	Dataset              Dataset     `json:"dataset,omitempty"`
	MeasurementTechnique interface{} `json:"measurementTechnique,omitempty"` // URL, Text
	CreativeWork
}

type FDAcategoryX struct {
	DrugPregnancyCategory
}

type WebAPI struct {
	Documentation interface{} `json:"documentation,omitempty"` // CreativeWork, URL
	Service
}

type StadiumOrArena struct {
	CivicStructure
	SportsActivityLocation
}

type SoundtrackAlbum struct {
	MusicAlbumProductionType
}

type ClothingStore struct {
	Store
}

type OrderReturned struct {
	OrderStatus
}

type Dataset struct {
	Spatial               Place        `json:"spatial,omitempty"`
	Issn                  Text         `json:"issn,omitempty"`
	Temporal              DateTime     `json:"temporal,omitempty"`
	IncludedDataCatalog   DataCatalog  `json:"includedDataCatalog,omitempty"`
	MeasurementTechnique  interface{}  `json:"measurementTechnique,omitempty"` // URL, Text
	VariableMeasured      interface{}  `json:"variableMeasured,omitempty"`     // PropertyValue, Text
	IncludedInDataCatalog DataCatalog  `json:"includedInDataCatalog,omitempty"`
	Catalog               DataCatalog  `json:"catalog,omitempty"`
	DatasetTimeInterval   DateTime     `json:"datasetTimeInterval,omitempty"`
	Distribution          DataDownload `json:"distribution,omitempty"`
	VariablesMeasured     interface{}  `json:"variablesMeasured,omitempty"` // PropertyValue, Text
	CreativeWork
}

type BoardingPolicyType struct {
	Enumeration
}

type Gynecologic struct {
	MedicalBusiness
	MedicalSpecialty
}

type PaymentCard struct {
	FloorLimit         MonetaryAmount `json:"floorLimit,omitempty"`
	ContactlessPayment Boolean        `json:"contactlessPayment,omitempty"`
	CashBack           interface{}    `json:"cashBack,omitempty"` // Number, Boolean
	FinancialProduct
	PaymentMethod
}

type TelevisionChannel struct {
	BroadcastChannel
}

type WriteAction struct {
	Language Language `json:"language,omitempty"`
	CreateAction
}

type Casino struct {
	EntertainmentBusiness
}

type Toxicologic struct {
	MedicalSpecialty
}

type OccupationalActivity struct {
	PhysicalActivityCategory
}

type UserComments struct {
	CommentTime interface{}  `json:"commentTime,omitempty"` // Date, DateTime
	ReplyToUrl  URL          `json:"replyToUrl,omitempty"`
	CommentText Text         `json:"commentText,omitempty"`
	Discusses   CreativeWork `json:"discusses,omitempty"`
	Creator     interface{}  `json:"creator,omitempty"` // Organization, Person
	UserInteraction
}

type House struct {
	NumberOfRooms interface{} `json:"numberOfRooms,omitempty"` // QuantitativeValue, Number
	Accommodation
}

type AlbumRelease struct {
	MusicAlbumReleaseType
}

type Service struct {
	ServiceAudience  Audience                  `json:"serviceAudience,omitempty"`
	AvailableChannel ServiceChannel            `json:"availableChannel,omitempty"`
	Logo             interface{}               `json:"logo,omitempty"` // URL, ImageObject
	Review           Review                    `json:"review,omitempty"`
	Offers           Offer                     `json:"offers,omitempty"`
	ProviderMobility Text                      `json:"providerMobility,omitempty"`
	ServiceOutput    Thing                     `json:"serviceOutput,omitempty"`
	Brand            interface{}               `json:"brand,omitempty"`       // Organization, Brand
	IsRelatedTo      interface{}               `json:"isRelatedTo,omitempty"` // Service, Product
	AreaServed       interface{}               `json:"areaServed,omitempty"`  // GeoShape, Text, Place, AdministrativeArea
	HoursAvailable   OpeningHoursSpecification `json:"hoursAvailable,omitempty"`
	HasOfferCatalog  OfferCatalog              `json:"hasOfferCatalog,omitempty"`
	TermsOfService   interface{}               `json:"termsOfService,omitempty"` // Text, URL
	AggregateRating  AggregateRating           `json:"aggregateRating,omitempty"`
	Produces         Thing                     `json:"produces,omitempty"`
	ServiceType      Text                      `json:"serviceType,omitempty"`
	Award            Text                      `json:"award,omitempty"`
	Audience         Audience                  `json:"audience,omitempty"`
	ServiceArea      interface{}               `json:"serviceArea,omitempty"` // Place, GeoShape, AdministrativeArea
	Provider         interface{}               `json:"provider,omitempty"`    // Organization, Person
	Broker           interface{}               `json:"broker,omitempty"`      // Person, Organization
	IsSimilarTo      interface{}               `json:"isSimilarTo,omitempty"` // Product, Service
	Intangible
}

type Map struct {
	MapType MapCategoryType `json:"mapType,omitempty"`
	CreativeWork
}

type GameServer struct {
	PlayersOnline Integer          `json:"playersOnline,omitempty"`
	Game          VideoGame        `json:"game,omitempty"`
	ServerStatus  GameServerStatus `json:"serverStatus,omitempty"`
	Intangible
}

type SocialEvent struct {
	Event
}

type Substance struct {
	MaximumIntake    MaximumDoseSchedule `json:"maximumIntake,omitempty"`
	ActiveIngredient Text                `json:"activeIngredient,omitempty"`
	MedicalEntity
}

type CohortStudy struct {
	MedicalObservationalStudyDesign
}

type FurnitureStore struct {
	Store
}

type Duration struct {
	Quantity
}

type WarrantyScope struct {
	Enumeration
}

type Pediatric struct {
	MedicalBusiness
	MedicalSpecialty
}

type Movie struct {
	Trailer           VideoObject  `json:"trailer,omitempty"`
	CountryOfOrigin   Country      `json:"countryOfOrigin,omitempty"`
	SubtitleLanguage  interface{}  `json:"subtitleLanguage,omitempty"` // Language, Text
	MusicBy           interface{}  `json:"musicBy,omitempty"`          // MusicGroup, Person
	Director          Person       `json:"director,omitempty"`
	Directors         Person       `json:"directors,omitempty"`
	Actor             Person       `json:"actor,omitempty"`
	Actors            Person       `json:"actors,omitempty"`
	ProductionCompany Organization `json:"productionCompany,omitempty"`
	CreativeWork
}

type ImageGallery struct {
	CollectionPage
}

type AcceptAction struct {
	AllocateAction
}

type Float struct {
	Number
}

type CategoryCode struct {
	InCodeSet interface{} `json:"inCodeSet,omitempty"` // CategoryCodeSet, URL
	CodeValue Text        `json:"codeValue,omitempty"`
	Intangible
}

type ReservationPackage struct {
	SubReservation Reservation `json:"subReservation,omitempty"`
	Reservation
}

type Recruiting struct {
	MedicalStudyStatus
}

type Gastroenterologic struct {
	MedicalSpecialty
}

type APIReference struct {
	ProgrammingModel      Text `json:"programmingModel,omitempty"`
	ExecutableLibraryName Text `json:"executableLibraryName,omitempty"`
	AssemblyVersion       Text `json:"assemblyVersion,omitempty"`
	Assembly              Text `json:"assembly,omitempty"`
	TargetPlatform        Text `json:"targetPlatform,omitempty"`
	TechArticle
}

type MedicalGuideline struct {
	GuidelineSubject MedicalEntity        `json:"guidelineSubject,omitempty"`
	EvidenceOrigin   Text                 `json:"evidenceOrigin,omitempty"`
	GuidelineDate    Date                 `json:"guidelineDate,omitempty"`
	EvidenceLevel    MedicalEvidenceLevel `json:"evidenceLevel,omitempty"`
	MedicalEntity
}

type VisualArtsEvent struct {
	Event
}

type Specialty struct {
	Enumeration
}

type OfferCatalog struct {
	ItemList
}

type Landform struct {
	Place
}

type MedicalProcedure struct {
	Status        interface{}          `json:"status,omitempty"` // EventStatusType, MedicalStudyStatus, Text
	Followup      Text                 `json:"followup,omitempty"`
	HowPerformed  Text                 `json:"howPerformed,omitempty"`
	BodyLocation  Text                 `json:"bodyLocation,omitempty"`
	ProcedureType MedicalProcedureType `json:"procedureType,omitempty"`
	Indication    MedicalIndication    `json:"indication,omitempty"`
	Outcome       interface{}          `json:"outcome,omitempty"`     // MedicalEntity, Text
	Preparation   interface{}          `json:"preparation,omitempty"` // Text, MedicalEntity
	MedicalEntity
}

type ReserveAction struct {
	PlanAction
}

type OpinionNewsArticle struct {
	NewsArticle
}

type PublicationVolume struct {
	VolumeNumber interface{} `json:"volumeNumber,omitempty"` // Integer, Text
	CreativeWork
}

type CompilationAlbum struct {
	MusicAlbumProductionType
}

type SingleBlindedTrial struct {
	MedicalTrialDesign
}

type ComedyEvent struct {
	Event
}

type RejectAction struct {
	AllocateAction
}

type Resort struct {
	LodgingBusiness
}

type Periodical struct {
	CreativeWorkSeries
}

type TransitMap struct {
	MapCategoryType
}

type Crematorium struct {
	CivicStructure
}

type AdministrativeArea struct {
	Place
}

type DeliveryMethod struct {
	Enumeration
}

type LiveAlbum struct {
	MusicAlbumProductionType
}

type MedicineSystem struct {
	MedicalEnumeration
}

type OfficialLegalValue struct {
	LegalValueLevel
}

type Dermatology struct {
	MedicalBusiness
	MedicalSpecialty
}

type CheckoutPage struct {
	WebPage
}

type AddAction struct {
	UpdateAction
}

type DryCleaningOrLaundry struct {
	LocalBusiness
}

type TaxiStand struct {
	CivicStructure
}

type NotInForce struct {
	LegalForceStatus
}

type AutomotiveBusiness struct {
	LocalBusiness
}

type Bone struct {
	AnatomicalStructure
}

type MedicalEntity struct {
	Guideline            MedicalGuideline `json:"guideline,omitempty"`
	RelevantSpecialty    MedicalSpecialty `json:"relevantSpecialty,omitempty"`
	Study                MedicalStudy     `json:"study,omitempty"`
	MedicineSystem       MedicineSystem   `json:"medicineSystem,omitempty"`
	RecognizingAuthority Organization     `json:"recognizingAuthority,omitempty"`
	Code                 MedicalCode      `json:"code,omitempty"`
	LegalStatus          interface{}      `json:"legalStatus,omitempty"` // DrugLegalStatus, Text, MedicalEnumeration
	Thing
}

type PsychologicalTreatment struct {
	TherapeuticProcedure
}

type AutoDealer struct {
	AutomotiveBusiness
}

type MoveAction struct {
	ToLocation   Place `json:"toLocation,omitempty"`
	FromLocation Place `json:"fromLocation,omitempty"`
	Action
}

type IceCreamShop struct {
	FoodEstablishment
}

type Flexibility struct {
	PhysicalActivityCategory
}

type Car struct {
	AcrissCode Text              `json:"acrissCode,omitempty"`
	RoofLoad   QuantitativeValue `json:"roofLoad,omitempty"`
	Vehicle
}

type RecommendedDoseSchedule struct {
	DoseSchedule
}

type Beach struct {
	CivicStructure
}

type DatedMoneySpecification struct {
	EndDate   interface{} `json:"endDate,omitempty"`   // Date, DateTime
	StartDate interface{} `json:"startDate,omitempty"` // DateTime, Date
	StructuredValue
}

type ComedyClub struct {
	EntertainmentBusiness
}

type RepaymentSpecification struct {
	NumberOfLoanPayments   Number         `json:"numberOfLoanPayments,omitempty"`
	LoanPaymentFrequency   Number         `json:"loanPaymentFrequency,omitempty"`
	EarlyPrepaymentPenalty MonetaryAmount `json:"earlyPrepaymentPenalty,omitempty"`
	LoanPaymentAmount      MonetaryAmount `json:"loanPaymentAmount,omitempty"`
	DownPayment            interface{}    `json:"downPayment,omitempty"` // MonetaryAmount, Number
	StructuredValue
}

type Retail struct {
	DrugCostCategory
}

type FastFoodRestaurant struct {
	FoodEstablishment
}

type SportsEvent struct {
	HomeTeam   interface{} `json:"homeTeam,omitempty"`   // SportsTeam, Person
	Competitor interface{} `json:"competitor,omitempty"` // Person, SportsTeam
	AwayTeam   interface{} `json:"awayTeam,omitempty"`   // SportsTeam, Person
	Event
}

type School struct {
	EducationalOrganization
}

type WorkersUnion struct {
	Organization
}

type PostalAddress struct {
	AddressCountry      interface{} `json:"addressCountry,omitempty"` // Country, Text
	AddressLocality     Text        `json:"addressLocality,omitempty"`
	PostalCode          Text        `json:"postalCode,omitempty"`
	StreetAddress       Text        `json:"streetAddress,omitempty"`
	PostOfficeBoxNumber Text        `json:"postOfficeBoxNumber,omitempty"`
	AddressRegion       Text        `json:"addressRegion,omitempty"`
	ContactPoint
}

type AutoPartsStore struct {
	AutomotiveBusiness
	Store
}

type Report struct {
	ReportNumber Text `json:"reportNumber,omitempty"`
	Article
}

type PropertyValueSpecification struct {
	DefaultValue   interface{} `json:"defaultValue,omitempty"` // Text, Thing
	ValueRequired  Boolean     `json:"valueRequired,omitempty"`
	MinValue       Number      `json:"minValue,omitempty"`
	ValueName      Text        `json:"valueName,omitempty"`
	ValuePattern   Text        `json:"valuePattern,omitempty"`
	StepValue      Number      `json:"stepValue,omitempty"`
	ReadonlyValue  Boolean     `json:"readonlyValue,omitempty"`
	ValueMaxLength Number      `json:"valueMaxLength,omitempty"`
	MaxValue       Number      `json:"maxValue,omitempty"`
	MultipleValues Boolean     `json:"multipleValues,omitempty"`
	ValueMinLength Number      `json:"valueMinLength,omitempty"`
	Intangible
}

type LegalService struct {
	LocalBusiness
}

type MedicalConditionStage struct {
	StageAsNumber  Number `json:"stageAsNumber,omitempty"`
	SubStageSuffix Text   `json:"subStageSuffix,omitempty"`
	MedicalIntangible
}

type LowSaltDiet struct {
	RestrictedDiet
}

type ScholarlyArticle struct {
	Article
}

type GeoShape struct {
	Line           Text        `json:"line,omitempty"`
	PostalCode     Text        `json:"postalCode,omitempty"`
	Polygon        Text        `json:"polygon,omitempty"`
	Address        interface{} `json:"address,omitempty"` // PostalAddress, Text
	Circle         Text        `json:"circle,omitempty"`
	Box            Text        `json:"box,omitempty"`
	AddressCountry interface{} `json:"addressCountry,omitempty"` // Country, Text
	Elevation      interface{} `json:"elevation,omitempty"`      // Text, Number
	StructuredValue
}

type Artery struct {
	SupplyTo       AnatomicalStructure `json:"supplyTo,omitempty"`
	ArterialBranch AnatomicalStructure `json:"arterialBranch,omitempty"`
	Source         AnatomicalStructure `json:"source,omitempty"`
	Vessel
}

type DeactivateAction struct {
	ControlAction
}

type TireShop struct {
	Store
}

type Quantity struct {
	Intangible
}

type Thing struct {
	AdditionalType            URL         `json:"additionalType,omitempty"`
	Identifier                interface{} `json:"identifier,omitempty"` // PropertyValue, Text, URL
	SameAs                    URL         `json:"sameAs,omitempty"`
	PotentialAction           Action      `json:"potentialAction,omitempty"`
	Url                       URL         `json:"url,omitempty"`
	Description               Text        `json:"description,omitempty"`
	Name                      Text        `json:"name,omitempty"`
	MainEntityOfPage          interface{} `json:"mainEntityOfPage,omitempty"` // URL, CreativeWork
	DisambiguatingDescription Text        `json:"disambiguatingDescription,omitempty"`
	Image                     interface{} `json:"image,omitempty"` // ImageObject, URL
	AlternateName             Text        `json:"alternateName,omitempty"`
}

type BusinessEvent struct {
	Event
}

type Prion struct {
	InfectiousAgentClass
}

type Airport struct {
	IcaoCode Text `json:"icaoCode,omitempty"`
	IataCode Text `json:"iataCode,omitempty"`
	CivicStructure
}

type Playground struct {
	CivicStructure
}

type Neuro struct {
	PhysicalExam
}

type GolfCourse struct {
	SportsActivityLocation
}

type BedAndBreakfast struct {
	LodgingBusiness
}

type ChildCare struct {
	LocalBusiness
}

type FourWheelDriveConfiguration struct {
	DriveWheelConfigurationValue
}

type ReadPermission struct {
	DigitalDocumentPermissionType
}

type AchieveAction struct {
	Action
}

type MedicalRiskScore struct {
	Algorithm Text `json:"algorithm,omitempty"`
	MedicalRiskEstimator
}

type DanceGroup struct {
	PerformingGroup
}

type FilmAction struct {
	CreateAction
}

type MusicReleaseFormatType struct {
	Enumeration
}

type Ultrasound struct {
	MedicalImagingTechnique
}

type CompleteDataFeed struct {
	DataFeed
}

type ExercisePlan struct {
	AdditionalVariable Text        `json:"additionalVariable,omitempty"`
	Workload           interface{} `json:"workload,omitempty"`          // QualitativeValue, Energy
	Intensity          interface{} `json:"intensity,omitempty"`         // QualitativeValue, Text
	RestPeriods        interface{} `json:"restPeriods,omitempty"`       // Text, QualitativeValue
	ActivityDuration   interface{} `json:"activityDuration,omitempty"`  // QualitativeValue, Duration
	ActivityFrequency  interface{} `json:"activityFrequency,omitempty"` // Text, QualitativeValue
	ExerciseType       Text        `json:"exerciseType,omitempty"`
	Repetitions        interface{} `json:"repetitions,omitempty"` // QualitativeValue, Number
	CreativeWork
	PhysicalActivity
}

type UseAction struct {
	ConsumeAction
}

type AnimalShelter struct {
	LocalBusiness
}

type OnSitePickup struct {
	DeliveryMethod
}

type EvidenceLevelB struct {
	MedicalEvidenceLevel
}

type PublicHealth struct {
	MedicalBusiness
	MedicalSpecialty
}

type Therapeutic struct {
	MedicalDevicePurpose
}

type MedicalTestPanel struct {
	SubTest MedicalTest `json:"subTest,omitempty"`
	MedicalTest
}

type SpreadsheetDigitalDocument struct {
	DigitalDocument
}

type CriticReview struct {
	Review
}

type Recipe struct {
	Ingredients        Text                 `json:"ingredients,omitempty"`
	RecipeIngredient   Text                 `json:"recipeIngredient,omitempty"`
	RecipeCuisine      Text                 `json:"recipeCuisine,omitempty"`
	RecipeCategory     Text                 `json:"recipeCategory,omitempty"`
	RecipeYield        interface{}          `json:"recipeYield,omitempty"` // Text, QuantitativeValue
	CookTime           Duration             `json:"cookTime,omitempty"`
	RecipeInstructions interface{}          `json:"recipeInstructions,omitempty"` // CreativeWork, ItemList, Text
	Nutrition          NutritionInformation `json:"nutrition,omitempty"`
	SuitableForDiet    RestrictedDiet       `json:"suitableForDiet,omitempty"`
	CookingMethod      Text                 `json:"cookingMethod,omitempty"`
	HowTo
}

type CreativeWorkSeason struct {
	SeasonNumber      interface{}        `json:"seasonNumber,omitempty"` // Integer, Text
	ProductionCompany Organization       `json:"productionCompany,omitempty"`
	StartDate         interface{}        `json:"startDate,omitempty"` // DateTime, Date
	EndDate           interface{}        `json:"endDate,omitempty"`   // Date, DateTime
	Director          Person             `json:"director,omitempty"`
	PartOfSeries      CreativeWorkSeries `json:"partOfSeries,omitempty"`
	Episode           Episode            `json:"episode,omitempty"`
	Trailer           VideoObject        `json:"trailer,omitempty"`
	Actor             Person             `json:"actor,omitempty"`
	NumberOfEpisodes  Integer            `json:"numberOfEpisodes,omitempty"`
	Episodes          Episode            `json:"episodes,omitempty"`
	CreativeWork
}

type Code struct {
	CreativeWork
}

type DrugCost struct {
	CostOrigin         Text               `json:"costOrigin,omitempty"`
	CostCategory       DrugCostCategory   `json:"costCategory,omitempty"`
	DrugUnit           Text               `json:"drugUnit,omitempty"`
	CostCurrency       Text               `json:"costCurrency,omitempty"`
	ApplicableLocation AdministrativeArea `json:"applicableLocation,omitempty"`
	CostPerUnit        interface{}        `json:"costPerUnit,omitempty"` // Number, Text, QualitativeValue
	MedicalEnumeration
}

type TelevisionStation struct {
	LocalBusiness
}

type TVEpisode struct {
	CountryOfOrigin  Country     `json:"countryOfOrigin,omitempty"`
	SubtitleLanguage interface{} `json:"subtitleLanguage,omitempty"` // Language, Text
	PartOfTVSeries   TVSeries    `json:"partOfTVSeries,omitempty"`
	Episode
}

type ActiveNotRecruiting struct {
	MedicalStudyStatus
}

type Definitive struct {
	LegalValueLevel
}

type Zoo struct {
	CivicStructure
}

type Electrician struct {
	HomeAndConstructionBusiness
}

type PlayAction struct {
	Audience Audience `json:"audience,omitempty"`
	Event    Event    `json:"event,omitempty"`
	Action
}

type ContactPage struct {
	WebPage
}

type PercutaneousProcedure struct {
	MedicalProcedureType
}

type Reservation struct {
	ReservationId         Text                  `json:"reservationId,omitempty"`
	UnderName             interface{}           `json:"underName,omitempty"`  // Organization, Person
	TotalPrice            interface{}           `json:"totalPrice,omitempty"` // Number, Text, PriceSpecification
	BookingTime           DateTime              `json:"bookingTime,omitempty"`
	PriceCurrency         Text                  `json:"priceCurrency,omitempty"`
	ModifiedTime          DateTime              `json:"modifiedTime,omitempty"`
	ReservationStatus     ReservationStatusType `json:"reservationStatus,omitempty"`
	ReservationFor        Thing                 `json:"reservationFor,omitempty"`
	ProgramMembershipUsed ProgramMembership     `json:"programMembershipUsed,omitempty"`
	Provider              interface{}           `json:"provider,omitempty"` // Organization, Person
	Broker                interface{}           `json:"broker,omitempty"`   // Person, Organization
	ReservedTicket        Ticket                `json:"reservedTicket,omitempty"`
	BookingAgent          interface{}           `json:"bookingAgent,omitempty"` // Organization, Person
	Intangible
}

type ActionStatusType struct {
	Enumeration
}

type BusStation struct {
	CivicStructure
}

type VideoGallery struct {
	CollectionPage
}

type MobilePhoneStore struct {
	Store
}

type OpenTrial struct {
	MedicalTrialDesign
}

type InviteAction struct {
	Event Event `json:"event,omitempty"`
	CommunicateAction
}

type PlaceboControlledTrial struct {
	MedicalTrialDesign
}

type NewsMediaOrganization struct {
	VerificationFactCheckingPolicy  interface{} `json:"verificationFactCheckingPolicy,omitempty"`  // CreativeWork, URL
	EthicsPolicy                    interface{} `json:"ethicsPolicy,omitempty"`                    // URL, CreativeWork
	UnnamedSourcesPolicy            interface{} `json:"unnamedSourcesPolicy,omitempty"`            // URL, CreativeWork
	CorrectionsPolicy               interface{} `json:"correctionsPolicy,omitempty"`               // URL, CreativeWork
	MissionCoveragePrioritiesPolicy interface{} `json:"missionCoveragePrioritiesPolicy,omitempty"` // CreativeWork, URL
	Masthead                        interface{} `json:"masthead,omitempty"`                        // URL, CreativeWork
	ActionableFeedbackPolicy        interface{} `json:"actionableFeedbackPolicy,omitempty"`        // URL, CreativeWork
	DiversityPolicy                 interface{} `json:"diversityPolicy,omitempty"`                 // CreativeWork, URL
	Organization
}

type CoverArt struct {
	VisualArtwork
}

type RealEstateAgent struct {
	LocalBusiness
}

type DrinkAction struct {
	ConsumeAction
}

type Rating struct {
	RatingValue interface{} `json:"ratingValue,omitempty"` // Number, Text
	BestRating  interface{} `json:"bestRating,omitempty"`  // Text, Number
	Author      interface{} `json:"author,omitempty"`      // Organization, Person
	WorstRating interface{} `json:"worstRating,omitempty"` // Text, Number
	Intangible
}

type AutoRental struct {
	AutomotiveBusiness
}

type HowToDirection struct {
	BeforeMedia MediaObject `json:"beforeMedia,omitempty"`
	PrepTime    Duration    `json:"prepTime,omitempty"`
	PerformTime Duration    `json:"performTime,omitempty"`
	Tool        interface{} `json:"tool,omitempty"` // HowToTool, Text
	AfterMedia  MediaObject `json:"afterMedia,omitempty"`
	Supply      interface{} `json:"supply,omitempty"` // HowToSupply, Text
	TotalTime   Duration    `json:"totalTime,omitempty"`
	DuringMedia MediaObject `json:"duringMedia,omitempty"`
	ListItem
}

type Optician struct {
	MedicalBusiness
}

type ReactAction struct {
	AssessAction
}

type MusicGroup struct {
	MusicGroupMember Person         `json:"musicGroupMember,omitempty"`
	Genre            interface{}    `json:"genre,omitempty"` // URL, Text
	Tracks           MusicRecording `json:"tracks,omitempty"`
	Albums           MusicAlbum     `json:"albums,omitempty"`
	Track            interface{}    `json:"track,omitempty"` // ItemList, MusicRecording
	Album            MusicAlbum     `json:"album,omitempty"`
	PerformingGroup
}

type OrderPaymentDue struct {
	OrderStatus
}

type RadioSeason struct {
	CreativeWorkSeason
}

type InstallAction struct {
	ConsumeAction
}

type FDAcategoryB struct {
	DrugPregnancyCategory
}

type BlogPosting struct {
	SocialMediaPosting
}

type CompletedActionStatus struct {
	ActionStatusType
}

type LiveBlogPosting struct {
	CoverageEndTime   DateTime    `json:"coverageEndTime,omitempty"`
	LiveBlogUpdate    BlogPosting `json:"liveBlogUpdate,omitempty"`
	CoverageStartTime DateTime    `json:"coverageStartTime,omitempty"`
	BlogPosting
}

type LaserDiscFormat struct {
	MusicReleaseFormatType
}

type NGO struct {
	Organization
}

type TheaterGroup struct {
	PerformingGroup
}

type ItemListOrderAscending struct {
	ItemListOrderType
}

type DataDownload struct {
	MeasurementTechnique interface{} `json:"measurementTechnique,omitempty"` // URL, Text
	MediaObject
}

type TouristInformationCenter struct {
	LocalBusiness
}

type UserLikes struct {
	UserInteraction
}

type Midwifery struct {
	MedicalBusiness
	MedicalSpecialty
}

type Pond struct {
	BodyOfWater
}

type TechArticle struct {
	ProficiencyLevel Text `json:"proficiencyLevel,omitempty"`
	Dependencies     Text `json:"dependencies,omitempty"`
	Article
}

type Course struct {
	EducationalCredentialAwarded interface{}    `json:"educationalCredentialAwarded,omitempty"` // URL, Text
	HasCourseInstance            CourseInstance `json:"hasCourseInstance,omitempty"`
	CourseCode                   Text           `json:"courseCode,omitempty"`
	CoursePrerequisites          interface{}    `json:"coursePrerequisites,omitempty"` // Text, AlignmentObject, Course
	CreativeWork
}

type Series struct {
	CreativeWork
}

type DriveWheelConfigurationValue struct {
	QualitativeValue
}

type MRI struct {
	MedicalImagingTechnique
}

type VenueMap struct {
	MapCategoryType
}

type ReplaceAction struct {
	Replacee Thing `json:"replacee,omitempty"`
	Replacer Thing `json:"replacer,omitempty"`
	UpdateAction
}

type TattooParlor struct {
	HealthAndBeautyBusiness
}

type RadioChannel struct {
	BroadcastChannel
}

type Book struct {
	Isbn          Text           `json:"isbn,omitempty"`
	Illustrator   Person         `json:"illustrator,omitempty"`
	BookEdition   Text           `json:"bookEdition,omitempty"`
	NumberOfPages Integer        `json:"numberOfPages,omitempty"`
	BookFormat    BookFormatType `json:"bookFormat,omitempty"`
	Abridged      Boolean        `json:"abridged,omitempty"`
	CreativeWork
}

type MedicalRiskEstimator struct {
	IncludedRiskFactor MedicalRiskFactor `json:"includedRiskFactor,omitempty"`
	EstimatesRiskOf    MedicalEntity     `json:"estimatesRiskOf,omitempty"`
	MedicalEntity
}

type SeatingMap struct {
	MapCategoryType
}

type Energy struct {
	Quantity
}

type DigitalAudioTapeFormat struct {
	MusicReleaseFormatType
}

type Demand struct {
	EligibleRegion            interface{}         `json:"eligibleRegion,omitempty"` // GeoShape, Text, Place
	AvailableAtOrFrom         Place               `json:"availableAtOrFrom,omitempty"`
	Warranty                  WarrantyPromise     `json:"warranty,omitempty"`
	ItemOffered               interface{}         `json:"itemOffered,omitempty"` // Product, Service
	EligibleQuantity          QuantitativeValue   `json:"eligibleQuantity,omitempty"`
	Mpn                       Text                `json:"mpn,omitempty"`
	ValidThrough              DateTime            `json:"validThrough,omitempty"`
	IneligibleRegion          interface{}         `json:"ineligibleRegion,omitempty"` // Place, GeoShape, Text
	AreaServed                interface{}         `json:"areaServed,omitempty"`       // GeoShape, Text, Place, AdministrativeArea
	AvailableDeliveryMethod   DeliveryMethod      `json:"availableDeliveryMethod,omitempty"`
	EligibleTransactionVolume PriceSpecification  `json:"eligibleTransactionVolume,omitempty"`
	Sku                       Text                `json:"sku,omitempty"`
	EligibleCustomerType      BusinessEntityType  `json:"eligibleCustomerType,omitempty"`
	EligibleDuration          QuantitativeValue   `json:"eligibleDuration,omitempty"`
	Gtin13                    Text                `json:"gtin13,omitempty"`
	BusinessFunction          BusinessFunction    `json:"businessFunction,omitempty"`
	PriceSpecification        PriceSpecification  `json:"priceSpecification,omitempty"`
	AvailabilityStarts        DateTime            `json:"availabilityStarts,omitempty"`
	Gtin8                     Text                `json:"gtin8,omitempty"`
	IncludesObject            TypeAndQuantityNode `json:"includesObject,omitempty"`
	SerialNumber              Text                `json:"serialNumber,omitempty"`
	InventoryLevel            QuantitativeValue   `json:"inventoryLevel,omitempty"`
	AcceptedPaymentMethod     interface{}         `json:"acceptedPaymentMethod,omitempty"` // LoanOrCredit, PaymentMethod
	Availability              ItemAvailability    `json:"availability,omitempty"`
	AdvanceBookingRequirement QuantitativeValue   `json:"advanceBookingRequirement,omitempty"`
	Gtin14                    Text                `json:"gtin14,omitempty"`
	AvailabilityEnds          DateTime            `json:"availabilityEnds,omitempty"`
	Seller                    interface{}         `json:"seller,omitempty"` // Person, Organization
	DeliveryLeadTime          QuantitativeValue   `json:"deliveryLeadTime,omitempty"`
	ValidFrom                 DateTime            `json:"validFrom,omitempty"`
	ItemCondition             OfferItemCondition  `json:"itemCondition,omitempty"`
	Gtin12                    Text                `json:"gtin12,omitempty"`
	Intangible
}

type OpeningHoursSpecification struct {
	DayOfWeek    DayOfWeek `json:"dayOfWeek,omitempty"`
	Opens        Time      `json:"opens,omitempty"`
	ValidThrough DateTime  `json:"validThrough,omitempty"`
	ValidFrom    DateTime  `json:"validFrom,omitempty"`
	Closes       Time      `json:"closes,omitempty"`
	StructuredValue
}

type SportsOrganization struct {
	Sport interface{} `json:"sport,omitempty"` // Text, URL
	Organization
}

type Suspended struct {
	MedicalStudyStatus
}

type CollegeOrUniversity struct {
	EducationalOrganization
}

type RadioClip struct {
	Clip
}

type SelfStorage struct {
	LocalBusiness
}

type Brewery struct {
	FoodEstablishment
}

type HousePainter struct {
	HomeAndConstructionBusiness
}

type NailSalon struct {
	HealthAndBeautyBusiness
}

type DeliveryChargeSpecification struct {
	IneligibleRegion        interface{}    `json:"ineligibleRegion,omitempty"` // Place, GeoShape, Text
	AreaServed              interface{}    `json:"areaServed,omitempty"`       // GeoShape, Text, Place, AdministrativeArea
	AppliesToDeliveryMethod DeliveryMethod `json:"appliesToDeliveryMethod,omitempty"`
	EligibleRegion          interface{}    `json:"eligibleRegion,omitempty"` // GeoShape, Text, Place
	PriceSpecification
}

type ComputerLanguage struct {
	Intangible
}

type SingleFamilyResidence struct {
	NumberOfRooms interface{}       `json:"numberOfRooms,omitempty"` // QuantitativeValue, Number
	Occupancy     QuantitativeValue `json:"occupancy,omitempty"`
	House
}

type Quotation struct {
	SpokenByCharacter interface{} `json:"spokenByCharacter,omitempty"` // Organization, Person
	CreativeWork
}

type Audience struct {
	GeographicArea AdministrativeArea `json:"geographicArea,omitempty"`
	AudienceType   Text               `json:"audienceType,omitempty"`
	Intangible
}

type MenuItem struct {
	Nutrition       NutritionInformation `json:"nutrition,omitempty"`
	Offers          Offer                `json:"offers,omitempty"`
	SuitableForDiet RestrictedDiet       `json:"suitableForDiet,omitempty"`
	MenuAddOn       interface{}          `json:"menuAddOn,omitempty"` // MenuSection, MenuItem
	Intangible
}

type RearWheelDriveConfiguration struct {
	DriveWheelConfigurationValue
}

type BorrowAction struct {
	Lender interface{} `json:"lender,omitempty"` // Person, Organization
	TransferAction
}

type FMRadioChannel struct {
	RadioChannel
}

type MedicalObservationalStudy struct {
	StudyDesign MedicalObservationalStudyDesign `json:"studyDesign,omitempty"`
	MedicalStudy
}

type RadioStation struct {
	LocalBusiness
}

type Hostel struct {
	LodgingBusiness
}

type FireStation struct {
	EmergencyService
	CivicStructure
}

type ControlAction struct {
	Action
}

type ApplyAction struct {
	OrganizeAction
}

type WesternConventional struct {
	MedicineSystem
}

type EBook struct {
	BookFormatType
}

type Monday struct {
	DayOfWeek
}

type RsvpResponseType struct {
	Enumeration
}

type LikeAction struct {
	ReactAction
}

type Legislation struct {
	LegislationChangedBy      Legislation      `json:"legislationChangedBy,omitempty"`
	LegislationDateVersion    Date             `json:"legislationDateVersion,omitempty"`
	LegislationApplies        Legislation      `json:"legislationApplies,omitempty"`
	LegislationLegalForce     LegalForceStatus `json:"legislationLegalForce,omitempty"`
	LegislationPassedBy       interface{}      `json:"legislationPassedBy,omitempty"` // Organization, Person
	LegislationConsolidates   Legislation      `json:"legislationConsolidates,omitempty"`
	LegislationAppliedBy      Legislation      `json:"legislationAppliedBy,omitempty"`
	LegislationTransposes     Legislation      `json:"legislationTransposes,omitempty"`
	LegislationIdentifier     Text             `json:"legislationIdentifier,omitempty"`
	LegislationResponsible    interface{}      `json:"legislationResponsible,omitempty"` // Person, Organization
	LegislationTransposedBy   Legislation      `json:"legislationTransposedBy,omitempty"`
	LegislationDate           Date             `json:"legislationDate,omitempty"`
	LegislationChanges        Legislation      `json:"legislationChanges,omitempty"`
	LegislationType           interface{}      `json:"legislationType,omitempty"` // CategoryCode, Text
	LegislationConsolidatedBy Legislation      `json:"legislationConsolidatedBy,omitempty"`
	CreativeWork
}

type BookFormatType struct {
	Enumeration
}

type FinancialProduct struct {
	AnnualPercentageRate            interface{} `json:"annualPercentageRate,omitempty"`            // QuantitativeValue, Number
	InterestRate                    interface{} `json:"interestRate,omitempty"`                    // QuantitativeValue, Number
	FeesAndCommissionsSpecification interface{} `json:"feesAndCommissionsSpecification,omitempty"` // Text, URL
	Service
}

type BrokerageAccount struct {
	InvestmentOrDeposit
}

type QuoteAction struct {
	TradeAction
}

type MedicalCondition struct {
	ExpectedPrognosis     Text                  `json:"expectedPrognosis,omitempty"`
	Pathophysiology       Text                  `json:"pathophysiology,omitempty"`
	PrimaryPrevention     MedicalTherapy        `json:"primaryPrevention,omitempty"`
	Epidemiology          Text                  `json:"epidemiology,omitempty"`
	Drug                  Drug                  `json:"drug,omitempty"`
	SignOrSymptom         MedicalSignOrSymptom  `json:"signOrSymptom,omitempty"`
	RiskFactor            MedicalRiskFactor     `json:"riskFactor,omitempty"`
	PossibleComplication  Text                  `json:"possibleComplication,omitempty"`
	SecondaryPrevention   MedicalTherapy        `json:"secondaryPrevention,omitempty"`
	AssociatedAnatomy     interface{}           `json:"associatedAnatomy,omitempty"` // AnatomicalSystem, SuperficialAnatomy, AnatomicalStructure
	NaturalProgression    Text                  `json:"naturalProgression,omitempty"`
	DifferentialDiagnosis DDxElement            `json:"differentialDiagnosis,omitempty"`
	Cause                 MedicalCause          `json:"cause,omitempty"`
	PossibleTreatment     MedicalTherapy        `json:"possibleTreatment,omitempty"`
	Stage                 MedicalConditionStage `json:"stage,omitempty"`
	Subtype               Text                  `json:"subtype,omitempty"`
	TypicalTest           MedicalTest           `json:"typicalTest,omitempty"`
	Status                interface{}           `json:"status,omitempty"` // EventStatusType, MedicalStudyStatus, Text
	MedicalEntity
}

type FoodService struct {
	Service
}

type ListItem struct {
	NextItem     ListItem    `json:"nextItem,omitempty"`
	PreviousItem ListItem    `json:"previousItem,omitempty"`
	Item         Thing       `json:"item,omitempty"`
	Position     interface{} `json:"position,omitempty"` // Integer, Text
	Intangible
}

type AnaerobicActivity struct {
	PhysicalActivityCategory
}

type Online struct {
	GameServerStatus
}

type BookSeries struct {
	CreativeWorkSeries
}

type AudiobookFormat struct {
	BookFormatType
}

type GardenStore struct {
	Store
}

type Event struct {
	DoorTime                  DateTime        `json:"doorTime,omitempty"`
	PreviousStartDate         Date            `json:"previousStartDate,omitempty"`
	SubEvent                  Event           `json:"subEvent,omitempty"`
	Offers                    Offer           `json:"offers,omitempty"`
	TypicalAgeRange           Text            `json:"typicalAgeRange,omitempty"`
	AggregateRating           AggregateRating `json:"aggregateRating,omitempty"`
	Attendee                  interface{}     `json:"attendee,omitempty"`  // Person, Organization
	Organizer                 interface{}     `json:"organizer,omitempty"` // Organization, Person
	Attendees                 interface{}     `json:"attendees,omitempty"` // Organization, Person
	About                     Thing           `json:"about,omitempty"`
	Performers                interface{}     `json:"performers,omitempty"` // Organization, Person
	WorkPerformed             CreativeWork    `json:"workPerformed,omitempty"`
	Audience                  Audience        `json:"audience,omitempty"`
	EndDate                   interface{}     `json:"endDate,omitempty"`  // Date, DateTime
	Location                  interface{}     `json:"location,omitempty"` // Place, PostalAddress, Text
	Actor                     Person          `json:"actor,omitempty"`
	Director                  Person          `json:"director,omitempty"`
	Translator                interface{}     `json:"translator,omitempty"` // Organization, Person
	RemainingAttendeeCapacity Integer         `json:"remainingAttendeeCapacity,omitempty"`
	SubEvents                 Event           `json:"subEvents,omitempty"`
	StartDate                 interface{}     `json:"startDate,omitempty"`   // DateTime, Date
	Contributor               interface{}     `json:"contributor,omitempty"` // Person, Organization
	SuperEvent                Event           `json:"superEvent,omitempty"`
	Composer                  interface{}     `json:"composer,omitempty"` // Organization, Person
	RecordedIn                CreativeWork    `json:"recordedIn,omitempty"`
	WorkFeatured              CreativeWork    `json:"workFeatured,omitempty"`
	Funder                    interface{}     `json:"funder,omitempty"`    // Person, Organization
	Performer                 interface{}     `json:"performer,omitempty"` // Person, Organization
	EventStatus               EventStatusType `json:"eventStatus,omitempty"`
	MaximumAttendeeCapacity   Integer         `json:"maximumAttendeeCapacity,omitempty"`
	IsAccessibleForFree       Boolean         `json:"isAccessibleForFree,omitempty"`
	Review                    Review          `json:"review,omitempty"`
	Thing
}

type VideoGame struct {
	PlayMode     GamePlayMode `json:"playMode,omitempty"`
	Director     Person       `json:"director,omitempty"`
	CheatCode    CreativeWork `json:"cheatCode,omitempty"`
	Directors    Person       `json:"directors,omitempty"`
	MusicBy      interface{}  `json:"musicBy,omitempty"` // MusicGroup, Person
	Trailer      VideoObject  `json:"trailer,omitempty"`
	Actor        Person       `json:"actor,omitempty"`
	GameServer   GameServer   `json:"gameServer,omitempty"`
	GamePlatform interface{}  `json:"gamePlatform,omitempty"` // Thing, Text, URL
	Actors       Person       `json:"actors,omitempty"`
	GameTip      CreativeWork `json:"gameTip,omitempty"`
	Game
	SoftwareApplication
}

type RecyclingCenter struct {
	LocalBusiness
}

type Corporation struct {
	TickerSymbol Text `json:"tickerSymbol,omitempty"`
	Organization
}

type Wholesale struct {
	DrugCostCategory
}

type Consortium struct {
	Organization
}

type CommentAction struct {
	ResultComment Comment `json:"resultComment,omitempty"`
	CommunicateAction
}

type Diet struct {
	PhysiologicalBenefits Text        `json:"physiologicalBenefits,omitempty"`
	Endorsers             interface{} `json:"endorsers,omitempty"` // Person, Organization
	Overview              Text        `json:"overview,omitempty"`
	ExpertConsiderations  Text        `json:"expertConsiderations,omitempty"`
	DietFeatures          Text        `json:"dietFeatures,omitempty"`
	Risks                 Text        `json:"risks,omitempty"`
	LifestyleModification
	CreativeWork
}

type ComicCoverArt struct {
	CoverArt
	ComicStory
}

type Skin struct {
	PhysicalExam
}

type OnlineOnly struct {
	ItemAvailability
}

type TheaterEvent struct {
	Event
}

type OrderCancelled struct {
	OrderStatus
}

type PerformAction struct {
	EntertainmentBusiness EntertainmentBusiness `json:"entertainmentBusiness,omitempty"`
	PlayAction
}

type SinglePlayer struct {
	GamePlayMode
}

type ParcelDelivery struct {
	ExpectedArrivalFrom  DateTime       `json:"expectedArrivalFrom,omitempty"`
	TrackingNumber       Text           `json:"trackingNumber,omitempty"`
	PartOfOrder          Order          `json:"partOfOrder,omitempty"`
	ExpectedArrivalUntil DateTime       `json:"expectedArrivalUntil,omitempty"`
	HasDeliveryMethod    DeliveryMethod `json:"hasDeliveryMethod,omitempty"`
	DeliveryAddress      PostalAddress  `json:"deliveryAddress,omitempty"`
	Provider             interface{}    `json:"provider,omitempty"` // Organization, Person
	ItemShipped          Product        `json:"itemShipped,omitempty"`
	TrackingUrl          URL            `json:"trackingUrl,omitempty"`
	OriginAddress        PostalAddress  `json:"originAddress,omitempty"`
	Carrier              Organization   `json:"carrier,omitempty"`
	DeliveryStatus       DeliveryEvent  `json:"deliveryStatus,omitempty"`
	Intangible
}

type ReportedDoseSchedule struct {
	DoseSchedule
}

type DrugClass struct {
	Drug Drug `json:"drug,omitempty"`
	MedicalEnumeration
}

type Taxi struct {
	Service
}

type Neurologic struct {
	MedicalSpecialty
}

type WebPage struct {
	LastReviewed       Date           `json:"lastReviewed,omitempty"`
	SignificantLinks   URL            `json:"significantLinks,omitempty"`
	MainContentOfPage  WebPageElement `json:"mainContentOfPage,omitempty"`
	ReviewedBy         interface{}    `json:"reviewedBy,omitempty"` // Organization, Person
	SignificantLink    URL            `json:"significantLink,omitempty"`
	Breadcrumb         interface{}    `json:"breadcrumb,omitempty"` // Text, BreadcrumbList
	PrimaryImageOfPage ImageObject    `json:"primaryImageOfPage,omitempty"`
	RelatedLink        URL            `json:"relatedLink,omitempty"`
	Specialty          Specialty      `json:"specialty,omitempty"`
	Speakable          interface{}    `json:"speakable,omitempty"` // SpeakableSpecification, URL
	CreativeWork
}

type FailedActionStatus struct {
	ActionStatusType
}

type PropertyValue struct {
	ValueReference       interface{} `json:"valueReference,omitempty"` // Enumeration, PropertyValue, StructuredValue, QuantitativeValue, QualitativeValue
	Value                interface{} `json:"value,omitempty"`          // StructuredValue, Number, Boolean, Text
	MaxValue             Number      `json:"maxValue,omitempty"`
	UnitCode             interface{} `json:"unitCode,omitempty"` // URL, Text
	UnitText             Text        `json:"unitText,omitempty"`
	MeasurementTechnique interface{} `json:"measurementTechnique,omitempty"` // URL, Text
	PropertyID           interface{} `json:"propertyID,omitempty"`           // URL, Text
	MinValue             Number      `json:"minValue,omitempty"`
	StructuredValue
}

type BuyAction struct {
	WarrantyPromise WarrantyPromise `json:"warrantyPromise,omitempty"`
	Vendor          interface{}     `json:"vendor,omitempty"` // Organization, Person
	Seller          interface{}     `json:"seller,omitempty"` // Person, Organization
	TradeAction
}

type MensClothingStore struct {
	Store
}

type PartiallyInForce struct {
	LegalForceStatus
}

type MusicAlbumReleaseType struct {
	Enumeration
}

type HearingImpairedSupported struct {
	ContactPointOption
}

type CT struct {
	MedicalImagingTechnique
}

type PostOffice struct {
	GovernmentOffice
}

type LegislationObject struct {
	LegislationLegalValue LegalValueLevel `json:"legislationLegalValue,omitempty"`
	Legislation
	MediaObject
}

type VeterinaryCare struct {
	MedicalOrganization
}

type ProductModel struct {
	PredecessorOf ProductModel `json:"predecessorOf,omitempty"`
	SuccessorOf   ProductModel `json:"successorOf,omitempty"`
	IsVariantOf   ProductModel `json:"isVariantOf,omitempty"`
	Product
}

type MoneyTransfer struct {
	Amount          interface{} `json:"amount,omitempty"`          //
	BeneficiaryBank interface{} `json:"beneficiaryBank,omitempty"` // Text, BankOrCreditUnion
	TransferAction
}

type WPAdBlock struct {
	WebPageElement
}

type GeospatialGeometry struct {
	GeospatiallyDisjoint   interface{} `json:"geospatiallyDisjoint,omitempty"`   // Place, GeospatialGeometry
	GeospatiallyContains   interface{} `json:"geospatiallyContains,omitempty"`   // GeospatialGeometry, Place
	GeospatiallyTouches    interface{} `json:"geospatiallyTouches,omitempty"`    // GeospatialGeometry, Place
	GeospatiallyWithin     interface{} `json:"geospatiallyWithin,omitempty"`     // GeospatialGeometry, Place
	GeospatiallyCrosses    interface{} `json:"geospatiallyCrosses,omitempty"`    // Place, GeospatialGeometry
	GeospatiallyCoveredBy  interface{} `json:"geospatiallyCoveredBy,omitempty"`  // Place, GeospatialGeometry
	GeospatiallyIntersects interface{} `json:"geospatiallyIntersects,omitempty"` // GeospatialGeometry, Place
	GeospatiallyEquals     interface{} `json:"geospatiallyEquals,omitempty"`     // GeospatialGeometry, Place
	GeospatiallyOverlaps   interface{} `json:"geospatiallyOverlaps,omitempty"`   // GeospatialGeometry, Place
	GeospatiallyCovers     interface{} `json:"geospatiallyCovers,omitempty"`     // Place, GeospatialGeometry
	Intangible
}

type ScreeningEvent struct {
	WorkPresented    Movie       `json:"workPresented,omitempty"`
	SubtitleLanguage interface{} `json:"subtitleLanguage,omitempty"` // Language, Text
	VideoFormat      Text        `json:"videoFormat,omitempty"`
	Event
}

type MedicalIndication struct {
	MedicalEntity
}

type MultiCenterTrial struct {
	MedicalTrialDesign
}

type OwnershipInfo struct {
	OwnedThrough DateTime    `json:"ownedThrough,omitempty"`
	TypeOfGood   interface{} `json:"typeOfGood,omitempty"`   // Service, Product
	AcquiredFrom interface{} `json:"acquiredFrom,omitempty"` // Person, Organization
	OwnedFrom    DateTime    `json:"ownedFrom,omitempty"`
	StructuredValue
}

type UsedCondition struct {
	OfferItemCondition
}

type OnDemandEvent struct {
	PublicationEvent
}

type RestrictedDiet struct {
	Enumeration
}

type AnalysisNewsArticle struct {
	NewsArticle
}

type LocalBusiness struct {
	PaymentAccepted    Text         `json:"paymentAccepted,omitempty"`
	PriceRange         Text         `json:"priceRange,omitempty"`
	CurrenciesAccepted Text         `json:"currenciesAccepted,omitempty"`
	BranchOf           Organization `json:"branchOf,omitempty"`
	OpeningHours       Text         `json:"openingHours,omitempty"`
	Place
	Organization
}

type UserDownloads struct {
	UserInteraction
}

type BusReservation struct {
	Reservation
}

type Eye struct {
	PhysicalExam
}

type Discontinued struct {
	ItemAvailability
}

type Pharmacy struct {
	MedicalBusiness
}

type GatedResidenceCommunity struct {
	Residence
}

type BusinessFunction struct {
	Enumeration
}

type Wednesday struct {
	DayOfWeek
}

type AudioObject struct {
	Transcript Text `json:"transcript,omitempty"`
	MediaObject
}

type InteractAction struct {
	Action
}

type RiverBodyOfWater struct {
	BodyOfWater
}

type WholesaleStore struct {
	Store
}

type Waterfall struct {
	BodyOfWater
}

type Restaurant struct {
	FoodEstablishment
}

type State struct {
	AdministrativeArea
}
