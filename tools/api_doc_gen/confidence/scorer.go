package confidence

func Score(evidence Evidence) Assessment {
	raw := 0
	signals := []Signal{}
	add := func(key string, label string, weight int, positive bool, detail string) {
		raw += weight
		signals = append(signals, Signal{Key: key, Label: label, Weight: weight, Positive: positive, Detail: detail})
	}

	addonCount := len(evidence.AddonsSeenIn)
	if addonCount >= 4 {
		add("addon_spread_4_plus", "Seen in 4 or more addons", WeightAppearsFourOrMoreAddons, true, "Cross-addon spread is strong.")
	} else if addonCount >= 2 {
		add("addon_spread_2_3", "Seen in 2 to 3 addons", WeightAppearsTwoToThreeAddons, true, "Cross-addon spread is present but limited.")
	}
	if evidence.DefaultUIPresence {
		add("default_ui_presence", "Matches default UI or extracted base UI surface", WeightDefaultUIPresence, true, "Symbol aligns with known default-interface namespaces.")
	}
	if evidence.XMLAttributeUsageCount > 0 {
		add("xml_handler_attribute", "Used directly in XML handler attributes", WeightXMLHandlerAttributeUsage, true, "XML exposure suggests an engine-level contract.")
	}
	if evidence.GlobalUsageCount > 0 && evidence.LocalDefinitionCount == 0 {
		add("global_call_no_local_def", "Called globally with no local definition", WeightGlobalFunctionNoLocalDefinition, true, "No addon-local definition was observed in the generated corpus.")
	}
	if evidence.KnownEngineNamespace {
		add("known_engine_namespace", "Matches a known engine namespace", WeightKnownEngineNamespace, true, "Namespace shape matches WAR engine APIs.")
	}
	if evidence.EventBindingPresence {
		add("event_binding", "Used in event registration or dispatch", WeightEventPattern, true, "Observed in event-driven engine hooks.")
	}
	if evidence.TOCInitReferenceCount > 0 {
		add("init_flow", "Referenced from initialization flow", WeightInitializationFlow, true, "Lifecycle reconstruction references this symbol.")
	}
	if evidence.UsedByXMLAndLua {
		add("xml_and_lua", "Observed in both XML and Lua paths", WeightXMLAndLua, true, "Cross-source linkage reinforces platform-level usage.")
	}
	if evidence.ConsistentRole && addonCount >= 2 {
		add("consistent_role", "Role is consistent across addons", WeightConsistentRole, true, "The same symbol serves the same kind of job across addons.")
	}
	if evidence.ConsistentArguments {
		add("consistent_arguments", "Argument pattern is consistent", WeightConsistentArguments, true, "Observed argument positions remain stable.")
	}
	if evidence.ConsistentReturns {
		add("consistent_returns", "Return usage is consistent", WeightConsistentReturns, true, "Observed as a stable query-style API.")
	}
	if evidence.SlashCommandPresence {
		add("slash_pattern", "Appears in slash command registration patterns", WeightSlashCommandPattern, true, "Observed in shared command registration flows.")
	}
	if evidence.DocumentationReferenceCount > 0 {
		add("documentation_reference", "Referenced by generated docs or reference files", WeightDocumentationReference, true, "The symbol is reinforced outside a single call page.")
	}
	if len(evidence.SourceKinds) >= 3 {
		add("cross_source_reinforcement", "Reinforced across multiple generated source types", WeightCrossSourceReinforcement, true, "Evidence comes from several independent addon-api source types.")
	}

	if evidence.LocalDefinitionCount > 0 && addonCount <= 1 {
		add("single_addon_local_definition", "Defined locally in one addon", WeightSingleAddonLocalDefinition, false, "Observed local definition count suggests addon-local ownership.")
	}
	if addonCount <= 1 && evidence.ProjectSpecificName {
		add("single_addon_project_specific", "Only appears once with project-specific naming", WeightSingleAddonProjectSpecific, false, "Name shape strongly resembles addon-local code.")
	}
	if evidence.LocalHelperOnly {
		add("local_helper_only", "Only appears as a local helper", WeightLocalHelperOnly, false, "No shared or structural evidence was found beyond local helper usage.")
	}
	if evidence.ProjectSpecificName && !evidence.KnownEngineNamespace {
		add("addon_prefix_penalty", "Contains addon-specific naming", WeightAddonPrefixPenalty, false, "Name contains addon-local root tokens.")
	}
	if addonCount <= 1 && evidence.XMLUsageCount == 0 && evidence.GlobalUsageCount == 0 && evidence.LuaUsageCount <= 1 {
		add("single_internal_usage", "Only referenced by one internal path", WeightSingleInternalModuleOnly, false, "Structural evidence outside a single internal path is missing.")
	}
	if evidence.WeakUsageOnly {
		add("weak_single_usage", "Only one weak usage site", WeightWeakSingleUsage, false, "Evidence is too shallow to trust as platform API.")
	}
	if evidence.ConflictingSignatures {
		add("conflicting_signatures", "Conflicting signatures across usages", WeightConflictingSignatures, false, "Observed arity or argument shape conflicts across usages.")
	}
	if evidence.ConflictingRoles {
		add("conflicting_roles", "Conflicting semantic roles", WeightConflictingRoles, false, "The same name appears to play incompatible roles.")
	}
	if evidence.WrapperLikely {
		add("wrapper_penalty", "Likely wrapper around another API", WeightWrapperPenalty, false, "The symbol looks like addon glue over a more primitive API.")
	}
	if evidence.NeverOutsideLocalGraph {
		add("never_outside_local_graph", "Never appears outside addon-local flow", WeightNeverOutsideLocalGraph, false, "No XML, global, or cross-addon reinforcement was found.")
	}
	if evidence.PlaceholderLikeName {
		add("placeholder_or_computed_name", "Looks like a placeholder or computed expression", WeightPlaceholderOrComputedName, false, "The symbol name looks like a variable placeholder, interpolation artifact, or computed expression rather than a stable API contract.")
	}

	final := ClampScore(raw)
	if evidence.KnownEngineNamespace && final < ScoreMediumMin && !hasStrongContradiction(evidence) {
		overrideWeight := ScoreMediumMin - final
		if overrideWeight > 0 {
			signals = append(signals, Signal{Key: "known_namespace_override", Label: "Known namespace override", Weight: overrideWeight, Positive: true, Detail: "Known engine namespaces are promoted to at least MEDIUM when strong contradictory evidence is absent."})
			final = ScoreMediumMin
		}
	}

	assessment := Assessment{
		SymbolName: evidence.SymbolName,
		RawScore:   raw,
		FinalScore: final,
		Level:      LevelForScore(final),
		Signals:    signals,
		Evidence:   evidence,
	}
	assessment.Rationale = BuildRationale(assessment)
	return assessment
}

func ShouldPromote(assessment Assessment) bool {
	evidence := assessment.Evidence
	if evidence.PlaceholderLikeName {
		return false
	}
	if evidence.ProjectSpecificName && !evidence.KnownEngineNamespace {
		return false
	}
	if assessment.Level == LevelHigh {
		return true
	}
	if assessment.Level != LevelMedium {
		return false
	}
	return evidence.KnownEngineNamespace || evidence.DefaultUIPresence || evidence.EventBindingPresence || evidence.XMLAttributeUsageCount > 0 || evidence.UsedByXMLAndLua
}

func IsRejectedAddonLocal(assessment Assessment) bool {
	evidence := assessment.Evidence
	if ShouldPromote(assessment) {
		return false
	}
	return evidence.LocalDefinitionCount > 0 || evidence.ProjectSpecificName || evidence.PlaceholderLikeName || evidence.LocalHelperOnly || evidence.NeverOutsideLocalGraph || evidence.WrapperLikely
}

func hasStrongContradiction(evidence Evidence) bool {
	return evidence.LocalHelperOnly || evidence.NeverOutsideLocalGraph || evidence.ConflictingRoles || evidence.WrapperLikely || evidence.PlaceholderLikeName || (evidence.ProjectSpecificName && len(evidence.AddonsSeenIn) <= 1)
}
