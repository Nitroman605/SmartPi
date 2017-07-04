package smartpi
// Register addresses},

var ADE7878REG = map[string][]byte{
	"AIGAIN": []byte{0x43, 0x80},
	"AVGAIN": []byte{0x43, 0x81},
	"BIGAIN": []byte{0x43, 0x82},
	"BVGAIN": []byte{0x43, 0x83},
	"CIGAIN": []byte{0x43, 0x84},
	"CVGAIN": []byte{0x43, 0x85},
	"NIGAIN": []byte{0x43, 0x86},
	"AIRMSOS": []byte{0x43, 0x87},
	"AVRMSOS": []byte{0x43, 0x88},
	"BIRMSOS": []byte{0x43, 0x89},
	"BVRMSOS": []byte{0x43, 0x8A},
	"CIRMSOS": []byte{0x43, 0x8B},
	"CVRMSOS": []byte{0x43, 0x8C},
	"NIRMSOS": []byte{0x43, 0x8D},
	"AVAGAIN": []byte{0x43, 0x8E},
	"BVAGAIN": []byte{0x43, 0x8F},
	"CVAGAIN": []byte{0x43, 0x90},
	"AWGAIN": []byte{0x43, 0x91},
	"AWATTOS": []byte{0x43, 0x92},
	"BWGAIN": []byte{0x43, 0x93},
	"BWATTOS": []byte{0x43, 0x94},
	"CWGAIN": []byte{0x43, 0x95},
	"CWATTOS": []byte{0x43, 0x96},
	"AVARGAIN": []byte{0x43, 0x97},
	"AVAROS": []byte{0x43, 0x98},
	"BVARGAIN": []byte{0x43, 0x99},
	"BVAROS": []byte{0x43, 0x9A},
	"CVARGAIN": []byte{0x43, 0x9B},
	"CVAROS": []byte{0x43, 0x9C},
	"AFWGAIN": []byte{0x43, 0x9D},
	"AFWATTOS": []byte{0x43, 0x9E},
	"BFWGAIN": []byte{0x43, 0x9F},
	"BFWATTOS": []byte{0x43, 0xA0},
	"CFWGAIN": []byte{0x43, 0xA1},
	"CFWATTOS": []byte{0x43, 0xA3},
	"AFVARGAIN": []byte{0x43, 0x9E},
	"AFVAROS": []byte{0x43, 0xA4},
	"BFVARGAIN": []byte{0x43, 0xA5},
	"BFVAROS": []byte{0x43, 0xA6},
	"CFVARGAIN": []byte{0x43, 0xA7},
	"CFVAROS": []byte{0x43, 0xA8},
	"VATHR1": []byte{0x43, 0xA9},
	"VATHR0": []byte{0x43, 0xAA},
	"WTHR1": []byte{0x43, 0xAB},
	"WTHR0": []byte{0x43, 0xAC},
	"VARTHR1": []byte{0x43, 0xAD},
	"VARTHR0": []byte{0x43, 0xAE},
	"VANOLOAD": []byte{0x43, 0xB0},
	"APNOLOAD": []byte{0x43, 0xB1},
	"VARNOLOAD": []byte{0x43, 0xB2},
	"VLEVEL": []byte{0x43, 0xB3},
	"DICOEFF": []byte{0x43, 0xB5},
	"HPFDIS": []byte{0x43, 0xB6},
	"ISUM": []byte{0x43, 0xBF},
	"AIRMS": []byte{0x43, 0xC0},
	"AVRMS": []byte{0x43, 0xC1},
	"BIRMS": []byte{0x43, 0xC2},
	"BVRMS": []byte{0x43, 0xC3},
	"CIRMS": []byte{0x43, 0xC4},
	"CVRMS": []byte{0x43, 0xC5},
	"NIRMS": []byte{0x43, 0xC6},
	"RUN": []byte{	0xE2, 0x28},
	"AWATTHR": []byte{0xE4, 0x00},
	"BWATTHR": []byte{0xE4, 0x01},
	"CWATTHR": []byte{0xE4, 0x02},
	"AFWATTHR": []byte{0xE4, 0x03},
	"BFWATTHR": []byte{0xE4, 0x04},
	"CFWATTHR": []byte{0xE4, 0x05},
	"AVARHR": []byte{0xE4, 0x06},
	"BVARHR": []byte{0xE4, 0x07},
	"CVARHR": []byte{0xE4, 0x08},
	"AFVARHR": []byte{0xE4, 0x09},
	"BFVARHR": []byte{0xE4, 0x0A},
	"CFVARHR": []byte{0xE4, 0x0B},
	"AVAHR": []byte{0xE4, 0x0C},
	"BVAHR": []byte{0xE4, 0x0D},
	"CVAHR": []byte{0xE4, 0x0E},
	"IPEAK": []byte{0xE5, 0x00},
	"VPEAK": []byte{0xE5, 0x01},
	"STATUS0": []byte{0xE5, 0x02},
	"STATUS1": []byte{0xE5, 0x03},
	"AIMAV": []byte{0xE5, 0x04},
	"BIMAV": []byte{0xE5, 0x05},
	"CIMAV": []byte{0xE5, 0x06},
	"OILVL": []byte{0xE5, 0x07},
	"OVLVL": []byte{0xE5, 0x08},
	"SAGLVL": []byte{0xE5, 0x09},
	"MASK0": []byte{0xE5, 0x0A},
	"MASK1": []byte{0xE5, 0x0B},
	"IAWV": []byte{0xE5, 0x0C},
	"IBWV": []byte{0xE5, 0x0D},
	"ICWV": []byte{0xE5, 0x0E},
	"INWV": []byte{0xE5, 0x0F},
	"VAWV": []byte{0xE5, 0x10},
	"VBWV": []byte{0xE5, 0x11},
	"VCWV": []byte{0xE5, 0x12},
	"AWATT": []byte{0xE5, 0x13},
	"BWATT": []byte{0xE5, 0x14},
	"CWATT": []byte{0xE5, 0x15},
	"AVAR": []byte{0xE5, 0x16},
	"BVAR": []byte{0xE5, 0x17},
	"CVAR": []byte{0xE5, 0x18},
	"AVA": []byte{	0xE5, 0x19},
	"BVA": []byte{	0xE5, 0x1A},
	"CVA": []byte{	0xE5, 0x1B},
	"CHECKSUM": []byte{0xE5, 0x1F},
	"VNOM": []byte{0xE5, 0x20},
	"PHSTATUS": []byte{0xE6, 0x00},
	"ANGLE0": []byte{0xE6, 0x01},
	"ANGLE1": []byte{0xE6, 0x02},
	"ANGLE2": []byte{0xE6, 0x03},
	"PERIOD": []byte{0xE6, 0x07},
	"PHNOLOAD": []byte{0xE6, 0x08},
	"LINECYC": []byte{0xE6, 0x0C},
	"ZXTOUT": []byte{0xE6, 0x0D},
	"COMPMODE": []byte{0xE6, 0x0E},
	"GAIN": []byte{0xE6, 0x0F},
	"CFMODE": []byte{0xE6, 0x10},
	"CF1DEN": []byte{0xE6, 0x11},
	"CF2DEN": []byte{0xE6, 0x12},
	"CF3DEN": []byte{0xE6, 0x13},
	"APHCAL": []byte{0xE6, 0x14},
	"BPHCAL": []byte{0xE6, 0x15},
	"CPHCAL": []byte{0xE6, 0x16},
	"PHSIGN": []byte{0xE6, 0x17},
	"CONFIG": []byte{0xE6, 0x18},
	"MMODE": []byte{0xE7, 0x00},
	"ACCMODE": []byte{0xE7, 0x01},
	"LCYCMODE": []byte{0xE7, 0x02},
	"PEAKCYC": []byte{0xE7, 0x03},
	"SAGCYC": []byte{0xE7, 0x04},
	"CFCYC": []byte{0xE7, 0x05},
	"HSDC_CFG": []byte{0xE7, 0x06},
	"VERSION": []byte{0xE7, 0x07},
	"LPOILVL": []byte{0xEC, 0x00},
	"CONFIG2": []byte{0xEC, 0x01},
}