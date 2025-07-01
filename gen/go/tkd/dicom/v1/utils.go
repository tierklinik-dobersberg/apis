package dicomv1

import (
	"fmt"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

func ProtoVR(t tag.Tag, vr string) ValueRepresentation {
	if t == tag.Item {
		return ValueRepresentation_VR_ITEM
	}

	if t == tag.PixelData {
		return ValueRepresentation_VR_PIXEL_DATA
	}

	switch vr {
	case "DA":
		return ValueRepresentation_VR_DATE
	case "AT":
		return ValueRepresentation_VR_TAG_LIST
	case "OW", "OB":
		return ValueRepresentation_VR_BYTES
	case "LT", "UT":
		return ValueRepresentation_VR_STRING
	case "UL":
		return ValueRepresentation_VR_UINT32_LIST
	case "SL":
		return ValueRepresentation_VR_INT32_LIST
	case "US":
		return ValueRepresentation_VR_UINT16_LIST
	case "SS":
		return ValueRepresentation_VR_INT16_LIST
	case "FL":
		return ValueRepresentation_VR_FLOAT32_LIST
	case "FD":
		return ValueRepresentation_VR_FLOAT64_LIST
	case "SQ":
		return ValueRepresentation_VR_SEQUENCE
	case "UN":
		return ValueRepresentation_VR_UNSPECIFIED
	default:
		return ValueRepresentation_VR_STRING_LIST
	}
}

func ValueRepresentationProto(vr tag.VRKind) ValueRepresentation {
	switch vr {
	case tag.VRBytes:
		return ValueRepresentation_VR_BYTES
	case tag.VRDate:
		return ValueRepresentation_VR_DATE
	case tag.VRFloat32List:
		return ValueRepresentation_VR_FLOAT32_LIST
	case tag.VRFloat64List:
		return ValueRepresentation_VR_FLOAT64_LIST
	case tag.VRInt16List:
		return ValueRepresentation_VR_INT16_LIST
	case tag.VRInt32List:
		return ValueRepresentation_VR_INT32_LIST
	case tag.VRItem:
		return ValueRepresentation_VR_ITEM
	case tag.VRPixelData:
		return ValueRepresentation_VR_PIXEL_DATA
	case tag.VRSequence:
		return ValueRepresentation_VR_SEQUENCE
	case tag.VRString:
		return ValueRepresentation_VR_STRING
	case tag.VRStringList:
		return ValueRepresentation_VR_STRING_LIST
	case tag.VRTagList:
		return ValueRepresentation_VR_TAG_LIST
	case tag.VRUInt16List:
		return ValueRepresentation_VR_UINT16_LIST
	case tag.VRUInt32List:
		return ValueRepresentation_VR_UINT32_LIST
	default:
		return ValueRepresentation_VR_UNSPECIFIED
	}
}

func ElementProto(el *dicom.Element) (*Element, error) {
	info, err := tag.Find(el.Tag)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup tag info: %w", err)
	}

	pb := &Element{
		Tag:            el.Tag.Uint32(),
		TagName:        info.Keyword,
		TagDisplayName: info.Name,
		Vr:             ValueRepresentationProto(el.ValueRepresentation),
		RawVr:          el.RawValueRepresentation,
	}

	var value isValue_Value
	switch v := el.Value.GetValue().(type) {
	case []string:
		value = &Value_Strings{
			Strings: &StringList{
				Values: v,
			},
		}
	case []byte:
		value = &Value_Bytes{
			Bytes: v,
		}

	case []int:
		elements := make([]int32, 0, len(v))
		for _, i := range v {
			elements = append(elements, int32(i))
		}

		value = &Value_Ints{
			Ints: &IntList{
				Values: elements,
			},
		}

	case dicom.PixelDataInfo:

	case []*dicom.Element:
		elements := make([]*Element, 0, len(v))
		for _, e := range v {
			epb, err := ElementProto(e)
			if err != nil {
				return nil, err
			}

			elements = append(elements, epb)
		}

		value = &Value_SequenceItem{
			SequenceItem: &SequenceItem{
				Values: elements,
			},
		}

	case []dicom.SequenceItemValue:
		elements := make([]*SequenceItem, 0, len(v))
		for _, s := range v {
			var el []*Element

			for _, e := range s.GetValue().([]*dicom.Element) {
				epb, err := ElementProto(e)
				if err != nil {
					return nil, err
				}

				el = append(el, epb)
			}

			elements = append(elements, &SequenceItem{
				Values: el,
			})
		}

		value = &Value_Sequences{
			Sequences: &Sequences{
				Values: elements,
			},
		}

	case []float64:
		value = &Value_Floats{
			Floats: &FloatList{
				Values: v,
			},
		}

	default:
	}

	pb.Value = &Value{
		Value: value,
	}

	return pb, nil
}
