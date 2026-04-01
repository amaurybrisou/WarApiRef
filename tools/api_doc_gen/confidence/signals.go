package confidence

type Signal struct {
	Key      string
	Label    string
	Weight   int
	Positive bool
	Detail   string
}

type Evidence struct {
	SymbolName                  string
	SymbolKind                  string
	AddonsSeenIn                []string
	FilesSeenIn                 []string
	XMLUsageCount               int
	LuaUsageCount               int
	GlobalUsageCount            int
	LocalDefinitionCount        int
	NamespacesDetected          []string
	KnownEngineNamespace        bool
	DefaultUIPresence           bool
	EventBindingPresence        bool
	ExampleLocations            []string
	XMLAttributeUsageCount      int
	DocumentationReferenceCount int
	TOCInitReferenceCount       int
	UsedByXMLAndLua             bool
	ConsistentRole              bool
	ConsistentArguments         bool
	ConsistentReturns           bool
	SlashCommandPresence        bool
	WeakUsageOnly               bool
	ProjectSpecificName         bool
	PlaceholderLikeName         bool
	ConflictingSignatures       bool
	ConflictingRoles            bool
	WrapperLikely               bool
	NeverOutsideLocalGraph      bool
	LocalHelperOnly             bool
	SourceKinds                 []string
}

type Assessment struct {
	SymbolName string
	RawScore   int
	FinalScore int
	Level      Level
	Signals    []Signal
	Evidence   Evidence
	Rationale  string
}

type WeightRow struct {
	Condition string
	Weight    int
	Category  string
}

func WeightRows() []WeightRow {
	return []WeightRow{
		{Condition: "Appears in 4 or more distinct addons", Weight: WeightAppearsFourOrMoreAddons, Category: "positive"},
		{Condition: "Appears in 2 to 3 distinct addons", Weight: WeightAppearsTwoToThreeAddons, Category: "positive"},
		{Condition: "Appears in default interface or base UI namespace", Weight: WeightDefaultUIPresence, Category: "positive"},
		{Condition: "Appears directly in XML handler attributes", Weight: WeightXMLHandlerAttributeUsage, Category: "positive"},
		{Condition: "Appears as a global call with no local definition", Weight: WeightGlobalFunctionNoLocalDefinition, Category: "positive"},
		{Condition: "Matches a known engine namespace", Weight: WeightKnownEngineNamespace, Category: "positive"},
		{Condition: "Used in event registration or dispatch patterns", Weight: WeightEventPattern, Category: "positive"},
		{Condition: "Referenced from initialization flow evidence", Weight: WeightInitializationFlow, Category: "positive"},
		{Condition: "Observed from both XML and Lua paths", Weight: WeightXMLAndLua, Category: "positive"},
		{Condition: "Consistent role across multiple addons", Weight: WeightConsistentRole, Category: "positive"},
		{Condition: "Argument pattern is consistent across usages", Weight: WeightConsistentArguments, Category: "positive"},
		{Condition: "Return usage is consistent across usages", Weight: WeightConsistentReturns, Category: "positive"},
		{Condition: "Observed in slash command registration patterns", Weight: WeightSlashCommandPattern, Category: "positive"},
		{Condition: "Referenced by generated docs or reference files", Weight: WeightDocumentationReference, Category: "positive"},
		{Condition: "Reinforced across multiple generated source types", Weight: WeightCrossSourceReinforcement, Category: "positive"},
		{Condition: "Defined locally in only one addon", Weight: WeightSingleAddonLocalDefinition, Category: "negative"},
		{Condition: "Only appears in one addon with project-specific naming", Weight: WeightSingleAddonProjectSpecific, Category: "negative"},
		{Condition: "Only appears as a local helper", Weight: WeightLocalHelperOnly, Category: "negative"},
		{Condition: "Contains addon-specific prefix or suffix", Weight: WeightAddonPrefixPenalty, Category: "negative"},
		{Condition: "Only referenced by one internal module with no XML or global usage", Weight: WeightSingleInternalModuleOnly, Category: "negative"},
		{Condition: "Inferred from one weak usage site", Weight: WeightWeakSingleUsage, Category: "negative"},
		{Condition: "Conflicting signatures across usages", Weight: WeightConflictingSignatures, Category: "negative"},
		{Condition: "Conflicting semantic roles across addons", Weight: WeightConflictingRoles, Category: "negative"},
		{Condition: "Likely wrapper around another platform API", Weight: WeightWrapperPenalty, Category: "negative"},
		{Condition: "Never appears outside addon-local flow", Weight: WeightNeverOutsideLocalGraph, Category: "negative"},
		{Condition: "Looks like a placeholder or computed expression", Weight: WeightPlaceholderOrComputedName, Category: "negative"},
	}
}
