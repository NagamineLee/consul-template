// generated by stringer -type=typeCode; DO NOT EDIT

package protocol

import "fmt"

const (
	_typeCode_name_0 = "tcNulltcTinyinttcSmallinttcInttcBiginttcDecimaltcRealtcDoubletcChartcVarchartcNchartcNvarchartcBinarytcVarbinarytcDatetcTimetcTimestamp"
	_typeCode_name_1 = "tcClobtcNclobtcBlobtcBooleantcStringtcNstringtcBlocatortcNlocatortcBstring"
	_typeCode_name_2 = "tcVarchar2tcVarchar3tcNvarchar3tcVarbinary3"
	_typeCode_name_3 = "tcSmalldecimal"
	_typeCode_name_4 = "tcTexttcShorttexttcBintext"
	_typeCode_name_5 = "tcAlphanum"
	_typeCode_name_6 = "tcLongdatetcSeconddatetcDaydatetcSecondtime"
	_typeCode_name_7 = "tcGeometrytcPoint"
	_typeCode_name_8 = "tcPointz"
)

var (
	_typeCode_index_0 = [...]uint8{0, 6, 15, 25, 30, 38, 47, 53, 61, 67, 76, 83, 93, 101, 112, 118, 124, 135}
	_typeCode_index_1 = [...]uint8{0, 6, 13, 19, 28, 36, 45, 55, 65, 74}
	_typeCode_index_2 = [...]uint8{0, 10, 20, 31, 43}
	_typeCode_index_3 = [...]uint8{0, 14}
	_typeCode_index_4 = [...]uint8{0, 6, 17, 26}
	_typeCode_index_5 = [...]uint8{0, 10}
	_typeCode_index_6 = [...]uint8{0, 10, 22, 31, 43}
	_typeCode_index_7 = [...]uint8{0, 10, 17}
	_typeCode_index_8 = [...]uint8{0, 8}
)

func (i typeCode) String() string {
	switch {
	case 0 <= i && i <= 16:
		return _typeCode_name_0[_typeCode_index_0[i]:_typeCode_index_0[i+1]]
	case 25 <= i && i <= 33:
		i -= 25
		return _typeCode_name_1[_typeCode_index_1[i]:_typeCode_index_1[i+1]]
	case 35 <= i && i <= 38:
		i -= 35
		return _typeCode_name_2[_typeCode_index_2[i]:_typeCode_index_2[i+1]]
	case i == 47:
		return _typeCode_name_3
	case 51 <= i && i <= 53:
		i -= 51
		return _typeCode_name_4[_typeCode_index_4[i]:_typeCode_index_4[i+1]]
	case i == 55:
		return _typeCode_name_5
	case 61 <= i && i <= 64:
		i -= 61
		return _typeCode_name_6[_typeCode_index_6[i]:_typeCode_index_6[i+1]]
	case 74 <= i && i <= 75:
		i -= 74
		return _typeCode_name_7[_typeCode_index_7[i]:_typeCode_index_7[i+1]]
	case i == 80:
		return _typeCode_name_8
	default:
		return fmt.Sprintf("typeCode(%d)", i)
	}
}
