package client

type operation interface {
	Process(state *albionState)
}

//go:generate stringer -type=OperationType
type OperationType uint16

const (
	opUnused = iota
	opPing
	opJoin
	opCreateAccount
	opLogin
	opSendCrashLog
	opCreateCharacter
	opDeleteCharacter
	opSelectCharacter
	opRedeemKeycode
	opGetGameServerByCluster
	opGetSubscriptionDetails
	opGetActiveSubscription
	opGetSubscriptionUrl
	opGetPurchaseGoldUrl
	opGetBuyTrialDetails
	opGetReferralSeasonDetails
	opGetAvailableTrialKeys
	opMove
	opAttackStart
	opCastStart
	opCastCancel
	opTerminateToggleSpell
	opChannelingCancel
	opAttackBuildingStart
	opInventoryDestroyItem
	opInventoryMoveItem
	opInventorySplitStack
	opChangeCluster
	opConsoleCommand
	opChatMessage
	opReportClientError
	opRegisterToObject
	opUnRegisterFromObject
	opCraftBuildingChangeSettings
	opCraftBuildingTakeMoney
	opRepairBuildingChangeSettings
	opRepairBuildingTakeMoney
	opActionBuildingChangeSettings
	opHarvestStart
	opHarvestCancel
	opTakeSilver
	opActionOnBuildingStart
	opActionOnBuildingCancel
	opItemRerollQualityStart
	opItemRerollQualityCancel
	opInstallResourceStart
	opInstallResourceCancel
	opInstallSilver
	opBuildingFillNutrition
	opBuildingChangeRenovationState
	opBuildingBuySkin
	opBuildingClaim
	opBuildingGiveup
	opBuildingNutritionSilverStorageDeposit
	opBuildingNutritionSilverStorageWithdraw
	opBuildingNutritionSilverRewardSet
	opConstructionSiteCreate
	opPlaceableItemPlace
	opPlaceableItemPlaceCancel
	opPlaceableObjectPickup
	opFurnitureObjectUse
	opFarmableHarvest
	opFarmableFinishGrownItem
	opFarmableDestroy
	opFarmableGetProduct
	opFarmableFill
	opLaborerObjectPlace
	opLaborerObjectPlaceCancel
	opTearDownConstructionSite
	opCastleGateUse
	opAuctionCreateOffer
	opAuctionCreateRequest
	opAuctionGetOffers
	opAuctionGetRequests
	opAuctionBuyOffer
	opAuctionAbortAuction
	opAuctionModifyAuction
	opAuctionAbortOffer
	opAuctionAbortRequest
	opAuctionSellRequest
	opAuctionGetFinishedAuctions
	opAuctionFetchAuction
	opAuctionGetMyOpenOffers
	opAuctionGetMyOpenRequests
	opAuctionGetMyOpenAuctions
	opAuctionGetItemsAverage
	opAuctionGetItemAverageStats
	opAuctionGetItemAverageValue
	opContainerOpen
	opContainerClose
	opContainerManageSubContainer
	opRespawn
	opSuicide
	opJoinGuild
	opLeaveGuild
	opCreateGuild
	opInviteToGuild
	opDeclineGuildInvitation
	opKickFromGuild
	opDuellingChallengePlayer
	opDuellingAcceptChallenge
	opDuellingDenyChallenge
	opChangeClusterTax
	opClaimTerritory
	opGiveUpTerritory
	opChangeTerritoryAccessRights
	opGetMonolithInfo
	opGetClaimInfo
	opGetAttackInfo
	opGetTerritorySeasonPoints
	opGetAttackSchedule
	opScheduleAttack
	opGetMatches
	opGetMatchDetails
	opJoinMatch
	opLeaveMatch
	opChangeChatSettings
	opLogoutStart
	opLogoutCancel
	opClaimOrbStart
	opClaimOrbCancel
	opMatchLootChestOpeningStart
	opMatchLootChestOpeningCancel
	opDepositToGuildAccount
	opWithdrawalFromAccount
	opChangeGuildPayUpkeepFlag
	opChangeGuildTax
	opGetMyTerritories
	opMorganaCommand
	opGetServerInfo
	opInviteMercenaryToMatch
	opSubscribeToCluster
	opAnswerMercenaryInvitation
	opGetCharacterEquipment
	opGetCharacterSteamAchievements
	opGetCharacterStats
	opGetKillHistoryDetails
	opLearnMasteryLevel
	opReSpecAchievement
	opChangeAvatar
	opGetRankings
	opGetRank
	opGetGvgSeasonRankings
	opGetGvgSeasonRank
	opGetGvgSeasonHistoryRankings
	opGetGvgSeasonGuildMemberHistory
	opKickFromGvGMatch
	opGetChestLogs
	opGetAccessRightLogs
	opGetGuildAccountLogs
	opInviteToPlayerTrade
	opPlayerTradeCancel
	opPlayerTradeInvitationAccept
	opPlayerTradeAddItem
	opPlayerTradeRemoveItem
	opPlayerTradeAcceptTrade
	opPlayerTradeSetSilverOrGold
	opSendMiniMapPing
	opStuck
	opBuyRealEstate
	opClaimRealEstate
	opGiveUpRealEstate
	opChangeRealEstateOutline
	opGetMailInfos
	opReadMail
	opSendNewMail
	opDeleteMail
	opClaimAttachmentFromMail
	opUpdateLfgInfo
	opGetLfgInfos
	opGetMyGuildLfgInfo
	opGetLfgDescriptionText
	opLfgApplyToGuild
	opAnswerLfgGuildApplication
	opGetClusterInfo
	opRegisterChatPeer
	opSendChatMessage
	opJoinChatChannel
	opLeaveChatChannel
	opSendWhisperMessage
	opSay
	opPlayEmote
	opStopEmote
	opGetClusterMapInfo
	opAccessRightsChangeSettings
	opMount
	opMountCancel
	opBuyJourney
	opSetSaleStatusForEstate
	opResolveGuildOrPlayerName
	opGetRespawnInfos
	opMakeHome
	opLeaveHome
	opResurrectionReply
	opAllianceCreate
	opAllianceDisband
	opAllianceGetMemberInfos
	opAllianceInvite
	opAllianceAnswerInvitation
	opAllianceCancelInvitation
	opAllianceKickGuild
	opAllianceLeave
	opAllianceChangeGoldPaymentFlag
	opAllianceGetDetailInfo
	opGetIslandInfos
	opAbandonMyIsland
	opBuyMyIsland
	opBuyGuildIsland
	opAbandonGuildIsland
	opUpgradeMyIsland
	opUpgradeGuildIsland
	opTerritoryFillNutrition
	opTeleportBack
	opPartyInvitePlayer
	opPartyAnswerInvitation
	opPartyLeave
	opPartyKickPlayer
	opPartyMakeLeader
	opPartyChangeLootSetting
	opPartyMarkObject
	opPartySetRole
	opGetGuildMOTD
	opSetGuildMOTD
	opExitEnterStart
	opExitEnterCancel
	opQuestGiverRequest
	opGoldMarketGetBuyOffer
	opGoldMarketGetBuyOfferFromSilver
	opGoldMarketGetSellOffer
	opGoldMarketGetSellOfferFromSilver
	opGoldMarketBuyGold
	opGoldMarketSellGold
	opGoldMarketCreateSellOrder
	opGoldMarketCreateBuyOrder
	opGoldMarketGetInfos
	opGoldMarketCancelOrder
	opGoldMarketGetAverageInfo
	opSiegeCampClaimStart
	opSiegeCampClaimCancel
	opTreasureChestUsingStart
	opTreasureChestUsingCancel
	opLaborerStartJob
	opLaborerTakeJobLoot
	opLaborerDismiss
	opLaborerMove
	opLaborerBuyItem
	opLaborerUpgrade
	opBuyPremium
	opBuyTrial
	opRealEstateGetAuctionData
	opRealEstateBidOnAuction
	opGetSiegeCampCooldown
	opFriendInvite
	opFriendAnswerInvitation
	opFriendCancelnvitation
	opFriendRemove
	opInventoryStack
	opInventorySort
	opEquipmentItemChangeSpell
	opExpeditionRegister
	opExpeditionRegisterCancel
	opJoinExpedition
	opDeclineExpeditionInvitation
	opVoteStart
	opVoteDoVote
	opRatingDoRate
	opEnteringExpeditionStart
	opEnteringExpeditionCancel
	opActivateExpeditionCheckPoint
	opArenaRegister
	opArenaRegisterCancel
	opArenaLeave
	opJoinArenaMatch
	opDeclineArenaInvitation
	opEnteringArenaStart
	opEnteringArenaCancel
	opArenaCustomMatch
	opArenaCustomMatchCreate
	opUpdateCharacterStatement
	opBoostFarmable
	opGetStrikeHistory
	opUseFunction
	opUsePortalEntrance
	opQueryPortalBinding
	opClaimPaymentTransaction
	opChangeUseFlag
	opClientPerformanceStats
	opExtendedHardwareStats
	opTerritoryClaimStart
	opTerritoryClaimCancel
	opRequestAppStoreProducts
	opVerifyProductPurchase
	opQueryGuildPlayerStats
	opTrackAchievements
	opSetAchievementsAutoLearn
	opDepositItemToGuildCurrency
	opWithdrawalItemFromGuildCurrency
	opAuctionSellSpecificItemRequest
	opFishingStart
	opFishingCasting
	opFishingCast
	opFishingCatch
	opFishingPull
	opFishingGiveLine
	opFishingFinish
	opFishingCancel
	opCreateGuildAccessTag
	opDeleteGuildAccessTag
	opRenameGuildAccessTag
	opFlagGuildAccessTagGuildPermission
	opAssignGuildAccessTag
	opRemoveGuildAccessTagFromPlayer
	opModifyGuildAccessTagEditors
	opRequestPublicAccessTags
	opChangeAccessTagPublicFlag
	opUpdateGuildAccessTag
	opSteamStartMicrotransaction
	opSteamFinishMicrotransaction
	opSteamIdHasActiveAccount
	opCheckEmailAccountState
	opLinkAccountToSteamId
	opBuyGvgSeasonBooster
	opChangeFlaggingPrepare
	opOverCharge
	opOverChargeEnd
	opRequestTrusted
	opChangeGuildLogo
	opPartyFinderRegisterForUpdates
	opPartyFinderUnregisterForUpdates
	opPartyFinderEnlistNewPartySearch
	opPartyFinderDeletePartySearch
	opPartyFinderChangePartySearch
	opPartyFinderChangeRole
	opPartyFinderApplyForGroup
	opPartyFinderAcceptOrDeclineApplyForGroup
	opPartyFinderGetEquipmentSnapshot
	opPartyFinderRegisterApplicants
	opPartyFinderUnregisterApplicants
	opPartyFinderFulltextSearch
	opGetPersonalSeasonTrackerData
	opUseConsumableFromInventory
	opClaimPersonalSeasonReward
	opEasyAntiCheatMessageToServer
	opRetaliationAttackClaimTerritory
)
