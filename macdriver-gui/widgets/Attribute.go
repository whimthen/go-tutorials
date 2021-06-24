package widgets

type NSLayoutConstraintAttribute uint
type NSLayoutConstraintRelation uint

const (
	NSLayoutConstraintAttributeNotAnAttribute NSLayoutConstraintAttribute = iota
	NSLayoutConstraintAttributeLeft
	NSLayoutConstraintAttributeRight
	NSLayoutConstraintAttributeTop
	NSLayoutConstraintAttributeBottom
	NSLayoutConstraintAttributeLeading
	NSLayoutConstraintAttributeTrailing
	NSLayoutConstraintAttributeWidth
	NSLayoutConstraintAttributeHeight
	NSLayoutConstraintAttributeCenterX
	NSLayoutConstraintAttributeCenterY
	NSLayoutConstraintAttributeLastBaseline
	NSLayoutConstraintAttributeFirstBaseline
	NSLayoutConstraintAttributeLeftMargin
	NSLayoutConstraintAttributeRightMargin
	NSLayoutConstraintAttributeTopMargin
	NSLayoutConstraintAttributeBottomMargin
	NSLayoutConstraintAttributeLeadingMargin
	NSLayoutConstraintAttributeTrailingMargin
	NSLayoutConstraintAttributeCenterXWithinMargins
	NSLayoutConstraintAttributeCenterYWithinMargins
)

const (
	//NSLayoutConstraintRelationLessThanOrEqual NSLayoutConstraintRelation = iota - 1
	NSLayoutConstraintRelationLessEqual              = 0
	NSLayoutConstraintRelationLessGreaterThanOrEqual = 1
)
