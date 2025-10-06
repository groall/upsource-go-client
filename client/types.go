package client

// Minimal DTOs extracted from docs for typed helpers we expose.

type CurrentUserResponseDTO struct {
	UserID                              string          `json:"userId"`
	Login                               string          `json:"login"`
	Name                                string          `json:"name"`
	IsServerAdmin                       bool            `json:"isServerAdmin"`
	IsGuestUser                         bool            `json:"isGuestUser"`
	IsTestUser                          bool            `json:"isTestUser"`
	CanCreateProjects                   bool            `json:"canCreateProjects"`
	EmailStatus                         EmailStatusEnum `json:"emailStatus"`
	AdminPermissionsInProjects          []string        `json:"adminPermissionsInProjects"`
	ReviewEditPermissionsInProjects     []string        `json:"reviewEditPermissionsInProjects"`
	ReviewViewPermissionsInProjects     []string        `json:"reviewViewPermissionsInProjects"`
	CodeContributePermissionsInProjects []string        `json:"codeContributePermissionsInProjects"`
	IsEULAAccepted                      *bool           `json:"isEULAAccepted"`
}

type ProjectIdListDTO struct {
	ProjectID []string `json:"projectId"`
}

type ShortProjectInfoDTO struct {
	ProjectID            string             `json:"projectId"`
	ProjectName          string             `json:"projectName"`
	HeadHash             string             `json:"headHash,omitempty"`
	IsReady              bool               `json:"isReady"`
	LastCommitDate       int64              `json:"lastCommitDate,omitempty"`
	LastCommitAuthorName string             `json:"lastCommitAuthorName,omitempty"`
	IconURL              string             `json:"iconUrl,omitempty"`
	Group                *ProjectGroupDTO   `json:"group,omitempty"`
	Founder              *ProjectFounderDTO `json:"founder,omitempty"`
	LastDayCommits       int32              `json:"lastDayCommits,omitempty"`
	LastMonthCommits     int32              `json:"lastMonthCommits,omitempty"`
	TotalCommits         int32              `json:"totalCommits,omitempty"`
	IsArchived           *bool              `json:"isArchived,omitempty"`
}

type ShortProjectInfoListDTO struct {
	Project []ShortProjectInfoDTO `json:"project"`
}

type RevisionsListRequestDTO struct {
	ProjectID    string `json:"projectId"`
	Limit        int32  `json:"limit"`
	Skip         int32  `json:"skip,omitempty"`
	RequestGraph *bool  `json:"requestGraph,omitempty"`
}

type RevisionDescriptorListDTO struct {
	Revision []RevisionInfoDTO     `json:"revision"`
	Graph    *RevisionListGraphDTO `json:"graph,omitempty"`
	HeadHash string                `json:"headHash,omitempty"`
	Query    string                `json:"query,omitempty"`
}

type VoidMessage struct {
	VoidField int32 `json:"voidField,omitempty"`
}

type ProjectIdDTO struct {
	ProjectID string `json:"projectId"`
}

type RevisionIdDTO struct {
	RevisionID string `json:"revisionId"`
}

type FileInRevisionDTO struct {
	ProjectID  string `json:"projectId"`
	RevisionID string `json:"revisionId"`
	FileName   string `json:"fileName"`
}

type RevisionInProjectDTO struct {
	ProjectID  string `json:"projectId"`
	RevisionID string `json:"revisionId"`
}

type ReviewIdDTO struct {
	ProjectID string `json:"projectId"`
	ReviewID  string `json:"reviewId"`
}

type RangeDTO struct {
	StartOffset int32 `json:"startOffset"`
	EndOffset   int32 `json:"endOffset"`
}

// From Misc.json

type WebhookEventEnum int

const (
	CommentAdded               WebhookEventEnum = 1
	ReviewCreated              WebhookEventEnum = 2
	ReviewParticipantCompleted WebhookEventEnum = 3
	NewReviewParticipant       WebhookEventEnum = 4
	RevisionAddedToReview      WebhookEventEnum = 5
	NewRevision                WebhookEventEnum = 6
	ReviewStateChanged         WebhookEventEnum = 7
	ReviewParticipantRemoved   WebhookEventEnum = 8
	MergedToDefaultBranch      WebhookEventEnum = 9
	NewBranch                  WebhookEventEnum = 10
	ReviewRemoved              WebhookEventEnum = 11
	BranchRemoved              WebhookEventEnum = 12
	ReactionToggled            WebhookEventEnum = 13
	CommentEdited              WebhookEventEnum = 14
	CommentRemoved             WebhookEventEnum = 15
	ReviewLabelChanged         WebhookEventEnum = 16
)

type TestPOP3MailboxStatusEnum int

const (
	OK               TestPOP3MailboxStatusEnum = 1
	LoginFailed      TestPOP3MailboxStatusEnum = 2
	ConnectionFailed TestPOP3MailboxStatusEnum = 3
)

type UserAgreementTextDTO struct {
	Text string `json:"text"`
}

type ExportDataResponseDTO struct {
	ServerPath string `json:"serverPath"`
}

type CheckCanCreateProjectRequestDTO struct {
	HubProjectID string `json:"hubProjectId"`
}

type CheckCanCreateProjectResponseDTO struct {
	IsAllowed bool   `json:"isAllowed"`
	Message   string `json:"message,omitempty"`
}

type GetSettingRequestDTO struct {
	Key string `json:"key"`
}

type GetProjectSettingRequestDTO struct {
	ProjectID string               `json:"projectId"`
	Request   GetSettingRequestDTO `json:"request"`
}

type GetSettingResponseDTO struct {
	Value string `json:"value,omitempty"`
}

type SetSettingRequestDTO struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SetProjectSettingRequestDTO struct {
	ProjectID string               `json:"projectId"`
	Request   SetSettingRequestDTO `json:"request"`
}

type SetProjectWebhooksRequestDTO struct {
	ProjectID string                     `json:"projectId"`
	Triggers  []ProjectWebhookTriggerDTO `json:"triggers"`
}

type ProjectWebhookTriggerDTO struct {
	Events []WebhookEventEnum `json:"events"`
	Urls   []string           `json:"urls"`
}

type TestPOP3MailboxRequestDTO struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TestPOP3MailboxResponseDTO struct {
	Status  TestPOP3MailboxStatusEnum `json:"status"`
	Message string                    `json:"message,omitempty"`
}

type ReadOnlyModeDTO struct {
	IsReadOnly  bool   `json:"isReadOnly"`
	Description string `json:"description,omitempty"`
}

type EmailBouncesResponseDTO struct {
	Emails []string `json:"emails"`
}

type UpdateUserTimezoneDTO struct {
	UserID         string `json:"userId"`
	TimezoneOffset int32  `json:"timezoneOffset"`
}

type UserActionNotificationDTO struct {
	UserID string `json:"userId"`
}

type AchievementDTO struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	Description      string `json:"description,omitempty"`
	Date             int64  `json:"date,omitempty"`
	IsUnread         bool   `json:"isUnread,omitempty"`
	IsUnlocked       bool   `json:"isUnlocked,omitempty"`
}

type AchievementsListDTO struct {
	Items []AchievementDTO `json:"items"`
}

type UserAchievementsRequestDTO struct {
	UserID string `json:"userId"`
}

type UserAchievementsCountDTO struct {
	Count int32 `json:"count"`
}

type SetMottoRequestDTO struct {
	Motto string `json:"motto,omitempty"`
}

// From Users.json

type RoleInReviewEnum int

const (
	Author   RoleInReviewEnum = 1
	Reviewer RoleInReviewEnum = 2
	Watcher  RoleInReviewEnum = 3
)

type EmailStatusEnum int

const (
	EmailStatusOK           EmailStatusEnum = 1
	EmailStatusNotVerified  EmailStatusEnum = 2
	EmailStatusNotSpecified EmailStatusEnum = 3
	EmailStatusHardBounce   EmailStatusEnum = 4
)

type UsersForMentionRequestDTO struct {
	ProjectID  string `json:"projectId"`
	ReviewID   string `json:"reviewId,omitempty"`
	RevisionID string `json:"revisionId,omitempty"`
	FileName   string `json:"fileName,omitempty"`
	Query      string `json:"query,omitempty"`
	Limit      int32  `json:"limit"`
}

type UsersForReviewRequestDTO struct {
	ReviewID ReviewIdDTO      `json:"reviewId"`
	Role     RoleInReviewEnum `json:"role"`
	Query    string           `json:"query,omitempty"`
	Limit    int32            `json:"limit"`
	Timeout  int64            `json:"timeout,omitempty"`
}

type UsersListDTO struct {
	Me                         string   `json:"me,omitempty"`
	SuggestedUsers             []string `json:"suggestedUsers"`
	SuggestedUserRelevance     []int32  `json:"suggestedUserRelevance"`
	SuggestedUserOpenedReviews []int32  `json:"suggestedUserOpenedReviews"`
	Committers                 []string `json:"committers"`
	Others                     []string `json:"others"`
	HasMore                    bool     `json:"hasMore"`
}

type ProjectUserGroupsRequestDTO struct {
	ProjectID string `json:"projectId"`
	Query     string `json:"query,omitempty"`
	Limit     int32  `json:"limit"`
}

type UserGroupsIdsListDTO struct {
	IDs []string `json:"ids"`
}

type UserGroupDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	UsersCount int32  `json:"usersCount"`
}

type UserGroupsListDTO struct {
	Groups  []UserGroupDTO `json:"groups"`
	HasMore bool           `json:"hasMore"`
}

type UserInfoRequestDTO struct {
	IDs []string `json:"ids"`
}

type UserInfoResponseDTO struct {
	Infos []FullUserInfoDTO `json:"infos"`
}

type FullUserInfoDTO struct {
	UserID      string `json:"userId"`
	Name        string `json:"name"`
	IsResolved  bool   `json:"isResolved"`
	IsMe        bool   `json:"isMe"`
	IsOnline    bool   `json:"isOnline,omitempty"`
	AvatarURL   string `json:"avatarUrl,omitempty"`
	ProfileURL  string `json:"profileUrl,omitempty"`
	Email       string `json:"email,omitempty"`
	Login       string `json:"login,omitempty"`
	AbsentUntil int64  `json:"absentUntil,omitempty"`
}

type UsersPresenceInfoResponseDTO struct {
	Presences []UserPresenceInfoDTO `json:"presences"`
}

type UserPresenceInfoDTO struct {
	UserID      string `json:"userId"`
	IsOnline    bool   `json:"isOnline"`
	AbsentUntil int64  `json:"absentUntil,omitempty"`
}

type BindVcsUsernameRequestDTO struct {
	UserID string `json:"userId"`
}

type UserProjectsRequestDTO struct {
	UserID                string `json:"userId"`
	CommitTimestampFilter int64  `json:"commitTimestampFilter"`
}

type UserProjectInfoDTO struct {
	ProjectID      string `json:"projectId"`
	LastCommitTime int64  `json:"lastCommitTime"`
}

type UserProjectsResponseDTO struct {
	Projects []UserProjectInfoDTO `json:"projects"`
}

type UserAbsenceRequestDTO struct {
	AbsentUntil int64  `json:"absentUntil,omitempty"`
	UserID      string `json:"userId,omitempty"`
}

type FindUsersRequestDTO struct {
	ProjectID string `json:"projectId,omitempty"`
	Pattern   string `json:"pattern"`
	Limit     int32  `json:"limit"`
}

type UsersForReviewDTO struct {
	Result        UsersListDTO `json:"result"`
	RelevantUser  []string     `json:"relevantUser"`
	UserRelevance []int32      `json:"userRelevance"`
}

// From Errors.json

type RPCErrors int

const (
	RPCOK                        RPCErrors = 0
	RPCUnknown                   RPCErrors = 1
	RPCInvalidRequest            RPCErrors = 100
	RPCAuthenticationError       RPCErrors = 101
	RPCAccessDenied              RPCErrors = 102
	RPCSecurityTokenExpired      RPCErrors = 103
	RPCSecurityTokenRequired     RPCErrors = 104
	RPCRequestRejected           RPCErrors = 105
	RPCNotFound                  RPCErrors = 106
	RPCMethodNotSupported        RPCErrors = 107
	RPCInternalError             RPCErrors = 108
	RPCTaskCancelled             RPCErrors = 109
	RPCServerReadOnly            RPCErrors = 110
	RPCNotSupportedInProduction  RPCErrors = 1000
	RPCUpsourceException         RPCErrors = 1001
	RPCFileNotFoundInRevision    RPCErrors = 1002
	RPCRevisionNotIndexed        RPCErrors = 1003
	RPCUpsourceNotReadyException RPCErrors = 2001
	RPCFileIsTooBig              RPCErrors = 2002
)

// From Events.json

type ReviewState int

const (
	ReviewStateOpen   ReviewState = 0
	ReviewStateClosed ReviewState = 1
)

type ParticipantState int

const (
	ParticipantStateUnread   ParticipantState = 0
	ParticipantStateRead     ParticipantState = 1
	ParticipantStateAccepted ParticipantState = 2
	ParticipantStateRejected ParticipantState = 3
)

type NotificationReason int

const (
	NotificationReasonUnknown                  NotificationReason = 0
	NotificationReasonCommentInAuthorFeed      NotificationReason = 1
	NotificationReasonNotifyCommitAuthor       NotificationReason = 2
	NotificationReasonMention                  NotificationReason = 3
	NotificationReasonHashTagSubscription      NotificationReason = 4
	NotificationReasonDiscussionIsStarred      NotificationReason = 5
	NotificationReasonParticipatedInDiscussion NotificationReason = 6
	NotificationReasonParticipatedInReview     NotificationReason = 7
	NotificationReasonReply                    NotificationReason = 8
)

type UserIdBean struct {
	UserID    string `json:"userId,omitempty"`
	UserName  string `json:"userName,omitempty"`
	UserEmail string `json:"userEmail,omitempty"`
}

type FeedEventBean struct {
	UserID       *UserIdBean  `json:"userId,omitempty"`
	UserIDs      []UserIdBean `json:"userIds"`
	ReviewNumber int32        `json:"reviewNumber,omitempty"`
	ReviewID     string       `json:"reviewId,omitempty"`
	Date         int64        `json:"date"`
	Actor        UserIdBean   `json:"actor"`
	FeedEventID  string       `json:"feedEventId,omitempty"`
}

type ReviewCreatedFeedEventBean struct {
	Base      FeedEventBean `json:"base"`
	Revisions []string      `json:"revisions"`
	Branch    string        `json:"branch,omitempty"`
}

type DiscussionFeedEventBean struct {
	Base               FeedEventBean      `json:"base"`
	NotificationReason NotificationReason `json:"notificationReason"`
	DiscussionID       string             `json:"discussionId"`
	CommentID          string             `json:"commentId,omitempty"`
	IsEdit             *bool              `json:"isEdit,omitempty"`
	ResolveAction      *bool              `json:"resolveAction,omitempty"`
	CommentText        string             `json:"commentText,omitempty"`
	IsDeletion         *bool              `json:"isDeletion,omitempty"`
	ProjectID          string             `json:"projectId,omitempty"`
	RevisionID         string             `json:"revisionId,omitempty"`
	FileName           string             `json:"fileName,omitempty"`
}

type NewParticipantInReviewFeedEventBean struct {
	Base        FeedEventBean    `json:"base"`
	Participant UserIdBean       `json:"participant"`
	Role        RoleInReviewEnum `json:"role"`
}

type ParticipantStateChangedFeedEventBean struct {
	Base        FeedEventBean    `json:"base"`
	Participant UserIdBean       `json:"participant"`
	OldState    ParticipantState `json:"oldState"`
	NewState    ParticipantState `json:"newState"`
}

type RemovedParticipantFromReviewFeedEventBean struct {
	Base        FeedEventBean    `json:"base"`
	Participant UserIdBean       `json:"participant"`
	FormerRole  RoleInReviewEnum `json:"formerRole"`
}

type ReviewRemovedFeedEventBean struct {
	Base     FeedEventBean `json:"base"`
	ReviewID string        `json:"reviewId"`
}

type ReviewStateChangedFeedEventBean struct {
	Base     FeedEventBean `json:"base"`
	OldState ReviewState   `json:"oldState"`
	NewState ReviewState   `json:"newState"`
}

type RevisionAddedToReviewFeedEventBean struct {
	Base        FeedEventBean `json:"base"`
	RevisionID  string        `json:"revisionId,omitempty"`
	RevisionIDs []string      `json:"revisionIds"`
}

type RevisionRemovedFromReviewFeedEventBean struct {
	Base        FeedEventBean `json:"base"`
	RevisionIDs []string      `json:"revisionIds"`
}

type ReviewStoppedBranchTrackingFeedEventBean struct {
	Base   FeedEventBean `json:"base"`
	Branch string        `json:"branch"`
}

type ReviewSquashedFeedEventBean struct {
	Base       FeedEventBean `json:"base"`
	RevisionID string        `json:"revisionId"`
}

type PullRequestMergedFeedEventBean struct {
	Base        FeedEventBean `json:"base"`
	PullRequest string        `json:"pullRequest"`
}

type ReviewDeadlineUpdatedFeedEventBean struct {
	Base     FeedEventBean `json:"base"`
	Deadline *int64        `json:"deadline,omitempty"`
}

type NewRevisionEventBean struct {
	RevisionID string   `json:"revisionId"`
	Branches   []string `json:"branches"`
	Author     string   `json:"author,omitempty"`
	Message    string   `json:"message,omitempty"`
	Date       *int64   `json:"date,omitempty"`
}

type MergedToDefaultBranchEventBean struct {
	CommitsCount int32    `json:"commitsCount"`
	Branches     []string `json:"branches"`
}

type NewBranchEventBean struct {
	Name string `json:"name"`
}

type DeleteBranchEventBean struct {
	Name string `json:"name"`
}

type ReactionToggledEventBean struct {
	DiscussionID string     `json:"discussionId"`
	CommentID    string     `json:"commentId"`
	ReactionID   string     `json:"reactionId"`
	WasAdded     bool       `json:"wasAdded"`
	Actor        UserIdBean `json:"actor"`
}

type ReviewLabelChangedEventBean struct {
	ReviewID  string     `json:"reviewId"`
	LabelID   string     `json:"labelId"`
	LabelName string     `json:"labelName"`
	WasAdded  bool       `json:"wasAdded"`
	Actor     UserIdBean `json:"actor"`
}

// From Projects.json

type RevisionStateEnum int

const (
	None     RevisionStateEnum = 1
	Found    RevisionStateEnum = 2
	Imported RevisionStateEnum = 3
)

type ReviewStateEnum int

const (
	ReviewStateEnumOpen   ReviewStateEnum = 1
	ReviewStateEnumClosed ReviewStateEnum = 2
)

type ReadEnum int

const (
	ReadEnumRead           ReadEnum = 1
	ReadEnumUnread         ReadEnum = 2
	ReadEnumManuallyUnread ReadEnum = 3
)

type ParticipantStateEnum int

const (
	ParticipantStateEnumUnread   ParticipantStateEnum = 1
	ParticipantStateEnumRead     ParticipantStateEnum = 2
	ParticipantStateEnumAccepted ParticipantStateEnum = 3
	ParticipantStateEnumRejected ParticipantStateEnum = 4
)

type FeedTypeEnum int

const (
	FeedTypeEnumFeed   FeedTypeEnum = 1
	FeedTypeEnumReview FeedTypeEnum = 2
)

type DiffTypeEnum int

const (
	DiffTypeEnumAdded     DiffTypeEnum = 1
	DiffTypeEnumRemoved   DiffTypeEnum = 2
	DiffTypeEnumModified  DiffTypeEnum = 3
	DiffTypeEnumCommented DiffTypeEnum = 4
)

type TestConnectionStatusEnum int

const (
	TestConnectionStatusEnumSuccess TestConnectionStatusEnum = 1
	TestConnectionStatusEnumFailure TestConnectionStatusEnum = 2
	TestConnectionStatusEnumTimeout TestConnectionStatusEnum = 3
)

type ProjectStateEnum int

const (
	ProjectStateEnumNotStarted   ProjectStateEnum = 1
	ProjectStateEnumInitializing ProjectStateEnum = 2
	ProjectStateEnumSuccess      ProjectStateEnum = 3
	ProjectStateEnumFailure      ProjectStateEnum = 4
)

type AnalyzerProblemSeverityEnum int

const (
	AnalyzerProblemSeverityEnumInfo    AnalyzerProblemSeverityEnum = 1
	AnalyzerProblemSeverityEnumWarning AnalyzerProblemSeverityEnum = 2
	AnalyzerProblemSeverityEnumError   AnalyzerProblemSeverityEnum = 3
)

type BuildStatusEnum int

const (
	BuildStatusEnumSuccess    BuildStatusEnum = 1
	BuildStatusEnumFailed     BuildStatusEnum = 2
	BuildStatusEnumInProgress BuildStatusEnum = 3
)

type SyncResultEnum int

const (
	SyncResultEnumOk                     SyncResultEnum = 1
	SyncResultEnumProjectNotSynchronized SyncResultEnum = 2
	SyncResultEnumCommentNotSynchronized SyncResultEnum = 3
	SyncResultEnumReviewNotSynchronized  SyncResultEnum = 4
	SyncResultEnumAccessTokenNotProvided SyncResultEnum = 5
	SyncResultEnumGithubMisconfiguration SyncResultEnum = 6
	SyncResultEnumUnsupportedMethod      SyncResultEnum = 7
	SyncResultEnumNotFound               SyncResultEnum = 8
	SyncResultEnumInternalError          SyncResultEnum = 9
	SyncResultEnumRateLimitExceeded      SyncResultEnum = 10
	SyncResultEnumIoError                SyncResultEnum = 11
)

type AppConfigurationProblemEnum int

const (
	AppConfigurationProblemEnumIncompatibleHub   AppConfigurationProblemEnum = 1
	AppConfigurationProblemEnumMissingAuthModule AppConfigurationProblemEnum = 2
)

type OwnershipSummaryEnum int

const (
	OwnershipSummaryEnumOK        OwnershipSummaryEnum = 1
	OwnershipSummaryEnumONE_MAJOR OwnershipSummaryEnum = 2
	OwnershipSummaryEnumALL_MINOR OwnershipSummaryEnum = 3
)

type ConflictTypeEnum int

const (
	ConflictTypeEnumNO_CONFLICT         ConflictTypeEnum = 1
	ConflictTypeEnumCAN_BE_RESOLVED     ConflictTypeEnum = 2
	ConflictTypeEnumCAN_NOT_BE_RESOLVED ConflictTypeEnum = 3
)

type RevisionReachabilityEnum int

const (
	RevisionReachabilityEnumReachable    RevisionReachabilityEnum = 1
	RevisionReachabilityEnumUnknown      RevisionReachabilityEnum = 2
	RevisionReachabilityEnumNotReachable RevisionReachabilityEnum = 3
)

type ReviewFeedSortEnum int

const (
	ReviewFeedSortEnumNatural     ReviewFeedSortEnum = 1
	ReviewFeedSortEnumLastUpdated ReviewFeedSortEnum = 2
)

type ProjectGroupDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type FindProjectsRequestDTO struct {
	Pattern string `json:"pattern"`
	Limit   int32  `json:"limit"`
	IsExact *bool  `json:"isExact,omitempty"`
}

type ProjectFounderDTO struct {
	UserID string `json:"userId"`
	Date   int64  `json:"date"`
}

type ProjectInfoDTO struct {
	ProjectName                string                          `json:"projectName"`
	ProjectID                  string                          `json:"projectId"`
	HeadHash                   string                          `json:"headHash"`
	CodeReviewIdPattern        string                          `json:"codeReviewIdPattern"`
	ExternalLinks              []ExternalLinkDTO               `json:"externalLinks"`
	IssueTrackerConnections    []ExternalLinkDTO               `json:"issueTrackerConnections"`
	ProjectModelType           string                          `json:"projectModelType"`
	DefaultEffectiveCharset    string                          `json:"defaultEffectiveCharset"`
	DefaultBranch              string                          `json:"defaultBranch,omitempty"`
	IssueTrackerDetails        *IssueTrackerProjectDetailsDTO  `json:"issueTrackerDetails,omitempty"`
	IsConnectedToGithub        bool                            `json:"isConnectedToGithub"`
	IsConnectedToGitlab        bool                            `json:"isConnectedToGitlab"`
	IconURL                    string                          `json:"iconUrl,omitempty"`
	Group                      *ProjectGroupDTO                `json:"group,omitempty"`
	IsArchived                 *bool                           `json:"isArchived,omitempty"`
	IssueTrackerProjectDetails []IssueTrackerProjectDetailsDTO `json:"issueTrackerProjectDetails"`
}

type ProjectReadmeResponseDTO struct {
	Text       string `json:"text"`
	FileName   string `json:"fileName"`
	RevisionID string `json:"revisionId"`
}

type CodeReviewPattern struct {
	ProjectID string `json:"projectId"`
	Pattern   string `json:"pattern"`
}

type CodeReviewPatternsDTO struct {
	Patterns []CodeReviewPattern `json:"patterns"`
}

type ExternalLinkDTO struct {
	URL    string `json:"url"`
	Prefix string `json:"prefix"`
}

type ProjectSettingsDTO struct {
	ProjectName                     string                           `json:"projectName"`
	VcsSettings                     string                           `json:"vcsSettings,omitempty"`
	CheckIntervalSeconds            int64                            `json:"checkIntervalSeconds,omitempty"`
	ProjectModel                    ProjectModel                     `json:"projectModel"`
	CodeReviewIdPattern             string                           `json:"codeReviewIdPattern"`
	RunInspections                  *bool                            `json:"runInspections,omitempty"`
	InspectionProfileName           string                           `json:"inspectionProfileName,omitempty"`
	MavenSettings                   string                           `json:"mavenSettings,omitempty"`
	MavenProfiles                   string                           `json:"mavenProfiles,omitempty"`
	MavenJdkName                    string                           `json:"mavenJdkName,omitempty"`
	ModelConversionSystemProperties string                           `json:"modelConversionSystemProperties,omitempty"`
	GradleProperties                string                           `json:"gradleProperties,omitempty"`
	GradleInitScript                string                           `json:"gradleInitScript,omitempty"`
	ExternalLinks                   []ExternalLinkDTO                `json:"externalLinks"`
	IssueTrackerProviderSettings    *IssueTrackerProviderSettingsDTO `json:"issueTrackerProviderSettings,omitempty"`
	UserManagementURL               string                           `json:"userManagementUrl,omitempty"`
	DefaultEncoding                 string                           `json:"defaultEncoding,omitempty"`
	DefaultBranch                   string                           `json:"defaultBranch,omitempty"`
	IgnoredFilesInReview            []string                         `json:"ignoredFilesInReview"`
	SkipFileContentsImport          []string                         `json:"skipFileContentsImport"`
	JavascriptLanguageLevel         string                           `json:"javascriptLanguageLevel,omitempty"`
	PhpLanguageLevel                string                           `json:"phpLanguageLevel,omitempty"`
	PhpExternalDependencies         []string                         `json:"phpExternalDependencies"`
	PhpComposerScript               []string                         `json:"phpComposerScript"`
	PythonLanguageLevel             string                           `json:"pythonLanguageLevel,omitempty"`
	BuildStatusReceiveToken         string                           `json:"buildStatusReceiveToken,omitempty"`
	AuthorCanCloseReview            *bool                            `json:"authorCanCloseReview,omitempty"`
	AuthorCanDeleteReview           *bool                            `json:"authorCanDeleteReview,omitempty"`
	LimitResolveDiscussion          *bool                            `json:"limitResolveDiscussion,omitempty"`
	ReviewDeadlineHours             int32                            `json:"reviewDeadlineHours,omitempty"`
	AddMergeCommitsToBranchReview   *bool                            `json:"addMergeCommitsToBranchReview,omitempty"`
	IsArchived                      *bool                            `json:"isArchived,omitempty"`
}

type CreateProjectRequestDTO struct {
	NewProjectID string                 `json:"newProjectId"`
	Settings     ProjectSettingsDTO     `json:"settings"`
	Custom       []SetSettingRequestDTO `json:"custom"`
}

type EditProjectRequestDTO struct {
	ProjectID string             `json:"projectId"`
	Settings  ProjectSettingsDTO `json:"settings"`
}

type ProjectModel struct {
	Type         string `json:"type"`
	PathToModel  string `json:"pathToModel,omitempty"`
	DefaultJdkID string `json:"defaultJdkId,omitempty"`
}

type TestConnectionRequestDTO struct {
	ProjectID   string `json:"projectId,omitempty"`
	VcsSettings string `json:"vcsSettings"`
}

type TestConnectionResponseDTO struct {
	Status  TestConnectionStatusEnum `json:"status"`
	Message string                   `json:"message,omitempty"`
}

type ReactionTargetDTO struct {
	ProjectID    string `json:"projectId"`
	DiscussionID string `json:"discussionId,omitempty"`
	CommentID    string `json:"commentId,omitempty"`
	RevisionID   string `json:"revisionId,omitempty"`
	ReviewID     string `json:"reviewId,omitempty"`
}

type ReactionDTO struct {
	ID       string   `json:"id"`
	UsersIDs []string `json:"usersIds"`
}

type SimpleDiscussionCounterDTO struct {
	Count           int32 `json:"count"`
	HasUnresolved   bool  `json:"hasUnresolved"`
	UnresolvedCount int32 `json:"unresolvedCount"`
	ResolvedCount   int32 `json:"resolvedCount"`
}

type RevisionInfoDTO struct {
	ProjectID             string                   `json:"projectId"`
	RevisionID            string                   `json:"revisionId"`
	RevisionDate          int64                    `json:"revisionDate"`
	EffectiveRevisionDate int64                    `json:"effectiveRevisionDate"`
	RevisionCommitMessage string                   `json:"revisionCommitMessage"`
	State                 RevisionStateEnum        `json:"state"`
	VcsRevisionID         string                   `json:"vcsRevisionId"`
	ShortRevisionID       string                   `json:"shortRevisionId"`
	AuthorID              string                   `json:"authorId"`
	Reachability          RevisionReachabilityEnum `json:"reachability"`
	Tags                  []string                 `json:"tags"`
	BranchHeadLabel       []string                 `json:"branchHeadLabel"`
	ParentRevisionIDs     []string                 `json:"parentRevisionIds"`
}

type MatchingRevisionsResponseDTO struct {
	Revision []RevisionInfoDTO `json:"revision"`
}

type RevisionsInReviewResponseDTO struct {
	AllRevisions              RevisionDescriptorListDTO `json:"allRevisions"`
	NewRevisions              RevisionsSetDTO           `json:"newRevisions"`
	HasMissingRevisions       bool                      `json:"hasMissingRevisions"`
	CanSquash                 bool                      `json:"canSquash"`
	BranchHint                string                    `json:"branchHint,omitempty"`
	CanTrackBranch            *bool                     `json:"canTrackBranch,omitempty"`
	CanCherryPickWithTeamCity *bool                     `json:"canCherryPickWithTeamCity,omitempty"`
}

type RevisionSuggestDTO struct {
	RevisionID string  `json:"revisionId"`
	Score      float32 `json:"score"`
}

type SuggestedRevisionListDTO struct {
	Suggest []RevisionSuggestDTO `json:"suggest"`
}

type RevisionsSetDTO struct {
	Revisions []string `json:"revisions"`
	SelectAll *bool    `json:"selectAll,omitempty"`
}

type ReviewSummaryChangesRequestDTO struct {
	ReviewID  ReviewIdDTO      `json:"reviewId"`
	Revisions *RevisionsSetDTO `json:"revisions,omitempty"`
}

type ReviewInspectionsDiffRequestDTO struct {
	ReviewID  ReviewIdDTO      `json:"reviewId"`
	Revisions *RevisionsSetDTO `json:"revisions,omitempty"`
}

type RevisionsListFilteredRequestDTO struct {
	ProjectID    string `json:"projectId"`
	Query        string `json:"query"`
	Limit        int32  `json:"limit"`
	Skip         int32  `json:"skip,omitempty"`
	RequestGraph *bool  `json:"requestGraph,omitempty"`
}

type UserRevisionsListRequestDTO struct {
	UserID string `json:"userId"`
	Limit  int32  `json:"limit"`
	Query  string `json:"query,omitempty"`
}

type RevisionsListUpdateRequestDTO struct {
	ProjectID        string   `json:"projectId"`
	Limit            int32    `json:"limit"`
	Hash             string   `json:"hash,omitempty"`
	MessageSubstring []string `json:"messageSubstring"`
}

type RevisionsListUpdateResponseDTO struct {
	Revision []RevisionInfoDTO `json:"revision"`
	Hash     string            `json:"hash"`
}

type ProjectItemsListDTO struct {
	Items []ProjectTreeItemDTO `json:"items"`
}

type ProjectTreeItemDTO struct {
	DisplayName string `json:"displayName"`
	DisplayType string `json:"displayType"`
	IsDirectory bool   `json:"isDirectory"`
	IsModule    bool   `json:"isModule"`
	FileID      string `json:"fileId"`
	HasChildren bool   `json:"hasChildren"`
}

type RevisionListGraphDTO struct {
	Width int32                     `json:"width"`
	Rows  []RevisionListGraphRowDTO `json:"rows"`
}

type RevisionListGraphRowDTO struct {
	Nodes []RevisionListGraphNodeDTO `json:"nodes"`
	Edges []RevisionListGraphEdgeDTO `json:"edges"`
}

type RevisionListGraphNodeDTO struct {
	Position int32 `json:"position"`
	Color    int32 `json:"color"`
}

type RevisionListGraphEdgeDTO struct {
	Position   int32 `json:"position"`
	ToPosition int32 `json:"toPosition"`
	IsUp       bool  `json:"isUp"`
	IsSolid    bool  `json:"isSolid"`
	HasArrow   bool  `json:"hasArrow"`
	Color      int32 `json:"color"`
}

type AnchorDTO struct {
	Range            *RangeDTO `json:"range,omitempty"`
	FileID           string    `json:"fileId,omitempty"`
	RevisionID       string    `json:"revisionId,omitempty"`
	InlineInRevision string    `json:"inlineInRevision,omitempty"`
}

type CommentDTO struct {
	DiscussionID   string          `json:"discussionId"`
	CommentID      string          `json:"commentId"`
	Text           string          `json:"text"`
	AuthorID       string          `json:"authorId"`
	Date           int64           `json:"date"`
	ParentID       string          `json:"parentId,omitempty"`
	IsEditable     bool            `json:"isEditable"`
	MarkupType     string          `json:"markupType,omitempty"`
	IsSynchronized bool            `json:"isSynchronized"`
	SyncResult     *SyncResultEnum `json:"syncResult,omitempty"`
	IsRead         bool            `json:"isRead"`
	Reactions      []ReactionDTO   `json:"reactions"`
}

type UpdateCommentResponseDTO struct {
	SyncResult *SyncResultEnum `json:"syncResult,omitempty"`
}

type RemoveCommentResponseDTO struct {
	SyncResult *SyncResultEnum `json:"syncResult,omitempty"`
}

type CloseReviewResponseDTO struct {
	SyncResult *SyncResultEnum `json:"syncResult,omitempty"`
}

type RenameReviewResponseDTO struct {
	SyncResult *SyncResultEnum `json:"syncResult,omitempty"`
}

type SanitizeHTMLRequestDTO struct {
	HTML          string `json:"html"`
	IsBasicMarkup bool   `json:"isBasicMarkup"`
}

type SanitizeHTMLResponseDTO struct {
	SanitizedHTML string `json:"sanitizedHTML"`
}

type ReviewListDTO struct {
	Reviews    []ReviewDescriptorDTO `json:"reviews"`
	HasMore    bool                  `json:"hasMore"`
	TotalCount int32                 `json:"totalCount"`
}

type ReviewsRequestDTO struct {
	Limit     int32  `json:"limit"`
	Query     string `json:"query,omitempty"`
	SortBy    string `json:"sortBy,omitempty"`
	ProjectID string `json:"projectId,omitempty"`
	Skip      int32  `json:"skip,omitempty"`
}

type StartBranchTrackingRequestDTO struct {
	ReviewID ReviewIdDTO `json:"reviewId"`
	Branch   string      `json:"branch"`
}

type StopBranchTrackingRequestDTO struct {
	ReviewID ReviewIdDTO `json:"reviewId"`
	Branch   string      `json:"branch,omitempty"`
}

type DiscussionInReviewDTO struct {
	DiscussionID string      `json:"discussionId"`
	ReviewID     ReviewIdDTO `json:"reviewId"`
}

type ShortReviewInfoDTO struct {
	ReviewID       ReviewIdDTO        `json:"reviewId"`
	Title          string             `json:"title"`
	State          ReviewStateEnum    `json:"state"`
	Branch         []string           `json:"branch"`
	CompletionRate *CompletionRateDTO `json:"completionRate,omitempty"`
	Labels         []LabelDTO         `json:"labels"`
}

type CompletionRateDTO struct {
	CompletedCount int32 `json:"completedCount"`
	ReviewersCount int32 `json:"reviewersCount"`
	HasConcern     bool  `json:"hasConcern"`
}

type ReviewDescriptorDTO struct {
	ReviewID          ReviewIdDTO                 `json:"reviewId"`
	Title             string                      `json:"title"`
	Description       string                      `json:"description,omitempty"`
	Participants      []ParticipantInReviewDTO    `json:"participants"`
	State             ReviewStateEnum             `json:"state"`
	IsUnread          *bool                       `json:"isUnread,omitempty"`
	IsReadyToClose    *bool                       `json:"isReadyToClose,omitempty"`
	Branch            []string                    `json:"branch"`
	Issue             []IssueIdDTO                `json:"issue"`
	IsRemoved         *bool                       `json:"isRemoved,omitempty"`
	CreatedAt         int64                       `json:"createdAt"`
	CreatedBy         string                      `json:"createdBy,omitempty"`
	UpdatedAt         int64                       `json:"updatedAt"`
	CompletionRate    *CompletionRateDTO          `json:"completionRate,omitempty"`
	DiscussionCounter *SimpleDiscussionCounterDTO `json:"discussionCounter,omitempty"`
	Deadline          *int64                      `json:"deadline,omitempty"`
	IsMuted           *bool                       `json:"isMuted,omitempty"`
	Labels            []LabelDTO                  `json:"labels"`
	MergeFromBranch   string                      `json:"mergeFromBranch,omitempty"`
	MergeToBranch     string                      `json:"mergeToBranch,omitempty"`
}

type LabelDTO struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	ColorID string `json:"colorId,omitempty"`
}

type LabelsRequestDTO struct {
	ProjectID string `json:"projectId,omitempty"`
}

type LabelsListDTO struct {
	HidePredefinedLabels bool       `json:"hidePredefinedLabels"`
	PredefinedLabels     []LabelDTO `json:"predefinedLabels"`
	CustomLabels         []LabelDTO `json:"customLabels"`
}

type EditLabelsRequestDTO struct {
	ProjectID string   `json:"projectId"`
	LabelID   []string `json:"labelId"`
}

type EditLabelRequestDTO struct {
	ProjectID string   `json:"projectId"`
	Label     LabelDTO `json:"label"`
}

type HidePredefinedLabelsRequestDTO struct {
	ProjectID string `json:"projectId"`
	DoHide    bool   `json:"doHide"`
}

type DiscussionInFileDTO struct {
	DiscussionID string              `json:"discussionId"`
	Anchor       AnchorDTO           `json:"anchor"`
	Origin       *AnchorDTO          `json:"origin,omitempty"`
	Comments     []CommentDTO        `json:"comments"`
	Read         *ReadEnum           `json:"read,omitempty"`
	IsStarred    *bool               `json:"isStarred,omitempty"`
	Review       *ShortReviewInfoDTO `json:"review,omitempty"`
	Labels       []LabelDTO          `json:"labels"`
	Issue        string              `json:"issue,omitempty"`
	IsResolved   *bool               `json:"isResolved,omitempty"`
	SyncResult   *SyncResultEnum     `json:"syncResult,omitempty"`
}

type CreateReviewRequestDTO struct {
	ProjectID       string   `json:"projectId"`
	Title           string   `json:"title,omitempty"`
	Revisions       []string `json:"revisions"`
	Branch          string   `json:"branch,omitempty"`
	MergeFromBranch string   `json:"mergeFromBranch,omitempty"`
	MergeToBranch   string   `json:"mergeToBranch,omitempty"`
}

type CloseReviewRequestDTO struct {
	ReviewID  ReviewIdDTO `json:"reviewId"`
	IsFlagged bool        `json:"isFlagged"`
}

type RenameReviewRequestDTO struct {
	ReviewID ReviewIdDTO `json:"reviewId"`
	Text     string      `json:"text"`
}

type EditReviewDescriptionRequestDTO struct {
	ReviewID ReviewIdDTO `json:"reviewId"`
	Text     string      `json:"text"`
}

type EditReviewDescriptionResponseDTO struct {
	SyncResult *SyncResultEnum `json:"syncResult,omitempty"`
}

type RevisionsInReviewDTO struct {
	ReviewID   ReviewIdDTO `json:"reviewId"`
	RevisionID []string    `json:"revisionId"`
}

type ParticipantInReviewRequestDTO struct {
	ReviewID    ReviewIdDTO            `json:"reviewId"`
	Participant ParticipantInReviewDTO `json:"participant"`
}

type AddGroupToReviewRequestDTO struct {
	ReviewID ReviewIdDTO      `json:"reviewId"`
	GroupID  string           `json:"groupId"`
	Role     RoleInReviewEnum `json:"role"`
}

type ParticipantInReviewDTO struct {
	UserID string                `json:"userId"`
	Role   RoleInReviewEnum      `json:"role"`
	State  *ParticipantStateEnum `json:"state,omitempty"`
}

type UpdateParticipantInReviewRequestDTO struct {
	ReviewID ReviewIdDTO          `json:"reviewId"`
	State    ParticipantStateEnum `json:"state"`
	UserID   string               `json:"userId,omitempty"`
}

type UpdateParticipantInReviewResponseDTO struct {
	SyncResult *SyncResultEnum `json:"syncResult,omitempty"`
}

type AddGroupToReviewResponseDTO struct {
	AddedUsersCount int32 `json:"addedUsersCount"`
}

type ToggleReviewMutedRequestDTO struct {
	ReviewID ReviewIdDTO `json:"reviewId"`
	DoMute   bool        `json:"doMute"`
}

type ReviewDeadlineRequestDTO struct {
	ReviewID ReviewIdDTO `json:"reviewId"`
	Deadline *int64      `json:"deadline,omitempty"`
}

type ReviewTargetBranchDTO struct {
	ReviewID     ReviewIdDTO `json:"reviewId"`
	TargetBranch string      `json:"targetBranch"`
}

type SetReviewTargetBranchResponseDTO struct {
	SyncResult *SyncResultEnum `json:"syncResult,omitempty"`
}

type ResolveDiscussionResponseDTO struct {
	SyncResult *SyncResultEnum `json:"syncResult,omitempty"`
}

type UpdateReviewLabelRequestDTO struct {
	ProjectID string       `json:"projectId"`
	ReviewID  *ReviewIdDTO `json:"reviewId,omitempty"`
	Label     LabelDTO     `json:"label"`
}

type UpdateReviewLabelResponseDTO struct {
	Label      LabelDTO        `json:"label"`
	SyncResult *SyncResultEnum `json:"syncResult,omitempty"`
}

type FileInReviewDTO struct {
	ReviewID ReviewIdDTO       `json:"reviewId"`
	File     FileInRevisionDTO `json:"file"`
}

type DiscussionsInFileDTO struct {
	Discussions []DiscussionInFileDTO `json:"discussions"`
}

type DiscussionsGroupDTO struct {
	FileID      string               `json:"fileId,omitempty"`
	Discussions DiscussionsInFileDTO `json:"discussions"`
}

type DiscussionsInRevisionDTO struct {
	Groups []DiscussionsGroupDTO `json:"groups"`
}

type CreateDiscussionRequestDTO struct {
	Anchor     AnchorDTO    `json:"anchor"`
	ReviewID   *ReviewIdDTO `json:"reviewId,omitempty"`
	Text       string       `json:"text"`
	ProjectID  string       `json:"projectId"`
	MarkupType string       `json:"markupType,omitempty"`
	Labels     []LabelDTO   `json:"labels"`
}

type DiscussionIdDTO struct {
	ProjectID    string `json:"projectId"`
	DiscussionID string `json:"discussionId"`
}

type ResolveDiscussionRequestDTO struct {
	ProjectID    string `json:"projectId"`
	DiscussionID string `json:"discussionId"`
	IsResolved   bool   `json:"isResolved"`
	Revision     string `json:"revision,omitempty"`
}

type UpdateDiscussionLabelRequestDTO struct {
	ProjectID    string   `json:"projectId"`
	DiscussionID string   `json:"discussionId"`
	Label        LabelDTO `json:"label"`
}

type AddCommentRequestDTO struct {
	ProjectID    string `json:"projectId"`
	DiscussionID string `json:"discussionId"`
	Text         string `json:"text"`
	ParentID     string `json:"parentId,omitempty"`
	MarkupType   string `json:"markupType,omitempty"`
}

type UpdateCommentRequestDTO struct {
	ProjectID    string `json:"projectId"`
	DiscussionID string `json:"discussionId"`
	CommentID    string `json:"commentId"`
	Text         string `json:"text"`
	MarkupType   string `json:"markupType,omitempty"`
}

type RemoveCommentRequestDTO struct {
	ProjectID    string `json:"projectId"`
	DiscussionID string `json:"discussionId"`
	CommentID    string `json:"commentId"`
}

type UpdateTaskListRequestDTO struct {
	ProjectID    string `json:"projectId"`
	DiscussionID string `json:"discussionId"`
	CommentID    string `json:"commentId"`
	ItemIndex    int32  `json:"itemIndex"`
	ItemValue    bool   `json:"itemValue"`
}

type DiscussionsInProjectRequestDTO struct {
	ProjectID string `json:"projectId"`
	Query     string `json:"query"`
	Limit     int32  `json:"limit"`
	Skip      int32  `json:"skip,omitempty"`
}

type DiscussionsInProjectDTO struct {
	Discussions []DiscussionInFileDTO `json:"discussions"`
	HasMore     *bool                 `json:"hasMore,omitempty"`
	TotalCount  *int32                `json:"totalCount,omitempty"`
}

type ToggleReactionRequestDTO struct {
	Target     ReactionTargetDTO `json:"target"`
	ReactionID string            `json:"reactionId"`
	DoAdd      bool              `json:"doAdd"`
}

type RevisionListDTO struct {
	ProjectID  string   `json:"projectId"`
	RevisionID []string `json:"revisionId"`
}

type RevisionReviewInfoDTO struct {
	ReviewInfo *ShortReviewInfoDTO `json:"reviewInfo,omitempty"`
}

type RevisionReviewInfoListDTO struct {
	ReviewInfo []RevisionReviewInfoDTO `json:"reviewInfo"`
}

type ReviewSuggestDTO struct {
	ReviewInfo *ShortReviewInfoDTO `json:"reviewInfo,omitempty"`
	Score      *float32            `json:"score,omitempty"`
}

type SuggestedReviewListDTO struct {
	Suggest []ReviewSuggestDTO `json:"suggest"`
}

type UncommittedRevisionDTO struct {
	ProjectID        string   `json:"projectId"`
	ParentRevisionID string   `json:"parentRevisionId"`
	Files            []string `json:"files"`
}

type RevisionBuildStatusKeyDTO struct {
	Status      BuildStatusEnum `json:"status"`
	Name        string          `json:"name,omitempty"`
	URL         string          `json:"url,omitempty"`
	Description string          `json:"description,omitempty"`
}

type RevisionBuildStatusDTO struct {
	ProjectID  string                      `json:"projectId"`
	RevisionID string                      `json:"revisionId"`
	Keys       []RevisionBuildStatusKeyDTO `json:"keys"`
}

type RevisionBuildStatusListDTO struct {
	BuildStatus []RevisionBuildStatusDTO `json:"buildStatus"`
}

type RevisionDiscussionCountersRequestDTO struct {
	Revisions      RevisionListDTO `json:"revisions"`
	FileNameFilter string          `json:"fileNameFilter,omitempty"`
}

type RevisionDiscussionCountersListDTO struct {
	Counter []SimpleDiscussionCounterDTO `json:"counter"`
}

type FeedRequestDTO struct {
	Limit          int32               `json:"limit"`
	Type           FeedTypeEnum        `json:"type"`
	ProjectID      string              `json:"projectId,omitempty"`
	ReviewID       string              `json:"reviewId,omitempty"`
	SearchQuery    string              `json:"searchQuery,omitempty"`
	ReviewFeedSort *ReviewFeedSortEnum `json:"reviewFeedSort,omitempty"`
}

type DiscussionInFeedDTO struct {
	ProjectID            string              `json:"projectId"`
	DiscussionID         string              `json:"discussionId"`
	Anchor               AnchorDTO           `json:"anchor"`
	Comments             []CommentDTO        `json:"comments"`
	Review               *ShortReviewInfoDTO `json:"review,omitempty"`
	Labels               []LabelDTO          `json:"labels"`
	Read                 *ReadEnum           `json:"read,omitempty"`
	IsStarred            *bool               `json:"isStarred,omitempty"`
	FirstUnreadCommentID string              `json:"firstUnreadCommentId,omitempty"`
	Issue                string              `json:"issue,omitempty"`
	IsResolved           *bool               `json:"isResolved,omitempty"`
	ResolvedBy           string              `json:"resolvedBy,omitempty"`
}

type DiscussionInFileWithFileDTO struct {
	RevisionID       string              `json:"revisionId,omitempty"`
	FileName         string              `json:"fileName"`
	DiscussionInFile DiscussionInFileDTO `json:"discussionInFile"`
}

type DiscussionsInFilesDTO struct {
	Discussions []DiscussionInFileWithFileDTO `json:"discussions"`
}

type ParticipantStateChangedDTO struct {
	Participant string               `json:"participant"`
	OldState    ParticipantStateEnum `json:"oldState"`
	NewState    ParticipantStateEnum `json:"newState"`
}

type ReviewStateChangedDTO struct {
	OldState ReviewStateEnum `json:"oldState"`
	NewState ReviewStateEnum `json:"newState"`
}

type FeedItemDTO struct {
	FeedItemID                   string                      `json:"feedItemId"`
	ProjectID                    string                      `json:"projectId"`
	Discussion                   *DiscussionInFeedDTO        `json:"discussion,omitempty"`
	AddedRevisions               []RevisionInfoDTO           `json:"addedRevisions"`
	RemovedRevisions             []RevisionInfoDTO           `json:"removedRevisions"`
	NewParticipantInReview       *ParticipantInReviewDTO     `json:"newParticipantInReview,omitempty"`
	RemovedParticipantFromReview string                      `json:"removedParticipantFromReview,omitempty"`
	ParticipantStateChanged      *ParticipantStateChangedDTO `json:"participantStateChanged,omitempty"`
	CreatedReview                *ShortReviewInfoDTO         `json:"createdReview,omitempty"`
	ModifiedReview               *ShortReviewInfoDTO         `json:"modifiedReview,omitempty"`
	RemovedReview                *ReviewIdDTO                `json:"removedReview,omitempty"`
	ReviewStateChanged           *ReviewStateChangedDTO      `json:"reviewStateChanged,omitempty"`
	BranchTrackingStopped        string                      `json:"branchTrackingStopped,omitempty"`
	UpdatedDeadline              *int64                      `json:"updatedDeadline,omitempty"`
	PullRequest                  string                      `json:"pullRequest,omitempty"`
	Date                         int64                       `json:"date"`
	ActorID                      string                      `json:"actorId"`
	SquashedToRevision           *RevisionInfoDTO            `json:"squashedToRevision,omitempty"`
	RemovedDiscussionID          string                      `json:"removedDiscussionId,omitempty"`
}

type FeedDTO struct {
	FeedItems []FeedItemDTO `json:"feedItems"`
	HasMore   bool          `json:"hasMore"`
}

type UpdateDiscussionFlagRequestDTO struct {
	ProjectID    string `json:"projectId"`
	DiscussionID string `json:"discussionId"`
	IsFlagged    bool   `json:"isFlagged"`
}

type RevisionsDiffRequestDTO struct {
	ProjectID      string `json:"projectId"`
	BaseRevision   string `json:"baseRevision"`
	SecondRevision string `json:"secondRevision"`
}

type ReviewSummaryChangesResponseDTO struct {
	Diff            *RevisionsDiffDTO    `json:"diff,omitempty"`
	Annotation      string               `json:"annotation,omitempty"`
	IgnoredFiles    []string             `json:"ignoredFiles"`
	FileDiffSummary []FileDiffSummaryDTO `json:"fileDiffSummary"`
}

type FileDiffSummaryDTO struct {
	File         FileInRevisionDTO `json:"file"`
	AddedLines   int32             `json:"addedLines"`
	RemovedLines int32             `json:"removedLines"`
}

type RevisionsDiffDTO struct {
	Diff []RevisionDiffItemDTO `json:"diff"`
}

type RevisionDiffItemDTO struct {
	ProjectID    string             `json:"projectId"`
	DiffType     DiffTypeEnum       `json:"diffType"`
	NewFile      FileInRevisionDTO  `json:"newFile"`
	OldFile      *FileInRevisionDTO `json:"oldFile,omitempty"`
	FileIcon     string             `json:"fileIcon,omitempty"`
	IsRead       bool               `json:"isRead"`
	ConflictType *ConflictTypeEnum  `json:"conflictType,omitempty"`
}

type RevisionInBranchDTO struct {
	Revision RevisionInfoDTO `json:"revision"`
	Branch   string          `json:"branch,omitempty"`
}

type FileMetaResponseDTO struct {
	IsUpToDate                   bool                  `json:"isUpToDate"`
	IsDeleted                    bool                  `json:"isDeleted"`
	IsMerged                     bool                  `json:"isMerged"`
	LastModifiedRevision         RevisionInfoDTO       `json:"lastModifiedRevision"`
	IsLastModifiedRevisionMerged bool                  `json:"isLastModifiedRevisionMerged"`
	DeletedInRevision            *RevisionInfoDTO      `json:"deletedInRevision,omitempty"`
	DeletedInBranch              string                `json:"deletedInBranch,omitempty"`
	ModifiedInParallelBranches   []RevisionInBranchDTO `json:"modifiedInParallelBranches"`
	DefaultBranch                string                `json:"defaultBranch,omitempty"`
}

type FileHistoryRequestDTO struct {
	File  FileInRevisionDTO `json:"file"`
	Limit int32             `json:"limit"`
	Skip  int32             `json:"skip,omitempty"`
}

type FileHistoryResponseDTO struct {
	History []FileHistoryItemDTO  `json:"history"`
	Graph   *RevisionListGraphDTO `json:"graph,omitempty"`
	HasMore bool                  `json:"hasMore"`
}

type FileHistoryItemDTO struct {
	DiffType DiffTypeEnum    `json:"diffType"`
	Revision RevisionInfoDTO `json:"revision"`
	FileName string          `json:"fileName"`
}

type FileAnnotationResponseDTO struct {
	Retrospective []FileAnnotationSectionDTO `json:"retrospective"`
	Perspective   []FileAnnotationSectionDTO `json:"perspective"`
}

type FileAnnotationSectionDTO struct {
	StartLine             int32                     `json:"startLine"`
	LineCount             int32                     `json:"lineCount"`
	Revision              RevisionInfoDTO           `json:"revision"`
	FilePath              string                    `json:"filePath"`
	PriorChangeAnnotation *FileAnnotationSectionDTO `json:"priorChangeAnnotation,omitempty"`
}

type FileDiffRequestDTO struct {
	LeftFile                *FileInRevisionDTO `json:"leftFile,omitempty"`
	RightFile               *FileInRevisionDTO `json:"rightFile,omitempty"`
	IgnoreWhitespace        bool               `json:"ignoreWhitespace"`
	IsLeftFileDefinedAsNull *bool              `json:"isLeftFileDefinedAsNull,omitempty"`
	ContextLines            int32              `json:"contextLines,omitempty"`
}

type FileDiffResponseDTO struct {
	LeftFile  *FileInRevisionDTO    `json:"leftFile,omitempty"`
	RightFile *FileInRevisionDTO    `json:"rightFile,omitempty"`
	Fragments []FileDiffFragmentDTO `json:"fragments"`
}

type FileDiffFragmentDTO struct {
	LeftFileStartLine  int32      `json:"leftFileStartLine"`
	LeftFileLineCount  int32      `json:"leftFileLineCount"`
	RightFileStartLine int32      `json:"rightFileStartLine"`
	RightFileLineCount int32      `json:"rightFileLineCount"`
	IsUnchanged        bool       `json:"isUnchanged"`
	AddedRanges        []RangeDTO `json:"addedRanges"`
	DeletedRanges      []RangeDTO `json:"deletedRanges"`
}

type FileMergeInlineDiffRequestDTO struct {
	FileID           FileInRevisionDTO `json:"fileId"`
	SourceRevisionID string            `json:"sourceRevisionId"`
	BaseBranch       string            `json:"baseBranch"`
	DiffType         DiffTypeEnum      `json:"diffType"`
	IgnoreWhitespace bool              `json:"ignoreWhitespace"`
	ContextLines     int32             `json:"contextLines,omitempty"`
}

type RevisionChangesRequestDTO struct {
	Revision            RevisionInProjectDTO `json:"revision"`
	CompareToRevisionID string               `json:"compareToRevisionId,omitempty"`
	Limit               int32                `json:"limit"`
	Skip                int32                `json:"skip,omitempty"`
}

type FileInlineDiffResponseDTO struct {
	IsIdentical         bool                       `json:"isIdentical"`
	Text                string                     `json:"text"`
	OldFile             *FileInRevisionDTO         `json:"oldFile,omitempty"`
	NewFile             *FileInRevisionDTO         `json:"newFile,omitempty"`
	ContentType         FileContentTypeDTO         `json:"contentType"`
	AddedLines          []int32                    `json:"addedLines"`
	RemovedLines        []int32                    `json:"removedLines"`
	ModifiedLines       []int32                    `json:"modifiedLines"`
	CollapsedLines      []RangeDTO                 `json:"collapsedLines"`
	AddedRanges         []RangeDTO                 `json:"addedRanges"`
	RemovedRanges       []RangeDTO                 `json:"removedRanges"`
	IsSyntaxSupported   *bool                      `json:"isSyntaxSupported,omitempty"`
	SyntaxMarkup        []TextMarkupDTO            `json:"syntaxMarkup"`
	DiffToOldDocument   []RangeMappingDTO          `json:"diffToOldDocument"`
	DiffToNewDocument   []RangeMappingDTO          `json:"diffToNewDocument"`
	OldLineNumbers      []int32                    `json:"oldLineNumbers"`
	NewLineNumbers      []int32                    `json:"newLineNumbers"`
	Annotation          []FileAnnotationSectionDTO `json:"annotation"`
	HasUnrelatedChanges *bool                      `json:"hasUnrelatedChanges,omitempty"`
	ConflictType        *ConflictTypeEnum          `json:"conflictType,omitempty"`
}

type FileInReviewDiffRequestDTO struct {
	File                 FileInReviewDTO  `json:"file"`
	IgnoreWhitespace     bool             `json:"ignoreWhitespace"`
	Revisions            *RevisionsSetDTO `json:"revisions,omitempty"`
	ShowUnrelatedChanges *bool            `json:"showUnrelatedChanges,omitempty"`
	ContextLines         int32            `json:"contextLines,omitempty"`
}

type FileInReviewReadStatusRequestDTO struct {
	ReviewID     ReviewIdDTO     `json:"reviewId"`
	File         string          `json:"file"`
	Revisions    RevisionsSetDTO `json:"revisions"`
	MarkAsUnread *bool           `json:"markAsUnread,omitempty"`
}

type ReviewSummaryDiscussionsRequestDTO struct {
	ReviewID  ReviewIdDTO      `json:"reviewId"`
	Revisions *RevisionsSetDTO `json:"revisions,omitempty"`
}

type FileFragmentAuthorsRequestDTO struct {
	File      FileInRevisionDTO `json:"file"`
	StartLine int32             `json:"startLine"`
	EndLine   int32             `json:"endLine"`
}

type RangeMappingDTO struct {
	From RangeDTO `json:"from"`
	To   RangeDTO `json:"to"`
}

type ProjectConfigurationResponseDTO struct {
	SupportedVcs []string `json:"supportedVcs"`
}

type VcsHostingRequestDTO struct {
	ProjectID string `json:"projectId,omitempty"`
}

type VcsHostingResponseDTO struct {
	Service []VcsHostingServiceDTO `json:"service"`
}

type VcsHostingServiceDTO struct {
	ServiceID           string   `json:"serviceId"`
	ServiceURL          string   `json:"serviceUrl"`
	RepoIDs             []string `json:"repoIds"`
	CanCreateRepository bool     `json:"canCreateRepository"`
}

type StacktraceDTO struct {
	ProjectID  string `json:"projectId"`
	Stacktrace string `json:"stacktrace"`
	RevisionID string `json:"revisionId,omitempty"`
}

type StacktracePositionDTO struct {
	LineNumber        int32     `json:"lineNumber"`
	FullPath          string    `json:"fullPath"`
	VcsCommitID       string    `json:"vcsCommitId,omitempty"`
	InterpolatedRange *RangeDTO `json:"interpolatedRange,omitempty"`
}

type StacktracePositionsDTO struct {
	Positions []StacktracePositionDTO `json:"positions"`
}

type FileOwnershipSummaryDTO struct {
	FilePath string               `json:"filePath"`
	State    OwnershipSummaryEnum `json:"state"`
	UserID   string               `json:"userId,omitempty"`
}

type ReviewOwnershipSummaryDTO struct {
	Files []FileOwnershipSummaryDTO `json:"files"`
}

type ParticipantProgressDTO struct {
	UserID            string `json:"userId"`
	ReadFilesCount    int32  `json:"readFilesCount"`
	LastSeenTimestamp int64  `json:"lastSeenTimestamp"`
}

type ReviewProgressDTO struct {
	Participants []ParticipantProgressDTO `json:"participants"`
	TotalFiles   int32                    `json:"totalFiles"`
}

type RevisionOwnershipSummaryDTO struct {
	RevisionID string `json:"revisionId"`
	UserID     string `json:"userId"`
}

type RevisionsOwnershipSummaryDTO struct {
	RevisionOwner []RevisionOwnershipSummaryDTO `json:"revisionOwner"`
}

type RevisionsExternalInspectionsDiffResponseDTO struct {
	Diff []RevisionExternalInspectionsDiffResponseDTO `json:"diff"`
}

type RevisionExternalInspectionsDiffResponseDTO struct {
	ProjectID  string                                   `json:"projectId"`
	RevisionID string                                   `json:"revisionId"`
	Files      []FileExternalInspectionsDiffResponseDTO `json:"files"`
}

type FileExternalInspectionsDiffResponseDTO struct {
	FileID     string               `json:"fileId"`
	Introduced *InspectionsCountDTO `json:"introduced,omitempty"`
	Fixed      *InspectionsCountDTO `json:"fixed,omitempty"`
}

type InspectionsCountDTO struct {
	Errors       int32 `json:"errors"`
	Warnings     int32 `json:"warnings"`
	WeakWarnings int32 `json:"weakWarnings"`
	Infos        int32 `json:"infos"`
}

type ExternalInspectionDTO struct {
	RevisionID string                 `json:"revisionId"`
	FileName   string                 `json:"fileName"`
	Line       int32                  `json:"line"`
	Severity   InspectionSeverityEnum `json:"severity"`
	Message    string                 `json:"message"`
}

type CompareRequestDTO struct {
	ProjectID    string `json:"projectId"`
	LeftLocator  string `json:"leftLocator"`
	RightLocator string `json:"rightLocator"`
}

type ExampleComparison struct {
	LeftLocator    string `json:"leftLocator"`
	RightLocator   string `json:"rightLocator"`
	LastCommitTime int64  `json:"lastCommitTime"`
}

type BranchingRevisionSuggestion struct {
	RevisionID      string `json:"revisionId"`
	PathToLeftSize  int32  `json:"pathToLeftSize"`
	PathToRightSize int32  `json:"pathToRightSize"`
}

type CompareInfoDTO struct {
	LeftRevisionID  string                       `json:"leftRevisionId"`
	RightRevisionID string                       `json:"rightRevisionId"`
	Heads           []string                     `json:"heads"`
	CommitsCount    int32                        `json:"commitsCount"`
	FilesCount      int32                        `json:"filesCount"`
	Examples        []ExampleComparison          `json:"examples"`
	Suggestion      *BranchingRevisionSuggestion `json:"suggestion,omitempty"`
}

type CompareGraphDTO struct {
	Revisions              []RevisionInfoDTO    `json:"revisions"`
	Graph                  RevisionListGraphDTO `json:"graph"`
	PathToBaseLength       int32                `json:"pathToBaseLength"`
	FullPathToBaseLength   int32                `json:"fullPathToBaseLength"`
	PathToSecondLength     int32                `json:"pathToSecondLength"`
	FullPathToSecondLength int32                `json:"fullPathToSecondLength"`
}

type CanCreateBranchReviewDTO struct {
	IsAllowed bool   `json:"isAllowed"`
	Message   string `json:"message,omitempty"`
}

type BranchStatsDTO struct {
	ParentBranch  string `json:"parentBranch,omitempty"`
	CommitsAhead  int32  `json:"commitsAhead"`
	CommitsBehind int32  `json:"commitsBehind"`
}

type BranchMergeInfoDTO struct {
	MergeFiles []string `json:"mergeFiles"`
}

type BranchInfoDTO struct {
	CommitsCount      int32                    `json:"commitsCount"`
	FilesCount        int32                    `json:"filesCount"`
	BranchingRevision string                   `json:"branchingRevision"`
	HeadRevision      RevisionInfoDTO          `json:"headRevision"`
	ReviewInfo        *ReviewDescriptorDTO     `json:"reviewInfo,omitempty"`
	CanCreateReview   CanCreateBranchReviewDTO `json:"canCreateReview"`
	Stats             BranchStatsDTO           `json:"stats"`
	MergeInfo         BranchMergeInfoDTO       `json:"mergeInfo"`
	IsPullRequest     bool                     `json:"isPullRequest"`
}

type BranchRequestDTO struct {
	ProjectID string `json:"projectId"`
	Branch    string `json:"branch"`
}

type BranchGraphDTO struct {
	Revisions []RevisionInfoDTO    `json:"revisions"`
	Graph     RevisionListGraphDTO `json:"graph"`
}

type RevisionBranchesResponseDTO struct {
	BranchName []string `json:"branchName"`
}

type FindCommitsRequestDTO struct {
	Commits        []FindCommitsRequestPatternDTO `json:"commits"`
	RequestChanges *bool                          `json:"requestChanges,omitempty"`
	Limit          int32                          `json:"limit,omitempty"`
}

type FindCommitsResponseDTO struct {
	Commits []FindCommitsResponseCommitsDTO `json:"commits"`
}

type FindCommitsRequestPatternDTO struct {
	RevisionID      string `json:"revisionId,omitempty"`
	ProjectID       string `json:"projectId,omitempty"`
	MessageFragment string `json:"messageFragment,omitempty"`
	Author          string `json:"author,omitempty"`
	CommitTime      int64  `json:"commitTime,omitempty"`
}

type FindCommitsResponseCommitsDTO struct {
	Commits []FindCommitsResponseCommitDTO `json:"commits"`
}

type FindCommitsResponseCommitDTO struct {
	ProjectID   string            `json:"projectId"`
	ProjectName string            `json:"projectName"`
	Revision    RevisionInfoDTO   `json:"revision"`
	Changes     *RevisionsDiffDTO `json:"changes,omitempty"`
}

type BranchesRequestDTO struct {
	ProjectID string `json:"projectId"`
	Query     string `json:"query"`
	Limit     int32  `json:"limit"`
	SortBy    string `json:"sortBy,omitempty"`
}

type BranchDTO struct {
	ProjectID    string          `json:"projectId"`
	Name         string          `json:"name"`
	LastRevision RevisionInfoDTO `json:"lastRevision"`
	IsDefault    bool            `json:"isDefault"`
	Stats        BranchStatsDTO  `json:"stats"`
	IsHosted     bool            `json:"isHosted"`
	ReviewID     *ReviewIdDTO    `json:"reviewId,omitempty"`
}

type BranchListDTO struct {
	Branch                []BranchDTO `json:"branch"`
	HasMore               bool        `json:"hasMore"`
	TotalBranches         int32       `json:"totalBranches"`
	DefaultBranch         string      `json:"defaultBranch,omitempty"`
	IsDefaultBranchExists *bool       `json:"isDefaultBranchExists,omitempty"`
}

type RoleDetailsDTO struct {
	ID          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type AllRolesResponseDTO struct {
	RoleDetails []RoleDetailsDTO `json:"roleDetails"`
}

type RoleDTO struct {
	Key      string `json:"key"`
	IsGlobal bool   `json:"isGlobal"`
	IsGroup  bool   `json:"isGroup"`
}

type UserRolesDTO struct {
	UserID string    `json:"userId"`
	Roles  []RoleDTO `json:"roles"`
}

type UsersRolesRequestDTO struct {
	ProjectID string `json:"projectId"`
	Offset    int32  `json:"offset"`
	PageSize  int32  `json:"pageSize"`
	Query     string `json:"query,omitempty"`
}

type UsersRolesResponseDTO struct {
	UserRoles []UserRolesDTO `json:"userRoles"`
	HasMore   bool           `json:"hasMore"`
}

type DeleteRoleRequestDTO struct {
	ProjectID string `json:"projectId"`
	UserID    string `json:"userId"`
	RoleKey   string `json:"roleKey"`
}

type AddRoleRequestDTO struct {
	ProjectID string `json:"projectId"`
	UserID    string `json:"userId"`
	RoleKey   string `json:"roleKey"`
}

type InviteUserRequestDTO struct {
	ProjectID string `json:"projectId"`
	Email     string `json:"email"`
}

type InviteUserResponseDTO struct {
	User      *UserRolesDTO `json:"user,omitempty"`
	IsInvited bool          `json:"isInvited"`
}

type FindPullRequestDTO struct {
	ProjectID string `json:"projectId"`
	Branch    string `json:"branch"`
}

type VcsRepoListDTO struct {
	Repo []VcsRepoDTO `json:"repo"`
}

type VcsRepoDTO struct {
	ID  string   `json:"id"`
	URL []string `json:"url"`
}

// IssueTrackerProviderSettingsDTO is a placeholder for now

type IssueTrackerProviderSettingsDTO struct{}

// IssueTrackerProjectDetailsDTO is a placeholder for now

type IssueTrackerProjectDetailsDTO struct{}

// InspectionSeverityEnum is a placeholder for now

type InspectionSeverityEnum int

type TimeUnitEnum int

const (
	HOUR    TimeUnitEnum = 1
	DAY     TimeUnitEnum = 2
	WEEK    TimeUnitEnum = 3
	MONTH   TimeUnitEnum = 4
	QUARTER TimeUnitEnum = 5
	YEAR    TimeUnitEnum = 6
)

type ReviewCoverageStateEnum int

const (
	CLOSED ReviewCoverageStateEnum = 1
	OPEN   ReviewCoverageStateEnum = 2
	ALL    ReviewCoverageStateEnum = 3
)

type AnalyzerStats struct {
	MinCommitTime        int64 `json:"minCommitTime"`
	MaxCommitTime        int64 `json:"maxCommitTime"`
	TotalCommits         int32 `json:"totalCommits"`
	MinIndexedCommitTime int64 `json:"minIndexedCommitTime"`
	MaxIndexedCommitTime int64 `json:"maxIndexedCommitTime"`
	TotalIndexedCommits  int32 `json:"totalIndexedCommits"`
	IsProjectModelKnown  bool  `json:"isProjectModelKnown"`
}

type TimeValue struct {
	Time  int64 `json:"time"`
	Value int32 `json:"value"`
}

type CommitterUserInfo struct {
	Committer string `json:"committer"`
	UserID    string `json:"userId"`
}

type ProjectActivityRequestDTO struct {
	ProjectID     string `json:"projectId"`
	Module        string `json:"module,omitempty"`
	ReferenceTime int64  `json:"referenceTime"`
}

type ProjectActivityDTO struct {
	Items   []TimeValue   `json:"items"`
	Modules []string      `json:"modules"`
	Stats   AnalyzerStats `json:"stats"`
}

type ContributorsDistributionRequestDTO struct {
	ProjectID  string  `json:"projectId"`
	Module     string  `json:"module,omitempty"`
	FromTime   int64   `json:"fromTime"`
	ToTime     int64   `json:"toTime"`
	TimePoints []int64 `json:"timePoints"`
}

type CommitterTimeValue struct {
	Committer string      `json:"committer"`
	Items     []TimeValue `json:"items"`
}

type ContributorsDistributionDTO struct {
	Items      []CommitterTimeValue `json:"items"`
	TimePoints []int64              `json:"timePoints"`
	Users      []CommitterUserInfo  `json:"users"`
}

type ResponsibilityDistributionRequestDTO struct {
	ProjectID string `json:"projectId"`
	FromTime  int64  `json:"fromTime"`
	ToTime    int64  `json:"toTime"`
}

type ModuleValue struct {
	Module string `json:"module"`
	Value  int32  `json:"value"`
}

type CommitterModuleValue struct {
	Committer string        `json:"committer"`
	Items     []ModuleValue `json:"items"`
}

type ResponsibilityDistributionDTO struct {
	Items   []CommitterModuleValue `json:"items"`
	Users   []CommitterUserInfo    `json:"users"`
	Modules []string               `json:"modules"`
	Stats   AnalyzerStats          `json:"stats"`
}

type ProjectCommittersDTO struct {
	Users []CommitterUserInfo `json:"users"`
}

type UserActivityRequestDTO struct {
	ProjectID     string       `json:"projectId"`
	Period        TimeUnitEnum `json:"period"`
	ReferenceTime int64        `json:"referenceTime"`
	Committers    []string     `json:"committers"`
}

type UserActivityDTO struct {
	Items []TimeValue   `json:"items"`
	Stats AnalyzerStats `json:"stats"`
}

type ModulesDistributionRequestDTO struct {
	ProjectID  string   `json:"projectId"`
	FromTime   int64    `json:"fromTime"`
	ToTime     int64    `json:"toTime"`
	Committers []string `json:"committers"`
	TimePoints []int64  `json:"timePoints"`
}

type ModuleTimeValue struct {
	Module string      `json:"module"`
	Items  []TimeValue `json:"items"`
}

type ModulesDistributionDTO struct {
	Items      []ModuleTimeValue `json:"items"`
	TimePoints []int64           `json:"timePoints"`
	Modules    []string          `json:"modules"`
}

type CommitsSummaryRequestDTO struct {
	ProjectID  string   `json:"projectId"`
	FromTime   int64    `json:"fromTime"`
	ToTime     int64    `json:"toTime"`
	Committers []string `json:"committers"`
}

type CommitsSummaryDTO struct {
	TotalCommits   int32 `json:"totalCommits"`
	OffsiteCommits int32 `json:"offsiteCommits"`
	ModulesTouched int32 `json:"modulesTouched"`
}

type CommitsDetailsRequestDTO struct {
	ProjectID  string   `json:"projectId"`
	FromTime   int64    `json:"fromTime"`
	ToTime     int64    `json:"toTime"`
	Committers []string `json:"committers"`
	Module     string   `json:"module,omitempty"`
}

type CommitInfo struct {
	RevisionID      string `json:"revisionId"`
	Description     string `json:"description"`
	Time            int64  `json:"time"`
	ShortRevisionID string `json:"shortRevisionId"`
}

type CommitsDetailsDTO struct {
	Commits []CommitInfo `json:"commits"`
}

type UserValue struct {
	UserID string `json:"userId"`
	Value  int32  `json:"value"`
}

type LabelStats struct {
	Label         LabelDTO    `json:"label"`
	CountsPerUser []UserValue `json:"countsPerUser"`
}

type CommentsStatisticsDTO struct {
	ReviewComments  int32        `json:"reviewComments"`
	TotalComments   int32        `json:"totalComments"`
	TotalCounts     []UserValue  `json:"totalCounts"`
	UnlabeledCounts []UserValue  `json:"unlabeledCounts"`
	StatsPerLabel   []LabelStats `json:"statsPerLabel"`
}

type ReviewStatisticsRequestDTO struct {
	ProjectID string `json:"projectId"`
	FromTime  int64  `json:"fromTime,omitempty"`
	ToTime    int64  `json:"toTime,omitempty"`
}

type ReviewStatisticsDTO struct {
	OpenReviews                     int32                 `json:"openReviews"`
	ClosedReviews                   int32                 `json:"closedReviews"`
	AllRevisions                    int32                 `json:"allRevisions"`
	RevisionsCoveredByOpenReviews   int32                 `json:"revisionsCoveredByOpenReviews"`
	RevisionsCoveredByClosedReviews int32                 `json:"revisionsCoveredByClosedReviews"`
	AuthorStatsByReviews            []UserValue           `json:"authorStatsByReviews"`
	ReviewerStatsByReviews          []UserValue           `json:"reviewerStatsByReviews"`
	AuthorStatsByRevisions          []UserValue           `json:"authorStatsByRevisions"`
	ReviewerStatsByRevisions        []UserValue           `json:"reviewerStatsByRevisions"`
	CommentsStats                   CommentsStatisticsDTO `json:"commentsStats"`
}

type ReviewCoverageRequestDTO struct {
	ProjectID     string       `json:"projectId"`
	Period        TimeUnitEnum `json:"period"`
	ReferenceTime int64        `json:"referenceTime"`
}

type ReviewCoverageDTO struct {
	AllRevisions             []TimeValue `json:"allRevisions"`
	RevisionsCoveredByOpen   []TimeValue `json:"revisionsCoveredByOpen"`
	RevisionsCoveredByClosed []TimeValue `json:"revisionsCoveredByClosed"`
}

type FileHistoryChartRequestDTO struct {
	ProjectID   string `json:"projectId"`
	FramesCount int32  `json:"framesCount,omitempty"`
}

type FileHistoryFrameDTO struct {
	Timestamp int64   `json:"timestamp"`
	Deltas    []int32 `json:"deltas"`
}

type FileHistoryChartDTO struct {
	MaxAge     int64                 `json:"maxAge"`
	MaxEdits   int32                 `json:"maxEdits"`
	Paths      []string              `json:"paths"`
	FirstEdits []int64               `json:"firstEdits"`
	Frames     []FileHistoryFrameDTO `json:"frames"`
}

type ProjectTreeMapRequestDTO struct {
	ProjectID  string `json:"projectId"`
	RevisionID string `json:"revisionId"`
}

type ProjectTreeMapEntryDTO struct {
	Path                string `json:"path"`
	Size                int32  `json:"size"`
	ModificationsCount  int32  `json:"modificationsCount"`
	IsBinary            bool   `json:"isBinary"`
	TimeSinceLastChange int64  `json:"timeSinceLastChange"`
}

type ProjectTreeMapDTO struct {
	Entries []ProjectTreeMapEntryDTO `json:"entries"`
}

type ProjectPulseRequestDTO struct {
	ProjectID  string  `json:"projectId"`
	FromTime   int64   `json:"fromTime"`
	ToTime     int64   `json:"toTime"`
	TimePoints []int64 `json:"timePoints"`
}

type PulseResponseDTO struct {
	UserValues []int32 `json:"userValues"`
	AllValues  []int32 `json:"allValues"`
}

type UserPulseRequestDTO struct {
	UserID     string  `json:"userId"`
	FromTime   int64   `json:"fromTime"`
	ToTime     int64   `json:"toTime"`
	TimePoints []int64 `json:"timePoints"`
}

type ReviewTimeStatisticsRequestDTO struct {
	UserID string `json:"userId"`
}

type ReviewTimeStatisticsDTO struct {
	Distribution []int64 `json:"distribution"`
}

type ReviewersGraphRequestDTO struct {
	ProjectID string `json:"projectId"`
	FromTime  int64  `json:"fromTime,omitempty"`
}

type ReviewersGraphLinkDTO struct {
	From               int32 `json:"from"`
	To                 int32 `json:"to"`
	Strength           int64 `json:"strength"`
	TotalReviewsCount  int32 `json:"totalReviewsCount"`
	RecentReviewsCount int32 `json:"recentReviewsCount"`
}

type ReviewersGraphDTO struct {
	UserIDs []string                `json:"userIds"`
	Links   []ReviewersGraphLinkDTO `json:"links"`
}

// From FindUsages.json

type StubIdDTO struct {
	FileID    FileInRevisionDTO `json:"fileId"`
	StubIndex int32             `json:"stubIndex"`
}

type PsiElementIdDTO struct {
	FileID     FileInRevisionDTO `json:"fileId"`
	Range      RangeDTO          `json:"range"`
	LocalIndex int32             `json:"localIndex"`
}

type FilePathToRootDTO struct {
	FileName   string              `json:"fileName"`
	PathToRoot ProjectItemsListDTO `json:"pathToRoot"`
}

type FindUsagesResponseDTO struct {
	FilePathToRoot []FilePathToRootDTO `json:"filePathToRoot"`
	Item           []FindUsagesItemDTO `json:"item"`
}

type FindHierarchyResultDTO struct {
	Ancestors  []NavigationTargetItemDTO `json:"ancestors"`
	Inheritors []NavigationTargetItemDTO `json:"inheritors"`
}

type FindUsagesItemDTO struct {
	NavigationTarget     NavigationTargetItemDTO `json:"navigationTarget"`
	PreviewText          string                  `json:"previewText"`
	StartOffsetInPreview int32                   `json:"startOffsetInPreview"`
	EndOffsetInPreview   int32                   `json:"endOffsetInPreview"`
	LineNumber           int32                   `json:"lineNumber"`
	IsImportant          bool                    `json:"isImportant"`
}

type NavigationTargetItemDTO struct {
	FileID            FileInRevisionDTO     `json:"fileId"`
	StartOffset       int32                 `json:"startOffset"`
	EndOffset         int32                 `json:"endOffset"`
	StubIndex         int32                 `json:"stubIndex"`
	TargetDescription *TargetDescriptionDTO `json:"targetDescription,omitempty"`
}

type TargetDescriptionDTO struct {
	ImageName            string `json:"imageName,omitempty"`
	Image                string `json:"image,omitempty"`
	TargetPresentation   string `json:"targetPresentation,omitempty"`
	LocationPresentation string `json:"locationPresentation,omitempty"`
}

type TextSearchRequestDTO struct {
	ProjectID      string `json:"projectId,omitempty"`
	Query          string `json:"query"`
	FilenameFilter string `json:"filenameFilter,omitempty"`
	IgnoreDeleted  bool   `json:"ignoreDeleted"`
	IgnoreArchived bool   `json:"ignoreArchived"`
}

type TextSearchProjectDetails struct {
	ProjectName string `json:"projectName"`
}

type TextSearchSnippet struct {
	Text            string   `json:"text"`
	StartLineNumber int32    `json:"startLineNumber"`
	RangeInSnippet  RangeDTO `json:"rangeInSnippet"`
	RangeInFile     RangeDTO `json:"rangeInFile"`
}

type TextSearchItem struct {
	File           FileInRevisionDTO        `json:"file"`
	RevisionInfo   RevisionInfoDTO          `json:"revisionInfo"`
	ContentType    FileContentTypeDTO       `json:"contentType"`
	IsDeleted      bool                     `json:"isDeleted"`
	ProjectDetails TextSearchProjectDetails `json:"projectDetails"`
	Snippets       []TextSearchSnippet      `json:"snippets"`
	TotalMatches   int32                    `json:"totalMatches"`
}

type TextSearchResponseDTO struct {
	Items []TextSearchItem `json:"items"`
}

type GotoFileRequestDTO struct {
	ProjectID  string       `json:"projectId,omitempty"`
	RevisionID string       `json:"revisionId,omitempty"`
	ReviewID   *ReviewIdDTO `json:"reviewId,omitempty"`
	Pattern    string       `json:"pattern"`
	Limit      int32        `json:"limit"`
}

type GotoFileItemDTO struct {
	ProjectID    string `json:"projectId"`
	RevisionID   string `json:"revisionId"`
	FileName     string `json:"fileName"`
	IsDeleted    bool   `json:"isDeleted"`
	LastModified int64  `json:"lastModified,omitempty"`
	Branch       string `json:"branch,omitempty"`
}

type GotoFileResponseDTO struct {
	Items   []GotoFileItemDTO `json:"items"`
	HasMore bool              `json:"hasMore"`
}

type FindBranchRequestDTO struct {
	ProjectID string `json:"projectId"`
	Pattern   string `json:"pattern"`
	Limit     int32  `json:"limit"`
}

type FindBranchResponseDTO struct {
	Branches []string `json:"branches"`
	HasMore  bool     `json:"hasMore"`
}

type ElementDocumentationDTO struct {
	DocString string `json:"docString,omitempty"`
}

type UsagesDiffFullRequestDTO struct {
	OldElement PsiElementIdDTO `json:"oldElement"`
	NewElement PsiElementIdDTO `json:"newElement"`
}

type UsagesDiffByAnotherRevisionRequestDTO struct {
	OriginElement     PsiElementIdDTO `json:"originElement"`
	AnotherRevisionID string          `json:"anotherRevisionId"`
	IsAnotherOld      bool            `json:"isAnotherOld"`
}

type UsagesDiffByBoundaryRevisionsRequestDTO struct {
	OriginElement      PsiElementIdDTO `json:"originElement"`
	NewRevisionID      string          `json:"newRevisionId"`
	BoundaryRevisionID string          `json:"boundaryRevisionId"`
}

type UsagesDiffRequestDTO struct {
	FullRequest                *UsagesDiffFullRequestDTO                `json:"fullRequest,omitempty"`
	RequestByAnotherRevision   *UsagesDiffByAnotherRevisionRequestDTO   `json:"requestByAnotherRevision,omitempty"`
	RequestByBoundaryRevisions *UsagesDiffByBoundaryRevisionsRequestDTO `json:"requestByBoundaryRevisions,omitempty"`
}

type SameUsagesDTO struct {
	OldUsage FindUsagesItemDTO `json:"oldUsage"`
	NewUsage FindUsagesItemDTO `json:"newUsage"`
}

type UsagesDiffDTO struct {
	OldFilePathToRoot *FilePathToRootDTO  `json:"oldFilePathToRoot,omitempty"`
	NewFilePathToRoot *FilePathToRootDTO  `json:"newFilePathToRoot,omitempty"`
	RemovedUsage      []FindUsagesItemDTO `json:"removedUsage"`
	SameUsage         []SameUsagesDTO     `json:"sameUsage"`
	AddedUsage        []FindUsagesItemDTO `json:"addedUsage"`
}

type UsagesDiffResponseDTO struct {
	OldRevisionID string          `json:"oldRevisionId"`
	NewRevisionID string          `json:"newRevisionId"`
	Usages        []UsagesDiffDTO `json:"usages"`
}

// IssueFieldValueIdDTO is a field-value pair for creating an issue.
type IssueFieldValueIdDTO struct {
	// FieldID is the field ID (see IssueFieldDTO).
	FieldID string `json:"fieldId"`
	// ValueID is the value ID (see IssueFieldValueDTO).
	ValueID string `json:"valueId"`
}

// CreateIssueFromDiscussionRequestDTO is the request to create an issue from a discussion.
type CreateIssueFromDiscussionRequestDTO struct {
	ProjectID    string                 `json:"projectId"`
	DiscussionID string                 `json:"discussionId"`
	FieldValue   []IssueFieldValueIdDTO `json:"fieldValue"`
}

// IssueTrackerProviderDTO describes an issue tracker provider.
// See /upsource_docs/~api_doc/reference/IssueTrackers.html#messages.IssueTrackerProviderDTO
type IssueTrackerProviderDTO struct {
	// ProviderKey is the issue tracker project key (JIRA, YouTrack).
	ProviderKey string `json:"providerKey"`
	// ScriptText is the text of the issue tracker configuration script (must be a React component).
	ScriptText string `json:"scriptText"`
	// IsAutoSetup indicates whether issue tracker integration can be set up automatically.
	IsAutoSetup bool `json:"isAutoSetup"`
}

// IssueTrackerProvidersListDTO is a list of available issue tracker providers.
// See /upsource_docs/~api_doc/reference/IssueTrackers.html#messages.IssueTrackerProvidersListDTO
type IssueTrackerProvidersListDTO struct {
	// Providers is a list of issue tracker providers.
	Providers []IssueTrackerProviderDTO `json:"providers,omitempty"`
}

// IssueInfoRequestDTO is a request for detailed issue information.
// See /upsource_docs/~api_doc/reference/IssueTrackers.html#messages.IssueInfoRequestDTO
type IssueInfoRequestDTO struct {
	// ProjectID is the project ID in Upsource.
	ProjectID string `json:"projectId"`
	// IssueID is the issue ID.
	IssueID IssueIdDTO `json:"issueId"`
}

// IssueIdDTO describes an issue ID.
// See /upsource_docs/~api_doc/reference/IssueTrackers.html#messages.IssueIdDTO
type IssueIdDTO struct {
	// IssueID is the issue ID.
	IssueID string `json:"issueId"`
	// IssueLink is the issue link.
	IssueLink string `json:"issueLink,omitempty"`
	// IsCreatedFromUpsource indicates whether the issue has been created from Upsource.
	IsCreatedFromUpsource bool `json:"isCreatedFromUpsource,omitempty"`
}

// CreateIssueFromReviewRequestDTO is a request to create an issue from a review.
type CreateIssueFromReviewRequestDTO struct {
	ProjectID  string                 `json:"projectId"`
	ReviewKey  string                 `json:"reviewKey"`
	FieldValue []IssueFieldValueIdDTO `json:"fieldValue"`
}

// IssueInfoDTO contains detailed issue information.
// See /upsource_docs/~api_doc/reference/IssueTrackers.html#messages.IssueInfoDTO
type IssueInfoDTO struct {
	// IssueID is the issue ID.
	IssueID string `json:"issueId"`
	// IssueLink is the issue URL.
	IssueLink string `json:"issueLink,omitempty"`
	// Summary is the issue summary.
	Summary string `json:"summary,omitempty"`
	// IsResolved indicates whether the issue has been resolved.
	IsResolved bool `json:"isResolved,omitempty"`
	// Field is the HTML presentation of issue fields (Issue Type, Priority, State).
	Field []string `json:"field,omitempty"`
}

// InspectionsDiffDTO contains the diff of inspections between two revisions.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.InspectionsDiffDTO
type InspectionsDiffDTO struct {
	// DiffForFile is a list of inspection diffs per file.
	DiffForFile []InspectionsDiffForFileDTO `json:"diffForFile,omitempty"`
}

// InspectionsDiffForFileDTO contains the diff of inspections for a single file.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.InspectionsDiffForFileDTO
type InspectionsDiffForFileDTO struct {
	// Path is a full path to the file starting with a slash (e.g. /directory/file.txt).
	Path string `json:"path"`
	// Introduced is a list of introduced problems.
	Introduced []InspectionCodeMarkupItemDTO `json:"introduced,omitempty"`
	// Fixed is a list of fixed problems.
	Fixed []InspectionCodeMarkupItemDTO `json:"fixed,omitempty"`
	// FixedOriginalRanges is a list of ranges pointing to fixed problems in the original revision.
	FixedOriginalRanges []RangeDTO `json:"fixedOriginalRanges,omitempty"`
}

// InspectionCodeMarkupItemDTO describes a single inspection result.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.InspectionCodeMarkupItemDTO
type InspectionCodeMarkupItemDTO struct {
	// Range is the text range of the inspection.
	Range RangeDTO `json:"range"`
	// TextAttribute describes how the inspected text should be highlighted.
	TextAttribute TextAttributeDTO `json:"textAttribute"`
	// Message is the inspection message.
	Message string `json:"message"`
	// Severity is the severity of the inspection.
	Severity InspectionSeverityEnum `json:"severity"`
	// TextAttributesKey is the name (if any) of this markup key, given by IntelliJ IDEA.
	TextAttributesKey string `json:"textAttributesKey,omitempty"`
}

// TextAttributeDTO describes text attributes for highlighting.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.TextAttributeDTO
type TextAttributeDTO struct {
	// FgColor is the foreground color.
	FgColor string `json:"fgColor,omitempty"`
	// BgColor is the background color.
	BgColor string `json:"bgColor,omitempty"`
	// FontStyle is the font style (italic, bold).
	FontStyle string `json:"fontStyle,omitempty"`
	// EffectStyle is the text effect (underline, wave underline, strikeout, etc.).
	EffectStyle string `json:"effectStyle,omitempty"`
	// EffectColor is the effect color.
	EffectColor string `json:"effectColor,omitempty"`
	// ErrorStripeColor is the error stripe color.
	ErrorStripeColor string `json:"errorStripeColor,omitempty"`
}

// TextMarkupDTO describes a single text markup item.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.TextMarkupDTO
type TextMarkupDTO struct {
	// Range is the text range.
	Range RangeDTO `json:"range"`
	// TextAttribute describes how the text should be highlighted.
	TextAttribute TextAttributeDTO `json:"textAttribute"`
}

// FileContributorsResponseDTO contains the file contributors.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.FileContributorsResponseDTO
type FileContributorsResponseDTO struct {
	// AuthorIds are the user IDs of the contributors.
	AuthorIds []string `json:"authorIds,omitempty"`
	// AuthoredChangesCounts is the number of changes made to the file.
	AuthoredChangesCounts []int32 `json:"authoredChangesCounts,omitempty"`
	// ReviewedChangesCounts is the number of reviews performed on the file.
	ReviewedChangesCounts []int32 `json:"reviewedChangesCounts,omitempty"`
	// MajorContributorUserId is the major contributor to the file, if there is one.
	MajorContributorUserId string `json:"majorContributorUserId,omitempty"`
}

// FileContentResponseDTO contains the contents of a file.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.FileContentResponseDTO
type FileContentResponseDTO struct {
	// ContentType describes the file type.
	ContentType FileContentTypeDTO `json:"contentType"`
	// FileContent contains the file content.
	FileContent *FileContentDTO `json:"fileContent,omitempty"`
}

// FileContentTypeDTO describes the file type.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.FileContentTypeDTO
type FileContentTypeDTO struct {
	// IsText indicates whether the file is presentable as text.
	IsText bool `json:"isText"`
	// IsDirectory indicates whether the file is a directory.
	IsDirectory bool `json:"isDirectory"`
	// IsGenerated indicates whether file contents are generated.
	IsGenerated bool `json:"isGenerated"`
	// CanDownload indicates whether the file can be downloaded.
	CanDownload bool `json:"canDownload"`
	// FileType is the file extension.
	FileType string `json:"fileType"`
}

// FileContentDTO contains the file text and markup.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.FileContentDTO
type FileContentDTO struct {
	// Text is the file text.
	Text string `json:"text"`
	// Foldings are the foldable ranges.
	Foldings []FoldingInfoDTO `json:"foldings,omitempty"`
	// IsSyntaxSupported indicates whether syntax markup is available for this file type.
	IsSyntaxSupported bool `json:"isSyntaxSupported,omitempty"`
	// Syntax is the syntax markup.
	Syntax []TextMarkupDTO `json:"syntax,omitempty"`
}

// FilePsiRequestDTO is a request for PSI (Program Structure Interface) information for a file.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.FilePsiRequestDTO
type FilePsiRequestDTO struct {
	File               FileInRevisionDTO `json:"file"`
	RequestReferences  bool              `json:"requestReferences,omitempty"`
	RequestGutterMarks bool              `json:"requestGutterMarks,omitempty"`
	RequestInspections bool              `json:"requestInspections,omitempty"`
}

// FilePsiResponseDTO is a response with PSI information for a file.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.FilePsiResponseDTO
type FilePsiResponseDTO struct {
	// HasPsi indicates whether the code model is available for a file.
	HasPsi bool `json:"hasPsi,omitempty"`
	// PsiStatusMessage is a PSI status message, e.g. a failure to provide code intelligence or another message regarding the file in general.
	PsiStatusMessage string `json:"psiStatusMessage,omitempty"`
	// ReferenceMarkup contains the reference markup.
	ReferenceMarkup *FileReferenceCodeMarkupDTO `json:"referenceMarkup,omitempty"`
	// TextMarkup contains the text markup.
	TextMarkup *FileTextMarkupDTO `json:"textMarkup,omitempty"`
	// GutterMarks contains the gutter marks.
	GutterMarks *GutterCodeMarkupsDTO `json:"gutterMarks,omitempty"`
	// Inspections contains the code inspections.
	Inspections *CodeInspectionsDTO `json:"inspections,omitempty"`
	// ExternalInspections contains the external inspections.
	ExternalInspections []InspectionCodeMarkupItemDTO `json:"externalInspections,omitempty"`
}

// FileReferenceCodeMarkupDTO contains reference code markup.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.FileReferenceCodeMarkupDTO
type FileReferenceCodeMarkupDTO struct {
	Markup                 []ReferenceCodeMarkupItemDTO `json:"markup,omitempty"`
	NavigationPointsTable  []MarkupNavigationPointDTO   `json:"navigationPointsTable,omitempty"`
	FileNameTable          []FileInRevisionDTO          `json:"fileNameTable,omitempty"`
	LocalDeclarationRanges []LocalDeclarationRangeDTO   `json:"localDeclarationRanges,omitempty"`
}

// LocalDeclarationRangeDTO describes a local declaration range.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.LocalDeclarationRangeDTO
type LocalDeclarationRangeDTO struct {
	TargetID    int32 `json:"targetId"`
	StartOffset int32 `json:"startOffset"`
	EndOffset   int32 `json:"endOffset"`
}

// FileTextMarkupDTO contains text markup.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.FileTextMarkupDTO
type FileTextMarkupDTO struct {
	Markup []TextMarkupDTO `json:"markup,omitempty"`
}

// GutterCodeMarkupsDTO contains gutter code markups.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.GutterCodeMarkupsDTO
type GutterCodeMarkupsDTO struct {
	Items []GutterCodeMarkupItemDTO `json:"items,omitempty"`
}

// CodeInspectionsDTO contains code inspections.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.CodeInspectionsDTO
type CodeInspectionsDTO struct {
	Inspections []InspectionCodeMarkupItemDTO `json:"inspections,omitempty"`
	FileLevel   FileWarningLevelEnum          `json:"fileLevel"`
}

// FileWarningLevelEnum describes PSI inspections level
type FileWarningLevelEnum int

const (
	FileWarningLevelEnumNone    FileWarningLevelEnum = 1
	FileWarningLevelEnumOk      FileWarningLevelEnum = 2
	FileWarningLevelEnumWarning FileWarningLevelEnum = 3
	FileWarningLevelEnumError   FileWarningLevelEnum = 4
)

// GutterMarkEnum describes the properties of a gutter marker
type GutterMarkEnum int

// ReferenceCodeMarkupItemDTO describes a single reference code markup item.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.ReferenceCodeMarkupItemDTO
type ReferenceCodeMarkupItemDTO struct {
	Range           RangeDTO `json:"range"`
	MarkupID        int32    `json:"markupId"`
	TargetID        int32    `json:"targetId"`
	CapabilityFlags int32    `json:"capabilityFlags"`
	LocalIndex      int32    `json:"localIndex"`
	Hash            string   `json:"hash,omitempty"`
}

// MarkupNavigationPointDTO describes a navigation point in the markup.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.MarkupNavigationPointDTO
type MarkupNavigationPointDTO struct {
	TargetID    int32 `json:"targetId"`
	FileID      int32 `json:"fileId"`
	StartOffset int32 `json:"startOffset"`
	EndOffset   int32 `json:"endOffset"`
	StubIndex   int32 `json:"stubIndex"`
}

// GutterCodeMarkupItemDTO describes a single gutter code markup item.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.GutterCodeMarkupItemDTO
type GutterCodeMarkupItemDTO struct {
	Range                 RangeDTO         `json:"range"`
	GutterMark            []GutterMarkEnum `json:"gutterMark,omitempty"`
	RelatedSemanticMarkup int32            `json:"relatedSemanticMarkup"`
}

// FoldingInfoDTO describes a foldable range.
// See /upsource_docs/~api_doc/reference/FileOrDirectoryContent.html#messages.FoldingInfoDTO
type FoldingInfoDTO struct {
	// FromOffset is the start offset of the foldable range.
	FromOffset int32 `json:"fromOffset"`
	// ToOffset is the end offset of the foldable range.
	ToOffset int32 `json:"toOffset"`
	// IsInitialCollapsed indicates whether the range is initially collapsed.
	IsInitialCollapsed bool `json:"isInitialCollapsed"`
	// CollapsedText is the text to be displayed in place of the folded range.
	CollapsedText string `json:"collapsedText"`
}
