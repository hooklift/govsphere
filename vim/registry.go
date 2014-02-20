package vim

var registry map[string]func() interface{}

func init() {
	registry = map[string]func() interface{}{

		"AboutInfo": func() interface{} { return &AboutInfo{} },

		"AccountCreatedEvent": func() interface{} { return &AccountCreatedEvent{} },

		"AccountRemovedEvent": func() interface{} { return &AccountRemovedEvent{} },

		"AccountUpdatedEvent": func() interface{} { return &AccountUpdatedEvent{} },

		"Action": func() interface{} { return &Action{} },

		"ActionParameter": func() interface{} { return &ActionParameter{} },

		"ActionType": func() interface{} { return &ActionType{} },

		"ActiveDirectoryFault": func() interface{} { return &ActiveDirectoryFault{} },

		"ActiveDirectoryProfile": func() interface{} { return &ActiveDirectoryProfile{} },

		"AdminDisabled": func() interface{} { return &AdminDisabled{} },

		"AdminNotDisabled": func() interface{} { return &AdminNotDisabled{} },

		"AdminPasswordNotChangedEvent": func() interface{} { return &AdminPasswordNotChangedEvent{} },

		"AffinityConfigured": func() interface{} { return &AffinityConfigured{} },

		"AffinityType": func() interface{} { return &AffinityType{} },

		"AfterStartupTaskScheduler": func() interface{} { return &AfterStartupTaskScheduler{} },

		"AgentInstallFailed": func() interface{} { return &AgentInstallFailed{} },

		"AgentInstallFailedReason": func() interface{} { return &AgentInstallFailedReason{} },

		"Alarm": func() interface{} { return &Alarm{} },

		"AlarmAcknowledgedEvent": func() interface{} { return &AlarmAcknowledgedEvent{} },

		"AlarmAction": func() interface{} { return &AlarmAction{} },

		"AlarmActionTriggeredEvent": func() interface{} { return &AlarmActionTriggeredEvent{} },

		"AlarmClearedEvent": func() interface{} { return &AlarmClearedEvent{} },

		"AlarmCreatedEvent": func() interface{} { return &AlarmCreatedEvent{} },

		"AlarmDescription": func() interface{} { return &AlarmDescription{} },

		"AlarmEmailCompletedEvent": func() interface{} { return &AlarmEmailCompletedEvent{} },

		"AlarmEmailFailedEvent": func() interface{} { return &AlarmEmailFailedEvent{} },

		"AlarmEvent": func() interface{} { return &AlarmEvent{} },

		"AlarmEventArgument": func() interface{} { return &AlarmEventArgument{} },

		"AlarmExpression": func() interface{} { return &AlarmExpression{} },

		"AlarmInfo": func() interface{} { return &AlarmInfo{} },

		"AlarmManager": func() interface{} { return &AlarmManager{} },

		"AlarmReconfiguredEvent": func() interface{} { return &AlarmReconfiguredEvent{} },

		"AlarmRemovedEvent": func() interface{} { return &AlarmRemovedEvent{} },

		"AlarmScriptCompleteEvent": func() interface{} { return &AlarmScriptCompleteEvent{} },

		"AlarmScriptFailedEvent": func() interface{} { return &AlarmScriptFailedEvent{} },

		"AlarmSetting": func() interface{} { return &AlarmSetting{} },

		"AlarmSnmpCompletedEvent": func() interface{} { return &AlarmSnmpCompletedEvent{} },

		"AlarmSnmpFailedEvent": func() interface{} { return &AlarmSnmpFailedEvent{} },

		"AlarmSpec": func() interface{} { return &AlarmSpec{} },

		"AlarmState": func() interface{} { return &AlarmState{} },

		"AlarmStatusChangedEvent": func() interface{} { return &AlarmStatusChangedEvent{} },

		"AlarmTriggeringAction": func() interface{} { return &AlarmTriggeringAction{} },

		"AlarmTriggeringActionTransitionSpec": func() interface{} { return &AlarmTriggeringActionTransitionSpec{} },

		"AllVirtualMachinesLicensedEvent": func() interface{} { return &AllVirtualMachinesLicensedEvent{} },

		"AlreadyAuthenticatedSessionEvent": func() interface{} { return &AlreadyAuthenticatedSessionEvent{} },

		"AlreadyBeingManaged": func() interface{} { return &AlreadyBeingManaged{} },

		"AlreadyConnected": func() interface{} { return &AlreadyConnected{} },

		"AlreadyExists": func() interface{} { return &AlreadyExists{} },

		"AlreadyUpgraded": func() interface{} { return &AlreadyUpgraded{} },

		"AndAlarmExpression": func() interface{} { return &AndAlarmExpression{} },

		"AnswerFile": func() interface{} { return &AnswerFile{} },

		"AnswerFileCreateSpec": func() interface{} { return &AnswerFileCreateSpec{} },

		"AnswerFileOptionsCreateSpec": func() interface{} { return &AnswerFileOptionsCreateSpec{} },

		"AnswerFileSerializedCreateSpec": func() interface{} { return &AnswerFileSerializedCreateSpec{} },

		"AnswerFileStatusError": func() interface{} { return &AnswerFileStatusError{} },

		"AnswerFileStatusResult": func() interface{} { return &AnswerFileStatusResult{} },

		"AnswerFileUpdateFailed": func() interface{} { return &AnswerFileUpdateFailed{} },

		"AnswerFileUpdateFailure": func() interface{} { return &AnswerFileUpdateFailure{} },

		"ApplicationQuiesceFault": func() interface{} { return &ApplicationQuiesceFault{} },

		"ApplyProfile": func() interface{} { return &ApplyProfile{} },

		"ApplyStorageRecommendationResult": func() interface{} { return &ApplyStorageRecommendationResult{} },

		"ArrayUpdateOperation": func() interface{} { return &ArrayUpdateOperation{} },

		"ArrayUpdateSpec": func() interface{} { return &ArrayUpdateSpec{} },

		"AuthMinimumAdminPermission": func() interface{} { return &AuthMinimumAdminPermission{} },

		"AuthenticationProfile": func() interface{} { return &AuthenticationProfile{} },

		"AuthorizationDescription": func() interface{} { return &AuthorizationDescription{} },

		"AuthorizationEvent": func() interface{} { return &AuthorizationEvent{} },

		"AuthorizationManager": func() interface{} { return &AuthorizationManager{} },

		"AuthorizationPrivilege": func() interface{} { return &AuthorizationPrivilege{} },

		"AuthorizationRole": func() interface{} { return &AuthorizationRole{} },

		"AutoStartAction": func() interface{} { return &AutoStartAction{} },

		"AutoStartDefaults": func() interface{} { return &AutoStartDefaults{} },

		"AutoStartPowerInfo": func() interface{} { return &AutoStartPowerInfo{} },

		"AutoStartWaitHeartbeatSetting": func() interface{} { return &AutoStartWaitHeartbeatSetting{} },

		"BackupBlobReadFailure": func() interface{} { return &BackupBlobReadFailure{} },

		"BackupBlobWriteFailure": func() interface{} { return &BackupBlobWriteFailure{} },

		"BadUsernameSessionEvent": func() interface{} { return &BadUsernameSessionEvent{} },

		"BlockedByFirewall": func() interface{} { return &BlockedByFirewall{} },

		"BoolOption": func() interface{} { return &BoolOption{} },

		"BoolPolicy": func() interface{} { return &BoolPolicy{} },

		"CAMServerRefusedConnection": func() interface{} { return &CAMServerRefusedConnection{} },

		"CanceledHostOperationEvent": func() interface{} { return &CanceledHostOperationEvent{} },

		"CannotAccessFile": func() interface{} { return &CannotAccessFile{} },

		"CannotAccessLocalSource": func() interface{} { return &CannotAccessLocalSource{} },

		"CannotAccessNetwork": func() interface{} { return &CannotAccessNetwork{} },

		"CannotAccessVmComponent": func() interface{} { return &CannotAccessVmComponent{} },

		"CannotAccessVmConfig": func() interface{} { return &CannotAccessVmConfig{} },

		"CannotAccessVmDevice": func() interface{} { return &CannotAccessVmDevice{} },

		"CannotAccessVmDisk": func() interface{} { return &CannotAccessVmDisk{} },

		"CannotAddHostWithFTVmAsStandalone": func() interface{} { return &CannotAddHostWithFTVmAsStandalone{} },

		"CannotAddHostWithFTVmToDifferentCluster": func() interface{} { return &CannotAddHostWithFTVmToDifferentCluster{} },

		"CannotAddHostWithFTVmToNonHACluster": func() interface{} { return &CannotAddHostWithFTVmToNonHACluster{} },

		"CannotChangeDrsBehaviorForFtSecondary": func() interface{} { return &CannotChangeDrsBehaviorForFtSecondary{} },

		"CannotChangeHaSettingsForFtSecondary": func() interface{} { return &CannotChangeHaSettingsForFtSecondary{} },

		"CannotChangeVsanClusterUuid": func() interface{} { return &CannotChangeVsanClusterUuid{} },

		"CannotChangeVsanNodeUuid": func() interface{} { return &CannotChangeVsanNodeUuid{} },

		"CannotCreateFile": func() interface{} { return &CannotCreateFile{} },

		"CannotDecryptPasswords": func() interface{} { return &CannotDecryptPasswords{} },

		"CannotDeleteFile": func() interface{} { return &CannotDeleteFile{} },

		"CannotDisableDrsOnClustersWithVApps": func() interface{} { return &CannotDisableDrsOnClustersWithVApps{} },

		"CannotDisableSnapshot": func() interface{} { return &CannotDisableSnapshot{} },

		"CannotDisconnectHostWithFaultToleranceVm": func() interface{} { return &CannotDisconnectHostWithFaultToleranceVm{} },

		"CannotModifyConfigCpuRequirements": func() interface{} { return &CannotModifyConfigCpuRequirements{} },

		"CannotMoveFaultToleranceVm": func() interface{} { return &CannotMoveFaultToleranceVm{} },

		"CannotMoveFaultToleranceVmMoveType": func() interface{} { return &CannotMoveFaultToleranceVmMoveType{} },

		"CannotMoveHostWithFaultToleranceVm": func() interface{} { return &CannotMoveHostWithFaultToleranceVm{} },

		"CannotMoveVmWithDeltaDisk": func() interface{} { return &CannotMoveVmWithDeltaDisk{} },

		"CannotMoveVmWithNativeDeltaDisk": func() interface{} { return &CannotMoveVmWithNativeDeltaDisk{} },

		"CannotMoveVsanEnabledHost": func() interface{} { return &CannotMoveVsanEnabledHost{} },

		"CannotPlaceWithoutPrerequisiteMoves": func() interface{} { return &CannotPlaceWithoutPrerequisiteMoves{} },

		"CannotPowerOffVmInCluster": func() interface{} { return &CannotPowerOffVmInCluster{} },

		"CannotPowerOffVmInClusterOperation": func() interface{} { return &CannotPowerOffVmInClusterOperation{} },

		"CannotReconfigureVsanWhenHaEnabled": func() interface{} { return &CannotReconfigureVsanWhenHaEnabled{} },

		"CannotUseNetwork": func() interface{} { return &CannotUseNetwork{} },

		"CannotUseNetworkReason": func() interface{} { return &CannotUseNetworkReason{} },

		"Capability": func() interface{} { return &Capability{} },

		"CheckResult": func() interface{} { return &CheckResult{} },

		"CheckTestType": func() interface{} { return &CheckTestType{} },

		"ChoiceOption": func() interface{} { return &ChoiceOption{} },

		"ClockSkew": func() interface{} { return &ClockSkew{} },

		"CloneFromSnapshotNotSupported": func() interface{} { return &CloneFromSnapshotNotSupported{} },

		"ClusterAction": func() interface{} { return &ClusterAction{} },

		"ClusterActionHistory": func() interface{} { return &ClusterActionHistory{} },

		"ClusterAffinityRuleSpec": func() interface{} { return &ClusterAffinityRuleSpec{} },

		"ClusterAntiAffinityRuleSpec": func() interface{} { return &ClusterAntiAffinityRuleSpec{} },

		"ClusterAttemptedVmInfo": func() interface{} { return &ClusterAttemptedVmInfo{} },

		"ClusterComplianceCheckedEvent": func() interface{} { return &ClusterComplianceCheckedEvent{} },

		"ClusterComputeResource": func() interface{} { return &ClusterComputeResource{} },

		"ClusterComputeResourceSummary": func() interface{} { return &ClusterComputeResourceSummary{} },

		"ClusterConfigInfo": func() interface{} { return &ClusterConfigInfo{} },

		"ClusterConfigInfoEx": func() interface{} { return &ClusterConfigInfoEx{} },

		"ClusterConfigSpec": func() interface{} { return &ClusterConfigSpec{} },

		"ClusterConfigSpecEx": func() interface{} { return &ClusterConfigSpecEx{} },

		"ClusterCreatedEvent": func() interface{} { return &ClusterCreatedEvent{} },

		"ClusterDasAamHostInfo": func() interface{} { return &ClusterDasAamHostInfo{} },

		"ClusterDasAamNodeState": func() interface{} { return &ClusterDasAamNodeState{} },

		"ClusterDasAamNodeStateDasState": func() interface{} { return &ClusterDasAamNodeStateDasState{} },

		"ClusterDasAdmissionControlInfo": func() interface{} { return &ClusterDasAdmissionControlInfo{} },

		"ClusterDasAdmissionControlPolicy": func() interface{} { return &ClusterDasAdmissionControlPolicy{} },

		"ClusterDasAdvancedRuntimeInfo": func() interface{} { return &ClusterDasAdvancedRuntimeInfo{} },

		"ClusterDasConfigInfo": func() interface{} { return &ClusterDasConfigInfo{} },

		"ClusterDasConfigInfoHBDatastoreCandidate": func() interface{} { return &ClusterDasConfigInfoHBDatastoreCandidate{} },

		"ClusterDasConfigInfoServiceState": func() interface{} { return &ClusterDasConfigInfoServiceState{} },

		"ClusterDasConfigInfoVmMonitoringState": func() interface{} { return &ClusterDasConfigInfoVmMonitoringState{} },

		"ClusterDasData": func() interface{} { return &ClusterDasData{} },

		"ClusterDasDataSummary": func() interface{} { return &ClusterDasDataSummary{} },

		"ClusterDasFailoverLevelAdvancedRuntimeInfo": func() interface{} { return &ClusterDasFailoverLevelAdvancedRuntimeInfo{} },

		"ClusterDasFailoverLevelAdvancedRuntimeInfoHostSlots": func() interface{} { return &ClusterDasFailoverLevelAdvancedRuntimeInfoHostSlots{} },

		"ClusterDasFailoverLevelAdvancedRuntimeInfoSlotInfo": func() interface{} { return &ClusterDasFailoverLevelAdvancedRuntimeInfoSlotInfo{} },

		"ClusterDasFailoverLevelAdvancedRuntimeInfoVmSlots": func() interface{} { return &ClusterDasFailoverLevelAdvancedRuntimeInfoVmSlots{} },

		"ClusterDasFdmAvailabilityState": func() interface{} { return &ClusterDasFdmAvailabilityState{} },

		"ClusterDasFdmHostState": func() interface{} { return &ClusterDasFdmHostState{} },

		"ClusterDasHostInfo": func() interface{} { return &ClusterDasHostInfo{} },

		"ClusterDasHostRecommendation": func() interface{} { return &ClusterDasHostRecommendation{} },

		"ClusterDasVmConfigInfo": func() interface{} { return &ClusterDasVmConfigInfo{} },

		"ClusterDasVmConfigSpec": func() interface{} { return &ClusterDasVmConfigSpec{} },

		"ClusterDasVmSettings": func() interface{} { return &ClusterDasVmSettings{} },

		"ClusterDasVmSettingsIsolationResponse": func() interface{} { return &ClusterDasVmSettingsIsolationResponse{} },

		"ClusterDasVmSettingsRestartPriority": func() interface{} { return &ClusterDasVmSettingsRestartPriority{} },

		"ClusterDestroyedEvent": func() interface{} { return &ClusterDestroyedEvent{} },

		"ClusterDpmConfigInfo": func() interface{} { return &ClusterDpmConfigInfo{} },

		"ClusterDpmHostConfigInfo": func() interface{} { return &ClusterDpmHostConfigInfo{} },

		"ClusterDpmHostConfigSpec": func() interface{} { return &ClusterDpmHostConfigSpec{} },

		"ClusterDrsConfigInfo": func() interface{} { return &ClusterDrsConfigInfo{} },

		"ClusterDrsFaults": func() interface{} { return &ClusterDrsFaults{} },

		"ClusterDrsFaultsFaultsByVirtualDisk": func() interface{} { return &ClusterDrsFaultsFaultsByVirtualDisk{} },

		"ClusterDrsFaultsFaultsByVm": func() interface{} { return &ClusterDrsFaultsFaultsByVm{} },

		"ClusterDrsMigration": func() interface{} { return &ClusterDrsMigration{} },

		"ClusterDrsRecommendation": func() interface{} { return &ClusterDrsRecommendation{} },

		"ClusterDrsVmConfigInfo": func() interface{} { return &ClusterDrsVmConfigInfo{} },

		"ClusterDrsVmConfigSpec": func() interface{} { return &ClusterDrsVmConfigSpec{} },

		"ClusterEnterMaintenanceResult": func() interface{} { return &ClusterEnterMaintenanceResult{} },

		"ClusterEvent": func() interface{} { return &ClusterEvent{} },

		"ClusterFailoverHostAdmissionControlInfo": func() interface{} { return &ClusterFailoverHostAdmissionControlInfo{} },

		"ClusterFailoverHostAdmissionControlInfoHostStatus": func() interface{} { return &ClusterFailoverHostAdmissionControlInfoHostStatus{} },

		"ClusterFailoverHostAdmissionControlPolicy": func() interface{} { return &ClusterFailoverHostAdmissionControlPolicy{} },

		"ClusterFailoverLevelAdmissionControlInfo": func() interface{} { return &ClusterFailoverLevelAdmissionControlInfo{} },

		"ClusterFailoverLevelAdmissionControlPolicy": func() interface{} { return &ClusterFailoverLevelAdmissionControlPolicy{} },

		"ClusterFailoverResourcesAdmissionControlInfo": func() interface{} { return &ClusterFailoverResourcesAdmissionControlInfo{} },

		"ClusterFailoverResourcesAdmissionControlPolicy": func() interface{} { return &ClusterFailoverResourcesAdmissionControlPolicy{} },

		"ClusterFixedSizeSlotPolicy": func() interface{} { return &ClusterFixedSizeSlotPolicy{} },

		"ClusterGroupInfo": func() interface{} { return &ClusterGroupInfo{} },

		"ClusterGroupSpec": func() interface{} { return &ClusterGroupSpec{} },

		"ClusterHostGroup": func() interface{} { return &ClusterHostGroup{} },

		"ClusterHostPowerAction": func() interface{} { return &ClusterHostPowerAction{} },

		"ClusterHostRecommendation": func() interface{} { return &ClusterHostRecommendation{} },

		"ClusterInitialPlacementAction": func() interface{} { return &ClusterInitialPlacementAction{} },

		"ClusterMigrationAction": func() interface{} { return &ClusterMigrationAction{} },

		"ClusterNotAttemptedVmInfo": func() interface{} { return &ClusterNotAttemptedVmInfo{} },

		"ClusterOvercommittedEvent": func() interface{} { return &ClusterOvercommittedEvent{} },

		"ClusterPowerOnVmOption": func() interface{} { return &ClusterPowerOnVmOption{} },

		"ClusterPowerOnVmResult": func() interface{} { return &ClusterPowerOnVmResult{} },

		"ClusterProfile": func() interface{} { return &ClusterProfile{} },

		"ClusterProfileCompleteConfigSpec": func() interface{} { return &ClusterProfileCompleteConfigSpec{} },

		"ClusterProfileConfigInfo": func() interface{} { return &ClusterProfileConfigInfo{} },

		"ClusterProfileConfigServiceCreateSpec": func() interface{} { return &ClusterProfileConfigServiceCreateSpec{} },

		"ClusterProfileConfigSpec": func() interface{} { return &ClusterProfileConfigSpec{} },

		"ClusterProfileCreateSpec": func() interface{} { return &ClusterProfileCreateSpec{} },

		"ClusterProfileManager": func() interface{} { return &ClusterProfileManager{} },

		"ClusterProfileServiceType": func() interface{} { return &ClusterProfileServiceType{} },

		"ClusterRecommendation": func() interface{} { return &ClusterRecommendation{} },

		"ClusterReconfiguredEvent": func() interface{} { return &ClusterReconfiguredEvent{} },

		"ClusterRuleInfo": func() interface{} { return &ClusterRuleInfo{} },

		"ClusterRuleSpec": func() interface{} { return &ClusterRuleSpec{} },

		"ClusterSlotPolicy": func() interface{} { return &ClusterSlotPolicy{} },

		"ClusterStatusChangedEvent": func() interface{} { return &ClusterStatusChangedEvent{} },

		"ClusterVmGroup": func() interface{} { return &ClusterVmGroup{} },

		"ClusterVmHostRuleInfo": func() interface{} { return &ClusterVmHostRuleInfo{} },

		"ClusterVmToolsMonitoringSettings": func() interface{} { return &ClusterVmToolsMonitoringSettings{} },

		"CollectorAddressUnset": func() interface{} { return &CollectorAddressUnset{} },

		"ComplianceFailure": func() interface{} { return &ComplianceFailure{} },

		"ComplianceLocator": func() interface{} { return &ComplianceLocator{} },

		"ComplianceProfile": func() interface{} { return &ComplianceProfile{} },

		"ComplianceResult": func() interface{} { return &ComplianceResult{} },

		"ComplianceResultStatus": func() interface{} { return &ComplianceResultStatus{} },

		"CompositePolicyOption": func() interface{} { return &CompositePolicyOption{} },

		"ComputeResource": func() interface{} { return &ComputeResource{} },

		"ComputeResourceConfigInfo": func() interface{} { return &ComputeResourceConfigInfo{} },

		"ComputeResourceConfigSpec": func() interface{} { return &ComputeResourceConfigSpec{} },

		"ComputeResourceEventArgument": func() interface{} { return &ComputeResourceEventArgument{} },

		"ComputeResourceHostSPBMLicenseInfo": func() interface{} { return &ComputeResourceHostSPBMLicenseInfo{} },

		"ComputeResourceHostSPBMLicenseInfoHostSPBMLicenseState": func() interface{} { return &ComputeResourceHostSPBMLicenseInfoHostSPBMLicenseState{} },

		"ComputeResourceSummary": func() interface{} { return &ComputeResourceSummary{} },

		"ConcurrentAccess": func() interface{} { return &ConcurrentAccess{} },

		"ConfigSpecOperation": func() interface{} { return &ConfigSpecOperation{} },

		"ConfigTarget": func() interface{} { return &ConfigTarget{} },

		"ConflictingConfiguration": func() interface{} { return &ConflictingConfiguration{} },

		"ConflictingConfigurationConfig": func() interface{} { return &ConflictingConfigurationConfig{} },

		"ConflictingDatastoreFound": func() interface{} { return &ConflictingDatastoreFound{} },

		"ConnectedIso": func() interface{} { return &ConnectedIso{} },

		"ContainerView": func() interface{} { return &ContainerView{} },

		"CpuCompatibilityUnknown": func() interface{} { return &CpuCompatibilityUnknown{} },

		"CpuHotPlugNotSupported": func() interface{} { return &CpuHotPlugNotSupported{} },

		"CpuIncompatible": func() interface{} { return &CpuIncompatible{} },

		"CpuIncompatible1ECX": func() interface{} { return &CpuIncompatible1ECX{} },

		"CpuIncompatible81EDX": func() interface{} { return &CpuIncompatible81EDX{} },

		"CreateTaskAction": func() interface{} { return &CreateTaskAction{} },

		"CustomFieldDef": func() interface{} { return &CustomFieldDef{} },

		"CustomFieldDefAddedEvent": func() interface{} { return &CustomFieldDefAddedEvent{} },

		"CustomFieldDefEvent": func() interface{} { return &CustomFieldDefEvent{} },

		"CustomFieldDefRemovedEvent": func() interface{} { return &CustomFieldDefRemovedEvent{} },

		"CustomFieldDefRenamedEvent": func() interface{} { return &CustomFieldDefRenamedEvent{} },

		"CustomFieldEvent": func() interface{} { return &CustomFieldEvent{} },

		"CustomFieldStringValue": func() interface{} { return &CustomFieldStringValue{} },

		"CustomFieldValue": func() interface{} { return &CustomFieldValue{} },

		"CustomFieldValueChangedEvent": func() interface{} { return &CustomFieldValueChangedEvent{} },

		"CustomFieldsManager": func() interface{} { return &CustomFieldsManager{} },

		"CustomizationAdapterMapping": func() interface{} { return &CustomizationAdapterMapping{} },

		"CustomizationAutoIpV6Generator": func() interface{} { return &CustomizationAutoIpV6Generator{} },

		"CustomizationCustomIpGenerator": func() interface{} { return &CustomizationCustomIpGenerator{} },

		"CustomizationCustomIpV6Generator": func() interface{} { return &CustomizationCustomIpV6Generator{} },

		"CustomizationCustomName": func() interface{} { return &CustomizationCustomName{} },

		"CustomizationDhcpIpGenerator": func() interface{} { return &CustomizationDhcpIpGenerator{} },

		"CustomizationDhcpIpV6Generator": func() interface{} { return &CustomizationDhcpIpV6Generator{} },

		"CustomizationEvent": func() interface{} { return &CustomizationEvent{} },

		"CustomizationFailed": func() interface{} { return &CustomizationFailed{} },

		"CustomizationFault": func() interface{} { return &CustomizationFault{} },

		"CustomizationFixedIp": func() interface{} { return &CustomizationFixedIp{} },

		"CustomizationFixedIpV6": func() interface{} { return &CustomizationFixedIpV6{} },

		"CustomizationFixedName": func() interface{} { return &CustomizationFixedName{} },

		"CustomizationGlobalIPSettings": func() interface{} { return &CustomizationGlobalIPSettings{} },

		"CustomizationGuiRunOnce": func() interface{} { return &CustomizationGuiRunOnce{} },

		"CustomizationGuiUnattended": func() interface{} { return &CustomizationGuiUnattended{} },

		"CustomizationIPSettings": func() interface{} { return &CustomizationIPSettings{} },

		"CustomizationIPSettingsIpV6AddressSpec": func() interface{} { return &CustomizationIPSettingsIpV6AddressSpec{} },

		"CustomizationIdentification": func() interface{} { return &CustomizationIdentification{} },

		"CustomizationIdentitySettings": func() interface{} { return &CustomizationIdentitySettings{} },

		"CustomizationIpGenerator": func() interface{} { return &CustomizationIpGenerator{} },

		"CustomizationIpV6Generator": func() interface{} { return &CustomizationIpV6Generator{} },

		"CustomizationLicenseDataMode": func() interface{} { return &CustomizationLicenseDataMode{} },

		"CustomizationLicenseFilePrintData": func() interface{} { return &CustomizationLicenseFilePrintData{} },

		"CustomizationLinuxIdentityFailed": func() interface{} { return &CustomizationLinuxIdentityFailed{} },

		"CustomizationLinuxOptions": func() interface{} { return &CustomizationLinuxOptions{} },

		"CustomizationLinuxPrep": func() interface{} { return &CustomizationLinuxPrep{} },

		"CustomizationName": func() interface{} { return &CustomizationName{} },

		"CustomizationNetBIOSMode": func() interface{} { return &CustomizationNetBIOSMode{} },

		"CustomizationNetworkSetupFailed": func() interface{} { return &CustomizationNetworkSetupFailed{} },

		"CustomizationOptions": func() interface{} { return &CustomizationOptions{} },

		"CustomizationPassword": func() interface{} { return &CustomizationPassword{} },

		"CustomizationPending": func() interface{} { return &CustomizationPending{} },

		"CustomizationPrefixName": func() interface{} { return &CustomizationPrefixName{} },

		"CustomizationSpec": func() interface{} { return &CustomizationSpec{} },

		"CustomizationSpecInfo": func() interface{} { return &CustomizationSpecInfo{} },

		"CustomizationSpecItem": func() interface{} { return &CustomizationSpecItem{} },

		"CustomizationSpecManager": func() interface{} { return &CustomizationSpecManager{} },

		"CustomizationStartedEvent": func() interface{} { return &CustomizationStartedEvent{} },

		"CustomizationStatelessIpV6Generator": func() interface{} { return &CustomizationStatelessIpV6Generator{} },

		"CustomizationSucceeded": func() interface{} { return &CustomizationSucceeded{} },

		"CustomizationSysprep": func() interface{} { return &CustomizationSysprep{} },

		"CustomizationSysprepFailed": func() interface{} { return &CustomizationSysprepFailed{} },

		"CustomizationSysprepRebootOption": func() interface{} { return &CustomizationSysprepRebootOption{} },

		"CustomizationSysprepText": func() interface{} { return &CustomizationSysprepText{} },

		"CustomizationUnknownFailure": func() interface{} { return &CustomizationUnknownFailure{} },

		"CustomizationUnknownIpGenerator": func() interface{} { return &CustomizationUnknownIpGenerator{} },

		"CustomizationUnknownIpV6Generator": func() interface{} { return &CustomizationUnknownIpV6Generator{} },

		"CustomizationUnknownName": func() interface{} { return &CustomizationUnknownName{} },

		"CustomizationUserData": func() interface{} { return &CustomizationUserData{} },

		"CustomizationVirtualMachineName": func() interface{} { return &CustomizationVirtualMachineName{} },

		"CustomizationWinOptions": func() interface{} { return &CustomizationWinOptions{} },

		"DVPortConfigInfo": func() interface{} { return &DVPortConfigInfo{} },

		"DVPortConfigSpec": func() interface{} { return &DVPortConfigSpec{} },

		"DVPortNotSupported": func() interface{} { return &DVPortNotSupported{} },

		"DVPortSetting": func() interface{} { return &DVPortSetting{} },

		"DVPortState": func() interface{} { return &DVPortState{} },

		"DVPortStatus": func() interface{} { return &DVPortStatus{} },

		"DVPortStatusVmDirectPathGen2InactiveReasonNetwork": func() interface{} { return &DVPortStatusVmDirectPathGen2InactiveReasonNetwork{} },

		"DVPortStatusVmDirectPathGen2InactiveReasonOther": func() interface{} { return &DVPortStatusVmDirectPathGen2InactiveReasonOther{} },

		"DVPortgroupConfigInfo": func() interface{} { return &DVPortgroupConfigInfo{} },

		"DVPortgroupConfigSpec": func() interface{} { return &DVPortgroupConfigSpec{} },

		"DVPortgroupCreatedEvent": func() interface{} { return &DVPortgroupCreatedEvent{} },

		"DVPortgroupDestroyedEvent": func() interface{} { return &DVPortgroupDestroyedEvent{} },

		"DVPortgroupEvent": func() interface{} { return &DVPortgroupEvent{} },

		"DVPortgroupPolicy": func() interface{} { return &DVPortgroupPolicy{} },

		"DVPortgroupReconfiguredEvent": func() interface{} { return &DVPortgroupReconfiguredEvent{} },

		"DVPortgroupRenamedEvent": func() interface{} { return &DVPortgroupRenamedEvent{} },

		"DVPortgroupSelection": func() interface{} { return &DVPortgroupSelection{} },

		"DVSBackupRestoreCapability": func() interface{} { return &DVSBackupRestoreCapability{} },

		"DVSCapability": func() interface{} { return &DVSCapability{} },

		"DVSConfigInfo": func() interface{} { return &DVSConfigInfo{} },

		"DVSConfigSpec": func() interface{} { return &DVSConfigSpec{} },

		"DVSContactInfo": func() interface{} { return &DVSContactInfo{} },

		"DVSCreateSpec": func() interface{} { return &DVSCreateSpec{} },

		"DVSFailureCriteria": func() interface{} { return &DVSFailureCriteria{} },

		"DVSFeatureCapability": func() interface{} { return &DVSFeatureCapability{} },

		"DVSHealthCheckCapability": func() interface{} { return &DVSHealthCheckCapability{} },

		"DVSHealthCheckConfig": func() interface{} { return &DVSHealthCheckConfig{} },

		"DVSHostLocalPortInfo": func() interface{} { return &DVSHostLocalPortInfo{} },

		"DVSManagerDvsConfigTarget": func() interface{} { return &DVSManagerDvsConfigTarget{} },

		"DVSNameArrayUplinkPortPolicy": func() interface{} { return &DVSNameArrayUplinkPortPolicy{} },

		"DVSNetworkResourceManagementCapability": func() interface{} { return &DVSNetworkResourceManagementCapability{} },

		"DVSNetworkResourcePool": func() interface{} { return &DVSNetworkResourcePool{} },

		"DVSNetworkResourcePoolAllocationInfo": func() interface{} { return &DVSNetworkResourcePoolAllocationInfo{} },

		"DVSNetworkResourcePoolConfigSpec": func() interface{} { return &DVSNetworkResourcePoolConfigSpec{} },

		"DVSPolicy": func() interface{} { return &DVSPolicy{} },

		"DVSRollbackCapability": func() interface{} { return &DVSRollbackCapability{} },

		"DVSRuntimeInfo": func() interface{} { return &DVSRuntimeInfo{} },

		"DVSSecurityPolicy": func() interface{} { return &DVSSecurityPolicy{} },

		"DVSSelection": func() interface{} { return &DVSSelection{} },

		"DVSSummary": func() interface{} { return &DVSSummary{} },

		"DVSTrafficShapingPolicy": func() interface{} { return &DVSTrafficShapingPolicy{} },

		"DVSUplinkPortPolicy": func() interface{} { return &DVSUplinkPortPolicy{} },

		"DVSVendorSpecificConfig": func() interface{} { return &DVSVendorSpecificConfig{} },

		"DailyTaskScheduler": func() interface{} { return &DailyTaskScheduler{} },

		"DasAdmissionControlDisabledEvent": func() interface{} { return &DasAdmissionControlDisabledEvent{} },

		"DasAdmissionControlEnabledEvent": func() interface{} { return &DasAdmissionControlEnabledEvent{} },

		"DasAgentFoundEvent": func() interface{} { return &DasAgentFoundEvent{} },

		"DasAgentUnavailableEvent": func() interface{} { return &DasAgentUnavailableEvent{} },

		"DasClusterIsolatedEvent": func() interface{} { return &DasClusterIsolatedEvent{} },

		"DasConfigFault": func() interface{} { return &DasConfigFault{} },

		"DasConfigFaultDasConfigFaultReason": func() interface{} { return &DasConfigFaultDasConfigFaultReason{} },

		"DasDisabledEvent": func() interface{} { return &DasDisabledEvent{} },

		"DasEnabledEvent": func() interface{} { return &DasEnabledEvent{} },

		"DasHeartbeatDatastoreInfo": func() interface{} { return &DasHeartbeatDatastoreInfo{} },

		"DasHostFailedEvent": func() interface{} { return &DasHostFailedEvent{} },

		"DasHostIsolatedEvent": func() interface{} { return &DasHostIsolatedEvent{} },

		"DasVmPriority": func() interface{} { return &DasVmPriority{} },

		"DatabaseError": func() interface{} { return &DatabaseError{} },

		"DatabaseSizeEstimate": func() interface{} { return &DatabaseSizeEstimate{} },

		"DatabaseSizeParam": func() interface{} { return &DatabaseSizeParam{} },

		"Datacenter": func() interface{} { return &Datacenter{} },

		"DatacenterConfigInfo": func() interface{} { return &DatacenterConfigInfo{} },

		"DatacenterConfigSpec": func() interface{} { return &DatacenterConfigSpec{} },

		"DatacenterCreatedEvent": func() interface{} { return &DatacenterCreatedEvent{} },

		"DatacenterEvent": func() interface{} { return &DatacenterEvent{} },

		"DatacenterEventArgument": func() interface{} { return &DatacenterEventArgument{} },

		"DatacenterMismatch": func() interface{} { return &DatacenterMismatch{} },

		"DatacenterMismatchArgument": func() interface{} { return &DatacenterMismatchArgument{} },

		"DatacenterRenamedEvent": func() interface{} { return &DatacenterRenamedEvent{} },

		"Datastore": func() interface{} { return &Datastore{} },

		"DatastoreAccessible": func() interface{} { return &DatastoreAccessible{} },

		"DatastoreCapability": func() interface{} { return &DatastoreCapability{} },

		"DatastoreCapacityIncreasedEvent": func() interface{} { return &DatastoreCapacityIncreasedEvent{} },

		"DatastoreDestroyedEvent": func() interface{} { return &DatastoreDestroyedEvent{} },

		"DatastoreDiscoveredEvent": func() interface{} { return &DatastoreDiscoveredEvent{} },

		"DatastoreDuplicatedEvent": func() interface{} { return &DatastoreDuplicatedEvent{} },

		"DatastoreEvent": func() interface{} { return &DatastoreEvent{} },

		"DatastoreEventArgument": func() interface{} { return &DatastoreEventArgument{} },

		"DatastoreFileCopiedEvent": func() interface{} { return &DatastoreFileCopiedEvent{} },

		"DatastoreFileDeletedEvent": func() interface{} { return &DatastoreFileDeletedEvent{} },

		"DatastoreFileEvent": func() interface{} { return &DatastoreFileEvent{} },

		"DatastoreFileMovedEvent": func() interface{} { return &DatastoreFileMovedEvent{} },

		"DatastoreHostMount": func() interface{} { return &DatastoreHostMount{} },

		"DatastoreIORMReconfiguredEvent": func() interface{} { return &DatastoreIORMReconfiguredEvent{} },

		"DatastoreInfo": func() interface{} { return &DatastoreInfo{} },

		"DatastoreMountPathDatastorePair": func() interface{} { return &DatastoreMountPathDatastorePair{} },

		"DatastoreNamespaceManager": func() interface{} { return &DatastoreNamespaceManager{} },

		"DatastoreNotWritableOnHost": func() interface{} { return &DatastoreNotWritableOnHost{} },

		"DatastoreOption": func() interface{} { return &DatastoreOption{} },

		"DatastorePrincipalConfigured": func() interface{} { return &DatastorePrincipalConfigured{} },

		"DatastoreRemovedOnHostEvent": func() interface{} { return &DatastoreRemovedOnHostEvent{} },

		"DatastoreRenamedEvent": func() interface{} { return &DatastoreRenamedEvent{} },

		"DatastoreRenamedOnHostEvent": func() interface{} { return &DatastoreRenamedOnHostEvent{} },

		"DatastoreSummary": func() interface{} { return &DatastoreSummary{} },

		"DatastoreSummaryMaintenanceModeState": func() interface{} { return &DatastoreSummaryMaintenanceModeState{} },

		"DateTimeProfile": func() interface{} { return &DateTimeProfile{} },

		"DayOfWeek": func() interface{} { return &DayOfWeek{} },

		"DeltaDiskFormatNotSupported": func() interface{} { return &DeltaDiskFormatNotSupported{} },

		"Description": func() interface{} { return &Description{} },

		"DestinationSwitchFull": func() interface{} { return &DestinationSwitchFull{} },

		"DestinationVsanDisabled": func() interface{} { return &DestinationVsanDisabled{} },

		"DeviceBackedVirtualDiskSpec": func() interface{} { return &DeviceBackedVirtualDiskSpec{} },

		"DeviceBackingNotSupported": func() interface{} { return &DeviceBackingNotSupported{} },

		"DeviceControllerNotSupported": func() interface{} { return &DeviceControllerNotSupported{} },

		"DeviceHotPlugNotSupported": func() interface{} { return &DeviceHotPlugNotSupported{} },

		"DeviceNotFound": func() interface{} { return &DeviceNotFound{} },

		"DeviceNotSupported": func() interface{} { return &DeviceNotSupported{} },

		"DeviceNotSupportedReason": func() interface{} { return &DeviceNotSupportedReason{} },

		"DeviceUnsupportedForVmPlatform": func() interface{} { return &DeviceUnsupportedForVmPlatform{} },

		"DeviceUnsupportedForVmVersion": func() interface{} { return &DeviceUnsupportedForVmVersion{} },

		"DiagnosticManager": func() interface{} { return &DiagnosticManager{} },

		"DiagnosticManagerBundleInfo": func() interface{} { return &DiagnosticManagerBundleInfo{} },

		"DiagnosticManagerLogCreator": func() interface{} { return &DiagnosticManagerLogCreator{} },

		"DiagnosticManagerLogDescriptor": func() interface{} { return &DiagnosticManagerLogDescriptor{} },

		"DiagnosticManagerLogFormat": func() interface{} { return &DiagnosticManagerLogFormat{} },

		"DiagnosticManagerLogHeader": func() interface{} { return &DiagnosticManagerLogHeader{} },

		"DiagnosticPartitionStorageType": func() interface{} { return &DiagnosticPartitionStorageType{} },

		"DiagnosticPartitionType": func() interface{} { return &DiagnosticPartitionType{} },

		"DirectoryNotEmpty": func() interface{} { return &DirectoryNotEmpty{} },

		"DisableAdminNotSupported": func() interface{} { return &DisableAdminNotSupported{} },

		"DisallowedChangeByService": func() interface{} { return &DisallowedChangeByService{} },

		"DisallowedChangeByServiceDisallowedChange": func() interface{} { return &DisallowedChangeByServiceDisallowedChange{} },

		"DisallowedDiskModeChange": func() interface{} { return &DisallowedDiskModeChange{} },

		"DisallowedMigrationDeviceAttached": func() interface{} { return &DisallowedMigrationDeviceAttached{} },

		"DisallowedOperationOnFailoverHost": func() interface{} { return &DisallowedOperationOnFailoverHost{} },

		"DiskChangeExtent": func() interface{} { return &DiskChangeExtent{} },

		"DiskChangeInfo": func() interface{} { return &DiskChangeInfo{} },

		"DiskHasPartitions": func() interface{} { return &DiskHasPartitions{} },

		"DiskIsLastRemainingNonSSD": func() interface{} { return &DiskIsLastRemainingNonSSD{} },

		"DiskIsNonLocal": func() interface{} { return &DiskIsNonLocal{} },

		"DiskIsUSB": func() interface{} { return &DiskIsUSB{} },

		"DiskMoveTypeNotSupported": func() interface{} { return &DiskMoveTypeNotSupported{} },

		"DiskNotSupported": func() interface{} { return &DiskNotSupported{} },

		"DiskTooSmall": func() interface{} { return &DiskTooSmall{} },

		"DistributedVirtualPort": func() interface{} { return &DistributedVirtualPort{} },

		"DistributedVirtualPortgroup": func() interface{} { return &DistributedVirtualPortgroup{} },

		"DistributedVirtualPortgroupInfo": func() interface{} { return &DistributedVirtualPortgroupInfo{} },

		"DistributedVirtualPortgroupMetaTagName": func() interface{} { return &DistributedVirtualPortgroupMetaTagName{} },

		"DistributedVirtualPortgroupPortgroupType": func() interface{} { return &DistributedVirtualPortgroupPortgroupType{} },

		"DistributedVirtualSwitch": func() interface{} { return &DistributedVirtualSwitch{} },

		"DistributedVirtualSwitchHostInfrastructureTrafficClass": func() interface{} { return &DistributedVirtualSwitchHostInfrastructureTrafficClass{} },

		"DistributedVirtualSwitchHostMember": func() interface{} { return &DistributedVirtualSwitchHostMember{} },

		"DistributedVirtualSwitchHostMemberBacking": func() interface{} { return &DistributedVirtualSwitchHostMemberBacking{} },

		"DistributedVirtualSwitchHostMemberConfigInfo": func() interface{} { return &DistributedVirtualSwitchHostMemberConfigInfo{} },

		"DistributedVirtualSwitchHostMemberConfigSpec": func() interface{} { return &DistributedVirtualSwitchHostMemberConfigSpec{} },

		"DistributedVirtualSwitchHostMemberHostComponentState": func() interface{} { return &DistributedVirtualSwitchHostMemberHostComponentState{} },

		"DistributedVirtualSwitchHostMemberPnicBacking": func() interface{} { return &DistributedVirtualSwitchHostMemberPnicBacking{} },

		"DistributedVirtualSwitchHostMemberPnicSpec": func() interface{} { return &DistributedVirtualSwitchHostMemberPnicSpec{} },

		"DistributedVirtualSwitchHostMemberRuntimeState": func() interface{} { return &DistributedVirtualSwitchHostMemberRuntimeState{} },

		"DistributedVirtualSwitchHostProductSpec": func() interface{} { return &DistributedVirtualSwitchHostProductSpec{} },

		"DistributedVirtualSwitchInfo": func() interface{} { return &DistributedVirtualSwitchInfo{} },

		"DistributedVirtualSwitchKeyedOpaqueBlob": func() interface{} { return &DistributedVirtualSwitchKeyedOpaqueBlob{} },

		"DistributedVirtualSwitchManager": func() interface{} { return &DistributedVirtualSwitchManager{} },

		"DistributedVirtualSwitchManagerCompatibilityResult": func() interface{} { return &DistributedVirtualSwitchManagerCompatibilityResult{} },

		"DistributedVirtualSwitchManagerDvsProductSpec": func() interface{} { return &DistributedVirtualSwitchManagerDvsProductSpec{} },

		"DistributedVirtualSwitchManagerHostArrayFilter": func() interface{} { return &DistributedVirtualSwitchManagerHostArrayFilter{} },

		"DistributedVirtualSwitchManagerHostContainer": func() interface{} { return &DistributedVirtualSwitchManagerHostContainer{} },

		"DistributedVirtualSwitchManagerHostContainerFilter": func() interface{} { return &DistributedVirtualSwitchManagerHostContainerFilter{} },

		"DistributedVirtualSwitchManagerHostDvsFilterSpec": func() interface{} { return &DistributedVirtualSwitchManagerHostDvsFilterSpec{} },

		"DistributedVirtualSwitchManagerHostDvsMembershipFilter": func() interface{} { return &DistributedVirtualSwitchManagerHostDvsMembershipFilter{} },

		"DistributedVirtualSwitchManagerImportResult": func() interface{} { return &DistributedVirtualSwitchManagerImportResult{} },

		"DistributedVirtualSwitchNicTeamingPolicyMode": func() interface{} { return &DistributedVirtualSwitchNicTeamingPolicyMode{} },

		"DistributedVirtualSwitchPortConnectee": func() interface{} { return &DistributedVirtualSwitchPortConnectee{} },

		"DistributedVirtualSwitchPortConnecteeConnecteeType": func() interface{} { return &DistributedVirtualSwitchPortConnecteeConnecteeType{} },

		"DistributedVirtualSwitchPortConnection": func() interface{} { return &DistributedVirtualSwitchPortConnection{} },

		"DistributedVirtualSwitchPortCriteria": func() interface{} { return &DistributedVirtualSwitchPortCriteria{} },

		"DistributedVirtualSwitchPortStatistics": func() interface{} { return &DistributedVirtualSwitchPortStatistics{} },

		"DistributedVirtualSwitchProductSpec": func() interface{} { return &DistributedVirtualSwitchProductSpec{} },

		"DistributedVirtualSwitchProductSpecOperationType": func() interface{} { return &DistributedVirtualSwitchProductSpecOperationType{} },

		"DomainNotFound": func() interface{} { return &DomainNotFound{} },

		"DpmBehavior": func() interface{} { return &DpmBehavior{} },

		"DrsBehavior": func() interface{} { return &DrsBehavior{} },

		"DrsDisabledEvent": func() interface{} { return &DrsDisabledEvent{} },

		"DrsDisabledOnVm": func() interface{} { return &DrsDisabledOnVm{} },

		"DrsEnabledEvent": func() interface{} { return &DrsEnabledEvent{} },

		"DrsEnteredStandbyModeEvent": func() interface{} { return &DrsEnteredStandbyModeEvent{} },

		"DrsEnteringStandbyModeEvent": func() interface{} { return &DrsEnteringStandbyModeEvent{} },

		"DrsExitStandbyModeFailedEvent": func() interface{} { return &DrsExitStandbyModeFailedEvent{} },

		"DrsExitedStandbyModeEvent": func() interface{} { return &DrsExitedStandbyModeEvent{} },

		"DrsExitingStandbyModeEvent": func() interface{} { return &DrsExitingStandbyModeEvent{} },

		"DrsInjectorWorkloadCorrelationState": func() interface{} { return &DrsInjectorWorkloadCorrelationState{} },

		"DrsInvocationFailedEvent": func() interface{} { return &DrsInvocationFailedEvent{} },

		"DrsRecommendationReasonCode": func() interface{} { return &DrsRecommendationReasonCode{} },

		"DrsRecoveredFromFailureEvent": func() interface{} { return &DrsRecoveredFromFailureEvent{} },

		"DrsResourceConfigureFailedEvent": func() interface{} { return &DrsResourceConfigureFailedEvent{} },

		"DrsResourceConfigureSyncedEvent": func() interface{} { return &DrsResourceConfigureSyncedEvent{} },

		"DrsRuleComplianceEvent": func() interface{} { return &DrsRuleComplianceEvent{} },

		"DrsRuleViolationEvent": func() interface{} { return &DrsRuleViolationEvent{} },

		"DrsVmMigratedEvent": func() interface{} { return &DrsVmMigratedEvent{} },

		"DrsVmPoweredOnEvent": func() interface{} { return &DrsVmPoweredOnEvent{} },

		"DrsVmotionIncompatibleFault": func() interface{} { return &DrsVmotionIncompatibleFault{} },

		"DuplicateDisks": func() interface{} { return &DuplicateDisks{} },

		"DuplicateIpDetectedEvent": func() interface{} { return &DuplicateIpDetectedEvent{} },

		"DuplicateName": func() interface{} { return &DuplicateName{} },

		"DuplicateVsanNetworkInterface": func() interface{} { return &DuplicateVsanNetworkInterface{} },

		"DvpgImportEvent": func() interface{} { return &DvpgImportEvent{} },

		"DvpgRestoreEvent": func() interface{} { return &DvpgRestoreEvent{} },

		"DvsAcceptNetworkRuleAction": func() interface{} { return &DvsAcceptNetworkRuleAction{} },

		"DvsApplyOperationFault": func() interface{} { return &DvsApplyOperationFault{} },

		"DvsApplyOperationFaultFaultOnObject": func() interface{} { return &DvsApplyOperationFaultFaultOnObject{} },

		"DvsCopyNetworkRuleAction": func() interface{} { return &DvsCopyNetworkRuleAction{} },

		"DvsCreatedEvent": func() interface{} { return &DvsCreatedEvent{} },

		"DvsDestroyedEvent": func() interface{} { return &DvsDestroyedEvent{} },

		"DvsDropNetworkRuleAction": func() interface{} { return &DvsDropNetworkRuleAction{} },

		"DvsEvent": func() interface{} { return &DvsEvent{} },

		"DvsEventArgument": func() interface{} { return &DvsEventArgument{} },

		"DvsFault": func() interface{} { return &DvsFault{} },

		"DvsFilterConfig": func() interface{} { return &DvsFilterConfig{} },

		"DvsFilterConfigSpec": func() interface{} { return &DvsFilterConfigSpec{} },

		"DvsFilterOnFailure": func() interface{} { return &DvsFilterOnFailure{} },

		"DvsFilterParameter": func() interface{} { return &DvsFilterParameter{} },

		"DvsFilterPolicy": func() interface{} { return &DvsFilterPolicy{} },

		"DvsGreEncapNetworkRuleAction": func() interface{} { return &DvsGreEncapNetworkRuleAction{} },

		"DvsHealthStatusChangeEvent": func() interface{} { return &DvsHealthStatusChangeEvent{} },

		"DvsHostBackInSyncEvent": func() interface{} { return &DvsHostBackInSyncEvent{} },

		"DvsHostJoinedEvent": func() interface{} { return &DvsHostJoinedEvent{} },

		"DvsHostLeftEvent": func() interface{} { return &DvsHostLeftEvent{} },

		"DvsHostStatusUpdated": func() interface{} { return &DvsHostStatusUpdated{} },

		"DvsHostVNicProfile": func() interface{} { return &DvsHostVNicProfile{} },

		"DvsHostWentOutOfSyncEvent": func() interface{} { return &DvsHostWentOutOfSyncEvent{} },

		"DvsImportEvent": func() interface{} { return &DvsImportEvent{} },

		"DvsIpNetworkRuleQualifier": func() interface{} { return &DvsIpNetworkRuleQualifier{} },

		"DvsIpPort": func() interface{} { return &DvsIpPort{} },

		"DvsIpPortRange": func() interface{} { return &DvsIpPortRange{} },

		"DvsLogNetworkRuleAction": func() interface{} { return &DvsLogNetworkRuleAction{} },

		"DvsMacNetworkRuleQualifier": func() interface{} { return &DvsMacNetworkRuleQualifier{} },

		"DvsMacRewriteNetworkRuleAction": func() interface{} { return &DvsMacRewriteNetworkRuleAction{} },

		"DvsMergedEvent": func() interface{} { return &DvsMergedEvent{} },

		"DvsNetworkRuleAction": func() interface{} { return &DvsNetworkRuleAction{} },

		"DvsNetworkRuleDirectionType": func() interface{} { return &DvsNetworkRuleDirectionType{} },

		"DvsNetworkRuleQualifier": func() interface{} { return &DvsNetworkRuleQualifier{} },

		"DvsNotAuthorized": func() interface{} { return &DvsNotAuthorized{} },

		"DvsOperationBulkFault": func() interface{} { return &DvsOperationBulkFault{} },

		"DvsOperationBulkFaultFaultOnHost": func() interface{} { return &DvsOperationBulkFaultFaultOnHost{} },

		"DvsOutOfSyncHostArgument": func() interface{} { return &DvsOutOfSyncHostArgument{} },

		"DvsPortBlockedEvent": func() interface{} { return &DvsPortBlockedEvent{} },

		"DvsPortConnectedEvent": func() interface{} { return &DvsPortConnectedEvent{} },

		"DvsPortCreatedEvent": func() interface{} { return &DvsPortCreatedEvent{} },

		"DvsPortDeletedEvent": func() interface{} { return &DvsPortDeletedEvent{} },

		"DvsPortDisconnectedEvent": func() interface{} { return &DvsPortDisconnectedEvent{} },

		"DvsPortEnteredPassthruEvent": func() interface{} { return &DvsPortEnteredPassthruEvent{} },

		"DvsPortExitedPassthruEvent": func() interface{} { return &DvsPortExitedPassthruEvent{} },

		"DvsPortJoinPortgroupEvent": func() interface{} { return &DvsPortJoinPortgroupEvent{} },

		"DvsPortLeavePortgroupEvent": func() interface{} { return &DvsPortLeavePortgroupEvent{} },

		"DvsPortLinkDownEvent": func() interface{} { return &DvsPortLinkDownEvent{} },

		"DvsPortLinkUpEvent": func() interface{} { return &DvsPortLinkUpEvent{} },

		"DvsPortReconfiguredEvent": func() interface{} { return &DvsPortReconfiguredEvent{} },

		"DvsPortRuntimeChangeEvent": func() interface{} { return &DvsPortRuntimeChangeEvent{} },

		"DvsPortUnblockedEvent": func() interface{} { return &DvsPortUnblockedEvent{} },

		"DvsPortVendorSpecificStateChangeEvent": func() interface{} { return &DvsPortVendorSpecificStateChangeEvent{} },

		"DvsProfile": func() interface{} { return &DvsProfile{} },

		"DvsPuntNetworkRuleAction": func() interface{} { return &DvsPuntNetworkRuleAction{} },

		"DvsRateLimitNetworkRuleAction": func() interface{} { return &DvsRateLimitNetworkRuleAction{} },

		"DvsReconfiguredEvent": func() interface{} { return &DvsReconfiguredEvent{} },

		"DvsRenamedEvent": func() interface{} { return &DvsRenamedEvent{} },

		"DvsRestoreEvent": func() interface{} { return &DvsRestoreEvent{} },

		"DvsScopeViolated": func() interface{} { return &DvsScopeViolated{} },

		"DvsServiceConsoleVNicProfile": func() interface{} { return &DvsServiceConsoleVNicProfile{} },

		"DvsSingleIpPort": func() interface{} { return &DvsSingleIpPort{} },

		"DvsSystemTrafficNetworkRuleQualifier": func() interface{} { return &DvsSystemTrafficNetworkRuleQualifier{} },

		"DvsTrafficFilterConfig": func() interface{} { return &DvsTrafficFilterConfig{} },

		"DvsTrafficFilterConfigSpec": func() interface{} { return &DvsTrafficFilterConfigSpec{} },

		"DvsTrafficRule": func() interface{} { return &DvsTrafficRule{} },

		"DvsTrafficRuleset": func() interface{} { return &DvsTrafficRuleset{} },

		"DvsUpdateTagNetworkRuleAction": func() interface{} { return &DvsUpdateTagNetworkRuleAction{} },

		"DvsUpgradeAvailableEvent": func() interface{} { return &DvsUpgradeAvailableEvent{} },

		"DvsUpgradeInProgressEvent": func() interface{} { return &DvsUpgradeInProgressEvent{} },

		"DvsUpgradeRejectedEvent": func() interface{} { return &DvsUpgradeRejectedEvent{} },

		"DvsUpgradedEvent": func() interface{} { return &DvsUpgradedEvent{} },

		"DvsVNicProfile": func() interface{} { return &DvsVNicProfile{} },

		"DynamicArray": func() interface{} { return &DynamicArray{} },

		"DynamicData": func() interface{} { return &DynamicData{} },

		"DynamicProperty": func() interface{} { return &DynamicProperty{} },

		"EVCAdmissionFailed": func() interface{} { return &EVCAdmissionFailed{} },

		"EVCAdmissionFailedCPUFeaturesForMode": func() interface{} { return &EVCAdmissionFailedCPUFeaturesForMode{} },

		"EVCAdmissionFailedCPUModel": func() interface{} { return &EVCAdmissionFailedCPUModel{} },

		"EVCAdmissionFailedCPUModelForMode": func() interface{} { return &EVCAdmissionFailedCPUModelForMode{} },

		"EVCAdmissionFailedCPUVendor": func() interface{} { return &EVCAdmissionFailedCPUVendor{} },

		"EVCAdmissionFailedCPUVendorUnknown": func() interface{} { return &EVCAdmissionFailedCPUVendorUnknown{} },

		"EVCAdmissionFailedHostDisconnected": func() interface{} { return &EVCAdmissionFailedHostDisconnected{} },

		"EVCAdmissionFailedHostSoftware": func() interface{} { return &EVCAdmissionFailedHostSoftware{} },

		"EVCAdmissionFailedHostSoftwareForMode": func() interface{} { return &EVCAdmissionFailedHostSoftwareForMode{} },

		"EVCAdmissionFailedVmActive": func() interface{} { return &EVCAdmissionFailedVmActive{} },

		"EVCMode": func() interface{} { return &EVCMode{} },

		"EightHostLimitViolated": func() interface{} { return &EightHostLimitViolated{} },

		"ElementDescription": func() interface{} { return &ElementDescription{} },

		"EnteredMaintenanceModeEvent": func() interface{} { return &EnteredMaintenanceModeEvent{} },

		"EnteredStandbyModeEvent": func() interface{} { return &EnteredStandbyModeEvent{} },

		"EnteringMaintenanceModeEvent": func() interface{} { return &EnteringMaintenanceModeEvent{} },

		"EnteringStandbyModeEvent": func() interface{} { return &EnteringStandbyModeEvent{} },

		"EntityBackup": func() interface{} { return &EntityBackup{} },

		"EntityBackupConfig": func() interface{} { return &EntityBackupConfig{} },

		"EntityEventArgument": func() interface{} { return &EntityEventArgument{} },

		"EntityImportType": func() interface{} { return &EntityImportType{} },

		"EntityPrivilege": func() interface{} { return &EntityPrivilege{} },

		"EntityType": func() interface{} { return &EntityType{} },

		"EnumDescription": func() interface{} { return &EnumDescription{} },

		"EnvironmentBrowser": func() interface{} { return &EnvironmentBrowser{} },

		"ErrorUpgradeEvent": func() interface{} { return &ErrorUpgradeEvent{} },

		"EvaluationLicenseSource": func() interface{} { return &EvaluationLicenseSource{} },

		"Event": func() interface{} { return &Event{} },

		"EventAlarmExpression": func() interface{} { return &EventAlarmExpression{} },

		"EventAlarmExpressionComparison": func() interface{} { return &EventAlarmExpressionComparison{} },

		"EventAlarmExpressionComparisonOperator": func() interface{} { return &EventAlarmExpressionComparisonOperator{} },

		"EventArgDesc": func() interface{} { return &EventArgDesc{} },

		"EventArgument": func() interface{} { return &EventArgument{} },

		"EventCategory": func() interface{} { return &EventCategory{} },

		"EventDescription": func() interface{} { return &EventDescription{} },

		"EventDescriptionEventDetail": func() interface{} { return &EventDescriptionEventDetail{} },

		"EventEventSeverity": func() interface{} { return &EventEventSeverity{} },

		"EventEx": func() interface{} { return &EventEx{} },

		"EventFilterSpec": func() interface{} { return &EventFilterSpec{} },

		"EventFilterSpecByEntity": func() interface{} { return &EventFilterSpecByEntity{} },

		"EventFilterSpecByTime": func() interface{} { return &EventFilterSpecByTime{} },

		"EventFilterSpecByUsername": func() interface{} { return &EventFilterSpecByUsername{} },

		"EventFilterSpecRecursionOption": func() interface{} { return &EventFilterSpecRecursionOption{} },

		"EventHistoryCollector": func() interface{} { return &EventHistoryCollector{} },

		"EventManager": func() interface{} { return &EventManager{} },

		"ExitMaintenanceModeEvent": func() interface{} { return &ExitMaintenanceModeEvent{} },

		"ExitStandbyModeFailedEvent": func() interface{} { return &ExitStandbyModeFailedEvent{} },

		"ExitedStandbyModeEvent": func() interface{} { return &ExitedStandbyModeEvent{} },

		"ExitingStandbyModeEvent": func() interface{} { return &ExitingStandbyModeEvent{} },

		"ExpiredAddonLicense": func() interface{} { return &ExpiredAddonLicense{} },

		"ExpiredEditionLicense": func() interface{} { return &ExpiredEditionLicense{} },

		"ExpiredFeatureLicense": func() interface{} { return &ExpiredFeatureLicense{} },

		"ExtExtendedProductInfo": func() interface{} { return &ExtExtendedProductInfo{} },

		"ExtManagedEntityInfo": func() interface{} { return &ExtManagedEntityInfo{} },

		"ExtSolutionManagerInfo": func() interface{} { return &ExtSolutionManagerInfo{} },

		"ExtSolutionManagerInfoTabInfo": func() interface{} { return &ExtSolutionManagerInfoTabInfo{} },

		"ExtendedDescription": func() interface{} { return &ExtendedDescription{} },

		"ExtendedElementDescription": func() interface{} { return &ExtendedElementDescription{} },

		"ExtendedEvent": func() interface{} { return &ExtendedEvent{} },

		"ExtendedEventPair": func() interface{} { return &ExtendedEventPair{} },

		"ExtendedFault": func() interface{} { return &ExtendedFault{} },

		"ExtensibleManagedObject": func() interface{} { return &ExtensibleManagedObject{} },

		"Extension": func() interface{} { return &Extension{} },

		"ExtensionClientInfo": func() interface{} { return &ExtensionClientInfo{} },

		"ExtensionEventTypeInfo": func() interface{} { return &ExtensionEventTypeInfo{} },

		"ExtensionFaultTypeInfo": func() interface{} { return &ExtensionFaultTypeInfo{} },

		"ExtensionHealthInfo": func() interface{} { return &ExtensionHealthInfo{} },

		"ExtensionManager": func() interface{} { return &ExtensionManager{} },

		"ExtensionManagerIpAllocationUsage": func() interface{} { return &ExtensionManagerIpAllocationUsage{} },

		"ExtensionOvfConsumerInfo": func() interface{} { return &ExtensionOvfConsumerInfo{} },

		"ExtensionPrivilegeInfo": func() interface{} { return &ExtensionPrivilegeInfo{} },

		"ExtensionResourceInfo": func() interface{} { return &ExtensionResourceInfo{} },

		"ExtensionServerInfo": func() interface{} { return &ExtensionServerInfo{} },

		"ExtensionTaskTypeInfo": func() interface{} { return &ExtensionTaskTypeInfo{} },

		"FailToEnableSPBM": func() interface{} { return &FailToEnableSPBM{} },

		"FailToLockFaultToleranceVMs": func() interface{} { return &FailToLockFaultToleranceVMs{} },

		"FailoverLevelRestored": func() interface{} { return &FailoverLevelRestored{} },

		"FaultToleranceAntiAffinityViolated": func() interface{} { return &FaultToleranceAntiAffinityViolated{} },

		"FaultToleranceCannotEditMem": func() interface{} { return &FaultToleranceCannotEditMem{} },

		"FaultToleranceConfigInfo": func() interface{} { return &FaultToleranceConfigInfo{} },

		"FaultToleranceCpuIncompatible": func() interface{} { return &FaultToleranceCpuIncompatible{} },

		"FaultToleranceNeedsThickDisk": func() interface{} { return &FaultToleranceNeedsThickDisk{} },

		"FaultToleranceNotLicensed": func() interface{} { return &FaultToleranceNotLicensed{} },

		"FaultToleranceNotSameBuild": func() interface{} { return &FaultToleranceNotSameBuild{} },

		"FaultTolerancePrimaryConfigInfo": func() interface{} { return &FaultTolerancePrimaryConfigInfo{} },

		"FaultTolerancePrimaryPowerOnNotAttempted": func() interface{} { return &FaultTolerancePrimaryPowerOnNotAttempted{} },

		"FaultToleranceSecondaryConfigInfo": func() interface{} { return &FaultToleranceSecondaryConfigInfo{} },

		"FaultToleranceSecondaryOpResult": func() interface{} { return &FaultToleranceSecondaryOpResult{} },

		"FaultToleranceVmNotDasProtected": func() interface{} { return &FaultToleranceVmNotDasProtected{} },

		"FcoeConfig": func() interface{} { return &FcoeConfig{} },

		"FcoeConfigFcoeCapabilities": func() interface{} { return &FcoeConfigFcoeCapabilities{} },

		"FcoeConfigFcoeSpecification": func() interface{} { return &FcoeConfigFcoeSpecification{} },

		"FcoeConfigVlanRange": func() interface{} { return &FcoeConfigVlanRange{} },

		"FcoeFault": func() interface{} { return &FcoeFault{} },

		"FcoeFaultPnicHasNoPortSet": func() interface{} { return &FcoeFaultPnicHasNoPortSet{} },

		"FeatureRequirementsNotMet": func() interface{} { return &FeatureRequirementsNotMet{} },

		"FibreChannelPortType": func() interface{} { return &FibreChannelPortType{} },

		"FileAlreadyExists": func() interface{} { return &FileAlreadyExists{} },

		"FileBackedPortNotSupported": func() interface{} { return &FileBackedPortNotSupported{} },

		"FileBackedVirtualDiskSpec": func() interface{} { return &FileBackedVirtualDiskSpec{} },

		"FileFault": func() interface{} { return &FileFault{} },

		"FileInfo": func() interface{} { return &FileInfo{} },

		"FileLocked": func() interface{} { return &FileLocked{} },

		"FileManager": func() interface{} { return &FileManager{} },

		"FileNameTooLong": func() interface{} { return &FileNameTooLong{} },

		"FileNotFound": func() interface{} { return &FileNotFound{} },

		"FileNotWritable": func() interface{} { return &FileNotWritable{} },

		"FileQuery": func() interface{} { return &FileQuery{} },

		"FileQueryFlags": func() interface{} { return &FileQueryFlags{} },

		"FileSystemMountInfoVStorageSupportStatus": func() interface{} { return &FileSystemMountInfoVStorageSupportStatus{} },

		"FileTooLarge": func() interface{} { return &FileTooLarge{} },

		"FileTransferInformation": func() interface{} { return &FileTransferInformation{} },

		"FilesystemQuiesceFault": func() interface{} { return &FilesystemQuiesceFault{} },

		"FirewallProfile": func() interface{} { return &FirewallProfile{} },

		"FirewallProfileRulesetProfile": func() interface{} { return &FirewallProfileRulesetProfile{} },

		"FloatOption": func() interface{} { return &FloatOption{} },

		"FloppyImageFileInfo": func() interface{} { return &FloppyImageFileInfo{} },

		"FloppyImageFileQuery": func() interface{} { return &FloppyImageFileQuery{} },

		"Folder": func() interface{} { return &Folder{} },

		"FolderEventArgument": func() interface{} { return &FolderEventArgument{} },

		"FolderFileInfo": func() interface{} { return &FolderFileInfo{} },

		"FolderFileQuery": func() interface{} { return &FolderFileQuery{} },

		"FtIssuesOnHost": func() interface{} { return &FtIssuesOnHost{} },

		"FtIssuesOnHostHostSelectionType": func() interface{} { return &FtIssuesOnHostHostSelectionType{} },

		"FullStorageVMotionNotSupported": func() interface{} { return &FullStorageVMotionNotSupported{} },

		"GeneralEvent": func() interface{} { return &GeneralEvent{} },

		"GeneralHostErrorEvent": func() interface{} { return &GeneralHostErrorEvent{} },

		"GeneralHostInfoEvent": func() interface{} { return &GeneralHostInfoEvent{} },

		"GeneralHostWarningEvent": func() interface{} { return &GeneralHostWarningEvent{} },

		"GeneralUserEvent": func() interface{} { return &GeneralUserEvent{} },

		"GeneralVmErrorEvent": func() interface{} { return &GeneralVmErrorEvent{} },

		"GeneralVmInfoEvent": func() interface{} { return &GeneralVmInfoEvent{} },

		"GeneralVmWarningEvent": func() interface{} { return &GeneralVmWarningEvent{} },

		"GenericDrsFault": func() interface{} { return &GenericDrsFault{} },

		"GenericVmConfigFault": func() interface{} { return &GenericVmConfigFault{} },

		"GhostDvsProxySwitchDetectedEvent": func() interface{} { return &GhostDvsProxySwitchDetectedEvent{} },

		"GhostDvsProxySwitchRemovedEvent": func() interface{} { return &GhostDvsProxySwitchRemovedEvent{} },

		"GlobalMessageChangedEvent": func() interface{} { return &GlobalMessageChangedEvent{} },

		"GroupAlarmAction": func() interface{} { return &GroupAlarmAction{} },

		"GuestAuthManager": func() interface{} { return &GuestAuthManager{} },

		"GuestAuthentication": func() interface{} { return &GuestAuthentication{} },

		"GuestAuthenticationChallenge": func() interface{} { return &GuestAuthenticationChallenge{} },

		"GuestComponentsOutOfDate": func() interface{} { return &GuestComponentsOutOfDate{} },

		"GuestDiskInfo": func() interface{} { return &GuestDiskInfo{} },

		"GuestFileAttributes": func() interface{} { return &GuestFileAttributes{} },

		"GuestFileInfo": func() interface{} { return &GuestFileInfo{} },

		"GuestFileManager": func() interface{} { return &GuestFileManager{} },

		"GuestFileType": func() interface{} { return &GuestFileType{} },

		"GuestInfo": func() interface{} { return &GuestInfo{} },

		"GuestInfoAppStateType": func() interface{} { return &GuestInfoAppStateType{} },

		"GuestInfoNamespaceGenerationInfo": func() interface{} { return &GuestInfoNamespaceGenerationInfo{} },

		"GuestListFileInfo": func() interface{} { return &GuestListFileInfo{} },

		"GuestNicInfo": func() interface{} { return &GuestNicInfo{} },

		"GuestOperationsFault": func() interface{} { return &GuestOperationsFault{} },

		"GuestOperationsManager": func() interface{} { return &GuestOperationsManager{} },

		"GuestOperationsUnavailable": func() interface{} { return &GuestOperationsUnavailable{} },

		"GuestOsDescriptor": func() interface{} { return &GuestOsDescriptor{} },

		"GuestOsDescriptorFirmwareType": func() interface{} { return &GuestOsDescriptorFirmwareType{} },

		"GuestOsDescriptorSupportLevel": func() interface{} { return &GuestOsDescriptorSupportLevel{} },

		"GuestPermissionDenied": func() interface{} { return &GuestPermissionDenied{} },

		"GuestPosixFileAttributes": func() interface{} { return &GuestPosixFileAttributes{} },

		"GuestProcessInfo": func() interface{} { return &GuestProcessInfo{} },

		"GuestProcessManager": func() interface{} { return &GuestProcessManager{} },

		"GuestProcessNotFound": func() interface{} { return &GuestProcessNotFound{} },

		"GuestProgramSpec": func() interface{} { return &GuestProgramSpec{} },

		"GuestScreenInfo": func() interface{} { return &GuestScreenInfo{} },

		"GuestStackInfo": func() interface{} { return &GuestStackInfo{} },

		"GuestWindowsFileAttributes": func() interface{} { return &GuestWindowsFileAttributes{} },

		"GuestWindowsProgramSpec": func() interface{} { return &GuestWindowsProgramSpec{} },

		"HAErrorsAtDest": func() interface{} { return &HAErrorsAtDest{} },

		"HbrManagerReplicationVmInfo": func() interface{} { return &HbrManagerReplicationVmInfo{} },

		"HealthStatusChangedEvent": func() interface{} { return &HealthStatusChangedEvent{} },

		"HealthSystemRuntime": func() interface{} { return &HealthSystemRuntime{} },

		"HistoryCollector": func() interface{} { return &HistoryCollector{} },

		"HostAccessRestrictedToManagementServer": func() interface{} { return &HostAccessRestrictedToManagementServer{} },

		"HostAccountSpec": func() interface{} { return &HostAccountSpec{} },

		"HostActiveDirectory": func() interface{} { return &HostActiveDirectory{} },

		"HostActiveDirectoryAuthentication": func() interface{} { return &HostActiveDirectoryAuthentication{} },

		"HostActiveDirectoryInfo": func() interface{} { return &HostActiveDirectoryInfo{} },

		"HostActiveDirectoryInfoDomainMembershipStatus": func() interface{} { return &HostActiveDirectoryInfoDomainMembershipStatus{} },

		"HostActiveDirectorySpec": func() interface{} { return &HostActiveDirectorySpec{} },

		"HostAddFailedEvent": func() interface{} { return &HostAddFailedEvent{} },

		"HostAddedEvent": func() interface{} { return &HostAddedEvent{} },

		"HostAdminDisableEvent": func() interface{} { return &HostAdminDisableEvent{} },

		"HostAdminEnableEvent": func() interface{} { return &HostAdminEnableEvent{} },

		"HostApplyProfile": func() interface{} { return &HostApplyProfile{} },

		"HostAuthenticationManager": func() interface{} { return &HostAuthenticationManager{} },

		"HostAuthenticationManagerInfo": func() interface{} { return &HostAuthenticationManagerInfo{} },

		"HostAuthenticationStore": func() interface{} { return &HostAuthenticationStore{} },

		"HostAuthenticationStoreInfo": func() interface{} { return &HostAuthenticationStoreInfo{} },

		"HostAutoStartManager": func() interface{} { return &HostAutoStartManager{} },

		"HostAutoStartManagerConfig": func() interface{} { return &HostAutoStartManagerConfig{} },

		"HostBIOSInfo": func() interface{} { return &HostBIOSInfo{} },

		"HostBlockAdapterTargetTransport": func() interface{} { return &HostBlockAdapterTargetTransport{} },

		"HostBlockHba": func() interface{} { return &HostBlockHba{} },

		"HostBootDevice": func() interface{} { return &HostBootDevice{} },

		"HostBootDeviceInfo": func() interface{} { return &HostBootDeviceInfo{} },

		"HostBootDeviceSystem": func() interface{} { return &HostBootDeviceSystem{} },

		"HostCacheConfigurationInfo": func() interface{} { return &HostCacheConfigurationInfo{} },

		"HostCacheConfigurationManager": func() interface{} { return &HostCacheConfigurationManager{} },

		"HostCacheConfigurationSpec": func() interface{} { return &HostCacheConfigurationSpec{} },

		"HostCapability": func() interface{} { return &HostCapability{} },

		"HostCapabilityFtUnsupportedReason": func() interface{} { return &HostCapabilityFtUnsupportedReason{} },

		"HostCapabilityVmDirectPathGen2UnsupportedReason": func() interface{} { return &HostCapabilityVmDirectPathGen2UnsupportedReason{} },

		"HostCnxFailedAccountFailedEvent": func() interface{} { return &HostCnxFailedAccountFailedEvent{} },

		"HostCnxFailedAlreadyManagedEvent": func() interface{} { return &HostCnxFailedAlreadyManagedEvent{} },

		"HostCnxFailedBadCcagentEvent": func() interface{} { return &HostCnxFailedBadCcagentEvent{} },

		"HostCnxFailedBadUsernameEvent": func() interface{} { return &HostCnxFailedBadUsernameEvent{} },

		"HostCnxFailedBadVersionEvent": func() interface{} { return &HostCnxFailedBadVersionEvent{} },

		"HostCnxFailedCcagentUpgradeEvent": func() interface{} { return &HostCnxFailedCcagentUpgradeEvent{} },

		"HostCnxFailedEvent": func() interface{} { return &HostCnxFailedEvent{} },

		"HostCnxFailedNetworkErrorEvent": func() interface{} { return &HostCnxFailedNetworkErrorEvent{} },

		"HostCnxFailedNoAccessEvent": func() interface{} { return &HostCnxFailedNoAccessEvent{} },

		"HostCnxFailedNoConnectionEvent": func() interface{} { return &HostCnxFailedNoConnectionEvent{} },

		"HostCnxFailedNoLicenseEvent": func() interface{} { return &HostCnxFailedNoLicenseEvent{} },

		"HostCnxFailedNotFoundEvent": func() interface{} { return &HostCnxFailedNotFoundEvent{} },

		"HostCnxFailedTimeoutEvent": func() interface{} { return &HostCnxFailedTimeoutEvent{} },

		"HostCommunication": func() interface{} { return &HostCommunication{} },

		"HostComplianceCheckedEvent": func() interface{} { return &HostComplianceCheckedEvent{} },

		"HostCompliantEvent": func() interface{} { return &HostCompliantEvent{} },

		"HostConfigAppliedEvent": func() interface{} { return &HostConfigAppliedEvent{} },

		"HostConfigChange": func() interface{} { return &HostConfigChange{} },

		"HostConfigChangeMode": func() interface{} { return &HostConfigChangeMode{} },

		"HostConfigChangeOperation": func() interface{} { return &HostConfigChangeOperation{} },

		"HostConfigFailed": func() interface{} { return &HostConfigFailed{} },

		"HostConfigFault": func() interface{} { return &HostConfigFault{} },

		"HostConfigInfo": func() interface{} { return &HostConfigInfo{} },

		"HostConfigManager": func() interface{} { return &HostConfigManager{} },

		"HostConfigSpec": func() interface{} { return &HostConfigSpec{} },

		"HostConfigSummary": func() interface{} { return &HostConfigSummary{} },

		"HostConnectFault": func() interface{} { return &HostConnectFault{} },

		"HostConnectInfo": func() interface{} { return &HostConnectInfo{} },

		"HostConnectInfoNetworkInfo": func() interface{} { return &HostConnectInfoNetworkInfo{} },

		"HostConnectSpec": func() interface{} { return &HostConnectSpec{} },

		"HostConnectedEvent": func() interface{} { return &HostConnectedEvent{} },

		"HostConnectionLostEvent": func() interface{} { return &HostConnectionLostEvent{} },

		"HostCpuIdInfo": func() interface{} { return &HostCpuIdInfo{} },

		"HostCpuInfo": func() interface{} { return &HostCpuInfo{} },

		"HostCpuPackage": func() interface{} { return &HostCpuPackage{} },

		"HostCpuPackageVendor": func() interface{} { return &HostCpuPackageVendor{} },

		"HostCpuPowerManagementInfo": func() interface{} { return &HostCpuPowerManagementInfo{} },

		"HostCpuPowerManagementInfoPolicyType": func() interface{} { return &HostCpuPowerManagementInfoPolicyType{} },

		"HostCpuSchedulerSystem": func() interface{} { return &HostCpuSchedulerSystem{} },

		"HostDasDisabledEvent": func() interface{} { return &HostDasDisabledEvent{} },

		"HostDasDisablingEvent": func() interface{} { return &HostDasDisablingEvent{} },

		"HostDasEnabledEvent": func() interface{} { return &HostDasEnabledEvent{} },

		"HostDasEnablingEvent": func() interface{} { return &HostDasEnablingEvent{} },

		"HostDasErrorEvent": func() interface{} { return &HostDasErrorEvent{} },

		"HostDasErrorEventHostDasErrorReason": func() interface{} { return &HostDasErrorEventHostDasErrorReason{} },

		"HostDasEvent": func() interface{} { return &HostDasEvent{} },

		"HostDasOkEvent": func() interface{} { return &HostDasOkEvent{} },

		"HostDatastoreBrowser": func() interface{} { return &HostDatastoreBrowser{} },

		"HostDatastoreBrowserSearchResults": func() interface{} { return &HostDatastoreBrowserSearchResults{} },

		"HostDatastoreBrowserSearchSpec": func() interface{} { return &HostDatastoreBrowserSearchSpec{} },

		"HostDatastoreConnectInfo": func() interface{} { return &HostDatastoreConnectInfo{} },

		"HostDatastoreExistsConnectInfo": func() interface{} { return &HostDatastoreExistsConnectInfo{} },

		"HostDatastoreNameConflictConnectInfo": func() interface{} { return &HostDatastoreNameConflictConnectInfo{} },

		"HostDatastoreSystem": func() interface{} { return &HostDatastoreSystem{} },

		"HostDatastoreSystemCapabilities": func() interface{} { return &HostDatastoreSystemCapabilities{} },

		"HostDateTimeConfig": func() interface{} { return &HostDateTimeConfig{} },

		"HostDateTimeInfo": func() interface{} { return &HostDateTimeInfo{} },

		"HostDateTimeSystem": func() interface{} { return &HostDateTimeSystem{} },

		"HostDateTimeSystemTimeZone": func() interface{} { return &HostDateTimeSystemTimeZone{} },

		"HostDevice": func() interface{} { return &HostDevice{} },

		"HostDhcpService": func() interface{} { return &HostDhcpService{} },

		"HostDhcpServiceConfig": func() interface{} { return &HostDhcpServiceConfig{} },

		"HostDhcpServiceSpec": func() interface{} { return &HostDhcpServiceSpec{} },

		"HostDiagnosticPartition": func() interface{} { return &HostDiagnosticPartition{} },

		"HostDiagnosticPartitionCreateDescription": func() interface{} { return &HostDiagnosticPartitionCreateDescription{} },

		"HostDiagnosticPartitionCreateOption": func() interface{} { return &HostDiagnosticPartitionCreateOption{} },

		"HostDiagnosticPartitionCreateSpec": func() interface{} { return &HostDiagnosticPartitionCreateSpec{} },

		"HostDiagnosticSystem": func() interface{} { return &HostDiagnosticSystem{} },

		"HostDigestInfo": func() interface{} { return &HostDigestInfo{} },

		"HostDigestInfoDigestMethodType": func() interface{} { return &HostDigestInfoDigestMethodType{} },

		"HostDirectoryStore": func() interface{} { return &HostDirectoryStore{} },

		"HostDirectoryStoreInfo": func() interface{} { return &HostDirectoryStoreInfo{} },

		"HostDisconnectedEvent": func() interface{} { return &HostDisconnectedEvent{} },

		"HostDisconnectedEventReasonCode": func() interface{} { return &HostDisconnectedEventReasonCode{} },

		"HostDiskConfigurationResult": func() interface{} { return &HostDiskConfigurationResult{} },

		"HostDiskDimensions": func() interface{} { return &HostDiskDimensions{} },

		"HostDiskDimensionsChs": func() interface{} { return &HostDiskDimensionsChs{} },

		"HostDiskDimensionsLba": func() interface{} { return &HostDiskDimensionsLba{} },

		"HostDiskMappingInfo": func() interface{} { return &HostDiskMappingInfo{} },

		"HostDiskMappingOption": func() interface{} { return &HostDiskMappingOption{} },

		"HostDiskMappingPartitionInfo": func() interface{} { return &HostDiskMappingPartitionInfo{} },

		"HostDiskMappingPartitionOption": func() interface{} { return &HostDiskMappingPartitionOption{} },

		"HostDiskPartitionAttributes": func() interface{} { return &HostDiskPartitionAttributes{} },

		"HostDiskPartitionBlockRange": func() interface{} { return &HostDiskPartitionBlockRange{} },

		"HostDiskPartitionInfo": func() interface{} { return &HostDiskPartitionInfo{} },

		"HostDiskPartitionInfoPartitionFormat": func() interface{} { return &HostDiskPartitionInfoPartitionFormat{} },

		"HostDiskPartitionInfoType": func() interface{} { return &HostDiskPartitionInfoType{} },

		"HostDiskPartitionLayout": func() interface{} { return &HostDiskPartitionLayout{} },

		"HostDiskPartitionSpec": func() interface{} { return &HostDiskPartitionSpec{} },

		"HostDnsConfig": func() interface{} { return &HostDnsConfig{} },

		"HostDnsConfigSpec": func() interface{} { return &HostDnsConfigSpec{} },

		"HostEnableAdminFailedEvent": func() interface{} { return &HostEnableAdminFailedEvent{} },

		"HostEsxAgentHostManager": func() interface{} { return &HostEsxAgentHostManager{} },

		"HostEsxAgentHostManagerConfigInfo": func() interface{} { return &HostEsxAgentHostManagerConfigInfo{} },

		"HostEvent": func() interface{} { return &HostEvent{} },

		"HostEventArgument": func() interface{} { return &HostEventArgument{} },

		"HostExtraNetworksEvent": func() interface{} { return &HostExtraNetworksEvent{} },

		"HostFeatureCapability": func() interface{} { return &HostFeatureCapability{} },

		"HostFeatureMask": func() interface{} { return &HostFeatureMask{} },

		"HostFeatureVersionInfo": func() interface{} { return &HostFeatureVersionInfo{} },

		"HostFeatureVersionKey": func() interface{} { return &HostFeatureVersionKey{} },

		"HostFibreChannelHba": func() interface{} { return &HostFibreChannelHba{} },

		"HostFibreChannelOverEthernetHba": func() interface{} { return &HostFibreChannelOverEthernetHba{} },

		"HostFibreChannelOverEthernetHbaLinkInfo": func() interface{} { return &HostFibreChannelOverEthernetHbaLinkInfo{} },

		"HostFibreChannelOverEthernetTargetTransport": func() interface{} { return &HostFibreChannelOverEthernetTargetTransport{} },

		"HostFibreChannelTargetTransport": func() interface{} { return &HostFibreChannelTargetTransport{} },

		"HostFileAccess": func() interface{} { return &HostFileAccess{} },

		"HostFileSystemMountInfo": func() interface{} { return &HostFileSystemMountInfo{} },

		"HostFileSystemVolume": func() interface{} { return &HostFileSystemVolume{} },

		"HostFileSystemVolumeInfo": func() interface{} { return &HostFileSystemVolumeInfo{} },

		"HostFirewallConfig": func() interface{} { return &HostFirewallConfig{} },

		"HostFirewallConfigRuleSetConfig": func() interface{} { return &HostFirewallConfigRuleSetConfig{} },

		"HostFirewallDefaultPolicy": func() interface{} { return &HostFirewallDefaultPolicy{} },

		"HostFirewallInfo": func() interface{} { return &HostFirewallInfo{} },

		"HostFirewallRule": func() interface{} { return &HostFirewallRule{} },

		"HostFirewallRuleDirection": func() interface{} { return &HostFirewallRuleDirection{} },

		"HostFirewallRulePortType": func() interface{} { return &HostFirewallRulePortType{} },

		"HostFirewallRuleProtocol": func() interface{} { return &HostFirewallRuleProtocol{} },

		"HostFirewallRuleset": func() interface{} { return &HostFirewallRuleset{} },

		"HostFirewallRulesetIpList": func() interface{} { return &HostFirewallRulesetIpList{} },

		"HostFirewallRulesetIpNetwork": func() interface{} { return &HostFirewallRulesetIpNetwork{} },

		"HostFirewallRulesetRulesetSpec": func() interface{} { return &HostFirewallRulesetRulesetSpec{} },

		"HostFirewallSystem": func() interface{} { return &HostFirewallSystem{} },

		"HostFirmwareSystem": func() interface{} { return &HostFirmwareSystem{} },

		"HostFlagInfo": func() interface{} { return &HostFlagInfo{} },

		"HostForceMountedInfo": func() interface{} { return &HostForceMountedInfo{} },

		"HostGetShortNameFailedEvent": func() interface{} { return &HostGetShortNameFailedEvent{} },

		"HostGraphicsInfo": func() interface{} { return &HostGraphicsInfo{} },

		"HostGraphicsInfoGraphicsType": func() interface{} { return &HostGraphicsInfoGraphicsType{} },

		"HostGraphicsManager": func() interface{} { return &HostGraphicsManager{} },

		"HostHardwareElementInfo": func() interface{} { return &HostHardwareElementInfo{} },

		"HostHardwareElementStatus": func() interface{} { return &HostHardwareElementStatus{} },

		"HostHardwareInfo": func() interface{} { return &HostHardwareInfo{} },

		"HostHardwareStatusInfo": func() interface{} { return &HostHardwareStatusInfo{} },

		"HostHardwareSummary": func() interface{} { return &HostHardwareSummary{} },

		"HostHealthStatusSystem": func() interface{} { return &HostHealthStatusSystem{} },

		"HostHostBusAdapter": func() interface{} { return &HostHostBusAdapter{} },

		"HostHyperThreadScheduleInfo": func() interface{} { return &HostHyperThreadScheduleInfo{} },

		"HostImageAcceptanceLevel": func() interface{} { return &HostImageAcceptanceLevel{} },

		"HostImageConfigManager": func() interface{} { return &HostImageConfigManager{} },

		"HostImageProfileSummary": func() interface{} { return &HostImageProfileSummary{} },

		"HostInAuditModeEvent": func() interface{} { return &HostInAuditModeEvent{} },

		"HostInDomain": func() interface{} { return &HostInDomain{} },

		"HostIncompatibleForFaultTolerance": func() interface{} { return &HostIncompatibleForFaultTolerance{} },

		"HostIncompatibleForFaultToleranceReason": func() interface{} { return &HostIncompatibleForFaultToleranceReason{} },

		"HostIncompatibleForRecordReplay": func() interface{} { return &HostIncompatibleForRecordReplay{} },

		"HostIncompatibleForRecordReplayReason": func() interface{} { return &HostIncompatibleForRecordReplayReason{} },

		"HostInternetScsiHba": func() interface{} { return &HostInternetScsiHba{} },

		"HostInternetScsiHbaAuthenticationCapabilities": func() interface{} { return &HostInternetScsiHbaAuthenticationCapabilities{} },

		"HostInternetScsiHbaAuthenticationProperties": func() interface{} { return &HostInternetScsiHbaAuthenticationProperties{} },

		"HostInternetScsiHbaChapAuthenticationType": func() interface{} { return &HostInternetScsiHbaChapAuthenticationType{} },

		"HostInternetScsiHbaDigestCapabilities": func() interface{} { return &HostInternetScsiHbaDigestCapabilities{} },

		"HostInternetScsiHbaDigestProperties": func() interface{} { return &HostInternetScsiHbaDigestProperties{} },

		"HostInternetScsiHbaDigestType": func() interface{} { return &HostInternetScsiHbaDigestType{} },

		"HostInternetScsiHbaDiscoveryCapabilities": func() interface{} { return &HostInternetScsiHbaDiscoveryCapabilities{} },

		"HostInternetScsiHbaDiscoveryProperties": func() interface{} { return &HostInternetScsiHbaDiscoveryProperties{} },

		"HostInternetScsiHbaIPCapabilities": func() interface{} { return &HostInternetScsiHbaIPCapabilities{} },

		"HostInternetScsiHbaIPProperties": func() interface{} { return &HostInternetScsiHbaIPProperties{} },

		"HostInternetScsiHbaNetworkBindingSupportType": func() interface{} { return &HostInternetScsiHbaNetworkBindingSupportType{} },

		"HostInternetScsiHbaParamValue": func() interface{} { return &HostInternetScsiHbaParamValue{} },

		"HostInternetScsiHbaSendTarget": func() interface{} { return &HostInternetScsiHbaSendTarget{} },

		"HostInternetScsiHbaStaticTarget": func() interface{} { return &HostInternetScsiHbaStaticTarget{} },

		"HostInternetScsiHbaStaticTargetTargetDiscoveryMethod": func() interface{} { return &HostInternetScsiHbaStaticTargetTargetDiscoveryMethod{} },

		"HostInternetScsiHbaTargetSet": func() interface{} { return &HostInternetScsiHbaTargetSet{} },

		"HostInternetScsiTargetTransport": func() interface{} { return &HostInternetScsiTargetTransport{} },

		"HostInventoryFull": func() interface{} { return &HostInventoryFull{} },

		"HostInventoryFullEvent": func() interface{} { return &HostInventoryFullEvent{} },

		"HostInventoryUnreadableEvent": func() interface{} { return &HostInventoryUnreadableEvent{} },

		"HostIpChangedEvent": func() interface{} { return &HostIpChangedEvent{} },

		"HostIpConfig": func() interface{} { return &HostIpConfig{} },

		"HostIpConfigIpV6Address": func() interface{} { return &HostIpConfigIpV6Address{} },

		"HostIpConfigIpV6AddressConfigType": func() interface{} { return &HostIpConfigIpV6AddressConfigType{} },

		"HostIpConfigIpV6AddressConfiguration": func() interface{} { return &HostIpConfigIpV6AddressConfiguration{} },

		"HostIpConfigIpV6AddressStatus": func() interface{} { return &HostIpConfigIpV6AddressStatus{} },

		"HostIpInconsistentEvent": func() interface{} { return &HostIpInconsistentEvent{} },

		"HostIpRouteConfig": func() interface{} { return &HostIpRouteConfig{} },

		"HostIpRouteConfigSpec": func() interface{} { return &HostIpRouteConfigSpec{} },

		"HostIpRouteEntry": func() interface{} { return &HostIpRouteEntry{} },

		"HostIpRouteOp": func() interface{} { return &HostIpRouteOp{} },

		"HostIpRouteTableConfig": func() interface{} { return &HostIpRouteTableConfig{} },

		"HostIpRouteTableInfo": func() interface{} { return &HostIpRouteTableInfo{} },

		"HostIpToShortNameFailedEvent": func() interface{} { return &HostIpToShortNameFailedEvent{} },

		"HostIpmiInfo": func() interface{} { return &HostIpmiInfo{} },

		"HostIsolationIpPingFailedEvent": func() interface{} { return &HostIsolationIpPingFailedEvent{} },

		"HostKernelModuleSystem": func() interface{} { return &HostKernelModuleSystem{} },

		"HostLicensableResourceInfo": func() interface{} { return &HostLicensableResourceInfo{} },

		"HostLicensableResourceKey": func() interface{} { return &HostLicensableResourceKey{} },

		"HostLicenseConnectInfo": func() interface{} { return &HostLicenseConnectInfo{} },

		"HostLicenseExpiredEvent": func() interface{} { return &HostLicenseExpiredEvent{} },

		"HostLicenseSpec": func() interface{} { return &HostLicenseSpec{} },

		"HostListSummary": func() interface{} { return &HostListSummary{} },

		"HostListSummaryQuickStats": func() interface{} { return &HostListSummaryQuickStats{} },

		"HostLocalAccountManager": func() interface{} { return &HostLocalAccountManager{} },

		"HostLocalAuthentication": func() interface{} { return &HostLocalAuthentication{} },

		"HostLocalAuthenticationInfo": func() interface{} { return &HostLocalAuthenticationInfo{} },

		"HostLocalFileSystemVolume": func() interface{} { return &HostLocalFileSystemVolume{} },

		"HostLocalFileSystemVolumeSpec": func() interface{} { return &HostLocalFileSystemVolumeSpec{} },

		"HostLocalPortCreatedEvent": func() interface{} { return &HostLocalPortCreatedEvent{} },

		"HostLowLevelProvisioningManagerDiskLayoutSpec": func() interface{} { return &HostLowLevelProvisioningManagerDiskLayoutSpec{} },

		"HostLowLevelProvisioningManagerReloadTarget": func() interface{} { return &HostLowLevelProvisioningManagerReloadTarget{} },

		"HostLowLevelProvisioningManagerSnapshotLayoutSpec": func() interface{} { return &HostLowLevelProvisioningManagerSnapshotLayoutSpec{} },

		"HostLowLevelProvisioningManagerVmMigrationStatus": func() interface{} { return &HostLowLevelProvisioningManagerVmMigrationStatus{} },

		"HostLowLevelProvisioningManagerVmRecoveryInfo": func() interface{} { return &HostLowLevelProvisioningManagerVmRecoveryInfo{} },

		"HostMaintenanceSpec": func() interface{} { return &HostMaintenanceSpec{} },

		"HostMemberHealthCheckResult": func() interface{} { return &HostMemberHealthCheckResult{} },

		"HostMemberRuntimeInfo": func() interface{} { return &HostMemberRuntimeInfo{} },

		"HostMemberUplinkHealthCheckResult": func() interface{} { return &HostMemberUplinkHealthCheckResult{} },

		"HostMemoryProfile": func() interface{} { return &HostMemoryProfile{} },

		"HostMemorySpec": func() interface{} { return &HostMemorySpec{} },

		"HostMemorySystem": func() interface{} { return &HostMemorySystem{} },

		"HostMissingNetworksEvent": func() interface{} { return &HostMissingNetworksEvent{} },

		"HostMonitoringStateChangedEvent": func() interface{} { return &HostMonitoringStateChangedEvent{} },

		"HostMountInfo": func() interface{} { return &HostMountInfo{} },

		"HostMountInfoInaccessibleReason": func() interface{} { return &HostMountInfoInaccessibleReason{} },

		"HostMountMode": func() interface{} { return &HostMountMode{} },

		"HostMultipathInfo": func() interface{} { return &HostMultipathInfo{} },

		"HostMultipathInfoFixedLogicalUnitPolicy": func() interface{} { return &HostMultipathInfoFixedLogicalUnitPolicy{} },

		"HostMultipathInfoLogicalUnit": func() interface{} { return &HostMultipathInfoLogicalUnit{} },

		"HostMultipathInfoLogicalUnitPolicy": func() interface{} { return &HostMultipathInfoLogicalUnitPolicy{} },

		"HostMultipathInfoLogicalUnitStorageArrayTypePolicy": func() interface{} { return &HostMultipathInfoLogicalUnitStorageArrayTypePolicy{} },

		"HostMultipathInfoPath": func() interface{} { return &HostMultipathInfoPath{} },

		"HostMultipathStateInfo": func() interface{} { return &HostMultipathStateInfo{} },

		"HostMultipathStateInfoPath": func() interface{} { return &HostMultipathStateInfoPath{} },

		"HostNasVolume": func() interface{} { return &HostNasVolume{} },

		"HostNasVolumeConfig": func() interface{} { return &HostNasVolumeConfig{} },

		"HostNasVolumeSpec": func() interface{} { return &HostNasVolumeSpec{} },

		"HostNatService": func() interface{} { return &HostNatService{} },

		"HostNatServiceConfig": func() interface{} { return &HostNatServiceConfig{} },

		"HostNatServiceNameServiceSpec": func() interface{} { return &HostNatServiceNameServiceSpec{} },

		"HostNatServicePortForwardSpec": func() interface{} { return &HostNatServicePortForwardSpec{} },

		"HostNatServiceSpec": func() interface{} { return &HostNatServiceSpec{} },

		"HostNetCapabilities": func() interface{} { return &HostNetCapabilities{} },

		"HostNetOffloadCapabilities": func() interface{} { return &HostNetOffloadCapabilities{} },

		"HostNetStackInstance": func() interface{} { return &HostNetStackInstance{} },

		"HostNetStackInstanceCongestionControlAlgorithmType": func() interface{} { return &HostNetStackInstanceCongestionControlAlgorithmType{} },

		"HostNetStackInstanceSystemStackKey": func() interface{} { return &HostNetStackInstanceSystemStackKey{} },

		"HostNetworkConfig": func() interface{} { return &HostNetworkConfig{} },

		"HostNetworkConfigNetStackSpec": func() interface{} { return &HostNetworkConfigNetStackSpec{} },

		"HostNetworkConfigResult": func() interface{} { return &HostNetworkConfigResult{} },

		"HostNetworkInfo": func() interface{} { return &HostNetworkInfo{} },

		"HostNetworkPolicy": func() interface{} { return &HostNetworkPolicy{} },

		"HostNetworkSecurityPolicy": func() interface{} { return &HostNetworkSecurityPolicy{} },

		"HostNetworkSystem": func() interface{} { return &HostNetworkSystem{} },

		"HostNetworkTrafficShapingPolicy": func() interface{} { return &HostNetworkTrafficShapingPolicy{} },

		"HostNewNetworkConnectInfo": func() interface{} { return &HostNewNetworkConnectInfo{} },

		"HostNicFailureCriteria": func() interface{} { return &HostNicFailureCriteria{} },

		"HostNicOrderPolicy": func() interface{} { return &HostNicOrderPolicy{} },

		"HostNicTeamingPolicy": func() interface{} { return &HostNicTeamingPolicy{} },

		"HostNoAvailableNetworksEvent": func() interface{} { return &HostNoAvailableNetworksEvent{} },

		"HostNoHAEnabledPortGroupsEvent": func() interface{} { return &HostNoHAEnabledPortGroupsEvent{} },

		"HostNoRedundantManagementNetworkEvent": func() interface{} { return &HostNoRedundantManagementNetworkEvent{} },

		"HostNonCompliantEvent": func() interface{} { return &HostNonCompliantEvent{} },

		"HostNotConnected": func() interface{} { return &HostNotConnected{} },

		"HostNotInClusterEvent": func() interface{} { return &HostNotInClusterEvent{} },

		"HostNotReachable": func() interface{} { return &HostNotReachable{} },

		"HostNtpConfig": func() interface{} { return &HostNtpConfig{} },

		"HostNumaInfo": func() interface{} { return &HostNumaInfo{} },

		"HostNumaNode": func() interface{} { return &HostNumaNode{} },

		"HostNumericSensorHealthState": func() interface{} { return &HostNumericSensorHealthState{} },

		"HostNumericSensorInfo": func() interface{} { return &HostNumericSensorInfo{} },

		"HostNumericSensorType": func() interface{} { return &HostNumericSensorType{} },

		"HostOpaqueNetworkInfo": func() interface{} { return &HostOpaqueNetworkInfo{} },

		"HostOpaqueSwitch": func() interface{} { return &HostOpaqueSwitch{} },

		"HostOvercommittedEvent": func() interface{} { return &HostOvercommittedEvent{} },

		"HostParallelScsiHba": func() interface{} { return &HostParallelScsiHba{} },

		"HostParallelScsiTargetTransport": func() interface{} { return &HostParallelScsiTargetTransport{} },

		"HostPatchManager": func() interface{} { return &HostPatchManager{} },

		"HostPatchManagerInstallState": func() interface{} { return &HostPatchManagerInstallState{} },

		"HostPatchManagerIntegrityStatus": func() interface{} { return &HostPatchManagerIntegrityStatus{} },

		"HostPatchManagerLocator": func() interface{} { return &HostPatchManagerLocator{} },

		"HostPatchManagerPatchManagerOperationSpec": func() interface{} { return &HostPatchManagerPatchManagerOperationSpec{} },

		"HostPatchManagerReason": func() interface{} { return &HostPatchManagerReason{} },

		"HostPatchManagerResult": func() interface{} { return &HostPatchManagerResult{} },

		"HostPatchManagerStatus": func() interface{} { return &HostPatchManagerStatus{} },

		"HostPatchManagerStatusPrerequisitePatch": func() interface{} { return &HostPatchManagerStatusPrerequisitePatch{} },

		"HostPathSelectionPolicyOption": func() interface{} { return &HostPathSelectionPolicyOption{} },

		"HostPciDevice": func() interface{} { return &HostPciDevice{} },

		"HostPciPassthruConfig": func() interface{} { return &HostPciPassthruConfig{} },

		"HostPciPassthruInfo": func() interface{} { return &HostPciPassthruInfo{} },

		"HostPciPassthruSystem": func() interface{} { return &HostPciPassthruSystem{} },

		"HostPlugStoreTopology": func() interface{} { return &HostPlugStoreTopology{} },

		"HostPlugStoreTopologyAdapter": func() interface{} { return &HostPlugStoreTopologyAdapter{} },

		"HostPlugStoreTopologyDevice": func() interface{} { return &HostPlugStoreTopologyDevice{} },

		"HostPlugStoreTopologyPath": func() interface{} { return &HostPlugStoreTopologyPath{} },

		"HostPlugStoreTopologyPlugin": func() interface{} { return &HostPlugStoreTopologyPlugin{} },

		"HostPlugStoreTopologyTarget": func() interface{} { return &HostPlugStoreTopologyTarget{} },

		"HostPortGroup": func() interface{} { return &HostPortGroup{} },

		"HostPortGroupConfig": func() interface{} { return &HostPortGroupConfig{} },

		"HostPortGroupPort": func() interface{} { return &HostPortGroupPort{} },

		"HostPortGroupProfile": func() interface{} { return &HostPortGroupProfile{} },

		"HostPortGroupSpec": func() interface{} { return &HostPortGroupSpec{} },

		"HostPosixAccountSpec": func() interface{} { return &HostPosixAccountSpec{} },

		"HostPowerOpFailed": func() interface{} { return &HostPowerOpFailed{} },

		"HostPowerOperationType": func() interface{} { return &HostPowerOperationType{} },

		"HostPowerPolicy": func() interface{} { return &HostPowerPolicy{} },

		"HostPowerSystem": func() interface{} { return &HostPowerSystem{} },

		"HostPrimaryAgentNotShortNameEvent": func() interface{} { return &HostPrimaryAgentNotShortNameEvent{} },

		"HostProfile": func() interface{} { return &HostProfile{} },

		"HostProfileAppliedEvent": func() interface{} { return &HostProfileAppliedEvent{} },

		"HostProfileCompleteConfigSpec": func() interface{} { return &HostProfileCompleteConfigSpec{} },

		"HostProfileConfigInfo": func() interface{} { return &HostProfileConfigInfo{} },

		"HostProfileConfigSpec": func() interface{} { return &HostProfileConfigSpec{} },

		"HostProfileHostBasedConfigSpec": func() interface{} { return &HostProfileHostBasedConfigSpec{} },

		"HostProfileManager": func() interface{} { return &HostProfileManager{} },

		"HostProfileManagerAnswerFileStatus": func() interface{} { return &HostProfileManagerAnswerFileStatus{} },

		"HostProfileManagerConfigTaskList": func() interface{} { return &HostProfileManagerConfigTaskList{} },

		"HostProfileSerializedHostProfileSpec": func() interface{} { return &HostProfileSerializedHostProfileSpec{} },

		"HostProxySwitch": func() interface{} { return &HostProxySwitch{} },

		"HostProxySwitchConfig": func() interface{} { return &HostProxySwitchConfig{} },

		"HostProxySwitchHostLagConfig": func() interface{} { return &HostProxySwitchHostLagConfig{} },

		"HostProxySwitchSpec": func() interface{} { return &HostProxySwitchSpec{} },

		"HostReconnectionFailedEvent": func() interface{} { return &HostReconnectionFailedEvent{} },

		"HostReliableMemoryInfo": func() interface{} { return &HostReliableMemoryInfo{} },

		"HostRemovedEvent": func() interface{} { return &HostRemovedEvent{} },

		"HostReplayUnsupportedReason": func() interface{} { return &HostReplayUnsupportedReason{} },

		"HostResignatureRescanResult": func() interface{} { return &HostResignatureRescanResult{} },

		"HostRuntimeInfo": func() interface{} { return &HostRuntimeInfo{} },

		"HostRuntimeInfoNetStackInstanceRuntimeInfo": func() interface{} { return &HostRuntimeInfoNetStackInstanceRuntimeInfo{} },

		"HostRuntimeInfoNetStackInstanceRuntimeInfoState": func() interface{} { return &HostRuntimeInfoNetStackInstanceRuntimeInfoState{} },

		"HostRuntimeInfoNetworkRuntimeInfo": func() interface{} { return &HostRuntimeInfoNetworkRuntimeInfo{} },

		"HostScsiDisk": func() interface{} { return &HostScsiDisk{} },

		"HostScsiDiskPartition": func() interface{} { return &HostScsiDiskPartition{} },

		"HostScsiTopology": func() interface{} { return &HostScsiTopology{} },

		"HostScsiTopologyInterface": func() interface{} { return &HostScsiTopologyInterface{} },

		"HostScsiTopologyLun": func() interface{} { return &HostScsiTopologyLun{} },

		"HostScsiTopologyTarget": func() interface{} { return &HostScsiTopologyTarget{} },

		"HostSecuritySpec": func() interface{} { return &HostSecuritySpec{} },

		"HostService": func() interface{} { return &HostService{} },

		"HostServiceConfig": func() interface{} { return &HostServiceConfig{} },

		"HostServiceInfo": func() interface{} { return &HostServiceInfo{} },

		"HostServicePolicy": func() interface{} { return &HostServicePolicy{} },

		"HostServiceSourcePackage": func() interface{} { return &HostServiceSourcePackage{} },

		"HostServiceSystem": func() interface{} { return &HostServiceSystem{} },

		"HostServiceTicket": func() interface{} { return &HostServiceTicket{} },

		"HostShortNameInconsistentEvent": func() interface{} { return &HostShortNameInconsistentEvent{} },

		"HostShortNameToIpFailedEvent": func() interface{} { return &HostShortNameToIpFailedEvent{} },

		"HostShutdownEvent": func() interface{} { return &HostShutdownEvent{} },

		"HostSnmpAgentCapability": func() interface{} { return &HostSnmpAgentCapability{} },

		"HostSnmpConfigSpec": func() interface{} { return &HostSnmpConfigSpec{} },

		"HostSnmpDestination": func() interface{} { return &HostSnmpDestination{} },

		"HostSnmpSystem": func() interface{} { return &HostSnmpSystem{} },

		"HostSnmpSystemAgentLimits": func() interface{} { return &HostSnmpSystemAgentLimits{} },

		"HostSriovConfig": func() interface{} { return &HostSriovConfig{} },

		"HostSriovInfo": func() interface{} { return &HostSriovInfo{} },

		"HostSslThumbprintInfo": func() interface{} { return &HostSslThumbprintInfo{} },

		"HostStandbyMode": func() interface{} { return &HostStandbyMode{} },

		"HostStatusChangedEvent": func() interface{} { return &HostStatusChangedEvent{} },

		"HostStorageArrayTypePolicyOption": func() interface{} { return &HostStorageArrayTypePolicyOption{} },

		"HostStorageDeviceInfo": func() interface{} { return &HostStorageDeviceInfo{} },

		"HostStorageElementInfo": func() interface{} { return &HostStorageElementInfo{} },

		"HostStorageOperationalInfo": func() interface{} { return &HostStorageOperationalInfo{} },

		"HostStorageSystem": func() interface{} { return &HostStorageSystem{} },

		"HostSyncFailedEvent": func() interface{} { return &HostSyncFailedEvent{} },

		"HostSystem": func() interface{} { return &HostSystem{} },

		"HostSystemConnectionState": func() interface{} { return &HostSystemConnectionState{} },

		"HostSystemHealthInfo": func() interface{} { return &HostSystemHealthInfo{} },

		"HostSystemIdentificationInfo": func() interface{} { return &HostSystemIdentificationInfo{} },

		"HostSystemIdentificationInfoIdentifier": func() interface{} { return &HostSystemIdentificationInfoIdentifier{} },

		"HostSystemInfo": func() interface{} { return &HostSystemInfo{} },

		"HostSystemPowerState": func() interface{} { return &HostSystemPowerState{} },

		"HostSystemReconnectSpec": func() interface{} { return &HostSystemReconnectSpec{} },

		"HostSystemResourceInfo": func() interface{} { return &HostSystemResourceInfo{} },

		"HostSystemSwapConfiguration": func() interface{} { return &HostSystemSwapConfiguration{} },

		"HostSystemSwapConfigurationDatastoreOption": func() interface{} { return &HostSystemSwapConfigurationDatastoreOption{} },

		"HostSystemSwapConfigurationDisabledOption": func() interface{} { return &HostSystemSwapConfigurationDisabledOption{} },

		"HostSystemSwapConfigurationHostCacheOption": func() interface{} { return &HostSystemSwapConfigurationHostCacheOption{} },

		"HostSystemSwapConfigurationHostLocalSwapOption": func() interface{} { return &HostSystemSwapConfigurationHostLocalSwapOption{} },

		"HostSystemSwapConfigurationSystemSwapOption": func() interface{} { return &HostSystemSwapConfigurationSystemSwapOption{} },

		"HostTargetTransport": func() interface{} { return &HostTargetTransport{} },

		"HostTpmAttestationReport": func() interface{} { return &HostTpmAttestationReport{} },

		"HostTpmBootSecurityOptionEventDetails": func() interface{} { return &HostTpmBootSecurityOptionEventDetails{} },

		"HostTpmCommandEventDetails": func() interface{} { return &HostTpmCommandEventDetails{} },

		"HostTpmDigestInfo": func() interface{} { return &HostTpmDigestInfo{} },

		"HostTpmEventDetails": func() interface{} { return &HostTpmEventDetails{} },

		"HostTpmEventLogEntry": func() interface{} { return &HostTpmEventLogEntry{} },

		"HostTpmOptionEventDetails": func() interface{} { return &HostTpmOptionEventDetails{} },

		"HostTpmSoftwareComponentEventDetails": func() interface{} { return &HostTpmSoftwareComponentEventDetails{} },

		"HostUnresolvedVmfsExtent": func() interface{} { return &HostUnresolvedVmfsExtent{} },

		"HostUnresolvedVmfsExtentUnresolvedReason": func() interface{} { return &HostUnresolvedVmfsExtentUnresolvedReason{} },

		"HostUnresolvedVmfsResignatureSpec": func() interface{} { return &HostUnresolvedVmfsResignatureSpec{} },

		"HostUnresolvedVmfsResolutionResult": func() interface{} { return &HostUnresolvedVmfsResolutionResult{} },

		"HostUnresolvedVmfsResolutionSpec": func() interface{} { return &HostUnresolvedVmfsResolutionSpec{} },

		"HostUnresolvedVmfsResolutionSpecVmfsUuidResolution": func() interface{} { return &HostUnresolvedVmfsResolutionSpecVmfsUuidResolution{} },

		"HostUnresolvedVmfsVolume": func() interface{} { return &HostUnresolvedVmfsVolume{} },

		"HostUnresolvedVmfsVolumeResolveStatus": func() interface{} { return &HostUnresolvedVmfsVolumeResolveStatus{} },

		"HostUpgradeFailedEvent": func() interface{} { return &HostUpgradeFailedEvent{} },

		"HostUserWorldSwapNotEnabledEvent": func() interface{} { return &HostUserWorldSwapNotEnabledEvent{} },

		"HostVFlashManager": func() interface{} { return &HostVFlashManager{} },

		"HostVFlashManagerVFlashCacheConfigInfo": func() interface{} { return &HostVFlashManagerVFlashCacheConfigInfo{} },

		"HostVFlashManagerVFlashCacheConfigInfoVFlashModuleConfigOption": func() interface{} { return &HostVFlashManagerVFlashCacheConfigInfoVFlashModuleConfigOption{} },

		"HostVFlashManagerVFlashCacheConfigSpec": func() interface{} { return &HostVFlashManagerVFlashCacheConfigSpec{} },

		"HostVFlashManagerVFlashConfigInfo": func() interface{} { return &HostVFlashManagerVFlashConfigInfo{} },

		"HostVFlashManagerVFlashResourceConfigInfo": func() interface{} { return &HostVFlashManagerVFlashResourceConfigInfo{} },

		"HostVFlashManagerVFlashResourceConfigSpec": func() interface{} { return &HostVFlashManagerVFlashResourceConfigSpec{} },

		"HostVFlashManagerVFlashResourceRunTimeInfo": func() interface{} { return &HostVFlashManagerVFlashResourceRunTimeInfo{} },

		"HostVFlashResourceConfigurationResult": func() interface{} { return &HostVFlashResourceConfigurationResult{} },

		"HostVMotionCompatibility": func() interface{} { return &HostVMotionCompatibility{} },

		"HostVMotionConfig": func() interface{} { return &HostVMotionConfig{} },

		"HostVMotionInfo": func() interface{} { return &HostVMotionInfo{} },

		"HostVMotionNetConfig": func() interface{} { return &HostVMotionNetConfig{} },

		"HostVMotionSystem": func() interface{} { return &HostVMotionSystem{} },

		"HostVfatVolume": func() interface{} { return &HostVfatVolume{} },

		"HostVffsSpec": func() interface{} { return &HostVffsSpec{} },

		"HostVffsVolume": func() interface{} { return &HostVffsVolume{} },

		"HostVirtualNic": func() interface{} { return &HostVirtualNic{} },

		"HostVirtualNicConfig": func() interface{} { return &HostVirtualNicConfig{} },

		"HostVirtualNicConnection": func() interface{} { return &HostVirtualNicConnection{} },

		"HostVirtualNicManager": func() interface{} { return &HostVirtualNicManager{} },

		"HostVirtualNicManagerInfo": func() interface{} { return &HostVirtualNicManagerInfo{} },

		"HostVirtualNicManagerNicType": func() interface{} { return &HostVirtualNicManagerNicType{} },

		"HostVirtualNicManagerNicTypeSelection": func() interface{} { return &HostVirtualNicManagerNicTypeSelection{} },

		"HostVirtualNicSpec": func() interface{} { return &HostVirtualNicSpec{} },

		"HostVirtualSwitch": func() interface{} { return &HostVirtualSwitch{} },

		"HostVirtualSwitchAutoBridge": func() interface{} { return &HostVirtualSwitchAutoBridge{} },

		"HostVirtualSwitchBeaconConfig": func() interface{} { return &HostVirtualSwitchBeaconConfig{} },

		"HostVirtualSwitchBondBridge": func() interface{} { return &HostVirtualSwitchBondBridge{} },

		"HostVirtualSwitchBridge": func() interface{} { return &HostVirtualSwitchBridge{} },

		"HostVirtualSwitchConfig": func() interface{} { return &HostVirtualSwitchConfig{} },

		"HostVirtualSwitchSimpleBridge": func() interface{} { return &HostVirtualSwitchSimpleBridge{} },

		"HostVirtualSwitchSpec": func() interface{} { return &HostVirtualSwitchSpec{} },

		"HostVmciAccessManagerAccessSpec": func() interface{} { return &HostVmciAccessManagerAccessSpec{} },

		"HostVmciAccessManagerMode": func() interface{} { return &HostVmciAccessManagerMode{} },

		"HostVmfsRescanResult": func() interface{} { return &HostVmfsRescanResult{} },

		"HostVmfsSpec": func() interface{} { return &HostVmfsSpec{} },

		"HostVmfsVolume": func() interface{} { return &HostVmfsVolume{} },

		"HostVnicConnectedToCustomizedDVPortEvent": func() interface{} { return &HostVnicConnectedToCustomizedDVPortEvent{} },

		"HostVsanInternalSystem": func() interface{} { return &HostVsanInternalSystem{} },

		"HostVsanInternalSystemCmmdsQuery": func() interface{} { return &HostVsanInternalSystemCmmdsQuery{} },

		"HostVsanSystem": func() interface{} { return &HostVsanSystem{} },

		"HostWwnChangedEvent": func() interface{} { return &HostWwnChangedEvent{} },

		"HostWwnConflictEvent": func() interface{} { return &HostWwnConflictEvent{} },

		"HotSnapshotMoveNotSupported": func() interface{} { return &HotSnapshotMoveNotSupported{} },

		"HourlyTaskScheduler": func() interface{} { return &HourlyTaskScheduler{} },

		"HttpNfcLease": func() interface{} { return &HttpNfcLease{} },

		"HttpNfcLeaseDatastoreLeaseInfo": func() interface{} { return &HttpNfcLeaseDatastoreLeaseInfo{} },

		"HttpNfcLeaseDeviceUrl": func() interface{} { return &HttpNfcLeaseDeviceUrl{} },

		"HttpNfcLeaseHostInfo": func() interface{} { return &HttpNfcLeaseHostInfo{} },

		"HttpNfcLeaseInfo": func() interface{} { return &HttpNfcLeaseInfo{} },

		"HttpNfcLeaseManifestEntry": func() interface{} { return &HttpNfcLeaseManifestEntry{} },

		"HttpNfcLeaseState": func() interface{} { return &HttpNfcLeaseState{} },

		"IDEDiskNotSupported": func() interface{} { return &IDEDiskNotSupported{} },

		"IORMNotSupportedHostOnDatastore": func() interface{} { return &IORMNotSupportedHostOnDatastore{} },

		"IScsiBootFailureEvent": func() interface{} { return &IScsiBootFailureEvent{} },

		"ImportHostAddFailure": func() interface{} { return &ImportHostAddFailure{} },

		"ImportOperationBulkFault": func() interface{} { return &ImportOperationBulkFault{} },

		"ImportOperationBulkFaultFaultOnImport": func() interface{} { return &ImportOperationBulkFaultFaultOnImport{} },

		"ImportSpec": func() interface{} { return &ImportSpec{} },

		"InUseFeatureManipulationDisallowed": func() interface{} { return &InUseFeatureManipulationDisallowed{} },

		"InaccessibleDatastore": func() interface{} { return &InaccessibleDatastore{} },

		"InaccessibleVFlashSource": func() interface{} { return &InaccessibleVFlashSource{} },

		"IncompatibleDefaultDevice": func() interface{} { return &IncompatibleDefaultDevice{} },

		"IncompatibleHostForFtSecondary": func() interface{} { return &IncompatibleHostForFtSecondary{} },

		"IncompatibleSetting": func() interface{} { return &IncompatibleSetting{} },

		"IncorrectFileType": func() interface{} { return &IncorrectFileType{} },

		"IncorrectHostInformation": func() interface{} { return &IncorrectHostInformation{} },

		"IncorrectHostInformationEvent": func() interface{} { return &IncorrectHostInformationEvent{} },

		"IndependentDiskVMotionNotSupported": func() interface{} { return &IndependentDiskVMotionNotSupported{} },

		"InfoUpgradeEvent": func() interface{} { return &InfoUpgradeEvent{} },

		"InheritablePolicy": func() interface{} { return &InheritablePolicy{} },

		"InsufficientAgentVmsDeployed": func() interface{} { return &InsufficientAgentVmsDeployed{} },

		"InsufficientCpuResourcesFault": func() interface{} { return &InsufficientCpuResourcesFault{} },

		"InsufficientDisks": func() interface{} { return &InsufficientDisks{} },

		"InsufficientFailoverResourcesEvent": func() interface{} { return &InsufficientFailoverResourcesEvent{} },

		"InsufficientFailoverResourcesFault": func() interface{} { return &InsufficientFailoverResourcesFault{} },

		"InsufficientHostCapacityFault": func() interface{} { return &InsufficientHostCapacityFault{} },

		"InsufficientHostCpuCapacityFault": func() interface{} { return &InsufficientHostCpuCapacityFault{} },

		"InsufficientHostMemoryCapacityFault": func() interface{} { return &InsufficientHostMemoryCapacityFault{} },

		"InsufficientMemoryResourcesFault": func() interface{} { return &InsufficientMemoryResourcesFault{} },

		"InsufficientPerCpuCapacity": func() interface{} { return &InsufficientPerCpuCapacity{} },

		"InsufficientResourcesFault": func() interface{} { return &InsufficientResourcesFault{} },

		"InsufficientStandbyCpuResource": func() interface{} { return &InsufficientStandbyCpuResource{} },

		"InsufficientStandbyMemoryResource": func() interface{} { return &InsufficientStandbyMemoryResource{} },

		"InsufficientStandbyResource": func() interface{} { return &InsufficientStandbyResource{} },

		"InsufficientStorageSpace": func() interface{} { return &InsufficientStorageSpace{} },

		"InsufficientVFlashResourcesFault": func() interface{} { return &InsufficientVFlashResourcesFault{} },

		"IntExpression": func() interface{} { return &IntExpression{} },

		"IntOption": func() interface{} { return &IntOption{} },

		"IntPolicy": func() interface{} { return &IntPolicy{} },

		"InternetScsiSnsDiscoveryMethod": func() interface{} { return &InternetScsiSnsDiscoveryMethod{} },

		"InvalidAffinitySettingFault": func() interface{} { return &InvalidAffinitySettingFault{} },

		"InvalidArgument": func() interface{} { return &InvalidArgument{} },

		"InvalidBmcRole": func() interface{} { return &InvalidBmcRole{} },

		"InvalidBundle": func() interface{} { return &InvalidBundle{} },

		"InvalidCAMCertificate": func() interface{} { return &InvalidCAMCertificate{} },

		"InvalidCAMServer": func() interface{} { return &InvalidCAMServer{} },

		"InvalidClientCertificate": func() interface{} { return &InvalidClientCertificate{} },

		"InvalidCollectorVersion": func() interface{} { return &InvalidCollectorVersion{} },

		"InvalidController": func() interface{} { return &InvalidController{} },

		"InvalidDasConfigArgument": func() interface{} { return &InvalidDasConfigArgument{} },

		"InvalidDasConfigArgumentEntryForInvalidArgument": func() interface{} { return &InvalidDasConfigArgumentEntryForInvalidArgument{} },

		"InvalidDasRestartPriorityForFtVm": func() interface{} { return &InvalidDasRestartPriorityForFtVm{} },

		"InvalidDatastore": func() interface{} { return &InvalidDatastore{} },

		"InvalidDatastorePath": func() interface{} { return &InvalidDatastorePath{} },

		"InvalidDatastoreState": func() interface{} { return &InvalidDatastoreState{} },

		"InvalidDeviceBacking": func() interface{} { return &InvalidDeviceBacking{} },

		"InvalidDeviceOperation": func() interface{} { return &InvalidDeviceOperation{} },

		"InvalidDeviceSpec": func() interface{} { return &InvalidDeviceSpec{} },

		"InvalidDiskFormat": func() interface{} { return &InvalidDiskFormat{} },

		"InvalidDrsBehaviorForFtVm": func() interface{} { return &InvalidDrsBehaviorForFtVm{} },

		"InvalidEditionEvent": func() interface{} { return &InvalidEditionEvent{} },

		"InvalidEditionLicense": func() interface{} { return &InvalidEditionLicense{} },

		"InvalidEvent": func() interface{} { return &InvalidEvent{} },

		"InvalidFolder": func() interface{} { return &InvalidFolder{} },

		"InvalidFormat": func() interface{} { return &InvalidFormat{} },

		"InvalidGuestLogin": func() interface{} { return &InvalidGuestLogin{} },

		"InvalidHostConnectionState": func() interface{} { return &InvalidHostConnectionState{} },

		"InvalidHostName": func() interface{} { return &InvalidHostName{} },

		"InvalidHostState": func() interface{} { return &InvalidHostState{} },

		"InvalidIndexArgument": func() interface{} { return &InvalidIndexArgument{} },

		"InvalidIpfixConfig": func() interface{} { return &InvalidIpfixConfig{} },

		"InvalidIpmiLoginInfo": func() interface{} { return &InvalidIpmiLoginInfo{} },

		"InvalidIpmiMacAddress": func() interface{} { return &InvalidIpmiMacAddress{} },

		"InvalidLicense": func() interface{} { return &InvalidLicense{} },

		"InvalidLocale": func() interface{} { return &InvalidLocale{} },

		"InvalidLogin": func() interface{} { return &InvalidLogin{} },

		"InvalidName": func() interface{} { return &InvalidName{} },

		"InvalidNasCredentials": func() interface{} { return &InvalidNasCredentials{} },

		"InvalidNetworkInType": func() interface{} { return &InvalidNetworkInType{} },

		"InvalidNetworkResource": func() interface{} { return &InvalidNetworkResource{} },

		"InvalidOperationOnSecondaryVm": func() interface{} { return &InvalidOperationOnSecondaryVm{} },

		"InvalidPowerState": func() interface{} { return &InvalidPowerState{} },

		"InvalidPrivilege": func() interface{} { return &InvalidPrivilege{} },

		"InvalidProfileReferenceHost": func() interface{} { return &InvalidProfileReferenceHost{} },

		"InvalidProfileReferenceHostReason": func() interface{} { return &InvalidProfileReferenceHostReason{} },

		"InvalidProperty": func() interface{} { return &InvalidProperty{} },

		"InvalidPropertyType": func() interface{} { return &InvalidPropertyType{} },

		"InvalidPropertyValue": func() interface{} { return &InvalidPropertyValue{} },

		"InvalidRequest": func() interface{} { return &InvalidRequest{} },

		"InvalidResourcePoolStructureFault": func() interface{} { return &InvalidResourcePoolStructureFault{} },

		"InvalidSnapshotFormat": func() interface{} { return &InvalidSnapshotFormat{} },

		"InvalidState": func() interface{} { return &InvalidState{} },

		"InvalidType": func() interface{} { return &InvalidType{} },

		"InvalidVmConfig": func() interface{} { return &InvalidVmConfig{} },

		"InventoryDescription": func() interface{} { return &InventoryDescription{} },

		"InventoryHasStandardAloneHosts": func() interface{} { return &InventoryHasStandardAloneHosts{} },

		"InventoryView": func() interface{} { return &InventoryView{} },

		"IpAddress": func() interface{} { return &IpAddress{} },

		"IpAddressProfile": func() interface{} { return &IpAddressProfile{} },

		"IpHostnameGeneratorError": func() interface{} { return &IpHostnameGeneratorError{} },

		"IpPool": func() interface{} { return &IpPool{} },

		"IpPoolAssociation": func() interface{} { return &IpPoolAssociation{} },

		"IpPoolIpPoolConfigInfo": func() interface{} { return &IpPoolIpPoolConfigInfo{} },

		"IpPoolManager": func() interface{} { return &IpPoolManager{} },

		"IpPoolManagerIpAllocation": func() interface{} { return &IpPoolManagerIpAllocation{} },

		"IpRange": func() interface{} { return &IpRange{} },

		"IpRouteProfile": func() interface{} { return &IpRouteProfile{} },

		"IscsiDependencyEntity": func() interface{} { return &IscsiDependencyEntity{} },

		"IscsiFault": func() interface{} { return &IscsiFault{} },

		"IscsiFaultInvalidVnic": func() interface{} { return &IscsiFaultInvalidVnic{} },

		"IscsiFaultPnicInUse": func() interface{} { return &IscsiFaultPnicInUse{} },

		"IscsiFaultVnicAlreadyBound": func() interface{} { return &IscsiFaultVnicAlreadyBound{} },

		"IscsiFaultVnicHasActivePaths": func() interface{} { return &IscsiFaultVnicHasActivePaths{} },

		"IscsiFaultVnicHasMultipleUplinks": func() interface{} { return &IscsiFaultVnicHasMultipleUplinks{} },

		"IscsiFaultVnicHasNoUplinks": func() interface{} { return &IscsiFaultVnicHasNoUplinks{} },

		"IscsiFaultVnicHasWrongUplink": func() interface{} { return &IscsiFaultVnicHasWrongUplink{} },

		"IscsiFaultVnicInUse": func() interface{} { return &IscsiFaultVnicInUse{} },

		"IscsiFaultVnicIsLastPath": func() interface{} { return &IscsiFaultVnicIsLastPath{} },

		"IscsiFaultVnicNotBound": func() interface{} { return &IscsiFaultVnicNotBound{} },

		"IscsiFaultVnicNotFound": func() interface{} { return &IscsiFaultVnicNotFound{} },

		"IscsiManager": func() interface{} { return &IscsiManager{} },

		"IscsiMigrationDependency": func() interface{} { return &IscsiMigrationDependency{} },

		"IscsiPortInfo": func() interface{} { return &IscsiPortInfo{} },

		"IscsiPortInfoPathStatus": func() interface{} { return &IscsiPortInfoPathStatus{} },

		"IscsiStatus": func() interface{} { return &IscsiStatus{} },

		"IsoImageFileInfo": func() interface{} { return &IsoImageFileInfo{} },

		"IsoImageFileQuery": func() interface{} { return &IsoImageFileQuery{} },

		"KernelModuleInfo": func() interface{} { return &KernelModuleInfo{} },

		"KernelModuleSectionInfo": func() interface{} { return &KernelModuleSectionInfo{} },

		"KeyAnyValue": func() interface{} { return &KeyAnyValue{} },

		"KeyValue": func() interface{} { return &KeyValue{} },

		"LargeRDMConversionNotSupported": func() interface{} { return &LargeRDMConversionNotSupported{} },

		"LargeRDMNotSupportedOnDatastore": func() interface{} { return &LargeRDMNotSupportedOnDatastore{} },

		"LatencySensitivity": func() interface{} { return &LatencySensitivity{} },

		"LatencySensitivitySensitivityLevel": func() interface{} { return &LatencySensitivitySensitivityLevel{} },

		"LegacyNetworkInterfaceInUse": func() interface{} { return &LegacyNetworkInterfaceInUse{} },

		"LicenseAssignmentFailed": func() interface{} { return &LicenseAssignmentFailed{} },

		"LicenseAssignmentFailedReason": func() interface{} { return &LicenseAssignmentFailedReason{} },

		"LicenseAssignmentManager": func() interface{} { return &LicenseAssignmentManager{} },

		"LicenseAssignmentManagerLicenseAssignment": func() interface{} { return &LicenseAssignmentManagerLicenseAssignment{} },

		"LicenseAvailabilityInfo": func() interface{} { return &LicenseAvailabilityInfo{} },

		"LicenseDiagnostics": func() interface{} { return &LicenseDiagnostics{} },

		"LicenseDowngradeDisallowed": func() interface{} { return &LicenseDowngradeDisallowed{} },

		"LicenseEntityNotFound": func() interface{} { return &LicenseEntityNotFound{} },

		"LicenseEvent": func() interface{} { return &LicenseEvent{} },

		"LicenseExpired": func() interface{} { return &LicenseExpired{} },

		"LicenseExpiredEvent": func() interface{} { return &LicenseExpiredEvent{} },

		"LicenseFeatureInfo": func() interface{} { return &LicenseFeatureInfo{} },

		"LicenseFeatureInfoSourceRestriction": func() interface{} { return &LicenseFeatureInfoSourceRestriction{} },

		"LicenseFeatureInfoState": func() interface{} { return &LicenseFeatureInfoState{} },

		"LicenseFeatureInfoUnit": func() interface{} { return &LicenseFeatureInfoUnit{} },

		"LicenseKeyEntityMismatch": func() interface{} { return &LicenseKeyEntityMismatch{} },

		"LicenseManager": func() interface{} { return &LicenseManager{} },

		"LicenseManagerEvaluationInfo": func() interface{} { return &LicenseManagerEvaluationInfo{} },

		"LicenseManagerLicenseInfo": func() interface{} { return &LicenseManagerLicenseInfo{} },

		"LicenseManagerLicenseKey": func() interface{} { return &LicenseManagerLicenseKey{} },

		"LicenseManagerState": func() interface{} { return &LicenseManagerState{} },

		"LicenseNonComplianceEvent": func() interface{} { return &LicenseNonComplianceEvent{} },

		"LicenseReservationInfo": func() interface{} { return &LicenseReservationInfo{} },

		"LicenseReservationInfoState": func() interface{} { return &LicenseReservationInfoState{} },

		"LicenseRestricted": func() interface{} { return &LicenseRestricted{} },

		"LicenseRestrictedEvent": func() interface{} { return &LicenseRestrictedEvent{} },

		"LicenseServerAvailableEvent": func() interface{} { return &LicenseServerAvailableEvent{} },

		"LicenseServerSource": func() interface{} { return &LicenseServerSource{} },

		"LicenseServerUnavailable": func() interface{} { return &LicenseServerUnavailable{} },

		"LicenseServerUnavailableEvent": func() interface{} { return &LicenseServerUnavailableEvent{} },

		"LicenseSource": func() interface{} { return &LicenseSource{} },

		"LicenseSourceUnavailable": func() interface{} { return &LicenseSourceUnavailable{} },

		"LicenseUsageInfo": func() interface{} { return &LicenseUsageInfo{} },

		"LimitExceeded": func() interface{} { return &LimitExceeded{} },

		"LinkDiscoveryProtocolConfig": func() interface{} { return &LinkDiscoveryProtocolConfig{} },

		"LinkDiscoveryProtocolConfigOperationType": func() interface{} { return &LinkDiscoveryProtocolConfigOperationType{} },

		"LinkDiscoveryProtocolConfigProtocolType": func() interface{} { return &LinkDiscoveryProtocolConfigProtocolType{} },

		"LinkLayerDiscoveryProtocolInfo": func() interface{} { return &LinkLayerDiscoveryProtocolInfo{} },

		"LinkProfile": func() interface{} { return &LinkProfile{} },

		"LinuxVolumeNotClean": func() interface{} { return &LinuxVolumeNotClean{} },

		"ListView": func() interface{} { return &ListView{} },

		"LocalDatastoreCreatedEvent": func() interface{} { return &LocalDatastoreCreatedEvent{} },

		"LocalDatastoreInfo": func() interface{} { return &LocalDatastoreInfo{} },

		"LocalLicenseSource": func() interface{} { return &LocalLicenseSource{} },

		"LocalTSMEnabledEvent": func() interface{} { return &LocalTSMEnabledEvent{} },

		"LocalizableMessage": func() interface{} { return &LocalizableMessage{} },

		"LocalizationManager": func() interface{} { return &LocalizationManager{} },

		"LocalizationManagerMessageCatalog": func() interface{} { return &LocalizationManagerMessageCatalog{} },

		"LocalizedMethodFault": func() interface{} { return &LocalizedMethodFault{} },

		"LockerMisconfiguredEvent": func() interface{} { return &LockerMisconfiguredEvent{} },

		"LockerReconfiguredEvent": func() interface{} { return &LockerReconfiguredEvent{} },

		"LogBundlingFailed": func() interface{} { return &LogBundlingFailed{} },

		"LongOption": func() interface{} { return &LongOption{} },

		"LongPolicy": func() interface{} { return &LongPolicy{} },

		"MacAddress": func() interface{} { return &MacAddress{} },

		"MacRange": func() interface{} { return &MacRange{} },

		"MaintenanceModeFileMove": func() interface{} { return &MaintenanceModeFileMove{} },

		"ManagedByInfo": func() interface{} { return &ManagedByInfo{} },

		"ManagedEntity": func() interface{} { return &ManagedEntity{} },

		"ManagedEntityEventArgument": func() interface{} { return &ManagedEntityEventArgument{} },

		"ManagedEntityStatus": func() interface{} { return &ManagedEntityStatus{} },

		"ManagedObjectNotFound": func() interface{} { return &ManagedObjectNotFound{} },

		"ManagedObjectReference": func() interface{} { return &ManagedObjectReference{} },

		"ManagedObjectView": func() interface{} { return &ManagedObjectView{} },

		"MemoryHotPlugNotSupported": func() interface{} { return &MemoryHotPlugNotSupported{} },

		"MemorySizeNotRecommended": func() interface{} { return &MemorySizeNotRecommended{} },

		"MemorySizeNotSupported": func() interface{} { return &MemorySizeNotSupported{} },

		"MemorySizeNotSupportedByDatastore": func() interface{} { return &MemorySizeNotSupportedByDatastore{} },

		"MemorySnapshotOnIndependentDisk": func() interface{} { return &MemorySnapshotOnIndependentDisk{} },

		"MethodAction": func() interface{} { return &MethodAction{} },

		"MethodActionArgument": func() interface{} { return &MethodActionArgument{} },

		"MethodAlreadyDisabledFault": func() interface{} { return &MethodAlreadyDisabledFault{} },

		"MethodDescription": func() interface{} { return &MethodDescription{} },

		"MethodDisabled": func() interface{} { return &MethodDisabled{} },

		"MethodFault": func() interface{} { return &MethodFault{} },

		"MethodNotFound": func() interface{} { return &MethodNotFound{} },

		"MetricAlarmExpression": func() interface{} { return &MetricAlarmExpression{} },

		"MetricAlarmOperator": func() interface{} { return &MetricAlarmOperator{} },

		"MigrationDisabled": func() interface{} { return &MigrationDisabled{} },

		"MigrationErrorEvent": func() interface{} { return &MigrationErrorEvent{} },

		"MigrationEvent": func() interface{} { return &MigrationEvent{} },

		"MigrationFault": func() interface{} { return &MigrationFault{} },

		"MigrationFeatureNotSupported": func() interface{} { return &MigrationFeatureNotSupported{} },

		"MigrationHostErrorEvent": func() interface{} { return &MigrationHostErrorEvent{} },

		"MigrationHostWarningEvent": func() interface{} { return &MigrationHostWarningEvent{} },

		"MigrationNotReady": func() interface{} { return &MigrationNotReady{} },

		"MigrationResourceErrorEvent": func() interface{} { return &MigrationResourceErrorEvent{} },

		"MigrationResourceWarningEvent": func() interface{} { return &MigrationResourceWarningEvent{} },

		"MigrationWarningEvent": func() interface{} { return &MigrationWarningEvent{} },

		"MismatchedBundle": func() interface{} { return &MismatchedBundle{} },

		"MismatchedNetworkPolicies": func() interface{} { return &MismatchedNetworkPolicies{} },

		"MismatchedVMotionNetworkNames": func() interface{} { return &MismatchedVMotionNetworkNames{} },

		"MissingBmcSupport": func() interface{} { return &MissingBmcSupport{} },

		"MissingController": func() interface{} { return &MissingController{} },

		"MissingIpPool": func() interface{} { return &MissingIpPool{} },

		"MissingLinuxCustResources": func() interface{} { return &MissingLinuxCustResources{} },

		"MissingNetworkIpConfig": func() interface{} { return &MissingNetworkIpConfig{} },

		"MissingObject": func() interface{} { return &MissingObject{} },

		"MissingPowerOffConfiguration": func() interface{} { return &MissingPowerOffConfiguration{} },

		"MissingPowerOnConfiguration": func() interface{} { return &MissingPowerOnConfiguration{} },

		"MissingProperty": func() interface{} { return &MissingProperty{} },

		"MissingWindowsCustResources": func() interface{} { return &MissingWindowsCustResources{} },

		"MksConnectionLimitReached": func() interface{} { return &MksConnectionLimitReached{} },

		"ModeInfo": func() interface{} { return &ModeInfo{} },

		"MonthlyByDayTaskScheduler": func() interface{} { return &MonthlyByDayTaskScheduler{} },

		"MonthlyByWeekdayTaskScheduler": func() interface{} { return &MonthlyByWeekdayTaskScheduler{} },

		"MonthlyTaskScheduler": func() interface{} { return &MonthlyTaskScheduler{} },

		"MountError": func() interface{} { return &MountError{} },

		"MtuMatchEvent": func() interface{} { return &MtuMatchEvent{} },

		"MtuMismatchEvent": func() interface{} { return &MtuMismatchEvent{} },

		"MultipathState": func() interface{} { return &MultipathState{} },

		"MultipleCertificatesVerifyFault": func() interface{} { return &MultipleCertificatesVerifyFault{} },

		"MultipleCertificatesVerifyFaultThumbprintData": func() interface{} { return &MultipleCertificatesVerifyFaultThumbprintData{} },

		"MultipleSnapshotsNotSupported": func() interface{} { return &MultipleSnapshotsNotSupported{} },

		"NASDatastoreCreatedEvent": func() interface{} { return &NASDatastoreCreatedEvent{} },

		"NamePasswordAuthentication": func() interface{} { return &NamePasswordAuthentication{} },

		"NamespaceFull": func() interface{} { return &NamespaceFull{} },

		"NamespaceLimitReached": func() interface{} { return &NamespaceLimitReached{} },

		"NamespaceWriteProtected": func() interface{} { return &NamespaceWriteProtected{} },

		"NasConfigFault": func() interface{} { return &NasConfigFault{} },

		"NasConnectionLimitReached": func() interface{} { return &NasConnectionLimitReached{} },

		"NasDatastoreInfo": func() interface{} { return &NasDatastoreInfo{} },

		"NasSessionCredentialConflict": func() interface{} { return &NasSessionCredentialConflict{} },

		"NasStorageProfile": func() interface{} { return &NasStorageProfile{} },

		"NasVolumeNotMounted": func() interface{} { return &NasVolumeNotMounted{} },

		"NegatableExpression": func() interface{} { return &NegatableExpression{} },

		"NetBIOSConfigInfo": func() interface{} { return &NetBIOSConfigInfo{} },

		"NetBIOSConfigInfoMode": func() interface{} { return &NetBIOSConfigInfoMode{} },

		"NetDhcpConfigInfo": func() interface{} { return &NetDhcpConfigInfo{} },

		"NetDhcpConfigInfoDhcpOptions": func() interface{} { return &NetDhcpConfigInfoDhcpOptions{} },

		"NetDhcpConfigSpec": func() interface{} { return &NetDhcpConfigSpec{} },

		"NetDhcpConfigSpecDhcpOptionsSpec": func() interface{} { return &NetDhcpConfigSpecDhcpOptionsSpec{} },

		"NetDnsConfigInfo": func() interface{} { return &NetDnsConfigInfo{} },

		"NetDnsConfigSpec": func() interface{} { return &NetDnsConfigSpec{} },

		"NetIpConfigInfo": func() interface{} { return &NetIpConfigInfo{} },

		"NetIpConfigInfoIpAddress": func() interface{} { return &NetIpConfigInfoIpAddress{} },

		"NetIpConfigInfoIpAddressOrigin": func() interface{} { return &NetIpConfigInfoIpAddressOrigin{} },

		"NetIpConfigInfoIpAddressStatus": func() interface{} { return &NetIpConfigInfoIpAddressStatus{} },

		"NetIpConfigSpec": func() interface{} { return &NetIpConfigSpec{} },

		"NetIpConfigSpecIpAddressSpec": func() interface{} { return &NetIpConfigSpecIpAddressSpec{} },

		"NetIpRouteConfigInfo": func() interface{} { return &NetIpRouteConfigInfo{} },

		"NetIpRouteConfigInfoGateway": func() interface{} { return &NetIpRouteConfigInfoGateway{} },

		"NetIpRouteConfigInfoIpRoute": func() interface{} { return &NetIpRouteConfigInfoIpRoute{} },

		"NetIpRouteConfigSpec": func() interface{} { return &NetIpRouteConfigSpec{} },

		"NetIpRouteConfigSpecGatewaySpec": func() interface{} { return &NetIpRouteConfigSpecGatewaySpec{} },

		"NetIpRouteConfigSpecIpRouteSpec": func() interface{} { return &NetIpRouteConfigSpecIpRouteSpec{} },

		"NetIpStackInfo": func() interface{} { return &NetIpStackInfo{} },

		"NetIpStackInfoDefaultRouter": func() interface{} { return &NetIpStackInfoDefaultRouter{} },

		"NetIpStackInfoEntryType": func() interface{} { return &NetIpStackInfoEntryType{} },

		"NetIpStackInfoNetToMedia": func() interface{} { return &NetIpStackInfoNetToMedia{} },

		"NetIpStackInfoPreference": func() interface{} { return &NetIpStackInfoPreference{} },

		"NetStackInstanceProfile": func() interface{} { return &NetStackInstanceProfile{} },

		"Network": func() interface{} { return &Network{} },

		"NetworkCopyFault": func() interface{} { return &NetworkCopyFault{} },

		"NetworkDisruptedAndConfigRolledBack": func() interface{} { return &NetworkDisruptedAndConfigRolledBack{} },

		"NetworkEventArgument": func() interface{} { return &NetworkEventArgument{} },

		"NetworkInaccessible": func() interface{} { return &NetworkInaccessible{} },

		"NetworkPolicyProfile": func() interface{} { return &NetworkPolicyProfile{} },

		"NetworkProfile": func() interface{} { return &NetworkProfile{} },

		"NetworkProfileDnsConfigProfile": func() interface{} { return &NetworkProfileDnsConfigProfile{} },

		"NetworkRollbackEvent": func() interface{} { return &NetworkRollbackEvent{} },

		"NetworkSummary": func() interface{} { return &NetworkSummary{} },

		"NetworksMayNotBeTheSame": func() interface{} { return &NetworksMayNotBeTheSame{} },

		"NicSettingMismatch": func() interface{} { return &NicSettingMismatch{} },

		"NoAccessUserEvent": func() interface{} { return &NoAccessUserEvent{} },

		"NoActiveHostInCluster": func() interface{} { return &NoActiveHostInCluster{} },

		"NoAvailableIp": func() interface{} { return &NoAvailableIp{} },

		"NoClientCertificate": func() interface{} { return &NoClientCertificate{} },

		"NoCompatibleDatastore": func() interface{} { return &NoCompatibleDatastore{} },

		"NoCompatibleHardAffinityHost": func() interface{} { return &NoCompatibleHardAffinityHost{} },

		"NoCompatibleHost": func() interface{} { return &NoCompatibleHost{} },

		"NoCompatibleHostWithAccessToDevice": func() interface{} { return &NoCompatibleHostWithAccessToDevice{} },

		"NoCompatibleSoftAffinityHost": func() interface{} { return &NoCompatibleSoftAffinityHost{} },

		"NoConnectedDatastore": func() interface{} { return &NoConnectedDatastore{} },

		"NoDatastoresConfiguredEvent": func() interface{} { return &NoDatastoresConfiguredEvent{} },

		"NoDiskFound": func() interface{} { return &NoDiskFound{} },

		"NoDiskSpace": func() interface{} { return &NoDiskSpace{} },

		"NoDisksToCustomize": func() interface{} { return &NoDisksToCustomize{} },

		"NoGateway": func() interface{} { return &NoGateway{} },

		"NoGuestHeartbeat": func() interface{} { return &NoGuestHeartbeat{} },

		"NoHost": func() interface{} { return &NoHost{} },

		"NoHostSuitableForFtSecondary": func() interface{} { return &NoHostSuitableForFtSecondary{} },

		"NoLicenseEvent": func() interface{} { return &NoLicenseEvent{} },

		"NoLicenseServerConfigured": func() interface{} { return &NoLicenseServerConfigured{} },

		"NoMaintenanceModeDrsRecommendationForVM": func() interface{} { return &NoMaintenanceModeDrsRecommendationForVM{} },

		"NoPeerHostFound": func() interface{} { return &NoPeerHostFound{} },

		"NoPermission": func() interface{} { return &NoPermission{} },

		"NoPermissionOnAD": func() interface{} { return &NoPermissionOnAD{} },

		"NoPermissionOnHost": func() interface{} { return &NoPermissionOnHost{} },

		"NoPermissionOnNasVolume": func() interface{} { return &NoPermissionOnNasVolume{} },

		"NoSubjectName": func() interface{} { return &NoSubjectName{} },

		"NoVcManagedIpConfigured": func() interface{} { return &NoVcManagedIpConfigured{} },

		"NoVirtualNic": func() interface{} { return &NoVirtualNic{} },

		"NoVmInVApp": func() interface{} { return &NoVmInVApp{} },

		"NonADUserRequired": func() interface{} { return &NonADUserRequired{} },

		"NonHomeRDMVMotionNotSupported": func() interface{} { return &NonHomeRDMVMotionNotSupported{} },

		"NonPersistentDisksNotSupported": func() interface{} { return &NonPersistentDisksNotSupported{} },

		"NonVIWorkloadDetectedOnDatastoreEvent": func() interface{} { return &NonVIWorkloadDetectedOnDatastoreEvent{} },

		"NonVmwareOuiMacNotSupportedHost": func() interface{} { return &NonVmwareOuiMacNotSupportedHost{} },

		"NotADirectory": func() interface{} { return &NotADirectory{} },

		"NotAFile": func() interface{} { return &NotAFile{} },

		"NotAuthenticated": func() interface{} { return &NotAuthenticated{} },

		"NotEnoughCpus": func() interface{} { return &NotEnoughCpus{} },

		"NotEnoughLicenses": func() interface{} { return &NotEnoughLicenses{} },

		"NotEnoughLogicalCpus": func() interface{} { return &NotEnoughLogicalCpus{} },

		"NotEnoughResourcesToStartVmEvent": func() interface{} { return &NotEnoughResourcesToStartVmEvent{} },

		"NotFound": func() interface{} { return &NotFound{} },

		"NotImplemented": func() interface{} { return &NotImplemented{} },

		"NotSupported": func() interface{} { return &NotSupported{} },

		"NotSupportedDeviceForFT": func() interface{} { return &NotSupportedDeviceForFT{} },

		"NotSupportedDeviceForFTDeviceType": func() interface{} { return &NotSupportedDeviceForFTDeviceType{} },

		"NotSupportedHost": func() interface{} { return &NotSupportedHost{} },

		"NotSupportedHostForVFlash": func() interface{} { return &NotSupportedHostForVFlash{} },

		"NotSupportedHostForVsan": func() interface{} { return &NotSupportedHostForVsan{} },

		"NotSupportedHostInCluster": func() interface{} { return &NotSupportedHostInCluster{} },

		"NotSupportedHostInDvs": func() interface{} { return &NotSupportedHostInDvs{} },

		"NotSupportedHostInHACluster": func() interface{} { return &NotSupportedHostInHACluster{} },

		"NotUserConfigurableProperty": func() interface{} { return &NotUserConfigurableProperty{} },

		"NumPortsProfile": func() interface{} { return &NumPortsProfile{} },

		"NumVirtualCoresPerSocketNotSupported": func() interface{} { return &NumVirtualCoresPerSocketNotSupported{} },

		"NumVirtualCpusExceedsLimit": func() interface{} { return &NumVirtualCpusExceedsLimit{} },

		"NumVirtualCpusIncompatible": func() interface{} { return &NumVirtualCpusIncompatible{} },

		"NumVirtualCpusIncompatibleReason": func() interface{} { return &NumVirtualCpusIncompatibleReason{} },

		"NumVirtualCpusNotSupported": func() interface{} { return &NumVirtualCpusNotSupported{} },

		"NumericRange": func() interface{} { return &NumericRange{} },

		"ObjectContent": func() interface{} { return &ObjectContent{} },

		"ObjectSpec": func() interface{} { return &ObjectSpec{} },

		"ObjectUpdate": func() interface{} { return &ObjectUpdate{} },

		"ObjectUpdateKind": func() interface{} { return &ObjectUpdateKind{} },

		"OnceTaskScheduler": func() interface{} { return &OnceTaskScheduler{} },

		"OpaqueNetwork": func() interface{} { return &OpaqueNetwork{} },

		"OpaqueNetworkSummary": func() interface{} { return &OpaqueNetworkSummary{} },

		"OpaqueNetworkTargetInfo": func() interface{} { return &OpaqueNetworkTargetInfo{} },

		"OperationDisabledByGuest": func() interface{} { return &OperationDisabledByGuest{} },

		"OperationDisallowedOnHost": func() interface{} { return &OperationDisallowedOnHost{} },

		"OperationNotSupportedByGuest": func() interface{} { return &OperationNotSupportedByGuest{} },

		"OptionDef": func() interface{} { return &OptionDef{} },

		"OptionManager": func() interface{} { return &OptionManager{} },

		"OptionProfile": func() interface{} { return &OptionProfile{} },

		"OptionType": func() interface{} { return &OptionType{} },

		"OptionValue": func() interface{} { return &OptionValue{} },

		"OrAlarmExpression": func() interface{} { return &OrAlarmExpression{} },

		"OutOfBounds": func() interface{} { return &OutOfBounds{} },

		"OutOfSyncDvsHost": func() interface{} { return &OutOfSyncDvsHost{} },

		"OvfAttribute": func() interface{} { return &OvfAttribute{} },

		"OvfConnectedDevice": func() interface{} { return &OvfConnectedDevice{} },

		"OvfConnectedDeviceFloppy": func() interface{} { return &OvfConnectedDeviceFloppy{} },

		"OvfConnectedDeviceIso": func() interface{} { return &OvfConnectedDeviceIso{} },

		"OvfConstraint": func() interface{} { return &OvfConstraint{} },

		"OvfConsumerCallbackFault": func() interface{} { return &OvfConsumerCallbackFault{} },

		"OvfConsumerCommunicationError": func() interface{} { return &OvfConsumerCommunicationError{} },

		"OvfConsumerFault": func() interface{} { return &OvfConsumerFault{} },

		"OvfConsumerInvalidSection": func() interface{} { return &OvfConsumerInvalidSection{} },

		"OvfConsumerOstNode": func() interface{} { return &OvfConsumerOstNode{} },

		"OvfConsumerOstNodeType": func() interface{} { return &OvfConsumerOstNodeType{} },

		"OvfConsumerOvfSection": func() interface{} { return &OvfConsumerOvfSection{} },

		"OvfConsumerPowerOnFault": func() interface{} { return &OvfConsumerPowerOnFault{} },

		"OvfConsumerUndeclaredSection": func() interface{} { return &OvfConsumerUndeclaredSection{} },

		"OvfConsumerUndefinedPrefix": func() interface{} { return &OvfConsumerUndefinedPrefix{} },

		"OvfConsumerValidationFault": func() interface{} { return &OvfConsumerValidationFault{} },

		"OvfCpuCompatibility": func() interface{} { return &OvfCpuCompatibility{} },

		"OvfCpuCompatibilityCheckNotSupported": func() interface{} { return &OvfCpuCompatibilityCheckNotSupported{} },

		"OvfCreateDescriptorParams": func() interface{} { return &OvfCreateDescriptorParams{} },

		"OvfCreateDescriptorResult": func() interface{} { return &OvfCreateDescriptorResult{} },

		"OvfCreateImportSpecParams": func() interface{} { return &OvfCreateImportSpecParams{} },

		"OvfCreateImportSpecParamsDiskProvisioningType": func() interface{} { return &OvfCreateImportSpecParamsDiskProvisioningType{} },

		"OvfCreateImportSpecResult": func() interface{} { return &OvfCreateImportSpecResult{} },

		"OvfDeploymentOption": func() interface{} { return &OvfDeploymentOption{} },

		"OvfDiskMappingNotFound": func() interface{} { return &OvfDiskMappingNotFound{} },

		"OvfDiskOrderConstraint": func() interface{} { return &OvfDiskOrderConstraint{} },

		"OvfDuplicateElement": func() interface{} { return &OvfDuplicateElement{} },

		"OvfDuplicatedElementBoundary": func() interface{} { return &OvfDuplicatedElementBoundary{} },

		"OvfDuplicatedPropertyIdExport": func() interface{} { return &OvfDuplicatedPropertyIdExport{} },

		"OvfDuplicatedPropertyIdImport": func() interface{} { return &OvfDuplicatedPropertyIdImport{} },

		"OvfElement": func() interface{} { return &OvfElement{} },

		"OvfElementInvalidValue": func() interface{} { return &OvfElementInvalidValue{} },

		"OvfExport": func() interface{} { return &OvfExport{} },

		"OvfExportFailed": func() interface{} { return &OvfExportFailed{} },

		"OvfFault": func() interface{} { return &OvfFault{} },

		"OvfFile": func() interface{} { return &OvfFile{} },

		"OvfFileItem": func() interface{} { return &OvfFileItem{} },

		"OvfHardwareCheck": func() interface{} { return &OvfHardwareCheck{} },

		"OvfHardwareExport": func() interface{} { return &OvfHardwareExport{} },

		"OvfHostResourceConstraint": func() interface{} { return &OvfHostResourceConstraint{} },

		"OvfHostValueNotParsed": func() interface{} { return &OvfHostValueNotParsed{} },

		"OvfImport": func() interface{} { return &OvfImport{} },

		"OvfImportFailed": func() interface{} { return &OvfImportFailed{} },

		"OvfInternalError": func() interface{} { return &OvfInternalError{} },

		"OvfInvalidPackage": func() interface{} { return &OvfInvalidPackage{} },

		"OvfInvalidValue": func() interface{} { return &OvfInvalidValue{} },

		"OvfInvalidValueConfiguration": func() interface{} { return &OvfInvalidValueConfiguration{} },

		"OvfInvalidValueEmpty": func() interface{} { return &OvfInvalidValueEmpty{} },

		"OvfInvalidValueFormatMalformed": func() interface{} { return &OvfInvalidValueFormatMalformed{} },

		"OvfInvalidValueReference": func() interface{} { return &OvfInvalidValueReference{} },

		"OvfInvalidVmName": func() interface{} { return &OvfInvalidVmName{} },

		"OvfManager": func() interface{} { return &OvfManager{} },

		"OvfManagerCommonParams": func() interface{} { return &OvfManagerCommonParams{} },

		"OvfMappedOsId": func() interface{} { return &OvfMappedOsId{} },

		"OvfMissingAttribute": func() interface{} { return &OvfMissingAttribute{} },

		"OvfMissingElement": func() interface{} { return &OvfMissingElement{} },

		"OvfMissingElementNormalBoundary": func() interface{} { return &OvfMissingElementNormalBoundary{} },

		"OvfMissingHardware": func() interface{} { return &OvfMissingHardware{} },

		"OvfNetworkInfo": func() interface{} { return &OvfNetworkInfo{} },

		"OvfNetworkMapping": func() interface{} { return &OvfNetworkMapping{} },

		"OvfNetworkMappingNotSupported": func() interface{} { return &OvfNetworkMappingNotSupported{} },

		"OvfNoHostNic": func() interface{} { return &OvfNoHostNic{} },

		"OvfNoSpaceOnController": func() interface{} { return &OvfNoSpaceOnController{} },

		"OvfNoSupportedHardwareFamily": func() interface{} { return &OvfNoSupportedHardwareFamily{} },

		"OvfOptionInfo": func() interface{} { return &OvfOptionInfo{} },

		"OvfParseDescriptorParams": func() interface{} { return &OvfParseDescriptorParams{} },

		"OvfParseDescriptorResult": func() interface{} { return &OvfParseDescriptorResult{} },

		"OvfProperty": func() interface{} { return &OvfProperty{} },

		"OvfPropertyExport": func() interface{} { return &OvfPropertyExport{} },

		"OvfPropertyNetwork": func() interface{} { return &OvfPropertyNetwork{} },

		"OvfPropertyNetworkExport": func() interface{} { return &OvfPropertyNetworkExport{} },

		"OvfPropertyQualifier": func() interface{} { return &OvfPropertyQualifier{} },

		"OvfPropertyQualifierDuplicate": func() interface{} { return &OvfPropertyQualifierDuplicate{} },

		"OvfPropertyQualifierIgnored": func() interface{} { return &OvfPropertyQualifierIgnored{} },

		"OvfPropertyType": func() interface{} { return &OvfPropertyType{} },

		"OvfPropertyValue": func() interface{} { return &OvfPropertyValue{} },

		"OvfResourceMap": func() interface{} { return &OvfResourceMap{} },

		"OvfSystemFault": func() interface{} { return &OvfSystemFault{} },

		"OvfToXmlUnsupportedElement": func() interface{} { return &OvfToXmlUnsupportedElement{} },

		"OvfUnableToExportDisk": func() interface{} { return &OvfUnableToExportDisk{} },

		"OvfUnexpectedElement": func() interface{} { return &OvfUnexpectedElement{} },

		"OvfUnknownDevice": func() interface{} { return &OvfUnknownDevice{} },

		"OvfUnknownDeviceBacking": func() interface{} { return &OvfUnknownDeviceBacking{} },

		"OvfUnknownEntity": func() interface{} { return &OvfUnknownEntity{} },

		"OvfUnsupportedAttribute": func() interface{} { return &OvfUnsupportedAttribute{} },

		"OvfUnsupportedAttributeValue": func() interface{} { return &OvfUnsupportedAttributeValue{} },

		"OvfUnsupportedDeviceBackingInfo": func() interface{} { return &OvfUnsupportedDeviceBackingInfo{} },

		"OvfUnsupportedDeviceBackingOption": func() interface{} { return &OvfUnsupportedDeviceBackingOption{} },

		"OvfUnsupportedDeviceExport": func() interface{} { return &OvfUnsupportedDeviceExport{} },

		"OvfUnsupportedDiskProvisioning": func() interface{} { return &OvfUnsupportedDiskProvisioning{} },

		"OvfUnsupportedElement": func() interface{} { return &OvfUnsupportedElement{} },

		"OvfUnsupportedElementValue": func() interface{} { return &OvfUnsupportedElementValue{} },

		"OvfUnsupportedPackage": func() interface{} { return &OvfUnsupportedPackage{} },

		"OvfUnsupportedSection": func() interface{} { return &OvfUnsupportedSection{} },

		"OvfUnsupportedSubType": func() interface{} { return &OvfUnsupportedSubType{} },

		"OvfUnsupportedType": func() interface{} { return &OvfUnsupportedType{} },

		"OvfValidateHostParams": func() interface{} { return &OvfValidateHostParams{} },

		"OvfValidateHostResult": func() interface{} { return &OvfValidateHostResult{} },

		"OvfWrongElement": func() interface{} { return &OvfWrongElement{} },

		"OvfWrongNamespace": func() interface{} { return &OvfWrongNamespace{} },

		"OvfXmlFormat": func() interface{} { return &OvfXmlFormat{} },

		"ParaVirtualSCSIController": func() interface{} { return &ParaVirtualSCSIController{} },

		"ParaVirtualSCSIControllerOption": func() interface{} { return &ParaVirtualSCSIControllerOption{} },

		"PasswordField": func() interface{} { return &PasswordField{} },

		"PatchAlreadyInstalled": func() interface{} { return &PatchAlreadyInstalled{} },

		"PatchBinariesNotFound": func() interface{} { return &PatchBinariesNotFound{} },

		"PatchInstallFailed": func() interface{} { return &PatchInstallFailed{} },

		"PatchIntegrityError": func() interface{} { return &PatchIntegrityError{} },

		"PatchMetadataCorrupted": func() interface{} { return &PatchMetadataCorrupted{} },

		"PatchMetadataInvalid": func() interface{} { return &PatchMetadataInvalid{} },

		"PatchMetadataNotFound": func() interface{} { return &PatchMetadataNotFound{} },

		"PatchMissingDependencies": func() interface{} { return &PatchMissingDependencies{} },

		"PatchNotApplicable": func() interface{} { return &PatchNotApplicable{} },

		"PatchSuperseded": func() interface{} { return &PatchSuperseded{} },

		"PerfCompositeMetric": func() interface{} { return &PerfCompositeMetric{} },

		"PerfCounterInfo": func() interface{} { return &PerfCounterInfo{} },

		"PerfEntityMetric": func() interface{} { return &PerfEntityMetric{} },

		"PerfEntityMetricBase": func() interface{} { return &PerfEntityMetricBase{} },

		"PerfEntityMetricCSV": func() interface{} { return &PerfEntityMetricCSV{} },

		"PerfFormat": func() interface{} { return &PerfFormat{} },

		"PerfInterval": func() interface{} { return &PerfInterval{} },

		"PerfMetricId": func() interface{} { return &PerfMetricId{} },

		"PerfMetricIntSeries": func() interface{} { return &PerfMetricIntSeries{} },

		"PerfMetricSeries": func() interface{} { return &PerfMetricSeries{} },

		"PerfMetricSeriesCSV": func() interface{} { return &PerfMetricSeriesCSV{} },

		"PerfProviderSummary": func() interface{} { return &PerfProviderSummary{} },

		"PerfQuerySpec": func() interface{} { return &PerfQuerySpec{} },

		"PerfSampleInfo": func() interface{} { return &PerfSampleInfo{} },

		"PerfStatsType": func() interface{} { return &PerfStatsType{} },

		"PerfSummaryType": func() interface{} { return &PerfSummaryType{} },

		"PerformanceDescription": func() interface{} { return &PerformanceDescription{} },

		"PerformanceManager": func() interface{} { return &PerformanceManager{} },

		"PerformanceManagerCounterLevelMapping": func() interface{} { return &PerformanceManagerCounterLevelMapping{} },

		"PerformanceManagerUnit": func() interface{} { return &PerformanceManagerUnit{} },

		"PerformanceStatisticsDescription": func() interface{} { return &PerformanceStatisticsDescription{} },

		"Permission": func() interface{} { return &Permission{} },

		"PermissionAddedEvent": func() interface{} { return &PermissionAddedEvent{} },

		"PermissionEvent": func() interface{} { return &PermissionEvent{} },

		"PermissionProfile": func() interface{} { return &PermissionProfile{} },

		"PermissionRemovedEvent": func() interface{} { return &PermissionRemovedEvent{} },

		"PermissionUpdatedEvent": func() interface{} { return &PermissionUpdatedEvent{} },

		"PhysCompatRDMNotSupported": func() interface{} { return &PhysCompatRDMNotSupported{} },

		"PhysicalNic": func() interface{} { return &PhysicalNic{} },

		"PhysicalNicCdpDeviceCapability": func() interface{} { return &PhysicalNicCdpDeviceCapability{} },

		"PhysicalNicCdpInfo": func() interface{} { return &PhysicalNicCdpInfo{} },

		"PhysicalNicConfig": func() interface{} { return &PhysicalNicConfig{} },

		"PhysicalNicHint": func() interface{} { return &PhysicalNicHint{} },

		"PhysicalNicHintInfo": func() interface{} { return &PhysicalNicHintInfo{} },

		"PhysicalNicIpHint": func() interface{} { return &PhysicalNicIpHint{} },

		"PhysicalNicLinkInfo": func() interface{} { return &PhysicalNicLinkInfo{} },

		"PhysicalNicNameHint": func() interface{} { return &PhysicalNicNameHint{} },

		"PhysicalNicProfile": func() interface{} { return &PhysicalNicProfile{} },

		"PhysicalNicResourcePoolSchedulerDisallowedReason": func() interface{} { return &PhysicalNicResourcePoolSchedulerDisallowedReason{} },

		"PhysicalNicSpec": func() interface{} { return &PhysicalNicSpec{} },

		"PhysicalNicVmDirectPathGen2SupportedMode": func() interface{} { return &PhysicalNicVmDirectPathGen2SupportedMode{} },

		"PlatformConfigFault": func() interface{} { return &PlatformConfigFault{} },

		"PnicUplinkProfile": func() interface{} { return &PnicUplinkProfile{} },

		"PodDiskLocator": func() interface{} { return &PodDiskLocator{} },

		"PodStorageDrsEntry": func() interface{} { return &PodStorageDrsEntry{} },

		"PolicyOption": func() interface{} { return &PolicyOption{} },

		"PortGroupConnecteeType": func() interface{} { return &PortGroupConnecteeType{} },

		"PortGroupProfile": func() interface{} { return &PortGroupProfile{} },

		"PosixUserSearchResult": func() interface{} { return &PosixUserSearchResult{} },

		"PowerOnFtSecondaryFailed": func() interface{} { return &PowerOnFtSecondaryFailed{} },

		"PowerOnFtSecondaryTimedout": func() interface{} { return &PowerOnFtSecondaryTimedout{} },

		"PowerSystemCapability": func() interface{} { return &PowerSystemCapability{} },

		"PowerSystemInfo": func() interface{} { return &PowerSystemInfo{} },

		"PrivilegeAvailability": func() interface{} { return &PrivilegeAvailability{} },

		"PrivilegePolicyDef": func() interface{} { return &PrivilegePolicyDef{} },

		"ProductComponentInfo": func() interface{} { return &ProductComponentInfo{} },

		"Profile": func() interface{} { return &Profile{} },

		"ProfileApplyProfileElement": func() interface{} { return &ProfileApplyProfileElement{} },

		"ProfileApplyProfileProperty": func() interface{} { return &ProfileApplyProfileProperty{} },

		"ProfileAssociatedEvent": func() interface{} { return &ProfileAssociatedEvent{} },

		"ProfileChangedEvent": func() interface{} { return &ProfileChangedEvent{} },

		"ProfileComplianceManager": func() interface{} { return &ProfileComplianceManager{} },

		"ProfileCompositeExpression": func() interface{} { return &ProfileCompositeExpression{} },

		"ProfileCompositePolicyOptionMetadata": func() interface{} { return &ProfileCompositePolicyOptionMetadata{} },

		"ProfileConfigInfo": func() interface{} { return &ProfileConfigInfo{} },

		"ProfileCreateSpec": func() interface{} { return &ProfileCreateSpec{} },

		"ProfileCreatedEvent": func() interface{} { return &ProfileCreatedEvent{} },

		"ProfileDeferredPolicyOptionParameter": func() interface{} { return &ProfileDeferredPolicyOptionParameter{} },

		"ProfileDescription": func() interface{} { return &ProfileDescription{} },

		"ProfileDescriptionSection": func() interface{} { return &ProfileDescriptionSection{} },

		"ProfileDissociatedEvent": func() interface{} { return &ProfileDissociatedEvent{} },

		"ProfileEvent": func() interface{} { return &ProfileEvent{} },

		"ProfileEventArgument": func() interface{} { return &ProfileEventArgument{} },

		"ProfileExecuteError": func() interface{} { return &ProfileExecuteError{} },

		"ProfileExecuteResult": func() interface{} { return &ProfileExecuteResult{} },

		"ProfileExecuteResultStatus": func() interface{} { return &ProfileExecuteResultStatus{} },

		"ProfileExpression": func() interface{} { return &ProfileExpression{} },

		"ProfileExpressionMetadata": func() interface{} { return &ProfileExpressionMetadata{} },

		"ProfileManager": func() interface{} { return &ProfileManager{} },

		"ProfileMetadata": func() interface{} { return &ProfileMetadata{} },

		"ProfileMetadataProfileSortSpec": func() interface{} { return &ProfileMetadataProfileSortSpec{} },

		"ProfileNumericComparator": func() interface{} { return &ProfileNumericComparator{} },

		"ProfileParameterMetadata": func() interface{} { return &ProfileParameterMetadata{} },

		"ProfilePolicy": func() interface{} { return &ProfilePolicy{} },

		"ProfilePolicyMetadata": func() interface{} { return &ProfilePolicyMetadata{} },

		"ProfilePolicyOptionMetadata": func() interface{} { return &ProfilePolicyOptionMetadata{} },

		"ProfileProfileStructure": func() interface{} { return &ProfileProfileStructure{} },

		"ProfileProfileStructureProperty": func() interface{} { return &ProfileProfileStructureProperty{} },

		"ProfilePropertyPath": func() interface{} { return &ProfilePropertyPath{} },

		"ProfileReferenceHostChangedEvent": func() interface{} { return &ProfileReferenceHostChangedEvent{} },

		"ProfileRemovedEvent": func() interface{} { return &ProfileRemovedEvent{} },

		"ProfileSerializedCreateSpec": func() interface{} { return &ProfileSerializedCreateSpec{} },

		"ProfileSimpleExpression": func() interface{} { return &ProfileSimpleExpression{} },

		"ProfileUpdateFailed": func() interface{} { return &ProfileUpdateFailed{} },

		"ProfileUpdateFailedUpdateFailure": func() interface{} { return &ProfileUpdateFailedUpdateFailure{} },

		"PropertyChange": func() interface{} { return &PropertyChange{} },

		"PropertyChangeOp": func() interface{} { return &PropertyChangeOp{} },

		"PropertyCollector": func() interface{} { return &PropertyCollector{} },

		"PropertyFilter": func() interface{} { return &PropertyFilter{} },

		"PropertyFilterSpec": func() interface{} { return &PropertyFilterSpec{} },

		"PropertyFilterUpdate": func() interface{} { return &PropertyFilterUpdate{} },

		"PropertySpec": func() interface{} { return &PropertySpec{} },

		"QuestionPending": func() interface{} { return &QuestionPending{} },

		"QuiesceDatastoreIOForHAFailed": func() interface{} { return &QuiesceDatastoreIOForHAFailed{} },

		"RDMConversionNotSupported": func() interface{} { return &RDMConversionNotSupported{} },

		"RDMNotPreserved": func() interface{} { return &RDMNotPreserved{} },

		"RDMNotSupported": func() interface{} { return &RDMNotSupported{} },

		"RDMNotSupportedOnDatastore": func() interface{} { return &RDMNotSupportedOnDatastore{} },

		"RDMPointsToInaccessibleDisk": func() interface{} { return &RDMPointsToInaccessibleDisk{} },

		"RawDiskNotSupported": func() interface{} { return &RawDiskNotSupported{} },

		"ReadHostResourcePoolTreeFailed": func() interface{} { return &ReadHostResourcePoolTreeFailed{} },

		"ReadOnlyDisksWithLegacyDestination": func() interface{} { return &ReadOnlyDisksWithLegacyDestination{} },

		"RebootRequired": func() interface{} { return &RebootRequired{} },

		"RecommendationReasonCode": func() interface{} { return &RecommendationReasonCode{} },

		"RecommendationType": func() interface{} { return &RecommendationType{} },

		"RecordReplayDisabled": func() interface{} { return &RecordReplayDisabled{} },

		"RecoveryEvent": func() interface{} { return &RecoveryEvent{} },

		"RecurrentTaskScheduler": func() interface{} { return &RecurrentTaskScheduler{} },

		"RemoteDeviceNotSupported": func() interface{} { return &RemoteDeviceNotSupported{} },

		"RemoteTSMEnabledEvent": func() interface{} { return &RemoteTSMEnabledEvent{} },

		"RemoveFailed": func() interface{} { return &RemoveFailed{} },

		"ReplicationConfigFault": func() interface{} { return &ReplicationConfigFault{} },

		"ReplicationDiskConfigFault": func() interface{} { return &ReplicationDiskConfigFault{} },

		"ReplicationDiskConfigFaultReasonForFault": func() interface{} { return &ReplicationDiskConfigFaultReasonForFault{} },

		"ReplicationFault": func() interface{} { return &ReplicationFault{} },

		"ReplicationIncompatibleWithFT": func() interface{} { return &ReplicationIncompatibleWithFT{} },

		"ReplicationInfoDiskSettings": func() interface{} { return &ReplicationInfoDiskSettings{} },

		"ReplicationInvalidOptions": func() interface{} { return &ReplicationInvalidOptions{} },

		"ReplicationNotSupportedOnHost": func() interface{} { return &ReplicationNotSupportedOnHost{} },

		"ReplicationVmConfigFault": func() interface{} { return &ReplicationVmConfigFault{} },

		"ReplicationVmConfigFaultReasonForFault": func() interface{} { return &ReplicationVmConfigFaultReasonForFault{} },

		"ReplicationVmFault": func() interface{} { return &ReplicationVmFault{} },

		"ReplicationVmFaultReasonForFault": func() interface{} { return &ReplicationVmFaultReasonForFault{} },

		"ReplicationVmProgressInfo": func() interface{} { return &ReplicationVmProgressInfo{} },

		"ReplicationVmState": func() interface{} { return &ReplicationVmState{} },

		"RequestCanceled": func() interface{} { return &RequestCanceled{} },

		"ResourceAllocationInfo": func() interface{} { return &ResourceAllocationInfo{} },

		"ResourceAllocationOption": func() interface{} { return &ResourceAllocationOption{} },

		"ResourceConfigOption": func() interface{} { return &ResourceConfigOption{} },

		"ResourceConfigSpec": func() interface{} { return &ResourceConfigSpec{} },

		"ResourceInUse": func() interface{} { return &ResourceInUse{} },

		"ResourceNotAvailable": func() interface{} { return &ResourceNotAvailable{} },

		"ResourcePlanningManager": func() interface{} { return &ResourcePlanningManager{} },

		"ResourcePool": func() interface{} { return &ResourcePool{} },

		"ResourcePoolCreatedEvent": func() interface{} { return &ResourcePoolCreatedEvent{} },

		"ResourcePoolDestroyedEvent": func() interface{} { return &ResourcePoolDestroyedEvent{} },

		"ResourcePoolEvent": func() interface{} { return &ResourcePoolEvent{} },

		"ResourcePoolEventArgument": func() interface{} { return &ResourcePoolEventArgument{} },

		"ResourcePoolMovedEvent": func() interface{} { return &ResourcePoolMovedEvent{} },

		"ResourcePoolQuickStats": func() interface{} { return &ResourcePoolQuickStats{} },

		"ResourcePoolReconfiguredEvent": func() interface{} { return &ResourcePoolReconfiguredEvent{} },

		"ResourcePoolResourceUsage": func() interface{} { return &ResourcePoolResourceUsage{} },

		"ResourcePoolRuntimeInfo": func() interface{} { return &ResourcePoolRuntimeInfo{} },

		"ResourcePoolSummary": func() interface{} { return &ResourcePoolSummary{} },

		"ResourceViolatedEvent": func() interface{} { return &ResourceViolatedEvent{} },

		"RestrictedVersion": func() interface{} { return &RestrictedVersion{} },

		"RetrieveOptions": func() interface{} { return &RetrieveOptions{} },

		"RetrieveResult": func() interface{} { return &RetrieveResult{} },

		"RoleAddedEvent": func() interface{} { return &RoleAddedEvent{} },

		"RoleEvent": func() interface{} { return &RoleEvent{} },

		"RoleEventArgument": func() interface{} { return &RoleEventArgument{} },

		"RoleRemovedEvent": func() interface{} { return &RoleRemovedEvent{} },

		"RoleUpdatedEvent": func() interface{} { return &RoleUpdatedEvent{} },

		"RollbackEvent": func() interface{} { return &RollbackEvent{} },

		"RollbackFailure": func() interface{} { return &RollbackFailure{} },

		"RuleViolation": func() interface{} { return &RuleViolation{} },

		"RunScriptAction": func() interface{} { return &RunScriptAction{} },

		"RuntimeFault": func() interface{} { return &RuntimeFault{} },

		"SSLDisabledFault": func() interface{} { return &SSLDisabledFault{} },

		"SSLVerifyFault": func() interface{} { return &SSLVerifyFault{} },

		"SSPIAuthentication": func() interface{} { return &SSPIAuthentication{} },

		"SSPIChallenge": func() interface{} { return &SSPIChallenge{} },

		"ScheduledHardwareUpgradeInfo": func() interface{} { return &ScheduledHardwareUpgradeInfo{} },

		"ScheduledHardwareUpgradeInfoHardwareUpgradePolicy": func() interface{} { return &ScheduledHardwareUpgradeInfoHardwareUpgradePolicy{} },

		"ScheduledHardwareUpgradeInfoHardwareUpgradeStatus": func() interface{} { return &ScheduledHardwareUpgradeInfoHardwareUpgradeStatus{} },

		"ScheduledTask": func() interface{} { return &ScheduledTask{} },

		"ScheduledTaskCompletedEvent": func() interface{} { return &ScheduledTaskCompletedEvent{} },

		"ScheduledTaskCreatedEvent": func() interface{} { return &ScheduledTaskCreatedEvent{} },

		"ScheduledTaskDescription": func() interface{} { return &ScheduledTaskDescription{} },

		"ScheduledTaskDetail": func() interface{} { return &ScheduledTaskDetail{} },

		"ScheduledTaskEmailCompletedEvent": func() interface{} { return &ScheduledTaskEmailCompletedEvent{} },

		"ScheduledTaskEmailFailedEvent": func() interface{} { return &ScheduledTaskEmailFailedEvent{} },

		"ScheduledTaskEvent": func() interface{} { return &ScheduledTaskEvent{} },

		"ScheduledTaskEventArgument": func() interface{} { return &ScheduledTaskEventArgument{} },

		"ScheduledTaskFailedEvent": func() interface{} { return &ScheduledTaskFailedEvent{} },

		"ScheduledTaskInfo": func() interface{} { return &ScheduledTaskInfo{} },

		"ScheduledTaskManager": func() interface{} { return &ScheduledTaskManager{} },

		"ScheduledTaskReconfiguredEvent": func() interface{} { return &ScheduledTaskReconfiguredEvent{} },

		"ScheduledTaskRemovedEvent": func() interface{} { return &ScheduledTaskRemovedEvent{} },

		"ScheduledTaskSpec": func() interface{} { return &ScheduledTaskSpec{} },

		"ScheduledTaskStartedEvent": func() interface{} { return &ScheduledTaskStartedEvent{} },

		"ScsiLun": func() interface{} { return &ScsiLun{} },

		"ScsiLunCapabilities": func() interface{} { return &ScsiLunCapabilities{} },

		"ScsiLunDescriptor": func() interface{} { return &ScsiLunDescriptor{} },

		"ScsiLunDescriptorQuality": func() interface{} { return &ScsiLunDescriptorQuality{} },

		"ScsiLunDurableName": func() interface{} { return &ScsiLunDurableName{} },

		"ScsiLunState": func() interface{} { return &ScsiLunState{} },

		"ScsiLunType": func() interface{} { return &ScsiLunType{} },

		"ScsiLunVStorageSupportStatus": func() interface{} { return &ScsiLunVStorageSupportStatus{} },

		"SeSparseVirtualDiskSpec": func() interface{} { return &SeSparseVirtualDiskSpec{} },

		"SearchIndex": func() interface{} { return &SearchIndex{} },

		"SecondaryVmAlreadyDisabled": func() interface{} { return &SecondaryVmAlreadyDisabled{} },

		"SecondaryVmAlreadyEnabled": func() interface{} { return &SecondaryVmAlreadyEnabled{} },

		"SecondaryVmAlreadyRegistered": func() interface{} { return &SecondaryVmAlreadyRegistered{} },

		"SecondaryVmNotRegistered": func() interface{} { return &SecondaryVmNotRegistered{} },

		"SecurityError": func() interface{} { return &SecurityError{} },

		"SecurityProfile": func() interface{} { return &SecurityProfile{} },

		"SelectionSet": func() interface{} { return &SelectionSet{} },

		"SelectionSpec": func() interface{} { return &SelectionSpec{} },

		"SendEmailAction": func() interface{} { return &SendEmailAction{} },

		"SendSNMPAction": func() interface{} { return &SendSNMPAction{} },

		"ServerLicenseExpiredEvent": func() interface{} { return &ServerLicenseExpiredEvent{} },

		"ServerStartedSessionEvent": func() interface{} { return &ServerStartedSessionEvent{} },

		"ServiceConsolePortGroupProfile": func() interface{} { return &ServiceConsolePortGroupProfile{} },

		"ServiceConsoleReservationInfo": func() interface{} { return &ServiceConsoleReservationInfo{} },

		"ServiceContent": func() interface{} { return &ServiceContent{} },

		"ServiceInstance": func() interface{} { return &ServiceInstance{} },

		"ServiceManager": func() interface{} { return &ServiceManager{} },

		"ServiceManagerServiceInfo": func() interface{} { return &ServiceManagerServiceInfo{} },

		"ServiceProfile": func() interface{} { return &ServiceProfile{} },

		"SessionEvent": func() interface{} { return &SessionEvent{} },

		"SessionManager": func() interface{} { return &SessionManager{} },

		"SessionManagerGenericServiceTicket": func() interface{} { return &SessionManagerGenericServiceTicket{} },

		"SessionManagerHttpServiceRequestSpec": func() interface{} { return &SessionManagerHttpServiceRequestSpec{} },

		"SessionManagerHttpServiceRequestSpecMethod": func() interface{} { return &SessionManagerHttpServiceRequestSpecMethod{} },

		"SessionManagerLocalTicket": func() interface{} { return &SessionManagerLocalTicket{} },

		"SessionManagerServiceRequestSpec": func() interface{} { return &SessionManagerServiceRequestSpec{} },

		"SessionManagerVmomiServiceRequestSpec": func() interface{} { return &SessionManagerVmomiServiceRequestSpec{} },

		"SessionTerminatedEvent": func() interface{} { return &SessionTerminatedEvent{} },

		"SharedBusControllerNotSupported": func() interface{} { return &SharedBusControllerNotSupported{} },

		"SharesInfo": func() interface{} { return &SharesInfo{} },

		"SharesLevel": func() interface{} { return &SharesLevel{} },

		"SharesOption": func() interface{} { return &SharesOption{} },

		"ShrinkDiskFault": func() interface{} { return &ShrinkDiskFault{} },

		"SimpleCommand": func() interface{} { return &SimpleCommand{} },

		"SimpleCommandEncoding": func() interface{} { return &SimpleCommandEncoding{} },

		"SingleIp": func() interface{} { return &SingleIp{} },

		"SingleMac": func() interface{} { return &SingleMac{} },

		"SlpDiscoveryMethod": func() interface{} { return &SlpDiscoveryMethod{} },

		"SnapshotCloneNotSupported": func() interface{} { return &SnapshotCloneNotSupported{} },

		"SnapshotCopyNotSupported": func() interface{} { return &SnapshotCopyNotSupported{} },

		"SnapshotDisabled": func() interface{} { return &SnapshotDisabled{} },

		"SnapshotFault": func() interface{} { return &SnapshotFault{} },

		"SnapshotIncompatibleDeviceInVm": func() interface{} { return &SnapshotIncompatibleDeviceInVm{} },

		"SnapshotLocked": func() interface{} { return &SnapshotLocked{} },

		"SnapshotMoveFromNonHomeNotSupported": func() interface{} { return &SnapshotMoveFromNonHomeNotSupported{} },

		"SnapshotMoveNotSupported": func() interface{} { return &SnapshotMoveNotSupported{} },

		"SnapshotMoveToNonHomeNotSupported": func() interface{} { return &SnapshotMoveToNonHomeNotSupported{} },

		"SnapshotNoChange": func() interface{} { return &SnapshotNoChange{} },

		"SnapshotRevertIssue": func() interface{} { return &SnapshotRevertIssue{} },

		"SoftRuleVioCorrectionDisallowed": func() interface{} { return &SoftRuleVioCorrectionDisallowed{} },

		"SoftRuleVioCorrectionImpact": func() interface{} { return &SoftRuleVioCorrectionImpact{} },

		"SsdDiskNotAvailable": func() interface{} { return &SsdDiskNotAvailable{} },

		"StateAlarmExpression": func() interface{} { return &StateAlarmExpression{} },

		"StateAlarmOperator": func() interface{} { return &StateAlarmOperator{} },

		"StaticRouteProfile": func() interface{} { return &StaticRouteProfile{} },

		"StorageDrsCannotMoveDiskInMultiWriterMode": func() interface{} { return &StorageDrsCannotMoveDiskInMultiWriterMode{} },

		"StorageDrsCannotMoveFTVm": func() interface{} { return &StorageDrsCannotMoveFTVm{} },

		"StorageDrsCannotMoveIndependentDisk": func() interface{} { return &StorageDrsCannotMoveIndependentDisk{} },

		"StorageDrsCannotMoveManuallyPlacedSwapFile": func() interface{} { return &StorageDrsCannotMoveManuallyPlacedSwapFile{} },

		"StorageDrsCannotMoveManuallyPlacedVm": func() interface{} { return &StorageDrsCannotMoveManuallyPlacedVm{} },

		"StorageDrsCannotMoveSharedDisk": func() interface{} { return &StorageDrsCannotMoveSharedDisk{} },

		"StorageDrsCannotMoveTemplate": func() interface{} { return &StorageDrsCannotMoveTemplate{} },

		"StorageDrsCannotMoveVmInUserFolder": func() interface{} { return &StorageDrsCannotMoveVmInUserFolder{} },

		"StorageDrsCannotMoveVmWithMountedCDROM": func() interface{} { return &StorageDrsCannotMoveVmWithMountedCDROM{} },

		"StorageDrsCannotMoveVmWithNoFilesInLayout": func() interface{} { return &StorageDrsCannotMoveVmWithNoFilesInLayout{} },

		"StorageDrsConfigInfo": func() interface{} { return &StorageDrsConfigInfo{} },

		"StorageDrsConfigSpec": func() interface{} { return &StorageDrsConfigSpec{} },

		"StorageDrsDatacentersCannotShareDatastore": func() interface{} { return &StorageDrsDatacentersCannotShareDatastore{} },

		"StorageDrsDisabledOnVm": func() interface{} { return &StorageDrsDisabledOnVm{} },

		"StorageDrsIoLoadBalanceConfig": func() interface{} { return &StorageDrsIoLoadBalanceConfig{} },

		"StorageDrsIolbDisabledInternally": func() interface{} { return &StorageDrsIolbDisabledInternally{} },

		"StorageDrsOptionSpec": func() interface{} { return &StorageDrsOptionSpec{} },

		"StorageDrsPodConfigInfo": func() interface{} { return &StorageDrsPodConfigInfo{} },

		"StorageDrsPodConfigInfoBehavior": func() interface{} { return &StorageDrsPodConfigInfoBehavior{} },

		"StorageDrsPodConfigSpec": func() interface{} { return &StorageDrsPodConfigSpec{} },

		"StorageDrsPodSelectionSpec": func() interface{} { return &StorageDrsPodSelectionSpec{} },

		"StorageDrsSpaceLoadBalanceConfig": func() interface{} { return &StorageDrsSpaceLoadBalanceConfig{} },

		"StorageDrsUnableToMoveFiles": func() interface{} { return &StorageDrsUnableToMoveFiles{} },

		"StorageDrsVmConfigInfo": func() interface{} { return &StorageDrsVmConfigInfo{} },

		"StorageDrsVmConfigSpec": func() interface{} { return &StorageDrsVmConfigSpec{} },

		"StorageIOAllocationInfo": func() interface{} { return &StorageIOAllocationInfo{} },

		"StorageIOAllocationOption": func() interface{} { return &StorageIOAllocationOption{} },

		"StorageIORMConfigOption": func() interface{} { return &StorageIORMConfigOption{} },

		"StorageIORMConfigSpec": func() interface{} { return &StorageIORMConfigSpec{} },

		"StorageIORMInfo": func() interface{} { return &StorageIORMInfo{} },

		"StorageIORMThresholdMode": func() interface{} { return &StorageIORMThresholdMode{} },

		"StorageMigrationAction": func() interface{} { return &StorageMigrationAction{} },

		"StoragePerformanceSummary": func() interface{} { return &StoragePerformanceSummary{} },

		"StoragePlacementAction": func() interface{} { return &StoragePlacementAction{} },

		"StoragePlacementResult": func() interface{} { return &StoragePlacementResult{} },

		"StoragePlacementSpec": func() interface{} { return &StoragePlacementSpec{} },

		"StoragePlacementSpecPlacementType": func() interface{} { return &StoragePlacementSpecPlacementType{} },

		"StoragePod": func() interface{} { return &StoragePod{} },

		"StoragePodSummary": func() interface{} { return &StoragePodSummary{} },

		"StorageProfile": func() interface{} { return &StorageProfile{} },

		"StorageRequirement": func() interface{} { return &StorageRequirement{} },

		"StorageResourceManager": func() interface{} { return &StorageResourceManager{} },

		"StorageVMotionNotSupported": func() interface{} { return &StorageVMotionNotSupported{} },

		"StorageVmotionIncompatible": func() interface{} { return &StorageVmotionIncompatible{} },

		"StringExpression": func() interface{} { return &StringExpression{} },

		"StringOption": func() interface{} { return &StringOption{} },

		"StringPolicy": func() interface{} { return &StringPolicy{} },

		"SuspendedRelocateNotSupported": func() interface{} { return &SuspendedRelocateNotSupported{} },

		"SwapDatastoreNotWritableOnHost": func() interface{} { return &SwapDatastoreNotWritableOnHost{} },

		"SwapDatastoreUnset": func() interface{} { return &SwapDatastoreUnset{} },

		"SwapPlacementOverrideNotSupported": func() interface{} { return &SwapPlacementOverrideNotSupported{} },

		"SwitchIpUnset": func() interface{} { return &SwitchIpUnset{} },

		"SwitchNotInUpgradeMode": func() interface{} { return &SwitchNotInUpgradeMode{} },

		"SystemError": func() interface{} { return &SystemError{} },

		"Tag": func() interface{} { return &Tag{} },

		"Task": func() interface{} { return &Task{} },

		"TaskDescription": func() interface{} { return &TaskDescription{} },

		"TaskEvent": func() interface{} { return &TaskEvent{} },

		"TaskFilterSpec": func() interface{} { return &TaskFilterSpec{} },

		"TaskFilterSpecByEntity": func() interface{} { return &TaskFilterSpecByEntity{} },

		"TaskFilterSpecByTime": func() interface{} { return &TaskFilterSpecByTime{} },

		"TaskFilterSpecByUsername": func() interface{} { return &TaskFilterSpecByUsername{} },

		"TaskFilterSpecRecursionOption": func() interface{} { return &TaskFilterSpecRecursionOption{} },

		"TaskFilterSpecTimeOption": func() interface{} { return &TaskFilterSpecTimeOption{} },

		"TaskHistoryCollector": func() interface{} { return &TaskHistoryCollector{} },

		"TaskInProgress": func() interface{} { return &TaskInProgress{} },

		"TaskInfo": func() interface{} { return &TaskInfo{} },

		"TaskInfoState": func() interface{} { return &TaskInfoState{} },

		"TaskManager": func() interface{} { return &TaskManager{} },

		"TaskReason": func() interface{} { return &TaskReason{} },

		"TaskReasonAlarm": func() interface{} { return &TaskReasonAlarm{} },

		"TaskReasonSchedule": func() interface{} { return &TaskReasonSchedule{} },

		"TaskReasonSystem": func() interface{} { return &TaskReasonSystem{} },

		"TaskReasonUser": func() interface{} { return &TaskReasonUser{} },

		"TaskScheduler": func() interface{} { return &TaskScheduler{} },

		"TaskTimeoutEvent": func() interface{} { return &TaskTimeoutEvent{} },

		"TeamingMatchEvent": func() interface{} { return &TeamingMatchEvent{} },

		"TeamingMisMatchEvent": func() interface{} { return &TeamingMisMatchEvent{} },

		"TemplateBeingUpgradedEvent": func() interface{} { return &TemplateBeingUpgradedEvent{} },

		"TemplateConfigFileInfo": func() interface{} { return &TemplateConfigFileInfo{} },

		"TemplateConfigFileQuery": func() interface{} { return &TemplateConfigFileQuery{} },

		"TemplateUpgradeEvent": func() interface{} { return &TemplateUpgradeEvent{} },

		"TemplateUpgradeFailedEvent": func() interface{} { return &TemplateUpgradeFailedEvent{} },

		"TemplateUpgradedEvent": func() interface{} { return &TemplateUpgradedEvent{} },

		"ThirdPartyLicenseAssignmentFailed": func() interface{} { return &ThirdPartyLicenseAssignmentFailed{} },

		"ThirdPartyLicenseAssignmentFailedReason": func() interface{} { return &ThirdPartyLicenseAssignmentFailedReason{} },

		"TicketedSessionAuthentication": func() interface{} { return &TicketedSessionAuthentication{} },

		"TimedOutHostOperationEvent": func() interface{} { return &TimedOutHostOperationEvent{} },

		"Timedout": func() interface{} { return &Timedout{} },

		"TooManyConcurrentNativeClones": func() interface{} { return &TooManyConcurrentNativeClones{} },

		"TooManyConsecutiveOverrides": func() interface{} { return &TooManyConsecutiveOverrides{} },

		"TooManyDevices": func() interface{} { return &TooManyDevices{} },

		"TooManyDisksOnLegacyHost": func() interface{} { return &TooManyDisksOnLegacyHost{} },

		"TooManyGuestLogons": func() interface{} { return &TooManyGuestLogons{} },

		"TooManyHosts": func() interface{} { return &TooManyHosts{} },

		"TooManyNativeCloneLevels": func() interface{} { return &TooManyNativeCloneLevels{} },

		"TooManyNativeClonesOnFile": func() interface{} { return &TooManyNativeClonesOnFile{} },

		"TooManySnapshotLevels": func() interface{} { return &TooManySnapshotLevels{} },

		"ToolsAlreadyUpgraded": func() interface{} { return &ToolsAlreadyUpgraded{} },

		"ToolsAutoUpgradeNotSupported": func() interface{} { return &ToolsAutoUpgradeNotSupported{} },

		"ToolsConfigInfo": func() interface{} { return &ToolsConfigInfo{} },

		"ToolsConfigInfoToolsLastInstallInfo": func() interface{} { return &ToolsConfigInfoToolsLastInstallInfo{} },

		"ToolsImageCopyFailed": func() interface{} { return &ToolsImageCopyFailed{} },

		"ToolsImageNotAvailable": func() interface{} { return &ToolsImageNotAvailable{} },

		"ToolsImageSignatureCheckFailed": func() interface{} { return &ToolsImageSignatureCheckFailed{} },

		"ToolsInstallationInProgress": func() interface{} { return &ToolsInstallationInProgress{} },

		"ToolsUnavailable": func() interface{} { return &ToolsUnavailable{} },

		"ToolsUpgradeCancelled": func() interface{} { return &ToolsUpgradeCancelled{} },

		"TraversalSpec": func() interface{} { return &TraversalSpec{} },

		"TypeDescription": func() interface{} { return &TypeDescription{} },

		"UnSupportedDatastoreForVFlash": func() interface{} { return &UnSupportedDatastoreForVFlash{} },

		"UncommittedUndoableDisk": func() interface{} { return &UncommittedUndoableDisk{} },

		"UnconfiguredPropertyValue": func() interface{} { return &UnconfiguredPropertyValue{} },

		"UncustomizableGuest": func() interface{} { return &UncustomizableGuest{} },

		"UnexpectedCustomizationFault": func() interface{} { return &UnexpectedCustomizationFault{} },

		"UnexpectedFault": func() interface{} { return &UnexpectedFault{} },

		"UnlicensedVirtualMachinesEvent": func() interface{} { return &UnlicensedVirtualMachinesEvent{} },

		"UnlicensedVirtualMachinesFoundEvent": func() interface{} { return &UnlicensedVirtualMachinesFoundEvent{} },

		"UnrecognizedHost": func() interface{} { return &UnrecognizedHost{} },

		"UnsharedSwapVMotionNotSupported": func() interface{} { return &UnsharedSwapVMotionNotSupported{} },

		"UnsupportedDatastore": func() interface{} { return &UnsupportedDatastore{} },

		"UnsupportedGuest": func() interface{} { return &UnsupportedGuest{} },

		"UnsupportedVimApiVersion": func() interface{} { return &UnsupportedVimApiVersion{} },

		"UnsupportedVmxLocation": func() interface{} { return &UnsupportedVmxLocation{} },

		"UnusedVirtualDiskBlocksNotScrubbed": func() interface{} { return &UnusedVirtualDiskBlocksNotScrubbed{} },

		"UpdateSet": func() interface{} { return &UpdateSet{} },

		"UpdateVirtualMachineFilesResult": func() interface{} { return &UpdateVirtualMachineFilesResult{} },

		"UpdateVirtualMachineFilesResultFailedVmFileInfo": func() interface{} { return &UpdateVirtualMachineFilesResultFailedVmFileInfo{} },

		"UpdatedAgentBeingRestartedEvent": func() interface{} { return &UpdatedAgentBeingRestartedEvent{} },

		"UpgradeEvent": func() interface{} { return &UpgradeEvent{} },

		"UpgradePolicy": func() interface{} { return &UpgradePolicy{} },

		"UplinkPortMtuNotSupportEvent": func() interface{} { return &UplinkPortMtuNotSupportEvent{} },

		"UplinkPortMtuSupportEvent": func() interface{} { return &UplinkPortMtuSupportEvent{} },

		"UplinkPortVlanTrunkedEvent": func() interface{} { return &UplinkPortVlanTrunkedEvent{} },

		"UplinkPortVlanUntrunkedEvent": func() interface{} { return &UplinkPortVlanUntrunkedEvent{} },

		"UserAssignedToGroup": func() interface{} { return &UserAssignedToGroup{} },

		"UserDirectory": func() interface{} { return &UserDirectory{} },

		"UserGroupProfile": func() interface{} { return &UserGroupProfile{} },

		"UserInputRequiredParameterMetadata": func() interface{} { return &UserInputRequiredParameterMetadata{} },

		"UserLoginSessionEvent": func() interface{} { return &UserLoginSessionEvent{} },

		"UserLogoutSessionEvent": func() interface{} { return &UserLogoutSessionEvent{} },

		"UserNotFound": func() interface{} { return &UserNotFound{} },

		"UserPasswordChanged": func() interface{} { return &UserPasswordChanged{} },

		"UserProfile": func() interface{} { return &UserProfile{} },

		"UserSearchResult": func() interface{} { return &UserSearchResult{} },

		"UserSession": func() interface{} { return &UserSession{} },

		"UserUnassignedFromGroup": func() interface{} { return &UserUnassignedFromGroup{} },

		"UserUpgradeEvent": func() interface{} { return &UserUpgradeEvent{} },

		"VAppAutoStartAction": func() interface{} { return &VAppAutoStartAction{} },

		"VAppCloneSpec": func() interface{} { return &VAppCloneSpec{} },

		"VAppCloneSpecNetworkMappingPair": func() interface{} { return &VAppCloneSpecNetworkMappingPair{} },

		"VAppCloneSpecProvisioningType": func() interface{} { return &VAppCloneSpecProvisioningType{} },

		"VAppCloneSpecResourceMap": func() interface{} { return &VAppCloneSpecResourceMap{} },

		"VAppConfigFault": func() interface{} { return &VAppConfigFault{} },

		"VAppConfigInfo": func() interface{} { return &VAppConfigInfo{} },

		"VAppConfigSpec": func() interface{} { return &VAppConfigSpec{} },

		"VAppEntityConfigInfo": func() interface{} { return &VAppEntityConfigInfo{} },

		"VAppIPAssignmentInfo": func() interface{} { return &VAppIPAssignmentInfo{} },

		"VAppIPAssignmentInfoAllocationSchemes": func() interface{} { return &VAppIPAssignmentInfoAllocationSchemes{} },

		"VAppIPAssignmentInfoIpAllocationPolicy": func() interface{} { return &VAppIPAssignmentInfoIpAllocationPolicy{} },

		"VAppIPAssignmentInfoProtocols": func() interface{} { return &VAppIPAssignmentInfoProtocols{} },

		"VAppNotRunning": func() interface{} { return &VAppNotRunning{} },

		"VAppOperationInProgress": func() interface{} { return &VAppOperationInProgress{} },

		"VAppOvfSectionInfo": func() interface{} { return &VAppOvfSectionInfo{} },

		"VAppOvfSectionSpec": func() interface{} { return &VAppOvfSectionSpec{} },

		"VAppProductInfo": func() interface{} { return &VAppProductInfo{} },

		"VAppProductSpec": func() interface{} { return &VAppProductSpec{} },

		"VAppPropertyFault": func() interface{} { return &VAppPropertyFault{} },

		"VAppPropertyInfo": func() interface{} { return &VAppPropertyInfo{} },

		"VAppPropertySpec": func() interface{} { return &VAppPropertySpec{} },

		"VAppTaskInProgress": func() interface{} { return &VAppTaskInProgress{} },

		"VFlashModuleNotSupported": func() interface{} { return &VFlashModuleNotSupported{} },

		"VFlashModuleNotSupportedReason": func() interface{} { return &VFlashModuleNotSupportedReason{} },

		"VFlashModuleVersionIncompatible": func() interface{} { return &VFlashModuleVersionIncompatible{} },

		"VMFSDatastoreCreatedEvent": func() interface{} { return &VMFSDatastoreCreatedEvent{} },

		"VMFSDatastoreExpandedEvent": func() interface{} { return &VMFSDatastoreExpandedEvent{} },

		"VMFSDatastoreExtendedEvent": func() interface{} { return &VMFSDatastoreExtendedEvent{} },

		"VMINotSupported": func() interface{} { return &VMINotSupported{} },

		"VMOnConflictDVPort": func() interface{} { return &VMOnConflictDVPort{} },

		"VMOnVirtualIntranet": func() interface{} { return &VMOnVirtualIntranet{} },

		"VMotionAcrossNetworkNotSupported": func() interface{} { return &VMotionAcrossNetworkNotSupported{} },

		"VMotionCompatibilityType": func() interface{} { return &VMotionCompatibilityType{} },

		"VMotionInterfaceIssue": func() interface{} { return &VMotionInterfaceIssue{} },

		"VMotionLicenseExpiredEvent": func() interface{} { return &VMotionLicenseExpiredEvent{} },

		"VMotionLinkCapacityLow": func() interface{} { return &VMotionLinkCapacityLow{} },

		"VMotionLinkDown": func() interface{} { return &VMotionLinkDown{} },

		"VMotionNotConfigured": func() interface{} { return &VMotionNotConfigured{} },

		"VMotionNotLicensed": func() interface{} { return &VMotionNotLicensed{} },

		"VMotionNotSupported": func() interface{} { return &VMotionNotSupported{} },

		"VMotionProtocolIncompatible": func() interface{} { return &VMotionProtocolIncompatible{} },

		"VMwareDVSConfigInfo": func() interface{} { return &VMwareDVSConfigInfo{} },

		"VMwareDVSConfigSpec": func() interface{} { return &VMwareDVSConfigSpec{} },

		"VMwareDVSFeatureCapability": func() interface{} { return &VMwareDVSFeatureCapability{} },

		"VMwareDVSHealthCheckCapability": func() interface{} { return &VMwareDVSHealthCheckCapability{} },

		"VMwareDVSHealthCheckConfig": func() interface{} { return &VMwareDVSHealthCheckConfig{} },

		"VMwareDVSMtuHealthCheckResult": func() interface{} { return &VMwareDVSMtuHealthCheckResult{} },

		"VMwareDVSPortSetting": func() interface{} { return &VMwareDVSPortSetting{} },

		"VMwareDVSPortgroupPolicy": func() interface{} { return &VMwareDVSPortgroupPolicy{} },

		"VMwareDVSPvlanConfigSpec": func() interface{} { return &VMwareDVSPvlanConfigSpec{} },

		"VMwareDVSPvlanMapEntry": func() interface{} { return &VMwareDVSPvlanMapEntry{} },

		"VMwareDVSTeamingHealthCheckConfig": func() interface{} { return &VMwareDVSTeamingHealthCheckConfig{} },

		"VMwareDVSTeamingHealthCheckResult": func() interface{} { return &VMwareDVSTeamingHealthCheckResult{} },

		"VMwareDVSTeamingMatchStatus": func() interface{} { return &VMwareDVSTeamingMatchStatus{} },

		"VMwareDVSVlanHealthCheckResult": func() interface{} { return &VMwareDVSVlanHealthCheckResult{} },

		"VMwareDVSVlanMtuHealthCheckConfig": func() interface{} { return &VMwareDVSVlanMtuHealthCheckConfig{} },

		"VMwareDVSVspanCapability": func() interface{} { return &VMwareDVSVspanCapability{} },

		"VMwareDVSVspanConfigSpec": func() interface{} { return &VMwareDVSVspanConfigSpec{} },

		"VMwareDVSVspanSessionType": func() interface{} { return &VMwareDVSVspanSessionType{} },

		"VMwareDvsLacpApiVersion": func() interface{} { return &VMwareDvsLacpApiVersion{} },

		"VMwareDvsLacpCapability": func() interface{} { return &VMwareDvsLacpCapability{} },

		"VMwareDvsLacpGroupConfig": func() interface{} { return &VMwareDvsLacpGroupConfig{} },

		"VMwareDvsLacpGroupSpec": func() interface{} { return &VMwareDvsLacpGroupSpec{} },

		"VMwareDvsLacpLoadBalanceAlgorithm": func() interface{} { return &VMwareDvsLacpLoadBalanceAlgorithm{} },

		"VMwareDvsLagIpfixConfig": func() interface{} { return &VMwareDvsLagIpfixConfig{} },

		"VMwareDvsLagVlanConfig": func() interface{} { return &VMwareDvsLagVlanConfig{} },

		"VMwareIpfixConfig": func() interface{} { return &VMwareIpfixConfig{} },

		"VMwareUplinkLacpMode": func() interface{} { return &VMwareUplinkLacpMode{} },

		"VMwareUplinkLacpPolicy": func() interface{} { return &VMwareUplinkLacpPolicy{} },

		"VMwareUplinkPortOrderPolicy": func() interface{} { return &VMwareUplinkPortOrderPolicy{} },

		"VMwareVspanPort": func() interface{} { return &VMwareVspanPort{} },

		"VMwareVspanSession": func() interface{} { return &VMwareVspanSession{} },

		"ValidateMigrationTestType": func() interface{} { return &ValidateMigrationTestType{} },

		"VcAgentUninstallFailedEvent": func() interface{} { return &VcAgentUninstallFailedEvent{} },

		"VcAgentUninstalledEvent": func() interface{} { return &VcAgentUninstalledEvent{} },

		"VcAgentUpgradeFailedEvent": func() interface{} { return &VcAgentUpgradeFailedEvent{} },

		"VcAgentUpgradedEvent": func() interface{} { return &VcAgentUpgradedEvent{} },

		"View": func() interface{} { return &View{} },

		"ViewManager": func() interface{} { return &ViewManager{} },

		"VimAccountPasswordChangedEvent": func() interface{} { return &VimAccountPasswordChangedEvent{} },

		"VimFault": func() interface{} { return &VimFault{} },

		"VirtualAHCIController": func() interface{} { return &VirtualAHCIController{} },

		"VirtualAHCIControllerOption": func() interface{} { return &VirtualAHCIControllerOption{} },

		"VirtualApp": func() interface{} { return &VirtualApp{} },

		"VirtualAppImportSpec": func() interface{} { return &VirtualAppImportSpec{} },

		"VirtualAppLinkInfo": func() interface{} { return &VirtualAppLinkInfo{} },

		"VirtualAppSummary": func() interface{} { return &VirtualAppSummary{} },

		"VirtualAppVAppState": func() interface{} { return &VirtualAppVAppState{} },

		"VirtualBusLogicController": func() interface{} { return &VirtualBusLogicController{} },

		"VirtualBusLogicControllerOption": func() interface{} { return &VirtualBusLogicControllerOption{} },

		"VirtualCdrom": func() interface{} { return &VirtualCdrom{} },

		"VirtualCdromAtapiBackingInfo": func() interface{} { return &VirtualCdromAtapiBackingInfo{} },

		"VirtualCdromAtapiBackingOption": func() interface{} { return &VirtualCdromAtapiBackingOption{} },

		"VirtualCdromIsoBackingInfo": func() interface{} { return &VirtualCdromIsoBackingInfo{} },

		"VirtualCdromIsoBackingOption": func() interface{} { return &VirtualCdromIsoBackingOption{} },

		"VirtualCdromOption": func() interface{} { return &VirtualCdromOption{} },

		"VirtualCdromPassthroughBackingInfo": func() interface{} { return &VirtualCdromPassthroughBackingInfo{} },

		"VirtualCdromPassthroughBackingOption": func() interface{} { return &VirtualCdromPassthroughBackingOption{} },

		"VirtualCdromRemoteAtapiBackingInfo": func() interface{} { return &VirtualCdromRemoteAtapiBackingInfo{} },

		"VirtualCdromRemoteAtapiBackingOption": func() interface{} { return &VirtualCdromRemoteAtapiBackingOption{} },

		"VirtualCdromRemotePassthroughBackingInfo": func() interface{} { return &VirtualCdromRemotePassthroughBackingInfo{} },

		"VirtualCdromRemotePassthroughBackingOption": func() interface{} { return &VirtualCdromRemotePassthroughBackingOption{} },

		"VirtualController": func() interface{} { return &VirtualController{} },

		"VirtualControllerOption": func() interface{} { return &VirtualControllerOption{} },

		"VirtualDevice": func() interface{} { return &VirtualDevice{} },

		"VirtualDeviceBackingInfo": func() interface{} { return &VirtualDeviceBackingInfo{} },

		"VirtualDeviceBackingOption": func() interface{} { return &VirtualDeviceBackingOption{} },

		"VirtualDeviceBusSlotInfo": func() interface{} { return &VirtualDeviceBusSlotInfo{} },

		"VirtualDeviceBusSlotOption": func() interface{} { return &VirtualDeviceBusSlotOption{} },

		"VirtualDeviceConfigSpec": func() interface{} { return &VirtualDeviceConfigSpec{} },

		"VirtualDeviceConfigSpecFileOperation": func() interface{} { return &VirtualDeviceConfigSpecFileOperation{} },

		"VirtualDeviceConfigSpecOperation": func() interface{} { return &VirtualDeviceConfigSpecOperation{} },

		"VirtualDeviceConnectInfo": func() interface{} { return &VirtualDeviceConnectInfo{} },

		"VirtualDeviceConnectInfoStatus": func() interface{} { return &VirtualDeviceConnectInfoStatus{} },

		"VirtualDeviceConnectOption": func() interface{} { return &VirtualDeviceConnectOption{} },

		"VirtualDeviceDeviceBackingInfo": func() interface{} { return &VirtualDeviceDeviceBackingInfo{} },

		"VirtualDeviceDeviceBackingOption": func() interface{} { return &VirtualDeviceDeviceBackingOption{} },

		"VirtualDeviceFileBackingInfo": func() interface{} { return &VirtualDeviceFileBackingInfo{} },

		"VirtualDeviceFileBackingOption": func() interface{} { return &VirtualDeviceFileBackingOption{} },

		"VirtualDeviceFileExtension": func() interface{} { return &VirtualDeviceFileExtension{} },

		"VirtualDeviceOption": func() interface{} { return &VirtualDeviceOption{} },

		"VirtualDevicePciBusSlotInfo": func() interface{} { return &VirtualDevicePciBusSlotInfo{} },

		"VirtualDevicePipeBackingInfo": func() interface{} { return &VirtualDevicePipeBackingInfo{} },

		"VirtualDevicePipeBackingOption": func() interface{} { return &VirtualDevicePipeBackingOption{} },

		"VirtualDeviceRemoteDeviceBackingInfo": func() interface{} { return &VirtualDeviceRemoteDeviceBackingInfo{} },

		"VirtualDeviceRemoteDeviceBackingOption": func() interface{} { return &VirtualDeviceRemoteDeviceBackingOption{} },

		"VirtualDeviceURIBackingInfo": func() interface{} { return &VirtualDeviceURIBackingInfo{} },

		"VirtualDeviceURIBackingOption": func() interface{} { return &VirtualDeviceURIBackingOption{} },

		"VirtualDeviceURIBackingOptionDirection": func() interface{} { return &VirtualDeviceURIBackingOptionDirection{} },

		"VirtualDisk": func() interface{} { return &VirtualDisk{} },

		"VirtualDiskAdapterType": func() interface{} { return &VirtualDiskAdapterType{} },

		"VirtualDiskAntiAffinityRuleSpec": func() interface{} { return &VirtualDiskAntiAffinityRuleSpec{} },

		"VirtualDiskBlocksNotFullyProvisioned": func() interface{} { return &VirtualDiskBlocksNotFullyProvisioned{} },

		"VirtualDiskCompatibilityMode": func() interface{} { return &VirtualDiskCompatibilityMode{} },

		"VirtualDiskConfigSpec": func() interface{} { return &VirtualDiskConfigSpec{} },

		"VirtualDiskDeltaDiskFormat": func() interface{} { return &VirtualDiskDeltaDiskFormat{} },

		"VirtualDiskDeltaDiskFormatsSupported": func() interface{} { return &VirtualDiskDeltaDiskFormatsSupported{} },

		"VirtualDiskFlatVer1BackingInfo": func() interface{} { return &VirtualDiskFlatVer1BackingInfo{} },

		"VirtualDiskFlatVer1BackingOption": func() interface{} { return &VirtualDiskFlatVer1BackingOption{} },

		"VirtualDiskFlatVer2BackingInfo": func() interface{} { return &VirtualDiskFlatVer2BackingInfo{} },

		"VirtualDiskFlatVer2BackingOption": func() interface{} { return &VirtualDiskFlatVer2BackingOption{} },

		"VirtualDiskId": func() interface{} { return &VirtualDiskId{} },

		"VirtualDiskManager": func() interface{} { return &VirtualDiskManager{} },

		"VirtualDiskMode": func() interface{} { return &VirtualDiskMode{} },

		"VirtualDiskModeNotSupported": func() interface{} { return &VirtualDiskModeNotSupported{} },

		"VirtualDiskOption": func() interface{} { return &VirtualDiskOption{} },

		"VirtualDiskOptionVFlashCacheConfigOption": func() interface{} { return &VirtualDiskOptionVFlashCacheConfigOption{} },

		"VirtualDiskPartitionedRawDiskVer2BackingInfo": func() interface{} { return &VirtualDiskPartitionedRawDiskVer2BackingInfo{} },

		"VirtualDiskPartitionedRawDiskVer2BackingOption": func() interface{} { return &VirtualDiskPartitionedRawDiskVer2BackingOption{} },

		"VirtualDiskRawDiskMappingVer1BackingInfo": func() interface{} { return &VirtualDiskRawDiskMappingVer1BackingInfo{} },

		"VirtualDiskRawDiskMappingVer1BackingOption": func() interface{} { return &VirtualDiskRawDiskMappingVer1BackingOption{} },

		"VirtualDiskRawDiskVer2BackingInfo": func() interface{} { return &VirtualDiskRawDiskVer2BackingInfo{} },

		"VirtualDiskRawDiskVer2BackingOption": func() interface{} { return &VirtualDiskRawDiskVer2BackingOption{} },

		"VirtualDiskSeSparseBackingInfo": func() interface{} { return &VirtualDiskSeSparseBackingInfo{} },

		"VirtualDiskSeSparseBackingOption": func() interface{} { return &VirtualDiskSeSparseBackingOption{} },

		"VirtualDiskSparseVer1BackingInfo": func() interface{} { return &VirtualDiskSparseVer1BackingInfo{} },

		"VirtualDiskSparseVer1BackingOption": func() interface{} { return &VirtualDiskSparseVer1BackingOption{} },

		"VirtualDiskSparseVer2BackingInfo": func() interface{} { return &VirtualDiskSparseVer2BackingInfo{} },

		"VirtualDiskSparseVer2BackingOption": func() interface{} { return &VirtualDiskSparseVer2BackingOption{} },

		"VirtualDiskSpec": func() interface{} { return &VirtualDiskSpec{} },

		"VirtualDiskType": func() interface{} { return &VirtualDiskType{} },

		"VirtualDiskVFlashCacheConfigInfo": func() interface{} { return &VirtualDiskVFlashCacheConfigInfo{} },

		"VirtualDiskVFlashCacheConfigInfoCacheConsistencyType": func() interface{} { return &VirtualDiskVFlashCacheConfigInfoCacheConsistencyType{} },

		"VirtualDiskVFlashCacheConfigInfoCacheMode": func() interface{} { return &VirtualDiskVFlashCacheConfigInfoCacheMode{} },

		"VirtualE1000": func() interface{} { return &VirtualE1000{} },

		"VirtualE1000Option": func() interface{} { return &VirtualE1000Option{} },

		"VirtualE1000e": func() interface{} { return &VirtualE1000e{} },

		"VirtualE1000eOption": func() interface{} { return &VirtualE1000eOption{} },

		"VirtualEnsoniq1371": func() interface{} { return &VirtualEnsoniq1371{} },

		"VirtualEnsoniq1371Option": func() interface{} { return &VirtualEnsoniq1371Option{} },

		"VirtualEthernetCard": func() interface{} { return &VirtualEthernetCard{} },

		"VirtualEthernetCardDVPortBackingOption": func() interface{} { return &VirtualEthernetCardDVPortBackingOption{} },

		"VirtualEthernetCardDistributedVirtualPortBackingInfo": func() interface{} { return &VirtualEthernetCardDistributedVirtualPortBackingInfo{} },

		"VirtualEthernetCardLegacyNetworkBackingInfo": func() interface{} { return &VirtualEthernetCardLegacyNetworkBackingInfo{} },

		"VirtualEthernetCardLegacyNetworkBackingOption": func() interface{} { return &VirtualEthernetCardLegacyNetworkBackingOption{} },

		"VirtualEthernetCardLegacyNetworkDeviceName": func() interface{} { return &VirtualEthernetCardLegacyNetworkDeviceName{} },

		"VirtualEthernetCardMacType": func() interface{} { return &VirtualEthernetCardMacType{} },

		"VirtualEthernetCardNetworkBackingInfo": func() interface{} { return &VirtualEthernetCardNetworkBackingInfo{} },

		"VirtualEthernetCardNetworkBackingOption": func() interface{} { return &VirtualEthernetCardNetworkBackingOption{} },

		"VirtualEthernetCardNotSupported": func() interface{} { return &VirtualEthernetCardNotSupported{} },

		"VirtualEthernetCardOpaqueNetworkBackingInfo": func() interface{} { return &VirtualEthernetCardOpaqueNetworkBackingInfo{} },

		"VirtualEthernetCardOpaqueNetworkBackingOption": func() interface{} { return &VirtualEthernetCardOpaqueNetworkBackingOption{} },

		"VirtualEthernetCardOption": func() interface{} { return &VirtualEthernetCardOption{} },

		"VirtualFloppy": func() interface{} { return &VirtualFloppy{} },

		"VirtualFloppyDeviceBackingInfo": func() interface{} { return &VirtualFloppyDeviceBackingInfo{} },

		"VirtualFloppyDeviceBackingOption": func() interface{} { return &VirtualFloppyDeviceBackingOption{} },

		"VirtualFloppyImageBackingInfo": func() interface{} { return &VirtualFloppyImageBackingInfo{} },

		"VirtualFloppyImageBackingOption": func() interface{} { return &VirtualFloppyImageBackingOption{} },

		"VirtualFloppyOption": func() interface{} { return &VirtualFloppyOption{} },

		"VirtualFloppyRemoteDeviceBackingInfo": func() interface{} { return &VirtualFloppyRemoteDeviceBackingInfo{} },

		"VirtualFloppyRemoteDeviceBackingOption": func() interface{} { return &VirtualFloppyRemoteDeviceBackingOption{} },

		"VirtualHardware": func() interface{} { return &VirtualHardware{} },

		"VirtualHardwareCompatibilityIssue": func() interface{} { return &VirtualHardwareCompatibilityIssue{} },

		"VirtualHardwareOption": func() interface{} { return &VirtualHardwareOption{} },

		"VirtualHardwareVersionNotSupported": func() interface{} { return &VirtualHardwareVersionNotSupported{} },

		"VirtualHdAudioCard": func() interface{} { return &VirtualHdAudioCard{} },

		"VirtualHdAudioCardOption": func() interface{} { return &VirtualHdAudioCardOption{} },

		"VirtualIDEController": func() interface{} { return &VirtualIDEController{} },

		"VirtualIDEControllerOption": func() interface{} { return &VirtualIDEControllerOption{} },

		"VirtualKeyboard": func() interface{} { return &VirtualKeyboard{} },

		"VirtualKeyboardOption": func() interface{} { return &VirtualKeyboardOption{} },

		"VirtualLsiLogicController": func() interface{} { return &VirtualLsiLogicController{} },

		"VirtualLsiLogicControllerOption": func() interface{} { return &VirtualLsiLogicControllerOption{} },

		"VirtualLsiLogicSASController": func() interface{} { return &VirtualLsiLogicSASController{} },

		"VirtualLsiLogicSASControllerOption": func() interface{} { return &VirtualLsiLogicSASControllerOption{} },

		"VirtualMachine": func() interface{} { return &VirtualMachine{} },

		"VirtualMachineAffinityInfo": func() interface{} { return &VirtualMachineAffinityInfo{} },

		"VirtualMachineAppHeartbeatStatusType": func() interface{} { return &VirtualMachineAppHeartbeatStatusType{} },

		"VirtualMachineBootOptions": func() interface{} { return &VirtualMachineBootOptions{} },

		"VirtualMachineBootOptionsBootableCdromDevice": func() interface{} { return &VirtualMachineBootOptionsBootableCdromDevice{} },

		"VirtualMachineBootOptionsBootableDevice": func() interface{} { return &VirtualMachineBootOptionsBootableDevice{} },

		"VirtualMachineBootOptionsBootableDiskDevice": func() interface{} { return &VirtualMachineBootOptionsBootableDiskDevice{} },

		"VirtualMachineBootOptionsBootableEthernetDevice": func() interface{} { return &VirtualMachineBootOptionsBootableEthernetDevice{} },

		"VirtualMachineBootOptionsBootableFloppyDevice": func() interface{} { return &VirtualMachineBootOptionsBootableFloppyDevice{} },

		"VirtualMachineCapability": func() interface{} { return &VirtualMachineCapability{} },

		"VirtualMachineCdromInfo": func() interface{} { return &VirtualMachineCdromInfo{} },

		"VirtualMachineCloneSpec": func() interface{} { return &VirtualMachineCloneSpec{} },

		"VirtualMachineCompatibilityChecker": func() interface{} { return &VirtualMachineCompatibilityChecker{} },

		"VirtualMachineConfigInfo": func() interface{} { return &VirtualMachineConfigInfo{} },

		"VirtualMachineConfigInfoDatastoreUrlPair": func() interface{} { return &VirtualMachineConfigInfoDatastoreUrlPair{} },

		"VirtualMachineConfigInfoNpivWwnType": func() interface{} { return &VirtualMachineConfigInfoNpivWwnType{} },

		"VirtualMachineConfigInfoOverheadInfo": func() interface{} { return &VirtualMachineConfigInfoOverheadInfo{} },

		"VirtualMachineConfigInfoSwapPlacementType": func() interface{} { return &VirtualMachineConfigInfoSwapPlacementType{} },

		"VirtualMachineConfigOption": func() interface{} { return &VirtualMachineConfigOption{} },

		"VirtualMachineConfigOptionDescriptor": func() interface{} { return &VirtualMachineConfigOptionDescriptor{} },

		"VirtualMachineConfigSpec": func() interface{} { return &VirtualMachineConfigSpec{} },

		"VirtualMachineConfigSpecNpivWwnOp": func() interface{} { return &VirtualMachineConfigSpecNpivWwnOp{} },

		"VirtualMachineConfigSummary": func() interface{} { return &VirtualMachineConfigSummary{} },

		"VirtualMachineConnectionState": func() interface{} { return &VirtualMachineConnectionState{} },

		"VirtualMachineConsolePreferences": func() interface{} { return &VirtualMachineConsolePreferences{} },

		"VirtualMachineCpuIdInfoSpec": func() interface{} { return &VirtualMachineCpuIdInfoSpec{} },

		"VirtualMachineDatastoreInfo": func() interface{} { return &VirtualMachineDatastoreInfo{} },

		"VirtualMachineDatastoreVolumeOption": func() interface{} { return &VirtualMachineDatastoreVolumeOption{} },

		"VirtualMachineDefaultPowerOpInfo": func() interface{} { return &VirtualMachineDefaultPowerOpInfo{} },

		"VirtualMachineDefinedProfileSpec": func() interface{} { return &VirtualMachineDefinedProfileSpec{} },

		"VirtualMachineDeviceRuntimeInfo": func() interface{} { return &VirtualMachineDeviceRuntimeInfo{} },

		"VirtualMachineDeviceRuntimeInfoDeviceRuntimeState": func() interface{} { return &VirtualMachineDeviceRuntimeInfoDeviceRuntimeState{} },

		"VirtualMachineDeviceRuntimeInfoVirtualEthernetCardRuntimeState": func() interface{} { return &VirtualMachineDeviceRuntimeInfoVirtualEthernetCardRuntimeState{} },

		"VirtualMachineDeviceRuntimeInfoVirtualEthernetCardRuntimeStateVmDirectPathGen2InactiveReasonOther": func() interface{} {
			return &VirtualMachineDeviceRuntimeInfoVirtualEthernetCardRuntimeStateVmDirectPathGen2InactiveReasonOther{}
		},

		"VirtualMachineDeviceRuntimeInfoVirtualEthernetCardRuntimeStateVmDirectPathGen2InactiveReasonVm": func() interface{} {
			return &VirtualMachineDeviceRuntimeInfoVirtualEthernetCardRuntimeStateVmDirectPathGen2InactiveReasonVm{}
		},

		"VirtualMachineDiskDeviceInfo": func() interface{} { return &VirtualMachineDiskDeviceInfo{} },

		"VirtualMachineDisplayTopology": func() interface{} { return &VirtualMachineDisplayTopology{} },

		"VirtualMachineEmptyProfileSpec": func() interface{} { return &VirtualMachineEmptyProfileSpec{} },

		"VirtualMachineFaultToleranceState": func() interface{} { return &VirtualMachineFaultToleranceState{} },

		"VirtualMachineFeatureRequirement": func() interface{} { return &VirtualMachineFeatureRequirement{} },

		"VirtualMachineFileInfo": func() interface{} { return &VirtualMachineFileInfo{} },

		"VirtualMachineFileLayout": func() interface{} { return &VirtualMachineFileLayout{} },

		"VirtualMachineFileLayoutDiskLayout": func() interface{} { return &VirtualMachineFileLayoutDiskLayout{} },

		"VirtualMachineFileLayoutEx": func() interface{} { return &VirtualMachineFileLayoutEx{} },

		"VirtualMachineFileLayoutExDiskLayout": func() interface{} { return &VirtualMachineFileLayoutExDiskLayout{} },

		"VirtualMachineFileLayoutExDiskUnit": func() interface{} { return &VirtualMachineFileLayoutExDiskUnit{} },

		"VirtualMachineFileLayoutExFileInfo": func() interface{} { return &VirtualMachineFileLayoutExFileInfo{} },

		"VirtualMachineFileLayoutExFileType": func() interface{} { return &VirtualMachineFileLayoutExFileType{} },

		"VirtualMachineFileLayoutExSnapshotLayout": func() interface{} { return &VirtualMachineFileLayoutExSnapshotLayout{} },

		"VirtualMachineFileLayoutSnapshotLayout": func() interface{} { return &VirtualMachineFileLayoutSnapshotLayout{} },

		"VirtualMachineFlagInfo": func() interface{} { return &VirtualMachineFlagInfo{} },

		"VirtualMachineFlagInfoMonitorType": func() interface{} { return &VirtualMachineFlagInfoMonitorType{} },

		"VirtualMachineFlagInfoVirtualExecUsage": func() interface{} { return &VirtualMachineFlagInfoVirtualExecUsage{} },

		"VirtualMachineFlagInfoVirtualMmuUsage": func() interface{} { return &VirtualMachineFlagInfoVirtualMmuUsage{} },

		"VirtualMachineFloppyInfo": func() interface{} { return &VirtualMachineFloppyInfo{} },

		"VirtualMachineGuestOsFamily": func() interface{} { return &VirtualMachineGuestOsFamily{} },

		"VirtualMachineGuestOsIdentifier": func() interface{} { return &VirtualMachineGuestOsIdentifier{} },

		"VirtualMachineGuestState": func() interface{} { return &VirtualMachineGuestState{} },

		"VirtualMachineGuestSummary": func() interface{} { return &VirtualMachineGuestSummary{} },

		"VirtualMachineHtSharing": func() interface{} { return &VirtualMachineHtSharing{} },

		"VirtualMachineIdeDiskDeviceInfo": func() interface{} { return &VirtualMachineIdeDiskDeviceInfo{} },

		"VirtualMachineIdeDiskDevicePartitionInfo": func() interface{} { return &VirtualMachineIdeDiskDevicePartitionInfo{} },

		"VirtualMachineImportSpec": func() interface{} { return &VirtualMachineImportSpec{} },

		"VirtualMachineLegacyNetworkSwitchInfo": func() interface{} { return &VirtualMachineLegacyNetworkSwitchInfo{} },

		"VirtualMachineMemoryAllocationPolicy": func() interface{} { return &VirtualMachineMemoryAllocationPolicy{} },

		"VirtualMachineMemoryReservationInfo": func() interface{} { return &VirtualMachineMemoryReservationInfo{} },

		"VirtualMachineMemoryReservationSpec": func() interface{} { return &VirtualMachineMemoryReservationSpec{} },

		"VirtualMachineMessage": func() interface{} { return &VirtualMachineMessage{} },

		"VirtualMachineMetadataManagerVmMetadata": func() interface{} { return &VirtualMachineMetadataManagerVmMetadata{} },

		"VirtualMachineMetadataManagerVmMetadataInput": func() interface{} { return &VirtualMachineMetadataManagerVmMetadataInput{} },

		"VirtualMachineMetadataManagerVmMetadataOp": func() interface{} { return &VirtualMachineMetadataManagerVmMetadataOp{} },

		"VirtualMachineMetadataManagerVmMetadataOwner": func() interface{} { return &VirtualMachineMetadataManagerVmMetadataOwner{} },

		"VirtualMachineMetadataManagerVmMetadataOwnerOwner": func() interface{} { return &VirtualMachineMetadataManagerVmMetadataOwnerOwner{} },

		"VirtualMachineMetadataManagerVmMetadataResult": func() interface{} { return &VirtualMachineMetadataManagerVmMetadataResult{} },

		"VirtualMachineMksTicket": func() interface{} { return &VirtualMachineMksTicket{} },

		"VirtualMachineMovePriority": func() interface{} { return &VirtualMachineMovePriority{} },

		"VirtualMachineNeedSecondaryReason": func() interface{} { return &VirtualMachineNeedSecondaryReason{} },

		"VirtualMachineNetworkInfo": func() interface{} { return &VirtualMachineNetworkInfo{} },

		"VirtualMachineNetworkShaperInfo": func() interface{} { return &VirtualMachineNetworkShaperInfo{} },

		"VirtualMachineParallelInfo": func() interface{} { return &VirtualMachineParallelInfo{} },

		"VirtualMachinePciPassthroughInfo": func() interface{} { return &VirtualMachinePciPassthroughInfo{} },

		"VirtualMachinePowerOffBehavior": func() interface{} { return &VirtualMachinePowerOffBehavior{} },

		"VirtualMachinePowerOpType": func() interface{} { return &VirtualMachinePowerOpType{} },

		"VirtualMachinePowerState": func() interface{} { return &VirtualMachinePowerState{} },

		"VirtualMachineProfileRawData": func() interface{} { return &VirtualMachineProfileRawData{} },

		"VirtualMachineProfileSpec": func() interface{} { return &VirtualMachineProfileSpec{} },

		"VirtualMachineProvisioningChecker": func() interface{} { return &VirtualMachineProvisioningChecker{} },

		"VirtualMachineQuestionInfo": func() interface{} { return &VirtualMachineQuestionInfo{} },

		"VirtualMachineQuickStats": func() interface{} { return &VirtualMachineQuickStats{} },

		"VirtualMachineRecordReplayState": func() interface{} { return &VirtualMachineRecordReplayState{} },

		"VirtualMachineRelocateDiskMoveOptions": func() interface{} { return &VirtualMachineRelocateDiskMoveOptions{} },

		"VirtualMachineRelocateSpec": func() interface{} { return &VirtualMachineRelocateSpec{} },

		"VirtualMachineRelocateSpecDiskLocator": func() interface{} { return &VirtualMachineRelocateSpecDiskLocator{} },

		"VirtualMachineRelocateTransformation": func() interface{} { return &VirtualMachineRelocateTransformation{} },

		"VirtualMachineRuntimeInfo": func() interface{} { return &VirtualMachineRuntimeInfo{} },

		"VirtualMachineRuntimeInfoDasProtectionState": func() interface{} { return &VirtualMachineRuntimeInfoDasProtectionState{} },

		"VirtualMachineScsiDiskDeviceInfo": func() interface{} { return &VirtualMachineScsiDiskDeviceInfo{} },

		"VirtualMachineScsiPassthroughInfo": func() interface{} { return &VirtualMachineScsiPassthroughInfo{} },

		"VirtualMachineScsiPassthroughType": func() interface{} { return &VirtualMachineScsiPassthroughType{} },

		"VirtualMachineSerialInfo": func() interface{} { return &VirtualMachineSerialInfo{} },

		"VirtualMachineSnapshot": func() interface{} { return &VirtualMachineSnapshot{} },

		"VirtualMachineSnapshotInfo": func() interface{} { return &VirtualMachineSnapshotInfo{} },

		"VirtualMachineSnapshotTree": func() interface{} { return &VirtualMachineSnapshotTree{} },

		"VirtualMachineSoundInfo": func() interface{} { return &VirtualMachineSoundInfo{} },

		"VirtualMachineSriovInfo": func() interface{} { return &VirtualMachineSriovInfo{} },

		"VirtualMachineStandbyActionType": func() interface{} { return &VirtualMachineStandbyActionType{} },

		"VirtualMachineStorageInfo": func() interface{} { return &VirtualMachineStorageInfo{} },

		"VirtualMachineStorageSummary": func() interface{} { return &VirtualMachineStorageSummary{} },

		"VirtualMachineSummary": func() interface{} { return &VirtualMachineSummary{} },

		"VirtualMachineTargetInfo": func() interface{} { return &VirtualMachineTargetInfo{} },

		"VirtualMachineTargetInfoConfigurationTag": func() interface{} { return &VirtualMachineTargetInfoConfigurationTag{} },

		"VirtualMachineTicket": func() interface{} { return &VirtualMachineTicket{} },

		"VirtualMachineTicketType": func() interface{} { return &VirtualMachineTicketType{} },

		"VirtualMachineToolsRunningStatus": func() interface{} { return &VirtualMachineToolsRunningStatus{} },

		"VirtualMachineToolsStatus": func() interface{} { return &VirtualMachineToolsStatus{} },

		"VirtualMachineToolsVersionStatus": func() interface{} { return &VirtualMachineToolsVersionStatus{} },

		"VirtualMachineUsageOnDatastore": func() interface{} { return &VirtualMachineUsageOnDatastore{} },

		"VirtualMachineUsbInfo": func() interface{} { return &VirtualMachineUsbInfo{} },

		"VirtualMachineUsbInfoFamily": func() interface{} { return &VirtualMachineUsbInfoFamily{} },

		"VirtualMachineUsbInfoSpeed": func() interface{} { return &VirtualMachineUsbInfoSpeed{} },

		"VirtualMachineVFlashModuleInfo": func() interface{} { return &VirtualMachineVFlashModuleInfo{} },

		"VirtualMachineVMCIDevice": func() interface{} { return &VirtualMachineVMCIDevice{} },

		"VirtualMachineVMCIDeviceOption": func() interface{} { return &VirtualMachineVMCIDeviceOption{} },

		"VirtualMachineVMIROM": func() interface{} { return &VirtualMachineVMIROM{} },

		"VirtualMachineVideoCard": func() interface{} { return &VirtualMachineVideoCard{} },

		"VirtualMachineVideoCardUse3dRenderer": func() interface{} { return &VirtualMachineVideoCardUse3dRenderer{} },

		"VirtualMachineWipeResult": func() interface{} { return &VirtualMachineWipeResult{} },

		"VirtualNicManagerNetConfig": func() interface{} { return &VirtualNicManagerNetConfig{} },

		"VirtualPCIController": func() interface{} { return &VirtualPCIController{} },

		"VirtualPCIControllerOption": func() interface{} { return &VirtualPCIControllerOption{} },

		"VirtualPCIPassthrough": func() interface{} { return &VirtualPCIPassthrough{} },

		"VirtualPCIPassthroughDeviceBackingInfo": func() interface{} { return &VirtualPCIPassthroughDeviceBackingInfo{} },

		"VirtualPCIPassthroughDeviceBackingOption": func() interface{} { return &VirtualPCIPassthroughDeviceBackingOption{} },

		"VirtualPCIPassthroughOption": func() interface{} { return &VirtualPCIPassthroughOption{} },

		"VirtualPCNet32": func() interface{} { return &VirtualPCNet32{} },

		"VirtualPCNet32Option": func() interface{} { return &VirtualPCNet32Option{} },

		"VirtualPS2Controller": func() interface{} { return &VirtualPS2Controller{} },

		"VirtualPS2ControllerOption": func() interface{} { return &VirtualPS2ControllerOption{} },

		"VirtualParallelPort": func() interface{} { return &VirtualParallelPort{} },

		"VirtualParallelPortDeviceBackingInfo": func() interface{} { return &VirtualParallelPortDeviceBackingInfo{} },

		"VirtualParallelPortDeviceBackingOption": func() interface{} { return &VirtualParallelPortDeviceBackingOption{} },

		"VirtualParallelPortFileBackingInfo": func() interface{} { return &VirtualParallelPortFileBackingInfo{} },

		"VirtualParallelPortFileBackingOption": func() interface{} { return &VirtualParallelPortFileBackingOption{} },

		"VirtualParallelPortOption": func() interface{} { return &VirtualParallelPortOption{} },

		"VirtualPointingDevice": func() interface{} { return &VirtualPointingDevice{} },

		"VirtualPointingDeviceBackingOption": func() interface{} { return &VirtualPointingDeviceBackingOption{} },

		"VirtualPointingDeviceDeviceBackingInfo": func() interface{} { return &VirtualPointingDeviceDeviceBackingInfo{} },

		"VirtualPointingDeviceHostChoice": func() interface{} { return &VirtualPointingDeviceHostChoice{} },

		"VirtualPointingDeviceOption": func() interface{} { return &VirtualPointingDeviceOption{} },

		"VirtualSATAController": func() interface{} { return &VirtualSATAController{} },

		"VirtualSATAControllerOption": func() interface{} { return &VirtualSATAControllerOption{} },

		"VirtualSCSIController": func() interface{} { return &VirtualSCSIController{} },

		"VirtualSCSIControllerOption": func() interface{} { return &VirtualSCSIControllerOption{} },

		"VirtualSCSIPassthrough": func() interface{} { return &VirtualSCSIPassthrough{} },

		"VirtualSCSIPassthroughDeviceBackingInfo": func() interface{} { return &VirtualSCSIPassthroughDeviceBackingInfo{} },

		"VirtualSCSIPassthroughDeviceBackingOption": func() interface{} { return &VirtualSCSIPassthroughDeviceBackingOption{} },

		"VirtualSCSIPassthroughOption": func() interface{} { return &VirtualSCSIPassthroughOption{} },

		"VirtualSCSISharing": func() interface{} { return &VirtualSCSISharing{} },

		"VirtualSIOController": func() interface{} { return &VirtualSIOController{} },

		"VirtualSIOControllerOption": func() interface{} { return &VirtualSIOControllerOption{} },

		"VirtualSerialPort": func() interface{} { return &VirtualSerialPort{} },

		"VirtualSerialPortDeviceBackingInfo": func() interface{} { return &VirtualSerialPortDeviceBackingInfo{} },

		"VirtualSerialPortDeviceBackingOption": func() interface{} { return &VirtualSerialPortDeviceBackingOption{} },

		"VirtualSerialPortEndPoint": func() interface{} { return &VirtualSerialPortEndPoint{} },

		"VirtualSerialPortFileBackingInfo": func() interface{} { return &VirtualSerialPortFileBackingInfo{} },

		"VirtualSerialPortFileBackingOption": func() interface{} { return &VirtualSerialPortFileBackingOption{} },

		"VirtualSerialPortOption": func() interface{} { return &VirtualSerialPortOption{} },

		"VirtualSerialPortPipeBackingInfo": func() interface{} { return &VirtualSerialPortPipeBackingInfo{} },

		"VirtualSerialPortPipeBackingOption": func() interface{} { return &VirtualSerialPortPipeBackingOption{} },

		"VirtualSerialPortThinPrintBackingInfo": func() interface{} { return &VirtualSerialPortThinPrintBackingInfo{} },

		"VirtualSerialPortThinPrintBackingOption": func() interface{} { return &VirtualSerialPortThinPrintBackingOption{} },

		"VirtualSerialPortURIBackingInfo": func() interface{} { return &VirtualSerialPortURIBackingInfo{} },

		"VirtualSerialPortURIBackingOption": func() interface{} { return &VirtualSerialPortURIBackingOption{} },

		"VirtualSoundBlaster16": func() interface{} { return &VirtualSoundBlaster16{} },

		"VirtualSoundBlaster16Option": func() interface{} { return &VirtualSoundBlaster16Option{} },

		"VirtualSoundCard": func() interface{} { return &VirtualSoundCard{} },

		"VirtualSoundCardDeviceBackingInfo": func() interface{} { return &VirtualSoundCardDeviceBackingInfo{} },

		"VirtualSoundCardDeviceBackingOption": func() interface{} { return &VirtualSoundCardDeviceBackingOption{} },

		"VirtualSoundCardOption": func() interface{} { return &VirtualSoundCardOption{} },

		"VirtualSriovEthernetCard": func() interface{} { return &VirtualSriovEthernetCard{} },

		"VirtualSriovEthernetCardOption": func() interface{} { return &VirtualSriovEthernetCardOption{} },

		"VirtualSriovEthernetCardSriovBackingInfo": func() interface{} { return &VirtualSriovEthernetCardSriovBackingInfo{} },

		"VirtualSriovEthernetCardSriovBackingOption": func() interface{} { return &VirtualSriovEthernetCardSriovBackingOption{} },

		"VirtualSwitchProfile": func() interface{} { return &VirtualSwitchProfile{} },

		"VirtualSwitchSelectionProfile": func() interface{} { return &VirtualSwitchSelectionProfile{} },

		"VirtualUSB": func() interface{} { return &VirtualUSB{} },

		"VirtualUSBController": func() interface{} { return &VirtualUSBController{} },

		"VirtualUSBControllerOption": func() interface{} { return &VirtualUSBControllerOption{} },

		"VirtualUSBControllerPciBusSlotInfo": func() interface{} { return &VirtualUSBControllerPciBusSlotInfo{} },

		"VirtualUSBOption": func() interface{} { return &VirtualUSBOption{} },

		"VirtualUSBRemoteClientBackingInfo": func() interface{} { return &VirtualUSBRemoteClientBackingInfo{} },

		"VirtualUSBRemoteClientBackingOption": func() interface{} { return &VirtualUSBRemoteClientBackingOption{} },

		"VirtualUSBRemoteHostBackingInfo": func() interface{} { return &VirtualUSBRemoteHostBackingInfo{} },

		"VirtualUSBRemoteHostBackingOption": func() interface{} { return &VirtualUSBRemoteHostBackingOption{} },

		"VirtualUSBUSBBackingInfo": func() interface{} { return &VirtualUSBUSBBackingInfo{} },

		"VirtualUSBUSBBackingOption": func() interface{} { return &VirtualUSBUSBBackingOption{} },

		"VirtualUSBXHCIController": func() interface{} { return &VirtualUSBXHCIController{} },

		"VirtualUSBXHCIControllerOption": func() interface{} { return &VirtualUSBXHCIControllerOption{} },

		"VirtualVMIROMOption": func() interface{} { return &VirtualVMIROMOption{} },

		"VirtualVideoCardOption": func() interface{} { return &VirtualVideoCardOption{} },

		"VirtualVmxnet": func() interface{} { return &VirtualVmxnet{} },

		"VirtualVmxnet2": func() interface{} { return &VirtualVmxnet2{} },

		"VirtualVmxnet2Option": func() interface{} { return &VirtualVmxnet2Option{} },

		"VirtualVmxnet3": func() interface{} { return &VirtualVmxnet3{} },

		"VirtualVmxnet3Option": func() interface{} { return &VirtualVmxnet3Option{} },

		"VirtualVmxnetOption": func() interface{} { return &VirtualVmxnetOption{} },

		"VirtualizationManager": func() interface{} { return &VirtualizationManager{} },

		"VlanProfile": func() interface{} { return &VlanProfile{} },

		"VmAcquiredMksTicketEvent": func() interface{} { return &VmAcquiredMksTicketEvent{} },

		"VmAcquiredTicketEvent": func() interface{} { return &VmAcquiredTicketEvent{} },

		"VmAlreadyExistsInDatacenter": func() interface{} { return &VmAlreadyExistsInDatacenter{} },

		"VmAutoRenameEvent": func() interface{} { return &VmAutoRenameEvent{} },

		"VmBeingClonedEvent": func() interface{} { return &VmBeingClonedEvent{} },

		"VmBeingClonedNoFolderEvent": func() interface{} { return &VmBeingClonedNoFolderEvent{} },

		"VmBeingCreatedEvent": func() interface{} { return &VmBeingCreatedEvent{} },

		"VmBeingDeployedEvent": func() interface{} { return &VmBeingDeployedEvent{} },

		"VmBeingHotMigratedEvent": func() interface{} { return &VmBeingHotMigratedEvent{} },

		"VmBeingMigratedEvent": func() interface{} { return &VmBeingMigratedEvent{} },

		"VmBeingRelocatedEvent": func() interface{} { return &VmBeingRelocatedEvent{} },

		"VmCloneEvent": func() interface{} { return &VmCloneEvent{} },

		"VmCloneFailedEvent": func() interface{} { return &VmCloneFailedEvent{} },

		"VmClonedEvent": func() interface{} { return &VmClonedEvent{} },

		"VmConfigFault": func() interface{} { return &VmConfigFault{} },

		"VmConfigFileInfo": func() interface{} { return &VmConfigFileInfo{} },

		"VmConfigFileQuery": func() interface{} { return &VmConfigFileQuery{} },

		"VmConfigFileQueryFilter": func() interface{} { return &VmConfigFileQueryFilter{} },

		"VmConfigFileQueryFlags": func() interface{} { return &VmConfigFileQueryFlags{} },

		"VmConfigIncompatibleForFaultTolerance": func() interface{} { return &VmConfigIncompatibleForFaultTolerance{} },

		"VmConfigIncompatibleForRecordReplay": func() interface{} { return &VmConfigIncompatibleForRecordReplay{} },

		"VmConfigInfo": func() interface{} { return &VmConfigInfo{} },

		"VmConfigMissingEvent": func() interface{} { return &VmConfigMissingEvent{} },

		"VmConfigSpec": func() interface{} { return &VmConfigSpec{} },

		"VmConnectedEvent": func() interface{} { return &VmConnectedEvent{} },

		"VmCreatedEvent": func() interface{} { return &VmCreatedEvent{} },

		"VmDasBeingResetEvent": func() interface{} { return &VmDasBeingResetEvent{} },

		"VmDasBeingResetEventReasonCode": func() interface{} { return &VmDasBeingResetEventReasonCode{} },

		"VmDasBeingResetWithScreenshotEvent": func() interface{} { return &VmDasBeingResetWithScreenshotEvent{} },

		"VmDasResetFailedEvent": func() interface{} { return &VmDasResetFailedEvent{} },

		"VmDasUpdateErrorEvent": func() interface{} { return &VmDasUpdateErrorEvent{} },

		"VmDasUpdateOkEvent": func() interface{} { return &VmDasUpdateOkEvent{} },

		"VmDateRolledBackEvent": func() interface{} { return &VmDateRolledBackEvent{} },

		"VmDeployFailedEvent": func() interface{} { return &VmDeployFailedEvent{} },

		"VmDeployedEvent": func() interface{} { return &VmDeployedEvent{} },

		"VmDisconnectedEvent": func() interface{} { return &VmDisconnectedEvent{} },

		"VmDiscoveredEvent": func() interface{} { return &VmDiscoveredEvent{} },

		"VmDiskFailedEvent": func() interface{} { return &VmDiskFailedEvent{} },

		"VmDiskFileInfo": func() interface{} { return &VmDiskFileInfo{} },

		"VmDiskFileQuery": func() interface{} { return &VmDiskFileQuery{} },

		"VmDiskFileQueryFilter": func() interface{} { return &VmDiskFileQueryFilter{} },

		"VmDiskFileQueryFlags": func() interface{} { return &VmDiskFileQueryFlags{} },

		"VmEmigratingEvent": func() interface{} { return &VmEmigratingEvent{} },

		"VmEndRecordingEvent": func() interface{} { return &VmEndRecordingEvent{} },

		"VmEndReplayingEvent": func() interface{} { return &VmEndReplayingEvent{} },

		"VmEvent": func() interface{} { return &VmEvent{} },

		"VmEventArgument": func() interface{} { return &VmEventArgument{} },

		"VmFailedMigrateEvent": func() interface{} { return &VmFailedMigrateEvent{} },

		"VmFailedRelayoutEvent": func() interface{} { return &VmFailedRelayoutEvent{} },

		"VmFailedRelayoutOnVmfs2DatastoreEvent": func() interface{} { return &VmFailedRelayoutOnVmfs2DatastoreEvent{} },

		"VmFailedStartingSecondaryEvent": func() interface{} { return &VmFailedStartingSecondaryEvent{} },

		"VmFailedStartingSecondaryEventFailureReason": func() interface{} { return &VmFailedStartingSecondaryEventFailureReason{} },

		"VmFailedToPowerOffEvent": func() interface{} { return &VmFailedToPowerOffEvent{} },

		"VmFailedToPowerOnEvent": func() interface{} { return &VmFailedToPowerOnEvent{} },

		"VmFailedToRebootGuestEvent": func() interface{} { return &VmFailedToRebootGuestEvent{} },

		"VmFailedToResetEvent": func() interface{} { return &VmFailedToResetEvent{} },

		"VmFailedToShutdownGuestEvent": func() interface{} { return &VmFailedToShutdownGuestEvent{} },

		"VmFailedToStandbyGuestEvent": func() interface{} { return &VmFailedToStandbyGuestEvent{} },

		"VmFailedToSuspendEvent": func() interface{} { return &VmFailedToSuspendEvent{} },

		"VmFailedUpdatingSecondaryConfig": func() interface{} { return &VmFailedUpdatingSecondaryConfig{} },

		"VmFailoverFailed": func() interface{} { return &VmFailoverFailed{} },

		"VmFaultToleranceConfigIssue": func() interface{} { return &VmFaultToleranceConfigIssue{} },

		"VmFaultToleranceConfigIssueReasonForIssue": func() interface{} { return &VmFaultToleranceConfigIssueReasonForIssue{} },

		"VmFaultToleranceConfigIssueWrapper": func() interface{} { return &VmFaultToleranceConfigIssueWrapper{} },

		"VmFaultToleranceInvalidFileBacking": func() interface{} { return &VmFaultToleranceInvalidFileBacking{} },

		"VmFaultToleranceInvalidFileBackingDeviceType": func() interface{} { return &VmFaultToleranceInvalidFileBackingDeviceType{} },

		"VmFaultToleranceIssue": func() interface{} { return &VmFaultToleranceIssue{} },

		"VmFaultToleranceOpIssuesList": func() interface{} { return &VmFaultToleranceOpIssuesList{} },

		"VmFaultToleranceStateChangedEvent": func() interface{} { return &VmFaultToleranceStateChangedEvent{} },

		"VmFaultToleranceTooManyVMsOnHost": func() interface{} { return &VmFaultToleranceTooManyVMsOnHost{} },

		"VmFaultToleranceTurnedOffEvent": func() interface{} { return &VmFaultToleranceTurnedOffEvent{} },

		"VmFaultToleranceVmTerminatedEvent": func() interface{} { return &VmFaultToleranceVmTerminatedEvent{} },

		"VmGuestRebootEvent": func() interface{} { return &VmGuestRebootEvent{} },

		"VmGuestShutdownEvent": func() interface{} { return &VmGuestShutdownEvent{} },

		"VmGuestStandbyEvent": func() interface{} { return &VmGuestStandbyEvent{} },

		"VmHealthMonitoringStateChangedEvent": func() interface{} { return &VmHealthMonitoringStateChangedEvent{} },

		"VmHostAffinityRuleViolation": func() interface{} { return &VmHostAffinityRuleViolation{} },

		"VmInstanceUuidAssignedEvent": func() interface{} { return &VmInstanceUuidAssignedEvent{} },

		"VmInstanceUuidChangedEvent": func() interface{} { return &VmInstanceUuidChangedEvent{} },

		"VmInstanceUuidConflictEvent": func() interface{} { return &VmInstanceUuidConflictEvent{} },

		"VmLimitLicense": func() interface{} { return &VmLimitLicense{} },

		"VmLogFileInfo": func() interface{} { return &VmLogFileInfo{} },

		"VmLogFileQuery": func() interface{} { return &VmLogFileQuery{} },

		"VmMacAssignedEvent": func() interface{} { return &VmMacAssignedEvent{} },

		"VmMacChangedEvent": func() interface{} { return &VmMacChangedEvent{} },

		"VmMacConflictEvent": func() interface{} { return &VmMacConflictEvent{} },

		"VmMaxFTRestartCountReached": func() interface{} { return &VmMaxFTRestartCountReached{} },

		"VmMaxRestartCountReached": func() interface{} { return &VmMaxRestartCountReached{} },

		"VmMessageErrorEvent": func() interface{} { return &VmMessageErrorEvent{} },

		"VmMessageEvent": func() interface{} { return &VmMessageEvent{} },

		"VmMessageWarningEvent": func() interface{} { return &VmMessageWarningEvent{} },

		"VmMetadataManagerFault": func() interface{} { return &VmMetadataManagerFault{} },

		"VmMigratedEvent": func() interface{} { return &VmMigratedEvent{} },

		"VmMonitorIncompatibleForFaultTolerance": func() interface{} { return &VmMonitorIncompatibleForFaultTolerance{} },

		"VmNoCompatibleHostForSecondaryEvent": func() interface{} { return &VmNoCompatibleHostForSecondaryEvent{} },

		"VmNoNetworkAccessEvent": func() interface{} { return &VmNoNetworkAccessEvent{} },

		"VmNvramFileInfo": func() interface{} { return &VmNvramFileInfo{} },

		"VmNvramFileQuery": func() interface{} { return &VmNvramFileQuery{} },

		"VmOrphanedEvent": func() interface{} { return &VmOrphanedEvent{} },

		"VmPodConfigForPlacement": func() interface{} { return &VmPodConfigForPlacement{} },

		"VmPortGroupProfile": func() interface{} { return &VmPortGroupProfile{} },

		"VmPowerOffOnIsolationEvent": func() interface{} { return &VmPowerOffOnIsolationEvent{} },

		"VmPowerOnDisabled": func() interface{} { return &VmPowerOnDisabled{} },

		"VmPoweredOffEvent": func() interface{} { return &VmPoweredOffEvent{} },

		"VmPoweredOnEvent": func() interface{} { return &VmPoweredOnEvent{} },

		"VmPoweringOnWithCustomizedDVPortEvent": func() interface{} { return &VmPoweringOnWithCustomizedDVPortEvent{} },

		"VmPrimaryFailoverEvent": func() interface{} { return &VmPrimaryFailoverEvent{} },

		"VmReconfiguredEvent": func() interface{} { return &VmReconfiguredEvent{} },

		"VmRegisteredEvent": func() interface{} { return &VmRegisteredEvent{} },

		"VmRelayoutSuccessfulEvent": func() interface{} { return &VmRelayoutSuccessfulEvent{} },

		"VmRelayoutUpToDateEvent": func() interface{} { return &VmRelayoutUpToDateEvent{} },

		"VmReloadFromPathEvent": func() interface{} { return &VmReloadFromPathEvent{} },

		"VmReloadFromPathFailedEvent": func() interface{} { return &VmReloadFromPathFailedEvent{} },

		"VmRelocateFailedEvent": func() interface{} { return &VmRelocateFailedEvent{} },

		"VmRelocateSpecEvent": func() interface{} { return &VmRelocateSpecEvent{} },

		"VmRelocatedEvent": func() interface{} { return &VmRelocatedEvent{} },

		"VmRemoteConsoleConnectedEvent": func() interface{} { return &VmRemoteConsoleConnectedEvent{} },

		"VmRemoteConsoleDisconnectedEvent": func() interface{} { return &VmRemoteConsoleDisconnectedEvent{} },

		"VmRemovedEvent": func() interface{} { return &VmRemovedEvent{} },

		"VmRenamedEvent": func() interface{} { return &VmRenamedEvent{} },

		"VmRequirementsExceedCurrentEVCModeEvent": func() interface{} { return &VmRequirementsExceedCurrentEVCModeEvent{} },

		"VmResettingEvent": func() interface{} { return &VmResettingEvent{} },

		"VmResourcePoolMovedEvent": func() interface{} { return &VmResourcePoolMovedEvent{} },

		"VmResourceReallocatedEvent": func() interface{} { return &VmResourceReallocatedEvent{} },

		"VmRestartedOnAlternateHostEvent": func() interface{} { return &VmRestartedOnAlternateHostEvent{} },

		"VmResumingEvent": func() interface{} { return &VmResumingEvent{} },

		"VmSecondaryAddedEvent": func() interface{} { return &VmSecondaryAddedEvent{} },

		"VmSecondaryDisabledBySystemEvent": func() interface{} { return &VmSecondaryDisabledBySystemEvent{} },

		"VmSecondaryDisabledEvent": func() interface{} { return &VmSecondaryDisabledEvent{} },

		"VmSecondaryEnabledEvent": func() interface{} { return &VmSecondaryEnabledEvent{} },

		"VmSecondaryStartedEvent": func() interface{} { return &VmSecondaryStartedEvent{} },

		"VmShutdownOnIsolationEvent": func() interface{} { return &VmShutdownOnIsolationEvent{} },

		"VmShutdownOnIsolationEventOperation": func() interface{} { return &VmShutdownOnIsolationEventOperation{} },

		"VmSnapshotFileInfo": func() interface{} { return &VmSnapshotFileInfo{} },

		"VmSnapshotFileQuery": func() interface{} { return &VmSnapshotFileQuery{} },

		"VmStartRecordingEvent": func() interface{} { return &VmStartRecordingEvent{} },

		"VmStartReplayingEvent": func() interface{} { return &VmStartReplayingEvent{} },

		"VmStartingEvent": func() interface{} { return &VmStartingEvent{} },

		"VmStartingSecondaryEvent": func() interface{} { return &VmStartingSecondaryEvent{} },

		"VmStaticMacConflictEvent": func() interface{} { return &VmStaticMacConflictEvent{} },

		"VmStoppingEvent": func() interface{} { return &VmStoppingEvent{} },

		"VmSuspendedEvent": func() interface{} { return &VmSuspendedEvent{} },

		"VmSuspendingEvent": func() interface{} { return &VmSuspendingEvent{} },

		"VmTimedoutStartingSecondaryEvent": func() interface{} { return &VmTimedoutStartingSecondaryEvent{} },

		"VmToolsUpgradeFault": func() interface{} { return &VmToolsUpgradeFault{} },

		"VmUnsupportedStartingEvent": func() interface{} { return &VmUnsupportedStartingEvent{} },

		"VmUpgradeCompleteEvent": func() interface{} { return &VmUpgradeCompleteEvent{} },

		"VmUpgradeFailedEvent": func() interface{} { return &VmUpgradeFailedEvent{} },

		"VmUpgradingEvent": func() interface{} { return &VmUpgradingEvent{} },

		"VmUuidAssignedEvent": func() interface{} { return &VmUuidAssignedEvent{} },

		"VmUuidChangedEvent": func() interface{} { return &VmUuidChangedEvent{} },

		"VmUuidConflictEvent": func() interface{} { return &VmUuidConflictEvent{} },

		"VmValidateMaxDevice": func() interface{} { return &VmValidateMaxDevice{} },

		"VmWwnAssignedEvent": func() interface{} { return &VmWwnAssignedEvent{} },

		"VmWwnChangedEvent": func() interface{} { return &VmWwnChangedEvent{} },

		"VmWwnConflict": func() interface{} { return &VmWwnConflict{} },

		"VmWwnConflictEvent": func() interface{} { return &VmWwnConflictEvent{} },

		"VmfsAlreadyMounted": func() interface{} { return &VmfsAlreadyMounted{} },

		"VmfsAmbiguousMount": func() interface{} { return &VmfsAmbiguousMount{} },

		"VmfsDatastoreAllExtentOption": func() interface{} { return &VmfsDatastoreAllExtentOption{} },

		"VmfsDatastoreBaseOption": func() interface{} { return &VmfsDatastoreBaseOption{} },

		"VmfsDatastoreCreateSpec": func() interface{} { return &VmfsDatastoreCreateSpec{} },

		"VmfsDatastoreExpandSpec": func() interface{} { return &VmfsDatastoreExpandSpec{} },

		"VmfsDatastoreExtendSpec": func() interface{} { return &VmfsDatastoreExtendSpec{} },

		"VmfsDatastoreInfo": func() interface{} { return &VmfsDatastoreInfo{} },

		"VmfsDatastoreMultipleExtentOption": func() interface{} { return &VmfsDatastoreMultipleExtentOption{} },

		"VmfsDatastoreOption": func() interface{} { return &VmfsDatastoreOption{} },

		"VmfsDatastoreSingleExtentOption": func() interface{} { return &VmfsDatastoreSingleExtentOption{} },

		"VmfsDatastoreSpec": func() interface{} { return &VmfsDatastoreSpec{} },

		"VmfsMountFault": func() interface{} { return &VmfsMountFault{} },

		"VmotionInterfaceNotEnabled": func() interface{} { return &VmotionInterfaceNotEnabled{} },

		"VmwareDistributedVirtualSwitch": func() interface{} { return &VmwareDistributedVirtualSwitch{} },

		"VmwareDistributedVirtualSwitchPvlanPortType": func() interface{} { return &VmwareDistributedVirtualSwitchPvlanPortType{} },

		"VmwareDistributedVirtualSwitchPvlanSpec": func() interface{} { return &VmwareDistributedVirtualSwitchPvlanSpec{} },

		"VmwareDistributedVirtualSwitchTrunkVlanSpec": func() interface{} { return &VmwareDistributedVirtualSwitchTrunkVlanSpec{} },

		"VmwareDistributedVirtualSwitchVlanIdSpec": func() interface{} { return &VmwareDistributedVirtualSwitchVlanIdSpec{} },

		"VmwareDistributedVirtualSwitchVlanSpec": func() interface{} { return &VmwareDistributedVirtualSwitchVlanSpec{} },

		"VmwareUplinkPortTeamingPolicy": func() interface{} { return &VmwareUplinkPortTeamingPolicy{} },

		"VnicPortArgument": func() interface{} { return &VnicPortArgument{} },

		"VolumeEditorError": func() interface{} { return &VolumeEditorError{} },

		"VramLimitLicense": func() interface{} { return &VramLimitLicense{} },

		"VsanClusterConfigInfo": func() interface{} { return &VsanClusterConfigInfo{} },

		"VsanClusterConfigInfoHostDefaultInfo": func() interface{} { return &VsanClusterConfigInfoHostDefaultInfo{} },

		"VsanClusterUuidMismatch": func() interface{} { return &VsanClusterUuidMismatch{} },

		"VsanDiskFault": func() interface{} { return &VsanDiskFault{} },

		"VsanDiskIssueType": func() interface{} { return &VsanDiskIssueType{} },

		"VsanFault": func() interface{} { return &VsanFault{} },

		"VsanHostClusterStatus": func() interface{} { return &VsanHostClusterStatus{} },

		"VsanHostClusterStatusState": func() interface{} { return &VsanHostClusterStatusState{} },

		"VsanHostClusterStatusStateCompletionEstimate": func() interface{} { return &VsanHostClusterStatusStateCompletionEstimate{} },

		"VsanHostConfigInfo": func() interface{} { return &VsanHostConfigInfo{} },

		"VsanHostConfigInfoClusterInfo": func() interface{} { return &VsanHostConfigInfoClusterInfo{} },

		"VsanHostConfigInfoNetworkInfo": func() interface{} { return &VsanHostConfigInfoNetworkInfo{} },

		"VsanHostConfigInfoNetworkInfoPortConfig": func() interface{} { return &VsanHostConfigInfoNetworkInfoPortConfig{} },

		"VsanHostConfigInfoStorageInfo": func() interface{} { return &VsanHostConfigInfoStorageInfo{} },

		"VsanHostDecommissionMode": func() interface{} { return &VsanHostDecommissionMode{} },

		"VsanHostDecommissionModeObjectAction": func() interface{} { return &VsanHostDecommissionModeObjectAction{} },

		"VsanHostDiskMapResult": func() interface{} { return &VsanHostDiskMapResult{} },

		"VsanHostDiskMapping": func() interface{} { return &VsanHostDiskMapping{} },

		"VsanHostDiskResult": func() interface{} { return &VsanHostDiskResult{} },

		"VsanHostDiskResultState": func() interface{} { return &VsanHostDiskResultState{} },

		"VsanHostHealthState": func() interface{} { return &VsanHostHealthState{} },

		"VsanHostIpConfig": func() interface{} { return &VsanHostIpConfig{} },

		"VsanHostMembershipInfo": func() interface{} { return &VsanHostMembershipInfo{} },

		"VsanHostNodeState": func() interface{} { return &VsanHostNodeState{} },

		"VsanHostRuntimeInfo": func() interface{} { return &VsanHostRuntimeInfo{} },

		"VsanHostRuntimeInfoDiskIssue": func() interface{} { return &VsanHostRuntimeInfoDiskIssue{} },

		"VspanDestPortConflict": func() interface{} { return &VspanDestPortConflict{} },

		"VspanPortConflict": func() interface{} { return &VspanPortConflict{} },

		"VspanPortMoveFault": func() interface{} { return &VspanPortMoveFault{} },

		"VspanPortPromiscChangeFault": func() interface{} { return &VspanPortPromiscChangeFault{} },

		"VspanPortgroupPromiscChangeFault": func() interface{} { return &VspanPortgroupPromiscChangeFault{} },

		"VspanPortgroupTypeChangeFault": func() interface{} { return &VspanPortgroupTypeChangeFault{} },

		"VspanPromiscuousPortNotSupported": func() interface{} { return &VspanPromiscuousPortNotSupported{} },

		"VspanSameSessionPortConflict": func() interface{} { return &VspanSameSessionPortConflict{} },

		"WaitOptions": func() interface{} { return &WaitOptions{} },

		"WakeOnLanNotSupported": func() interface{} { return &WakeOnLanNotSupported{} },

		"WakeOnLanNotSupportedByVmotionNIC": func() interface{} { return &WakeOnLanNotSupportedByVmotionNIC{} },

		"WarningUpgradeEvent": func() interface{} { return &WarningUpgradeEvent{} },

		"WeekOfMonth": func() interface{} { return &WeekOfMonth{} },

		"WeeklyTaskScheduler": func() interface{} { return &WeeklyTaskScheduler{} },

		"WillLoseHAProtection": func() interface{} { return &WillLoseHAProtection{} },

		"WillLoseHAProtectionResolution": func() interface{} { return &WillLoseHAProtectionResolution{} },

		"WillModifyConfigCpuRequirements": func() interface{} { return &WillModifyConfigCpuRequirements{} },

		"WillResetSnapshotDirectory": func() interface{} { return &WillResetSnapshotDirectory{} },

		"WinNetBIOSConfigInfo": func() interface{} { return &WinNetBIOSConfigInfo{} },

		"WipeDiskFault": func() interface{} { return &WipeDiskFault{} },
	}
}
