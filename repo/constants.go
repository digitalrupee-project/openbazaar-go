package repo

const (
	NotifierTypeChatMessage                   NotificationType = "chatMessage"
	NotifierTypeChatRead                      NotificationType = "chatRead"
	NotifierTypeChatTyping                    NotificationType = "chatTyping"
	NotifierTypeCompletionNotification        NotificationType = "orderComplete"
	NotifierTypeDisputeAcceptedNotification   NotificationType = "disputeAccepted"
	NotifierTypeDisputeCloseNotification      NotificationType = "disputeClose"
	NotifierTypeDisputeOpenNotification       NotificationType = "disputeOpen"
	NotifierTypeDisputeUpdateNotification     NotificationType = "disputeUpdate"
	NotifierTypeFindModeratorResponse         NotificationType = "findModeratorResponse"
	NotifierTypeFollowNotification            NotificationType = "follow"
	NotifierTypeFulfillmentNotification       NotificationType = "fulfillment"
	NotifierTypeIncomingTransaction           NotificationType = "incomingTransaction"
	NotifierTypeModeratorAddNotification      NotificationType = "moderatorAdd"
	NotifierTypeModeratorRemoveNotification   NotificationType = "moderatorRemove"
	NotifierTypeOrderCancelNotification       NotificationType = "cancel"
	NotifierTypeOrderConfirmationNotification NotificationType = "orderConfirmation"
	NotifierTypeOrderDeclinedNotification     NotificationType = "orderDeclined"
	NotifierTypeOrderNewNotification          NotificationType = "order"
	NotifierTypePaymentNotification           NotificationType = "payment"
	NotifierTypePremarshalledNotifier         NotificationType = "premarshalledNotifier"
	NotifierTypeProcessingErrorNotification   NotificationType = "processingError"
	NotifierTypeRefundNotification            NotificationType = "refund"
	NotifierTypeStatusUpdateNotification      NotificationType = "statusUpdate"
	NotifierTypeTestNotification              NotificationType = "testNotification"
	NotifierTypeUnfollowNotification          NotificationType = "unfollow"

	// DisputeAging
	NotifierTypeDisputeAgedZeroDays       NotificationType = "disputeAgedZeroDays"
	NotifierTypeDisputeAgedFifteenDays    NotificationType = "disputeAgedFifteenDays"
	NotifierTypeDisputeAgedFourtyDays     NotificationType = "disputeAgedFourtyDays"
	NotifierTypeDisputeAgedFourtyFourDays NotificationType = "disputeAgedFourtyFourDays"
	NotifierTypeDisputeAgedFourtyFiveDays NotificationType = "disputeAgedFourtyFiveDays"

	// PurchaseAging
	NotifierTypePurchaseAgedZeroDays       NotificationType = "purchaseAgedZeroDays"
	NotifierTypePurchaseAgedFifteenDays    NotificationType = "purchaseAgedFifteenDays"
	NotifierTypePurchaseAgedFourtyDays     NotificationType = "purchaseAgedFourtyDays"
	NotifierTypePurchaseAgedFourtyFourDays NotificationType = "purchaseAgedFourtyFourDays"
	NotifierTypePurchaseAgedFourtyFiveDays NotificationType = "purchaseAgedFourtyFiveDays"
	// End Notification Types
)

type NotificationType string

func (t NotificationType) String() string { return string(t) }
