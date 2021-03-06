// Code generated by entc, DO NOT EDIT.

package kernel

func (ic *InstructCreate) SetInstruct(input *Instruct) *InstructCreate {
	ic.SetCurrentUnix(input.CurrentUnix)
	ic.SetUpdatedUnix(input.UpdatedUnix)
	return ic
}

func (ic *InstructCreate) SetInstructWithOptional(input *Instruct) *InstructCreate {
	ic.SetCurrentUnix(input.CurrentUnix)
	ic.SetUpdatedUnix(input.UpdatedUnix)
	return ic
}

func (iuo *InstructUpdateOne) SetInstruct(input *Instruct) *InstructUpdateOne {
	iuo.SetCurrentUnix(input.CurrentUnix)
	iuo.SetUpdatedUnix(input.UpdatedUnix)
	return iuo
}

func (iu *InstructUpdate) SetInstruct(input *Instruct) *InstructUpdate {
	iu.SetCurrentUnix(input.CurrentUnix)
	iu.SetUpdatedUnix(input.UpdatedUnix)
	return iu
}

func (iu *InstructUpdate) SetInstructWithOptional(input *Instruct) *InstructUpdate {
	iu.SetCurrentUnix(input.CurrentUnix)
	iu.SetUpdatedUnix(input.UpdatedUnix)
	return iu
}

func (pc *PinCreate) SetPin(input *Pin) *PinCreate {
	pc.SetRid(input.Rid)
	pc.SetStatus(input.Status)
	pc.SetStep(input.Step)
	pc.SetPriority(input.Priority)
	pc.SetRelate(input.Relate)
	pc.SetUpdatedUnix(input.UpdatedUnix)
	return pc
}

func (pc *PinCreate) SetPinWithOptional(input *Pin) *PinCreate {
	pc.SetRid(input.Rid)
	pc.SetStatus(input.Status)
	pc.SetStep(input.Step)
	pc.SetPriority(input.Priority)
	pc.SetRelate(input.Relate)
	pc.SetUpdatedUnix(input.UpdatedUnix)
	pc.SetComment(input.Comment)
	return pc
}

func (puo *PinUpdateOne) SetPin(input *Pin) *PinUpdateOne {
	puo.SetRid(input.Rid)
	puo.SetStatus(input.Status)
	puo.SetStep(input.Step)
	puo.SetPriority(input.Priority)
	puo.SetRelate(input.Relate)
	puo.SetUpdatedUnix(input.UpdatedUnix)
	return puo
}

func (pu *PinUpdate) SetPin(input *Pin) *PinUpdate {
	pu.SetRid(input.Rid)
	pu.SetStatus(input.Status)
	pu.SetStep(input.Step)
	pu.SetPriority(input.Priority)
	pu.SetRelate(input.Relate)
	pu.SetUpdatedUnix(input.UpdatedUnix)
	return pu
}

func (pu *PinUpdate) SetPinWithOptional(input *Pin) *PinUpdate {
	pu.SetRid(input.Rid)
	pu.SetStatus(input.Status)
	pu.SetStep(input.Step)
	pu.SetPriority(input.Priority)
	pu.SetRelate(input.Relate)
	pu.SetUpdatedUnix(input.UpdatedUnix)
	pu.SetComment(input.Comment)
	return pu
}

func (vc *VersionCreate) SetVersion(input *Version) *VersionCreate {
	vc.SetCurrent(input.Current)
	vc.SetLast(input.Last)
	return vc
}

func (vc *VersionCreate) SetVersionWithOptional(input *Version) *VersionCreate {
	vc.SetCurrent(input.Current)
	vc.SetLast(input.Last)
	return vc
}

func (vuo *VersionUpdateOne) SetVersion(input *Version) *VersionUpdateOne {
	vuo.SetCurrent(input.Current)
	vuo.SetLast(input.Last)
	return vuo
}

func (vu *VersionUpdate) SetVersion(input *Version) *VersionUpdate {
	vu.SetCurrent(input.Current)
	vu.SetLast(input.Last)
	return vu
}

func (vu *VersionUpdate) SetVersionWithOptional(input *Version) *VersionUpdate {
	vu.SetCurrent(input.Current)
	vu.SetLast(input.Last)
	return vu
}
