package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//SecurityStatus msg type = f.
type SecurityStatus struct {
	message.Message
}

//SecurityStatusBuilder builds SecurityStatus messages.
type SecurityStatusBuilder struct {
	message.MessageBuilder
}

//CreateSecurityStatusBuilder returns an initialized SecurityStatusBuilder with specified required fields.
func CreateSecurityStatusBuilder() SecurityStatusBuilder {
	var builder SecurityStatusBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.BuildMsgType("f"))
	return builder
}

//SecurityStatusReqID is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityStatusReqID() (*field.SecurityStatusReqID, errors.MessageRejectError) {
	f := new(field.SecurityStatusReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatusReqID reads a SecurityStatusReqID from SecurityStatus.
func (m SecurityStatus) GetSecurityStatusReqID(f *field.SecurityStatusReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for SecurityStatus.
func (m SecurityStatus) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from SecurityStatus.
func (m SecurityStatus) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for SecurityStatus.
func (m SecurityStatus) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from SecurityStatus.
func (m SecurityStatus) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from SecurityStatus.
func (m SecurityStatus) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from SecurityStatus.
func (m SecurityStatus) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for SecurityStatus.
func (m SecurityStatus) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from SecurityStatus.
func (m SecurityStatus) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for SecurityStatus.
func (m SecurityStatus) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from SecurityStatus.
func (m SecurityStatus) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for SecurityStatus.
func (m SecurityStatus) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from SecurityStatus.
func (m SecurityStatus) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from SecurityStatus.
func (m SecurityStatus) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for SecurityStatus.
func (m SecurityStatus) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from SecurityStatus.
func (m SecurityStatus) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for SecurityStatus.
func (m SecurityStatus) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from SecurityStatus.
func (m SecurityStatus) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for SecurityStatus.
func (m SecurityStatus) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from SecurityStatus.
func (m SecurityStatus) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for SecurityStatus.
func (m SecurityStatus) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from SecurityStatus.
func (m SecurityStatus) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for SecurityStatus.
func (m SecurityStatus) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from SecurityStatus.
func (m SecurityStatus) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for SecurityStatus.
func (m SecurityStatus) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from SecurityStatus.
func (m SecurityStatus) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for SecurityStatus.
func (m SecurityStatus) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from SecurityStatus.
func (m SecurityStatus) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for SecurityStatus.
func (m SecurityStatus) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from SecurityStatus.
func (m SecurityStatus) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for SecurityStatus.
func (m SecurityStatus) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from SecurityStatus.
func (m SecurityStatus) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for SecurityStatus.
func (m SecurityStatus) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from SecurityStatus.
func (m SecurityStatus) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for SecurityStatus.
func (m SecurityStatus) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from SecurityStatus.
func (m SecurityStatus) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for SecurityStatus.
func (m SecurityStatus) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from SecurityStatus.
func (m SecurityStatus) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for SecurityStatus.
func (m SecurityStatus) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from SecurityStatus.
func (m SecurityStatus) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for SecurityStatus.
func (m SecurityStatus) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from SecurityStatus.
func (m SecurityStatus) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for SecurityStatus.
func (m SecurityStatus) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from SecurityStatus.
func (m SecurityStatus) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for SecurityStatus.
func (m SecurityStatus) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from SecurityStatus.
func (m SecurityStatus) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for SecurityStatus.
func (m SecurityStatus) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from SecurityStatus.
func (m SecurityStatus) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for SecurityStatus.
func (m SecurityStatus) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from SecurityStatus.
func (m SecurityStatus) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for SecurityStatus.
func (m SecurityStatus) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from SecurityStatus.
func (m SecurityStatus) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for SecurityStatus.
func (m SecurityStatus) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from SecurityStatus.
func (m SecurityStatus) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from SecurityStatus.
func (m SecurityStatus) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for SecurityStatus.
func (m SecurityStatus) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from SecurityStatus.
func (m SecurityStatus) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from SecurityStatus.
func (m SecurityStatus) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from SecurityStatus.
func (m SecurityStatus) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from SecurityStatus.
func (m SecurityStatus) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from SecurityStatus.
func (m SecurityStatus) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from SecurityStatus.
func (m SecurityStatus) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for SecurityStatus.
func (m SecurityStatus) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from SecurityStatus.
func (m SecurityStatus) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for SecurityStatus.
func (m SecurityStatus) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from SecurityStatus.
func (m SecurityStatus) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for SecurityStatus.
func (m SecurityStatus) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from SecurityStatus.
func (m SecurityStatus) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for SecurityStatus.
func (m SecurityStatus) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from SecurityStatus.
func (m SecurityStatus) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for SecurityStatus.
func (m SecurityStatus) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from SecurityStatus.
func (m SecurityStatus) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for SecurityStatus.
func (m SecurityStatus) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from SecurityStatus.
func (m SecurityStatus) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for SecurityStatus.
func (m SecurityStatus) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from SecurityStatus.
func (m SecurityStatus) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from SecurityStatus.
func (m SecurityStatus) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for SecurityStatus.
func (m SecurityStatus) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from SecurityStatus.
func (m SecurityStatus) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for SecurityStatus.
func (m SecurityStatus) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from SecurityStatus.
func (m SecurityStatus) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for SecurityStatus.
func (m SecurityStatus) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from SecurityStatus.
func (m SecurityStatus) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for SecurityStatus.
func (m SecurityStatus) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from SecurityStatus.
func (m SecurityStatus) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for SecurityStatus.
func (m SecurityStatus) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from SecurityStatus.
func (m SecurityStatus) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for SecurityStatus.
func (m SecurityStatus) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from SecurityStatus.
func (m SecurityStatus) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for SecurityStatus.
func (m SecurityStatus) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from SecurityStatus.
func (m SecurityStatus) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for SecurityStatus.
func (m SecurityStatus) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from SecurityStatus.
func (m SecurityStatus) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for SecurityStatus.
func (m SecurityStatus) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from SecurityStatus.
func (m SecurityStatus) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for SecurityStatus.
func (m SecurityStatus) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from SecurityStatus.
func (m SecurityStatus) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for SecurityStatus.
func (m SecurityStatus) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from SecurityStatus.
func (m SecurityStatus) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from SecurityStatus.
func (m SecurityStatus) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for SecurityStatus.
func (m SecurityStatus) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from SecurityStatus.
func (m SecurityStatus) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for SecurityStatus.
func (m SecurityStatus) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from SecurityStatus.
func (m SecurityStatus) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from SecurityStatus.
func (m SecurityStatus) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from SecurityStatus.
func (m SecurityStatus) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from SecurityStatus.
func (m SecurityStatus) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for SecurityStatus.
func (m SecurityStatus) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from SecurityStatus.
func (m SecurityStatus) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for SecurityStatus.
func (m SecurityStatus) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from SecurityStatus.
func (m SecurityStatus) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for SecurityStatus.
func (m SecurityStatus) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from SecurityStatus.
func (m SecurityStatus) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for SecurityStatus.
func (m SecurityStatus) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from SecurityStatus.
func (m SecurityStatus) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for SecurityStatus.
func (m SecurityStatus) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from SecurityStatus.
func (m SecurityStatus) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for SecurityStatus.
func (m SecurityStatus) OptPayoutAmount() (*field.OptPayoutAmount, errors.MessageRejectError) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from SecurityStatus.
func (m SecurityStatus) GetOptPayoutAmount(f *field.OptPayoutAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for SecurityStatus.
func (m SecurityStatus) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from SecurityStatus.
func (m SecurityStatus) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for SecurityStatus.
func (m SecurityStatus) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from SecurityStatus.
func (m SecurityStatus) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for SecurityStatus.
func (m SecurityStatus) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from SecurityStatus.
func (m SecurityStatus) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for SecurityStatus.
func (m SecurityStatus) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from SecurityStatus.
func (m SecurityStatus) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for SecurityStatus.
func (m SecurityStatus) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from SecurityStatus.
func (m SecurityStatus) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for SecurityStatus.
func (m SecurityStatus) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from SecurityStatus.
func (m SecurityStatus) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for SecurityStatus.
func (m SecurityStatus) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from SecurityStatus.
func (m SecurityStatus) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for SecurityStatus.
func (m SecurityStatus) ValuationMethod() (*field.ValuationMethod, errors.MessageRejectError) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from SecurityStatus.
func (m SecurityStatus) GetValuationMethod(f *field.ValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for SecurityStatus.
func (m SecurityStatus) ContractMultiplierUnit() (*field.ContractMultiplierUnit, errors.MessageRejectError) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from SecurityStatus.
func (m SecurityStatus) GetContractMultiplierUnit(f *field.ContractMultiplierUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for SecurityStatus.
func (m SecurityStatus) FlowScheduleType() (*field.FlowScheduleType, errors.MessageRejectError) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from SecurityStatus.
func (m SecurityStatus) GetFlowScheduleType(f *field.FlowScheduleType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for SecurityStatus.
func (m SecurityStatus) RestructuringType() (*field.RestructuringType, errors.MessageRejectError) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from SecurityStatus.
func (m SecurityStatus) GetRestructuringType(f *field.RestructuringType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for SecurityStatus.
func (m SecurityStatus) Seniority() (*field.Seniority, errors.MessageRejectError) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from SecurityStatus.
func (m SecurityStatus) GetSeniority(f *field.Seniority) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for SecurityStatus.
func (m SecurityStatus) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from SecurityStatus.
func (m SecurityStatus) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for SecurityStatus.
func (m SecurityStatus) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from SecurityStatus.
func (m SecurityStatus) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for SecurityStatus.
func (m SecurityStatus) AttachmentPoint() (*field.AttachmentPoint, errors.MessageRejectError) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from SecurityStatus.
func (m SecurityStatus) GetAttachmentPoint(f *field.AttachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for SecurityStatus.
func (m SecurityStatus) DetachmentPoint() (*field.DetachmentPoint, errors.MessageRejectError) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from SecurityStatus.
func (m SecurityStatus) GetDetachmentPoint(f *field.DetachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for SecurityStatus.
func (m SecurityStatus) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from SecurityStatus.
func (m SecurityStatus) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for SecurityStatus.
func (m SecurityStatus) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from SecurityStatus.
func (m SecurityStatus) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for SecurityStatus.
func (m SecurityStatus) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from SecurityStatus.
func (m SecurityStatus) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecision) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for SecurityStatus.
func (m SecurityStatus) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from SecurityStatus.
func (m SecurityStatus) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for SecurityStatus.
func (m SecurityStatus) OptPayoutType() (*field.OptPayoutType, errors.MessageRejectError) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from SecurityStatus.
func (m SecurityStatus) GetOptPayoutType(f *field.OptPayoutType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for SecurityStatus.
func (m SecurityStatus) NoComplexEvents() (*field.NoComplexEvents, errors.MessageRejectError) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from SecurityStatus.
func (m SecurityStatus) GetNoComplexEvents(f *field.NoComplexEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryForm is a non-required field for SecurityStatus.
func (m SecurityStatus) DeliveryForm() (*field.DeliveryForm, errors.MessageRejectError) {
	f := new(field.DeliveryForm)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryForm reads a DeliveryForm from SecurityStatus.
func (m SecurityStatus) GetDeliveryForm(f *field.DeliveryForm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PctAtRisk is a non-required field for SecurityStatus.
func (m SecurityStatus) PctAtRisk() (*field.PctAtRisk, errors.MessageRejectError) {
	f := new(field.PctAtRisk)
	err := m.Body.Get(f)
	return f, err
}

//GetPctAtRisk reads a PctAtRisk from SecurityStatus.
func (m SecurityStatus) GetPctAtRisk(f *field.PctAtRisk) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrAttrib is a non-required field for SecurityStatus.
func (m SecurityStatus) NoInstrAttrib() (*field.NoInstrAttrib, errors.MessageRejectError) {
	f := new(field.NoInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrAttrib reads a NoInstrAttrib from SecurityStatus.
func (m SecurityStatus) GetNoInstrAttrib(f *field.NoInstrAttrib) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for SecurityStatus.
func (m SecurityStatus) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from SecurityStatus.
func (m SecurityStatus) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for SecurityStatus.
func (m SecurityStatus) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from SecurityStatus.
func (m SecurityStatus) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for SecurityStatus.
func (m SecurityStatus) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from SecurityStatus.
func (m SecurityStatus) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for SecurityStatus.
func (m SecurityStatus) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from SecurityStatus.
func (m SecurityStatus) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for SecurityStatus.
func (m SecurityStatus) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from SecurityStatus.
func (m SecurityStatus) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnsolicitedIndicator is a non-required field for SecurityStatus.
func (m SecurityStatus) UnsolicitedIndicator() (*field.UnsolicitedIndicator, errors.MessageRejectError) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetUnsolicitedIndicator reads a UnsolicitedIndicator from SecurityStatus.
func (m SecurityStatus) GetUnsolicitedIndicator(f *field.UnsolicitedIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityTradingStatus is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityTradingStatus() (*field.SecurityTradingStatus, errors.MessageRejectError) {
	f := new(field.SecurityTradingStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityTradingStatus reads a SecurityTradingStatus from SecurityStatus.
func (m SecurityStatus) GetSecurityTradingStatus(f *field.SecurityTradingStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FinancialStatus is a non-required field for SecurityStatus.
func (m SecurityStatus) FinancialStatus() (*field.FinancialStatus, errors.MessageRejectError) {
	f := new(field.FinancialStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetFinancialStatus reads a FinancialStatus from SecurityStatus.
func (m SecurityStatus) GetFinancialStatus(f *field.FinancialStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CorporateAction is a non-required field for SecurityStatus.
func (m SecurityStatus) CorporateAction() (*field.CorporateAction, errors.MessageRejectError) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}

//GetCorporateAction reads a CorporateAction from SecurityStatus.
func (m SecurityStatus) GetCorporateAction(f *field.CorporateAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HaltReasonInt is a non-required field for SecurityStatus.
func (m SecurityStatus) HaltReasonInt() (*field.HaltReasonInt, errors.MessageRejectError) {
	f := new(field.HaltReasonInt)
	err := m.Body.Get(f)
	return f, err
}

//GetHaltReasonInt reads a HaltReasonInt from SecurityStatus.
func (m SecurityStatus) GetHaltReasonInt(f *field.HaltReasonInt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InViewOfCommon is a non-required field for SecurityStatus.
func (m SecurityStatus) InViewOfCommon() (*field.InViewOfCommon, errors.MessageRejectError) {
	f := new(field.InViewOfCommon)
	err := m.Body.Get(f)
	return f, err
}

//GetInViewOfCommon reads a InViewOfCommon from SecurityStatus.
func (m SecurityStatus) GetInViewOfCommon(f *field.InViewOfCommon) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DueToRelated is a non-required field for SecurityStatus.
func (m SecurityStatus) DueToRelated() (*field.DueToRelated, errors.MessageRejectError) {
	f := new(field.DueToRelated)
	err := m.Body.Get(f)
	return f, err
}

//GetDueToRelated reads a DueToRelated from SecurityStatus.
func (m SecurityStatus) GetDueToRelated(f *field.DueToRelated) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BuyVolume is a non-required field for SecurityStatus.
func (m SecurityStatus) BuyVolume() (*field.BuyVolume, errors.MessageRejectError) {
	f := new(field.BuyVolume)
	err := m.Body.Get(f)
	return f, err
}

//GetBuyVolume reads a BuyVolume from SecurityStatus.
func (m SecurityStatus) GetBuyVolume(f *field.BuyVolume) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SellVolume is a non-required field for SecurityStatus.
func (m SecurityStatus) SellVolume() (*field.SellVolume, errors.MessageRejectError) {
	f := new(field.SellVolume)
	err := m.Body.Get(f)
	return f, err
}

//GetSellVolume reads a SellVolume from SecurityStatus.
func (m SecurityStatus) GetSellVolume(f *field.SellVolume) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HighPx is a non-required field for SecurityStatus.
func (m SecurityStatus) HighPx() (*field.HighPx, errors.MessageRejectError) {
	f := new(field.HighPx)
	err := m.Body.Get(f)
	return f, err
}

//GetHighPx reads a HighPx from SecurityStatus.
func (m SecurityStatus) GetHighPx(f *field.HighPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LowPx is a non-required field for SecurityStatus.
func (m SecurityStatus) LowPx() (*field.LowPx, errors.MessageRejectError) {
	f := new(field.LowPx)
	err := m.Body.Get(f)
	return f, err
}

//GetLowPx reads a LowPx from SecurityStatus.
func (m SecurityStatus) GetLowPx(f *field.LowPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a non-required field for SecurityStatus.
func (m SecurityStatus) LastPx() (*field.LastPx, errors.MessageRejectError) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from SecurityStatus.
func (m SecurityStatus) GetLastPx(f *field.LastPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for SecurityStatus.
func (m SecurityStatus) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SecurityStatus.
func (m SecurityStatus) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Adjustment is a non-required field for SecurityStatus.
func (m SecurityStatus) Adjustment() (*field.Adjustment, errors.MessageRejectError) {
	f := new(field.Adjustment)
	err := m.Body.Get(f)
	return f, err
}

//GetAdjustment reads a Adjustment from SecurityStatus.
func (m SecurityStatus) GetAdjustment(f *field.Adjustment) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SecurityStatus.
func (m SecurityStatus) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SecurityStatus.
func (m SecurityStatus) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SecurityStatus.
func (m SecurityStatus) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SecurityStatus.
func (m SecurityStatus) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SecurityStatus.
func (m SecurityStatus) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FirstPx is a non-required field for SecurityStatus.
func (m SecurityStatus) FirstPx() (*field.FirstPx, errors.MessageRejectError) {
	f := new(field.FirstPx)
	err := m.Body.Get(f)
	return f, err
}

//GetFirstPx reads a FirstPx from SecurityStatus.
func (m SecurityStatus) GetFirstPx(f *field.FirstPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityTradingEvent is a non-required field for SecurityStatus.
func (m SecurityStatus) SecurityTradingEvent() (*field.SecurityTradingEvent, errors.MessageRejectError) {
	f := new(field.SecurityTradingEvent)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityTradingEvent reads a SecurityTradingEvent from SecurityStatus.
func (m SecurityStatus) GetSecurityTradingEvent(f *field.SecurityTradingEvent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDBookType is a non-required field for SecurityStatus.
func (m SecurityStatus) MDBookType() (*field.MDBookType, errors.MessageRejectError) {
	f := new(field.MDBookType)
	err := m.Body.Get(f)
	return f, err
}

//GetMDBookType reads a MDBookType from SecurityStatus.
func (m SecurityStatus) GetMDBookType(f *field.MDBookType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketDepth is a non-required field for SecurityStatus.
func (m SecurityStatus) MarketDepth() (*field.MarketDepth, errors.MessageRejectError) {
	f := new(field.MarketDepth)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketDepth reads a MarketDepth from SecurityStatus.
func (m SecurityStatus) GetMarketDepth(f *field.MarketDepth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for SecurityStatus.
func (m SecurityStatus) MarketID() (*field.MarketID, errors.MessageRejectError) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from SecurityStatus.
func (m SecurityStatus) GetMarketID(f *field.MarketID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for SecurityStatus.
func (m SecurityStatus) MarketSegmentID() (*field.MarketSegmentID, errors.MessageRejectError) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from SecurityStatus.
func (m SecurityStatus) GetMarketSegmentID(f *field.MarketSegmentID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for SecurityStatus.
func (m SecurityStatus) ApplID() (*field.ApplID, errors.MessageRejectError) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from SecurityStatus.
func (m SecurityStatus) GetApplID(f *field.ApplID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for SecurityStatus.
func (m SecurityStatus) ApplSeqNum() (*field.ApplSeqNum, errors.MessageRejectError) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from SecurityStatus.
func (m SecurityStatus) GetApplSeqNum(f *field.ApplSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for SecurityStatus.
func (m SecurityStatus) ApplLastSeqNum() (*field.ApplLastSeqNum, errors.MessageRejectError) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from SecurityStatus.
func (m SecurityStatus) GetApplLastSeqNum(f *field.ApplLastSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for SecurityStatus.
func (m SecurityStatus) ApplResendFlag() (*field.ApplResendFlag, errors.MessageRejectError) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from SecurityStatus.
func (m SecurityStatus) GetApplResendFlag(f *field.ApplResendFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}
