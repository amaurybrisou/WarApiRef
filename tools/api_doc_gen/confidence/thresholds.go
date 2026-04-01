package confidence

type Level string

const (
	LevelHigh   Level = "HIGH"
	LevelMedium Level = "MEDIUM"
	LevelLow    Level = "LOW"
)

const (
	ScoreMin       = 0
	ScoreMax       = 100
	ScoreHighMin   = 70
	ScoreMediumMin = 40
)

const (
	WeightAppearsFourOrMoreAddons         = 30
	WeightAppearsTwoToThreeAddons         = 18
	WeightDefaultUIPresence               = 35
	WeightXMLHandlerAttributeUsage        = 30
	WeightGlobalFunctionNoLocalDefinition = 20
	WeightKnownEngineNamespace            = 25
	WeightEventPattern                    = 18
	WeightInitializationFlow              = 10
	WeightXMLAndLua                       = 20
	WeightConsistentRole                  = 15
	WeightConsistentArguments             = 10
	WeightConsistentReturns               = 8
	WeightSlashCommandPattern             = 10
	WeightDocumentationReference          = 25
	WeightCrossSourceReinforcement        = 20

	WeightSingleAddonLocalDefinition = -35
	WeightSingleAddonProjectSpecific = -25
	WeightLocalHelperOnly            = -30
	WeightAddonPrefixPenalty         = -18
	WeightSingleInternalModuleOnly   = -15
	WeightWeakSingleUsage            = -20
	WeightConflictingSignatures      = -15
	WeightConflictingRoles           = -20
	WeightWrapperPenalty             = -12
	WeightNeverOutsideLocalGraph     = -25
	WeightPlaceholderOrComputedName  = -100
)

func ClampScore(raw int) int {
	if raw < ScoreMin {
		return ScoreMin
	}
	if raw > ScoreMax {
		return ScoreMax
	}
	return raw
}

func LevelForScore(score int) Level {
	if score >= ScoreHighMin {
		return LevelHigh
	}
	if score >= ScoreMediumMin {
		return LevelMedium
	}
	return LevelLow
}
