package taapi

// Indicator represents a technical analysis indicator
type Indicator string

const (
	IndicatorRSI         Indicator = "rsi"
	IndicatorMACD        Indicator = "macd"
	IndicatorEMA         Indicator = "ema"
	IndicatorSMA         Indicator = "sma"
	IndicatorBBANDS      Indicator = "bbands"
	IndicatorSTOCH       Indicator = "stoch"
	IndicatorSTOCHRSI    Indicator = "stochrsi"
	IndicatorATR         Indicator = "atr"
	IndicatorADX         Indicator = "adx"
	IndicatorCCI         Indicator = "cci"
	IndicatorAROON       Indicator = "aroon"
	IndicatorMFI         Indicator = "mfi"
	IndicatorOBV         Indicator = "obv"
	IndicatorSAR         Indicator = "sar"
	IndicatorSUPERTREND  Indicator = "supertrend"
	IndicatorICHIMOKU    Indicator = "ichimoku"
	IndicatorVWAP        Indicator = "vwap"
	IndicatorHMA         Indicator = "hma"
	IndicatorWMA         Indicator = "wma"
	IndicatorDEMA        Indicator = "dema"
	IndicatorTEMA        Indicator = "tema"
	IndicatorWILLIAMS    Indicator = "williams"
	IndicatorUO          Indicator = "uo"
	IndicatorROC         Indicator = "roc"
	IndicatorBBP         Indicator = "bbp"
	IndicatorAO          Indicator = "ao"
	IndicatorCMF         Indicator = "cmf"
	IndicatorKELTNER     Indicator = "keltner"
	IndicatorDONCHIAN    Indicator = "donchian"
	IndicatorPIVOT       Indicator = "pivot"
	IndicatorFIBONACCI   Indicator = "fibonacci"
	IndicatorVOLUME      Indicator = "volume"
	IndicatorCANDLE      Indicator = "candle"
)

// String returns the string representation of the indicator
func (i Indicator) String() string {
	return string(i)
}

// IsValid checks if the indicator is valid
func (i Indicator) IsValid() bool {
	switch i {
	case IndicatorRSI, IndicatorMACD, IndicatorEMA, IndicatorSMA,
		IndicatorBBANDS, IndicatorSTOCH, IndicatorSTOCHRSI, IndicatorATR,
		IndicatorADX, IndicatorCCI, IndicatorAROON, IndicatorMFI,
		IndicatorOBV, IndicatorSAR, IndicatorSUPERTREND, IndicatorICHIMOKU,
		IndicatorVWAP, IndicatorHMA, IndicatorWMA, IndicatorDEMA,
		IndicatorTEMA, IndicatorWILLIAMS, IndicatorUO, IndicatorROC,
		IndicatorBBP, IndicatorAO, IndicatorCMF, IndicatorKELTNER,
		IndicatorDONCHIAN, IndicatorPIVOT, IndicatorFIBONACCI, IndicatorVOLUME,
		IndicatorCANDLE:
		return true
	}
	return false
}
